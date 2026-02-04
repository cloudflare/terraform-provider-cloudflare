package queue_consumer_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/queues"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_queue_consumer", &resource.Sweeper{
		Name: "cloudflare_queue_consumer",
		F:    testSweepCloudflareQueueConsumers,
	})
}

func testSweepCloudflareQueueConsumers(r string) error {
	ctx := context.Background()
	clientV1, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}
	clientV6 := acctest.SharedClient()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping queue consumers sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	// List all queues using v1 client
	queuesResp, _, err := clientV1.ListQueues(ctx, cfv1.AccountIdentifier(accountID), cfv1.ListQueuesParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch queues: %s", err))
		return fmt.Errorf("failed to fetch queues: %w", err)
	}

	if len(queuesResp) == 0 {
		tflog.Info(ctx, "No queues found, skipping queue consumers sweep")
		return nil
	}

	// For each queue, list and delete its consumers
	for _, queue := range queuesResp {
		// Only process queues that would be swept themselves
		if !utils.ShouldSweepResource(queue.Name) {
			continue
		}

		consumersPage, err := clientV6.Queues.Consumers.List(ctx, queue.ID, queues.ConsumerListParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to fetch consumers for queue %s: %s", queue.Name, err))
			continue
		}

		for _, consumer := range consumersPage.Result {
			tflog.Info(ctx, fmt.Sprintf("Deleting queue consumer: %s (queue: %s, account: %s)", consumer.ConsumerID, queue.Name, accountID))
			_, err := clientV6.Queues.Consumers.Delete(ctx, queue.ID, consumer.ConsumerID, queues.ConsumerDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete queue consumer %s: %s", consumer.ConsumerID, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted queue consumer: %s", consumer.ConsumerID))
		}
	}

	return nil
}

func TestAccCloudflareQueueConsumer_Worker_UpdateDeadLetterQueue(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	queueName := "test-queue-" + rnd
	dlq1 := "dlq-1-" + rnd
	dlq2 := "dlq-2-" + rnd
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
				Config: testAccCheckCloudflareQueueConsumerWorkerWithDeadLetter(rnd, accountID, queueName, dlq1),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("worker")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dead_letter_queue"), knownvalue.StringExact(dlq1)),
				},
			},
			{
				Config: testAccCheckCloudflareQueueConsumerWorkerWithDeadLetterUpdate(rnd, accountID, queueName, dlq1, dlq2),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("worker")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dead_letter_queue"), knownvalue.StringExact(dlq2)),
				},
			},
		},
	})
}

func TestAccCloudflareQueueConsumer_HttpPull_UpdateDeadLetterQueue(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	queueName := "test-queue-" + rnd
	dlq1 := "dlq-1-" + rnd
	dlq2 := "dlq-2-" + rnd
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
				Config: testAccCheckCloudflareQueueConsumerHttpPullWithDeadLetter(rnd, accountID, queueName, dlq1),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("http_pull")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dead_letter_queue"), knownvalue.StringExact(dlq1)),
				},
			},
			{
				Config: testAccCheckCloudflareQueueConsumerHttpPullWithDeadLetterUpdate(rnd, accountID, queueName, dlq1, dlq2),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("http_pull")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dead_letter_queue"), knownvalue.StringExact(dlq2)),
				},
			},
		},
	})
}

func TestAccCloudflareQueueConsumer_HttpPull_UpdateSettings(t *testing.T) {
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
				Config: testAccCheckCloudflareQueueConsumerHttpPullWithSettings(rnd, accountID, queueName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("http_pull")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("batch_size"), knownvalue.Float64Exact(10)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("max_retries"), knownvalue.Float64Exact(3)),
				},
			},
			{
				Config: testAccCheckCloudflareQueueConsumerHttpPullWithSettingsUpdate(rnd, accountID, queueName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("http_pull")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("batch_size"), knownvalue.Float64Exact(20)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("max_retries"), knownvalue.Float64Exact(5)),
				},
			},
		},
	})
}

func testSweepCloudflareQueueConsumer(r string) error {
	ctx := context.Background()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping queue consumers sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

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
		tflog.Error(ctx, fmt.Sprintf("Failed to list queues: %s", err))
		return fmt.Errorf("failed to list queues: %w", err)
	}

	if len(queueList.Result) == 0 {
		tflog.Info(ctx, "No queues to sweep consumers from")
		return nil
	}

	// For each queue, list and delete consumers
	for _, queue := range queueList.Result {
		consumers, err := client.Queues.Consumers.List(ctx, queue.QueueID, queues.ConsumerListParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to list consumers for queue %s: %s", queue.QueueID, err))
			continue
		}

		for _, consumer := range consumers.Result {
			tflog.Info(ctx, fmt.Sprintf("Deleting queue consumer: %s (queue: %s) (account: %s)", consumer.ConsumerID, queue.QueueID, accountID))
			_, err := client.Queues.Consumers.Delete(ctx, queue.QueueID, consumer.ConsumerID, queues.ConsumerDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete queue consumer %s: %s", consumer.ConsumerID, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted queue consumer: %s", consumer.ConsumerID))
		}
	}

	return nil
}

