package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testEmailRoutingRuleConfig(resourceID, zoneID string, enabled bool, priority int) string {
	return fmt.Sprintf(`
		resource "cloudflare_email_routing_rule" "%[1]s" {
		  zone_id = "%[2]s"
		  enabled = "%[3]t"
          priority = "%[4]d"
		  name = "terraform rule"
		  matcher {
			field  = "to"
			type = "literal"
			value = "test@example.com"
		  }

		  action {
			type = "forward"
			value = ["destinationaddress@example.net"]
		  }
	}
		`, resourceID, zoneID, enabled, priority)
}

func TestAccTestEmailRoutingRule(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_email_routing_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingRuleConfig(rnd, zoneID, true, 10),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "priority", "10"),
					resource.TestCheckResourceAttr(name, "name", "terraform rule"),

					resource.TestCheckResourceAttr(name, "matcher.0.type", "literal"),
					resource.TestCheckResourceAttr(name, "matcher.0.field", "to"),
					resource.TestCheckResourceAttr(name, "matcher.0.value", "test@example.com"),

					resource.TestCheckResourceAttr(name, "action.0.type", "forward"),
					resource.TestCheckResourceAttr(name, "action.0.value.#", "1"),
					resource.TestCheckResourceAttr(name, "action.0.value.0", "destinationaddress@example.net"),
				),
			},
		},
	})
}
