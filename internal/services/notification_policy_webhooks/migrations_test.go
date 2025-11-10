package notification_policy_webhooks_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateNotificationPolicyWebhooksBasic tests migration of a basic webhook from v4 to v5
func TestMigrateNotificationPolicyWebhooksBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the notification
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	webhookName := "tf-test-webhook-basic"
	webhookURL := "https://postman-echo.com/post"
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_notification_policy_webhooks" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  url        = "%[4]s"
}`, rnd, accountID, webhookName, webhookURL)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Resource name stays the same - cloudflare_notification_policy_webhooks
				statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(webhookName)),
				statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL)),
				statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("type"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateNotificationPolicyWebhooksWithSecret tests migration of a webhook with optional secret
func TestMigrateNotificationPolicyWebhooksWithSecret(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	webhookName := "tf-test-webhook-with-secret"
	webhookURL := "https://postman-echo.com/post"
	webhookSecret := "test-secret-12345"
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_notification_policy_webhooks" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  url        = "%[4]s"
  secret     = "%[5]s"
}`, rnd, accountID, webhookName, webhookURL, webhookSecret)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(webhookName)),
				statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL)),
				statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("secret"), knownvalue.StringExact(webhookSecret)),
				statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
			}),
		},
	})
}

