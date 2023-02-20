package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testEmailRoutingRuleCatchAllConfig(resourceID, zoneID string, enabled bool) string {
	return fmt.Sprintf(`
		resource "cloudflare_email_routing_catch_all" "%[1]s" {
		  zone_id = "%[2]s"
		  enabled = "%[3]t"
		  name = "terraform rule catch all"

		  matcher {
			type  = "all"
		  }

		  action {
			type = "forward"
			value = ["destinationaddress@example.net"]
		  }
	}
		`, resourceID, zoneID, enabled)
}

func TestAccTestEmailRoutingCatchAll(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_email_routing_catch_all." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingRuleCatchAllConfig(rnd, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", "terraform rule catch all"),

					resource.TestCheckResourceAttr(name, "matcher.0.type", "all"),

					resource.TestCheckResourceAttr(name, "action.0.type", "forward"),
					resource.TestCheckResourceAttr(name, "action.0.value.#", "1"),
					resource.TestCheckResourceAttr(name, "action.0.value.0", "destinationaddress@example.net"),
				),
			},
		},
	})
}
