package cloudflare

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareAccountRoles(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_account_roles.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCloudflareAccountRolesConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCloudflareAccountRolesDataSourceId(name),
					resource.TestCheckResourceAttr(name, "roles.#", "1"),
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

func testAccCloudflareAccountRolesConfig(name string) string {
	return fmt.Sprintf(`data "cloudflare_account_roles" "%[1]s" {}`, name)
}
