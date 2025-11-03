package argo_smart_routing_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// TestMigrateArgo_SmartRoutingOnly tests migration when only smart_routing attribute is present
// This test verifies that:
// - cloudflare_argo resource is split into cloudflare_argo_smart_routing
// - smart_routing attribute is renamed to value
// - moved block is created for state migration
func TestMigrateArgo_SmartRoutingOnly(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_argo_smart_routing." + rnd
	tmpDir := t.TempDir()

	// V4 config using cloudflare_argo with smart_routing
	v4Config := fmt.Sprintf(`
resource "cloudflare_argo" "%[1]s" {
  zone_id        = "%[2]s"
  smart_routing  = "on"
}`, rnd, zoneID)

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
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				// Verify smart_routing was renamed to value
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
			}),
		},
	})
}

// TestMigrateArgo_TieredCachingOnly tests migration when only tiered_caching attribute is present
// This test verifies that:
// - cloudflare_argo resource is split into cloudflare_argo_tiered_caching
// - tiered_caching attribute is renamed to value
func TestMigrateArgo_TieredCachingOnly(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_argo_tiered_caching." + rnd
	tmpDir := t.TempDir()

	// V4 config using cloudflare_argo with tiered_caching
	v4Config := fmt.Sprintf(`
resource "cloudflare_argo" "%[1]s" {
  zone_id         = "%[2]s"
  tiered_caching  = "on"
}`, rnd, zoneID)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				// Verify tiered_caching was renamed to value
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
			}),
		},
	})
}

// TestMigrateArgo_BothAttributes tests migration when both attributes are present
// This test verifies that:
// - cloudflare_argo splits into both cloudflare_argo_smart_routing and cloudflare_argo_tiered_caching
// - Each resource gets the appropriate value
// - Tiered caching resource gets _tiered suffix to avoid naming conflicts
func TestMigrateArgo_BothAttributes(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	smartRoutingResourceName := "cloudflare_argo_smart_routing." + rnd
	tieredCachingResourceName := "cloudflare_argo_tiered_caching." + rnd + "_tiered"
	tmpDir := t.TempDir()

	// V4 config using cloudflare_argo with both attributes
	v4Config := fmt.Sprintf(`
resource "cloudflare_argo" "%[1]s" {
  zone_id         = "%[2]s"
  smart_routing   = "on"
  tiered_caching  = "on"
}`, rnd, zoneID)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify smart_routing resource
				statecheck.ExpectKnownValue(smartRoutingResourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(smartRoutingResourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				// Verify tiered_caching resource with _tiered suffix
				statecheck.ExpectKnownValue(tieredCachingResourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(tieredCachingResourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
			}),
		},
	})
}

// TestMigrateArgo_NoAttributes tests migration when neither attribute is present
// This test verifies that:
// - A smart_routing resource is created with value = "off" as default
func TestMigrateArgo_NoAttributes(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_argo_smart_routing." + rnd
	tmpDir := t.TempDir()

	// V4 config using cloudflare_argo with no attributes (minimal config)
	v4Config := fmt.Sprintf(`
resource "cloudflare_argo" "%[1]s" {
  zone_id = "%[2]s"
}`, rnd, zoneID)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				// Verify default value of "off" is set
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
			}),
		},
	})
}

// TestMigrateArgo_SmartRoutingOff tests migration with smart_routing = "off"
func TestMigrateArgo_SmartRoutingOff(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_argo_smart_routing." + rnd
	tmpDir := t.TempDir()

	// V4 config with smart_routing explicitly set to "off"
	v4Config := fmt.Sprintf(`
resource "cloudflare_argo" "%[1]s" {
  zone_id       = "%[2]s"
  smart_routing = "off"
}`, rnd, zoneID)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
			}),
		},
	})
}
