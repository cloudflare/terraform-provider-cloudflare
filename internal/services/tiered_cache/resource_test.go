package tiered_cache_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testTieredCacheConfig(rnd, zoneID, cacheType string) string {
	return acctest.LoadTestCase("tieredcacheconfig.tf", rnd, zoneID, cacheType)
}

func TestAccCloudflareTieredCache_Smart(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_tiered_cache." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testTieredCacheConfig(rnd, zoneID, "on"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "value", "on"),
				),
			},
		},
	})
}
