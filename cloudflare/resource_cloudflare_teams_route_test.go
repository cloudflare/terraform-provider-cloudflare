package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareTeamsRouteExists(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_teams_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var TunnelRoute cloudflare.TunnelRoute

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckAccount(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRouteSimple(rnd, rnd, accountID, "10.0.0.20/32"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareTeamsRouteExists(name, &TunnelRoute),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttrSet(name, "tunnel_id"),
					resource.TestCheckResourceAttr(name, "network", "10.0.0.20/32"),
					resource.TestCheckResourceAttr(name, "comment", rnd),
				),
			},
		},
	})
}

func testAccCheckCloudflareTeamsRouteExists(name string, route *cloudflare.TunnelRoute) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Teams route is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundTunnelRoute, err := client.GetTunnelRouteForIP(context.Background(), cloudflare.TunnelRoutesForIPParams{
			Network: rs.Primary.ID,
		})

		if err != nil {
			return err
		}

		*route = foundTunnelRoute

		return nil
	}
}

func TestAccCloudflareTeamsRouteUpdateComment(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_teams_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var TunnelRoute cloudflare.TunnelRoute

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckAccount(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRouteSimple(rnd, rnd, accountID, "10.0.0.10/32"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareTeamsRouteExists(name, &TunnelRoute),
					resource.TestCheckResourceAttr(name, "comment", rnd),
				),
			},
			{
				Config: testAccCloudflareTeamsRouteSimple(rnd, rnd+"-updated", accountID, "10.0.0.10/32"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareTeamsRouteExists(name, &TunnelRoute),
					resource.TestCheckResourceAttr(name, "comment", rnd+"-updated"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsRouteSimple(ID, comment, accountID, network string) string {
	return fmt.Sprintf(`
resource "cloudflare_argo_tunnel" "%[1]s" {
	account_id = "%[3]s"
	name       = "%[1]s"
	secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

resource "cloudflare_teams_route" "%[1]s" {
    account_id = "%[3]s"
    tunnel_id = cloudflare_argo_tunnel.%[1]s.id
    network = "%[4]s"
    comment = "%[2]s"
}`, ID, comment, accountID, network)
}
