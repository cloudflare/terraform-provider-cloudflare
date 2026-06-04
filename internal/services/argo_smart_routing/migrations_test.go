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
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateArgoSmartRoutingOnly tests migration from cloudflare_argo with only smart_routing attribute
func TestMigrateArgoSmartRoutingOn(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_argo_smart_routing." + rnd
	tmpDir := t.TempDir()

	// V4 config - cloudflare_argo with smart_routing only
	v4Config := fmt.Sprintf(`
resource "cloudflare_argo" "%[1]s" {
  zone_id       = "%[2]s"
  smart_routing = "on"
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
				// Verify resource migrated to cloudflare_argo_smart_routing
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				// Verify computed fields are present
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateArgoSmartRoutingOff tests migration with smart_routing set to off
func TestMigrateArgoSmartRoutingOff(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_argo_smart_routing." + rnd
	tmpDir := t.TempDir()

	// V4 config - cloudflare_argo with smart_routing off
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateArgoNeitherAttribute tests migration when neither attribute exists (defaults to smart_routing off)
func TestMigrateArgoNeitherAttribute(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_argo_smart_routing." + rnd
	tmpDir := t.TempDir()

	// V4 config - cloudflare_argo with neither smart_routing nor tiered_caching
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
				// Should default to smart_routing with value off
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("off")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateArgoBothAttributes tests migration when both smart_routing and tiered_caching exist
func TestMigrateArgoBothAttributes(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	smartRoutingResourceName := "cloudflare_argo_smart_routing." + rnd
	tieredCachingResourceName := "cloudflare_argo_tiered_caching." + rnd + "_tiered"
	tmpDir := t.TempDir()

	// V4 config - cloudflare_argo with both attributes
	v4Config := fmt.Sprintf(`
resource "cloudflare_argo" "%[1]s" {
  zone_id        = "%[2]s"
  smart_routing  = "on"
  tiered_caching = "on"
}`, rnd, zoneID)

	stateChecks := []statecheck.StateCheck{
		// Verify smart_routing resource
		statecheck.ExpectKnownValue(smartRoutingResourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(smartRoutingResourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(smartRoutingResourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
		statecheck.ExpectKnownValue(smartRoutingResourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
		// Verify tiered_caching resource with _tiered suffix
		statecheck.ExpectKnownValue(tieredCachingResourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(tieredCachingResourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(tieredCachingResourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
		statecheck.ExpectKnownValue(tieredCachingResourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
	}

	steps := []resource.TestStep{
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
	}

	// Steps 2-3: Run migration and verify state (allows creates for split resources)
	steps = append(steps, acctest.MigrationV2TestStepAllowCreate(t, v4Config, tmpDir, "4.52.1", "v4", "v5", stateChecks)...)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps:      steps,
	})
}

// TestMigrateArgoWithVariables tests migration with variable references
func TestMigrateArgoWithVariables(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_argo_smart_routing." + rnd
	tmpDir := t.TempDir()

	// V4 config with variable reference
	v4Config := fmt.Sprintf(`
variable "zone_id" {
  type    = string
  default = "%[2]s"
}

resource "cloudflare_argo" "%[1]s" {
  zone_id       = var.zone_id
  smart_routing = "on"
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateArgoWithDependsOn tests migration with depends_on meta-argument
func TestMigrateArgoWithDependsOn(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_argo_smart_routing." + rnd
	tmpDir := t.TempDir()

	// V4 config with depends_on to test meta-argument preservation
	v4Config := fmt.Sprintf(`
variable "zone_id" {
  type    = string
  default = "%[2]s"
}

resource "cloudflare_argo" "%[1]s" {
  zone_id       = var.zone_id
  smart_routing = "on"

  depends_on = [var.zone_id]
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateArgoWithLifecycle tests migration with lifecycle meta-argument
func TestMigrateArgoWithLifecycle(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_argo_smart_routing." + rnd
	tmpDir := t.TempDir()

	// V4 config with lifecycle block to test preservation
	v4Config := fmt.Sprintf(`
resource "cloudflare_argo" "%[1]s" {
  zone_id       = "%[2]s"
  smart_routing = "on"

  lifecycle {
    ignore_changes = [smart_routing]
  }
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateArgoWithConditional tests migration with conditional expression
func TestMigrateArgoWithConditional(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_argo_smart_routing." + rnd
	tmpDir := t.TempDir()

	// V4 config with conditional expression
	v4Config := fmt.Sprintf(`
variable "enable_smart_routing" {
  type    = bool
  default = true
}

resource "cloudflare_argo" "%[1]s" {
  zone_id       = "%[2]s"
  smart_routing = var.enable_smart_routing ? "on" : "off"
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
			// Step 2: Run migration and verify state (should evaluate to "on" since default is true)
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("editable"), knownvalue.Bool(true)),
			}),
		},
	})
}
