package v500_test

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_project/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

//go:embed testdata/v4_minimal.tf
var v4MinimalTF string

//go:embed testdata/v4_build_config.tf
var v4BuildConfigTF string

//go:embed testdata/v4_build_config_complete.tf
var v4BuildConfigCompleteTF string

//go:embed testdata/v4_deployment_configs.tf
var v4DeploymentConfigsTF string

//go:embed testdata/v4_compatibility_flags.tf
var v4CompatibilityFlagsTF string

//go:embed testdata/v4_full.tf
var v4FullTF string

//go:embed testdata/v4_always_use_latest.tf
var v4AlwaysUseLatestTF string

//go:embed testdata/v4_service_bindings.tf
var v4ServiceBindingsTF string

// TestMigratePagesProject_V4ToV5_Minimal tests basic migration with minimal config
func TestMigratePagesProject_V4ToV5_Minimal(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-minimal-%s", rnd)

	v4Config := fmt.Sprintf(v4MinimalTF, rnd, accountID, projectName)

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
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
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
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-build-%s", rnd)

	v4Config := fmt.Sprintf(v4BuildConfigTF, rnd, accountID, projectName)

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
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
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
func TestMigratePagesProject_V4ToV5_WithSourceAndConfig(t *testing.T) {
	t.Skip("Skipping test that requires valid GitHub repository - API validates repository existence")

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
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
    config {
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
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
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
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-deploy-%s", rnd)

	v4Config := fmt.Sprintf(v4DeploymentConfigsTF, rnd, accountID, projectName)

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
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
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
// within deployment_configs attribute in state
func TestMigratePagesProject_V4ToV5_WithEnvironmentVariables(t *testing.T) {
	t.Skip("Skipping test - v4 provider doesn't store environment_variables/secrets in deployment_configs attribute properly")

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-envvars-%s", rnd)

	// V4 config with environment_variables and secrets
	// Note: Including preview config to match API behavior (API creates both preview and production)
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  deployment_configs {
    preview {
      compatibility_date = "2024-01-01"
    }
    production {
      environment_variables = {
        NODE_ENV = "production"
        API_URL  = "https://api.example.com"
      }
      secrets = {
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
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
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
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-bindings-%s", rnd)
	bucketName := fmt.Sprintf("test-bucket-%s", rnd)

	// V4 config with bindings (kv_namespaces, d1_databases, r2_buckets, service_binding)
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  build_config {
    build_command = ""
  }

  deployment_configs {
    preview {
      compatibility_date = "2026-01-14"
    }
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

	migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
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
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
		}, migrationSteps...),
	})
}

// TestMigratePagesProject_V4ToV5_FullResource tests complete resource with all features
func TestMigratePagesProject_V4ToV5_FullResource(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-full-%s", rnd)

	v4Config := fmt.Sprintf(v4FullTF, rnd, accountID, projectName)

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
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
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

// TestMigratePagesProject_V4ToV5_WithBuildConfigComplete tests all build_config attributes
func TestMigratePagesProject_V4ToV5_WithBuildConfigComplete(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-build-complete-%s", rnd)

	v4Config := fmt.Sprintf(v4BuildConfigCompleteTF, rnd, accountID, projectName)

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
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("build_caching"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("build_command"), knownvalue.StringExact("npm run build")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("destination_dir"), knownvalue.StringExact("dist")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("root_dir"), knownvalue.StringExact("/frontend")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("web_analytics_tag"), knownvalue.StringExact("my-tag")),
			}),
		},
	})
}

