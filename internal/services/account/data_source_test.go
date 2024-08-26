package account_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareAccounts(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_accounts.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccountsConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCloudflareAccountsSize(name),
				),
			},
		},
	})
}

func testAccCloudflareAccountsConfig(name string) string {
	return fmt.Sprintf(`data "cloudflare_accounts" "%[1]s" { }`, name)
}

func testAccCloudflareAccountsSize(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		var (
			accountsSize int
			err          error
		)

		if accountsSize, err = strconv.Atoi(a["accounts.#"]); err != nil {
			return err
		}

		if accountsSize < 1 {
			return fmt.Errorf("accounts count seems suspicious: %d", accountsSize)
		}

		return nil
	}
}
