package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var currentProviderVersion = internal.PackageVersion

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v4_with_advertisement.tf
var v4WithAdvertisementConfig string

//go:embed testdata/v4_minimal.tf
var v4MinimalConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

// buildFirstStep constructs step 1 for the migration test.
//
// For from_v4_latest: the v4 provider's Create is a no-op (sets ID from prefix_id and
// reads the existing prefix), so the resource must be imported using ImportState.
//
// For from_v5: the v5 provider can create prefixes via the API, so a normal apply is used.
func buildFirstStep(tc struct {
	name    string
	version string
}, resourceName, testConfig, accountID, prefixID string) resource.TestStep {
	if tc.version == currentProviderVersion {
		return resource.TestStep{
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			Config:                   testConfig,
		}
	}
	return resource.TestStep{
		ExternalProviders: map[string]resource.ExternalProvider{
			"cloudflare": {
				Source:            "cloudflare/cloudflare",
				VersionConstraint: tc.version,
			},
		},
		Config:            testConfig,
		ResourceName:      resourceName,
		ImportState:       true,
		ImportStateId:     fmt.Sprintf("%s/%s", accountID, prefixID),
		ImportStateVerify: true,
	}
}

// TestMigrateBYOIPPrefix_Basic tests basic migration with description.
//
// from_v4_latest: verifies state upgrade (0→1) via UpgradeFromV0, field rename
// (prefix_id→id), removal of advertisement, and API-populated computed fields.
//
// from_v5: verifies an empty plan when upgrading within v5 (no schema change).
func TestMigrateBYOIPPrefix_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")
	cidr := os.Getenv("CLOUDFLARE_BYO_IP_CIDR")
	asn, _ := strconv.ParseInt(os.Getenv("CLOUDFLARE_BYO_IP_ASN"), 10, 64)

	testCases := []struct {
		name    string
		version string
	}{
		{name: "from_v4_latest", version: acctest.GetLastV4Version()},
		{name: "from_v5", version: currentProviderVersion},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_byo_ip_prefix." + rnd
			tmpDir := t.TempDir()
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var testConfig string
			if tc.version == currentProviderVersion {
				if accountID == "" || cidr == "" || os.Getenv("CLOUDFLARE_BYO_IP_ASN") == "" {
					t.Skip("CLOUDFLARE_ACCOUNT_ID, CLOUDFLARE_BYO_IP_ASN, and CLOUDFLARE_BYO_IP_CIDR must be set")
				}
				testConfig = fmt.Sprintf(v5BasicConfig, rnd, accountID, asn, cidr)
			} else {
				if accountID == "" || prefixID == "" {
					t.Skip("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_BYO_IP_PREFIX_ID must be set")
				}
				testConfig = fmt.Sprintf(v4BasicConfig, rnd, accountID, prefixID)
			}

			stateChecks := []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Migration test prefix")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("advertised"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("asn"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cidr"), knownvalue.NotNull()),
			}
			// For v4 migration, verify the exact prefix ID was preserved
			if tc.version != currentProviderVersion {
				stateChecks = append(stateChecks,
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(prefixID)),
				)
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					buildFirstStep(tc, resourceName, testConfig, accountID, prefixID),
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, stateChecks),
				},
			})
		})
	}
}

// TestMigrateBYOIPPrefix_WithAdvertisement tests that the v4 advertisement field is
// correctly dropped during migration and replaced by the v5 computed advertised field.
func TestMigrateBYOIPPrefix_WithAdvertisement(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")
	cidr := os.Getenv("CLOUDFLARE_BYO_IP_CIDR")
	asn, _ := strconv.ParseInt(os.Getenv("CLOUDFLARE_BYO_IP_ASN"), 10, 64)

	testCases := []struct {
		name    string
		version string
	}{
		{name: "from_v4_latest", version: acctest.GetLastV4Version()},
		{name: "from_v5", version: currentProviderVersion},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_byo_ip_prefix." + rnd
			tmpDir := t.TempDir()
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var testConfig string
			if tc.version == currentProviderVersion {
				// from_v5: advertisement doesn't exist in v5; use basic v5 config
				if accountID == "" || cidr == "" || os.Getenv("CLOUDFLARE_BYO_IP_ASN") == "" {
					t.Skip("CLOUDFLARE_ACCOUNT_ID, CLOUDFLARE_BYO_IP_ASN, and CLOUDFLARE_BYO_IP_CIDR must be set")
				}
				testConfig = fmt.Sprintf(v5BasicConfig, rnd, accountID, asn, cidr)
			} else {
				if accountID == "" || prefixID == "" {
					t.Skip("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_BYO_IP_PREFIX_ID must be set")
				}
				testConfig = fmt.Sprintf(v4WithAdvertisementConfig, rnd, accountID, prefixID)
			}

			stateChecks := []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// v4 advertisement is removed; v5 advertised is populated by API
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("advertised"), knownvalue.NotNull()),
			}
			if tc.version != currentProviderVersion {
				stateChecks = append(stateChecks,
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(prefixID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Migration test with advertisement")),
				)
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					buildFirstStep(tc, resourceName, testConfig, accountID, prefixID),
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, stateChecks),
				},
			})
		})
	}
}

// TestMigrateBYOIPPrefix_Minimal tests migration with only the required v4 fields,
// verifying that v5 computed fields (asn, cidr, advertised) are populated from the API.
func TestMigrateBYOIPPrefix_Minimal(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")
	cidr := os.Getenv("CLOUDFLARE_BYO_IP_CIDR")
	asn, _ := strconv.ParseInt(os.Getenv("CLOUDFLARE_BYO_IP_ASN"), 10, 64)

	testCases := []struct {
		name    string
		version string
	}{
		{name: "from_v4_latest", version: acctest.GetLastV4Version()},
		{name: "from_v5", version: currentProviderVersion},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_byo_ip_prefix." + rnd
			tmpDir := t.TempDir()
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var testConfig string
			if tc.version == currentProviderVersion {
				// from_v5: v5 requires asn and cidr; use basic v5 config
				if accountID == "" || cidr == "" || os.Getenv("CLOUDFLARE_BYO_IP_ASN") == "" {
					t.Skip("CLOUDFLARE_ACCOUNT_ID, CLOUDFLARE_BYO_IP_ASN, and CLOUDFLARE_BYO_IP_CIDR must be set")
				}
				testConfig = fmt.Sprintf(v5BasicConfig, rnd, accountID, asn, cidr)
			} else {
				if accountID == "" || prefixID == "" {
					t.Skip("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_BYO_IP_PREFIX_ID must be set")
				}
				testConfig = fmt.Sprintf(v4MinimalConfig, rnd, accountID, prefixID)
			}

			stateChecks := []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// v5 computed fields populated from API after UpgradeState (or initial Read for v5 create)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("asn"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cidr"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("advertised"), knownvalue.NotNull()),
			}
			if tc.version != currentProviderVersion {
				stateChecks = append(stateChecks,
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(prefixID)),
				)
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					buildFirstStep(tc, resourceName, testConfig, accountID, prefixID),
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, stateChecks),
				},
			})
		})
	}
}
