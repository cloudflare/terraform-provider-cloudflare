package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func testEmailRoutingAddressConfig(resourceID, accountID, email string) string {
	return fmt.Sprintf(`
		resource "cloudflare_email_routing_address" "%[1]s" {
		  account_id = "%[2]s"
		  email = "%[3]s"
		}
		`, resourceID, accountID, email)
}

func TestAccTestEmailRoutingAddress(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_email_routing_address." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	//resourceCloudflareEmailRoutingAddress
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingAddressConfig(rnd, accountID, "user@example.com"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "email", "user@example.com"),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
				),
			},
		},
	})
}
