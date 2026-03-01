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
//go:embed testdata/v4_smart_routing_only.tf
var v4SmartRoutingOnlyConfig string

//go:embed testdata/v5_smart_routing_only.tf
var v5SmartRoutingOnlyConfig string

//go:embed testdata/v4_both_attributes.tf
var v4BothAttributesConfig string

//go:embed testdata/v5_both_attributes.tf
var v5BothAttributesConfig string

//go:embed testdata/v4_neither_attribute.tf
var v4NeitherAttributeConfig string

//go:embed testdata/v5_neither_attribute.tf
var v5NeitherAttributeConfig string

// argoSmartRoutingMigrationTestStep creates a test step for argo_smart_routing migration
func argoSmartRoutingMigrationTestStep(t *testing.T, testConfig string, tmpDir string, exactVersion string, sourceVersion string, targetVersion string, stateChecks []statecheck.StateCheck) resource.TestStep {
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

// TestMigrateArgoSmartRouting_V4ToV5_SmartRoutingOnly tests migration when only smart_routing is set
// This covers Scenario 1: smart_routing attribute only
func TestMigrateArgoSmartRouting_V4ToV5_SmartRoutingOnly(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: os.Getenv("LAST_V4_VERSION"),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4SmartRoutingOnlyConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5SmartRoutingOnlyConfig, rnd, zoneID)
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
					argoSmartRoutingMigrationTestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify zone_id is preserved
							statecheck.ExpectKnownValue(
								"cloudflare_argo_smart_routing."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							// Verify smart_routing → value transformation
							statecheck.ExpectKnownValue(
								"cloudflare_argo_smart_routing."+rnd,
								tfjsonpath.New("value"),
								knownvalue.StringExact("on"),
							),
							// Verify id changed from checksum to zone_id
							statecheck.ExpectKnownValue(
								"cloudflare_argo_smart_routing."+rnd,
								tfjsonpath.New("id"),
								knownvalue.StringExact(zoneID),
							),
							// Verify editable is set (computed field)
							statecheck.ExpectKnownValue(
								"cloudflare_argo_smart_routing."+rnd,
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

// TestMigrateArgoSmartRouting_V4ToV5_BothAttributes tests migration when both smart_routing and tiered_caching are set
// This covers Scenario 2: Both attributes (smart_routing gets moved block in config)
func TestMigrateArgoSmartRouting_V4ToV5_BothAttributes(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: os.Getenv("LAST_V4_VERSION"),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4BothAttributesConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5BothAttributesConfig, rnd, zoneID)
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

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
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
					argoSmartRoutingMigrationTestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_argo_smart_routing."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_argo_smart_routing."+rnd,
								tfjsonpath.New("value"),
								knownvalue.StringExact("on"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_argo_smart_routing."+rnd,
								tfjsonpath.New("id"),
								knownvalue.StringExact(zoneID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_argo_smart_routing."+rnd,
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

// TestMigrateArgoSmartRouting_V4ToV5_NeitherAttribute tests migration when neither attribute is set
// This covers Scenario 3: No attributes (defaults to smart_routing with value="off")
func TestMigrateArgoSmartRouting_V4ToV5_NeitherAttribute(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: os.Getenv("LAST_V4_VERSION"),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4NeitherAttributeConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5NeitherAttributeConfig, rnd, zoneID)
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

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
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
					argoSmartRoutingMigrationTestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_argo_smart_routing."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							// Verify default value is "off" when smart_routing not set
							statecheck.ExpectKnownValue(
								"cloudflare_argo_smart_routing."+rnd,
								tfjsonpath.New("value"),
								knownvalue.StringExact("off"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_argo_smart_routing."+rnd,
								tfjsonpath.New("id"),
								knownvalue.StringExact(zoneID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_argo_smart_routing."+rnd,
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
