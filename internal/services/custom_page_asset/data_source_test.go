package custom_page_asset_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareCustomPageAssetDataSource_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := "data.cloudflare_custom_page_asset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPageAssetDataSourceConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("description"), knownvalue.StringExact("Data source test asset")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://example.com/error.html")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("id"), knownvalue.StringExact(rnd)),
				},
			},
		},
	})
}

func testAccCustomPageAssetDataSourceConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("datasource_basic.tf", rnd, accountID)
}
