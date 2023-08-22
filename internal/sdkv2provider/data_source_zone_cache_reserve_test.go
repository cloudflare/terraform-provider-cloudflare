package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataCloudflareZoneCacheReserve_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_zone_cache_reserve.%s", rnd)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudflareZoneCacheReserveUpdate(t, zoneID, true)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataCloudflareZoneCacheReserveConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZoneCacheReserveValuesUpdated(zoneID, true),
					resource.TestCheckResourceAttrSet(name, consts.ZoneIDSchemaKey),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
		},
	})
}

func testAccDataCloudflareZoneCacheReserveConfig(zoneID, name string) string {
	return fmt.Sprintf(`
		data "cloudflare_zone_cache_reserve" "%[2]s" {
			zone_id = "%[1]s"
		}`, zoneID, name)
}
