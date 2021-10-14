package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareAccessGroupsAccountLevel(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_access_groups.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupsConfig(rnd, "account_id", accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCloudflareAccessGroupsDataSourceId(name),
					resource.TestCheckResourceAttr(name, "groups.#", "2"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroupsZoneLevel(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_access_groups.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupsConfig(rnd, "zone_id", zoneID),
				Check: resource.ComposeTestCheckFunc(
					testAccCloudflareAccessGroupsDataSourceId(name),
					resource.TestCheckResourceAttr(name, "groups.#", "2"),
				),
			},
		},
	})
}

func testAccCloudflareAccessGroupsDataSourceId(n string) resource.TestCheckFunc {
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

func testAccCloudflareAccessGroupsConfig(name string, accessor string, accessorID string) string {
	return fmt.Sprintf(`data "cloudflare_access_groups" "%[1]s" {
		%[2]s = "%[3]s"
	}`, name, accessor, accessorID)
}

