package cloudflare

import (
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	defaultScriptContent = "addEventListener('fetch', event => {event.respondWith(fetch(event.request))})"
)

func TestAccCloudflareWorkerRoute_SingleScriptNonEnt(t *testing.T) {
	// Temporarily unset CLOUDFLARE_ACCOUNT_ID if it is set in order
	// to test non-ENT behavior
	if os.Getenv("CLOUDFLARE_ACCOUNT_ID") != "" {
		defer func(accountId string) {
			os.Setenv("CLOUDFLARE_ACCOUNT_ID", accountId)
		}(os.Getenv("CLOUDFLARE_ACCOUNT_ID"))
		os.Setenv("CLOUDFLARE_ACCOUNT_ID", "")
	}

	testAccCloudflareWorkerRoute_SingleScript(t, nil)
}

// ENT customers should still be able to use the single-script
// configuration format if they want to
func TestAccCloudflareWorkerRoute_SingleScriptEnt(t *testing.T) {
	testAccCloudflareWorkerRoute_SingleScript(t, testAccPreCheckAccount)
}

func testAccCloudflareWorkerRoute_SingleScript(t *testing.T, preCheck preCheckFunc) {
	var route cloudflare.WorkerRoute
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	routeRnd := generateRandomResourceName()
	routeName := "cloudflare_worker_route." + routeRnd
	pattern1 := fmt.Sprintf("%s/%s", zoneName, generateRandomResourceName())
	pattern2 := fmt.Sprintf("%s/%s", zoneName, generateRandomResourceName())

	// We also create a script in order to test routes since routes
	// need to point to a script
	scriptRnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			if preCheck != nil {
				preCheck(t)
			}
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWorkerRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerRouteConfigSingleScriptInitial(zoneID, routeRnd, scriptRnd, pattern1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerRouteExists(routeName, &route),
					resource.TestCheckResourceAttr(routeName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(routeName, "pattern", pattern1),
					resource.TestCheckResourceAttr(routeName, "enabled", "true"),
					resource.TestCheckNoResourceAttr(routeName, "script_name"),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerRouteConfigSingleScriptUpdate(zoneID, routeRnd, scriptRnd, pattern2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerRouteExists(routeName, &route),
					resource.TestCheckResourceAttr(routeName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(routeName, "pattern", pattern2),
					resource.TestCheckResourceAttr(routeName, "enabled", "false"),
					resource.TestCheckNoResourceAttr(routeName, "script_name"),
				),
			},
		},
	})
}

func testAccCheckCloudflareWorkerRouteConfigSingleScriptInitial(zoneID, routeRnd, scriptRnd, pattern string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_route" "%[2]s" {
  zone_id = "%[1]s"
  pattern = "%[4]s"
  enabled = true
  depends_on = ["cloudflare_worker_script.%[3]s"]
}

resource "cloudflare_worker_script" "%[3]s" {
  zone_id = "%[1]s"
  content = "%[5]s"
}`, zoneID, routeRnd, scriptRnd, pattern, defaultScriptContent)
}

func testAccCheckCloudflareWorkerRouteConfigSingleScriptUpdate(zoneID, routeRnd, scriptRnd, pattern string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_route" "%[2]s" {
  zone_id = "%[1]s"
  pattern = "%[4]s"
  depends_on = ["cloudflare_worker_script.%[3]s"]
}

resource "cloudflare_worker_script" "%[3]s" {
  zone_id = "%[1]s"
  content = "%[5]s"
}`, zoneID, routeRnd, scriptRnd, pattern, defaultScriptContent)
}

func TestAccCloudflareWorkerRoute_MultiScriptEnt(t *testing.T) {
	t.Parallel()

	var route cloudflare.WorkerRoute
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	routeRnd := generateRandomResourceName()
	routeName := "cloudflare_worker_route." + routeRnd
	pattern1 := fmt.Sprintf("%s/%s", zoneName, generateRandomResourceName())
	pattern2 := fmt.Sprintf("%s/%s", zoneName, generateRandomResourceName())

	// We also create a script in order to test routes since routes
	// need to point to a script
	scriptRnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWorkerRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerRouteConfigMultiScriptInitial(zoneID, routeRnd, scriptRnd, pattern1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerRouteExists(routeName, &route),
					resource.TestCheckResourceAttr(routeName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(routeName, "pattern", pattern1),
					resource.TestCheckResourceAttr(routeName, "script_name", scriptRnd),
					resource.TestCheckNoResourceAttr(routeName, "enabled"),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerRouteConfigMultiScriptUpdate(zoneID, routeRnd, scriptRnd, pattern2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerRouteExists(routeName, &route),
					resource.TestCheckResourceAttr(routeName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(routeName, "pattern", pattern2),
					resource.TestCheckResourceAttr(routeName, "script_name", ""),
					resource.TestCheckNoResourceAttr(routeName, "enabled"),
				),
			},
		},
	})
}

