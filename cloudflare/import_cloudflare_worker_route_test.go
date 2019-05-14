package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudflareWorkerRoute_Import(t *testing.T) {
	var route cloudflare.WorkerRoute
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	routeRnd := acctest.RandString(10)
	routeName := "cloudflare_worker_route." + routeRnd
	pattern := fmt.Sprintf("%s/%s", zone, acctest.RandString(10))

	// We also create a script in order to test routes since routes
	// need to point to a script
	scriptRnd := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWorkerRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerRouteConfigMultiScriptInitial(zone, routeRnd, scriptRnd, pattern),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerRouteExists(routeName, &route),
				),
			},
			{
				ResourceName:        routeName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zone),
				ImportState:         true,
				ImportStateVerify:   true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerRouteExists(routeName, &route),
				),
			},
		},
	})
}
