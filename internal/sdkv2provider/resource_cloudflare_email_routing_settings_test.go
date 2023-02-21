package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testEmailRoutingSettingsConfig(resourceID, zoneID string, enabled bool) string {
	return fmt.Sprintf(`
		resource "cloudflare_email_routing_settings" "%[1]s" {
		  zone_id = "%[2]s"
		  enabled = "%[3]t"
		}
		`, resourceID, zoneID, enabled)
}

func TestAccTestEmailRoutingSettings(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_email_routing_settings." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingSettingsConfig(rnd, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}
