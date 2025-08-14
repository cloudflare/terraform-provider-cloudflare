package argo_tiered_caching_test

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// TestAccCloudflareArgoTieredCaching_Migration_FromTieredCacheGeneric_MultiVersion tests migration
// from v4 cloudflare_tiered_cache with cache_type="generic" to v5 cloudflare_argo_tiered_caching.
// This validates the state move functionality where:
// 1. The resource type changes from tiered_cache to argo_tiered_caching
// 2. The cache_type="generic" attribute is removed and value="on" is set
// 3. A moved block is generated to handle the state transfer
func TestAccCloudflareArgoTieredCaching_Migration_FromTieredCacheGeneric_MultiVersion(t *testing.T) {
	// Based on breaking changes analysis:
	// - All breaking changes for tiered_cache happened between 4.x and 5.0.0
	// - cache_type was removed and replaced with value
	// - Resource split: generic type becomes argo_tiered_caching in v5
	testCases := []struct {
		name    string
		version string
	}{
		{
			name:    "from_v4_52_1", // Last v4 release
			version: "4.52.1",
		},
		// Note: We don't test v5 versions here because argo_tiered_caching
		// doesn't exist in v4, only in v5. The migration creates it.
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := acctest.TestAccCloudflareZoneID
			rnd := utils.GenerateRandomResourceName()
			// Note: resource name changes from tiered_cache to argo_tiered_caching
			oldResourceName := fmt.Sprintf("cloudflare_tiered_cache.%s", rnd)
			newResourceName := fmt.Sprintf("cloudflare_argo_tiered_caching.%s", rnd)
			tmpDir := t.TempDir()

			// V4 config with cache_type="generic"
			v4Config := fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[2]s" {
  zone_id    = "%[1]s"
  cache_type = "generic"
}
`, zoneID, rnd)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Credentials(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with v4 provider using tiered_cache with generic
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: v4Config,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(oldResourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(oldResourceName, tfjsonpath.New("cache_type"), knownvalue.StringExact("generic")),
						},
					},
					// Step 2: Run migration which should:
					// - Transform tiered_cache to argo_tiered_caching
					// - Create moved block
					// - Transform cache_type="generic" to value="on"
					acctest.MigrationTestStep(t, v4Config, tmpDir, tc.version, []statecheck.StateCheck{
						// After migration, the resource is now argo_tiered_caching
						statecheck.ExpectKnownValue(newResourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(newResourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
					}),
					{
						// Step 3: Import to verify state consistency
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ResourceName:             newResourceName,
						ImportState:              true,
						ImportStateVerify:        true,
					},
					{
						// Step 4: Apply migrated config to ensure no changes needed
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(newResourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(newResourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
							statecheck.ExpectKnownValue(newResourceName, tfjsonpath.New("editable"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(newResourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
						},
					},
				},
			})
		})
	}
}

// TestAccCloudflareTieredCache_Migration_Smart_MultiVersion tests migration from v4 to v5
// where cache_type="smart" stays as cloudflare_tiered_cache but transforms to value="on"
func TestAccCloudflareTieredCache_Migration_Smart_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(zoneID, rnd string) string
	}{
		{
			name:    "from_v4_52_1",
			version: "4.52.1",
			configFn: func(zoneID, rnd string) string {
				return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[2]s" {
  zone_id    = "%[1]s"
  cache_type = "smart"
}
`, zoneID, rnd)
			},
		},
		{
			name:    "from_v5_0_0", // v5 already has value instead of cache_type
			version: "5.0.0",
			configFn: func(zoneID, rnd string) string {
				return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[2]s" {
  zone_id = "%[1]s"
  value   = "on"
}
`, zoneID, rnd)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := acctest.TestAccCloudflareZoneID
			rnd := utils.GenerateRandomResourceName()
			resourceName := fmt.Sprintf("cloudflare_tiered_cache.%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(zoneID, rnd)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Credentials(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
						ConfigStateChecks: func() []statecheck.StateCheck {
							checks := []statecheck.StateCheck{
								statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							}
							// v4 has cache_type, v5 has value
							if tc.version == "4.52.1" {
								checks = append(checks, statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cache_type"), knownvalue.StringExact("smart")))
							} else {
								checks = append(checks, statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")))
							}
							return checks
						}(),
					},
					// Step 2: Migrate - resource stays as tiered_cache but cache_type → value
					acctest.MigrationTestStep(t, testConfig, tmpDir, tc.version, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
					}),
					{
						// Step 3: Import to verify
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ResourceName:             resourceName,
						ImportState:              true,
						ImportStateVerify:        true,
					},
					{
						// Step 4: Apply migrated config
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
						},
					},
				},
			})
		})
	}
}

// TestAccCloudflareTieredCache_Migration_Off_MultiVersion tests migration from v4 to v5
// where cache_type="off" stays as cloudflare_tiered_cache with value="off"
func TestAccCloudflareTieredCache_Migration_Off_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(zoneID, rnd string) string
	}{
		{
			name:    "from_v4_52_1",
			version: "4.52.1",
			configFn: func(zoneID, rnd string) string {
				return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[2]s" {
  zone_id    = "%[1]s"
  cache_type = "off"
}
`, zoneID, rnd)
			},
		},
		{
			name:    "from_v5_0_0",
			version: "5.0.0",
			configFn: func(zoneID, rnd string) string {
				return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[2]s" {
  zone_id = "%[1]s"
  value   = "off"
}
`, zoneID, rnd)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := acctest.TestAccCloudflareZoneID
			rnd := utils.GenerateRandomResourceName()
			resourceName := fmt.Sprintf("cloudflare_tiered_cache.%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(zoneID, rnd)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Credentials(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
						ConfigStateChecks: func() []statecheck.StateCheck {
							checks := []statecheck.StateCheck{
								statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							}
							// v4 has cache_type, v5 has value
							if tc.version == "4.52.1" {
								checks = append(checks, statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cache_type"), knownvalue.StringExact("off")))
							} else {
								checks = append(checks, statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")))
							}
							return checks
						}(),
					},
					// Step 2: Migrate - stays as tiered_cache, cache_type="off" → value="off"
					acctest.MigrationTestStep(t, testConfig, tmpDir, tc.version, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
					}),
					{
						// Step 3: Import to verify
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ResourceName:             resourceName,
						ImportState:              true,
						ImportStateVerify:        true,
					},
					{
						// Step 4: Apply migrated config
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
						},
					},
				},
			})
		})
	}
}