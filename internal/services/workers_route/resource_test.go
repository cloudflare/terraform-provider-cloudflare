package workers_route_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func init() {
	resource.AddTestSweepers("cloudflare_workers_route", &resource.Sweeper{
		Name: "cloudflare_workers_route",
		F:    testSweepCloudflareWorkersRoute,
	})
}

func testSweepCloudflareWorkersRoute(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	
	if zoneID == "" {
		// Skip sweeping if no zone ID is set
		return nil
	}

	// List all routes in the zone
	page, err := client.Workers.Routes.List(ctx, workers.RouteListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		return fmt.Errorf("failed to list workers routes: %w", err)
	}

	// Delete all routes in the test zone
	// Note: In a test environment, we assume all routes can be deleted
	for page != nil {
		for _, route := range page.Result {
			_, err := client.Workers.Routes.Delete(ctx, route.ID, workers.RouteDeleteParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				// Log but continue sweeping other routes
				continue
			}
		}
		
		page, err = page.GetNextPage()
		if err != nil {
			break
		}
	}

	return nil
}

func TestAccCloudflareWorkersRoute_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_workers_route." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkersRouteConfig(rnd, accountID, zoneID, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.%s/*", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.StringExact(rnd)),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccCloudflareWorkersRouteConfigUpdate(rnd, accountID, zoneID, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s-updated.%s/*", rnd, domain))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.StringExact(rnd+"-updated")),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
			},
		},
	})
}

func testAccCheckCloudflareWorkersRouteDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_route" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		_, err := client.Workers.Routes.Get(context.Background(), rs.Primary.ID, workers.RouteGetParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err == nil {
			return fmt.Errorf("workers route still exists")
		}
	}

	return nil
}

func testAccCloudflareWorkersRouteConfig(rnd, accountID, zoneID, domain string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID, zoneID, domain)
}

func testAccCloudflareWorkersRouteConfigUpdate(rnd, accountID, zoneID, domain string) string {
	return acctest.LoadTestCase("basic_update.tf", rnd, accountID, zoneID, domain)
}

func TestAccCloudflareWorkersRoute_NoScript(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_workers_route." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkersRouteConfigNoScript(rnd, zoneID, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.%s/*", rnd, domain))),
					// Script should be null/empty when not set
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script"), knownvalue.Null()),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
			},
		},
	})
}

func testAccCloudflareWorkersRouteConfigNoScript(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("no_script.tf", rnd, zoneID, domain)
}