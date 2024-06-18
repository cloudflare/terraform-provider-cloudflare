package regional_hostname_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
)

var (
	zoneID = os.Getenv("CLOUDFLARE_ZONE_ID")
)

func TestAccCloudflareRegionalHostname_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_regional_hostname." + rnd
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
