package gateway_app_types_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareGatewayAppTypes_DataSource(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareGatewayAppTypesDataSourceConfig(accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.cloudflare_gateway_app_types.my_app_types", consts.AccountIDSchemaKey),
					resource.TestCheckResourceAttrSet("data.cloudflare_gateway_app_types.my_app_types", "app_types.#"),
				),
			},
		},
	})
}

func testAccCheckCloudflareGatewayAppTypesDataSourceConfig(accountID string) string {
	return fmt.Sprintf(`
data "cloudflare_gateway_app_types" "my_app_types" {
    account_id = "%s"
}
`, accountID)
}
