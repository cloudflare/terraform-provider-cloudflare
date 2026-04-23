package notification_policy_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateNotificationPolicyFromV4Basic tests basic migration with minimal fields
func TestMigrateNotificationPolicyFromV4Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy." + rnd
	webhookName := rnd + "_webhook"
	tmpDir := t.TempDir()

	// Create webhook integration first, then reference it in the policy
	v4Config := fmt.Sprintf(`
resource "cloudflare_notification_policy_webhooks" "%[1]s" {
  account_id = "%[2]s"
  name       = "test-webhook-%[1]s"
  url        = "https://www.cloudflare.com/cdn-cgi/trace"
}

resource "cloudflare_notification_policy" "%[3]s" {
  account_id  = "%[2]s"
  name        = "test-%[3]s"
  description = "Basic notification policy for migration testing"
  enabled     = true
  alert_type  = "universal_ssl_event_type"

  webhooks_integration {
    id   = cloudflare_notification_policy_webhooks.%[1]s.id
    name = cloudflare_notification_policy_webhooks.%[1]s.name
  }
}`, webhookName, accountID, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				// Note: v4 has UUID normalization issues (stores without hyphens, API returns with hyphens)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: true, // v4 has UUID normalization quirk
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("test-%s", rnd))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Basic notification policy for migration testing")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("alert_type"), knownvalue.StringExact("universal_ssl_event_type")),
				// Verify webhooks_integration transformed to mechanisms.webhooks
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mechanisms").AtMapKey("webhooks"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mechanisms").AtMapKey("webhooks").AtSliceIndex(0).AtMapKey("id"), knownvalue.NotNull()),
				// Verify timestamp fields are preserved after migration
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateNotificationPolicyFromV4MultipleSameTypeIntegrations tests multiple webhooks integrations
func TestMigrateNotificationPolicyFromV4MultipleWebhookIntegrations(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy." + rnd
	webhook1 := rnd + "_webhook1"
	webhook2 := rnd + "_webhook2"
	webhook3 := rnd + "_webhook3"
	tmpDir := t.TempDir()

	// Create 3 webhook integrations first, then reference them in the policy
	v4Config := fmt.Sprintf(`
resource "cloudflare_notification_policy_webhooks" "%[1]s" {
  account_id = "%[4]s"
  name       = "test-webhook-%[1]s"
  url        = "https://www.cloudflare.com/cdn-cgi/trace"
}

resource "cloudflare_notification_policy_webhooks" "%[2]s" {
  account_id = "%[4]s"
  name       = "test-webhook-%[2]s"
  url        = "https://www.cloudflare.com/cdn-cgi/trace"
}

resource "cloudflare_notification_policy_webhooks" "%[3]s" {
  account_id = "%[4]s"
  name       = "test-webhook-%[3]s"
  url        = "https://www.cloudflare.com/cdn-cgi/trace"
}

resource "cloudflare_notification_policy" "%[5]s" {
  account_id  = "%[4]s"
  name        = "test-%[5]s-multi-webhook"
  description = "Multiple webhooks integration testing"
  enabled     = true
  alert_type  = "universal_ssl_event_type"

  webhooks_integration {
    id   = cloudflare_notification_policy_webhooks.%[1]s.id
    name = cloudflare_notification_policy_webhooks.%[1]s.name
  }

  webhooks_integration {
    id   = cloudflare_notification_policy_webhooks.%[2]s.id
    name = cloudflare_notification_policy_webhooks.%[2]s.name
  }

  webhooks_integration {
    id   = cloudflare_notification_policy_webhooks.%[3]s.id
    name = cloudflare_notification_policy_webhooks.%[3]s.name
  }
}`, webhook1, webhook2, webhook3, accountID, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				// Note: v4 has UUID normalization issues (stores without hyphens, API returns with hyphens)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: true, // v4 has UUID normalization quirk
			},
			// Step 2: Run migration and verify multiple webhook integrations in array
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Multiple webhooks integration testing")),
				// Verify all three webhook integrations preserved in array
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mechanisms").AtMapKey("webhooks"), knownvalue.ListSizeExact(3)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mechanisms").AtMapKey("webhooks").AtSliceIndex(0).AtMapKey("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mechanisms").AtMapKey("webhooks").AtSliceIndex(1).AtMapKey("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mechanisms").AtMapKey("webhooks").AtSliceIndex(2).AtMapKey("id"), knownvalue.NotNull()),
				// Verify timestamp fields are preserved after migration
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateNotificationPolicyFromV4EnabledFalse tests CRITICAL enabled=false preservation
func TestMigrateNotificationPolicyFromV4EnabledFalse(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy." + rnd
	webhookName := rnd + "_webhook"
	tmpDir := t.TempDir()

	// Create webhook integration first, then reference it in the policy
	v4Config := fmt.Sprintf(`
resource "cloudflare_notification_policy_webhooks" "%[1]s" {
  account_id = "%[2]s"
  name       = "test-webhook-%[1]s"
  url        = "https://www.cloudflare.com/cdn-cgi/trace"
}

resource "cloudflare_notification_policy" "%[3]s" {
  account_id  = "%[2]s"
  name        = "test-%[3]s-disabled"
  description = "Disabled policy testing (enabled=false preservation)"
  enabled     = false
  alert_type  = "universal_ssl_event_type"

  webhooks_integration {
    id   = cloudflare_notification_policy_webhooks.%[1]s.id
    name = cloudflare_notification_policy_webhooks.%[1]s.name
  }
}`, webhookName, accountID, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				// Note: v4 has UUID normalization issues (stores without hyphens, API returns with hyphens)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: true, // v4 has UUID normalization quirk
			},
			// Step 2: Run migration and verify enabled=false is preserved
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Disabled policy testing (enabled=false preservation)")),
				// CRITICAL: enabled=false must be preserved (v5 defaults to true)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mechanisms").AtMapKey("webhooks"), knownvalue.ListSizeExact(1)),
				// Verify timestamp fields are preserved after migration
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateNotificationPolicyFromV4IntegrationNameRemoval tests that integration 'name' field is removed in v5
func TestMigrateNotificationPolicyFromV4IntegrationNameRemoval(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy." + rnd
	webhookName := rnd + "_webhook"
	tmpDir := t.TempDir()

	// V4 config with both 'id' and 'name' in webhooks_integration
	// The 'name' field should be removed during migration to v5
	v4Config := fmt.Sprintf(`
resource "cloudflare_notification_policy_webhooks" "%[1]s" {
  account_id = "%[2]s"
  name       = "test-webhook-%[1]s"
  url        = "https://www.cloudflare.com/cdn-cgi/trace"
}

resource "cloudflare_notification_policy" "%[3]s" {
  account_id  = "%[2]s"
  name        = "test-%[3]s-name-removal"
  description = "Testing integration name field removal"
  enabled     = true
  alert_type  = "universal_ssl_event_type"

  webhooks_integration {
    id   = cloudflare_notification_policy_webhooks.%[1]s.id
    name = cloudflare_notification_policy_webhooks.%[1]s.name
  }
}`, webhookName, accountID, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				// Note: v4 has UUID normalization issues (stores without hyphens, API returns with hyphens)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: true, // v4 has UUID normalization quirk
			},
			// Step 2: Run migration and verify 'name' field is removed from integration
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Testing integration name field removal")),
				// Verify webhooks_integration transformed to mechanisms.webhooks
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mechanisms").AtMapKey("webhooks"), knownvalue.ListSizeExact(1)),
				// CRITICAL: Verify only 'id' field exists in v5 (no 'name' field)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mechanisms").AtMapKey("webhooks").AtSliceIndex(0).AtMapKey("id"), knownvalue.NotNull()),
				// Verify timestamp fields are preserved after migration
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateNotificationPolicyFromV4WithFilters tests filters block migration from v4 to v5
func TestMigrateNotificationPolicyFromV4WithFilters(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy." + rnd
	webhookName := rnd + "_webhook"
	tmpDir := t.TempDir()

	// V4 config with filters block (TypeList/MaxItems:1 in v4)
	// Uses billing_usage_alert which supports product and limit filters
	// This tests the structural migration: v4 block syntax → v5 attribute syntax
	v4Config := fmt.Sprintf(`
resource "cloudflare_notification_policy_webhooks" "%[1]s" {
  account_id = "%[2]s"
  name       = "test-webhook-%[1]s"
  url        = "https://www.cloudflare.com/cdn-cgi/trace"
}

resource "cloudflare_notification_policy" "%[3]s" {
  account_id  = "%[2]s"
  name        = "test-%[3]s-filters"
  description = "Testing filters block migration"
  enabled     = true
  alert_type  = "billing_usage_alert"

  filters {
    product = ["worker_requests"]
    limit   = ["100"]
  }

  webhooks_integration {
    id   = cloudflare_notification_policy_webhooks.%[1]s.id
    name = cloudflare_notification_policy_webhooks.%[1]s.name
  }
}`, webhookName, accountID, rnd)

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
				Config:             v4Config,
				ExpectNonEmptyPlan: true, // v4 has UUID normalization quirk
			},
			// Step 2: Run migration and verify filters structure is properly transformed
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Testing filters block migration")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("alert_type"), knownvalue.StringExact("billing_usage_alert")),
				// CRITICAL: Verify filters block (v4: TypeList/MaxItems:1) → v5: SingleNestedAttribute
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("filters"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("filters").AtMapKey("product"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("filters").AtMapKey("limit"), knownvalue.ListSizeExact(1)),
				// Verify webhooks integration
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mechanisms").AtMapKey("webhooks"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mechanisms").AtMapKey("webhooks").AtSliceIndex(0).AtMapKey("id"), knownvalue.NotNull()),
				// Verify timestamp fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified"), knownvalue.NotNull()),
			}),
		},
	})
}
