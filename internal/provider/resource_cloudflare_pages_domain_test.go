package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func testPagesDomainConfig(resourceID, accountID, projectName, domain string) string {
	return fmt.Sprintf(`
		resource "cloudflare_pages_domain" "%[1]s" {
		  account_id = "%[2]s"
		  project_name = "%[3]s"
		  domain = "%[4]s"
		}
		`, resourceID, accountID, projectName, domain)
}

func TestPagesDomain(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_pages_domain." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	//resourceCloudflarePagesDomain
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testPagesDomainConfig(rnd, accountID, "this-is-my-project-01", "example.com"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "project_name", "this-is-my-project-01"),
					resource.TestCheckResourceAttr(name, "domain", "example.com"),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
				),
			},
		},
	})
}
