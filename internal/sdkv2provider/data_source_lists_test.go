package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareListDataSource(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()

	name := "data.cloudflare_lists." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareListDataSource(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
				),
			},
		},
	})
}

func testAccCheckCloudflareListDataSource(accountID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  kind        = "redirect"
}

data "cloudflare_lists" "%[1]s" {
  account_id = "%[2]s"
}`, name, accountID)
}
