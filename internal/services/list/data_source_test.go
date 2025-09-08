package list_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareListWithItems_IP_Datasource(t *testing.T) {

	rndIP := utils.GenerateRandomResourceName()

	datasourceName := fmt.Sprintf("data.cloudflare_list.%s", rndIP)

	descriptionIP := fmt.Sprintf("description.%s", rndIP)

	listNameIP := fmt.Sprintf("%s%s", listTestPrefix, rndIP)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareListWithIPItemsDataSource(rndIP, listNameIP, descriptionIP, accountID, "ip", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "name", listNameIP),
					resource.TestCheckResourceAttr(datasourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(datasourceName, "description", descriptionIP),
					resource.TestCheckResourceAttr(datasourceName, "kind", "ip"),
					resource.TestCheckTypeSetElemNestedAttrs(datasourceName, "items.*", map[string]string{
						"ip": "1.1.1.1",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(datasourceName, "items.*", map[string]string{
						"ip": "1.1.1.2",
					}),
					resource.TestCheckResourceAttr(datasourceName, "items.#", "2"),
				),
			},

			{
				Config: testAccCheckCloudflareListWithIPItemsDataSource(rndIP, listNameIP, descriptionIP, accountID, "ip", "1.1.1.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "name", listNameIP),
					resource.TestCheckResourceAttr(datasourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(datasourceName, "description", descriptionIP),
					resource.TestCheckResourceAttr(datasourceName, "kind", "ip"),
					resource.TestCheckTypeSetElemNestedAttrs(datasourceName, "items.*", map[string]string{
						"ip": "1.1.1.1",
					}),
					resource.TestCheckResourceAttr(datasourceName, "items.#", "1"),
				),
			},
		},
	})
}

func testAccCheckCloudflareListWithIPItemsDataSource(resourceName, listName, description, accountID, kind, search string) string {
	return acctest.LoadTestCase("listwithipitems_datasource.tf", resourceName, listName, description, accountID, kind, search)
}
