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

// Embed test configs
//go:embed testdata/v4_on.tf
var v4OnConfig string

//go:embed testdata/v5_on.tf
var v5OnConfig string

//go:embed testdata/v4_off.tf
var v4OffConfig string

//go:embed testdata/v5_off.tf
var v5OffConfig string

// TestMigrateRegionalTieredCache_V4ToV5_On tests migration with value="on" with DUAL test cases
func TestMigrateRegionalTieredCache_V4ToV5_On(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4OnConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5OnConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// For v5 tests, use local provider; for v4 tests, use external provider
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
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state transformation
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify core fields
							statecheck.ExpectKnownValue(
								"cloudflare_regional_tiered_cache."+rnd,
								tfjsonpath.New(consts.ZoneIDSchemaKey),
								knownvalue.StringExact(zoneID),
							),
							// Verify value migrated correctly
							statecheck.ExpectKnownValue(
								"cloudflare_regional_tiered_cache."+rnd,
								tfjsonpath.New("value"),
								knownvalue.StringExact("on"),
							),
							// Verify ID equals zone_id (critical transformation)
							statecheck.ExpectKnownValue(
								"cloudflare_regional_tiered_cache."+rnd,
								tfjsonpath.New("id"),
								knownvalue.StringExact(zoneID),
							),
							// Note: editable and modified_on are computed fields
							// They will be refreshed from API, so we don't validate exact values
						},
					),
				},
			})
		})
	}
}

// TestMigrateRegionalTieredCache_V4ToV5_Off tests migration with value="off" with DUAL test cases
func TestMigrateRegionalTieredCache_V4ToV5_Off(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4OffConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5OffConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// For v5 tests, use local provider; for v4 tests, use external provider
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
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state transformation
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify core fields
							statecheck.ExpectKnownValue(
								"cloudflare_regional_tiered_cache."+rnd,
								tfjsonpath.New(consts.ZoneIDSchemaKey),
								knownvalue.StringExact(zoneID),
							),
							// Verify value migrated correctly
							statecheck.ExpectKnownValue(
								"cloudflare_regional_tiered_cache."+rnd,
								tfjsonpath.New("value"),
								knownvalue.StringExact("off"),
							),
							// Verify ID equals zone_id (critical transformation)
							statecheck.ExpectKnownValue(
								"cloudflare_regional_tiered_cache."+rnd,
								tfjsonpath.New("id"),
								knownvalue.StringExact(zoneID),
							),
							// Note: editable and modified_on are computed fields
							// They will be refreshed from API, so we don't validate exact values
						},
					),
				},
			})
		})
	}
}
