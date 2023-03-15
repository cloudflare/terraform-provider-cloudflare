package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_tunnel_route", &resource.Sweeper{
		Name: "cloudflare_tunnel_route",
		F:    testSweepCloudflareTunnelRoute,
	})
}

func testSweepCloudflareTunnelRoute(r string) error {
	ctx := context.Background()
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	tunnelRoutes, err := client.ListTunnelRoutes(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.TunnelRoutesListParams{})
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

func TestAccCloudflareTunnelRoute_Exists(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_tunnel_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var TunnelRoute cloudflare.TunnelRoute

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd, accountID, "10.0.0.20/32"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareTunnelRouteExists(name, &TunnelRoute),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttr(name, "network", "10.0.0.20/32"),
					resource.TestCheckResourceAttr(name, "comment", rnd),
				),
			},
		},
	})
}

func testAccCheckCloudflareTunnelRouteExists(name string, route *cloudflare.TunnelRoute) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Tunnel route is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundTunnelRoute, err := client.ListTunnelRoutes(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), cloudflare.TunnelRoutesListParams{
			IsDeleted:     cloudflare.BoolPtr(false),
			NetworkSubset: rs.Primary.ID,
		})

		if err != nil {
			return err
		}

		*route = foundTunnelRoute[0]

		return nil
	}
}

func TestAccCloudflareTunnelRoute_UpdateComment(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_tunnel_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var TunnelRoute cloudflare.TunnelRoute

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd, accountID, "10.0.0.10/32"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareTunnelRouteExists(name, &TunnelRoute),
					resource.TestCheckResourceAttr(name, "comment", rnd),
				),
			},
			{
				Config: testAccCloudflareTunnelRouteSimple(rnd, rnd+"-updated", accountID, "10.0.0.10/32"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareTunnelRouteExists(name, &TunnelRoute),
					resource.TestCheckResourceAttr(name, "comment", rnd+"-updated"),
				),
			},
		},
	})
}

func testAccCloudflareTunnelRouteSimple(ID, comment, accountID, network string) string {
	return fmt.Sprintf(`
resource "cloudflare_tunnel" "%[1]s" {
	account_id = "%[3]s"
	name       = "%[1]s"
	secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

resource "cloudflare_tunnel_route" "%[1]s" {
    account_id = "%[3]s"
    tunnel_id = cloudflare_tunnel.%[1]s.id
    network = "%[4]s"
    comment = "%[2]s"
}`, ID, comment, accountID, network)
}
