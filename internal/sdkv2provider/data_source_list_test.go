package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareListDataSource(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()

	name := "data.cloudflare_list." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareListCreate(accountID, rnd),
			},
			{
				Config: testAccCheckCloudflareListDataSource(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "kind", "ip"),
					resource.TestCheckResourceAttr(name, "numitems", "0"),
				),
			},
		},
	})
}

func testAccCheckCloudflareListCreate(accountID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  kind        = "ip"
}
`, name, accountID)
}

func testAccCheckCloudflareListDataSource(accountID, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  kind        = "ip"
}

data "cloudflare_list" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}`, name, accountID)
}
