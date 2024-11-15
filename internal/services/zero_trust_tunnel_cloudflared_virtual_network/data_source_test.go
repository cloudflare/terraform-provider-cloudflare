package zero_trust_tunnel_cloudflared_virtual_network_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareTunneVirtualNetworkDatasource_MatchName(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareTunnelVirtualNetworkMatchName(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_zero_trust_tunnel_cloudflared_virtual_network."+rnd, "comment", "test"),
				),
			},
		},
	})
}

func testCloudflareTunnelVirtualNetworkMatchName(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return acctest.LoadTestCase("cloudflaretunnelvirtualnetworkmatchname.tf", accountID, name)
}
