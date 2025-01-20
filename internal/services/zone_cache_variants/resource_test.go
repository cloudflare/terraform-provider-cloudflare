package zone_cache_variants_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_zone_cache_variants", &resource.Sweeper{
		Name: "cloudflare_zone_cache_variants",
		F:    testSweepCloudflareZoneCacheVariants,
	})
}

func testSweepCloudflareZoneCacheVariants(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
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
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_cache_variants.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneCacheVariants_OneExt(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", zoneID),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "value.avif.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "value.avif.*", "image/avif"),
					resource.TestCheckTypeSetElemAttr(name, "value.avif.*", "image/webp"),
					resource.TestCheckNoResourceAttr(name, "value.bmp.#"),
					resource.TestCheckNoResourceAttr(name, "value.gif.#"),
					resource.TestCheckNoResourceAttr(name, "value.jpeg.#"),
					resource.TestCheckNoResourceAttr(name, "value.jpg.#"),
					resource.TestCheckNoResourceAttr(name, "value.jp2.#"),
					resource.TestCheckNoResourceAttr(name, "value.jpg2.#"),
					resource.TestCheckNoResourceAttr(name, "value.png.#"),
					resource.TestCheckNoResourceAttr(name, "value.tif.#"),
					resource.TestCheckNoResourceAttr(name, "value.tiff.#"),
					resource.TestCheckNoResourceAttr(name, "value.webp.#"),
				),
			},
		},
	})
}

func TestAccCloudflareZoneCacheVariants_AllExt(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_cache_variants.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneCacheVariants_AllExt(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", zoneID),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "value.avif.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "value.avif.*", "image/avif"),
					resource.TestCheckTypeSetElemAttr(name, "value.avif.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "value.bmp.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "value.bmp.*", "image/bmp"),
					resource.TestCheckTypeSetElemAttr(name, "value.bmp.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "value.gif.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "value.gif.*", "image/gif"),
					resource.TestCheckTypeSetElemAttr(name, "value.gif.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "value.jpeg.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "value.jpeg.*", "image/jpeg"),
					resource.TestCheckTypeSetElemAttr(name, "value.jpeg.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "value.jpg.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "value.jpg.*", "image/jpg"),
					resource.TestCheckTypeSetElemAttr(name, "value.jpg.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "value.jp2.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "value.jp2.*", "image/jp2"),
					resource.TestCheckTypeSetElemAttr(name, "value.jp2.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "value.jpg2.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "value.jpg2.*", "image/jpg2"),
					resource.TestCheckTypeSetElemAttr(name, "value.jpg2.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "value.png.#", "1"),
					resource.TestCheckTypeSetElemAttr(name, "value.png.*", "image/png"),
					resource.TestCheckResourceAttr(name, "value.tif.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "value.tif.*", "image/tif"),
					resource.TestCheckTypeSetElemAttr(name, "value.tif.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "value.tiff.#", "2"),
					resource.TestCheckTypeSetElemAttr(name, "value.tiff.*", "image/tiff"),
					resource.TestCheckTypeSetElemAttr(name, "value.tiff.*", "image/webp"),
					resource.TestCheckResourceAttr(name, "value.webp.#", "1"),
					resource.TestCheckTypeSetElemAttr(name, "value.webp.*", "image/webp"),
				),
			},
		},
	})
}

func testAccCloudflareZoneCacheVariants_OneExt(zoneID, name string) string {
	return acctest.LoadTestCase("oneext.tf", zoneID, name)
}

func testAccCloudflareZoneCacheVariants_AllExt(zoneID, name string) string {
	return acctest.LoadTestCase("allext.tf", zoneID, name)
}
