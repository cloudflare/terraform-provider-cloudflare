package magic_wan_ipsec_tunnel_test

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
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_ipsec_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	psk := "asdf1234"
	cfIP := utils.LookupMagicWanCfIP(t, accountID)
	config := testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID, psk, cfIP)
	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "customer_endpoint", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "cloudflare_endpoint", cfIP),
					resource.TestCheckResourceAttr(name, "interface_address", "10.212.0.9/31"),
					resource.TestCheckResourceAttr(name, "health_check.enabled", "true"),
					resource.TestCheckResourceAttr(name, "health_check.target.effective", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "health_check.type", "request"),
					resource.TestCheckResourceAttr(name, "health_check.direction", "unidirectional"),
					resource.TestCheckResourceAttr(name, "health_check.rate", "low"),
					resource.TestCheckResourceAttr(name, "psk", "asdf1234"),
					resource.TestCheckResourceAttr(name, "replay_protection", "true"),
				),
			},
			{
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // expect no change
			},
			{
				ResourceName: name,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					rs, ok := state.RootModule().Resources[name]
					if !ok {
						return "", fmt.Errorf("not found: %s", name)
					}
					return fmt.Sprintf("%s/%s", accountID, rs.Primary.ID), nil
				},
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psk"},
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
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_ipsec_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	psk := "asdf1234"
	cfIP := utils.LookupMagicWanCfIP(t, accountID)

	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID, psk, cfIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
				),
			},
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd+"-updated", accountID, psk, cfIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareIPsecTunnelUpdatePsk(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_ipsec_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	psk := "asdf1234"
	pskUpdated := "1234asd"
	cfIP := utils.LookupMagicWanCfIP(t, accountID)

	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID, psk, cfIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "psk", psk),
				),
			},
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID, pskUpdated, cfIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "psk", pskUpdated),
				),
			},
		},
	})
}

func testAccCheckCloudflareIPsecTunnelSimple(ID, description, accountID, psk, cfIP string) string {
	return acctest.LoadTestCase("ipsectunnelsimple.tf", ID, description, accountID, psk, cfIP)
}
