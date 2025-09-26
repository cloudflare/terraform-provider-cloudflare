package queue_consumer_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/queues"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
}

func testSweepCloudflareQueueConsumer(r string) error {
	ctx := context.Background()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	client := acctest.SharedClient()
	if client == nil {
		tflog.Error(ctx, "Failed to create Cloudflare client")
		return fmt.Errorf("failed to create Cloudflare client")
	}

	// List all queues first
	queueList, err := client.Queues.List(ctx, queues.QueueListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		return fmt.Errorf("failed to list queues: %w", err)
	}

	// For each queue, list and delete consumers
	for _, queue := range queueList.Result {
		consumers, err := client.Queues.Consumers.List(ctx, queue.QueueID, queues.ConsumerListParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			continue // Skip if we can't list consumers
		}

		for _, consumer := range consumers.Result {
			_, err := client.Queues.Consumers.Delete(ctx, queue.QueueID, consumer.ConsumerID, queues.ConsumerDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete queue consumer %s: %s", consumer.ConsumerID, err))
			}
		}
	}

	return nil
}

func TestAccCloudflareQueueConsumer_Basic(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	queueName := "test-queue-" + rnd
	resourceName := "cloudflare_queue_consumer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareQueueConsumerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareQueueConsumerBasic(rnd, accountID, queueName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consumer_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("worker")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact("test-worker")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareQueueConsumer_HttpPull(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	queueName := "test-queue-" + rnd
	resourceName := "cloudflare_queue_consumer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareQueueConsumerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareQueueConsumerHttpPull(rnd, accountID, queueName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consumer_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("http_pull")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareQueueConsumer_WithSettings(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	queueName := "test-queue-" + rnd
	resourceName := "cloudflare_queue_consumer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareQueueConsumerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareQueueConsumerWithSettings(rnd, accountID, queueName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consumer_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("worker")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact("test-worker")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings.batch_size"), knownvalue.Float64Exact(10)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings.max_retries"), knownvalue.Float64Exact(3)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings.max_wait_time_ms"), knownvalue.Float64Exact(5000)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func testAccCheckCloudflareQueueConsumerDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_queue_consumer" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		queueID := rs.Primary.Attributes["queue_id"]
		consumerID := rs.Primary.Attributes["consumer_id"]

		_, err := client.Queues.Consumers.Get(
			context.Background(),
			queueID,
			consumerID,
			queues.ConsumerGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("queue consumer still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareQueueConsumerBasic(rnd, accountID, queueName string) string {
	return fmt.Sprintf(`
resource "cloudflare_queue" "test_queue" {
  account_id = "%s"
  name       = "%s"
}

resource "cloudflare_queue_consumer" "%s" {
  account_id  = "%s"
  queue_id    = cloudflare_queue.test_queue.id
  type        = "worker"
  script_name = "test-worker"
}
`, accountID, queueName, rnd, accountID)
}

func testAccCheckCloudflareQueueConsumerHttpPull(rnd, accountID, queueName string) string {
	return fmt.Sprintf(`
resource "cloudflare_queue" "test_queue" {
  account_id = "%s"
  name       = "%s"
}

resource "cloudflare_queue_consumer" "%s" {
  account_id = "%s"
  queue_id   = cloudflare_queue.test_queue.id
  type       = "http_pull"
}
`, accountID, queueName, rnd, accountID)
}

func testAccCheckCloudflareQueueConsumerWithSettings(rnd, accountID, queueName string) string {
	return fmt.Sprintf(`
resource "cloudflare_queue" "test_queue" {
  account_id = "%s"
  name       = "%s"
}

resource "cloudflare_queue_consumer" "%s" {
  account_id  = "%s"
  queue_id    = cloudflare_queue.test_queue.id
  type        = "worker"
  script_name = "test-worker"
  
  settings {
    batch_size        = 10
    max_retries       = 3
    max_wait_time_ms  = 5000
  }
}
`, accountID, queueName, rnd, accountID)
}
