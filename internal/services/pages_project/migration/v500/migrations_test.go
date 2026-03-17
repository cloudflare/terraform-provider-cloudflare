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

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_project/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

//go:embed testdata/v4_minimal.tf
var v4MinimalTF string

//go:embed testdata/v5_minimal.tf
var v5MinimalTF string

//go:embed testdata/v4_build_config.tf
var v4BuildConfigTF string

//go:embed testdata/v5_build_config.tf
var v5BuildConfigTF string

//go:embed testdata/v4_build_config_complete.tf
var v4BuildConfigCompleteTF string

//go:embed testdata/v5_build_config_complete.tf
var v5BuildConfigCompleteTF string

//go:embed testdata/v4_deployment_configs.tf
var v4DeploymentConfigsTF string

//go:embed testdata/v5_deployment_configs.tf
var v5DeploymentConfigsTF string

//go:embed testdata/v4_compatibility_flags.tf
var v4CompatibilityFlagsTF string

//go:embed testdata/v5_compatibility_flags.tf
var v5CompatibilityFlagsTF string

//go:embed testdata/v4_full.tf
var v4FullTF string

//go:embed testdata/v5_full.tf
var v5FullTF string

//go:embed testdata/v4_always_use_latest.tf
var v4AlwaysUseLatestTF string

//go:embed testdata/v5_always_use_latest.tf
var v5AlwaysUseLatestTF string

//go:embed testdata/v4_service_bindings.tf
var v4ServiceBindingsTF string

//go:embed testdata/v5_service_bindings.tf
var v5ServiceBindingsTF string

//go:embed testdata/v4_durable_object_namespaces.tf
var v4DurableObjectNamespacesTF string

//go:embed testdata/v5_durable_object_namespaces.tf
var v5DurableObjectNamespacesTF string

//go:embed testdata/v4_default_changes.tf
var v4DefaultChangesTF string

//go:embed testdata/v5_default_changes.tf
var v5DefaultChangesTF string

//go:embed testdata/v4_usage_model_explicit.tf
var v4UsageModelExplicitTF string

//go:embed testdata/v5_usage_model_explicit.tf
var v5UsageModelExplicitTF string

//go:embed testdata/v4_source_and_config.tf
var v4SourceAndConfigTF string

//go:embed testdata/v4_environment_variables.tf
var v4EnvironmentVariablesTF string

//go:embed testdata/v4_bindings.tf
var v4BindingsTF string

//go:embed testdata/v4_multiple_bindings.tf
var v4MultipleBindingsTF string

// TestMigratePagesProject_V4ToV5_Minimal tests basic migration with minimal config
func TestMigratePagesProject_V4ToV5_Minimal(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, projectName string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v4MinimalTF, rnd, accountID, projectName)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v5MinimalTF, rnd, accountID, projectName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_pages_project." + rnd
			tmpDir := t.TempDir()
			projectName := fmt.Sprintf("tf-test-minimal-%s", rnd)
			testConfig := tc.configFn(rnd, accountID, projectName)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// For v5 tests, use local provider; for v4 tests, use external provider
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
					Config:             testConfig,
					ExpectNonEmptyPlan: false,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
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
		})
	}
}

