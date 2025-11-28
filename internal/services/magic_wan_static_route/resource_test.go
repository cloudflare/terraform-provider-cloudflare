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

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_magic_wan_static_route", &resource.Sweeper{
		Name: "cloudflare_magic_wan_static_route",
		F:    testSweepCloudflareMagicWanStaticRoute,
	})
}

func testSweepCloudflareMagicWanStaticRoute(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return fmt.Errorf("failed to create Cloudflare client: %w", clientErr)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping static routes sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	tflog.Info(ctx, "Starting to list static routes for sweeping")
	routes, err := client.ListMagicTransitStaticRoutes(ctx, accountID)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch static routes: %s", err))
		return fmt.Errorf("failed to fetch static routes: %w", err)
	}

	if len(routes) == 0 {
		tflog.Info(ctx, "No static routes to sweep")
		return nil
	}

	tflog.Info(ctx, fmt.Sprintf("Found %d static routes to sweep", len(routes)))

	deletedCount := 0
	failedCount := 0

	for _, route := range routes {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(route.Description) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting static route: %s (%s)", route.Description, route.ID))

		_, err := client.DeleteMagicTransitStaticRoute(ctx, accountID, route.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete static route %s: %s", route.ID, err))
			failedCount++
			continue
		}
		
		deletedCount++
		tflog.Info(ctx, fmt.Sprintf("Successfully deleted static route: %s", route.ID))
	}

	tflog.Info(ctx, fmt.Sprintf("Completed sweeping static routes: deleted %d, failed %d", deletedCount, failedCount))
	return nil
}

func TestAccCloudflareStaticRoute_Exists(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_static_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	cfIP := utils.LookupMagicWanCfIP(t, accountID)
	// Generate unique interface address and corresponding nexthop to avoid conflicts
	subnetBase := utils.RandIntRange(1, 254)
	interfaceAddr := fmt.Sprintf("10.214.%d.9/31", subnetBase)
	nexthop := fmt.Sprintf("10.214.%d.8", subnetBase) // Peer IP in /31 subnet
	config := testAccCheckCloudflareStaticRouteSimple(rnd, rnd, accountID, 100, cfIP, interfaceAddr, nexthop)

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
					resource.TestCheckResourceAttr(name, "nexthop", nexthop),
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
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_static_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	cfIP := utils.LookupMagicWanCfIP(t, accountID)
	// Generate unique interface address and corresponding nexthop to avoid conflicts
	subnetBase := utils.RandIntRange(1, 254)
	interfaceAddr := fmt.Sprintf("10.214.%d.9/31", subnetBase)
	nexthop := fmt.Sprintf("10.214.%d.8", subnetBase) // Peer IP in /31 subnet

	var StaticRoute cloudflare.MagicTransitStaticRoute

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd, accountID, 100, cfIP, interfaceAddr, nexthop),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					resource.TestCheckResourceAttr(name, "description", rnd),
				),
			},
			{
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd+"-updated", accountID, 100, cfIP, interfaceAddr, nexthop),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareStaticRoute_UpdateWeight(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_magic_wan_static_route.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	cfIP := utils.LookupMagicWanCfIP(t, accountID)
	// Generate unique interface address and corresponding nexthop to avoid conflicts
	subnetBase := utils.RandIntRange(1, 254)
	interfaceAddr := fmt.Sprintf("10.214.%d.9/31", subnetBase)
	nexthop := fmt.Sprintf("10.214.%d.8", subnetBase) // Peer IP in /31 subnet

	var StaticRoute cloudflare.MagicTransitStaticRoute
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd, accountID, 100, cfIP, interfaceAddr, nexthop),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareStaticRouteExists(name, &StaticRoute),
					resource.TestCheckResourceAttr(name, "weight", "100"),
				),
			},
			{
				PreConfig: func() {
					initialID = StaticRoute.ID
				},
				Config: testAccCheckCloudflareStaticRouteSimple(rnd, rnd+"-updated", accountID, 200, cfIP, interfaceAddr, nexthop),
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

func testAccCheckCloudflareStaticRouteSimple(ID, description, accountID string, weight int, cfIP, interfaceAddr, nexthop string) string {
	return acctest.LoadTestCase("staticroutesimple.tf", ID, description, accountID, weight, cfIP, interfaceAddr, nexthop)
}
