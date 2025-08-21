package workers_route_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// TestAccWorkersRouteMigrationFromV4Basic tests basic migration from v4 to v5 with script_name -> script rename
func TestMigrateWorkersRouteMigrationFromV4Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_route." + rnd
	tmpDir := t.TempDir()

	// V4 config using script_name attribute
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello from Worker!')) })"
}

resource "cloudflare_workers_route" "%[1]s" {
  zone_id     = "%[3]s"
  pattern     = "%[1]s.%[4]s/*"
  script_name = cloudflare_workers_script.%[1]s.name
}`, rnd, accountID, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.%s/*", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				},
			},
			// Step 2: Run migration and verify state transformation
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.%s/*", rnd, domain))),
				// Verify script_name -> script attribute rename
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.StringExact(rnd)),
			}),
		},
	})
}

// TestAccWorkersRouteMigrationFromV4NoScript tests migration of pattern-only routes without scripts
func TestMigrateWorkersRouteMigrationFromV4NoScript(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_route." + rnd
	tmpDir := t.TempDir()

	// V4 config with pattern only, no script_name (disabled route)
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_route" "%[1]s" {
  zone_id = "%[2]s"
  pattern = "%[1]s.%[3]s/*"
}`, rnd, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.%s/*", rnd, domain))),
					// No script_name should be present
				},
			},
			// Step 2: Run migration and verify state (no script attribute should be present)
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.%s/*", rnd, domain))),
				// No script attribute should be present for disabled routes
			}),
		},
	})
}

// TestAccWorkersRouteMigrationFromV4MultipleVersions tests migration from different v4 versions
func TestMigrateWorkersRouteMigrationFromV4MultipleVersions(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	testCases := []struct {
		name    string
		version string
	}{
		{
			name:    "from_v4_52_1",
			version: "4.52.1", // Last v4 release
		},
		{
			name:    "from_v4_50_0",
			version: "4.50.0", // Earlier v4 release
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_workers_route." + rnd
			tmpDir := t.TempDir()

			v4Config := fmt.Sprintf(`
resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello from Worker!')) })"
}

resource "cloudflare_workers_route" "%[1]s" {
  zone_id     = "%[3]s"
  pattern     = "%[1]s.%[4]s/*"
  script_name = cloudflare_workers_script.%[1]s.name
}`, rnd, accountID, zoneID, domain)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
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
						Config: v4Config,
					},
					acctest.MigrationTestStep(t, v4Config, tmpDir, tc.version, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.%s/*", rnd, domain))),
						// Verify script_name -> script migration works for all v4 versions
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.StringExact(rnd)),
					}),
				},
			})
		})
	}
}

// TestAccWorkersRouteMigrationFromV4ZoneScoped tests that zone scoping is preserved during migration
func TestMigrateWorkersRouteMigrationFromV4ZoneScoped(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_route." + rnd
	tmpDir := t.TempDir()

	// V4 config emphasizing zone-scoped nature
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_script" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Zone-scoped Worker')) })"
}

resource "cloudflare_workers_route" "%[1]s" {
  zone_id     = "%[3]s"
  pattern     = "*.%[4]s/api/*"
  script_name = cloudflare_workers_script.%[1]s.name
}`, rnd, accountID, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("*.%s/api/*", domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				},
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				// Verify zone scoping is preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("*.%s/api/*", domain))),
				// Verify attribute migration
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.StringExact(rnd)),
			}),
			{
				// Step 3: Import test to verify import functionality still works
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportState:              true,
				ImportStateVerify:        true,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					resourceState := state.RootModule().Resources[resourceName]
					return fmt.Sprintf("%s/%s", resourceState.Primary.Attributes["zone_id"], resourceState.Primary.ID), nil
				},
			},
		},
	})
}
