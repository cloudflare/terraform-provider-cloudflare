package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareWorkerCronTrigger_Basic(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_worker_cron_trigger.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkerCronTriggerConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "script_name", rnd),
					resource.TestCheckResourceAttr(name, "schedules.#", "2"),
				),
			},
		},
	})
}

func testAccCloudflareWorkerCronTriggerConfigBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
	account_id  = "%[2]s"
	name = "%[1]s"
	content = "addEventListener('fetch', event => {event.respondWith(new Response('test'))});"
}

resource "cloudflare_worker_cron_trigger" "%[1]s" {
	account_id  = "%[2]s"
	script_name = cloudflare_worker_script.%[1]s.name
	schedules   = [
		"*/5 * * * *",      # every 5 minutes
		"10 7 * * mon-fri", # 7:10am every weekday
	]
}
`, rnd, accountID)
}
