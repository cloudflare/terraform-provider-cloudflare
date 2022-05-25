package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareIPsecTunnelExists(t *testing.T) {
	skipMagicTransitTestForNonConfiguredDefaultZone(t)

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_ipsec_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "customer_endpoint", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "cloudflare_endpoint", "162.159.64.41"),
					resource.TestCheckResourceAttr(name, "interface_address", "10.212.0.9/31"),
				),
			},
		},
	})
}

func testAccCheckCloudflareIPsecTunnelExists(n string, tunnel *cloudflare.MagicTransitIPsecTunnel) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No IPsec tunnel is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundIPsecTunnel, err := client.GetMagicTransitIPsecTunnel(context.Background(), rs.Primary.Attributes["account_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		*tunnel = foundIPsecTunnel

		return nil
	}
}

func TestAccCloudflareIPsecTunnelUpdateDescription(t *testing.T) {
	skipMagicTransitTestForNonConfiguredDefaultZone(t)

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_ipsec_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
				),
			},
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd+"-updated", accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
				),
			},
		},
	})
}

func testAccCheckCloudflareIPsecTunnelSimple(ID, description, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ipsec_tunnel" "%[1]s" {
	account_id = "%[3]s"
	name = "%[2]s"
	customer_endpoint = "203.0.113.1"
	cloudflare_endpoint = "162.159.64.41"
	interface_address = "10.212.0.9/31"
	description = "%[2]s"
  }`, ID, description, accountID)
}
