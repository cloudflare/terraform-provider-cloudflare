package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Test URL that contains required tokens for all custom page types
const testCustomPagesURL = "https://custom-pages-basic.terraform-provider-acceptance-testing.workers.dev/"

// Embed test configs
//
//go:embed testdata/v4_account_basic.tf
var v4AccountBasicConfig string

//go:embed testdata/v5_account_basic.tf
var v5AccountBasicConfig string

//go:embed testdata/v4_zone_basic.tf
var v4ZoneBasicConfig string

//go:embed testdata/v5_zone_basic.tf
var v5ZoneBasicConfig string

/* Migration tests don't include every possible permutation, but do cover:
 * - At least one migration from v4 and one from v5 for each identifier
 * - Account-level and zone-level custom pages migrated from both v4 and v5
 * - Migrations of "default" and "customized" pages
 * - Field transformation: type → identifier
 * - State field default handling
 */

// TestMigrateCustomPages_AccountLevel_Basic tests account-level custom page migration
// with DUAL test cases (from_v4_latest and from_v5)
func TestMigrateCustomPages_AccountLevel_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, identifier, state, url string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, identifier, state, url string) string {
				return fmt.Sprintf(v4AccountBasicConfig, rnd, accountID, identifier, state, url)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, accountID, identifier, state, url string) string {
				return fmt.Sprintf(v5AccountBasicConfig, rnd, accountID, identifier, state, url)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_custom_pages." + rnd
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			tmpDir := t.TempDir()

			identifier := "500_errors"
			state := "customized"
			url := testCustomPagesURL

			testConfig := tc.configFn(rnd, accountID, identifier, state, url)
			sourceVer, targetVer := "v4", "v5"
			if tc.version != acctest.GetLastV4Version() {
				sourceVer, targetVer = "v5", "v5"
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific provider version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact(identifier)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact(state)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(url)),
					}),
				},
			})
		})
	}
}

// TestMigrateCustomPages_ZoneLevel_Basic tests zone-level custom page migration
// with DUAL test cases (from_v4_latest and from_v5)
func TestMigrateCustomPages_ZoneLevel_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, identifier, state, url string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, identifier, state, url string) string {
				return fmt.Sprintf(v4ZoneBasicConfig, rnd, zoneID, identifier, state, url)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, identifier, state, url string) string {
				return fmt.Sprintf(v5ZoneBasicConfig, rnd, zoneID, identifier, state, url)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_custom_pages." + rnd
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			tmpDir := t.TempDir()

			identifier := "500_errors"
			state := "customized"
			url := testCustomPagesURL

			testConfig := tc.configFn(rnd, zoneID, identifier, state, url)
			sourceVer, targetVer := "v4", "v5"
			if tc.version != acctest.GetLastV4Version() {
				sourceVer, targetVer = "v5", "v5"
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific provider version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact(identifier)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact(state)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(url)),
					}),
				},
			})
		})
	}
}

// TestMigrateCustomPages_V5_NoOp removed - v5→v5 upgrade is already covered
// by the "from_v5" test cases in TestMigrateCustomPages_AccountLevel_Basic
// and TestMigrateCustomPages_ZoneLevel_Basic above.
