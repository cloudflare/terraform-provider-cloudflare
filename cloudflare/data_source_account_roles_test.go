package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareAccountRoles(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_account_roles.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccountRolesConfig(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCloudflareAccountRolesDataSourceId(name),
					resource.TestCheckResourceAttr(name, "roles.#", "20"),
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
