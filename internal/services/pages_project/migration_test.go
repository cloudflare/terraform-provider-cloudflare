package pages_project_test

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

// TestMigratePagesProject_V4ToV5_Minimal tests basic migration with minimal config
func TestMigratePagesProject_V4ToV5_Minimal(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	projectName := fmt.Sprintf("tf-test-minimal-%s", rnd)

	// V4 config with only required fields
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"
}`, rnd, accountID, projectName)

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
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists (type unchanged)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("production_branch"), knownvalue.StringExact("main")),
				// Verify computed fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("subdomain"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigratePagesProject_V4ToV5_WithBuildConfig tests build_config block to attribute conversion
func TestMigratePagesProject_V4ToV5_WithBuildConfig(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	projectName := fmt.Sprintf("tf-test-build-%s", rnd)

	// V4 config with build_config block
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  build_config {
    build_command   = "npm run build"
    destination_dir = "public"
    root_dir        = "/"
  }
}`, rnd, accountID, projectName)

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
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
				// Verify build_config converted from block to object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("build_command"), knownvalue.StringExact("npm run build")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("destination_dir"), knownvalue.StringExact("public")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("root_dir"), knownvalue.StringExact("/")),
			}),
		},
	})
}

// TestMigratePagesProject_V4ToV5_WithSourceAndConfig tests source/config blocks and field rename
// NOTE: This test requires a valid GitHub repository. Skipping because API validates repository existence.

func TestMigratePagesProject_V4ToV5_WithSourceAndConfig(t *testing.T) {
	t.Skip("Skipping test that requires valid GitHub repository - API validates repository existence")

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	projectName := fmt.Sprintf("tf-test-source-%s", rnd)
	repoName := fmt.Sprintf("test-repo-%s", rnd)

	// V4 config with source and config blocks
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  source {
    type = "github"
    config = {
      owner                         = "cloudflare"
      repo_name                     = "%[4]s"
      production_branch             = "main"
      production_deployment_enabled = true
      pr_comments_enabled           = true
    }
  }
}`, rnd, accountID, projectName, repoName)

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
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
				// Verify source converted to object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("source"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("source").AtMapKey("type"), knownvalue.StringExact("github")),
				// Verify config converted to object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("source").AtMapKey("config"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("source").AtMapKey("config").AtMapKey("owner"), knownvalue.StringExact("cloudflare")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("source").AtMapKey("config").AtMapKey("repo_name"), knownvalue.StringExact(repoName)),
				// Verify field rename: production_deployment_enabled → production_deployments_enabled
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("source").AtMapKey("config").AtMapKey("production_deployments_enabled"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("source").AtMapKey("config").AtMapKey("pr_comments_enabled"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigratePagesProject_V4ToV5_WithDeploymentConfigsBasic tests deployment_configs with placement
func TestMigratePagesProject_V4ToV5_WithDeploymentConfigsBasic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	projectName := fmt.Sprintf("tf-test-deploy-%s", rnd)

	// V4 config with deployment_configs and nested placement
	// Note: Including preview config to match API behavior (API creates both preview and production)
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  deployment_configs {
    preview {
      compatibility_date = "2024-01-01"
      placement {
        mode = "smart"
      }
    }
    production {
      compatibility_date = "2024-01-01"
      placement {
        mode = "smart"
      }
    }
  }
}`, rnd, accountID, projectName)

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
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
				// Verify deployment_configs converted to object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs"), knownvalue.NotNull()),
				// Verify preview converted with v5 defaults (migration sets v5 defaults when not specified)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("compatibility_date"), knownvalue.StringExact("2024-01-01")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("usage_model"), knownvalue.StringExact("bundled")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("fail_open"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("placement"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("placement").AtMapKey("mode"), knownvalue.StringExact("smart")),
				// Verify production converted with v5 defaults
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("compatibility_date"), knownvalue.StringExact("2024-01-01")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("usage_model"), knownvalue.StringExact("bundled")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("fail_open"), knownvalue.Bool(false)),
				// Verify placement converted to object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("placement"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("placement").AtMapKey("mode"), knownvalue.StringExact("smart")),
			}),
		},
	})
}

