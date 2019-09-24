package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudflareWorkerRoute_Import(t *testing.T) {
	var route cloudflare.WorkerRoute
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	routeRnd := generateRandomResourceName()
	routeName := "cloudflare_worker_route." + routeRnd
	pattern := fmt.Sprintf("%s/%s", zone, generateRandomResourceName())

	// We also create a script in order to test routes since routes
	// need to point to a script
	scriptRnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWorkerRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerRouteConfigMultiScriptInitial(zoneID, routeRnd, scriptRnd, pattern),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerRouteExists(routeName, &route),
				),
			},
			{
				ResourceName:        routeName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerRouteExists(routeName, &route),
				),
			},
		},
	})
}
