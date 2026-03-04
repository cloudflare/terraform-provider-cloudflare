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

// Embed migration test configuration files
//
//go:embed testdata/v4_smart.tf
var v4SmartConfig string

//go:embed testdata/v4_off.tf
var v4OffConfig string

//go:embed testdata/v4_generic.tf
var v4GenericConfig string

//go:embed testdata/v5_on.tf
var v5OnConfig string

//go:embed testdata/v5_off.tf
var v5OffConfig string

//go:embed testdata/v5_variable.tf
var v5VariableConfig string

// TestMigrateTieredCache_Smart tests migration from v4 cache_type="smart" to v5 value="on"
// NOTE: The tf-migrate tool creates BOTH cloudflare_tiered_cache AND cloudflare_argo_tiered_caching
// resources. The argo_tiered_caching resource will be created on next apply (this is by design).
func TestMigrateTieredCache_Smart(t *testing.T) {
	testCases := []struct {
		name      string
		version   string
		configFn  func(rnd, zoneID string) string
		checkArgo bool // whether to check for argo_tiered_caching resource
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4SmartConfig, rnd, zoneID)
			},
			checkArgo: true, // v4 migration creates both resources
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5OnConfig, rnd, zoneID)
			},
			checkArgo: false, // v5 config only has tiered_cache
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			tieredCacheResourceName := "cloudflare_tiered_cache." + rnd
			argoTieredCachingResourceName := "cloudflare_argo_tiered_caching." + rnd
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			stateChecks := []statecheck.StateCheck{
				// Verify tiered_cache state transformation
				statecheck.ExpectKnownValue(tieredCacheResourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(tieredCacheResourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				statecheck.ExpectKnownValue(tieredCacheResourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
			}

			// Only check argo_tiered_caching for v4 migrations (where tf-migrate creates both resources)
			if tc.checkArgo {
				stateChecks = append(stateChecks,
					statecheck.ExpectKnownValue(argoTieredCachingResourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(argoTieredCachingResourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
					statecheck.ExpectKnownValue(argoTieredCachingResourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				)
			}

			// Build steps: Step 1 creates with specified provider version
			// For v5 tests, use local provider; for v4 tests, use external provider
			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				// Use local v5 provider (will create state)
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
			steps := []resource.TestStep{firstStep}

			// Steps 2-3: Run migration and verify state (allows creates for split resources)
			steps = append(steps, acctest.MigrationV2TestStepAllowCreate(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, stateChecks)...)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps:      steps,
			})
		})
	}
}

// TestMigrateTieredCache_Generic tests migration from v4 cache_type="generic"
// NOTE: The tf-migrate tool creates BOTH cloudflare_tiered_cache (with value="off") AND
// cloudflare_argo_tiered_caching (with value="on") resources.
func TestMigrateTieredCache_Generic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4GenericConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			tieredCacheResourceName := "cloudflare_tiered_cache." + rnd
			argoTieredCachingResourceName := "cloudflare_argo_tiered_caching." + rnd
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			stateChecks := []statecheck.StateCheck{
				// Verify tiered_cache state transformation (value="off" for generic)
				statecheck.ExpectKnownValue(tieredCacheResourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(tieredCacheResourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
				statecheck.ExpectKnownValue(tieredCacheResourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				// Verify new argo_tiered_caching resource was created (value="on")
				statecheck.ExpectKnownValue(argoTieredCachingResourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(argoTieredCachingResourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				statecheck.ExpectKnownValue(argoTieredCachingResourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
			}

			// Build steps: Step 1 creates with specified provider version
			steps := []resource.TestStep{
				{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				},
			}

			// Steps 2-3: Run migration and verify state (allows creates for split resources)
			steps = append(steps, acctest.MigrationV2TestStepAllowCreate(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, stateChecks)...)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps:      steps,
			})
		})
	}
}

// TestMigrateTieredCache_Off tests migration from v4 cache_type="off" to v5 value="off"
// NOTE: The tf-migrate tool creates BOTH cloudflare_tiered_cache AND cloudflare_argo_tiered_caching
// resources with value="off". The argo_tiered_caching resource will be created on next apply.
func TestMigrateTieredCache_Off(t *testing.T) {
	testCases := []struct {
		name      string
		version   string
		configFn  func(rnd, zoneID string) string
		checkArgo bool // whether to check for argo_tiered_caching resource
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4OffConfig, rnd, zoneID)
			},
			checkArgo: true, // v4 migration creates both resources
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5OffConfig, rnd, zoneID)
			},
			checkArgo: false, // v5 config only has tiered_cache
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			tieredCacheResourceName := "cloudflare_tiered_cache." + rnd
			argoTieredCachingResourceName := "cloudflare_argo_tiered_caching." + rnd
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			stateChecks := []statecheck.StateCheck{
				// Verify tiered_cache state transformation
				statecheck.ExpectKnownValue(tieredCacheResourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(tieredCacheResourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
				statecheck.ExpectKnownValue(tieredCacheResourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
			}

			// Only check argo_tiered_caching for v4 migrations (where tf-migrate creates both resources)
			if tc.checkArgo {
				stateChecks = append(stateChecks,
					statecheck.ExpectKnownValue(argoTieredCachingResourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(argoTieredCachingResourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
					statecheck.ExpectKnownValue(argoTieredCachingResourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				)
			}

			// Build steps: Step 1 creates with specified provider version
			// For v5 tests, use local provider; for v4 tests, use external provider
			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				// Use local v5 provider (will create state)
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
			steps := []resource.TestStep{firstStep}

			// Steps 2-3: Run migration and verify state (allows creates for split resources)
			steps = append(steps, acctest.MigrationV2TestStepAllowCreate(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, stateChecks)...)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps:      steps,
			})
		})
	}
}

// TestMigrateTieredCache_V5VersionUpgrade tests upgrading between v5 versions (no-op upgrade)
// This ensures that existing v5 state is compatible with the latest provider version
func TestMigrateTieredCache_V5VersionUpgrade(t *testing.T) {
	// Test different v5 versions to ensure state compatibility
	testCases := []struct {
		name    string
		version string
	}{
		{
			name:    "from_v5_9_0",
			version: "5.9.0",
		},
		{
			name:    "from_v5_10_0",
			version: "5.10.0",
		},
		{
			name:    "from_v5_12_0",
			version: "5.12.0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_tiered_cache." + rnd
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			tmpDir := t.TempDir()
			testConfig := fmt.Sprintf(v5OnConfig, rnd, zoneID)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific v5 provider version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
						},
					},
					{
						// Step 2: Upgrade to latest provider - should be a no-op
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						Config:                   testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
						},
					},
				},
			})
		})
	}
}

// TestMigrateTieredCache_V5VersionUpgrade_Off tests v5 version upgrade with value="off"
func TestMigrateTieredCache_V5VersionUpgrade_Off(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_tiered_cache." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v5OffConfig, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.9.0 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.9.0",
					},
				},
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				},
			},
			{
				// Step 2: Upgrade to latest provider - should be a no-op
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				},
			},
		},
	})
}

// TestMigrateTieredCache_V5VersionUpgrade_WithVariables tests v5 version upgrade with variable references
func TestMigrateTieredCache_V5VersionUpgrade_WithVariables(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_tiered_cache." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v5VariableConfig, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.9.0 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.9.0",
					},
				},
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				},
			},
			{
				// Step 2: Upgrade to latest provider - should be a no-op
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
				},
			},
		},
	})
}
