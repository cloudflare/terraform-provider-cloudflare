package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func testPagesProjectConfig(resourceID, accountID, projectName string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_domain" "%[1]s" {
		  account_id = "%[2]s"
		  project_name = "%[3]s"

		`, resourceID, accountID, projectName)
}

func testPagesProjectBuildConfig(resourceID, accountID, projectName string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_domain" "%[1]s" {
		  account_id = "%[2]s"
		  project_name = "%[3]s"
		  build_config = {
			build_command = "npm run build",
			destination_dir = "build"
			root_dir = "/"
			web_analytics_tag = "cee1c73f6e4743d0b5e6bb1a0bcaabcc"
			web_analytics_token = "021e1057c18547eca7b79f2516f06o7x"
		}
		`, resourceID, accountID, projectName)
}

func testPagesProjectDeploymentConfig(resourceID, accountID, projectName string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_domain" "%[1]s" {
		  account_id = "%[2]s"
		  project_name = "%[3]s"
		  deployment_configs {
		 	preview {
				environment_variables = {
					ENVIROMENT = "preview"
          		}
				compatibility_date = "2022-08-15"
				compatibility_flags = ["preview_flag"]
			},
        	production {
				environment_variables = {
					ENVIRONMENT = "production"
        		}
				compatibility_date = "2022-08-15"
				compatibility_flags = ["preview_flag"]
      		}
		}
		`, resourceID, accountID, projectName)
}

func TestAccTestPagesProject(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	//resourceCloudflarePagesProject
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectConfig(rnd, accountID, "this-is-my-project-01"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "project_name", "this-is-my-project-01"),
					resource.TestCheckResourceAttr(name, "domain", "example.com"),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
				),
			},
		},
	})
}

func TestAccTestPagesProjectBuildConfig(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	//resourceCloudflarePagesProject
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectBuildConfig(rnd, accountID, "this-is-my-project-01"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "project_name", "this-is-my-project-01"),
					resource.TestCheckResourceAttr(name, "domain", "example.com"),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
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

func TestAccTestPagesProjectDeploymentConfig(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	//resourceCloudflarePagesProject
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectDeploymentConfig(rnd, accountID, "this-is-my-project-01"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "project_name", "this-is-my-project-01"),
					resource.TestCheckResourceAttr(name, "domain", "example.com"),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.preview.0.environment_variables.0.ENVIRONMENT", "preview"),
					resource.TestCheckResourceAttr(name, "deployment_configs.0.production.0.environment_variables.0.ENVIRONMENT", "production"),
				),
			},
		},
	})
}