func testAccCheckCloudflareWorkerRouteConfigMultiScriptInitial(zoneID, routeRnd, scriptRnd, pattern string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_route" "%[2]s" {
  zone_id = "%[1]s"
  pattern = "%[4]s"
  script_name = "${cloudflare_worker_script.%[3]s.name}"
}

resource "cloudflare_worker_script" "%[3]s" {
  name = "%[3]s"
  content = "%[5]s"
}`, zoneID, routeRnd, scriptRnd, pattern, defaultScriptContent)
}

func testAccCheckCloudflareWorkerRouteConfigMultiScriptUpdate(zoneID, routeRnd, scriptRnd, pattern string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_route" "%[2]s" {
  zone_id = "%[1]s"
  pattern = "%[4]s"
}

resource "cloudflare_worker_script" "%[3]s" {
  name = "%[3]s"
  content = "%[5]s"
}`, zoneID, routeRnd, scriptRnd, pattern, defaultScriptContent)
}

func TestAccCloudflareWorkerRoute_MultiScriptDisabledRoute(t *testing.T) {
	t.Parallel()

	var route cloudflare.WorkerRoute
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	routeRnd := generateRandomResourceName()
	routeName := "cloudflare_worker_route." + routeRnd
	pattern := fmt.Sprintf("%s/%s", zoneName, generateRandomResourceName())

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWorkerRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerRouteConfigMultiScriptDisabledRoute(zoneID, routeRnd, pattern),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerRouteExists(routeName, &route),
					resource.TestCheckResourceAttr(routeName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(routeName, "pattern", pattern),
					resource.TestCheckNoResourceAttr(routeName, "script_name"),
					resource.TestCheckNoResourceAttr(routeName, "enabled"),
				),
			},
		},
	})
}

func testAccCheckCloudflareWorkerRouteConfigMultiScriptDisabledRoute(zoneID, routeRnd, pattern string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_route" "%[2]s" {
  zone_id = "%[1]s"
  pattern = "%[3]s"
}`, zoneID, routeRnd, pattern)
}

func getRouteFromApi(zoneID, routeId string) (cloudflare.WorkerRoute, error) {
	if zoneID == "" {
		return cloudflare.WorkerRoute{}, fmt.Errorf("zoneID is required to get a route")
	}
	if routeId == "" {
		return cloudflare.WorkerRoute{}, fmt.Errorf("routeId is required to get a route")
	}

	client := testAccProvider.Meta().(*cloudflare.API)
	resp, err := client.ListWorkerRoutes(zoneID)
	if err != nil {
		return cloudflare.WorkerRoute{}, err
	}

	var route cloudflare.WorkerRoute
	for _, r := range resp.Routes {
		if r.ID == routeId {
			route = r
			break
		}
	}

	return route, nil
}

func testAccCheckCloudflareWorkerRouteExists(n string, route *cloudflare.WorkerRoute) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Worker Route ID is set")
		}

		zoneID := rs.Primary.Attributes["zone_id"]
		routeId := rs.Primary.ID
		foundRoute, err := getRouteFromApi(zoneID, routeId)
		if err != nil {
			return err
		}

		if foundRoute.ID != routeId {
			return fmt.Errorf("Worker route with id %s not found", routeId)
		}

		*route = foundRoute
		return nil
	}
}

func testAccCheckCloudflareWorkerRouteDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_worker_route" {
			continue
		}

		zoneID := rs.Primary.Attributes["zone_id"]
		routeId := rs.Primary.ID
		route, err := getRouteFromApi(zoneID, routeId)

		if err != nil {
			return err
		}

		if route.ID != "" {
			return fmt.Errorf("Worker route with id %s still exists", route.ID)
		}

	}

	return nil
}
