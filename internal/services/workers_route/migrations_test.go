package workers_route_test

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

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// TestMigrateWorkersRouteMigrationFromV4Basic tests basic migration from v4 to v5
func TestMigrateWorkersRouteMigrationFromV4Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_route." + rnd
	tmpDir := t.TempDir()
	scriptName := fmt.Sprintf("test-script-%s", rnd)

	// V4 config using script_name attribute
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_route" "%[1]s" {
  zone_id     = "%[2]s"
  pattern     = "example.cfapi.net/*"
  script_name = "%[3]s"
}`, rnd, zoneID, scriptName)

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
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact("example.cfapi.net/*")),
				// Verify script_name -> script transformation
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.StringExact(scriptName)),
			}),
		},
	})
}

// TestMigrateWorkersRouteMigrationFromV4NoScript tests migration of route without script
func TestMigrateWorkersRouteMigrationFromV4NoScript(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_route." + rnd
	tmpDir := t.TempDir()

	// V4 config without script_name (optional attribute)
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_route" "%[1]s" {
  zone_id = "%[2]s"
  pattern = "example.cfapi.net/api/*"
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
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact("example.cfapi.net/api/*")),
				// Verify script attribute doesn't exist (wasn't set in v4)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.Null()),
			}),
		},
	})
}

// TestMigrateWorkersRouteMigrationFromV4SingleResource tests migration using singular resource name
func TestMigrateWorkersRouteMigrationFromV4SingleResource(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_worker_route." + rnd
	tmpDir := t.TempDir()
	scriptName := fmt.Sprintf("test-script-%s", rnd)

	// V4 config using old singular resource name
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_route" "%[1]s" {
  zone_id     = "%[2]s"
  pattern     = "*.cfapi.net/v1/*"
  script_name = "%[3]s"
}`, rnd, zoneID, scriptName)

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
			// Step 2: Run migration and verify state (note: resource name should be updated by Grit)
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact("*.cfapi.net/v1/*")),
				// Verify script_name -> script transformation
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.StringExact(scriptName)),
			}),
		},
	})
}