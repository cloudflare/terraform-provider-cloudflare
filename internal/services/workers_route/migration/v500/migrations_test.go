package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion
)

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_no_script.tf
var v4NoScriptConfig string

//go:embed testdata/v5_no_script.tf
var v5NoScriptConfig string

// ============================================================================
// Ported from migrations_test.go (service root) — original v4-only tests
// ============================================================================

// TestMigrateWorkersRoute_FromV4Basic tests basic migration: script_name→script.
// Ported from TestMigrateWorkersRouteMigrationFromV4Basic.
func TestMigrateWorkersRoute_FromV4Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_route." + rnd
	tmpDir := t.TempDir()
	scriptName := fmt.Sprintf("test-script-%s", rnd)
	version := acctest.GetLastV4Version()
	testConfig := fmt.Sprintf(v4BasicConfig, rnd, zoneID, accountID, scriptName)
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.cfapi.net/*", rnd))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.StringExact(scriptName)),
			}),
		},
	})
}

// TestMigrateWorkersRoute_FromV4NoScript tests migration of route without script.
// Ported from TestMigrateWorkersRouteMigrationFromV4NoScript.
func TestMigrateWorkersRoute_FromV4NoScript(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_route." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	testConfig := fmt.Sprintf(v4NoScriptConfig, rnd, zoneID)
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

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
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.cfapi.net/api/*", rnd))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.Null()),
			}),
		},
	})
}

// TestMigrateWorkersRoute_FromV4SingleResource tests migration using singular resource name.
// Ported from TestMigrateWorkersRouteMigrationFromV4SingleResource.
// Uses cloudflare_worker_route (singular) — tests the resource rename path via moved block.
func TestMigrateWorkersRoute_FromV4SingleResource(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_route." + rnd
	tmpDir := t.TempDir()
	scriptName := fmt.Sprintf("test-script-%s", rnd)
	version := acctest.GetLastV4Version()

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

	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.cfapi.net/v1/*", rnd))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.StringExact(scriptName)),
			}),
		},
	})
}

// ============================================================================
// New dual test cases (from_v4_latest / from_v5)
// ============================================================================

// TestMigrateWorkersRoute_Basic tests basic migration with dual test pattern.
func TestMigrateWorkersRoute_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, accountID, scriptName string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, accountID, scriptName string) string { return fmt.Sprintf(v4BasicConfig, rnd, zoneID, accountID, scriptName) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, zoneID, accountID, scriptName string) string { return fmt.Sprintf(v5BasicConfig, rnd, zoneID, accountID, scriptName) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			scriptName := fmt.Sprintf("test-script-%s", rnd)
			resourceName := "cloudflare_workers_route." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID, accountID, scriptName)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
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

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.cfapi.net/*", rnd))),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.StringExact(scriptName)),
					}),
				},
			})
		})
	}
}

// TestMigrateWorkersRoute_NoScript tests migration of route without script, dual pattern.
func TestMigrateWorkersRoute_NoScript(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4NoScriptConfig, rnd, zoneID) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v5NoScriptConfig, rnd, zoneID) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_workers_route." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
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

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.cfapi.net/api/*", rnd))),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.Null()),
					}),
				},
			})
		})
	}
}
