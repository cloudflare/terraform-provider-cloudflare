package workers_script_test

import (
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	scriptContent1    = `addEventListener('fetch', event => {event.respondWith(new Response('test 1'))});`
	scriptContent2    = `addEventListener('fetch', event => {event.respondWith(new Response('test 2'))});`
	moduleContent     = `export default { fetch() { return new Response('Hello world'); }, };`
	encodedWasm       = "AGFzbQEAAAAGgYCAgAAA" // wat source: `(module)`, so literally just an empty wasm module
	compatibilityDate = "2023-03-19"
	d1DatabaseID      = "ce8b95dc-b376-4ff8-9b9e-1801ed6d745d"
)

var (
	compatibilityFlags = []string{"nodejs_compat", "web_socket_compression"}
)

func TestAccCloudflareWorkerScript_ServiceWorker(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workers_script." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptConfigServiceWorkerInitial(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "script_name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent1),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigServiceWorkerUpdate(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "script_name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigServiceWorkerUpdateBinding(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "script_name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
				),
			},
		},
	})
}

func TestAccCloudflareWorkerScript_ModuleUpload(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workers_script." + rnd
	r2AccesKeyID := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_ID")
	r2AccesKeySecret := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_SECRET")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptUploadModule(rnd, accountID, r2AccesKeyID, r2AccesKeySecret),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "script_name", rnd),
					resource.TestCheckResourceAttr(name, "content", moduleContent),
					resource.TestCheckResourceAttr(name, "compatibility_date", compatibilityDate),
					resource.TestCheckResourceAttr(name, "compatibility_flags.#", "2"),
					resource.TestCheckResourceAttr(name, "compatibility_flags.0", compatibilityFlags[0]),
					// resource.TestCheckResourceAttr(name, "logpush", "true"),
					resource.TestCheckResourceAttr(name, "placement.mode", "smart"),
				),
			},
		},
	})
}

func testAccCheckCloudflareWorkerScriptConfigServiceWorkerInitial(rnd, accountID string) string {
	return acctest.LoadTestCase("service_worker_initial.tf", rnd, scriptContent1, accountID)
}

func testAccCheckCloudflareWorkerScriptConfigServiceWorkerUpdate(rnd, accountID string) string {
	return acctest.LoadTestCase("service_worker_update.tf", rnd, scriptContent2, accountID)
}

func testAccCheckCloudflareWorkerScriptConfigServiceWorkerUpdateBinding(rnd, accountID string) string {
	return acctest.LoadTestCase("service_worker_update_binding.tf", rnd, scriptContent2, encodedWasm, accountID)
}

func testAccCheckCloudflareWorkerScriptUploadModule(rnd, accountID, r2AccessKeyID, r2AccessKeySecret string) string {
	return acctest.LoadTestCase("module.tf", rnd, moduleContent, accountID, compatibilityDate, strings.Join(compatibilityFlags, `","`), r2AccessKeyID, r2AccessKeySecret, d1DatabaseID)
}