// TestMigratePagesProject_V4ToV5_WithDurableObjectNamespaces tests durable_object_namespaces binding migration
func TestMigratePagesProject_V4ToV5_WithDurableObjectNamespaces(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	doNamespaceID := os.Getenv("TEST_CLOUDFLARE_DO_NAMESPACE_ID")

	if doNamespaceID == "" {
		t.Skip("Skipping - requires TEST_CLOUDFLARE_DO_NAMESPACE_ID")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-do-%s", rnd)

	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  build_config {
    build_caching   = false
    build_command   = ""
    destination_dir = ""
    root_dir        = ""
  }

  deployment_configs {
    preview {
      compatibility_date = "2024-01-01"
    }
    production {
      compatibility_date = "2024-01-01"
      durable_object_namespaces = {
        MY_DO = "%[4]s"
      }
    }
  }
}`, rnd, accountID, projectName, doNamespaceID)

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
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("durable_object_namespaces").AtMapKey("MY_DO").AtMapKey("namespace_id"), knownvalue.StringExact(doNamespaceID)),
			}),
		},
	})
}

// TestMigratePagesProject_V4ToV5_WithCompatibilityFlags tests compatibility_flags list migration
func TestMigratePagesProject_V4ToV5_WithCompatibilityFlags(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-compat-%s", rnd)

	v4Config := fmt.Sprintf(v4CompatibilityFlagsTF, rnd, accountID, projectName)

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
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("compatibility_flags"), knownvalue.ListExact([]knownvalue.Check{
					knownvalue.StringExact("nodejs_compat"),
					knownvalue.StringExact("streams_enable_constructors"),
				})),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("compatibility_flags"), knownvalue.ListExact([]knownvalue.Check{
					knownvalue.StringExact("nodejs_compat"),
					knownvalue.StringExact("streams_enable_constructors"),
					knownvalue.StringExact("transformstream_enable_standard_constructor"),
				})),
			}),
		},
	})
}

// TestMigratePagesProject_V4ToV5_WithDefaultChanges tests that v4 defaults are preserved in v5
func TestMigratePagesProject_V4ToV5_WithDefaultChanges(t *testing.T) {

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-defaults-%s", rnd)

	// V4 defaults: usage_model="bundled", fail_open=false
	// V5 defaults: usage_model="standard", fail_open=true
	// Migration should preserve v4 defaults
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  deployment_configs {
    preview {
      compatibility_date = "2024-01-01"
    }
    production {
      compatibility_date = "2024-01-01"
    }
  }
}`, rnd, accountID, projectName)

	migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
		// V4 defaults should be preserved in migration
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("usage_model"), knownvalue.StringExact("bundled")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("fail_open"), knownvalue.Bool(false)),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
		}, migrationSteps...),
	})
}

// TestMigratePagesProject_V4ToV5_WithUsageModelExplicit tests explicit usage_model migration
func TestMigratePagesProject_V4ToV5_WithUsageModelExplicit(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-usage-%s", rnd)

	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  deployment_configs {
    preview {
      usage_model = "unbound"
    }
    production {
      usage_model = "standard"
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
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("usage_model"), knownvalue.StringExact("unbound")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("usage_model"), knownvalue.StringExact("standard")),
			}),
		},
	})
}

