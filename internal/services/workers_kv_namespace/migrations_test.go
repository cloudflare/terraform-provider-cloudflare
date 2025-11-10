package workers_kv_namespace_test

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

// TestMigrateWorkersKVNamespaceBasic tests basic migration from v4 to v5
func TestMigrateWorkersKVNamespaceBasic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_kv_namespace." + rnd
	tmpDir := t.TempDir()
	title := fmt.Sprintf("test-kv-namespace-%s", rnd)

	// V4 config - simple pass-through migration
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s" {
  account_id = "%[2]s"
  title      = "%[3]s"
}`, rnd, accountID, title)

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
				// Verify resource exists with same type (no rename)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(title)),
				// Verify new computed field is present in v5
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("supports_url_encoding"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateWorkersKVNamespaceWithSpecialChars tests migration with special characters in title
func TestMigrateWorkersKVNamespaceWithSpecialChars(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_kv_namespace." + rnd
	tmpDir := t.TempDir()
	// Title with spaces, dashes, and underscores
	title := fmt.Sprintf("Test KV Namespace_%s-2024", rnd)

	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s" {
  account_id = "%[2]s"
  title      = "%[3]s"
}`, rnd, accountID, title)

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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(title)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("supports_url_encoding"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateWorkersKVNamespaceMultiple tests migration of multiple KV namespaces in one config
func TestMigrateWorkersKVNamespaceMultiple(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	rnd1 := rnd + "1"
	rnd2 := rnd + "2"
	resourceName1 := "cloudflare_workers_kv_namespace." + rnd1
	resourceName2 := "cloudflare_workers_kv_namespace." + rnd2
	tmpDir := t.TempDir()
	title1 := fmt.Sprintf("test-kv-namespace-1-%s", rnd)
	title2 := fmt.Sprintf("test-kv-namespace-2-%s", rnd)

	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s" {
  account_id = "%[3]s"
  title      = "%[4]s"
}

resource "cloudflare_workers_kv_namespace" "%[2]s" {
  account_id = "%[3]s"
  title      = "%[5]s"
}`, rnd1, rnd2, accountID, title1, title2)

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
				// Verify first namespace
				statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("title"), knownvalue.StringExact(title1)),
				statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("supports_url_encoding"), knownvalue.NotNull()),
				// Verify second namespace
				statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("title"), knownvalue.StringExact(title2)),
				statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("supports_url_encoding"), knownvalue.NotNull()),
			}),
		},
	})
}
