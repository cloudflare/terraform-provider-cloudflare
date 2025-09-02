package notification_policy_webhooks_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareNotificationPolicyWebhooks_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the notification
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy_webhooks." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	webhookName := "my webhook destination for notifications"
	webhookURL := "https://httpbin.org/post"
	webhookSecret := "my-secret"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareNotificationPolicyWebhooksBasic(rnd, accountID, webhookName, webhookURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", webhookName),
					resource.TestCheckResourceAttr(resourceName, "url", webhookURL),
				),
			},
			{
				Config: testCloudflareNotificationPolicyWebhooksBasic(rnd, accountID, webhookName+" updated", webhookURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", webhookName+" updated"),
					resource.TestCheckResourceAttr(resourceName, "url", webhookURL),
				),
			},
			{
				Config: testCloudflareNotificationPolicyWebhooksBasicWithSecret(rnd, accountID, webhookName+" updated", webhookURL, webhookSecret),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", webhookName+" updated"),
					resource.TestCheckResourceAttr(resourceName, "url", webhookURL),
					resource.TestCheckResourceAttr(resourceName, "secret", webhookSecret),
				),
			},
			{
				Config: testCloudflareNotificationPolicyWebhooksBasicWithSecret(rnd, accountID, webhookName+" updated", webhookURL, webhookSecret+" updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", webhookName+" updated"),
					resource.TestCheckResourceAttr(resourceName, "url", webhookURL),
					resource.TestCheckResourceAttr(resourceName, "secret", webhookSecret+" updated"),
				),
			},
		},
	})
}

func testCloudflareNotificationPolicyWebhooksBasic(resourceID, accountID, name, url string) string {
	return acctest.LoadTestCase("basic.tf", resourceID, accountID, name, url)
}

func testCloudflareNotificationPolicyWebhooksBasicWithSecret(resourceID, accountID, name, url, secret string) string {
	return acctest.LoadTestCase("basic_with_secret.tf", resourceID, accountID, name, url, secret)
}
