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

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_multiple_entries.tf
var v4MultipleEntriesConfig string

//go:embed testdata/v5_multiple_entries.tf
var v5MultipleEntriesConfig string

//go:embed testdata/v4_minimal.tf
var v4MinimalConfig string

//go:embed testdata/v5_minimal.tf
var v5MinimalConfig string

//go:embed testdata/v4_complex_patterns.tf
var v4ComplexPatternsConfig string

//go:embed testdata/v5_complex_patterns.tf
var v5ComplexPatternsConfig string

//go:embed testdata/v4_no_description.tf
var v4NoDescriptionConfig string

//go:embed testdata/v5_16_basic.tf
var v516BasicConfig string

//go:embed testdata/v5_no_description.tf
var v5NoDescriptionConfig string

// TestMigrateZeroTrustDLPCustomProfile_V4ToV5_BasicProfile tests migration of a simple DLP profile
func TestMigrateZeroTrustDLPCustomProfile_V4ToV5_BasicProfile(t *testing.T) {
	testCases := []struct {
		name       string
		version    string
		disableMig bool // set migration mode="" so production upgrader runs (for older v5 releases)
		configFn   func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("account_id"),
							knownvalue.StringExact(accountID),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("name"),
							knownvalue.StringRegexp(regexp.MustCompile(".*-"+rnd+"$")),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("allowed_match_count"),
							knownvalue.Int64Exact(5),
						),
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDLPCustomProfile_V4ToV5_FromV516 tests upgrade from v5.16.0 to current.
// v5.16.0 had schema_version=0 and no context_awareness Computed default.
// Upgrading to current adds a Computed default for context_awareness, producing a one-time plan diff.
func TestMigrateZeroTrustDLPCustomProfile_V4ToV5_FromV516(t *testing.T) {
	os.Setenv("migration mode", "")
	t.Cleanup(func() { os.Setenv("migration mode", "1") })

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v516BasicConfig, rnd, accountID)
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
		// Step 2: Run migration — allow non-empty plan (context_awareness goes from null to computed default)
		{
			PreConfig: func() {
				acctest.WriteOutConfig(t, testConfig, tmpDir)
				acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)
			},
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			ConfigDirectory:          config.StaticDirectory(tmpDir),
		},
		// Step 3: Apply correction and verify final state is clean
		{
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			ConfigDirectory:          config.StaticDirectory(tmpDir),
			ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_dlp_custom_profile."+rnd,
					tfjsonpath.New("account_id"),
					knownvalue.StringExact(accountID),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_dlp_custom_profile."+rnd,
					tfjsonpath.New("name"),
					knownvalue.StringRegexp(regexp.MustCompile(".*-"+rnd+"$")),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_dlp_custom_profile."+rnd,
					tfjsonpath.New("allowed_match_count"),
					knownvalue.Int64Exact(5),
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

// TestMigrateZeroTrustDLPCustomProfile_V4ToV5_MultipleEntries tests migration with multiple entries
func TestMigrateZeroTrustDLPCustomProfile_V4ToV5_MultipleEntries(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4MultipleEntriesConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5MultipleEntriesConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("account_id"),
							knownvalue.StringExact(accountID),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("name"),
							knownvalue.StringExact("multi-pattern-"+rnd),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("allowed_match_count"),
							knownvalue.Int64Exact(10),
						),
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDLPCustomProfile_V4ToV5_MinimalProfile tests migration with minimal config
func TestMigrateZeroTrustDLPCustomProfile_V4ToV5_MinimalProfile(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4MinimalConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5MinimalConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("account_id"),
							knownvalue.StringExact(accountID),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("name"),
							knownvalue.StringExact("minimal-"+rnd),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("allowed_match_count"),
							knownvalue.Int64Exact(1),
						),
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDLPCustomProfile_V4ToV5_ComplexPatterns tests migration with complex validation patterns
func TestMigrateZeroTrustDLPCustomProfile_V4ToV5_ComplexPatterns(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4ComplexPatternsConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5ComplexPatternsConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("account_id"),
							knownvalue.StringExact(accountID),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("name"),
							knownvalue.StringExact("complex-"+rnd),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("allowed_match_count"),
							knownvalue.Int64Exact(3),
						),
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDLPCustomProfile_V4ToV5_NoDescription tests migration without description field
func TestMigrateZeroTrustDLPCustomProfile_V4ToV5_NoDescription(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4NoDescriptionConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5NoDescriptionConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("account_id"),
							knownvalue.StringExact(accountID),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_dlp_custom_profile."+rnd,
							tfjsonpath.New("allowed_match_count"),
							knownvalue.Int64Exact(0),
						),
					}),
				},
			})
		})
	}
}
