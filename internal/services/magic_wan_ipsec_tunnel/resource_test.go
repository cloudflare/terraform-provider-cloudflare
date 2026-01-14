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

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_magic_wan_ipsec_tunnel", &resource.Sweeper{
		Name: "cloudflare_magic_wan_ipsec_tunnel",
		F:    testSweepCloudflareMagicWanIPsecTunnel,
	})
}

func testSweepCloudflareMagicWanIPsecTunnel(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return fmt.Errorf("failed to create Cloudflare client: %w", clientErr)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping IPsec tunnels sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	tflog.Info(ctx, "Starting to list IPsec tunnels for sweeping")
	tunnels, err := client.ListMagicTransitIPsecTunnels(ctx, accountID)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch IPsec tunnels: %s", err))
		return fmt.Errorf("failed to fetch IPsec tunnels: %w", err)
	}

	if len(tunnels) == 0 {
		tflog.Info(ctx, "No IPsec tunnels to sweep")
		return nil
	}

	tflog.Info(ctx, fmt.Sprintf("Found %d IPsec tunnels to sweep", len(tunnels)))

	deletedCount := 0
	failedCount := 0

	for _, tunnel := range tunnels {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(tunnel.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting IPsec tunnel: %s (%s)", tunnel.Name, tunnel.ID))

		_, err := client.DeleteMagicTransitIPsecTunnel(ctx, accountID, tunnel.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete IPsec tunnel %s: %s", tunnel.ID, err))
			failedCount++
			continue
		}
		
		deletedCount++
		tflog.Info(ctx, fmt.Sprintf("Successfully deleted IPsec tunnel: %s", tunnel.ID))
	}

	tflog.Info(ctx, fmt.Sprintf("Completed sweeping IPsec tunnels: deleted %d, failed %d", deletedCount, failedCount))
	return nil
}

func TestAccCloudflareIPsecTunnelExists(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_ipsec_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	psk := "asdf1234"
	cfIP := utils.LookupMagicWanCfIP(t, accountID)
	// Use unique interface address per test to avoid conflicts
	interfaceAddr := "10.212.0.10/31"
	config := testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID, psk, cfIP, interfaceAddr)
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
					resource.TestCheckResourceAttr(name, "interface_address", "10.212.0.10/31"),
					resource.TestCheckResourceAttr(name, "health_check.enabled", "true"),
					resource.TestCheckResourceAttr(name, "health_check.target.effective", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "health_check.type", "request"),
					resource.TestCheckResourceAttr(name, "health_check.direction", "unidirectional"),
					resource.TestCheckResourceAttr(name, "health_check.rate", "low"),
					resource.TestCheckResourceAttr(name, "psk", "asdf1234"),
					resource.TestCheckResourceAttr(name, "replay_protection", "true"),
					resource.TestCheckResourceAttr(name, "automatic_return_routing", "true"),
					resource.TestCheckResourceAttr(name, "bgp.customer_asn", "65001"),
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
	// Use unique interface address per test to avoid conflicts
	interfaceAddr := "10.212.0.12/31"

	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID, psk, cfIP, interfaceAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
				),
			},
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd+"-updated", accountID, psk, cfIP, interfaceAddr),
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
	// Use unique interface address per test to avoid conflicts
	interfaceAddr := "10.212.0.14/31"

	var Tunnel cloudflare.MagicTransitIPsecTunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID, psk, cfIP, interfaceAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "psk", psk),
				),
			},
			{
				Config: testAccCheckCloudflareIPsecTunnelSimple(rnd, rnd, accountID, pskUpdated, cfIP, interfaceAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareIPsecTunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "psk", pskUpdated),
				),
			},
		},
	})
}

func testAccCheckCloudflareIPsecTunnelSimple(ID, description, accountID, psk, cfIP, interfaceAddr string) string {
	return acctest.LoadTestCase("ipsectunnelsimple.tf", ID, description, accountID, psk, cfIP, interfaceAddr)
}
