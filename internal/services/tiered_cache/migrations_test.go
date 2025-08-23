package tiered_cache_test

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

func TestAccCloudflareTieredCache_Migration_Smart(t *testing.T) {
	zoneID := acctest.TestAccCloudflareZoneID
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_tiered_cache." + rnd
	tmpDir := t.TempDir()

	// Test migration from multiple versions
	testCases := []struct {
		name     string
		version  string
		configFn func(string, string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflareTieredCacheMigrationConfigV4Smart,
		},
		{
			name:     "from_v5_0_0",
			version:  "5.0.0",
			configFn: testAccCloudflareTieredCacheMigrationConfigV5On,
		},
		{
			name:     "from_v5_7_1",
			version:  "5.7.1",
			configFn: testAccCloudflareTieredCacheMigrationConfigV5On,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			testConfig := tc.configFn(rnd, zoneID)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create tiered cache with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: tc.version,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							// For v4, check cache_type; for v5, check value
							func() statecheck.StateCheck {
								if tc.version[0] == '4' {
									return statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cache_type"), knownvalue.StringExact("smart"))
								}
								return statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on"))
							}(),
						},
					},
					// Step 2: Migrate to v5 provider
					acctest.MigrationTestStep(t, testConfig, tmpDir, tc.version, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
					}),
					{
						// Step 3: Import and verify
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ResourceName:             resourceName,
						ImportState:              true,
						ImportStateVerify:        true,
						ImportStateVerifyIgnore:  []string{"editable", "modified_on"},
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

func TestAccCloudflareTieredCache_Migration_Off(t *testing.T) {
	zoneID := acctest.TestAccCloudflareZoneID
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_tiered_cache." + rnd
	tmpDir := t.TempDir()

	// Test migration from multiple versions
	testCases := []struct {
		name     string
		version  string
		configFn func(string, string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflareTieredCacheMigrationConfigV4Off,
		},
		{
			name:     "from_v5_0_0",
			version:  "5.0.0",
			configFn: testAccCloudflareTieredCacheMigrationConfigV5Off,
		},
		{
			name:     "from_v5_7_1",
			version:  "5.7.1",
			configFn: testAccCloudflareTieredCacheMigrationConfigV5Off,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			testConfig := tc.configFn(rnd, zoneID)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create tiered cache with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: tc.version,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							// For v4, check cache_type; for v5, check value
							func() statecheck.StateCheck {
								if tc.version[0] == '4' {
									return statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cache_type"), knownvalue.StringExact("off"))
								}
								return statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off"))
							}(),
						},
					},
					// Step 2: Migrate to v5 provider
					acctest.MigrationTestStep(t, testConfig, tmpDir, tc.version, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
					}),
					{
						// Step 3: Import and verify
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ResourceName:             resourceName,
						ImportState:              true,
						ImportStateVerify:        true,
						ImportStateVerifyIgnore:  []string{"editable", "modified_on"},
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

func TestAccCloudflareTieredCache_Migration_Generic(t *testing.T) {
	zoneID := acctest.TestAccCloudflareZoneID
	rnd := utils.GenerateRandomResourceName()

	// For generic, it becomes argo_tiered_caching after migration
	sourceResourceName := "cloudflare_tiered_cache." + rnd
	targetResourceName := "cloudflare_argo_tiered_caching." + rnd
	tmpDir := t.TempDir()

	// Only test from v4 since generic->argo_tiered_caching is a v4->v5 migration
	testConfig := testAccCloudflareTieredCacheMigrationConfigV4Generic(rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create tiered cache with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(sourceResourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(sourceResourceName, tfjsonpath.New("cache_type"), knownvalue.StringExact("generic")),
				},
			},
			// Step 2: Migrate to v5 provider (will transform to argo_tiered_caching)
			acctest.MigrationTestStep(t, testConfig, tmpDir, "4.52.1", []statecheck.StateCheck{
				// After migration, check the argo_tiered_caching resource
				statecheck.ExpectKnownValue(targetResourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(targetResourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
			}),
			{
				// Step 3: Import and verify the new resource
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             targetResourceName,
				ImportState:              true,
				ImportStateVerify:        true,
				ImportStateVerifyIgnore:  []string{"editable", "modified_on"},
			},
			{
				// Step 4: Apply migrated config
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(targetResourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(targetResourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				},
			},
		},
	})
}

// V4 Configuration Functions

func testAccCloudflareTieredCacheMigrationConfigV4Smart(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id    = "%[2]s"
  cache_type = "smart"
}
`, rnd, zoneID)
}

func testAccCloudflareTieredCacheMigrationConfigV4Off(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id    = "%[2]s"
  cache_type = "off"
}
`, rnd, zoneID)
}

func testAccCloudflareTieredCacheMigrationConfigV4Generic(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id    = "%[2]s"
  cache_type = "generic"
}
`, rnd, zoneID)
}

// V5 Configuration Functions

func testAccCloudflareTieredCacheMigrationConfigV5On(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id = "%[2]s"
  value   = "on"
}
`, rnd, zoneID)
}

func testAccCloudflareTieredCacheMigrationConfigV5Off(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id = "%[2]s"
  value   = "off"
}
`, rnd, zoneID)
}
