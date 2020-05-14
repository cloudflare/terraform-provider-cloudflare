package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudflareWAFOverride(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_waf_override.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWAFOverrideBasicConfig(zoneID, zoneName, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "urls.#", "2"),
					resource.TestCheckResourceAttr(name, "urls.0", fmt.Sprintf("%s/basic-waf-override", zoneName)),
					resource.TestCheckResourceAttr(name, "urls.1", fmt.Sprintf("%s/another-basic-waf-override", zoneName)),
					resource.TestCheckResourceAttr(name, "rules.100015", "disable"),
					resource.TestCheckResourceAttr(name, "groups.ea8687e59929c1fd05ba97574ad43f77", "default"),
					resource.TestCheckResourceAttr(name, "rewrite_action.default", "block"),
					resource.TestCheckResourceAttr(name, "rewrite_action.challenge", "block"),
				),
			},
		},
	})
}

func testAccCheckCloudflareWAFOverrideBasicConfig(zoneID, zoneName, name string) string {
	return fmt.Sprintf(`
		resource "cloudflare_waf_override" "%[3]s" {
			zone_id = "%[1]s"
			urls = ["%[2]s/basic-waf-override", "%[2]s/another-basic-waf-override"]
			rules = {
				"100015": "disable"
			}
			groups = {
				"ea8687e59929c1fd05ba97574ad43f77": "default"
			}
			rewrite_action = {
				"default": "block",
  			"challenge": "block",
			}
		}`, zoneID, zoneName, name)
}
