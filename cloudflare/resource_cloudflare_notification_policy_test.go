package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareNotificationPolicy(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the notification
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceName := "cloudflare_notification_policy." + rnd
	updatedPolicyName := "updated test SSL policy from terraform provider"
	updatedPolicyDesc := "updated description"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testCheckCloudflareNotificationPolicy(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test SSL policy from terraform provider"),
					resource.TestCheckResourceAttr(resourceName, "description", "test description"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "alert_type", "universal_ssl_event_type"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
				),
			},
			{
				Config: testCheckCloudflareNotificationPolicyUpdated(rnd, updatedPolicyName, updatedPolicyDesc, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedPolicyName),
					resource.TestCheckResourceAttr(resourceName, "description", updatedPolicyDesc),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "alert_type", "universal_ssl_event_type"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
				),
			},
		},
	})
}

func testCheckCloudflareNotificationPolicy(name, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_notification_policy" "%[1]s" {
    name        = "test SSL policy from terraform provider"
    account_id  = "%[2]s"
    description = "test description"
    enabled     =  true
    alert_type  = "universal_ssl_event_type"
    email_integration {
      name =  ""
      id   =  "test@example.com"
    }
  }`, name, accountID)
}

func testCheckCloudflareNotificationPolicyUpdated(resName, policyName, policyDesc, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_notification_policy" "%[1]s" {
    name        = "%[2]s"
    account_id  = "%[4]s"
    description = "%[3]s"
    enabled     =  true
    alert_type  = "universal_ssl_event_type"
    email_integration {
      name =  ""
      id   =  "test@example.com"
    }
  }`, resName, policyName, policyDesc, accountID)
}

func TestAccCloudflareNotificationPolicyWithFiltersAttribute(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the notification
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceName := "cloudflare_notification_policy." + rnd
	updatedPolicyName := "updated Advanced Security Events Alert from terraform provider"
	updatedPolicyDesc := "updated description"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testCheckCloudflareNotificationPolicyWithFiltersAttribute(rnd, accountID, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test Advanced Security Events Alert from terraform provider"),
					resource.TestCheckResourceAttr(resourceName, "description", "test description"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "alert_type", "clickhouse_alert_fw_ent_anomaly"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "filters.0.services.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.services.0", "waf"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.zones.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.zones.0", zoneID),
				),
			},
			{
				Config: testCheckCloudflareNotificationPolicyWithFiltersAttributeUpdated(rnd, updatedPolicyName, updatedPolicyDesc, accountID, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedPolicyName),
					resource.TestCheckResourceAttr(resourceName, "description", updatedPolicyDesc),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "alert_type", "clickhouse_alert_fw_ent_anomaly"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "filters.0.services.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.services.0", "waf"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.services.1", "firewallrules"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.zones.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.zones.0", zoneID),
				),
			},
		},
	})
}

func testCheckCloudflareNotificationPolicyWithFiltersAttribute(name, accountID, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_notification_policy" "%[1]s" {
    name        = "test Advanced Security Events Alert from terraform provider"
    account_id  = "%[2]s"
    description = "test description"
    enabled     =  true
    alert_type  = "clickhouse_alert_fw_ent_anomaly"
    email_integration {
      name =  ""
      id   =  "test@example.com"
    }
    filters {
      services = [
        "waf",
      ]
      zones = ["%[3]s"]
    }
  }`, name, accountID, zoneID)
}

func testCheckCloudflareNotificationPolicyWithFiltersAttributeUpdated(resName, policyName, policyDesc, accountID, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_notification_policy" "%[1]s" {
    name        = "%[2]s"
    account_id  = "%[4]s"
    description = "%[3]s"
    enabled     =  true
    alert_type  = "clickhouse_alert_fw_ent_anomaly"
    email_integration {
      name =  ""
      id   =  "test@example.com"
    }
    filters {
      services = [
        "waf",
        "firewallrules",
      ]
      zones = ["%[5]s"]
    }
  }`, resName, policyName, policyDesc, accountID, zoneID)
}
