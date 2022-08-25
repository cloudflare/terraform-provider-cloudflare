package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func testEmailRoutingRuleConfig(resourceID, zoneID string, enabled bool, priority int) string {
	return fmt.Sprintf(`
		resource "cloudflare_email_routing_rule" "%[1]s" {
		  zone_id = "%[2]s"
		  enabled = "%[3]t"
          priority = "%[4]d"
		matchers = [
			{
				type = "literal"
				field = "to"
				value = "test@example.com"
			}
		]
		actions = [
			{
				type = "forward"
				value = ["destinationaddress@example.net"]
			}
		]
		}
		`, resourceID, zoneID, enabled, priority)
}

func TestAccTestEmailRoutingRule(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_email_routing_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	//resourceCloudflareEmailRoutingRule
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingRuleConfig(rnd, zoneID, true, 10),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enable", "true"),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "priority", "10"),
					resource.TestCheckResourceAttr(name, "matchers.0.type", "literal"),
					resource.TestCheckResourceAttr(name, "matchers.0.field", "to"),
					resource.TestCheckResourceAttr(name, "matchers.0.value", "test@example.com"),
					resource.TestCheckResourceAttr(name, "actions.0.type", "forward"),
					resource.TestCheckResourceAttr(name, "actions.0.value.0", "destinationaddress@example.net"),
				),
			},
		},
	})
}
