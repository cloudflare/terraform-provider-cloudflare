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

const testQueueName = "test-queue"

func init() {
	resource.AddTestSweepers("cloudflare_queue_consumer", &resource.Sweeper{
		Name: "cloudflare_queue_consumer",
		F:    testSweepCloudflareQueueConsumer,
	})
}

func testSweepCloudflareQueueConsumer(r string) error {
	ctx := context.Background()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	consumers, _, err := client.ListQueueConsumers(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListQueueConsumersParams{QueueName: testQueueName})
	if err != nil {
		return err
	}

	for _, consumer := range consumers {
		if err := client.DeleteQueueConsumer(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.DeleteQueueConsumerParams{QueueName: testQueueName, ConsumerName: consumer.ScriptName}); err != nil {
			return err
		}
	}

	if err := client.DeleteQueue(ctx, cloudflare.AccountIdentifier(accountID), testQueueName); err != nil {
		return err
	}

	return nil
}

func TestAccQueueConsumer_Basic(t *testing.T) {
	t.Parallel()

	var consumer cloudflare.QueueConsumer
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_queue_consumer." + rnd
	accountId := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudflareQueueConsumerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareQueueConsumer(rnd, accountId, testQueueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareQueueConsumerExists(rnd, &consumer),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
				),
			},
		},
	})
}

func testAccCloudflareQueueConsumerDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_queue_consumer" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		if accountID == "" {
			accountID = client.AccountID
		}

		consumers, _, err := client.ListQueueConsumers(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListQueueConsumersParams{QueueName: testQueueName})
		if err != nil {
			return err
		}

		for _, consumer := range consumers {
			if consumer.Name == rs.Primary.ID { // TODO: Use consumer ID when API supports it, as currently this won't work
				return fmt.Errorf("queue still exists but should not")
			}
		}
	}

	return nil
}

func testAccCheckCloudflareQueueConsumer(rnd, accountId, queueName string) string {
	return fmt.Sprintf(`
resource "cloudflare_queue_consumer" "%[1]s" {
	account_id = "%[2]s"
	queue_name = "%[3]s"
}`, rnd, accountId, queueName)
}

func testAccCheckCloudflareQueueConsumerExists(name string, consumer *cloudflare.QueueConsumer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)

		rs, ok := s.RootModule().Resources["cloudflare_queue_consumer."+name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		consumers, _, err := client.ListQueueConsumers(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListQueueConsumersParams{QueueName: testQueueName})
		if err != nil {
			return err
		}

		for _, c := range consumers {
			if c.Name == name {
				*consumer = c
				return nil
			}
		}

		return fmt.Errorf("consumer not found")
	}
}
