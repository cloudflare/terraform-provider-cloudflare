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
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed migration test configuration files
//
//go:embed testdata/v4_multiple_emails.tf
var v4MultipleEmailsConfig string

//go:embed testdata/v5_multiple_emails.tf
var v5MultipleEmailsConfig string

//go:embed testdata/v4_single_email.tf
var v4SingleEmailConfig string

//go:embed testdata/v5_single_email.tf
var v5SingleEmailConfig string

//go:embed testdata/v4_single_webhook.tf
var v4SingleWebhookConfig string

//go:embed testdata/v5_single_webhook.tf
var v5SingleWebhookConfig string

//go:embed testdata/v4_multiple_webhooks.tf
var v4MultipleWebhooksConfig string

//go:embed testdata/v5_multiple_webhooks.tf
var v5MultipleWebhooksConfig string

//go:embed testdata/v4_mixed_integrations.tf
var v4MixedIntegrationsConfig string

//go:embed testdata/v5_mixed_integrations.tf
var v5MixedIntegrationsConfig string

//go:embed testdata/v4_enabled_false.tf
var v4EnabledFalseConfig string

//go:embed testdata/v5_enabled_false.tf
var v5EnabledFalseConfig string

//go:embed testdata/v4_with_description.tf
var v4WithDescriptionConfig string

//go:embed testdata/v5_with_description.tf
var v5WithDescriptionConfig string

//go:embed testdata/v4_minimal.tf
var v4MinimalConfig string

//go:embed testdata/v5_minimal.tf
var v5MinimalConfig string

//go:embed testdata/v4_billing_filters.tf
var v4BillingFiltersConfig string

//go:embed testdata/v5_billing_filters.tf
var v5BillingFiltersConfig string

//go:embed testdata/v4_traffic_anomalies_filters.tf
var v4TrafficAnomaliesFiltersConfig string

//go:embed testdata/v5_traffic_anomalies_filters.tf
var v5TrafficAnomaliesFiltersConfig string

