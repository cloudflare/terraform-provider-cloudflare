package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareWorkersQueue_Basic(t *testing.T) {
	t.Parallel()
	var queue cloudflare.Queue
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_workers_queue." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudflareWorkersQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersQueue(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersQueueExists(rnd, &queue),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
				),
			},
		},
	})
}

func testAccCloudflareWorkersQueueDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_queue" {
			continue
		}

		accountID := rs.Primary.Attributes["account_id"]
		if accountID == "" {
			accountID = client.AccountID
		}

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

func testAccCheckCloudflareWorkersQueue(rName string) string {
	return fmt.Sprintf(`
resource "cloudflare_workers_queue" "%[1]s" {
	title = "%[1]s"
}`, rName)
}

func testAccCheckCloudflareWorkersQueueExists(name string, queue *cloudflare.Queue) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)

		rs, ok := s.RootModule().Resources["cloudflare_workers_queue."+name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}
		accountID := rs.Primary.Attributes["account_id"]
		if accountID == "" {
			accountID = client.AccountID
		}
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
