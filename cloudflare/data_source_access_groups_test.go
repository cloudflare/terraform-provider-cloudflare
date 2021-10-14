package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_access_groups", &resource.Sweeper{
		Name: "cloudflare_access_groups",
		F:    testSweepCloudflareAccessGroups,
	})
}

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
					resource.TestCheckResourceAttr(name, "groups.#", "1"),
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
			return fmt.Errorf("can't find Access Groups data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Snapshot Access Groups source ID not set")
		}
		return nil
	}
}

func testAccCloudflareAccessGroupsConfig(name string, accessor string, accessorID string) string {
	return fmt.Sprintf(`data "cloudflare_access_groups" "%[1]s" {
		%[2]s = "%[3]s"
	}
	
	%[4]s
	`, name, accessor, accessorID, createTestResources())
}

func createTestResources() string {
	return fmt.Sprintf(`resource "cloudflare_access_group" "test_account_group" {
		account_id     = "%[1]s"
		name           = "staging group"
	
		include {
			email = ["test_account@example.com"]
		}
	}
	
	resource "cloudflare_access_group" "test_zone_group" {
		zone_id        = "%[2]s"
		name           = "staging group"
	
		include {
			email = ["test_zone@example.com"]
		}
	}`, os.Getenv("CLOUDFLARE_ACCOUNT_ID"), os.Getenv("CLOUDFLARE_ZONE_ID"))
}
