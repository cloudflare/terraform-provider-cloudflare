package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testPagesDomainConfig(resourceID, accountID, projectName, domain string) string {
	return fmt.Sprintf(`

		resource "cloudflare_pages_project" "%[1]s" {
			account_id = "%[2]s"
			name = "%[3]s"
			production_branch = "main"
		}
		resource "cloudflare_pages_domain" "%[1]s" {
		  account_id = "%[2]s"
		  project_name = cloudflare_pages_project.%[1]s.name
		  domain = "%[4]s"
		}
		`, resourceID, accountID, projectName, domain)
}

func TestAccTestPagesDomain(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_pages_domain." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckEmail(t)
			testAccPreCheckApiKey(t)
			testAccPreCheckAccount(t)
			testAccPreCheckDomain(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesDomainConfig(rnd, accountID, "terraform-project"+rnd, rnd+os.Getenv("CLOUDFLARE_DOMAIN")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "project_name", "terraform-project"+rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+os.Getenv("CLOUDFLARE_DOMAIN")),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
				),
			},
		},
	})
}