// TestMigratePagesProject_V4ToV5_WithBuildConfig tests build_config block to attribute conversion
func TestMigratePagesProject_V4ToV5_WithBuildConfig(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, projectName string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v4BuildConfigTF, rnd, accountID, projectName)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v5BuildConfigTF, rnd, accountID, projectName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_pages_project." + rnd
			tmpDir := t.TempDir()
			projectName := fmt.Sprintf("tf-test-build-%s", rnd)
			testConfig := tc.configFn(rnd, accountID, projectName)
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
					Config:             testConfig,
					ExpectNonEmptyPlan: false,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
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
		})
	}
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
	v4Config := fmt.Sprintf(v4SourceAndConfigTF, rnd, accountID, projectName, repoName)

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
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, projectName string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v4DeploymentConfigsTF, rnd, accountID, projectName)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v5DeploymentConfigsTF, rnd, accountID, projectName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_pages_project." + rnd
			tmpDir := t.TempDir()
			projectName := fmt.Sprintf("tf-test-deploy-%s", rnd)
			testConfig := tc.configFn(rnd, accountID, projectName)
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
					Config:             testConfig,
					ExpectNonEmptyPlan: false,
				}
			}

			// Build state checks based on version being tested
			var stateChecks []statecheck.StateCheck

			// Common checks for both v4 and v5
			stateChecks = append(stateChecks,
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("compatibility_date"), knownvalue.StringExact("2024-01-01")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("placement"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("placement").AtMapKey("mode"), knownvalue.StringExact("smart")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("compatibility_date"), knownvalue.StringExact("2024-01-01")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("placement"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("placement").AtMapKey("mode"), knownvalue.StringExact("smart")),
			)

			// Version-specific checks for usage_model and fail_open
			if tc.version == currentProviderVersion {
				// v5 uses v5 defaults
				stateChecks = append(stateChecks,
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("usage_model"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("fail_open"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("usage_model"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("fail_open"), knownvalue.Bool(true)),
				)
			} else {
				// v4 migration preserves v4 defaults
				stateChecks = append(stateChecks,
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("usage_model"), knownvalue.StringExact("bundled")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("fail_open"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("usage_model"), knownvalue.StringExact("bundled")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("fail_open"), knownvalue.Bool(false)),
				)
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, stateChecks),
				},
			})
		})
	}
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
	v4Config := fmt.Sprintf(v4EnvironmentVariablesTF, rnd, accountID, projectName)

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
	v4Config := fmt.Sprintf(v4BindingsTF, rnd, accountID, projectName, kvNamespaceID, d1DatabaseID, bucketName, workerName)

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
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, projectName string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v4FullTF, rnd, accountID, projectName)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v5FullTF, rnd, accountID, projectName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_pages_project." + rnd
			tmpDir := t.TempDir()
			projectName := fmt.Sprintf("tf-test-full-%s", rnd)
			testConfig := tc.configFn(rnd, accountID, projectName)
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
					Config:             testConfig,
					ExpectNonEmptyPlan: false,
				}
			}

			// Build state checks based on version being tested
			var stateChecks []statecheck.StateCheck

			// Common checks for both v4 and v5
			stateChecks = append(stateChecks,
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
				// Verify build_config
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("build_command"), knownvalue.StringExact("npm run build")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("destination_dir"), knownvalue.StringExact("dist")),
				// Verify deployment_configs structure
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("compatibility_date"), knownvalue.StringExact("2024-01-01")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("placement"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("placement").AtMapKey("mode"), knownvalue.StringExact("smart")),
			)

			// Version-specific checks for usage_model and fail_open
			if tc.version == currentProviderVersion {
				// v5 uses v5 defaults
				stateChecks = append(stateChecks,
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("usage_model"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("fail_open"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("usage_model"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("fail_open"), knownvalue.Bool(true)),
				)
			} else {
				// v4 migration preserves v4 defaults
				stateChecks = append(stateChecks,
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("usage_model"), knownvalue.StringExact("bundled")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("fail_open"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("usage_model"), knownvalue.StringExact("bundled")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("fail_open"), knownvalue.Bool(false)),
				)
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, stateChecks),
				},
			})
		})
	}
}

// TestMigratePagesProject_V4ToV5_WithBuildConfigComplete tests all build_config attributes
func TestMigratePagesProject_V4ToV5_WithBuildConfigComplete(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, projectName string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v4BuildConfigCompleteTF, rnd, accountID, projectName)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v5BuildConfigCompleteTF, rnd, accountID, projectName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_pages_project." + rnd
			tmpDir := t.TempDir()
			projectName := fmt.Sprintf("tf-test-build-complete-%s", rnd)
			testConfig := tc.configFn(rnd, accountID, projectName)
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
					Config:             testConfig,
					ExpectNonEmptyPlan: false,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("build_caching"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("build_command"), knownvalue.StringExact("npm run build")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("destination_dir"), knownvalue.StringExact("dist")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("root_dir"), knownvalue.StringExact("/frontend")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("build_config").AtMapKey("web_analytics_tag"), knownvalue.StringExact("my-tag")),
			}),
				},
			})
		})
	}
}

