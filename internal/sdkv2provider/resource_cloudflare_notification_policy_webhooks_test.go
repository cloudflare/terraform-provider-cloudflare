package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareNotificationPolicyWebhooks(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the notification
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceName := "cloudflare_notification_policy_webhooks." + rnd
	webhooksDestination := "https://example.com"
	updatedWebhooksName := "my updated webhooks destination for notifications"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
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
	return fmt.Sprintf(`
  resource "cloudflare_notification_policy_webhooks" "%[1]s" {
	account_id  = "%[2]s"
    name        = "my webhooks destination for receiving Cloudflare notifications"
    url         = "https://example.com"
    secret      =  "my-secret-key"
  }`, name, accountID)
}

func testCheckCloudflareNotificationPolicyWebhooksUpdated(resName, webhooksName, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_notification_policy_webhooks" "%[1]s" {
	account_id  = "%[3]s"
    name        = "%[2]s"
	url         = "https://example.com"
  }`, resName, webhooksName, accountID)
}
