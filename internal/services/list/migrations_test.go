package list_test

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

/* Migration tests don't include every possible permutation, but do cover:
 * - At least one migration from v4 and one from v5 for each identifier
 * - TODO: Add specific test coverage notes
 */

func TestMigrateListMigrationFromV4BasicIP(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_list." + rnd
	tmpDir := t.TempDir()

	// V4 config
	v4Config := fmt.Sprintf(`
resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Test list created with v4 provider"
  kind        = "ip"

  item {
    value {
      ip = "192.0.2.1"
    }
    comment = "Test IP 1"
  }

  item {
    value {
      ip = "192.0.2.2"
    }
    comment = "Test IP 2"
  }
}`, rnd, accountID)

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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("ip")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Test list created with v4 provider")),
				// Verify items migrated correctly
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("ip"), knownvalue.StringExact("192.0.2.1")),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("comment"), knownvalue.StringExact("Test IP 1")),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_1", rnd), tfjsonpath.New("ip"), knownvalue.StringExact("192.0.2.2")),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_1", rnd), tfjsonpath.New("comment"), knownvalue.StringExact("Test IP 2")),
			}),
		},
	})
}

func TestMigrateListMigrationFromV4BasicASN(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_list." + rnd
	tmpDir := t.TempDir()

	// V4 config
	v4Config := fmt.Sprintf(`
resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Test list created with v4 provider"
  kind        = "asn"

  item {
    value {
      asn = 12456
    }
    comment = "Test asn 1"
  }

  item {
    value {
      asn = 789
    }
    comment = "Test ASN 2"
  }
}`, rnd, accountID)

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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("asn")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Test list created with v4 provider")),
				// Verify items migrated correctly
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("asn"), knownvalue.Int64Exact(789)),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("comment"), knownvalue.StringExact("Test ASN 2")),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_1", rnd), tfjsonpath.New("asn"), knownvalue.Int64Exact(12456)),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_1", rnd), tfjsonpath.New("comment"), knownvalue.StringExact("Test asn 1")),
			}),
		},
	})
}

func TestMigrateListMigrationFromV4BasicHostname(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_list." + rnd
	tmpDir := t.TempDir()

	// V4 config
	v4Config := fmt.Sprintf(`
resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Test list created with v4 provider"
  kind        = "hostname"

  item {
    value {
      hostname {
        url_hostname = "example.com"
      }
    }
    comment = "one"
  }
}`, rnd, accountID)

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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("hostname")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Test list created with v4 provider")),
				// Verify items migrated correctly
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("hostname"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("hostname").AtMapKey("url_hostname"), knownvalue.StringExact("example.com")),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("comment"), knownvalue.StringExact("one")),
			}),
		},
	})
}

func TestMigrateListMigrationFromV4BasicRedirect(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_list." + rnd
	tmpDir := t.TempDir()

	// V4 config
	v4Config := fmt.Sprintf(`
resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Test list created with v4 provider"
  kind        = "redirect"

  item {
	value {
	  redirect {
		include_subdomains    = "enabled"
		preserve_path_suffix  = "disabled"
		preserve_query_string = "enabled"
		source_url            = "example.com/foo"
		status_code           = 301
		subpath_matching      = "enabled"
		target_url            = "https://foo.example.com"
	  }
	}
	comment = "one"
  }
}`, rnd, accountID)

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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("redirect")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Test list created with v4 provider")),
				// Verify items migrated correctly
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("redirect"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("redirect").AtMapKey("include_subdomains"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("redirect").AtMapKey("preserve_path_suffix"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("redirect").AtMapKey("preserve_query_string"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("redirect").AtMapKey("subpath_matching"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("redirect").AtMapKey("source_url"), knownvalue.StringExact("example.com/foo")),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("redirect").AtMapKey("status_code"), knownvalue.Int32Exact(301)),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_list_item.%s_item_0", rnd), tfjsonpath.New("redirect").AtMapKey("target_url"), knownvalue.StringExact("https://foo.example.com")),
			}),
		},
	})
}
