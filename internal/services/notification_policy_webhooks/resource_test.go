package notification_policy_webhooks_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareNotificationPolicyWebhooks(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the notification
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy_webhooks." + rnd
	webhooksDestination := "https://example.com"
	updatedWebhooksName := "my updated webhooks destination for notifications"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCheckCloudflareNotificationPolicyWebhooks(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my webhooks destination for receiving Cloudflare notifications"),
					resource.TestCheckResourceAttr(resourceName, "url", webhooksDestination),
				),
			},
			{
				Config: testCheckCloudflareNotificationPolicyWebhooksUpdated(rnd, updatedWebhooksName, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedWebhooksName),
					resource.TestCheckResourceAttr(resourceName, "url", webhooksDestination),
					resource.TestCheckResourceAttr(resourceName, "type", "generic"),
				),
			},
		},
	})
}

func testCheckCloudflareNotificationPolicyWebhooks(name, accountID string) string {
	return acctest.LoadTestCase("checkcloudflarenotificationpolicywebhooks.tf", name, accountID)
}

func testCheckCloudflareNotificationPolicyWebhooksUpdated(resName, webhooksName, accountID string) string {
	return acctest.LoadTestCase("checkcloudflarenotificationpolicywebhooksupdated.tf", resName, webhooksName, accountID)
}
