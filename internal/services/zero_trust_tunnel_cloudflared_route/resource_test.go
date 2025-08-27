package zero_trust_tunnel_cloudflared_route_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_tunnel_cloudflared_route", &resource.Sweeper{
		Name: "cloudflare_zero_trust_tunnel_cloudflared_route",
		F:    testSweepCloudflareTunnelRoute,
	})
}

func testSweepCloudflareTunnelRoute(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}
	isDeleted := false
	tunnelRoutes, err := client.ListTunnelRoutes(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.TunnelRoutesListParams{IsDeleted: &isDeleted})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Tunnel Routes: %s", err))
	}

	if len(tunnelRoutes) == 0 {
		log.Print("[DEBUG] No Cloudflare Tunnel Routes to sweep")
		return nil
	}

	for _, tunnel := range tunnelRoutes {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Tunnel Route network: %s", tunnel.Network))
		//nolint:errcheck
		client.DeleteTunnelRoute(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.TunnelRoutesDeleteParams{Network: tunnel.Network, VirtualNetworkID: tunnel.TunnelID})
	}

	return nil
}

func TestAccCloudflareTunnelRoute_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Generate unique IP subnets to avoid conflicts
	subnet1 := fmt.Sprintf("10.%d.%d.10/32", utils.RandIntRange(10, 250), utils.RandIntRange(10, 250))
	subnet2 := fmt.Sprintf("20.%d.%d.20/32", utils.RandIntRange(10, 250), utils.RandIntRange(10, 250))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd, accountID, subnet1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttr(name, "network", subnet1),
					resource.TestCheckResourceAttr(name, "comment", rnd),
				),
			},
			// Update
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd, accountID, subnet2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("network"),
							knownvalue.StringExact(subnet2),
						),
					},
				},
				Check: resource.TestCheckResourceAttr(name, "network", subnet2),
			},
			// Re-applying same change does not produce drift
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd, accountID, subnet2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Import
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareTunnelRoute_UpdateComment(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Generate unique IP subnet to avoid conflicts
	subnet := fmt.Sprintf("10.%d.%d.10/32", utils.RandIntRange(10, 250), utils.RandIntRange(10, 250))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd, accountID, subnet),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "comment", rnd),
				),
			},
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd+"-updated", accountID, subnet),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "comment", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareTunnelRoute_NoComment(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Generate unique IP subnet to avoid conflicts
	cidr := fmt.Sprintf("10.%d.%d.1/32", utils.RandIntRange(10, 250), utils.RandIntRange(10, 250))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRouteNoComment(rnd, accountID, cidr),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttr(name, "network", cidr),
				),
			},
		},
	})
}

func TestAccCloudflareTunnelRoute_UpdateTunnel(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tunnel1 := fmt.Sprintf("%s_tun1", rnd)
	tunnel2 := fmt.Sprintf("%s_tun2", rnd)
	tunnel1Id := ""
	tunnel2Id := ""
	var tunnel1IdPtr *string = &tunnel1Id
	var tunnel2IdPtr *string = &tunnel2Id

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRouteUpdateTunnel(accountID, tunnel1, tunnel2, rnd, tunnel1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					getTunnelId(tunnel1, tunnel1IdPtr),
					getTunnelId(tunnel2, tunnel2IdPtr),
					resource.TestCheckResourceAttrPtr(name, "tunnel_id", tunnel1IdPtr),
				),
			},
			{
				Config: testAccCloudflareRouteUpdateTunnel(accountID, tunnel1, tunnel2, rnd, tunnel2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrPtr(name, "tunnel_id", tunnel2IdPtr),
				),
			},
		},
	})
}

func getTunnelId(tunnelName string, tunnel1Id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceName := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared.%s", tunnelName)

		// Get the resource instance from the state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		// Access the instance's attributes directly
		*tunnel1Id = rs.Primary.Attributes["id"]

		return nil
	}
}

func testAccCloudflareTunnelRouteSimple(ID, comment, accountID, network string) string {
	return acctest.LoadTestCase("route_simple.tf", ID, comment, accountID, network)
}

func testAccCloudflareRouteNoComment(ID, accountID, network string) string {
	return acctest.LoadTestCase("route_no_comment.tf", ID, accountID, network)
}

func testAccCloudflareRouteUpdateTunnel(accountID, tunnel1, tunnel2, ID, tunnel string) string {
	return acctest.LoadTestCase("route_update_tunnel.tf", accountID, tunnel1, tunnel2, ID, tunnel)
}
