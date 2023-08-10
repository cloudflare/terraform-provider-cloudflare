package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testPagesProjectSource(resourceID, accountID, projectName, repoOwner, repoName string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_project" "%[1]s" {
		  account_id = "%[2]s"
		  name = "%[1]s"
		  production_branch = "main"
		  source {
			type = "github"
			config {
				owner = "%[4]s"
				repo_name = "%[5]s"
				production_branch = "main"
				pr_comments_enabled = true
				deployments_enabled = true
				production_deployment_enabled = true
				preview_deployment_setting = "custom"
				preview_branch_includes = ["dev","preview"]
				preview_branch_excludes = ["main", "prod"]
			}
		  }
		}
		`, resourceID, accountID, projectName, repoOwner, repoName)
}

func testPagesProjectBuildConfig(resourceID, accountID string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_project" "%[1]s" {
		  account_id = "%[2]s"
		  name = "%[1]s"
		  production_branch = "main"
		  build_config {
			build_command = "npm run build"
			destination_dir = "build"
			root_dir = "/"
			web_analytics_tag = "cee1c73f6e4743d0b5e6bb1a0bcaabcc"
			web_analytics_token = "021e1057c18547eca7b79f2516f06o7x"
		  }
		}
		`, resourceID, accountID, resourceID)
}

func testPagesProjectDeploymentConfig(resourceID, accountID, projectName string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_project" "%[1]s" {
		  account_id = "%[2]s"
		  name = "%[3]s"
		  production_branch = "main"
		  deployment_configs {
		 	preview {
				environment_variables = {
					ENVIRONMENT = "preview"
				}
				secrets = {
					TURNSTILE_SECRET = "1x0000000000000000000000000000000AA"
				}
				kv_namespaces = {
					KV_BINDING = "5eb63bbbe01eeed093cb22bb8f5acdc3"
				}
				durable_object_namespaces = {
					DO_BINDING = "5eb63bbbe01eeed093cb22bb8f5acdc3"
				}
				r2_buckets = {
					R2_BINDING = "some-bucket"
				}
				d1_databases = {
					D1_BINDING = "445e2955-951a-4358-a35b-a4d0c813f63"
				}
				service_binding {
					name = "MY_SERVICE_BINDING"
					service = "my-service"
					environment = "preview"
				}
				compatibility_date = "2022-08-15"
				compatibility_flags = ["preview_flag"]
				fail_open = true
				always_use_latest_compatibility_date = true
				usage_model = "unbound"
			}
        	production {
				environment_variables = {
					ENVIRONMENT = "production"
					OTHER_VALUE = "other value"
				}
				secrets = {
					TURNSTILE_SECRET = "1x0000000000000000000000000000000AA"
					TURNSTILE_INVIS_SECRET = "2x0000000000000000000000000000000AA"
				}
				kv_namespaces = {
					KV_BINDING_1 = "5eb63bbbe01eeed093cb22bb8f5acdc3"
					KV_BINDING_2 = "3cdca5f8bb22bc390deee10ebbb36be5"
				}
				durable_object_namespaces = {
					DO_BINDING_1 = "5eb63bbbe01eeed093cb22bb8f5acdc3"
					DO_BINDING_2 = "3cdca5f8bb22bc390deee10ebbb36be5"
				}
				r2_buckets = {
					R2_BINDING_1 = "some-bucket"
					R2_BINDING_2 = "other-bucket"
				}
				d1_databases = {
					D1_BINDING_1 = "445e2955-951a-4358-a35b-a4d0c813f63"
					D1_BINDING_2 = "a399414b-c697-409a-a688-377db6433cd9"
				}
				service_binding {
					name = "MY_SERVICE_BINDING"
					service = "my-service"
					environment = "production"
				}
				compatibility_date = "2022-08-16"
				compatibility_flags = ["production_flag", "second flag"]
				fail_open = true
				always_use_latest_compatibility_date = false
				usage_model = "bundled"
				placement {
					mode = "smart"
				}
      		}
		}
		}
		`, resourceID, accountID, projectName)
}

func testPagesProjectDirectUpload(resourceID, accountID string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_project" "%[1]s" {
		  account_id = "%[2]s"
		  name = "%[1]s"
		  production_branch = "main"
		}

		`, resourceID, accountID)
}

