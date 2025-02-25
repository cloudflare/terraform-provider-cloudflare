package account_role_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareAccountRoles_Datasource(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_account_roles.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccountRoles_Config(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					testAccCheckExampleWidgetExists(name),
				),
			},
		},
	})
}

func testAccCheckExampleWidgetExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		i, _ := strconv.Atoi(rs.Primary.Attributes["result.#"])
		if i < 30 {
			return fmt.Errorf("role size is suspiciously low. should be > 30, got: %d", i)
		}

		return nil
	}
}

func testAccCloudflareAccountRoles_Config(rnd, accountID string) string {
	return acctest.LoadTestCase("datasource.tf", rnd, accountID)
}
