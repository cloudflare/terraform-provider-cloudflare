package sdkv2provider

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareAccountRoles(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_account_roles.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccountRolesConfig(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCloudflareAccountRolesDataSourceId(name),
					testAccCloudflareAccountRolesSize(name),
				),
			},
		},
	})
}

func testAccCloudflareAccountRolesDataSourceId(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		all := s.RootModule().Resources
		rs, ok := all[n]

		if !ok {
			return fmt.Errorf("can't find Account Roles data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Snapshot Account Roles source ID not set")
		}
		return nil
	}
}

func testAccCloudflareAccountRolesConfig(name string, accountID string) string {
	return fmt.Sprintf(`data "cloudflare_account_roles" "%[1]s" {
		account_id = "%[2]s"
	}`, name, accountID)
}

func testAccCloudflareAccountRolesSize(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		var (
			rolesSize int
			err       error
		)

		if rolesSize, err = strconv.Atoi(a["roles.#"]); err != nil {
			return err
		}

		if rolesSize < 20 {
			return fmt.Errorf("role count seems suspiciously low: %d", rolesSize)
		}

		return nil
	}
}
