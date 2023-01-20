package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareStaticRoute_Exists(t *testing.T) {
	skipMagicTransitTestForNonConfiguredDefaultZone(t)

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_static_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var StaticRoute cloudflare.MagicTransitStaticRoute

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd, accountID, 100),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "prefix", "10.100.0.0/24"),
					resource.TestCheckResourceAttr(name, "nexthop", "10.0.0.0"),
					resource.TestCheckResourceAttr(name, "priority", "100"),
					resource.TestCheckResourceAttr(name, "weight", "100"),
					resource.TestCheckResourceAttr(name, "colo_regions.0", "APAC"),
					resource.TestCheckResourceAttr(name, "colo_names.0", "den01"),
				),
			},
		},
	})
}

func testAccCheckCloudflareStaticRouteExists(n string, route *cloudflare.MagicTransitStaticRoute) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No static route is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundStaticRoute, err := client.GetMagicTransitStaticRoute(context.Background(), accountID, rs.Primary.ID)
		if err != nil {
			return err
		}

		*route = foundStaticRoute

		return nil
	}
}

func TestAccCloudflareStaticRoute_UpdateDescription(t *testing.T) {
	skipMagicTransitTestForNonConfiguredDefaultZone(t)

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_static_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var StaticRoute cloudflare.MagicTransitStaticRoute

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd, accountID, 100),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					resource.TestCheckResourceAttr(name, "description", rnd),
				),
			},
			{
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd+"-updated", accountID, 100),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareStaticRoute_UpdateWeight(t *testing.T) {
	skipMagicTransitTestForNonConfiguredDefaultZone(t)

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_static_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var StaticRoute cloudflare.MagicTransitStaticRoute
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd, accountID, 100),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					resource.TestCheckResourceAttr(name, "weight", "100"),
				),
			},
			{
				PreConfig: func() {
					initialID = StaticRoute.ID
				},
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd+"-updated", accountID, 200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					func(state *terraform.State) error {
						if initialID == StaticRoute.ID {
							return fmt.Errorf("forced recreation but Static Route got updated (id %q)", StaticRoute.ID)
						}
						return nil
					},
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
					resource.TestCheckResourceAttr(name, "weight", "200"),
				),
			},
		},
	})
}

func testAccCheckCloudflareStaticRouteSimple(ID, description, accountID string, weight int) string {
	return fmt.Sprintf(`
  resource "cloudflare_static_route" "%[1]s" {
	account_id = "%[3]s"
	prefix = "10.100.0.0/24"
	nexthop = "10.0.0.0"
	priority = "100"
	description = "%[2]s"
	weight = %[4]d
	colo_regions = ["APAC"]
	colo_names = ["den01"]
  }`, ID, description, accountID, weight)
}
