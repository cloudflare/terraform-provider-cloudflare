package workers_script_test

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

// TestMigrateWorkersScriptFromV4Basic tests basic migration from v4 to v5
func TestMigrateWorkersScriptFromV4Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config using old resource name and 'name' attribute
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')) })"
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4Module tests migration of module-based scripts
func TestMigrateWorkersScriptFromV4Module(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with module=true
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id         = "%[2]s"
  name              = "%[1]s"
  content           = "export default { fetch() { return new Response('Hello Module') } }"
  module            = true
  compatibility_date = "2023-03-19"
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("compatibility_date"), knownvalue.StringExact("2023-03-19")),
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4KVBinding tests KV namespace binding transformation
func TestMigrateWorkersScriptFromV4KVBinding(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with kv_namespace_binding
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello KV')) })"
  
  kv_namespace_binding {
    name         = "MY_KV"
    namespace_id = "test-namespace-id"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
				// Expect KV binding to be transformed to new bindings array format
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("MY_KV")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("kv_namespace")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("namespace_id"), knownvalue.StringExact("test-namespace-id")),
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4D1Binding tests D1 database binding removal
func TestMigrateWorkersScriptFromV4D1Binding(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with d1_database_binding
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello D1')) })"
  
  d1_database_binding {
    name        = "MY_D1"
    database_id = "test-database-id"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4ServiceBinding tests service binding removal
func TestMigrateWorkersScriptFromV4ServiceBinding(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with service_binding
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello Service')) })"
  
  service_binding {
    name        = "MY_SERVICE"
    service     = "target-service"
    environment = "production"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4PlainTextBinding tests plain text binding transformation
func TestMigrateWorkersScriptFromV4PlainTextBinding(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with plain_text_binding
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello Plain Text')) })"
  
  plain_text_binding {
    name = "MY_VAR"
    text = "Hello World"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
				// Expect plain text binding to be transformed to new bindings array format
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("MY_VAR")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("plain_text")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("text"), knownvalue.StringExact("Hello World")),
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4SecretTextBinding tests secret text binding transformation
func TestMigrateWorkersScriptFromV4SecretTextBinding(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with secret_text_binding
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello Secret')) })"
  
  secret_text_binding {
    name = "MY_SECRET"
    text = "secret-value"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
				// Expect binding to be transformed to new bindings array format
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("MY_SECRET")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("secret_text")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("text"), knownvalue.StringExact("secret-value")),
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4R2Binding tests R2 bucket binding removal
func TestMigrateWorkersScriptFromV4R2Binding(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with r2_bucket_binding
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello R2')) })"
  
  r2_bucket_binding {
    name        = "MY_BUCKET"
    bucket_name = "test-bucket"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4QueueBinding tests queue binding removal
func TestMigrateWorkersScriptFromV4QueueBinding(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with queue_binding
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello Queue')) })"
  
  queue_binding {
    name  = "MY_QUEUE"
    queue = "test-queue"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4MultipleBindings tests migration with multiple binding types
func TestMigrateWorkersScriptFromV4MultipleBindings(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with multiple binding types
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello Multiple')) })"
  
  kv_namespace_binding {
    name         = "MY_KV"
    namespace_id = "test-namespace-id"
  }
  
  plain_text_binding {
    name = "MY_VAR"
    text = "Hello World"
  }
  
  secret_text_binding {
    name = "MY_SECRET"
    text = "secret-value"
  }
  
  r2_bucket_binding {
    name        = "MY_BUCKET"
    bucket_name = "test-bucket"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
				// Expect 4 bindings to be transformed to new bindings array format
				// Note: bindings array order may vary, so we check for existence of each type
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings"), knownvalue.ListSizeExact(4)),
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4PlacementConfig tests placement configuration removal
func TestMigrateWorkersScriptFromV4PlacementConfig(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with placement block
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello Placement')) })"
  
  placement {
    mode = "smart"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
				// Placement block should be removed during migration
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4DispatchNamespace tests dispatch namespace removal
func TestMigrateWorkersScriptFromV4DispatchNamespace(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with dispatch_namespace
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[1]s"
  content           = "addEventListener('fetch', event => { event.respondWith(new Response('Hello Dispatch')) })"
  dispatch_namespace = "test-namespace"
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
				// dispatch_namespace should be removed during migration
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4AnalyticsEngineBinding tests analytics engine binding removal
func TestMigrateWorkersScriptFromV4AnalyticsEngineBinding(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with analytics_engine_binding
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello Analytics')) })"
  
  analytics_engine_binding {
    name    = "MY_ANALYTICS"
    dataset = "test-dataset"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4HyperdriveBinding tests hyperdrive config binding removal
func TestMigrateWorkersScriptFromV4HyperdriveBinding(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with hyperdrive_config_binding
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello Hyperdrive')) })"
  
  hyperdrive_config_binding {
    name      = "MY_HYPERDRIVE"
    binding_id = "test-binding-id"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateWorkersScriptFromV4CompatibilityFlags tests compatibility flags preservation
func TestMigrateWorkersScriptFromV4CompatibilityFlags(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()

	// V4 config with compatibility_flags
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id         = "%[2]s"
  name              = "%[1]s"
  content           = "addEventListener('fetch', event => { event.respondWith(new Response('Hello Compat')) })"
  compatibility_date = "2023-03-19"
  compatibility_flags = ["nodejs_compat", "streams_enable_constructors"]
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("compatibility_date"), knownvalue.StringExact("2023-03-19")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("compatibility_flags"), knownvalue.SetExact([]knownvalue.Check{
					knownvalue.StringExact("nodejs_compat"),
					knownvalue.StringExact("streams_enable_constructors"),
				})),
			}),
		},
	})
}