// TestMigratePagesProject_V4ToV5_WithMultipleBindings tests multiple bindings of different types
func TestMigratePagesProject_V4ToV5_WithMultipleBindings(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	kvNamespaceID := os.Getenv("TEST_CLOUDFLARE_KV_NAMESPACE_ID")
	d1DatabaseID := os.Getenv("TEST_CLOUDFLARE_D1_DATABASE_ID")
	doNamespaceID := os.Getenv("TEST_CLOUDFLARE_DO_NAMESPACE_ID")

	if kvNamespaceID == "" || d1DatabaseID == "" || doNamespaceID == "" {
		t.Skip("Skipping - requires TEST_CLOUDFLARE_KV_NAMESPACE_ID, TEST_CLOUDFLARE_D1_DATABASE_ID, and TEST_CLOUDFLARE_DO_NAMESPACE_ID")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-multi-%s", rnd)

	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"

  build_config {
    build_caching   = false
    build_command   = ""
    destination_dir = ""
    root_dir        = ""
  }

  deployment_configs {
    preview {
      compatibility_date = "2024-01-01"
    }
    production {
      compatibility_date = "2024-01-01"
      kv_namespaces = {
        KV_BINDING_1 = "%[4]s"
        KV_BINDING_2 = "%[4]s"
      }
      d1_databases = {
        DB_BINDING_1 = "%[5]s"
        DB_BINDING_2 = "%[5]s"
      }
      r2_buckets = {
        BUCKET_1 = "my-bucket-1"
        BUCKET_2 = "my-bucket-2"
      }
      durable_object_namespaces = {
        DO_BINDING_1 = "%[6]s"
        DO_BINDING_2 = "%[6]s"
      }
    }
  }
}`, rnd, accountID, projectName, kvNamespaceID, d1DatabaseID, doNamespaceID)

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
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				// Verify multiple KV bindings migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces").AtMapKey("KV_BINDING_1").AtMapKey("namespace_id"), knownvalue.StringExact(kvNamespaceID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces").AtMapKey("KV_BINDING_2").AtMapKey("namespace_id"), knownvalue.StringExact(kvNamespaceID)),
				// Verify multiple D1 bindings migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("d1_databases").AtMapKey("DB_BINDING_1").AtMapKey("id"), knownvalue.StringExact(d1DatabaseID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("d1_databases").AtMapKey("DB_BINDING_2").AtMapKey("id"), knownvalue.StringExact(d1DatabaseID)),
				// Verify multiple R2 bindings migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("r2_buckets").AtMapKey("BUCKET_1").AtMapKey("name"), knownvalue.StringExact("my-bucket-1")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("r2_buckets").AtMapKey("BUCKET_2").AtMapKey("name"), knownvalue.StringExact("my-bucket-2")),
				// Verify multiple DO bindings migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("durable_object_namespaces").AtMapKey("DO_BINDING_1").AtMapKey("namespace_id"), knownvalue.StringExact(doNamespaceID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("durable_object_namespaces").AtMapKey("DO_BINDING_2").AtMapKey("namespace_id"), knownvalue.StringExact(doNamespaceID)),
			}),
		},
	})
}

// TestMigratePagesProject_V4ToV5_WithMultipleServiceBindings tests multiple service_binding → services migration
func TestMigratePagesProject_V4ToV5_WithMultipleServiceBindings(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-services-%s", rnd)

	v4Config := fmt.Sprintf(v4ServiceBindingsTF, rnd, accountID, projectName)

	migrationStep := acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("services").AtMapKey("MY_SERVICE_1").AtMapKey("service"), knownvalue.StringExact("worker-1")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("services").AtMapKey("MY_SERVICE_1").AtMapKey("environment"), knownvalue.StringExact("production")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("services").AtMapKey("MY_SERVICE_2").AtMapKey("service"), knownvalue.StringExact("worker-2")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("services").AtMapKey("MY_SERVICE_3").AtMapKey("service"), knownvalue.StringExact("worker-3")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("services").AtMapKey("MY_SERVICE_3").AtMapKey("environment"), knownvalue.StringExact("staging")),
	})

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
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			migrationStep,
		},
	})
}

// TestMigratePagesProject_V4ToV5_WithAlwaysUseLatestCompatibilityDate tests always_use_latest_compatibility_date migration
func TestMigratePagesProject_V4ToV5_WithAlwaysUseLatestCompatibilityDate(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_pages_project." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)
	projectName := fmt.Sprintf("tf-test-latest-compat-%s", rnd)

	v4Config := fmt.Sprintf(v4AlwaysUseLatestTF, rnd, accountID, projectName)

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
						VersionConstraint: version,
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: false,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("always_use_latest_compatibility_date"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("always_use_latest_compatibility_date"), knownvalue.Bool(false)),
			}),
		},
	})
}

// =============================================================================
// Unit Tests for State Upgrader Logic
// =============================================================================

func TestUpgradeStateFromV0_BasicConversions(t *testing.T) {
	// Test MergeEnvVarsAndSecretsPreview
	t.Run("merge_env_vars_and_secrets_preview", func(t *testing.T) {
		ctx := context.Background()

		envVars, _ := types.MapValue(types.StringType, map[string]attr.Value{
			"NODE_ENV": types.StringValue("production"),
			"API_URL":  types.StringValue("https://api.example.com"),
		})

		secrets, _ := types.MapValue(types.StringType, map[string]attr.Value{
			"API_KEY":     types.StringValue("secret123"),
			"DB_PASSWORD": types.StringValue("dbpass456"),
		})

		result := v500.MergeEnvVarsAndSecretsPreview(ctx, envVars, secrets)
		if result == nil {
			t.Fatal("expected non-nil result")
		}

		// Verify environment variables are converted to plain_text
		if nodeEnv, ok := (*result)["NODE_ENV"]; !ok {
			t.Error("expected NODE_ENV in result")
		} else {
			if nodeEnv.Type.ValueString() != "plain_text" {
				t.Errorf("expected NODE_ENV type to be plain_text, got %s", nodeEnv.Type.ValueString())
			}
			if nodeEnv.Value.ValueString() != "production" {
				t.Errorf("expected NODE_ENV value to be production, got %s", nodeEnv.Value.ValueString())
			}
		}

		// Verify secrets are converted to secret_text
		if apiKey, ok := (*result)["API_KEY"]; !ok {
			t.Error("expected API_KEY in result")
		} else {
			if apiKey.Type.ValueString() != "secret_text" {
				t.Errorf("expected API_KEY type to be secret_text, got %s", apiKey.Type.ValueString())
			}
		}
	})

	// Test ConvertKVNamespacesV0ToV5Preview
	t.Run("convert_kv_namespaces_preview", func(t *testing.T) {
		ctx := context.Background()

		kvNamespaces, _ := types.MapValue(types.StringType, map[string]attr.Value{
			"MY_KV":      types.StringValue("kv-namespace-id-123"),
			"ANOTHER_KV": types.StringValue("kv-namespace-id-456"),
		})

		result := v500.ConvertKVNamespacesV0ToV5Preview(ctx, kvNamespaces)
		if result == nil {
			t.Fatal("expected non-nil result")
		}

		if myKV, ok := (*result)["MY_KV"]; !ok {
			t.Error("expected MY_KV in result")
		} else {
			if myKV.NamespaceID.ValueString() != "kv-namespace-id-123" {
				t.Errorf("expected namespace_id to be kv-namespace-id-123, got %s", myKV.NamespaceID.ValueString())
			}
		}
	})

	// Test ConvertD1DatabasesV0ToV5Preview
	t.Run("convert_d1_databases_preview", func(t *testing.T) {
		ctx := context.Background()

		d1Databases, _ := types.MapValue(types.StringType, map[string]attr.Value{
			"MY_DB": types.StringValue("d1-database-id-123"),
		})

		result := v500.ConvertD1DatabasesV0ToV5Preview(ctx, d1Databases)
		if result == nil {
			t.Fatal("expected non-nil result")
		}

		if myDB, ok := (*result)["MY_DB"]; !ok {
			t.Error("expected MY_DB in result")
		} else {
			if myDB.ID.ValueString() != "d1-database-id-123" {
				t.Errorf("expected id to be d1-database-id-123, got %s", myDB.ID.ValueString())
			}
		}
	})

	// Test ConvertR2BucketsV0ToV5Preview
	t.Run("convert_r2_buckets_preview", func(t *testing.T) {
		ctx := context.Background()

		r2Buckets, _ := types.MapValue(types.StringType, map[string]attr.Value{
			"MY_BUCKET": types.StringValue("my-bucket-name"),
		})

		result := v500.ConvertR2BucketsV0ToV5Preview(ctx, r2Buckets)
		if result == nil {
			t.Fatal("expected non-nil result")
		}

		if myBucket, ok := (*result)["MY_BUCKET"]; !ok {
			t.Error("expected MY_BUCKET in result")
		} else {
			if myBucket.Name.ValueString() != "my-bucket-name" {
				t.Errorf("expected name to be my-bucket-name, got %s", myBucket.Name.ValueString())
			}
		}
	})

	// Test ConvertDurableObjectNamespacesV0ToV5Preview
	t.Run("convert_durable_object_namespaces_preview", func(t *testing.T) {
		ctx := context.Background()

		doNamespaces, _ := types.MapValue(types.StringType, map[string]attr.Value{
			"MY_DO": types.StringValue("do-namespace-id-123"),
		})

		result := v500.ConvertDurableObjectNamespacesV0ToV5Preview(ctx, doNamespaces)
		if result == nil {
			t.Fatal("expected non-nil result")
		}

		if myDO, ok := (*result)["MY_DO"]; !ok {
			t.Error("expected MY_DO in result")
		} else {
			if myDO.NamespaceID.ValueString() != "do-namespace-id-123" {
				t.Errorf("expected namespace_id to be do-namespace-id-123, got %s", myDO.NamespaceID.ValueString())
			}
		}
	})

	// Test null/unknown map handling
	t.Run("null_map_returns_nil", func(t *testing.T) {
		ctx := context.Background()
		nullMap := types.MapNull(types.StringType)

		if result := v500.ConvertKVNamespacesV0ToV5Preview(ctx, nullMap); result != nil {
			t.Error("expected nil for null map")
		}

		if result := v500.ConvertD1DatabasesV0ToV5Preview(ctx, nullMap); result != nil {
			t.Error("expected nil for null map")
		}

		if result := v500.ConvertR2BucketsV0ToV5Preview(ctx, nullMap); result != nil {
			t.Error("expected nil for null map")
		}

		if result := v500.ConvertDurableObjectNamespacesV0ToV5Preview(ctx, nullMap); result != nil {
			t.Error("expected nil for null map")
		}
	})

	// Test empty map handling
	t.Run("empty_map_returns_nil", func(t *testing.T) {
		ctx := context.Background()
		emptyMap, _ := types.MapValue(types.StringType, map[string]attr.Value{})

		if result := v500.ConvertKVNamespacesV0ToV5Preview(ctx, emptyMap); result != nil {
			t.Error("expected nil for empty map")
		}

		if result := v500.MergeEnvVarsAndSecretsPreview(ctx, emptyMap, emptyMap); result != nil {
			t.Error("expected nil when both env vars and secrets are empty")
		}
	})
}
