package workers_kv_test

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

// TestMigrateWorkersKV_Basic tests migration of a basic Workers KV resource from v4 to v5
// Main change: key field renamed to key_name
func TestMigrateWorkersKV_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	namespaceName := fmt.Sprintf("tf-test-ns-%s", rnd)
	kvName := fmt.Sprintf("tf-test-kv-%s", rnd)
	keyName := "test_key"
	value := "test_value"
	tmpDir := t.TempDir()

	// V4 config using 'key' field
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s" {
  account_id = "%[2]s"
  title      = "%[3]s"
}

resource "cloudflare_workers_kv" "%[4]s" {
  account_id   = "%[2]s"
  namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
  key          = "%[5]s"
  value        = "%[6]s"
}`, namespaceName, accountID, namespaceName, kvName, keyName, value)

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
				// Verify the key field was renamed to key_name
				statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("key_name"), knownvalue.StringExact(keyName)),
				statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("value"), knownvalue.StringExact(value)),
				statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				// Verify namespace_id is preserved
				statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("namespace_id"), knownvalue.NotNull()),
				// Verify id field is preserved (it's the same as key_name)
				statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("id"), knownvalue.StringExact(keyName)),
			}),
		},
	})
}

// TestMigrateWorkersKV_SpecialCharacters tests migration with URL-encoded special characters
func TestMigrateWorkersKV_SpecialCharacters(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	namespaceName := fmt.Sprintf("tf-test-ns-%s", rnd)
	kvName := fmt.Sprintf("tf-test-kv-%s", rnd)
	// URL-encoded key with special characters
	keyName := "api/token/key"
	value := `{"api_key": "test123", "endpoint": "https://api.example.com"}`
	tmpDir := t.TempDir()

	// V4 config using 'key' field with special characters
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s" {
  account_id = "%[2]s"
  title      = "%[3]s"
}

resource "cloudflare_workers_kv" "%[4]s" {
  account_id   = "%[2]s"
  namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
  key          = "%[5]s"
  value        = %[6]q
}`, namespaceName, accountID, namespaceName, kvName, keyName, value)

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
				// Verify the key field was renamed to key_name and special characters preserved
				statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("key_name"), knownvalue.StringExact(keyName)),
				statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("value"), knownvalue.StringExact(value)),
				statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			}),
		},
	})
}

// TestMigrateWorkersKV_EmptyValue tests migration with an empty string value
func TestMigrateWorkersKV_EmptyValue(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	namespaceName := fmt.Sprintf("tf-test-ns-%s", rnd)
	kvName := fmt.Sprintf("tf-test-kv-%s", rnd)
	keyName := "empty_key"
	value := ""
	tmpDir := t.TempDir()

	// V4 config with empty value
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s" {
  account_id = "%[2]s"
  title      = "%[3]s"
}

resource "cloudflare_workers_kv" "%[4]s" {
  account_id   = "%[2]s"
  namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
  key          = "%[5]s"
  value        = "%[6]s"
}`, namespaceName, accountID, namespaceName, kvName, keyName, value)

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
				// Verify empty value is preserved
				statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("key_name"), knownvalue.StringExact(keyName)),
				statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("value"), knownvalue.StringExact("")),
				statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			}),
		},
	})
}
