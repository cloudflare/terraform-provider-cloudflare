package magic_wan_gre_tunnel_test

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
	resource.AddTestSweepers("cloudflare_magic_wan_gre_tunnel", &resource.Sweeper{
		Name: "cloudflare_magic_wan_gre_tunnel",
		F:    testSweepCloudflareMagicWanGRETunnel,
	})
}

func testSweepCloudflareMagicWanGRETunnel(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return fmt.Errorf("failed to create Cloudflare client: %w", clientErr)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping GRE tunnels sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	tflog.Info(ctx, "Starting to list GRE tunnels for sweeping")
	tunnels, err := client.ListMagicTransitGRETunnels(ctx, accountID)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch GRE tunnels: %s", err))
		return fmt.Errorf("failed to fetch GRE tunnels: %w", err)
	}

	if len(tunnels) == 0 {
		tflog.Info(ctx, "No GRE tunnels to sweep")
		return nil
	}

	tflog.Info(ctx, fmt.Sprintf("Found %d GRE tunnels to sweep", len(tunnels)))

	deletedCount := 0
	failedCount := 0

	for _, tunnel := range tunnels {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(tunnel.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting GRE tunnel: %s (%s)", tunnel.Name, tunnel.ID))

		_, err := client.DeleteMagicTransitGRETunnel(ctx, accountID, tunnel.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete GRE tunnel %s: %s", tunnel.ID, err))
			failedCount++
			continue
		}
		
		deletedCount++
		tflog.Info(ctx, fmt.Sprintf("Successfully deleted GRE tunnel: %s", tunnel.ID))
	}

	tflog.Info(ctx, fmt.Sprintf("Completed sweeping GRE tunnels: deleted %d, failed %d", deletedCount, failedCount))
	return nil
}

func TestAccCloudflareGRETunnelExists(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_gre_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	cfIP := utils.LookupMagicWanCfIP(t, accountID)
	// Use unique endpoints and interface address per test to avoid conflicts
	customerEndpoint := "203.0.113.10"
	interfaceAddr := "10.213.0.20/31"
	config := testAccCheckCloudflareGRETunnelSimple(rnd, rnd, rnd, accountID, cfIP, customerEndpoint, interfaceAddr)

	var Tunnel cloudflare.MagicTransitGRETunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "customer_gre_endpoint", "203.0.113.10"),
					resource.TestCheckResourceAttr(name, "cloudflare_gre_endpoint", cfIP),
					resource.TestCheckResourceAttr(name, "interface_address", "10.213.0.20/31"),
					resource.TestCheckResourceAttr(name, "health_check.target.effective", "203.0.113.10"),
					resource.TestCheckResourceAttr(name, "automatic_return_routing", "true"),
					resource.TestCheckResourceAttr(name, "bgp.customer_asn", "65002"),
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
				ImportState:       true,
				ImportStateVerify: true,
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

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundGRETunnel, err := client.GetMagicTransitGRETunnel(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err != nil {
			return err
		}

		*tunnel = foundGRETunnel

		return nil
	}
}

func TestAccCloudflareGRETunnelUpdateDescription(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_gre_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	cfIP := utils.LookupMagicWanCfIP(t, accountID)
	// Use unique endpoints and interface address per test to avoid conflicts
	customerEndpoint := "203.0.113.11"
	interfaceAddr := "10.213.0.22/31"

	var Tunnel cloudflare.MagicTransitGRETunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareGRETunnelSimple(rnd, rnd, rnd, accountID, cfIP, customerEndpoint, interfaceAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
				),
			},
			{
				Config: testAccCheckCloudflareGRETunnelSimple(rnd, rnd, rnd+"-updated", accountID, cfIP, customerEndpoint, interfaceAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareGRETunnelUpdateMulti(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_gre_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	cfIP := utils.LookupMagicWanCfIP(t, accountID)
	// Use unique endpoints and interface address per test to avoid conflicts
	customerEndpoint1 := "203.0.113.12"
	interfaceAddr1 := "10.213.0.24/31"
	customerEndpoint2 := "203.0.113.13"
	interfaceAddr2 := "10.213.0.26/31"

	var Tunnel cloudflare.MagicTransitGRETunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareGRETunnelSimple(rnd, rnd, rnd, accountID, cfIP, customerEndpoint1, interfaceAddr1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "customer_gre_endpoint", "203.0.113.12"),
					resource.TestCheckResourceAttr(name, "cloudflare_gre_endpoint", cfIP),
					resource.TestCheckResourceAttr(name, "interface_address", "10.213.0.24/31"),
					resource.TestCheckResourceAttr(name, "ttl", "64"),
					resource.TestCheckResourceAttr(name, "mtu", "1476"),
					resource.TestCheckResourceAttr(name, "health_check.enabled", "true"),
				),
			},
			{
				Config: testAccCheckCloudflareGRETunnelMultiUpdate(rnd, rnd, rnd+"-updated", accountID, cfIP, customerEndpoint2, interfaceAddr2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "customer_gre_endpoint", "203.0.113.13"),
					resource.TestCheckResourceAttr(name, "cloudflare_gre_endpoint", cfIP),
					resource.TestCheckResourceAttr(name, "interface_address", "10.213.0.26/31"),
					resource.TestCheckResourceAttr(name, "ttl", "65"),
					resource.TestCheckResourceAttr(name, "mtu", "1475"),
					resource.TestCheckResourceAttr(name, "health_check.enabled", "false"),
				),
			},
		},
	})
}

func testAccCheckCloudflareGRETunnelSimple(ID, name, description, accountID, cfIP, customerEndpoint, interfaceAddr string) string {
	return acctest.LoadTestCase("gretunnelsimple.tf", ID, name, description, accountID, cfIP, customerEndpoint, interfaceAddr)
}

func testAccCheckCloudflareGRETunnelMultiUpdate(ID, name, description, accountID, cfIP, customerEndpoint, interfaceAddr string) string {
	return acctest.LoadTestCase("gretunnelmultiupdate.tf", ID, name, description, accountID, cfIP, customerEndpoint, interfaceAddr)
}
