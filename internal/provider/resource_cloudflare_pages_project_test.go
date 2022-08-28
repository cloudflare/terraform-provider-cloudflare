package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccPreCheckPages(t *testing.T) {
	testAccPreCheckEmail(t)
	testAccPreCheckApiKey(t)
	testAccPreCheckAccount(t)
	pagesOwner := os.Getenv("CLOUDFLARE_PAGES_OWNER")
	pagesRepo := os.Getenv("CLOUDFLARE_PAGES_REPO")
	if pagesOwner == "" || pagesRepo == "" {
		t.Fatal("CLOUDFLARE_PAGES_OWNER and CLOUDFLARE_PAGES_REPO must be set for this acceptance test")
	}
}

func testPagesProjectSource(resourceID, accountID, projectName, repoOwner, repoName string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_project" "%[1]s" {
		  account_id = "%[2]s"
		  name = "%[3]s"
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

func testPagesProjectBuildConfig(resourceID, accountID, projectName, repoOwner, repoName string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_project" "%[1]s" {
		  account_id = "%[2]s"
		  name = "%[3]s"
		  production_branch = "main"
		  build_config {
			build_command = "npm run build"
			destination_dir = "build"
			root_dir = "/"
			web_analytics_tag = "cee1c73f6e4743d0b5e6bb1a0bcaabcc"
			web_analytics_token = "021e1057c18547eca7b79f2516f06o7x"
		  }
		}
		`, resourceID, accountID, projectName, repoOwner, repoName)
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
				compatibility_date = "2022-08-15"
				compatibility_flags = ["preview_flag"]
			}
        	production {
				environment_variables = {
					ENVIRONMENT = "production"
					OTHER_VALUE = "other value"
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
				compatibility_date = "2022-08-16"
				compatibility_flags = ["production_flag", "second flag"]
      		}
		}
		}
		`, resourceID, accountID, projectName)
}

func testPagesProjectDirectUpload(resourceID, accountID, projectName string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_project" "%[1]s" {
		  account_id = "%[2]s"
		  name = "%[3]s"
		  production_branch = "main"
		}
		`, resourceID, accountID, projectName)
}

func testPagesProjectPreviewOnly(resourceID, accountID, projectName string) string {
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
			   compatibility_date = "2022-08-15"
			   compatibility_flags = ["preview_flag"]
		   }
		}
	}
		`, resourceID, accountID, projectName)
}

