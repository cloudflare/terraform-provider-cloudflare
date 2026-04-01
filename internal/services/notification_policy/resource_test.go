package notification_policy_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_notification_policy", &resource.Sweeper{
		Name: "cloudflare_notification_policy",
		F:    testSweepCloudflareNotificationPolicies,
	})
}

func testSweepCloudflareNotificationPolicies(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping notification policies sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	policiesResp, err := client.ListNotificationPolicies(ctx, accountID)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch notification policies: %s", err))
		return fmt.Errorf("failed to fetch notification policies: %w", err)
	}

	if len(policiesResp.Result) == 0 {
		tflog.Info(ctx, "No notification policies to sweep")
		return nil
	}

	for _, policy := range policiesResp.Result {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(policy.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting notification policy: %s (account: %s)", policy.Name, accountID))
		_, err := client.DeleteNotificationPolicy(ctx, accountID, policy.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete notification policy %s: %s", policy.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted notification policy: %s", policy.ID))
	}

	return nil
}

func TestAccCloudflareNotificationPolicy_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the IP List
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy." + rnd
	updatedPolicyName := "updated test SSL policy from terraform provider"
	updatedPolicyDesc := "updated description"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCheckCloudflareNotificationPolicy(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test SSL policy from terraform provider"),
					resource.TestCheckResourceAttr(resourceName, "description", "test description"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "alert_type", "universal_ssl_event_type"),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "mechanisms.email.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "mechanisms.email.0.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "mechanisms.email.1.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "mechanisms.email.0.id", "test2@example.com"),
					resource.TestCheckResourceAttr(resourceName, "mechanisms.email.1.id", "test@example.com"),
				),
			},
			{
				Config: testCheckCloudflareNotificationPolicyUpdated(rnd, updatedPolicyName, updatedPolicyDesc, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedPolicyName),
					resource.TestCheckResourceAttr(resourceName, "description", updatedPolicyDesc),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "alert_type", "universal_ssl_event_type"),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "mechanisms.email.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "mechanisms.email.0.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "mechanisms.email.0.id", "test-updated@example.com"),
					resource.TestCheckResourceAttr(resourceName, "mechanisms.email.1.id", "test2-updated@example.com"),
				),
			},
		},
	})
}

func testCheckCloudflareNotificationPolicy(name, accountID string) string {
	return acctest.LoadTestCase("checkcloudflarenotificationpolicy.tf", name, accountID)
}

func testCheckCloudflareNotificationPolicyUpdated(resName, policyName, policyDesc, accountID string) string {
	return acctest.LoadTestCase("checkcloudflarenotificationpolicyupdated.tf", resName, policyName, policyDesc, accountID)
}

func TestAccCloudflareNotificationPolicy_WithFiltersAttribute(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the IP List
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy." + rnd
	updatedPolicyName := "updated workers usage notification"
	updatedPolicyDesc := "updated description"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCheckCloudflareNotificationPolicyWithFiltersAttribute(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "workers usage notification"),
					resource.TestCheckResourceAttr(resourceName, "description", "test description"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "alert_type", "billing_usage_alert"),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.product.*", "worker_requests"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.limit.*", "100"),
				),
			},
			{
				Config: testCheckCloudflareNotificationPolicyWithFiltersAttributeUpdated(rnd, updatedPolicyName, updatedPolicyDesc, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedPolicyName),
					resource.TestCheckResourceAttr(resourceName, "description", updatedPolicyDesc),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "alert_type", "billing_usage_alert"),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.product.*", "worker_requests"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.limit.*", "100"),
				),
			},
		},
	})
}

func TestAccCloudflareNotificationPolicy_WithSelectorsAttribute(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCheckCloudflareNotificationPolicyWithSelectors(rnd, accountID, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "traffic anomalies alert"),
					resource.TestCheckResourceAttr(resourceName, "description", "test description"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "alert_type", "traffic_anomalies_alert"),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.selectors.*", "total"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.alert_trigger_preferences.*", "zscore_drop"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.group_by.*", "zone"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.selectors.*", "total"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.where.*", "(origin_status_code eq 200)"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.zones.*", zoneID),
				),
			},
		},
	})
}

func testCheckCloudflareNotificationPolicyWithFiltersAttribute(name, accountID string) string {
	return acctest.LoadTestCase("checkcloudflarenotificationpolicywithfiltersattribute.tf", name, accountID)
}

func testCheckCloudflareNotificationPolicyWithFiltersAttributeUpdated(name, policyName, policyDesc, accountID string) string {
	return acctest.LoadTestCase("checkcloudflarenotificationpolicywithfiltersattributeupdated.tf", name, policyName, policyDesc, accountID)
}

// func TestFlattenExpandFilters(t *testing.T) {
// 	filters := map[string][]string{
// 		"services": {"waf", "firewallrules"},
// 		"zones":    {"abc123"},
// 	}
// 	flattenedFilters := flattenNotificationPolicyFilter(filters)
// 	expandedFilters := expandNotificationPolicyFilter(flattenedFilters)
// 	for k := range filters {
// 		sort.Strings(filters[k])
// 		sort.Strings(expandedFilters[k])
// 		assert.EqualValuesf(t, filters[k], expandedFilters[k], "values should equal without order")
// 	}
// }

func testCheckCloudflareNotificationPolicyWithSelectors(name, accountID, zoneID string) string {
	return acctest.LoadTestCase("checkcloudflarenotificationpolicywithselectors.tf", name, accountID, zoneID)
}

func TestAccCloudflareNotificationPolicy_RemappingAffectedComponents(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_notification_policy." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCheckCloudflareNotificationPolicyWithComponents(rnd, accountID, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "traffic anomalies alert"),
					resource.TestCheckResourceAttr(resourceName, "description", "test description"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "alert_type", "incident_alert"),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "filters.affected_components.0", "API"),
				),
			},
		},
	})
}

func testCheckCloudflareNotificationPolicyWithComponents(name, accountID, zoneID string) string {
	return acctest.LoadTestCase("checkcloudflarenotificationpolicywithcomponents.tf", name, accountID, zoneID)
}

func TestAccUpgradeNotificationPolicy_FromPublishedV5(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	config := testCheckCloudflareNotificationPolicy(rnd, accountID)

	// Stepping-stone upgrade: 5.16.0 → 5.18.0 → latest.
	//
	// v5.16.0 stored state at schema_version=0 and included the alert_interval attribute
	// which is absent from the v4 source schema used by the latest provider's UpgradeState[0]
	// handler. v5.18.0 introduced a no-op UpgradeState[0] that bumps the version to 1 using
	// the current v5 schema, making the state compatible with the latest provider's
	// UpgradeState[1] handler. Jumping directly from 5.16.0 to latest would fail because
	// the latest UpgradeState[0] handler's v4 source schema does not include alert_interval.
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Step 1: create with v5.16.0 (schema_version=0, includes alert_interval)
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.16.0",
					},
				},
				Config: config,
			},
			// Step 2: upgrade to v5.18.0 (bumps schema_version 0→1, no-op transform)
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.18.0",
					},
				},
				Config: config,
			},
			// Step 3: upgrade to latest (schema_version 1→500, no-op transform)
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}