// TestMigratePagesProject_V4ToV5_WithDurableObjectNamespaces tests durable_object_namespaces binding migration
func TestMigratePagesProject_V4ToV5_WithDurableObjectNamespaces(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	doNamespaceID := os.Getenv("TEST_CLOUDFLARE_DO_NAMESPACE_ID")

	if doNamespaceID == "" {
		t.Skip("Skipping - requires TEST_CLOUDFLARE_DO_NAMESPACE_ID")
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, projectName, doNamespaceID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, projectName, doNamespaceID string) string {
				return fmt.Sprintf(v4DurableObjectNamespacesTF, rnd, accountID, projectName, doNamespaceID)
			},
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, projectName, doNamespaceID string) string {
				return fmt.Sprintf(v5DurableObjectNamespacesTF, rnd, accountID, projectName, doNamespaceID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_pages_project." + rnd
			tmpDir := t.TempDir()
			projectName := fmt.Sprintf("tf-test-do-%s", rnd)
			testConfig := tc.configFn(rnd, accountID, projectName, doNamespaceID)
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
					Config:             testConfig,
					ExpectNonEmptyPlan: false,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("durable_object_namespaces").AtMapKey("MY_DO").AtMapKey("namespace_id"), knownvalue.StringExact(doNamespaceID)),
					}),
				},
			})
		})
	}
}

// TestMigratePagesProject_V4ToV5_WithCompatibilityFlags tests compatibility_flags list migration
func TestMigratePagesProject_V4ToV5_WithCompatibilityFlags(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, projectName string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v4CompatibilityFlagsTF, rnd, accountID, projectName)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v5CompatibilityFlagsTF, rnd, accountID, projectName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_pages_project." + rnd
			tmpDir := t.TempDir()
			projectName := fmt.Sprintf("tf-test-compat-%s", rnd)
			testConfig := tc.configFn(rnd, accountID, projectName)
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
					Config:             testConfig,
					ExpectNonEmptyPlan: false,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
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
		})
	}
}

// TestMigratePagesProject_V4ToV5_WithDefaultChanges tests that v4 defaults are preserved in v5
func TestMigratePagesProject_V4ToV5_WithDefaultChanges(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, projectName string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v4DefaultChangesTF, rnd, accountID, projectName)
			},
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v5DefaultChangesTF, rnd, accountID, projectName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_pages_project." + rnd
			tmpDir := t.TempDir()
			projectName := fmt.Sprintf("tf-test-defaults-%s", rnd)
			testConfig := tc.configFn(rnd, accountID, projectName)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// Build state checks based on version
			var stateChecks []statecheck.StateCheck
			if tc.version == currentProviderVersion {
				// V5 uses v5 defaults
				stateChecks = []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("usage_model"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("fail_open"), knownvalue.Bool(true)),
				}
			} else {
				// V4 migration preserves v4 defaults
				stateChecks = []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("usage_model"), knownvalue.StringExact("bundled")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("fail_open"), knownvalue.Bool(false)),
				}
			}

			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, stateChecks)

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
					Config:             testConfig,
					ExpectNonEmptyPlan: false,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps:      append([]resource.TestStep{firstStep}, migrationSteps...),
			})
		})
	}
}

// TestMigratePagesProject_V4ToV5_WithUsageModelExplicit tests explicit usage_model migration
func TestMigratePagesProject_V4ToV5_WithUsageModelExplicit(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, projectName string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v4UsageModelExplicitTF, rnd, accountID, projectName)
			},
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v5UsageModelExplicitTF, rnd, accountID, projectName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_pages_project." + rnd
			tmpDir := t.TempDir()
			projectName := fmt.Sprintf("tf-test-usage-%s", rnd)
			testConfig := tc.configFn(rnd, accountID, projectName)
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
					Config:             testConfig,
					ExpectNonEmptyPlan: false,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("usage_model"), knownvalue.StringExact("unbound")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("usage_model"), knownvalue.StringExact("standard")),
					}),
				},
			})
		})
	}
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

	v4Config := fmt.Sprintf(v4MultipleBindingsTF, rnd, accountID, projectName, kvNamespaceID, kvNamespaceID, d1DatabaseID, d1DatabaseID, doNamespaceID, doNamespaceID)

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
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, projectName string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v4AlwaysUseLatestTF, rnd, accountID, projectName)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, projectName string) string {
				return fmt.Sprintf(v5AlwaysUseLatestTF, rnd, accountID, projectName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_pages_project." + rnd
			tmpDir := t.TempDir()
			projectName := fmt.Sprintf("tf-test-latest-compat-%s", rnd)
			testConfig := tc.configFn(rnd, accountID, projectName)
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
					Config:             testConfig,
					ExpectNonEmptyPlan: false,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("always_use_latest_compatibility_date"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("always_use_latest_compatibility_date"), knownvalue.Bool(false)),
			}),
				},
			})
		})
	}
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
