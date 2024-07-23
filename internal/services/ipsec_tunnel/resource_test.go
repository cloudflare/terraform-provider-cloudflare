package ipsec_tunnel_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareIPsecTunnelExists(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Not configured for Magic Transit")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_ipsec_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	psk := "asdf1234"

	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(name, "health_check_direction", "unidirectional"),
					resource.TestCheckResourceAttr(name, "health_check_rate", "mid"),
					resource.TestCheckResourceAttr(name, "psk", "asdf1234"),
					resource.TestCheckResourceAttr(name, "allowNullCipher", "false"),
					resource.TestCheckResourceAttr(name, "replay_protection", "true"),
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

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundIPsecTunnel, err := client.GetMagicTransitIPsecTunnel(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err != nil {
			return err
		}

		*tunnel = foundIPsecTunnel

		return nil
	}
}

func TestAccCloudflareIPsecTunnelUpdateDescription(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Not configured for Magic Transit")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_ipsec_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	psk := "asdf1234"

	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
	acctest.TestAccSkipForDefaultZone(t, "Not configured for Magic Transit")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_ipsec_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	psk := "asdf1234"
	pskUpdated := "1234asd"

	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
	return acctest.LoadTestCase("ipsectunnelsimple.tf", ID, description, accountID, psk)
}
