package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	scriptContent1 = `addEventListener('fetch', event => {event.respondWith(new Response('test 1'))});`
	scriptContent2 = `addEventListener('fetch', event => {event.respondWith(new Response('test 2'))});`
	moduleContent  = `export default { fetch() { return new Response('Hello world'); }, };`
	encodedWasm    = "AGFzbQEAAAAGgYCAgAAA" // wat source: `(module)`, so literally just an empty wasm module
)

func TestAccCloudflareWorkerScript_MultiScriptEnt(t *testing.T) {
	t.Parallel()

	var script cloudflare.WorkerScript
	rnd := generateRandomResourceName()
	name := "cloudflare_worker_script." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
			testAccCheckCloudflareWorkerScriptCreateBucket(t, rnd)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptInitial(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, nil),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent1),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdate(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, nil),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdateBinding(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, []string{"MY_KV_NAMESPACE", "MY_PLAIN_TEXT", "MY_SECRET_TEXT", "MY_WASM", "MY_SERVICE_BINDING", "MY_BUCKET", "MY_QUEUE"}),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
				),
			},
		},
	})
}

func TestAccCloudflareWorkerScript_ModuleUpload(t *testing.T) {
	t.Parallel()

	var script cloudflare.WorkerScript
	rnd := generateRandomResourceName()
	name := "cloudflare_worker_script." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptUploadModule(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, nil),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", moduleContent),
				),
			},
		},
	})
}

// Create a bucket before creating a worker script binding.
// When a cloudflare_r2_bucket resource is added, we can switch to that instead
func testAccCheckCloudflareWorkerScriptCreateBucket(t *testing.T, rnd string) {
	client := testAccProvider.Meta().(*cloudflare.API)
	err := client.CreateR2Bucket(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.CreateR2BucketParameters{Name: rnd})
	if err != nil {
		t.Fatalf("unable to create test bucket named %s: %v", rnd, err)
	}
	t.Cleanup(func() {
		err := client.DeleteR2Bucket(context.Background(), cloudflare.AccountIdentifier(accountID), rnd)
		if err != nil {
			t.Errorf("Failed to clean up bucket named %s: %v", rnd, err)
		}
	})
}

func testAccCheckCloudflareWorkerScriptConfigMultiScriptInitial(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[3]s"
  name = "%[1]s"
  content = "%[2]s"
}`, rnd, scriptContent1, accountID)
}

func testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdate(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[3]s"
  name = "%[1]s"
  content = "%[2]s"
}`, rnd, scriptContent2, accountID)
}

func testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdateBinding(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s" {
	account_id = "%[4]s"
	title = "%[1]s"
}

resource "cloudflare_queue" "%[1]s" {
	account_id = "%[4]s"
	name = "%[1]s"
}

resource "cloudflare_worker_script" "%[1]s-service" {
	account_id = "%[4]s"
	name    = "%[1]s-service"
	content = "%[2]s"
}

resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[4]s"
  name    = "%[1]s"
  content = "%[2]s"

  kv_namespace_binding {
    name         = "MY_KV_NAMESPACE"
    namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
  }

  plain_text_binding {
    name = "MY_PLAIN_TEXT"
    text = "%[1]s"
  }

  secret_text_binding {
    name = "MY_SECRET_TEXT"
    text = "%[1]s"
  }

  webassembly_binding {
    name = "MY_WASM"
    module = "%[3]s"
  }

  r2_bucket_binding {
	name = "MY_BUCKET"
	bucket_name = "%[1]s"
  }

  service_binding {
	name = "MY_SERVICE_BINDING"
    service = cloudflare_worker_script.%[1]s-service.name
    environment = "production"
  }

  queue_binding {
    binding         = "MY_QUEUE"
    queue = cloudflare_queue.%[1]s.name
  }

}`, rnd, scriptContent2, encodedWasm, accountID)
}

func testAccCheckCloudflareWorkerScriptUploadModule(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[3]s"
  name = "%[1]s"
  content = "%[2]s"
  module = true
}`, rnd, moduleContent, accountID)
}

func testAccCheckCloudflareWorkerScriptExists(n string, script *cloudflare.WorkerScript, bindings []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Worker Script ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)

		r, err := client.GetWorker(context.Background(), cloudflare.AccountIdentifier(accountID), rs.Primary.Attributes["name"])
		if err != nil {
			return err
		}

		if r.Script == "" {
			return fmt.Errorf("Worker Script not found")
		}

		name := strings.Replace(n, "cloudflare_worker_script.", "", -1)
		foundBindings, err := getWorkerScriptBindings(context.Background(), accountID, name, client)
		if err != nil {
			return fmt.Errorf("cannot list script bindings: %w", err)
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
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_worker_script" {
			continue
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		r, _ := client.GetWorker(context.Background(), cloudflare.AccountIdentifier(accountID), rs.Primary.Attributes["name"])

		if r.Script != "" {
			return fmt.Errorf("worker script with id %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
