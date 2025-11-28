package workers_cron_trigger_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_workers_cron_trigger", &resource.Sweeper{
		Name: "cloudflare_workers_cron_trigger",
		F:    testSweepCloudflareWorkersCronTrigger,
	})
}

func testSweepCloudflareWorkersCronTrigger(r string) error {
	ctx := context.Background()
	// Workers Cron Trigger is a worker script-level configuration.
	// When worker scripts are swept, cron triggers are cleaned up automatically.
	// No sweeping required.
	tflog.Info(ctx, "Workers Cron Trigger doesn't require sweeping (worker script configuration)")
	return nil
}

func TestAccCloudflareWorkerCronTrigger_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_workers_cron_trigger.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkerCronTriggerConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "script_name", "mute-truth-fdb1"),
					resource.TestCheckResourceAttr(name, "schedules.#", "2"),
				),
			},
		},
	})
}

func testAccCloudflareWorkerCronTriggerConfigBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("workercrontriggerconfigbasic.tf", rnd, accountID)
}
