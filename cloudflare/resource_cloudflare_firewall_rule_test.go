package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccFirewallRuleSimple(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_firewall_rule." + rnd
	zone := os.Getenv("CLOUDFLARE_DOMAIN")

	filterQuoted := `(http.request.uri.path ~ \".*wp-login-` + rnd + `.php\" or http.request.uri.path ~ \".*xmlrpc.php\") and ip.src ne 192.0.2.1`
	// filterUnquoted := `(http.request.uri.path ~ ".*wp-login-` + rnd + `.php" or http.request.uri.path ~ ".*xmlrpc.php") and ip.src ne 192.0.2.1`

	fmt.Print(testFirewallRuleConfig(rnd, zone, "true", "this is notes", filterQuoted, "allow", "1"))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testFirewallRuleConfig(rnd, zone, "true", "this is notes", filterQuoted, "allow", "1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "description", "this is notes"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "action", "allow"),
					resource.TestCheckResourceAttr(name, "priority", "1"),
					resource.TestCheckResourceAttr(name, "zone", zone),
					resource.TestCheckResourceAttrSet(name, "zone_id"),
				),
			},
		},
	})
}

func testFirewallRuleConfig(resourceID, zoneName, paused, description, expression, action, priority string) string {
	return fmt.Sprintf(`
		resource "cloudflare_filter" "%[1]s" {
		  zone = "%[2]s"
		  paused = "%[3]s"
		  description = "%[4]s"
		  expression = "%[5]s"
		}
		resource "cloudflare_firewall_rule" "%[1]s" {
		  zone = "%[2]s"
		  paused = "%[3]s"
		  description = "%[4]s"
		  filter_id = "${cloudflare_filter.%[1]s.id}"
		  action = "%[6]s"
		  priority = %[7]s
		}
		`, resourceID, zoneName, paused, description, expression, action, priority)
}
