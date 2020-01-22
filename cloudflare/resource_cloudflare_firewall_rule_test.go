package cloudflare

import (
	"fmt"
	"log"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_firewall_rule", &resource.Sweeper{
		Name: "cloudflare_firewall_rule",
		F:    testSweepCloudflareFirewallRuleSweeper,
	})

	// Defined in the filter test but referenced here as the firewall uses
	// filters for the expressions and need to be cleaned out with the
	// firewall_rule resources.
	resource.AddTestSweepers("cloudflare_filter", &resource.Sweeper{
		Name: "cloudflare_filter",
		F:    testSweepCloudflareFilterSweeper,
	})
}

func testSweepCloudflareFirewallRuleSweeper(r string) error {
	client, clientErr := sharedClient()
	if clientErr != nil {
		log.Printf("[ERROR] Failed to create Cloudflare client: %s", clientErr)
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rules, rulesErr := client.FirewallRules(zoneID, cloudflare.PaginationOptions{})

	if rulesErr != nil {
		log.Printf("[ERROR] Failed to fetch Cloudflare firewall rules: %s", rulesErr)
	}

	for _, rule := range rules {
		err := client.DeleteFirewallRule(zoneID, rule.ID)

		if err != nil {
			log.Printf("[ERROR] Failed to delete Cloudflare firewall rule (%s) in zone ID: %s", rule.ID, zoneID)
		}
	}

	return nil
}

func TestAccFirewallRuleSimple(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_firewall_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	filterQuoted := `(http.request.uri.path ~ \".*wp-login-` + rnd + `.php\" or http.request.uri.path ~ \".*xmlrpc.php\") and ip.src ne 192.0.2.1`

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testFirewallRuleConfig(rnd, zoneID, "true", "this is notes", filterQuoted, "allow", "1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "description", "this is notes"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "action", "allow"),
					resource.TestCheckResourceAttr(name, "priority", "1"),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
				),
			},
		},
	})
}

func testFirewallRuleConfig(resourceID, zoneID, paused, description, expression, action, priority string) string {
	return fmt.Sprintf(`
		resource "cloudflare_filter" "%[1]s" {
		  zone_id = "%[2]s"
		  paused = "%[3]s"
		  description = "%[4]s"
		  expression = "%[5]s"
		}
		resource "cloudflare_firewall_rule" "%[1]s" {
		  zone_id = "%[2]s"
		  paused = "%[3]s"
		  description = "%[4]s"
		  filter_id = "${cloudflare_filter.%[1]s.id}"
		  action = "%[6]s"
		  priority = %[7]s
		}
		`, resourceID, zoneID, paused, description, expression, action, priority)
}
