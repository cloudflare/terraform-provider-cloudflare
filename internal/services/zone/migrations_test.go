package zone_test

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

// TestMigrateZoneMigrationFromV4Basic tests basic migration from v4 to v5
func TestMigrateZoneMigrationFromV4Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
	resourceName := "cloudflare_zone." + rnd
	tmpDir := t.TempDir()

	// V4 config using old attribute names
	v4Config := fmt.Sprintf(`
resource "cloudflare_zone" "%[1]s" {
  zone       = "%[2]s"
  account_id = "%[3]s"
}`, rnd, zoneName, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
				// Verify v4 attributes are removed/transformed
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
				// Verify computed attributes are present
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZoneMigrationFromV4Complete tests migration with all v4 attributes
func TestMigrateZoneMigrationFromV4Complete(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
	resourceName := "cloudflare_zone." + rnd
	tmpDir := t.TempDir()

	// V4 config with all optional attributes including removed ones
	// Use "full" type to avoid API restrictions with partial zones under cfapi.net
	v4Config := fmt.Sprintf(`
resource "cloudflare_zone" "%[1]s" {
  zone                = "%[2]s"
  account_id          = "%[3]s"
  paused              = true
  type                = "full"
  jump_start          = true
  plan                = "enterprise"
  vanity_name_servers = ["ns1.%[2]s", "ns2.%[2]s"]
}`, rnd, zoneName, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				// Verify v4→v5 transformations
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
				// Verify preserved attributes
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("paused"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("full")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("vanity_name_servers"), knownvalue.ListSizeExact(2)),
			}),
		},
	})
}

// TestMigrateZoneMigrationFromV4PlanFree tests migration with removed plan attribute
func TestMigrateZoneMigrationFromV4PlanFree(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
	resourceName := "cloudflare_zone." + rnd
	tmpDir := t.TempDir()

	// V4 config with plan that gets removed
	v4Config := fmt.Sprintf(`
resource "cloudflare_zone" "%[1]s" {
  zone       = "%[2]s"
  account_id = "%[3]s"
  plan       = "free"
  jump_start = false
}`, rnd, zoneName, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
			// Step 2: Run migration and verify removed attributes
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
				// V5 has computed plan object instead of configurable string
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("plan"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZoneMigrationFromV4Unicode tests migration with unicode domain names
func TestMigrateZoneMigrationFromV4Unicode(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zone." + rnd
	tmpDir := t.TempDir()

	// V4 config with unicode domain - use a working unicode domain under cfapi.net
	unicodeDomain := fmt.Sprintf("テスト-%s.cfapi.net", rnd)
	v4Config := fmt.Sprintf(`
resource "cloudflare_zone" "%[1]s" {
  zone       = "%[3]s"
  account_id = "%[2]s"
  type       = "full"
}`, rnd, accountID, unicodeDomain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
			// Step 2: Run migration and verify unicode handling
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(unicodeDomain)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("full")),
			}),
		},
	})
}

// TestMigrateZoneMigrationFromV4Secondary tests migration with secondary zone type
func TestMigrateZoneMigrationFromV4Secondary(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
	resourceName := "cloudflare_zone." + rnd
	tmpDir := t.TempDir()

	// V4 config with secondary zone type
	v4Config := fmt.Sprintf(`
resource "cloudflare_zone" "%[1]s" {
  zone       = "%[2]s"
  account_id = "%[3]s"
  type       = "secondary"
  paused     = false
}`, rnd, zoneName, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
			// Step 2: Run migration and verify secondary zone handling
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("secondary")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("paused"), knownvalue.Bool(false)),
			}),
		},
	})
}

// TestMigrateZoneMigrationFromV4MetaStructure tests migration handling of meta field changes
func TestMigrateZoneMigrationFromV4MetaStructure(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
	resourceName := "cloudflare_zone." + rnd
	tmpDir := t.TempDir()

	// Basic V4 config - meta is computed so we just need to verify it works
	v4Config := fmt.Sprintf(`
resource "cloudflare_zone" "%[1]s" {
  zone       = "%[2]s"
  account_id = "%[3]s"
  type       = "full"
}`, rnd, zoneName, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
			// Step 2: Run migration and verify meta structure
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
				// Verify meta is a structured object (v5 format)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("meta"), knownvalue.NotNull()),
			}),
		},
	})
}