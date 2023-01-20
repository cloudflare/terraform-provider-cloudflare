package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingAddressConfig(rnd, accountID, "user@example.com"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "email", "user@example.com"),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
				),
			},
		},
	})
}
