package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareZoneCacheReserve_Import(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_cache_reserve.%s", rnd)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudflareZoneCacheReserveUpdate(t, zoneID, true)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneCacheReserveConfig(zoneID, rnd, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, consts.ZoneIDSchemaKey),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
				),
			},
			{
				ImportState:       true,
				ImportStateId:     zoneID, // Ensure that a zone ID, not resource ID, is passed.
				ImportStateVerify: true,
				ResourceName:      name,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, consts.ZoneIDSchemaKey),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
		},
		CheckDestroy: testAccCheckCloudflareZoneCacheReserveDestroy(zoneID),
	})
}