// TestMigratePagesProject_V4ToV5_WithEnvironmentVariables tests env vars + secrets merge
// NOTE: Skipping because v4 provider doesn't properly store environment_variables and secrets
// within deployment_configs attribute in state
func TestMigratePagesProject_V4ToV5_WithEnvironmentVariables(t *testing.T) {
	t.Skip("Skipping test - v4 provider doesn't store environment_variables/secrets in deployment_configs attribute properly")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	projectName := fmt.Sprintf("tf-test-envvars-%s", rnd)

	// V4 config with environment_variables and secrets
	// Note: Including preview config to match API behavior (API creates both preview and production)
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  deployment_configs {
    preview = {
      compatibility_date = "2024-01-01"
    }
    production {
      environment_variables = {
        NODE_ENV = "production"
        API_URL  = "https://api.example.com"
      }
      secrets {
        API_KEY     = "secret123"
        DB_PASSWORD = "dbpass456"
      }
      placement {
        mode = "smart"
      }
    }
  }
}`, rnd, accountID, projectName)

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
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
				// Verify deployment_configs converted
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production"), knownvalue.NotNull()),
				// Verify env_vars exists (merged from environment_variables + secrets)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.NotNull()),
				// Verify environment_variables merged as plain_text
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("NODE_ENV"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("NODE_ENV").AtMapKey("type"), knownvalue.StringExact("plain_text")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("NODE_ENV").AtMapKey("value"), knownvalue.StringExact("production")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("API_URL"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("API_URL").AtMapKey("type"), knownvalue.StringExact("plain_text")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("API_URL").AtMapKey("value"), knownvalue.StringExact("https://api.example.com")),
				// Verify secrets merged as secret_text
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("API_KEY"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("API_KEY").AtMapKey("type"), knownvalue.StringExact("secret_text")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("DB_PASSWORD"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("DB_PASSWORD").AtMapKey("type"), knownvalue.StringExact("secret_text")),
			}),
		},
	})
}

// TestMigratePagesProject_V4ToV5_WithBindings tests TypeMap wrapping and service_binding conversion
func TestMigratePagesProject_V4ToV5_WithBindings(t *testing.T) {
	// This test requires actual KV namespace, D1 database, R2 bucket, and worker IDs
	// Skip if we don't have the necessary test resources
	kvNamespaceID := os.Getenv("TEST_CLOUDFLARE_KV_NAMESPACE_ID")
	d1DatabaseID := os.Getenv("TEST_CLOUDFLARE_D1_DATABASE_ID")
	workerName := os.Getenv("TEST_CLOUDFLARE_WORKER_NAME")

	if kvNamespaceID == "" || d1DatabaseID == "" || workerName == "" {
		t.Skip("Skipping binding test - requires TEST_CLOUDFLARE_KV_NAMESPACE_ID, TEST_CLOUDFLARE_D1_DATABASE_ID, and TEST_CLOUDFLARE_WORKER_NAME")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	projectName := fmt.Sprintf("tf-test-bindings-%s", rnd)
	bucketName := fmt.Sprintf("test-bucket-%s", rnd)

	// V4 config with bindings (kv_namespaces, d1_databases, r2_buckets, service_binding)
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  deployment_configs {
    production {
      kv_namespaces = {
        MY_KV = "%[4]s"
      }
      d1_databases = {
        MY_DB = "%[5]s"
      }
      r2_buckets = {
        MY_BUCKET = "%[6]s"
      }
      service_binding {
        name        = "MY_SERVICE"
        service     = "%[7]s"
        environment = "production"
      }
    }
  }
}`, rnd, accountID, projectName, kvNamespaceID, d1DatabaseID, bucketName, workerName)

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
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
				// Verify deployment_configs converted
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production"), knownvalue.NotNull()),
				// Verify kv_namespaces wrapped: "MY_KV": "id" → "MY_KV": {"namespace_id": "id"}
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces").AtMapKey("MY_KV"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces").AtMapKey("MY_KV").AtMapKey("namespace_id"), knownvalue.StringExact(kvNamespaceID)),
				// Verify d1_databases wrapped: "MY_DB": "id" → "MY_DB": {"id": "id"}
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("d1_databases"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("d1_databases").AtMapKey("MY_DB"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("d1_databases").AtMapKey("MY_DB").AtMapKey("id"), knownvalue.StringExact(d1DatabaseID)),
				// Verify r2_buckets wrapped: "MY_BUCKET": "name" → "MY_BUCKET": {"name": "name"}
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("r2_buckets"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("r2_buckets").AtMapKey("MY_BUCKET"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("r2_buckets").AtMapKey("MY_BUCKET").AtMapKey("name"), knownvalue.StringExact(bucketName)),
				// Verify service_binding converted to services map
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("services"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("services").AtMapKey("MY_SERVICE"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("services").AtMapKey("MY_SERVICE").AtMapKey("service"), knownvalue.StringExact(workerName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("services").AtMapKey("MY_SERVICE").AtMapKey("environment"), knownvalue.StringExact("production")),
			}),
		},
	})
}

// TestMigratePagesProject_V4ToV5_FullResource tests complete resource with all features
func TestMigratePagesProject_V4ToV5_FullResource(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	projectName := fmt.Sprintf("tf-test-full-%s", rnd)

	// V4 config with multiple features (build_config and deployment_configs)
	// Note: Removed source block because it requires valid GitHub repository
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  build_config {
    build_command   = "npm run build"
    destination_dir = "dist"
  }

  deployment_configs {
    preview {
      compatibility_date = "2024-01-01"
    }
    production {
      compatibility_date = "2024-01-01"
      placement {
        mode = "smart"
      }
    }
  }
}`, rnd, accountID, projectName)

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
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
				// Verify build_config
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("build_command"), knownvalue.StringExact("npm run build")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("destination_dir"), knownvalue.StringExact("dist")),
				// Verify deployment_configs with v5 defaults
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs"), knownvalue.NotNull()),
				// Verify preview with v5 defaults
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("usage_model"), knownvalue.StringExact("bundled")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("fail_open"), knownvalue.Bool(false)),
				// Verify production with v5 defaults
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("compatibility_date"), knownvalue.StringExact("2024-01-01")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("usage_model"), knownvalue.StringExact("bundled")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("fail_open"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("placement"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("placement").AtMapKey("mode"), knownvalue.StringExact("smart")),
			}),
		},
	})
}
