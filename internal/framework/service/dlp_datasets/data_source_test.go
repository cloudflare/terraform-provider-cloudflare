package dlp_datasets_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareDLPDatasets_DataSource(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareDlpDatasetsDataSourceConfig(accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.cloudflare_dlp_datasets.my_datasets", "account_id"),
				),
			},
		},
	})
}

func testAccCheckCloudflareDlpDatasetsDataSourceConfig(accountID string) string {
	return fmt.Sprintf(`
data "cloudflare_dlp_datasets" "my_datasets" {
    account_id = "%s"
}
`, accountID)
}
