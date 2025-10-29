package queue_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccCloudflareQueue_Settings_UpdateDeliveryPaused(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_queue." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t); acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareQueueWithPaused(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "queue_name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "settings.delivery_paused", "false"),
				),
			},
			{
				Config: testAccCheckCloudflareQueueWithPausedUpdate(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "queue_name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "settings.delivery_paused", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareQueue_Settings_UpdateRetention(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_queue." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t); acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareQueueWithRetention(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "queue_name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "settings.message_retention_period", "65"),
				),
			},
			{
				Config: testAccCheckCloudflareQueueWithRetentionUpdate(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "queue_name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "settings.message_retention_period", "60"),
				),
			},
		},
	})
}

func TestAccCloudflareQueue_Settings_UpdateDeliveryDelay(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_queue." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t); acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareQueueWithDeliveryDelay(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "queue_name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "settings.delivery_delay", "5"),
				),
			},
			{
				Config: testAccCheckCloudflareQueueWithDeliveryDelayUpdate(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "queue_name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "settings.delivery_delay", "10"),
				),
			},
		},
	})
}

func init() {
	// TODO: fixme
	//resource.AddTestSweepers("cloudflare_queue", &resource.Sweeper{
	//	Name: "cloudflare_queue",
	//	F:    testSweepCloudflareQueue,
	//})
}

func testSweepCloudflareQueue(r string) error {
	ctx := context.Background()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
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
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_queue." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t); acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareQueue(rnd, accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareQueueExists(rnd, &queue),
					resource.TestCheckResourceAttr(resourceName, "queue_name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
				),
			},
			{
				Config: testAccCheckCloudflareQueue(rnd, accountID, rnd+"-updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "queue_name", rnd+"-updated"),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
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
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

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
	return acctest.LoadTestCase("queue.tf", rnd, accountID, name)
}

func testAccCheckCloudflareQueueWithDeliveryDelay(rnd, accountID, name string) string {
	return acctest.LoadTestCase("queue_with_delivery_delay.tf", rnd, accountID, name)
}

func testAccCheckCloudflareQueueWithDeliveryDelayUpdate(rnd, accountID, name string) string {
	return acctest.LoadTestCase("queue_with_delivery_delay_update.tf", rnd, accountID, name)
}

func testAccCheckCloudflareQueueWithPaused(rnd, accountID, name string) string {
	return acctest.LoadTestCase("queue_with_paused.tf", rnd, accountID, name)
}

func testAccCheckCloudflareQueueWithPausedUpdate(rnd, accountID, name string) string {
	return acctest.LoadTestCase("queue_with_paused_update.tf", rnd, accountID, name)
}

func testAccCheckCloudflareQueueWithRetention(rnd, accountID, name string) string {
	return acctest.LoadTestCase("queue_with_retention.tf", rnd, accountID, name)
}

func testAccCheckCloudflareQueueWithRetentionUpdate(rnd, accountID, name string) string {
	return acctest.LoadTestCase("queue_with_retention_update.tf", rnd, accountID, name)
}

func testAccCheckCloudflareQueueExists(name string, queue *cloudflare.Queue) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}

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
