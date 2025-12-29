package workers_route_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

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
		tflog.Info(ctx, "Skipping workers routes sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	// List all routes in the zone
	page, err := client.Workers.Routes.List(ctx, workers.RouteListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list workers routes: %s", err))
		return fmt.Errorf("failed to list workers routes: %w", err)
	}

	hasRoutes := false
	// Delete test routes based on pattern matching
	for page != nil {
		for _, route := range page.Result {
			// Only delete routes with patterns matching test patterns
			// Test routes use patterns like: cftftest*.cfapi.net/*
			if !strings.Contains(route.Pattern, "cftftest") && !strings.Contains(route.Pattern, ".cfapi.net") {
				tflog.Debug(ctx, fmt.Sprintf("Skipping non-test route: %s", route.Pattern))
				continue
			}

			hasRoutes = true
			tflog.Info(ctx, fmt.Sprintf("Deleting worker route: %s (zone: %s)", route.Pattern, zoneID))
			_, err := client.Workers.Routes.Delete(ctx, route.ID, workers.RouteDeleteParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete worker route %s: %s", route.ID, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted worker route: %s", route.ID))
		}

		page, err = page.GetNextPage()
		if err != nil {
			break
		}
	}

	if !hasRoutes {
		tflog.Info(ctx, "No workers routes to sweep")
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
