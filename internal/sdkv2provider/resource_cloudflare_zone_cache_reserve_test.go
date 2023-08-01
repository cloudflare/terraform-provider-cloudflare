package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareZoneCacheReserve_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_cache_reserve.%s", rnd)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneCacheReserveConfig(zoneID, rnd, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZoneCacheReserveValuesUpdated(zoneID, true),
					resource.TestCheckResourceAttrSet(name, consts.ZoneIDSchemaKey),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
			{
				Config: testAccCloudflareZoneCacheReserveConfig(zoneID, rnd, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZoneCacheReserveValuesUpdated(zoneID, false),
					resource.TestCheckResourceAttrSet(name, consts.ZoneIDSchemaKey),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
				),
			},
		},
		CheckDestroy: testAccCheckCloudflareZoneCacheReserveDestroy(zoneID),
	})
}

func testAccCheckCloudflareZoneCacheReserveValuesUpdated(zoneID string, enable bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)

		params := cloudflare.GetCacheReserveParams{}
		output, err := client.GetCacheReserve(context.Background(), cloudflare.ZoneIdentifier(zoneID), params)
		if err != nil {
			return fmt.Errorf("unable to read Cache Reserve for zone %q: %w", zoneID, err)
		}

		// Default state for the Cache Reserve for any zone.
		var status, value = "disabled", cacheReserveDisabled

		if enable {
			status, value = "enabled", cacheReserveEnabled
		}
		if output.Value != value {
			return fmt.Errorf("expected Cache Reserve to be %q for zone: %s", status, zoneID)
		}

		return nil
	}
}

func testAccCheckCloudflareZoneCacheReserveDestroy(zoneID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)

		params := cloudflare.GetCacheReserveParams{}
		output, err := client.GetCacheReserve(context.Background(), cloudflare.ZoneIdentifier(zoneID), params)
		if err != nil {
			return fmt.Errorf("unable to read Cache Reserve for zone %q: %w", zoneID, err)
		}

		// Ensure that the Cache Reserve support has been correctly
		// disabled for a given zone, which is also the default.
		if output.Value == cacheReserveEnabled {
			return fmt.Errorf("unable to disable Cache Reserve for zone: %s", zoneID)
		}

		return nil
	}
}

func testAccCloudflareZoneCacheReserveUpdate(t *testing.T, zoneID string, enable bool) {
	client := testAccProvider.Meta().(*cloudflare.API)

	params := cloudflare.UpdateCacheReserveParams{
		Value: cacheReserveDisabled,
	}
	if enable {
		params.Value = cacheReserveEnabled
	}

	_, err := client.UpdateCacheReserve(context.Background(), cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		t.Errorf("unable to set Cache Reserve for zone %q: %s", zoneID, err)
	}
}

func testAccCloudflareZoneCacheReserveConfig(zoneID, name string, enable bool) string {
	return fmt.Sprintf(`
		resource "cloudflare_zone_cache_reserve" "%[2]s" {
			zone_id = "%[1]s"
			enabled = "%[3]t"
		}`, zoneID, name, enable)
}
