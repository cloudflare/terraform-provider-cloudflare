package regional_hostname_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
	return acctest.LoadTestCase("regionalhostnameconfig.tf", name, zoneID, zoneName, regionKey)
}
