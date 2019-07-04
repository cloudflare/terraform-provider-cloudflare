package cloudflare

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	scriptContent1 = `addEventListener('fetch', event => {event.respondWith(new Response('test 1'))});`
	scriptContent2 = `addEventListener('fetch', event => {event.respondWith(new Response('test 2'))});`
)

func TestAccCloudflareWorkerScript_SingleScriptNonEnt(t *testing.T) {
	// Temporarily unset CLOUDFLARE_ORG_ID if it is set in order
	// to test non-ENT behavior
	if os.Getenv("CLOUDFLARE_ORG_ID") != "" {
		defer func(orgId string) {
			os.Setenv("CLOUDFLARE_ORG_ID", orgId)
		}(os.Getenv("CLOUDFLARE_ORG_ID"))
		os.Setenv("CLOUDFLARE_ORG_ID", "")
	}

	testAccCloudflareWorkerScript_SingleScript(t, nil)
}

// ENT customers should still be able to use the single-script
// configuration format if they want to
func TestAccCloudflareWorkerScript_SingleScriptEnt(t *testing.T) {
	testAccCloudflareWorkerScript_SingleScript(t, testAccPreCheckOrg)
}

func testAccCloudflareWorkerScript_SingleScript(t *testing.T, preCheck preCheckFunc) {
	var script cloudflare.WorkerScript
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	name := "cloudflare_worker_script." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			if preCheck != nil {
				preCheck(t)
			}
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptConfigSingleScriptInitial(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script),
					resource.TestCheckResourceAttr(name, "zone", zone),
					resource.TestMatchResourceAttr(name, "zone_id", regexp.MustCompile("^[a-z0-9]{32}$")),
					resource.TestCheckResourceAttr(name, "content", scriptContent1),
					resource.TestCheckNoResourceAttr(name, "name"),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigSingleScriptUpdate(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script),
					resource.TestCheckResourceAttr(name, "zone", zone),
					resource.TestMatchResourceAttr(name, "zone_id", regexp.MustCompile("^[a-z0-9]{32}$")),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
					resource.TestCheckNoResourceAttr(name, "name"),
				),
			},
		},
	})
}

func testAccCheckCloudflareWorkerScriptConfigSingleScriptInitial(zone, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_script" "%[2]s" {
  zone = "%[1]s"
  content = "%[3]s"
}`, zone, rnd, scriptContent1)
}

func testAccCheckCloudflareWorkerScriptConfigSingleScriptUpdate(zone, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_script" "%[2]s" {
  zone = "%[1]s"
  content = "%[3]s"
}`, zone, rnd, scriptContent2)
}

func TestAccCloudflareWorkerScript_MultiScriptEnt(t *testing.T) {
	t.Parallel()

	var script cloudflare.WorkerScript
	rnd := generateRandomResourceName()
	name := "cloudflare_worker_script." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckOrg(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptInitial(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent1),
					resource.TestCheckNoResourceAttr(name, "zone"),
					resource.TestCheckNoResourceAttr(name, "zone_id"),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdate(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
					resource.TestCheckNoResourceAttr(name, "zone"),
					resource.TestCheckNoResourceAttr(name, "zone_id"),
				),
			},
		},
	})
}

func testAccCheckCloudflareWorkerScriptConfigMultiScriptInitial(rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  name = "%[1]s"
  content = "%[2]s"
}`, rnd, scriptContent1)
}

func testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdate(rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  name = "%[1]s"
  content = "%[2]s"
}`, rnd, scriptContent2)
}

func getRequestParamsFromResource(rs *terraform.ResourceState) cloudflare.WorkerRequestParams {
	var params cloudflare.WorkerRequestParams
	if rs.Primary.Attributes["name"] != "" {
		params = cloudflare.WorkerRequestParams{
			ScriptName: rs.Primary.Attributes["name"],
		}
	} else {
		params = cloudflare.WorkerRequestParams{
			ZoneID: rs.Primary.Attributes["zone_id"],
		}
	}

	return params
}

func testAccCheckCloudflareWorkerScriptExists(n string, script *cloudflare.WorkerScript) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Worker Script ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		params := getRequestParamsFromResource(rs)
		r, err := client.DownloadWorker(&params)
		if err != nil {
			return err
		}

		if r.Script == "" {
			return fmt.Errorf("Worker Script not found")
		}

		*script = r.WorkerScript
		return nil
	}
}

func testAccCheckCloudflareWorkerScriptDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_worker_script" {
			continue
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		params := getRequestParamsFromResource(rs)
		r, _ := client.DownloadWorker(&params)

		if r.Script != "" {
			return fmt.Errorf("Worker script with id %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
