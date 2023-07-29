package sdkv2provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataCloudflareZoneCacheReserve_Simple(t *testing.T) {
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

func TestAccDataCloudflareZoneCacheReserve_Error(t *testing.T) {
	rnd := generateRandomResourceName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataCloudflareZoneCacheReserveConfig("this is a test", rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("must be a valid Zone ID, got: this is a test")),
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
