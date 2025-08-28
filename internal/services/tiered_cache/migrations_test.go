package tiered_cache_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// Config generators for different provider versions
func tieredCacheConfigV4Smart(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id    = "%[2]s"
  cache_type = "smart"
}`, rnd, zoneID)
}

func tieredCacheConfigV4Off(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id    = "%[2]s"
  cache_type = "off"
}`, rnd, zoneID)
}

func tieredCacheConfigV4Generic(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id    = "%[2]s"
  cache_type = "generic"
}`, rnd, zoneID)
}

func tieredCacheConfigV5On(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id = "%[2]s"
  value   = "on"
}`, rnd, zoneID)
}

func tieredCacheConfigV5Off(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id = "%[2]s"
  value   = "off"
}`, rnd, zoneID)
}

// TestMigrateTieredCacheMultiVersion tests migration from multiple provider versions
func TestMigrateTieredCacheMultiVersion_Smart(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	testCases := []struct {
		name       string
		version    string
		configFunc func(rnd, zoneID string) string
	}{
		{
			name:       "from_v4_52_1",
			version:    "4.52.1",
			configFunc: tieredCacheConfigV4Smart,
		},
		{
			name:       "from_v5_0_0",
			version:    "5.0.0",
			configFunc: tieredCacheConfigV5On,
		},
		{
			name:       "from_v5_8_4", // Current stable release
			version:    "5.8.4",
			configFunc: tieredCacheConfigV5On,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_tiered_cache." + rnd
			tmpDir := t.TempDir()

			config := tc.configFunc(rnd, zoneID)

			// Build test steps
			steps := []resource.TestStep{}

			// Step 1: Create with specific provider version
			steps = append(steps, resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: config,
			})

			// Step 2: Run migration (for v4) or just upgrade provider (for v5)
			steps = append(steps,
				acctest.MigrationTestStep(t, config, tmpDir, tc.version, []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
				}),
			)

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

// TestMigrateTieredCacheMultiVersion tests migration with "off" value
func TestMigrateTieredCacheMultiVersion_Off(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	testCases := []struct {
		name       string
		version    string
		configFunc func(rnd, zoneID string) string
	}{
		{
			name:       "from_v4_52_1",
			version:    "4.52.1",
			configFunc: tieredCacheConfigV4Off,
		},
		{
			name:       "from_v5_0_0",
			version:    "5.0.0",
			configFunc: tieredCacheConfigV5Off,
		},
		{
			name:       "from_v5_8_4", // Current stable release
			version:    "5.8.4",
			configFunc: tieredCacheConfigV5Off,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_tiered_cache." + rnd
			tmpDir := t.TempDir()

			config := tc.configFunc(rnd, zoneID)

			// Build test steps
			steps := []resource.TestStep{}

			// Step 1: Create with specific provider version
			steps = append(steps, resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: config,
			})

			// Step 2: Run migration (for v4) or just upgrade provider (for v5)
			steps = append(steps,
				acctest.MigrationTestStep(t, config, tmpDir, tc.version, []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
				}),
			)

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

// TestMigrateTieredCache_Generic tests migration from v4 cache_type = "generic"
// to cloudflare_argo_tiered_caching resource with ResourceWithMoveState support
func TestMigrateTieredCache_Generic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	originalResourceName := "cloudflare_tiered_cache." + rnd
	migratedResourceName := "cloudflare_argo_tiered_caching." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()

	// V4 config using cache_type = "generic"
	v4Config := tieredCacheConfigV4Generic(rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(originalResourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(originalResourceName, tfjsonpath.New("cache_type"), knownvalue.StringExact("generic")),
					statecheck.ExpectKnownValue(originalResourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				},
			},
			// Step 2: Run migration and verify state transformation
			// After migration, the resource should be moved to cloudflare_argo_tiered_caching
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				// The resource should now exist as cloudflare_argo_tiered_caching with value="on"
				statecheck.ExpectKnownValue(migratedResourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(migratedResourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				statecheck.ExpectKnownValue(migratedResourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
			}),
		},
	})
}

// TestMigrateTieredCache_EdgeCases tests various edge cases in migration
func TestMigrateTieredCache_EdgeCases(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	// Test edge cases with value transformations
	edgeCases := []struct {
		name           string
		version        string
		configFunc     func(rnd, zoneID string) string
		expectedValue  string
		expectedChecks func(resourceName, zoneID, expectedValue string) []statecheck.StateCheck
	}{
		{
			name:          "v4_smart_to_on",
			version:       "4.52.1",
			configFunc:    tieredCacheConfigV4Smart,
			expectedValue: "on",
			expectedChecks: func(resourceName, zoneID, expectedValue string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(expectedValue)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				}
			},
		},
		{
			name:          "v4_off_to_off",
			version:       "4.52.1",
			configFunc:    tieredCacheConfigV4Off,
			expectedValue: "off",
			expectedChecks: func(resourceName, zoneID, expectedValue string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(expectedValue)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				}
			},
		},
		{
			name:          "v5_on_remains_on",
			version:       "5.0.0",
			configFunc:    tieredCacheConfigV5On,
			expectedValue: "on",
			expectedChecks: func(resourceName, zoneID, expectedValue string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(expectedValue)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				}
			},
		},
		{
			name:          "v5_off_remains_off",
			version:       "5.8.4",
			configFunc:    tieredCacheConfigV5Off,
			expectedValue: "off",
			expectedChecks: func(resourceName, zoneID, expectedValue string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(expectedValue)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				}
			},
		},
	}

	for _, tc := range edgeCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_tiered_cache." + rnd
			tmpDir := t.TempDir()

			config := tc.configFunc(rnd, zoneID)
			expectedChecks := tc.expectedChecks(resourceName, zoneID, tc.expectedValue)

			// Build test steps
			steps := []resource.TestStep{}

			// Step 1: Create with specific provider version
			steps = append(steps, resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: config,
			})

			// Step 2: Run migration (for v4) or just upgrade provider (for v5)
			steps = append(steps,
				acctest.MigrationTestStep(t, config, tmpDir, tc.version, expectedChecks),
			)

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
