package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareTunneVirtualNetworkMatchName(t *testing.T) {
	rnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareTunnelVirtualNetworkMatchName(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_tunnel_virtual_network.example", "comment", "test"),
				),
			},
		},
	})
}

func testCloudflareTunnelVirtualNetworkMatchName(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
resource "cloudflare_tunnel_virtual_network" "example" {
	account_id = "%[1]s"
	name       = %[2]q
	comment     = "test"
}
data "cloudflare_tunnel_virtual_network" "example" {
	account_id = cloudflare_tunnel_virtual_network.example.account_id
	name       = cloudflare_tunnel_virtual_network.example.name
}
`, accountID, name)
}
