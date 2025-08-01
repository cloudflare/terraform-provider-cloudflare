package queue_consumer_test

import (
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccQueueConsumer(t *testing.T) {

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_queue_consumer." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareQueueConsumerBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "queue_id", rnd),
					resource.TestCheckResourceAttr(resourceName, "script_name", "test_script"),
					resource.TestCheckResourceAttr(resourceName, "type", "worker"),
				),
			},
			{
				ResourceName: resourceName,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return strings.Join([]string{accountID, rnd, "default"}, "/"), nil
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCloudflareQueueConsumerBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("queueconsumerbasic.tf", rnd, accountID)
}
