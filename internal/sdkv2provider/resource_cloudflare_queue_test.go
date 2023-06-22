package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_queue", &resource.Sweeper{
		Name: "cloudflare_queue",
		F:    testSweepCloudflareQueue,
	})
}

func testSweepCloudflareQueue(r string) error {
	ctx := context.Background()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	resp, _, err := client.ListQueues(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListQueuesParams{})
	if err != nil {
		return err
	}

	for _, q := range resp {
		err := client.DeleteQueue(ctx, cloudflare.AccountIdentifier(accountID), q.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func TestAccCloudflareQueue_Basic(t *testing.T) {
	t.Parallel()
	var queue cloudflare.Queue
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_queue." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudflareQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareQueue(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareQueueExists(rnd, &queue),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
				),
			},
			{
				Config: testAccCheckCloudflareQueue(rnd, accountID, rnd+"-updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd+"-updated"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
				),
			},
			{
				ImportState:         true,
				ImportStateVerify:   true,
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func testAccCloudflareQueueDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_queue" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		resp, _, err := client.ListQueues(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListQueuesParams{})
		if err != nil {
			return err
		}

		for _, n := range resp {
			if n.ID == rs.Primary.ID {
				return fmt.Errorf("queue still exists but should not")
			}
		}
	}

	return nil
}

func testAccCheckCloudflareQueue(rnd, accountID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_queue" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
}`, rnd, accountID, name)
}

func testAccCheckCloudflareQueueExists(name string, queue *cloudflare.Queue) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)

		rs, ok := s.RootModule().Resources["cloudflare_queue."+name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		resp, _, err := client.ListQueues(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListQueuesParams{})
		if err != nil {
			return err
		}

		for _, q := range resp {
			if q.Name == name {
				*queue = q
				return nil
			}
		}

		return fmt.Errorf("queue not found")
	}
}
