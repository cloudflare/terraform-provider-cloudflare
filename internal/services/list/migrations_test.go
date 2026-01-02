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

// TestMigrateCloudflareListFromV4WithIPItems tests migration with IP item blocks
func TestMigrateCloudflareListFromV4WithIPItems(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
	resourceName := "cloudflare_list." + rnd
	tmpDir := t.TempDir()
	description := fmt.Sprintf("Test list %s", rnd)

	// V4 config with IP items
	v4Config := fmt.Sprintf(`
resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  kind        = "ip"
  description = "%[4]s"
  
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
}`, rnd, accountID, listName, description)

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
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("ip")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(description)),
				// Verify items were migrated to items array
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListSizeExact(2)),
				// Verify num_items is set correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("num_items"), knownvalue.Float64Exact(2)),
				// Check first item
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("ip"), knownvalue.StringExact("192.0.2.1")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("comment"), knownvalue.StringExact("Test IP 1")),
				// Check second item
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("ip"), knownvalue.StringExact("192.0.2.2")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("comment"), knownvalue.StringExact("Test IP 2")),
			}),
		},
	})
}

// TestMigrateCloudflareListFromV4WithASNItems tests migration with ASN item blocks
func TestMigrateCloudflareListFromV4WithASNItems(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
	resourceName := "cloudflare_list." + rnd
	tmpDir := t.TempDir()
	description := fmt.Sprintf("Test ASN list %s", rnd)

	// V4 config with ASN items
	v4Config := acctest.LoadTestCase("cloudflarelistv4withasnitems.tf", rnd, accountID, listName, description)

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
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("asn")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(description)),
				// Verify items were migrated to items array
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListSizeExact(2)),
				// Check first item
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("asn"), knownvalue.Int64Exact(12345)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("comment"), knownvalue.StringExact("Test ASN 1")),
				// Check second item (no comment)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("asn"), knownvalue.Int64Exact(67890)),
			}),
		},
	})
}

// TestMigrateCloudflareListFromV4WithHostnameItems tests migration with hostname item blocks
func TestMigrateCloudflareListFromV4WithHostnameItems(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
	resourceName := "cloudflare_list." + rnd
	tmpDir := t.TempDir()
	description := fmt.Sprintf("Test hostname list %s", rnd)

	// V4 config with hostname items
	v4Config := acctest.LoadTestCase("cloudflarelistv4withhostnameitems.tf", rnd, accountID, listName, description)

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
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("hostname")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(description)),
				// Verify items were migrated to items array
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListSizeExact(2)),
				// Check first item
				statecheck.ExpectKnownValue(
					resourceName,
					tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("hostname").AtMapKey("url_hostname"),
					knownvalue.StringExact("example.com"),
				),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("comment"), knownvalue.StringExact("Test hostname 1")),
				// Check second item
				statecheck.ExpectKnownValue(
					resourceName,
					tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("hostname").AtMapKey("url_hostname"),
					knownvalue.StringExact("test.example.com"),
				),
			}),
		},
	})
}

// TestMigrateCloudflareListFromV4WithRedirectItems tests migration with redirect item blocks including boolean conversion
func TestMigrateCloudflareListFromV4WithRedirectItems(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
	resourceName := "cloudflare_list." + rnd
	tmpDir := t.TempDir()
	description := fmt.Sprintf("Test redirect list %s", rnd)

	// V4 config with redirect items
	v4Config := acctest.LoadTestCase("cloudflarelistv4withredirectitems.tf", rnd, accountID, listName, description)

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
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("redirect")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(description)),
				// Verify items were migrated to items array
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListSizeExact(1)),
				// Check redirect item
				statecheck.ExpectKnownValue(
					resourceName,
					tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("source_url"),
					knownvalue.StringExact("https://example.com/old"),
				),
				statecheck.ExpectKnownValue(
					resourceName,
					tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("target_url"),
					knownvalue.StringExact("https://example.com/new"),
				),
				statecheck.ExpectKnownValue(
					resourceName,
					tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("status_code"),
					knownvalue.Int64Exact(301),
				),
				// Verify boolean conversions from "enabled"/"disabled" to true/false
				statecheck.ExpectKnownValue(
					resourceName,
					tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("include_subdomains"),
					knownvalue.Bool(true),
				),
				statecheck.ExpectKnownValue(
					resourceName,
					tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("subpath_matching"),
					knownvalue.Bool(false),
				),
				statecheck.ExpectKnownValue(
					resourceName,
					tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("preserve_query_string"),
					knownvalue.Bool(true),
				),
				statecheck.ExpectKnownValue(
					resourceName,
					tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("preserve_path_suffix"),
					knownvalue.Bool(false),
				),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("comment"), knownvalue.StringExact("Test redirect")),
			}),
		},
	})
}

// TestMigrateCloudflareListFromV4WithDynamicItems tests migration with dynamic item blocks
func TestMigrateCloudflareListFromV4WithDynamicItems(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
	resourceName := "cloudflare_list." + rnd
	tmpDir := t.TempDir()
	description := fmt.Sprintf("Test dynamic list %s", rnd)

	// V4 config with dynamic items
	v4Config := acctest.LoadTestCase("cloudflarelistv4withdynamicitems.tf", rnd, accountID, listName, description)

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
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("ip")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(description)),
				// Verify items were migrated from dynamic blocks to items array
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListSizeExact(2)),
				// Check dynamic items
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("ip"), knownvalue.StringExact("10.0.0.1")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("comment"), knownvalue.StringExact("Dynamic IP 1")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("ip"), knownvalue.StringExact("10.0.0.2")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("comment"), knownvalue.StringExact("Dynamic IP 2")),
			}),
		},
	})
}

// TestMigrateCloudflareListFromV4WithSeparateListItems tests migration with separate cloudflare_list_item resources
func TestMigrateCloudflareListFromV4WithSeparateListItems(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
	resourceName := "cloudflare_list." + rnd
	tmpDir := t.TempDir()
	description := fmt.Sprintf("Test list with separate items %s", rnd)

	// V4 config with separate list_item resources
	v4Config := acctest.LoadTestCase("cloudflarelistv4withseparatelistitems.tf", rnd, accountID, listName, description)

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
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("ip")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(description)),
				// Verify separate list_item resources were merged into items array
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListSizeExact(2)),
				// Check merged items
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("ip"), knownvalue.StringExact("172.16.0.1")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("comment"), knownvalue.StringExact("Separate item 1")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("ip"), knownvalue.StringExact("172.16.0.2")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("comment"), knownvalue.StringExact("Separate item 2")),
			}),
		},
	})
}
