package magic_wan_static_route_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareStaticRoute_Exists(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Not configured for Magic Transit")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_static_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	cfIP := utils.LookupMagicWanCfIP(t, accountID)
	config := testAccCheckCloudflareStaticRouteSimple(rnd, rnd, accountID, 100, cfIP)

	var StaticRoute cloudflare.MagicTransitStaticRoute

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					resource.TestCheckResourceAttr(name, "description", rnd),
					resource.TestCheckResourceAttr(name, "prefix", "10.100.0.0/24"),
					resource.TestCheckResourceAttr(name, "nexthop", "10.214.0.8"),
					resource.TestCheckResourceAttr(name, "priority", "100"),
					resource.TestCheckResourceAttr(name, "weight", "100"),
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
				ImportStateVerifyIgnore: []string{"created_on", "modified_on"}, // MOR-1878
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

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundStaticRoute, err := client.GetMagicTransitStaticRoute(context.Background(), accountID, rs.Primary.ID)
		if err != nil {
			return err
		}

		*route = foundStaticRoute

		return nil
	}
}

func TestAccCloudflareStaticRoute_UpdateDescription(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Not configured for Magic Transit")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_static_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	cfIP := utils.LookupMagicWanCfIP(t, accountID)

	var StaticRoute cloudflare.MagicTransitStaticRoute

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd, accountID, 100, cfIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					resource.TestCheckResourceAttr(name, "description", rnd),
				),
			},
			{
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd+"-updated", accountID, 100, cfIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareStaticRoute_UpdateWeight(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Not configured for Magic Transit")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_static_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	cfIP := utils.LookupMagicWanCfIP(t, accountID)

	var StaticRoute cloudflare.MagicTransitStaticRoute
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd, accountID, 100, cfIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					resource.TestCheckResourceAttr(name, "weight", "100"),
				),
			},
			{
				PreConfig: func() {
					initialID = StaticRoute.ID
				},
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd+"-updated", accountID, 200, cfIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					func(state *terraform.State) error {
						if initialID != StaticRoute.ID {
							return fmt.Errorf("Static Route should be updated but foreced created instead (id %q -> %q)", initialID, StaticRoute.ID)
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

func testAccCheckCloudflareStaticRouteSimple(ID, description, accountID string, weight int, cfIP string) string {
	return acctest.LoadTestCase("staticroutesimple.tf", ID, description, accountID, weight, cfIP)
}
