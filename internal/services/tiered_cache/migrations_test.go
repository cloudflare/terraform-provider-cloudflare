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

// TestMigrateTieredCache_Smart tests migration from v4 cache_type = "smart" to v5 value = "on"
func TestMigrateTieredCache_Smart(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_tiered_cache." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()

	// V4 config using cache_type = "smart"
	v4Config := fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id    = "%[2]s"
  cache_type = "smart"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		CheckDestroy: nil, // Migration tests don't need destroy checks
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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cache_type"), knownvalue.StringExact("smart")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				},
			},
			// Step 2: Run migration and verify state transformation
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
			}),
		},
	})
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
	v4Config := fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id    = "%[2]s"
  cache_type = "generic"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		CheckDestroy: nil, // Migration tests don't need destroy checks
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

// TestMigrateTieredCache_Off tests migration from v4 cache_type = "off" to v5 value = "off"
func TestMigrateTieredCache_Off(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_tiered_cache." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()

	// V4 config using cache_type = "off"
	v4Config := fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id    = "%[2]s"
  cache_type = "off"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		CheckDestroy: nil, // Migration tests don't need destroy checks
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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cache_type"), knownvalue.StringExact("off")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				},
			},
			// Step 2: Run migration and verify state transformation
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
			}),
		},
	})
}

// TestMigrateTieredCache_AllValues tests value transformations that work with the current resource
// NOTE: Generic cache_type requires migration to cloudflare_argo_tiered_caching resource (not tested here)
func TestMigrateTieredCache_AllValues(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	testCases := []struct {
		name        string
		v4Value     string
		v5Value     string
		description string
	}{
		{
			name:        "Smart_To_On",
			v4Value:     "smart",
			v5Value:     "on",
			description: "Migration from smart tiered cache to on",
		},
		{
			name:        "Off_To_Off",
			v4Value:     "off",
			v5Value:     "off", 
			description: "Migration from off tiered cache to off",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_tiered_cache." + rnd
			tmpDir := t.TempDir()

			v4Config := fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
  zone_id    = "%[2]s"
  cache_type = "%[3]s"
}`, rnd, zoneID, tc.v4Value)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: "4.52.1",
							},
						},
						Config: v4Config,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cache_type"), knownvalue.StringExact(tc.v4Value)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
						},
					},
					acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(tc.v5Value)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					}),
				},
			})
		})
	}
}