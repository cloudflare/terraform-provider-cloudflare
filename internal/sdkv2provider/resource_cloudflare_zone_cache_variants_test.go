package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_zone_cache_variants", &resource.Sweeper{
		Name: "cloudflare_zone_cache_variants",
		F:    testSweepCloudflareZoneCacheVariants,
	})
}

func testSweepCloudflareZoneCacheVariants(r string) error {
	ctx := context.Background()
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	tflog.Info(ctx, fmt.Sprintf("Deleting Zone Cache Variants for zone: %q", zoneID))
	client.DeleteZoneCacheVariants(context.Background(), zoneID)

	return nil
}

func TestAccCloudflareZoneCacheVariants_OneExt(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_cache_variants.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneCacheVariants_OneExt(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", zoneID),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "avif.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "avif.*", "image/avif"),
					resource.TestCheckTypeSetElemAttr(name, "avif.*", "image/webp"),
					resource.TestCheckNoResourceAttr(name, "bmp.#"),
					resource.TestCheckNoResourceAttr(name, "gif.#"),
					resource.TestCheckNoResourceAttr(name, "jpeg.#"),
					resource.TestCheckNoResourceAttr(name, "jpg.#"),
					resource.TestCheckNoResourceAttr(name, "jp2.#"),
					resource.TestCheckNoResourceAttr(name, "jpg2.#"),
					resource.TestCheckNoResourceAttr(name, "png.#"),
					resource.TestCheckNoResourceAttr(name, "tif.#"),
					resource.TestCheckNoResourceAttr(name, "tiff.#"),
					resource.TestCheckNoResourceAttr(name, "webp.#"),
				),
			},
		},
	})
}

func TestAccCloudflareZoneCacheVariants_AllExt(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_cache_variants.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneCacheVariants_AllExt(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", zoneID),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "avif.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "avif.*", "image/avif"),
					resource.TestCheckTypeSetElemAttr(name, "avif.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "bmp.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "bmp.*", "image/bmp"),
					resource.TestCheckTypeSetElemAttr(name, "bmp.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "gif.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "gif.*", "image/gif"),
					resource.TestCheckTypeSetElemAttr(name, "gif.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "jpeg.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "jpeg.*", "image/jpeg"),
					resource.TestCheckTypeSetElemAttr(name, "jpeg.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "jpg.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "jpg.*", "image/jpg"),
					resource.TestCheckTypeSetElemAttr(name, "jpg.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "jp2.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "jp2.*", "image/jp2"),
					resource.TestCheckTypeSetElemAttr(name, "jp2.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "jpg2.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "jpg2.*", "image/jpg2"),
					resource.TestCheckTypeSetElemAttr(name, "jpg2.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "png.#", "1"),
					resource.TestCheckTypeSetElemAttr(name, "png.*", "image/png"),
					resource.TestCheckResourceAttr(name, "tif.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "tif.*", "image/tif"),
					resource.TestCheckTypeSetElemAttr(name, "tif.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "tiff.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "tiff.*", "image/tiff"),
					resource.TestCheckTypeSetElemAttr(name, "tiff.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "webp.#", "1"),
					resource.TestCheckTypeSetElemAttr(name, "webp.*", "image/webp"),
				),
			},
		},
	})
}

func testAccCloudflareZoneCacheVariants_OneExt(zoneID, name string) string {
	return fmt.Sprintf(`
		resource "cloudflare_zone_cache_variants" "%[2]s" {
			zone_id = "%[1]s"
			avif = ["image/avif", "image/webp"]
		}`, zoneID, name)
}

func testAccCloudflareZoneCacheVariants_AllExt(zoneID, name string) string {
	return fmt.Sprintf(`
		resource "cloudflare_zone_cache_variants" "%[2]s" {
			zone_id = "%[1]s"
			avif = ["image/avif", "image/webp"]
			bmp = ["image/bmp", "image/webp"]
			gif = ["image/gif", "image/webp"]
			jpeg = ["image/jpeg", "image/webp"]
			jpg = ["image/jpg", "image/webp"]
			jp2 = ["image/jp2", "image/webp"]
			jpg2 = ["image/jpg2", "image/webp"]
			png = ["image/png"]
			tif = ["image/tif", "image/webp"]
			tiff = ["image/tiff", "image/webp"]
			webp = ["image/webp"]
		}`, zoneID, name)
}
