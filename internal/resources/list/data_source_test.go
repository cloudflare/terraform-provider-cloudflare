package list_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/consts"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
)

func TestAccCloudflareListsDataSource(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()

	name := "data.cloudflare_lists." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareListsDataSource(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "lists.0.numitems", "4"),
				),
			},
		},
	})
}

func testAccCheckCloudflareListsDataSource(accountID, name string) string {
	return fmt.Sprintf(`
	resource "cloudflare_list" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		description = "example list"
		kind = "ip"

		item {
		  value {
			ip = "192.0.2.0"
		  }
		  comment = "one"
		}

		item {
		  value {
			ip = "192.0.2.1"
		  }
		  comment = "two"
		}

		item {
		  value {
			ip = "192.0.2.2"
		  }
		  comment = "three"
		}

		item {
		  value {
			ip = "192.0.2.3"
		  }
		  comment = "four"
		}
	  }

data "cloudflare_lists" "%[1]s" {
  account_id = "%[2]s"
  depends_on = [ cloudflare_list.%[1]s ]
}`, name, accountID)
}
