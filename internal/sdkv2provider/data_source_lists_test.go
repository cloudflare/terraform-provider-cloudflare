package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareListsDataSource(t *testing.T) {
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
