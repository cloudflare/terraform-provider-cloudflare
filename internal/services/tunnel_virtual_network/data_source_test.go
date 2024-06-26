package tunnel_virtual_network_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareTunneVirtualNetwork_MatchName(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareTunnelVirtualNetworkMatchName(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_tunnel_virtual_network."+rnd, "comment", "test"),
				),
			},
		},
	})
}

func testCloudflareTunnelVirtualNetworkMatchName(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
resource "cloudflare_tunnel_virtual_network" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	comment     = "test"
}
data "cloudflare_tunnel_virtual_network" "%[2]s" {
	account_id = cloudflare_tunnel_virtual_network.%[2]s.account_id
	name       = cloudflare_tunnel_virtual_network.%[2]s.name
	depends_on = ["cloudflare_tunnel_virtual_network.%[2]s"]
}
`, accountID, name)
}
