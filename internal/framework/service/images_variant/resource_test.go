package images_variant_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareImagesVariant_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_images_variant.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var ImagesVariant cloudflare.ImagesVariant

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_Account(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareImagesVariantConfigurationBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareImagesVariantExists(resourceName, rnd, &ImagesVariant),
					resource.TestCheckResourceAttr(resourceName, "id", resourceName),
					resource.TestCheckResourceAttr(resourceName, "never_require_signed_urls", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.fit", "scale-down"),
					resource.TestCheckResourceAttr(resourceName, "options.metadata", "none"),
					resource.TestCheckResourceAttr(resourceName, "options.width", "500"),
					resource.TestCheckResourceAttr(resourceName, "options.height", "500"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCheckCloudflareImagesVariantExists(n string, name string, imagesVariant *cloudflare.ImagesVariant) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
		imagesVariantRS := s.RootModule().Resources["cloudflare_images_variant."+name]

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Variant ID is set")
		}

		client, err := acctest.SharedClient()
		if err != nil {
			return fmt.Errorf("error establishing client: %w", err)
		}

		foundImagesVariant, err := client.GetImagesVariant(context.Background(), cloudflare.AccountIdentifier(accountID), imagesVariantRS.Primary.ID)
		if err != nil {
			return err
		}

		*imagesVariant = foundImagesVariant

		return nil
	}
}

func testAccCloudflareImagesVariantConfigurationBasic(resName string, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_images_variant" "%[1]s" {
	account_id = "%[2]s"
	id = "%[1]s"
	never_require_signed_urls = true
	options = {
		fit = "scale-down"
		metadata = "none"
		width = 500
		height = 500
	}
}
	`, resName, accountID)
}
