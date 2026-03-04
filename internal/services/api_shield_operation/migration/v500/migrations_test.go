package v500_test

import (
	_ "embed"
	"fmt"
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

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed test configs
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_post_method.tf
var v4PostMethodConfig string

//go:embed testdata/v5_post_method.tf
var v5PostMethodConfig string

//go:embed testdata/v4_path_params.tf
var v4PathParamsConfig string

//go:embed testdata/v5_path_params.tf
var v5PathParamsConfig string

// TestMigrateAPIShieldOperation_V4ToV5_Basic tests basic migration with GET method
func TestMigrateAPIShieldOperation_V4ToV5_Basic(t *testing.T) {
	t.Skip("Migration not enabled yet")
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// For v5 tests, use local provider; for v4 tests, use external provider
			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				// Use local v5 provider (has GetSchemaVersion, will create version=1 state)
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// Use external v4 provider (will create version=0 state)
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
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						// Verify v4 fields migrated correctly
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("method"), knownvalue.StringExact("GET")),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("host"), knownvalue.StringExact("api.example.com")),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("endpoint"), knownvalue.StringExact("/api/v1/users")),
						// Verify new v5 computed fields exist (populated by API)
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("operation_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// TestMigrateAPIShieldOperation_V4ToV5_PostMethod tests migration with POST method
func TestMigrateAPIShieldOperation_V4ToV5_PostMethod(t *testing.T) {
	t.Skip("Migration not enabled yet")
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4PostMethodConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5PostMethodConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
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
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("method"), knownvalue.StringExact("POST")),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("host"), knownvalue.StringExact("api.example.com")),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("endpoint"), knownvalue.StringExact("/api/v1/users")),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("operation_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// TestMigrateAPIShieldOperation_V4ToV5_PathParams tests migration with path parameters
func TestMigrateAPIShieldOperation_V4ToV5_PathParams(t *testing.T) {
	t.Skip("Migration not enabled yet")
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4PathParamsConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5PathParamsConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
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
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("method"), knownvalue.StringExact("GET")),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("host"), knownvalue.StringExact("api.example.com")),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("endpoint"), knownvalue.StringExact("/api/v1/users/{var1}/posts/{var2}")),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("operation_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("cloudflare_api_shield_operation."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}
