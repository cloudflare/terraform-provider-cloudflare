package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed test configs
//go:embed testdata/v4_tiered_caching_only.tf
var v4TieredCachingOnlyConfig string

//go:embed testdata/v5_tiered_caching_only.tf
var v5TieredCachingOnlyConfig string

// argoTieredCachingMigrationTestStep creates a test step for argo_tiered_caching migration
func argoTieredCachingMigrationTestStep(t *testing.T, testConfig string, tmpDir string, exactVersion string, sourceVersion string, targetVersion string, stateChecks []statecheck.StateCheck) resource.TestStep {
	return resource.TestStep{
		PreConfig: func() {
			acctest.WriteOutConfig(t, testConfig, tmpDir)
			acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVersion, targetVersion)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				acctest.DebugNonEmptyPlan,
			},
		},
		ConfigStateChecks: stateChecks,
	}
}

// TestMigrateArgoTieredCaching_V4ToV5_TieredCachingOnly tests migration when only tiered_caching is set
// This covers Scenario 4: tiered_caching attribute only
func TestMigrateArgoTieredCaching_V4ToV5_TieredCachingOnly(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: os.Getenv("LAST_V4_VERSION"),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4TieredCachingOnlyConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5TieredCachingOnlyConfig, rnd, zoneID)
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
					argoTieredCachingMigrationTestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify zone_id is preserved
							statecheck.ExpectKnownValue(
								"cloudflare_argo_tiered_caching."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							// Verify tiered_caching → value transformation
							statecheck.ExpectKnownValue(
								"cloudflare_argo_tiered_caching."+rnd,
								tfjsonpath.New("value"),
								knownvalue.StringExact("on"),
							),
							// Verify id changed from checksum to zone_id
							statecheck.ExpectKnownValue(
								"cloudflare_argo_tiered_caching."+rnd,
								tfjsonpath.New("id"),
								knownvalue.StringExact(zoneID),
							),
							// Verify editable is set (computed field)
							statecheck.ExpectKnownValue(
								"cloudflare_argo_tiered_caching."+rnd,
								tfjsonpath.New("editable"),
								knownvalue.Bool(true),
							),
						},
					),
				},
			})
		})
	}
}
