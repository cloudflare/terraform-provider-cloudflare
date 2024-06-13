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

func TestAccCloudflareGRETunnelExists(t *testing.T) {
	skipMagicTransitTestForNonConfiguredDefaultZone(t)

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_gre_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var Tunnel cloudflare.MagicTransitGRETunnel

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareGRETunnelSimple(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "customer_gre_endpoint", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "cloudflare_gre_endpoint", "162.159.64.41"),
					resource.TestCheckResourceAttr(name, "interface_address", "10.212.0.9/31"),
					resource.TestCheckResourceAttr(name, "ttl", "64"),
					resource.TestCheckResourceAttr(name, "mtu", "1476"),
					resource.TestCheckResourceAttr(name, "health_check_enabled", "true"),
					resource.TestCheckResourceAttr(name, "health_check_target", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "health_check_type", "request"),
				),
			},
		},
	})
}

func testAccCheckCloudflareGRETunnelExists(n string, tunnel *cloudflare.MagicTransitGRETunnel) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No GRE tunnel is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundGRETunnel, err := client.GetMagicTransitGRETunnel(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err != nil {
			return err
		}

		*tunnel = foundGRETunnel

		return nil
	}
}

func TestAccCloudflareGRETunnelUpdateDescription(t *testing.T) {
	skipMagicTransitTestForNonConfiguredDefaultZone(t)

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_gre_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var Tunnel cloudflare.MagicTransitGRETunnel

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareGRETunnelSimple(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
				),
			},
			{
				Config: testAccCheckCloudflareGRETunnelSimple(rnd, rnd, rnd+"-updated", accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareGRETunnelUpdateMulti(t *testing.T) {
	skipMagicTransitTestForNonConfiguredDefaultZone(t)

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_gre_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var Tunnel cloudflare.MagicTransitGRETunnel

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareGRETunnelSimple(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "customer_gre_endpoint", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "cloudflare_gre_endpoint", "162.159.64.41"),
					resource.TestCheckResourceAttr(name, "interface_address", "10.212.0.9/31"),
					resource.TestCheckResourceAttr(name, "ttl", "64"),
					resource.TestCheckResourceAttr(name, "mtu", "1476"),
					resource.TestCheckResourceAttr(name, "health_check_enabled", "true"),
					resource.TestCheckResourceAttr(name, "health_check_target", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "health_check_type", "request"),
				),
			},
			{
				Config: testAccCheckCloudflareGRETunnelMultiUpdate(rnd, rnd, rnd+"-updated", accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "customer_gre_endpoint", "203.0.113.2"),
					resource.TestCheckResourceAttr(name, "cloudflare_gre_endpoint", "162.159.64.41"),
					resource.TestCheckResourceAttr(name, "interface_address", "10.212.0.11/31"),
					resource.TestCheckResourceAttr(name, "ttl", "65"),
					resource.TestCheckResourceAttr(name, "mtu", "1475"),
					resource.TestCheckResourceAttr(name, "health_check_enabled", "true"),
					resource.TestCheckResourceAttr(name, "health_check_target", "203.0.113.2"),
					resource.TestCheckResourceAttr(name, "health_check_type", "reply"),
				),
			},
		},
	})
}

func testAccCheckCloudflareGRETunnelSimple(ID, name, description, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_gre_tunnel" "%[1]s" {
	account_id = "%[4]s"
	name = "%[2]s"
	customer_gre_endpoint = "203.0.113.1"
	cloudflare_gre_endpoint = "162.159.64.41"
	interface_address = "10.212.0.9/31"
	description = "%[3]s"
    ttl = 64
    mtu = 1476
    health_check_enabled = true
    health_check_target = "203.0.113.1"
    health_check_type = "request"
  }`, ID, name, description, accountID)
}

func testAccCheckCloudflareGRETunnelMultiUpdate(ID, name, description, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_gre_tunnel" "%[1]s" {
	account_id = "%[4]s"
	name = "%[2]s"
	customer_gre_endpoint = "203.0.113.2"
	cloudflare_gre_endpoint = "162.159.64.41"
	interface_address = "10.212.0.11/31"
	description = "%[3]s"
    ttl = 65
    mtu = 1475
    health_check_target = "203.0.113.2"
    health_check_type = "reply"
  }`, ID, name, description, accountID)
}
