package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareIPsecTunnelExists(t *testing.T) {
	skipMagicTransitTestForNonConfiguredDefaultZone(t)

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_ipsec_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	psk := "asdf1234"

	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID, psk),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "customer_endpoint", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "cloudflare_endpoint", "162.159.64.41"),
					resource.TestCheckResourceAttr(name, "interface_address", "10.212.0.9/31"),
					resource.TestCheckResourceAttr(name, "health_check_enabled", "true"),
					resource.TestCheckResourceAttr(name, "health_check_target", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "health_check_type", "request"),
					resource.TestCheckResourceAttr(name, "psk", "asdf1234"),
					resource.TestCheckResourceAttr(name, "allowNullCipher", "false"),
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
		foundIPsecTunnel, err := client.GetMagicTransitIPsecTunnel(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
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
	psk := "asdf1234"

	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID, psk),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
				),
			},
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd+"-updated", accountID, psk),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareIPsecTunnelUpdatePsk(t *testing.T) {
	skipMagicTransitTestForNonConfiguredDefaultZone(t)

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_ipsec_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	psk := "asdf1234"
	pskUpdated := "1234asd"

	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID, psk),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "psk", psk),
				),
			},
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID, pskUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "psk", pskUpdated),
				),
			},
		},
	})
}

func testAccCheckCloudflareIPsecTunnelSimple(ID, description, accountID, psk string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ipsec_tunnel" "%[1]s" {
	account_id = "%[3]s"
	name = "%[2]s"
	customer_endpoint = "203.0.113.1"
	cloudflare_endpoint = "162.159.64.41"
	interface_address = "10.212.0.9/31"
	description = "%[2]s"
	health_check_enabled = true
	health_check_target = "203.0.113.1"
	health_check_type = "request"
	psk = "%[4]s"
	allow_null_cipher = false
  }`, ID, description, accountID, psk)
}
