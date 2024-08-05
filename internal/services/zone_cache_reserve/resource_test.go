package zone_cache_reserve_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareZoneCacheReserve_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_cache_reserve.%s", rnd)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneCacheReserveConfig(zoneID, rnd, "on"),
				Check: resource.ComposeTestCheckFunc(
					// testAccCheckCloudflareZoneCacheReserveValuesUpdated(zoneID, true),
					resource.TestCheckResourceAttrSet(name, consts.ZoneIDSchemaKey),
					resource.TestCheckResourceAttr(name, "value", "on"),
				),
			},
			{
				Config: testAccCloudflareZoneCacheReserveConfig(zoneID, rnd, "off"),
				Check: resource.ComposeTestCheckFunc(
					// testAccCheckCloudflareZoneCacheReserveValuesUpdated(zoneID, false),
					resource.TestCheckResourceAttrSet(name, consts.ZoneIDSchemaKey),
					resource.TestCheckResourceAttr(name, "value", "off"),
				),
			},
		},
		// CheckDestroy: testAccCheckCloudflareZoneCacheReserveDestroy(zoneID),
	})
}

// func testAccCheckCloudflareZoneCacheReserveValuesUpdated(zoneID string, enable bool) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
// 		if clientErr != nil {
// 			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
// 		}

// 		params := cloudflare.GetCacheReserveParams{}
// 		output, err := client.GetCacheReserve(context.Background(), cloudflare.ZoneIdentifier(zoneID), params)
// 		if err != nil {
// 			return fmt.Errorf("unable to read Cache Reserve for zone %q: %w", zoneID, err)
// 		}

// 		// Default state for the Cache Reserve for any zone.
// 		var status, value = "disabled", cacheReserveDisabled

// 		if enable {
// 			status, value = "enabled", cacheReserveEnabled
// 		}
// 		if output.Value != value {
// 			return fmt.Errorf("expected Cache Reserve to be %q for zone: %s", status, zoneID)
// 		}

// 		return nil
// 	}
// }

// func testAccCheckCloudflareZoneCacheReserveDestroy(zoneID string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
// 		if clientErr != nil {
// 			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
// 		}

// 		params := cloudflare.GetCacheReserveParams{}
// 		output, err := client.GetCacheReserve(context.Background(), cloudflare.ZoneIdentifier(zoneID), params)
// 		if err != nil {
// 			return fmt.Errorf("unable to read Cache Reserve for zone %q: %w", zoneID, err)
// 		}

// 		// Ensure that the Cache Reserve support has been correctly
// 		// disabled for a given zone, which is also the default.
// 		if output.Value == cacheReserveEnabled {
// 			return fmt.Errorf("unable to disable Cache Reserve for zone: %s", zoneID)
// 		}

// 		return nil
// 	}
// }

// func testAccCloudflareZoneCacheReserveUpdate(t *testing.T, zoneID string, enable bool) {
// 	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
// 	if clientErr != nil {
// 		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
// 	}

// 	params := cloudflare.UpdateCacheReserveParams{
// 		Value: cacheReserveDisabled,
// 	}
// 	if enable {
// 		params.Value = cacheReserveEnabled
// 	}

// 	_, err := client.UpdateCacheReserve(context.Background(), cloudflare.ZoneIdentifier(zoneID), params)
// 	if err != nil {
// 		t.Errorf("unable to set Cache Reserve for zone %q: %s", zoneID, err)
// 	}
// }

func testAccCloudflareZoneCacheReserveConfig(zoneID, name, enable string) string {
	return acctest.LoadTestCase("zonecachereserveconfig.tf", zoneID, name, enable)
}
