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

// TestMigrateWorkersRouteMigrationFromV4Basic tests basic migration from v4 to v5 
func TestMigrateWorkersRouteMigrationFromV4Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_route." + rnd
	tmpDir := t.TempDir()
	scriptName := fmt.Sprintf("test-script-%s", rnd)

	// V4 config using script_name attribute - create worker first, then route
	// Use unique pattern based on random name to avoid conflicts
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_script" "test_script_%[1]s" {
  account_id = "%[4]s"
  name       = "%[3]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}

resource "cloudflare_workers_route" "%[1]s" {
  zone_id     = "%[2]s"
  pattern     = "%[1]s.cfapi.net/*"
  script_name = "%[3]s"
  depends_on  = [cloudflare_workers_script.test_script_%[1]s]
}`, rnd, zoneID, scriptName, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_AccountID(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.cfapi.net/*", rnd))),
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
	// Use unique pattern based on random name to avoid conflicts
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_route" "%[1]s" {
  zone_id = "%[2]s"
  pattern = "%[1]s.cfapi.net/api/*"
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.cfapi.net/api/*", rnd))),
				// Verify script attribute doesn't exist (wasn't set in v4)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.Null()),
			}),
		},
	})
}

// TestMigrateWorkersRouteMigrationFromV4SingleResource tests migration using singular resource name
func TestMigrateWorkersRouteMigrationFromV4SingleResource(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_route." + rnd  // After migration, resource will be renamed
	tmpDir := t.TempDir()
	scriptName := fmt.Sprintf("test-script-%s", rnd)

	// V4 config using old singular resource name - create worker first, then route
	// Use unique pattern based on random name to avoid conflicts
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "test_script_%[1]s" {
  account_id = "%[4]s"
  name       = "%[3]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}

resource "cloudflare_worker_route" "%[1]s" {
  zone_id     = "%[2]s"
  pattern     = "%[1]s.cfapi.net/v1/*"
  script_name = "%[3]s"
  depends_on  = [cloudflare_worker_script.test_script_%[1]s]
}`, rnd, zoneID, scriptName, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_AccountID(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.cfapi.net/v1/*", rnd))),
				// Verify script_name -> script transformation
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.StringExact(scriptName)),
			}),
		},
	})
}
