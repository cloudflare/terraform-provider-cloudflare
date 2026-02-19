package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_with_secret.tf
var v4WithSecretConfig string

//go:embed testdata/v5_with_secret.tf
var v5WithSecretConfig string

// TestMigrateNotificationPolicyWebhooksBasic tests migration of a basic webhook from v4 to v5.
// This test verifies:
// 1. account_id, name, url, id, type are preserved
// 2. Resource name stays the same (no rename)
func TestMigrateNotificationPolicyWebhooksBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the notification
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	testCases := []struct {
		name           string
		version        string
		useDevProvider bool
		configFn       func(rnd, accountID, webhookURL string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, webhookURL string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID, "tf-test-webhook-basic", webhookURL)
			},
		},
		{
			name:           "from_v5",
			version:        currentProviderVersion,
			useDevProvider: true,
			configFn: func(rnd, accountID, webhookURL string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, "tf-test-webhook-basic", webhookURL)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			webhookURL := "https://postman-echo.com/post"
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, webhookURL)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var step1 resource.TestStep
			if tc.useDevProvider {
				step1 = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				step1 = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					step1,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Resource name stays the same - cloudflare_notification_policy_webhooks
							statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-test-webhook-basic")),
							statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL)),
							statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("type"), knownvalue.NotNull()),
						},
					),
				},
			})
		})
	}
}

// TestMigrateNotificationPolicyWebhooksWithSecret tests migration of a webhook with optional secret.
// This test verifies that the secret field is preserved through migration.
func TestMigrateNotificationPolicyWebhooksWithSecret(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	testCases := []struct {
		name           string
		version        string
		useDevProvider bool
		configFn       func(rnd, accountID, webhookURL, webhookSecret string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, webhookURL, webhookSecret string) string {
				return fmt.Sprintf(v4WithSecretConfig, rnd, accountID, "tf-test-webhook-with-secret", webhookURL, webhookSecret)
			},
		},
		{
			name:           "from_v5",
			version:        currentProviderVersion,
			useDevProvider: true,
			configFn: func(rnd, accountID, webhookURL, webhookSecret string) string {
				return fmt.Sprintf(v5WithSecretConfig, rnd, accountID, "tf-test-webhook-with-secret", webhookURL, webhookSecret)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			webhookURL := "https://postman-echo.com/post"
			webhookSecret := "test-secret-12345"
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, webhookURL, webhookSecret)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var step1 resource.TestStep
			if tc.useDevProvider {
				step1 = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				step1 = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					step1,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-test-webhook-with-secret")),
							statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("url"), knownvalue.StringExact(webhookURL)),
							statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("secret"), knownvalue.StringExact(webhookSecret)),
							statecheck.ExpectKnownValue("cloudflare_notification_policy_webhooks."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
						},
					),
				},
			})
		})
	}
}
