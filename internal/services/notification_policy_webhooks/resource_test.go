package notification_policy_webhooks_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/alerting"
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
	resource.AddTestSweepers("cloudflare_notification_policy_webhooks", &resource.Sweeper{
		Name: "cloudflare_notification_policy_webhooks",
		F:    testSweepCloudflareNotificationPolicyWebhooks,
	})
}

func testSweepCloudflareNotificationPolicyWebhooks(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping notification policy webhooks sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	webhooks, err := client.Alerting.Destinations.Webhooks.List(ctx, alerting.DestinationWebhookListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch notification policy webhooks: %s", err))
		return fmt.Errorf("failed to fetch notification policy webhooks: %w", err)
	}

	if len(webhooks.Result) == 0 {
		tflog.Info(ctx, "No notification policy webhooks to sweep")
		return nil
	}

	for _, webhook := range webhooks.Result {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(webhook.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting notification policy webhook: %s (account: %s)", webhook.Name, accountID))
		_, err := client.Alerting.Destinations.Webhooks.Delete(ctx, webhook.ID, alerting.DestinationWebhookDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete notification policy webhook %s: %s", webhook.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted notification policy webhook: %s", webhook.ID))
	}

	return nil
}

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
	webhookURL := "https://postman-echo.com/post"
	webhookSecret := "my-secret"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareNotificationPolicyWebhooksDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareNotificationPolicyWebhooksBasic(rnd, accountID, webhookName, webhookURL),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(webhookName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
			{
				Config: testCloudflareNotificationPolicyWebhooksBasic(rnd, accountID, webhookName+" updated", webhookURL),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(webhookName+" updated")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(webhookName+" updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
			{
				Config: testCloudflareNotificationPolicyWebhooksBasicWithSecret(rnd, accountID, webhookName+" updated", webhookURL, webhookSecret),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("secret"), knownvalue.StringExact(webhookSecret)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(webhookName+" updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("secret"), knownvalue.StringExact(webhookSecret)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
			{
				Config: testCloudflareNotificationPolicyWebhooksBasicWithSecret(rnd, accountID, webhookName+" updated", webhookURL, webhookSecret+" updated"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("secret"), knownvalue.StringExact(webhookSecret+" updated")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(webhookName+" updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("secret"), knownvalue.StringExact(webhookSecret+" updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"secret", "created_at", "type"},
			},
		},
	})
}

func TestAccCloudflareNotificationPolicyWebhooks_MinimalRequired(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy_webhooks." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	webhookName := "minimal webhook"
	webhookURL := "https://postman-echo.com/post"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareNotificationPolicyWebhooksDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareNotificationPolicyWebhooksMinimal(rnd, accountID, webhookName, webhookURL),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(webhookName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("secret"), knownvalue.Null()),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"secret", "created_at", "type"},
			},
		},
	})
}

func TestAccCloudflareNotificationPolicyWebhooks_SecretRemoval(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy_webhooks." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	webhookName := "webhook secret test"
	webhookURL := "https://postman-echo.com/post"
	webhookSecret := "test-secret"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareNotificationPolicyWebhooksDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareNotificationPolicyWebhooksBasicWithSecret(rnd, accountID, webhookName, webhookURL, webhookSecret),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(webhookName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("secret"), knownvalue.StringExact(webhookSecret)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
			{
				Config: testCloudflareNotificationPolicyWebhooksMinimal(rnd, accountID, webhookName, webhookURL),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(webhookName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("secret"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"secret", "created_at", "type"},
			},
		},
	})
}

func TestAccCloudflareNotificationPolicyWebhooks_URLUpdate(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy_webhooks." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	webhookName := "url update test"
	webhookURL1 := "https://postman-echo.com/post"
	webhookURL2 := "https://postman-echo.com/put"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareNotificationPolicyWebhooksDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareNotificationPolicyWebhooksMinimal(rnd, accountID, webhookName, webhookURL1),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(webhookName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
			{
				Config: testCloudflareNotificationPolicyWebhooksMinimal(rnd, accountID, webhookName, webhookURL2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL2)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(webhookName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
				ExpectNonEmptyPlan: true,
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

func testCloudflareNotificationPolicyWebhooksMinimal(resourceID, accountID, name, url string) string {
	return acctest.LoadTestCase("minimal.tf", resourceID, accountID, name, url)
}

func testAccCheckCloudflareNotificationPolicyWebhooksDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_notification_policy_webhooks" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.Alerting.Destinations.Webhooks.Get(
			context.Background(),
			rs.Primary.ID,
			alerting.DestinationWebhookGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("notification policy webhook still exists")
		}
	}

	return nil
}
