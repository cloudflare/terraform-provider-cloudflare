package image_variant_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareImageVariantDataSource_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	datasourceName := "data.cloudflare_image_variant." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccImageVariantDataSourceConfig(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(datasourceName, "variant_id", rnd),
					resource.TestCheckResourceAttr(datasourceName, "variant.options.fit", "scale-down"),
					resource.TestCheckResourceAttr(datasourceName, "variant.options.height", "480"),
					resource.TestCheckResourceAttr(datasourceName, "variant.options.width", "640"),
					resource.TestCheckResourceAttr(datasourceName, "variant.options.metadata", "none"),
				),
			},
		},
	})
}

func testAccImageVariantDataSourceConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("datasource_basic.tf", rnd, accountID)
}
