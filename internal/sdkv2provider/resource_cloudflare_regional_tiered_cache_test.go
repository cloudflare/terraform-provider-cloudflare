package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareRegionalTieredCache_Basic(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_regional_tiered_cache.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRegionalTieredCache(rnd, zoneID, "on"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "value", "on"),
				),
			},
			{
				Config: testAccCloudflareRegionalTieredCache(rnd, zoneID, "off"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "value", "off"),
				),
			},
			{
				Config: testAccCloudflareRegionalTieredCache(rnd, zoneID, "on"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "value", "on"),
				),
			},
			{
				Config:        testAccCloudflareRegionalTieredCache(rnd, zoneID, "on"),
				ResourceName:  name,
				ImportState:   true,
				ImportStateId: zoneID,
			},
		},
	})
}

func testAccCloudflareRegionalTieredCache(resourceName, zoneID, value string) string {
	return fmt.Sprintf(`
resource "cloudflare_regional_tiered_cache" "%[1]s" {
  zone_id = "%[2]s"
  value   = "%[3]s"
}
`, resourceName, zoneID, value)
}