// TestMigrateNotificationPolicy_MultipleEmails tests notification policy migration with multiple email integrations
// Covers: Multiple email_integration blocks → mechanisms.email array
func TestMigrateNotificationPolicy_MultipleEmails(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4MultipleEmailsConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5MultipleEmailsConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("enabled"),
								knownvalue.Bool(true),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("alert_type"),
								knownvalue.StringExact("universal_ssl_event_type"),
							),
							// Verify multiple email integrations
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("email"),
								knownvalue.SetSizeExact(2),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("webhooks"),
								knownvalue.Null(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("pagerduty"),
								knownvalue.Null(),
							),
							// Verify timestamps
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("created"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("modified"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateNotificationPolicy_SingleEmail tests notification policy migration with single email integration
// Covers: Single email_integration block → mechanisms.email single-element array
func TestMigrateNotificationPolicy_SingleEmail(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4SingleEmailConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5SingleEmailConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("alert_type"),
								knownvalue.StringExact("universal_ssl_event_type"),
							),
							// Verify single email integration
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("email"),
								knownvalue.SetExact([]knownvalue.Check{
									knownvalue.ObjectExact(map[string]knownvalue.Check{
										"id": knownvalue.StringExact("single-test@example.com"),
									}),
								}),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("webhooks"),
								knownvalue.Null(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("pagerduty"),
								knownvalue.Null(),
							),
							// Verify timestamps
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("created"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("modified"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// Continue in next message due to length...

// TestMigrateNotificationPolicy_SingleWebhook tests notification policy migration with single webhook integration
// Covers: Single webhooks_integration block → mechanisms.webhooks, name field removal, UUID normalization quirk
func TestMigrateNotificationPolicy_SingleWebhook(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(accountID, rnd, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(accountID, rnd, name string) string {
				return fmt.Sprintf(v4SingleWebhookConfig, accountID, rnd, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(accountID, rnd, name string) string {
				return fmt.Sprintf(v5SingleWebhookConfig, accountID, rnd, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(accountID, rnd, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// v4 provider has UUID normalization quirk with webhooks
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config:             testConfig,
					ExpectNonEmptyPlan: true, // v4 has UUID normalization quirk
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("alert_type"),
								knownvalue.StringExact("universal_ssl_event_type"),
							),
							// Verify single webhook integration (name field removed)
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("email"),
								knownvalue.Null(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("webhooks"),
								knownvalue.SetSizeExact(1),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("pagerduty"),
								knownvalue.Null(),
							),
							// Verify timestamps
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("created"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("modified"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateNotificationPolicy_MultipleWebhooks tests notification policy migration with multiple webhook integrations
// Covers: Multiple webhooks_integration blocks → mechanisms.webhooks array with 3 elements
func TestMigrateNotificationPolicy_MultipleWebhooks(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(accountID, rnd, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(accountID, rnd, name string) string {
				return fmt.Sprintf(v4MultipleWebhooksConfig, accountID, rnd, accountID, rnd, accountID, rnd, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(accountID, rnd, name string) string {
				return fmt.Sprintf(v5MultipleWebhooksConfig, accountID, rnd, accountID, rnd, accountID, rnd, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(accountID, rnd, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// v4 provider has UUID normalization quirk with webhooks
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config:             testConfig,
					ExpectNonEmptyPlan: true, // v4 has UUID normalization quirk
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("alert_type"),
								knownvalue.StringExact("universal_ssl_event_type"),
							),
							// Verify three webhook integrations
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("email"),
								knownvalue.Null(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("webhooks"),
								knownvalue.SetSizeExact(3),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("pagerduty"),
								knownvalue.Null(),
							),
							// Verify timestamps
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("created"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("modified"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}


// TestMigrateNotificationPolicy_MixedIntegrations tests notification policy migration with mixed integration types
// Covers: Multiple email_integration + webhooks_integration → mechanisms with both email and webhooks arrays
func TestMigrateNotificationPolicy_MixedIntegrations(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(accountID, rnd, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(accountID, rnd, name string) string {
				return fmt.Sprintf(v4MixedIntegrationsConfig, accountID, rnd, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(accountID, rnd, name string) string {
				return fmt.Sprintf(v5MixedIntegrationsConfig, accountID, rnd, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(accountID, rnd, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// v4 provider has UUID normalization quirk with webhooks
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config:             testConfig,
					ExpectNonEmptyPlan: true, // v4 has UUID normalization quirk
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("alert_type"),
								knownvalue.StringExact("expiring_service_token_alert"),
							),
							// Verify mixed integrations: 2 emails + 1 webhook
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("email"),
								knownvalue.SetSizeExact(2),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("webhooks"),
								knownvalue.SetSizeExact(1),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("pagerduty"),
								knownvalue.Null(),
							),
							// Verify timestamps
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("created"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("modified"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateNotificationPolicy_EnabledFalse tests CRITICAL enabled=false preservation during migration
// Covers: enabled=false must be preserved (v5 defaults to true, so this is critical for disabled policies)
func TestMigrateNotificationPolicy_EnabledFalse(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(accountID, rnd, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(accountID, rnd, name string) string {
				return fmt.Sprintf(v4EnabledFalseConfig, accountID, rnd, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(accountID, rnd, name string) string {
				return fmt.Sprintf(v5EnabledFalseConfig, accountID, rnd, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(accountID, rnd, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// v4 provider has UUID normalization quirk with webhooks
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config:             testConfig,
					ExpectNonEmptyPlan: true, // v4 has UUID normalization quirk
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							// CRITICAL: enabled=false must be preserved (v5 defaults to true)
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("enabled"),
								knownvalue.Bool(false),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("alert_type"),
								knownvalue.StringExact("expiring_service_token_alert"),
							),
							// Verify webhook mechanism
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("email"),
								knownvalue.Null(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("webhooks"),
								knownvalue.SetSizeExact(1),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("pagerduty"),
								knownvalue.Null(),
							),
							// Verify timestamps
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("created"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("modified"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateNotificationPolicy_WithDescription tests notification policy migration with description field
// Covers: Description field preservation during migration
func TestMigrateNotificationPolicy_WithDescription(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4WithDescriptionConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5WithDescriptionConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							// Verify description field preserved
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("description"),
								knownvalue.StringExact("This is a comprehensive test description field for migration testing purposes"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("alert_type"),
								knownvalue.StringExact("universal_ssl_event_type"),
							),
							// Verify email mechanism
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("email"),
								knownvalue.SetExact([]knownvalue.Check{
									knownvalue.ObjectExact(map[string]knownvalue.Check{
										"id": knownvalue.StringExact("description-test@example.com"),
									}),
								}),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("webhooks"),
								knownvalue.Null(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("pagerduty"),
								knownvalue.Null(),
							),
							// Verify timestamps
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("created"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("modified"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateNotificationPolicy_Minimal tests notification policy migration with only required fields
// Covers: Minimal configuration, absence of optional fields (description, filters)
func TestMigrateNotificationPolicy_Minimal(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4MinimalConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5MinimalConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("enabled"),
								knownvalue.Bool(true),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("alert_type"),
								knownvalue.StringExact("universal_ssl_event_type"),
							),
							// Verify email mechanism
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("email"),
								knownvalue.SetExact([]knownvalue.Check{
									knownvalue.ObjectExact(map[string]knownvalue.Check{
										"id": knownvalue.StringExact("test-minimal@example.com"),
									}),
								}),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("webhooks"),
								knownvalue.Null(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("pagerduty"),
								knownvalue.Null(),
							),
							// Verify optional fields are null/absent
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("description"),
								knownvalue.Null(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("filters"),
								knownvalue.Null(),
							),
							// Verify timestamps
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("created"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("modified"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateNotificationPolicy_BillingFilters tests notification policy migration with billing usage alert filters
// Covers: filters block (v4 TypeList/MaxItems:1) → v5 SingleNestedAttribute, product/limit filter fields
func TestMigrateNotificationPolicy_BillingFilters(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(accountID, rnd, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(accountID, rnd, name string) string {
				return fmt.Sprintf(v4BillingFiltersConfig, accountID, rnd, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(accountID, rnd, name string) string {
				return fmt.Sprintf(v5BillingFiltersConfig, accountID, rnd, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(accountID, rnd, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// v4 provider has UUID normalization quirk with webhooks
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config:             testConfig,
					ExpectNonEmptyPlan: true, // v4 has UUID normalization quirk
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("alert_type"),
								knownvalue.StringExact("billing_usage_alert"),
							),
							// CRITICAL: Verify filters block (v4: TypeList/MaxItems:1) → v5: SingleNestedAttribute
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("filters"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("filters").AtMapKey("product"),
								knownvalue.ListSizeExact(1),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("filters").AtMapKey("limit"),
								knownvalue.ListSizeExact(1),
							),
							// Verify webhook mechanism
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("email"),
								knownvalue.Null(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("webhooks"),
								knownvalue.SetSizeExact(1),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("pagerduty"),
								knownvalue.Null(),
							),
							// Verify timestamps
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("created"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("modified"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateNotificationPolicy_TrafficAnomaliesFilters tests notification policy migration with complex traffic anomalies filters
// Covers: Complex filters with selectors, alert_trigger_preferences, group_by, where, zones
func TestMigrateNotificationPolicy_TrafficAnomaliesFilters(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name, zoneID string) string {
				return fmt.Sprintf(v4TrafficAnomaliesFiltersConfig, rnd, accountID, name, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name, zoneID string) string {
				return fmt.Sprintf(v5TrafficAnomaliesFiltersConfig, rnd, accountID, name, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
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
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("alert_type"),
								knownvalue.StringExact("traffic_anomalies_alert"),
							),
							// Verify complex filters transformation
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("filters"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("filters").AtMapKey("zones"),
								knownvalue.ListSizeExact(1),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("filters").AtMapKey("selectors"),
								knownvalue.ListSizeExact(1),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("filters").AtMapKey("alert_trigger_preferences"),
								knownvalue.ListSizeExact(1),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("filters").AtMapKey("group_by"),
								knownvalue.ListSizeExact(1),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("filters").AtMapKey("where"),
								knownvalue.ListSizeExact(1),
							),
							// Verify email mechanism
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("email"),
								knownvalue.SetExact([]knownvalue.Check{
									knownvalue.ObjectExact(map[string]knownvalue.Check{
										"id": knownvalue.StringExact("traffic-test@example.com"),
									}),
								}),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("webhooks"),
								knownvalue.Null(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("mechanisms").AtMapKey("pagerduty"),
								knownvalue.Null(),
							),
							// Verify timestamps
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("created"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_notification_policy."+rnd,
								tfjsonpath.New("modified"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

