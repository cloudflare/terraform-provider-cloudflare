package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/terraform"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudflareZone(t *testing.T) {
	name := os.Getenv("CLOUDFLARE_DOMAIN")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZoneDataSourceID("data.cloudflare_zone.example_dot_com"),
					resource.TestCheckResourceAttr("data.cloudflare_zone.example_dot_com", "zone", name),
					resource.TestCheckResourceAttr("data.cloudflare_zone.example_dot_com", "paused", "false"),
				),
			},
		},
	})
}

func testAccCloudflareZoneConfig(zoneName string) string {
	return fmt.Sprintf(`
data "cloudflare_zone" "example_dot_com" {
	zone = "%[1]s"
}`, zoneName)
}

func testAccCheckCloudflareZoneDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find zone data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Snapshot zone source ID not set")
		}
		return nil
	}
}
