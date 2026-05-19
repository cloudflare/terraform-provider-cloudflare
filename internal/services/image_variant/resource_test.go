package image_variant_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccCloudflareImageVariant_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image_variant." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccImageVariantBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "options.fit", "scale-down"),
					resource.TestCheckResourceAttr(resourceName, "options.height", "480"),
					resource.TestCheckResourceAttr(resourceName, "options.width", "640"),
					resource.TestCheckResourceAttr(resourceName, "options.metadata", "none"),
					resource.TestCheckResourceAttr(resourceName, "never_require_signed_urls", "false"),
				),
			},
			},
	})
}

func testAccImageVariantBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID)
}
