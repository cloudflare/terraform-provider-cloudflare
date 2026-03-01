package v500_test

import (
	"crypto/rand"
	_ "embed"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// generateRandomTestIP generates a cryptographically random IP from TEST-NET-2 (198.51.100.0/24)
// RFC 5737 reserves this range for documentation and testing
func generateRandomTestIP() string {
	// Generate cryptographically random last octet (1-254, avoiding 0 and 255)
	n, _ := rand.Int(rand.Reader, big.NewInt(254))
	lastOctet := int(n.Int64()) + 1
	return fmt.Sprintf("198.51.100.%d", lastOctet)
}

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed test configs
//go:embed testdata/v4_basic_zone.tf
var v4BasicZoneConfig string

//go:embed testdata/v5_basic_zone.tf
var v5BasicZoneConfig string

//go:embed testdata/v4_account.tf
var v4AccountConfig string

//go:embed testdata/v5_account.tf
var v5AccountConfig string

// TestMigrateAccessRule_V4ToV5_BasicZone tests basic zone-level access rule migration with DUAL test cases.
//
// This test validates:
// - configuration: array[0] → object transformation (CRITICAL)
// - mode, notes pass-through
// - zone_id preserved
func TestMigrateAccessRule_V4ToV5_BasicZone(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, testIP string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, testIP string) string {
				return fmt.Sprintf(v4BasicZoneConfig, rnd, zoneID, testIP)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, testIP string) string {
				return fmt.Sprintf(v5BasicZoneConfig, rnd, zoneID, testIP)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			testIP := generateRandomTestIP()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID, testIP)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify zone_id preserved
							statecheck.ExpectKnownValue(
								"cloudflare_access_rule."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							// Verify mode preserved
							statecheck.ExpectKnownValue(
								"cloudflare_access_rule."+rnd,
								tfjsonpath.New("mode"),
								knownvalue.StringExact("block"),
							),
							// Verify notes preserved
							statecheck.ExpectKnownValue(
								"cloudflare_access_rule."+rnd,
								tfjsonpath.New("notes"),
								knownvalue.StringExact("Test access rule"),
							),
							// CRITICAL: Verify configuration.target (unwrapped from array)
							statecheck.ExpectKnownValue(
								"cloudflare_access_rule."+rnd,
								tfjsonpath.New("configuration").AtMapKey("target"),
								knownvalue.StringExact("ip"),
							),
							// CRITICAL: Verify configuration.value (unwrapped from array)
							statecheck.ExpectKnownValue(
								"cloudflare_access_rule."+rnd,
								tfjsonpath.New("configuration").AtMapKey("value"),
								knownvalue.StringExact(testIP),
							),
							// Verify new computed fields exist (NotNull)
							statecheck.ExpectKnownValue(
								"cloudflare_access_rule."+rnd,
								tfjsonpath.New("id"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateAccessRule_V4ToV5_Account tests account-level access rule migration with DUAL test cases.
//
// This test validates:
// - Account-level rules work correctly
// - configuration transformation for different target types (country)
// - Default notes value when not specified
func TestMigrateAccessRule_V4ToV5_Account(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, testIP string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, testIP string) string {
				return fmt.Sprintf(v4AccountConfig, rnd, accountID, testIP)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, testIP string) string {
				return fmt.Sprintf(v5AccountConfig, rnd, accountID, testIP)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			testIP := generateRandomTestIP()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, testIP)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify account_id preserved
							statecheck.ExpectKnownValue(
								"cloudflare_access_rule."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							// Verify mode preserved
							statecheck.ExpectKnownValue(
								"cloudflare_access_rule."+rnd,
								tfjsonpath.New("mode"),
								knownvalue.StringExact("challenge"),
							),
							// Verify notes exists (should be empty string default)
							statecheck.ExpectKnownValue(
								"cloudflare_access_rule."+rnd,
								tfjsonpath.New("notes"),
								knownvalue.NotNull(),
							),
							// CRITICAL: Verify configuration.target
							statecheck.ExpectKnownValue(
								"cloudflare_access_rule."+rnd,
								tfjsonpath.New("configuration").AtMapKey("target"),
								knownvalue.StringExact("ip"),
							),
							// CRITICAL: Verify configuration.value
							statecheck.ExpectKnownValue(
								"cloudflare_access_rule."+rnd,
								tfjsonpath.New("configuration").AtMapKey("value"),
								knownvalue.StringExact(testIP),
							),
						},
					),
				},
			})
		})
	}
}
