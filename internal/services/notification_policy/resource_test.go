package notification_policy_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

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
					resource.TestCheckResourceAttr(resourceName, "email_integration.#", "2"),
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
					resource.TestCheckResourceAttr(resourceName, "email_integration.#", "2"),
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
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.0.product.*", "worker_requests"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.0.limit.*", "100"),
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
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.0.product.*", "worker_requests"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.0.limit.*", "100"),
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
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.0.selectors.*", "total"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.0.alert_trigger_preferences.*", "zscore_drop"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.0.group_by.*", "zone"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.0.selectors.*", "total"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.0.where.*", "(origin_status_code eq 200)"),
					resource.TestCheckTypeSetElemAttr(resourceName, "filters.0.zones.*", zoneID),
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
					resource.TestCheckResourceAttr(resourceName, "filters.0.affected_components.0", "API"),
				),
			},
		},
	})
}

func testCheckCloudflareNotificationPolicyWithComponents(name, accountID, zoneID string) string {
	return acctest.LoadTestCase("checkcloudflarenotificationpolicywithcomponents.tf", name, accountID, zoneID)
}
