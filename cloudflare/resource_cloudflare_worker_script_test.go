package cloudflare

import (
	"fmt"
	"strings"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	scriptContent1 = `addEventListener('fetch', event => {event.respondWith(new Response('test 1'))});`
	scriptContent2 = `addEventListener('fetch', event => {event.respondWith(new Response('test 2'))});`
)

func TestAccCloudflareWorkerScript_MultiScriptEnt(t *testing.T) {
	t.Parallel()

	var script cloudflare.WorkerScript
	rnd := generateRandomResourceName()
	name := "cloudflare_worker_script." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptInitial(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, nil),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent1),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdate(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, nil),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdateBinding(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, []string{"KV_NAMESPACE", "PLAIN_TEXT", "SECRET_TEXT"}),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdateKvNamespaceBinding(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, []string{rnd, fmt.Sprintf("%s-copy", rnd)}),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
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

func testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdateBinding(rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s" {
  title = "%[1]s"
}

resource "cloudflare_worker_script" "%[1]s" {
  name = "%[1]s"
  content = "%[2]s"

  binding {
    name = "KV_NAMESPACE"
    kv_namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
  }

  binding {
    name = "PLAIN_TEXT"
    plain_text = "%[1]s"
  }

  binding {
    name = "SECRET_TEXT"
    secret_text = "%[1]s"
  }
}`, rnd, scriptContent2)
}

func testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdateKvNamespaceBinding(rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s" {
  title = "%[1]s"
}

resource "cloudflare_worker_script" "%[1]s" {
  name = "%[1]s"
  content = "%[2]s"

  kv_namespace_binding {
	name = "%[1]s"
	namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
  }

  kv_namespace_binding {
	name = "%[1]s-copy"
	namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
  }
}`, rnd, scriptContent2)
}

func getRequestParamsFromResource(rs *terraform.ResourceState) cloudflare.WorkerRequestParams {
	params := cloudflare.WorkerRequestParams{
		ScriptName: rs.Primary.Attributes["name"],
	}

	return params
}

func testAccCheckCloudflareWorkerScriptExists(n string, script *cloudflare.WorkerScript, bindings []string) resource.TestCheckFunc {
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

		name := strings.Replace(n, "cloudflare_worker_script.", "", -1)
		foundBindings, err := getWorkerScriptBindings(name, client)
		if err != nil {
			return fmt.Errorf("cannot list script bindings: %v", err)
		}

		for _, binding := range bindings {
			if _, ok := foundBindings[binding]; !ok {
				return fmt.Errorf("cannot find binding with name %s", binding)
			}
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