func TestAccCloudflarePagesProject_Basic(t *testing.T) {
	skipForDefaultAccount(t, "Pending investigation into automating the setup and teardown.")

	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	pagesOwner := os.Getenv("CLOUDFLARE_PAGES_OWNER")
	pagesRepo := os.Getenv("CLOUDFLARE_PAGES_REPO")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckPages(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectSource(rnd, accountID, rnd, pagesOwner, pagesRepo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "source.0.type", "github"),
					resource.TestCheckResourceAttr(name, "source.0.config.0.owner", pagesOwner),
					resource.TestCheckResourceAttr(name, "source.0.config.0.repo_name", pagesRepo),
					resource.TestCheckResourceAttr(name, "source.0.config.0.production_branch", "main"),
					resource.TestCheckResourceAttr(name, "source.0.config.0.pr_comments_enabled", "true"),
					resource.TestCheckResourceAttr(name, "source.0.config.0.deployments_enabled", "true"),
					resource.TestCheckResourceAttr(name, "source.0.config.0.production_deployment_enabled", "true"),
					resource.TestCheckResourceAttr(name, "source.0.config.0.preview_deployment_setting", "custom"),

					resource.TestCheckResourceAttr(name, "source.0.config.0.preview_branch_includes.#", "2"),
					resource.TestCheckResourceAttr(name, "source.0.config.0.preview_branch_includes.0", "dev"),
					resource.TestCheckResourceAttr(name, "source.0.config.0.preview_branch_includes.1", "preview"),

					resource.TestCheckResourceAttr(name, "source.0.config.0.preview_branch_excludes.#", "2"),
					resource.TestCheckResourceAttr(name, "source.0.config.0.preview_branch_excludes.0", "main"),
					resource.TestCheckResourceAttr(name, "source.0.config.0.preview_branch_excludes.1", "prod"),
				),
			},
		},
	})
}

func TestAccCloudflarePagesProject_BuildConfig(t *testing.T) {
	skipForDefaultAccount(t, "Pending investigation into automating the setup and teardown.")

	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckPages(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectBuildConfig(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "build_config.0.build_command", "npm run build"),
					resource.TestCheckResourceAttr(name, "build_config.0.destination_dir", "build"),
					resource.TestCheckResourceAttr(name, "build_config.0.root_dir", "/"),
					resource.TestCheckResourceAttr(name, "build_config.0.web_analytics_tag", "cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
					resource.TestCheckResourceAttr(name, "build_config.0.web_analytics_token", "021e1057c18547eca7b79f2516f06o7x"),
				),
			},
		},
	})
}

func TestAccCloudflarePagesProject_DeploymentConfig(t *testing.T) {
	skipForDefaultAccount(t, "Pending investigation into automating the setup and teardown.")

	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckPages(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectDeploymentConfig(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),

					// Preview
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.compatibility_date", "2022-08-15"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.compatibility_flags.#", "1"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.compatibility_flags.0", "preview_flag"),

					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.environment_variables.ENVIRONMENT", "preview"),

					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.secrets.%", "1"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.secrets.TURNSTILE_SECRET", "1x0000000000000000000000000000000AA"),

					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.kv_namespaces.%", "1"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.kv_namespaces.KV_BINDING", "5eb63bbbe01eeed093cb22bb8f5acdc3"),

					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.durable_object_namespaces.%", "1"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.durable_object_namespaces.DO_BINDING", "5eb63bbbe01eeed093cb22bb8f5acdc3"),

					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.d1_databases.%", "1"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.d1_databases.D1_BINDING", "445e2955-951a-4358-a35b-a4d0c813f63"),

					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.r2_buckets.%", "1"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.r2_buckets.R2_BINDING", "some-bucket"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.fail_open", "true"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.always_use_latest_compatibility_date", "true"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.usage_model", "unbound"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.placement.%", "0"),

					// Production
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.environment_variables.%", "2"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.environment_variables.ENVIRONMENT", "production"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.environment_variables.OTHER_VALUE", "other value"),

					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.secrets.%", "2"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.secrets.TURNSTILE_SECRET", "1x0000000000000000000000000000000AA"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.secrets.TURNSTILE_INVIS_SECRET", "2x0000000000000000000000000000000AA"),

					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.kv_namespaces.%", "2"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.kv_namespaces.KV_BINDING_1", "5eb63bbbe01eeed093cb22bb8f5acdc3"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.kv_namespaces.KV_BINDING_2", "3cdca5f8bb22bc390deee10ebbb36be5"),

					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.durable_object_namespaces.%", "2"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.durable_object_namespaces.DO_BINDING_1", "5eb63bbbe01eeed093cb22bb8f5acdc3"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.durable_object_namespaces.DO_BINDING_2", "3cdca5f8bb22bc390deee10ebbb36be5"),

					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.d1_databases.%", "2"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.d1_databases.D1_BINDING_1", "445e2955-951a-4358-a35b-a4d0c813f63"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.d1_databases.D1_BINDING_2", "a399414b-c697-409a-a688-377db6433cd9"),

					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.r2_buckets.%", "2"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.r2_buckets.R2_BINDING_1", "some-bucket"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.r2_buckets.R2_BINDING_2", "other-bucket"),

					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.compatibility_date", "2022-08-16"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.compatibility_flags.#", "2"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.compatibility_flags.0", "production_flag"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.compatibility_flags.1", "second flag"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.fail_open", "true"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.always_use_latest_compatibility_date", "false"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.usage_model", "bundled"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.placement.#", "1"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.placement.0.mode", "smart"),
				),
			},
		},
	})
}

func TestAccCloudflarePagesProject_DirectUpload(t *testing.T) {
	skipForDefaultAccount(t, "Pending investigation into automating the setup and teardown.")

	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID = os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectDirectUpload(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "production_branch", "main"),
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}
