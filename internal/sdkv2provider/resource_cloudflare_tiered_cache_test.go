package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testTieredCacheConfig(rnd, zoneID, cacheType string) string {
	return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
	zone_id = "%[2]s"
	cache_type = "%[3]s"
}
`, rnd, zoneID, cacheType)
}

func TestAccCloudflareTieredCache_Smart(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_tiered_cache." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testTieredCacheConfig(rnd, zoneID, "smart"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "cache_type", "smart"),
				),
			},
		},
	})
}

func TestAccCloudflareTieredCache_Generic(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_tiered_cache." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testTieredCacheConfig(rnd, zoneID, "generic"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "cache_type", "generic"),
				),
			},
		},
	})
}
