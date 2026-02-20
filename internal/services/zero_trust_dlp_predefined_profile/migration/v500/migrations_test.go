package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion
)

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v5_multiple_entries.tf
var v5MultipleEntriesConfig string

//go:embed testdata/v5_with_ocr.tf
var v5WithOCRConfig string

//go:embed testdata/v5_16_basic.tf
var v516BasicConfig string

func skipIfMissingPredefinedIDs(t *testing.T) (profileID, entryID string) {
	t.Helper()
	profileID = os.Getenv("CLOUDFLARE_DLP_PREDEFINED_PROFILE_ID")
	entryID = os.Getenv("CLOUDFLARE_DLP_PREDEFINED_ENTRY_ID")
	if profileID == "" || entryID == "" {
		t.Skip("CLOUDFLARE_DLP_PREDEFINED_PROFILE_ID and CLOUDFLARE_DLP_PREDEFINED_ENTRY_ID must be set")
	}
	return profileID, entryID
}

func skipIfMissingMultipleEntryIDs(t *testing.T) (profileID, entryID1, entryID2, entryID3 string) {
	t.Helper()
	profileID = os.Getenv("CLOUDFLARE_DLP_PREDEFINED_PROFILE_ID")
	entryID1 = os.Getenv("CLOUDFLARE_DLP_PREDEFINED_ENTRY_ID")
	entryID2 = os.Getenv("CLOUDFLARE_DLP_PREDEFINED_ENTRY_ID_2")
	entryID3 = os.Getenv("CLOUDFLARE_DLP_PREDEFINED_ENTRY_ID_3")
	if profileID == "" || entryID1 == "" || entryID2 == "" || entryID3 == "" {
		t.Skip("CLOUDFLARE_DLP_PREDEFINED_PROFILE_ID, CLOUDFLARE_DLP_PREDEFINED_ENTRY_ID, _ENTRY_ID_2, _ENTRY_ID_3 must be set")
	}
	return profileID, entryID1, entryID2, entryID3
}

// TestMigrateZeroTrustDLPPredefinedProfile_V5_Basic tests migration of a predefined DLP profile.
// Note: from_v4_latest is omitted because predefined profiles cannot be created via
// the v4 provider (must be imported) and the test framework's ImportState doesn't
// properly resolve ExternalProviders source addresses. The v4→v5 Transform logic
// is exercised by the from_v5 path (same Transform function) and tf-migrate integration tests.
func TestMigrateZeroTrustDLPPredefinedProfile_V5_Basic(t *testing.T) {
	profileID, entryID := skipIfMissingPredefinedIDs(t)

	testCases := []struct {
		name       string
		version    string
		disableMig bool
		configFn   func(rnd, accountID string) string
	}{
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, profileID, entryID)
			},
		},
		{
			name:       "from_v5_16",
			version:    "5.16.0",
			disableMig: true,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v516BasicConfig, rnd, accountID, profileID, entryID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.disableMig {
				os.Setenv("TF_MIG_TEST", "")
				t.Cleanup(func() { os.Setenv("TF_MIG_TEST", "1") })
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_dlp_predefined_profile."+rnd,
							tfjsonpath.New("account_id"),
							knownvalue.StringExact(accountID),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_dlp_predefined_profile."+rnd,
							tfjsonpath.New("profile_id"),
							knownvalue.StringExact(profileID),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_dlp_predefined_profile."+rnd,
							tfjsonpath.New("allowed_match_count"),
							knownvalue.Int64Exact(3),
						),
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDLPPredefinedProfile_V4ToV5_MultipleEntries tests migration with
// multiple entries, some enabled and some disabled.
func TestMigrateZeroTrustDLPPredefinedProfile_V4ToV5_MultipleEntries(t *testing.T) {
	profileID, entryID1, entryID2, _ := skipIfMissingMultipleEntryIDs(t)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v5MultipleEntriesConfig, rnd, accountID, profileID, entryID1, entryID2)
	sourceVer, targetVer := acctest.InferMigrationVersions(currentProviderVersion)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
			},
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, currentProviderVersion, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_dlp_predefined_profile."+rnd,
					tfjsonpath.New("account_id"),
					knownvalue.StringExact(accountID),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_dlp_predefined_profile."+rnd,
					tfjsonpath.New("enabled_entries"),
					knownvalue.ListSizeExact(2),
				),
			}),
		},
	})
}

// TestMigrateZeroTrustDLPPredefinedProfile_V4ToV5_WithOCR tests migration with ocr_enabled set.
func TestMigrateZeroTrustDLPPredefinedProfile_V4ToV5_WithOCR(t *testing.T) {
	profileID, entryID := skipIfMissingPredefinedIDs(t)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v5WithOCRConfig, rnd, accountID, profileID, entryID)
	sourceVer, targetVer := acctest.InferMigrationVersions(currentProviderVersion)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
			},
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, currentProviderVersion, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_dlp_predefined_profile."+rnd,
					tfjsonpath.New("account_id"),
					knownvalue.StringExact(accountID),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_dlp_predefined_profile."+rnd,
					tfjsonpath.New("ocr_enabled"),
					knownvalue.Bool(true),
				),
			}),
		},
	})
}

// TestMigrateZeroTrustDLPPredefinedProfile_V4ToV5_FromV516 tests upgrade from v5.16.0 to current.
func TestMigrateZeroTrustDLPPredefinedProfile_V4ToV5_FromV516(t *testing.T) {
	profileID, entryID := skipIfMissingPredefinedIDs(t)

	os.Setenv("TF_MIG_TEST", "")
	t.Cleanup(func() { os.Setenv("TF_MIG_TEST", "1") })

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v516BasicConfig, rnd, accountID, profileID, entryID)
	sourceVer, targetVer := acctest.InferMigrationVersions("5.16.0")

	steps := []resource.TestStep{
		{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "5.16.0",
				},
			},
			Config: testConfig,
		},
		{
			PreConfig: func() {
				acctest.WriteOutConfig(t, testConfig, tmpDir)
				acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)
			},
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			ConfigDirectory:          config.StaticDirectory(tmpDir),
		},
		{
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			ConfigDirectory:          config.StaticDirectory(tmpDir),
			ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_dlp_predefined_profile."+rnd,
					tfjsonpath.New("account_id"),
					knownvalue.StringExact(accountID),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_dlp_predefined_profile."+rnd,
					tfjsonpath.New("profile_id"),
					knownvalue.StringRegexp(regexp.MustCompile(".+")),
				),
			},
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps:      steps,
	})
}
