package cloudflare

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestZone(t *testing.T) {
	name := "cloudflare_zone.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig("test", "example.org", "true", "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", "example.org"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
					resource.TestCheckResourceAttr(name, "plan", "Free Website"),
				),
			},
		},
	})
}

func testZoneConfig(resourceID, zoneName, paused, jumpStart string) string {
	return fmt.Sprintf(`
				resource "cloudflare_zone" "%[1]s" {
                    zone = "%[2]s"
					paused = %[3]s
					jump_start = %[4]s
				}`, resourceID, zoneName, paused, jumpStart)
}