func TestAccCloudflareQueueConsumer_Worker(t *testing.T) {
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
				Config: testAccCheckCloudflareQueueConsumerWorker(rnd, accountID, queueName),
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

func TestAccCloudflareQueueConsumerWorker_WithSettings(t *testing.T) {
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
				Config: testAccCheckCloudflareQueueConsumerWorkerWithSettings(rnd, accountID, queueName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consumer_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("worker")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact("test-worker-consumer-worker-with-settings")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("batch_size"), knownvalue.Float64Exact(10)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("max_retries"), knownvalue.Float64Exact(3)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("max_wait_time_ms"), knownvalue.Float64Exact(5000)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareQueueConsumer_Worker_Update(t *testing.T) {
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
				Config: testAccCheckCloudflareQueueConsumerWorkerUpdateStart(rnd, accountID, queueName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consumer_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("worker")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact("test-worker-consumer-worker-update-start")),
				},
			},
			{
				Config: testAccCheckCloudflareQueueConsumerWorkerUpdate(rnd, accountID, queueName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("worker")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact("test-worker-consumer-worker-update")),
				},
			},
		},
	})
}

func TestAccCloudflareQueueConsumer_Worker_UpdateSettings(t *testing.T) {
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
				Config: testAccCheckCloudflareQueueConsumerWorkerWithSettingsUpdateStart(rnd, accountID, queueName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("batch_size"), knownvalue.Float64Exact(10)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("max_retries"), knownvalue.Float64Exact(3)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("max_wait_time_ms"), knownvalue.Float64Exact(5000)),
				},
			},
			{
				Config: testAccCheckCloudflareQueueConsumerWorkerWithSettingsUpdate(rnd, accountID, queueName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("batch_size"), knownvalue.Float64Exact(20)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("max_retries"), knownvalue.Float64Exact(5)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("max_wait_time_ms"), knownvalue.Float64Exact(8000)),
				},
			},
		},
	})
}

