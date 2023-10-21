package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareTunnelMatchName(t *testing.T) {
	rnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareTunnelMatchName(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_tunnel.example", "status", "inactive"),
				),
			},
		},
	})
}

func testCloudflareTunnelMatchName(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
resource "cloudflare_tunnel" "example" {
	account_id = "%[1]s"
	name       = %[2]q
	secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}
data "cloudflare_tunnel" "example" {
	account_id = cloudflare_tunnel.example.account_id
	name       = cloudflare_tunnel.example.name
}
`, accountID, name)
}
