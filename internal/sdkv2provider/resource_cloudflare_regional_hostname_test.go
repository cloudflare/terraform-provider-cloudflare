package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareRegionalHostname_Basic(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_regional_hostname." + rnd
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testRegionalHostnameConfig(rnd, zoneName, "ca"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hostname", zoneName),
					resource.TestCheckResourceAttr(name, "region_key", "ca"),
				),
			},
			{
				Config: testRegionalHostnameConfig(rnd, zoneName, "eu"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hostname", zoneName),
					resource.TestCheckResourceAttr(name, "region_key", "eu"),
				),
			},
		},
	})
}

func testRegionalHostnameConfig(name string, zoneName, regionKey string) string {
	return fmt.Sprintf(`
resource "cloudflare_regional_hostname" "%[1]s" {
	zone_id = "%[2]s"
	hostname = "%[3]s"
	region_key = "%[4]s"
}`, name, zoneID, zoneName, regionKey)
}