// TestAccCloudflareQueueConsumer_HttpPull_NoDiffAfterApply verifies that there is no diff
// when running terraform plan after an initial apply. This tests the fix for the issue
// where consumer_id was marked with RequiresReplace() which caused infinite replacements
// because the computed consumer_id value showed as "unknown" on subsequent plans.
// See: https://github.com/cloudflare/terraform-provider-cloudflare/issues/XXXX
func TestAccCloudflareQueueConsumer_HttpPull_NoDiffAfterApply(t *testing.T) {
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
			// Step 1: Initial apply
			{
				Config: testAccCheckCloudflareQueueConsumerHttpPullNoDiff(rnd, accountID, queueName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consumer_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("http_pull")),
				},
			},
			// Step 2: Same config - should produce no changes (empty plan)
			// This verifies that consumer_id doesn't cause unnecessary replacements
			{
				Config: testAccCheckCloudflareQueueConsumerHttpPullNoDiff(rnd, accountID, queueName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
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

func testAccCheckCloudflareQueueConsumerWorker(rnd, accountID, queueName string) string {
	return acctest.LoadTestCase("queueconsumer_worker.tf", accountID, queueName, rnd, accountID, accountID)
}

func testAccCheckCloudflareQueueConsumerWorkerUpdateStart(rnd, accountID, queueName string) string {
	return acctest.LoadTestCase("queueconsumer_worker_update_start.tf", accountID, queueName, rnd, accountID, accountID)
}

func testAccCheckCloudflareQueueConsumerWorkerUpdate(rnd, accountID, queueName string) string {
	return acctest.LoadTestCase("queueconsumer_worker_update.tf", accountID, queueName, rnd, accountID, accountID)
}

func testAccCheckCloudflareQueueConsumerWorkerWithSettings(rnd, accountID, queueName string) string {
	return acctest.LoadTestCase("queueconsumer_worker_with_settings.tf", accountID, queueName, rnd, accountID, accountID)
}

func testAccCheckCloudflareQueueConsumerWorkerWithSettingsUpdateStart(rnd, accountID, queueName string) string {
	return acctest.LoadTestCase("queueconsumer_worker_with_settings_update_start.tf", accountID, queueName, rnd, accountID, accountID)
}

func testAccCheckCloudflareQueueConsumerWorkerWithSettingsUpdate(rnd, accountID, queueName string) string {
	return acctest.LoadTestCase("queueconsumer_worker_with_settings_update.tf", accountID, queueName, rnd, accountID, accountID)
}

func testAccCheckCloudflareQueueConsumerHttpPull(rnd, accountID, queueName string) string {
	return acctest.LoadTestCase("queueconsumer_http.tf", accountID, queueName, rnd, accountID)
}

func testAccCheckCloudflareQueueConsumerHttpPullWithSettings(rnd, accountID, queueName string) string {
	return acctest.LoadTestCase("queueconsumer_http_with_settings.tf", accountID, queueName, rnd, accountID)
}

func testAccCheckCloudflareQueueConsumerHttpPullWithSettingsUpdate(rnd, accountID, queueName string) string {
	return acctest.LoadTestCase("queueconsumer_http_with_settings_update.tf", accountID, queueName, rnd, accountID)
}

func testAccCheckCloudflareQueueConsumerWorkerWithDeadLetter(rnd, accountID, queueName, dlqName string) string {
	return acctest.LoadTestCase("queueconsumer_worker_with_dead_letter.tf", accountID, queueName, accountID, dlqName, rnd, accountID, accountID)
}

func testAccCheckCloudflareQueueConsumerWorkerWithDeadLetterUpdate(rnd, accountID, queueName, dlqName1, dlqName2 string) string {
	return acctest.LoadTestCase("queueconsumer_worker_with_dead_letter_update.tf", accountID, queueName, accountID, dlqName1, accountID, dlqName2, rnd, accountID, accountID)
}

func testAccCheckCloudflareQueueConsumerHttpPullWithDeadLetter(rnd, accountID, queueName, dlqName string) string {
	return acctest.LoadTestCase("queueconsumer_http_with_dead_letter.tf", accountID, queueName, accountID, dlqName, rnd, accountID)
}

func testAccCheckCloudflareQueueConsumerHttpPullWithDeadLetterUpdate(rnd, accountID, queueName, dlqName1, dlqName2 string) string {
	return acctest.LoadTestCase("queueconsumer_http_with_dead_letter_update.tf", accountID, queueName, accountID, dlqName1, accountID, dlqName2, rnd, accountID)
}

func testAccCheckCloudflareQueueConsumerHttpPullNoDiff(rnd, accountID, queueName string) string {
	return acctest.LoadTestCase("queueconsumer_http_no_diff.tf", accountID, queueName, rnd, accountID)
}

func testAccCheckCloudflareQueueConsumerHttpPullReproduceBug(rnd, accountID, queueName string) string {
	return acctest.LoadTestCase("queueconsumer_http_reproduce_bug.tf", accountID, queueName, rnd, accountID)
}

// TestAccCloudflareQueueConsumer_HttpPull_ConsumerIdNoReplacement is a regression test
// that verifies the consumer_id computed field does not trigger resource replacement.
// Before the fix, consumer_id had RequiresReplace() which caused the resource to be
// destroyed and recreated on every terraform plan after initial apply.
func TestAccCloudflareQueueConsumer_HttpPull_ConsumerIdNoReplacement(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	queueName := "test-queue-" + rnd
	resourceName := "cloudflare_queue_consumer." + rnd

	var consumerIdStep1 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareQueueConsumerDestroy,
		Steps: []resource.TestStep{
			// Step 1: Initial apply - capture the consumer_id
			{
				Config: testAccCheckCloudflareQueueConsumerHttpPullReproduceBug(rnd, accountID, queueName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "consumer_id"),
					// Capture the consumer_id for comparison in step 2
					resource.TestCheckResourceAttrWith(resourceName, "consumer_id", func(value string) error {
						consumerIdStep1 = value
						return nil
					}),
				),
				// Allow non-empty plan due to settings/script drift (separate issue)
				// The key test is that there's NO REPLACEMENT in step 2
				ExpectNonEmptyPlan: true,
			},
			// Step 2: Same config - verify consumer_id is preserved (not recreated)
			// This would fail before the fix because consumer_id triggered replacement
			{
				Config: testAccCheckCloudflareQueueConsumerHttpPullReproduceBug(rnd, accountID, queueName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// Verify NO replacement action - this is the key assertion
						// Before the fix, this would show ResourceActionDestroyBeforeCreate
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					// Verify the consumer_id is the SAME as step 1 (no replacement occurred)
					resource.TestCheckResourceAttrWith(resourceName, "consumer_id", func(value string) error {
						if value != consumerIdStep1 {
							return fmt.Errorf("consumer_id changed from %s to %s - resource was replaced when it shouldn't have been", consumerIdStep1, value)
						}
						return nil
					}),
				),
				// Allow non-empty plan due to settings/script drift (separate issue)
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