func testPagesProjectProductionOnly(resourceID, accountID, projectName string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_project" "%[1]s" {
		  account_id = "%[2]s"
		  name = "%[3]s"
		  production_branch = "main"
		  deployment_configs {
			production {
			   environment_variables = {
				   ENVIRONMENT = "preview"
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
			   compatibility_date = "2022-08-15"
			   compatibility_flags = ["preview_flag"]
		   }
		}
	}
		`, resourceID, accountID, projectName)
}

func TestAccTestPagesProject(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	pagesOwner := os.Getenv("CLOUDFLARE_PAGES_OWNER")
	pagesRepo := os.Getenv("CLOUDFLARE_PAGES_REPO")
	sourceConfigPrefix := "source.0.config.0"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckPages(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectSource(rnd, accountID, rnd, pagesOwner, pagesRepo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "source.0.type", "github"),
					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".owner", pagesOwner),
					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".repo_name", pagesRepo),
					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".production_branch", "main"),
					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".pr_comments_enabled", "true"),
					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".deployments_enabled", "true"),
					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".production_deployment_enabled", "true"),
					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".preview_deployment_setting", "custom"),

					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".preview_branch_includes.#", "2"),
					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".preview_branch_includes.0", "dev"),
					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".preview_branch_includes.1", "preview"),

					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".preview_branch_excludes.#", "2"),
					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".preview_branch_excludes.0", "main"),
					resource.TestCheckResourceAttr(name, sourceConfigPrefix+".preview_branch_excludes.1", "prod"),
				),
			},
		},
	})
}

func TestAccTestPagesProjectBuildConfig(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	pagesOwner := os.Getenv("CLOUDFLARE_PAGES_OWNER")
	pagesRepo := os.Getenv("CLOUDFLARE_PAGES_REPO")
	buildConfigPrefix := "build_config.0"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckPages(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectBuildConfig(rnd, accountID, rnd, pagesOwner, pagesRepo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, buildConfigPrefix+".build_command", "npm run build"),
					resource.TestCheckResourceAttr(name, buildConfigPrefix+".destination_dir", "build"),
					resource.TestCheckResourceAttr(name, buildConfigPrefix+".root_dir", "/"),
					resource.TestCheckResourceAttr(name, buildConfigPrefix+".web_analytics_tag", "cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
					resource.TestCheckResourceAttr(name, buildConfigPrefix+".web_analytics_token", "021e1057c18547eca7b79f2516f06o7x"),
				),
			},
		},
	})
}

func TestAccTestPagesProjectDeploymentConfig(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	previewPrefix := "deployment_configs.0.preview.0"
	productionPrefix := "deployment_configs.0.production.0"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckPages(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectDeploymentConfig(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "account_id", accountID),

					// Preview
					resource.TestCheckResourceAttr(name, previewPrefix+".compatibility_date", "2022-08-15"),
					resource.TestCheckResourceAttr(name, previewPrefix+".compatibility_flags.#", "1"),
					resource.TestCheckResourceAttr(name, previewPrefix+".compatibility_flags.0", "preview_flag"),

					resource.TestCheckResourceAttr(name, previewPrefix+".environment_variables.%", "1"),
					resource.TestCheckResourceAttr(name, previewPrefix+".environment_variables.ENVIRONMENT", "preview"),

					resource.TestCheckResourceAttr(name, previewPrefix+".kv_namespaces.%", "1"),
					resource.TestCheckResourceAttr(name, previewPrefix+".kv_namespaces.KV_BINDING", "5eb63bbbe01eeed093cb22bb8f5acdc3"),

					resource.TestCheckResourceAttr(name, previewPrefix+".durable_object_namespaces.%", "1"),
					resource.TestCheckResourceAttr(name, previewPrefix+".durable_object_namespaces.DO_BINDING", "5eb63bbbe01eeed093cb22bb8f5acdc3"),

					resource.TestCheckResourceAttr(name, previewPrefix+".d1_databases.%", "1"),
					resource.TestCheckResourceAttr(name, previewPrefix+".d1_databases.D1_BINDING", "445e2955-951a-4358-a35b-a4d0c813f63"),

					resource.TestCheckResourceAttr(name, previewPrefix+".r2_buckets.%", "1"),
					resource.TestCheckResourceAttr(name, previewPrefix+".r2_buckets.R2_BINDING", "some-bucket"),

					// Production
					resource.TestCheckResourceAttr(name, productionPrefix+".environment_variables.%", "2"),
					resource.TestCheckResourceAttr(name, productionPrefix+".environment_variables.ENVIRONMENT", "production"),
					resource.TestCheckResourceAttr(name, productionPrefix+".environment_variables.OTHER_VALUE", "other value"),

					resource.TestCheckResourceAttr(name, productionPrefix+".kv_namespaces.%", "2"),
					resource.TestCheckResourceAttr(name, productionPrefix+".kv_namespaces.KV_BINDING_1", "5eb63bbbe01eeed093cb22bb8f5acdc3"),
					resource.TestCheckResourceAttr(name, productionPrefix+".kv_namespaces.KV_BINDING_2", "3cdca5f8bb22bc390deee10ebbb36be5"),

					resource.TestCheckResourceAttr(name, productionPrefix+".durable_object_namespaces.%", "2"),
					resource.TestCheckResourceAttr(name, productionPrefix+".durable_object_namespaces.DO_BINDING_1", "5eb63bbbe01eeed093cb22bb8f5acdc3"),
					resource.TestCheckResourceAttr(name, productionPrefix+".durable_object_namespaces.DO_BINDING_2", "3cdca5f8bb22bc390deee10ebbb36be5"),

					resource.TestCheckResourceAttr(name, productionPrefix+".d1_databases.%", "2"),
					resource.TestCheckResourceAttr(name, productionPrefix+".d1_databases.D1_BINDING_1", "445e2955-951a-4358-a35b-a4d0c813f63"),
					resource.TestCheckResourceAttr(name, productionPrefix+".d1_databases.D1_BINDING_2", "a399414b-c697-409a-a688-377db6433cd9"),

					resource.TestCheckResourceAttr(name, productionPrefix+".r2_buckets.%", "2"),
					resource.TestCheckResourceAttr(name, productionPrefix+".r2_buckets.R2_BINDING_1", "some-bucket"),
					resource.TestCheckResourceAttr(name, productionPrefix+".r2_buckets.R2_BINDING_2", "other-bucket"),

					resource.TestCheckResourceAttr(name, productionPrefix+".compatibility_date", "2022-08-16"),
					resource.TestCheckResourceAttr(name, productionPrefix+".compatibility_flags.#", "2"),
					resource.TestCheckResourceAttr(name, productionPrefix+".compatibility_flags.0", "production_flag"),
					resource.TestCheckResourceAttr(name, productionPrefix+".compatibility_flags.1", "second flag"),
				),
			},
		},
	})
}

func TestAccTestPagesProjectDirectUpload(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckPages(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectDirectUpload(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "production_branch", "main"),
				),
			},
		},
	})
}

func TestAccTestPagesProjectPreviewOnly(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	previewPrefix := "deployment_configs.0.preview.0"
	

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckPages(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectPreviewOnly(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "production_branch", "main"),
					resource.TestCheckResourceAttr(name, previewPrefix+".compatibility_date", "2022-08-15"),
					resource.TestCheckResourceAttr(name, previewPrefix+".compatibility_flags.#", "1"),
					resource.TestCheckResourceAttr(name, previewPrefix+".compatibility_flags.0", "preview_flag"),

					resource.TestCheckResourceAttr(name, previewPrefix+".environment_variables.%", "1"),
					resource.TestCheckResourceAttr(name, previewPrefix+".environment_variables.ENVIRONMENT", "preview"),

					resource.TestCheckResourceAttr(name, previewPrefix+".kv_namespaces.%", "1"),
					resource.TestCheckResourceAttr(name, previewPrefix+".kv_namespaces.KV_BINDING", "5eb63bbbe01eeed093cb22bb8f5acdc3"),

					resource.TestCheckResourceAttr(name, previewPrefix+".durable_object_namespaces.%", "1"),
					resource.TestCheckResourceAttr(name, previewPrefix+".durable_object_namespaces.DO_BINDING", "5eb63bbbe01eeed093cb22bb8f5acdc3"),

					resource.TestCheckResourceAttr(name, previewPrefix+".d1_databases.%", "1"),
					resource.TestCheckResourceAttr(name, previewPrefix+".d1_databases.D1_BINDING", "445e2955-951a-4358-a35b-a4d0c813f63"),

					resource.TestCheckResourceAttr(name, previewPrefix+".r2_buckets.%", "1"),
					resource.TestCheckResourceAttr(name, previewPrefix+".r2_buckets.R2_BINDING", "some-bucket"),
				),
			},
		},
	})
}

func TestAccTestPagesProjectProductionOnly(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	productionPrefix := "deployment_configs.0.production.0"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckPages(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectProductionOnly(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "production_branch", "main"),
					resource.TestCheckResourceAttr(name, productionPrefix+".compatibility_date", "2022-08-15"),
					resource.TestCheckResourceAttr(name, productionPrefix+".compatibility_flags.#", "1"),
					resource.TestCheckResourceAttr(name, productionPrefix+".compatibility_flags.0", "preview_flag"),

					resource.TestCheckResourceAttr(name, productionPrefix+".environment_variables.%", "1"),
					resource.TestCheckResourceAttr(name, productionPrefix+".environment_variables.ENVIRONMENT", "preview"),

					resource.TestCheckResourceAttr(name, productionPrefix+".kv_namespaces.%", "1"),
					resource.TestCheckResourceAttr(name, productionPrefix+".kv_namespaces.KV_BINDING", "5eb63bbbe01eeed093cb22bb8f5acdc3"),

					resource.TestCheckResourceAttr(name, productionPrefix+".durable_object_namespaces.%", "1"),
					resource.TestCheckResourceAttr(name, productionPrefix+".durable_object_namespaces.DO_BINDING", "5eb63bbbe01eeed093cb22bb8f5acdc3"),

					resource.TestCheckResourceAttr(name, productionPrefix+".d1_databases.%", "1"),
					resource.TestCheckResourceAttr(name, productionPrefix+".d1_databases.D1_BINDING", "445e2955-951a-4358-a35b-a4d0c813f63"),

					resource.TestCheckResourceAttr(name, productionPrefix+".r2_buckets.%", "1"),
					resource.TestCheckResourceAttr(name, productionPrefix+".r2_buckets.R2_BINDING", "some-bucket"),
				),
			},
		},
	})
}
