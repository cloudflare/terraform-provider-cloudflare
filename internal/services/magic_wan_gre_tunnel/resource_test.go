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

func TestAccCloudflareGRETunnelExists(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Not configured for Magic Transit.")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_gre_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	cfIP := utils.LookupMagicWanCfIP(t, accountID)
	config := testAccCheckCloudflareGRETunnelSimple(rnd, rnd, rnd, accountID, cfIP)

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
					resource.TestCheckResourceAttr(name, "customer_gre_endpoint", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "cloudflare_gre_endpoint", cfIP),
					resource.TestCheckResourceAttr(name, "interface_address", "10.213.0.9/31"),
					resource.TestCheckResourceAttr(name, "health_check.target.effective", "203.0.113.1"),
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
	acctest.TestAccSkipForDefaultZone(t, "Not configured for Magic Transit")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_gre_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	cfIP := utils.LookupMagicWanCfIP(t, accountID)

	var Tunnel cloudflare.MagicTransitGRETunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareGRETunnelSimple(rnd, rnd, rnd, accountID, cfIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
				),
			},
			{
				Config: testAccCheckCloudflareGRETunnelSimple(rnd, rnd, rnd+"-updated", accountID, cfIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareGRETunnelUpdateMulti(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Not configured for Magic Transit")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_gre_tunnel.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	cfIP := utils.LookupMagicWanCfIP(t, accountID)

	var Tunnel cloudflare.MagicTransitGRETunnel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareGRETunnelSimple(rnd, rnd, rnd, accountID, cfIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "customer_gre_endpoint", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "cloudflare_gre_endpoint", cfIP),
					resource.TestCheckResourceAttr(name, "interface_address", "10.213.0.9/31"),
					resource.TestCheckResourceAttr(name, "ttl", "64"),
					resource.TestCheckResourceAttr(name, "mtu", "1476"),
					resource.TestCheckResourceAttr(name, "health_check.enabled", "true"),
				),
			},
			{
				Config: testAccCheckCloudflareGRETunnelMultiUpdate(rnd, rnd, rnd+"-updated", accountID, cfIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareGRETunnelExists(name, &Tunnel),
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "customer_gre_endpoint", "203.0.113.2"),
					resource.TestCheckResourceAttr(name, "cloudflare_gre_endpoint", cfIP),
					resource.TestCheckResourceAttr(name, "interface_address", "10.213.0.11/31"),
					resource.TestCheckResourceAttr(name, "ttl", "65"),
					resource.TestCheckResourceAttr(name, "mtu", "1475"),
					resource.TestCheckResourceAttr(name, "health_check.enabled", "false"),
				),
			},
		},
	})
}

func testAccCheckCloudflareGRETunnelSimple(ID, name, description, accountID, cfIP string) string {
	return acctest.LoadTestCase("gretunnelsimple.tf", ID, name, description, accountID, cfIP)
}

func testAccCheckCloudflareGRETunnelMultiUpdate(ID, name, description, accountID, cfIP string) string {
	return acctest.LoadTestCase("gretunnelmultiupdate.tf", ID, name, description, accountID, cfIP)
}
