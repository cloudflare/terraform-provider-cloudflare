package account_role_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareAccountRoles_Datasource(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	listDataName := fmt.Sprintf("data.cloudflare_account_roles.%s", rnd)
	singleDataName := fmt.Sprintf("data.cloudflare_account_role.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccountRolesBasic(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					// verify LIST api call
					statecheck.ExpectKnownValue(listDataName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),

					// verify READ api call
					statecheck.ExpectKnownValue(singleDataName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(singleDataName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(singleDataName, tfjsonpath.New("name"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(singleDataName, tfjsonpath.New("permissions"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(singleDataName, "id", listDataName, "result.0.id"),
					resource.TestCheckResourceAttrPair(singleDataName, "name", listDataName, "result.0.name"),
					resource.TestCheckResourceAttrPair(singleDataName, "description", listDataName, "result.0.description"),
				),
			},
		},
	})
}

func testAccCloudflareAccountRolesBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
		data "cloudflare_account_roles" "%[1]s" {
		  account_id = "%[2]s"
		}
		
		data "cloudflare_account_role" "%[1]s" {
		  account_id = "%[2]s"
		  role_id    = data.cloudflare_account_roles.%[1]s.result[0].id
		}`, rnd, accountID,
	)
}
