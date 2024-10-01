package gateway_categories_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareGatewayCategories_DataSource(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for acceptance tests")
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareGatewayCategoriesDataSourceConfig(accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.cloudflare_gateway_categories.my_categories", consts.AccountIDSchemaKey),
					resource.TestCheckResourceAttrSet("data.cloudflare_gateway_categories.my_categories", "categories.#"),
				),
			},
		},
	})
}

func testAccCheckCloudflareGatewayCategoriesDataSourceConfig(accountID string) string {
	return fmt.Sprintf(`
data "cloudflare_gateway_categories" "my_categories" {
    account_id = "%s"
}
`, accountID)
}
