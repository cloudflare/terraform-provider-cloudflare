package cloudflare

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestCloudflareZoneLockdown(t *testing.T) {
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := acctest.RandString(10)
	name := "cloudflare_zone_lockdown." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneLockdownConfig(rnd, zone, "false", "this is notes", rnd+"."+zone+"/*", "ip", "198.51.100.4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone", zone),
					resource.TestMatchResourceAttr(name, "zone_id", regexp.MustCompile("^[a-z0-9]{32}$")),
					resource.TestCheckResourceAttr(name, "paused", "false"),
					resource.TestCheckResourceAttr(name, "description", "this is notes"),
					resource.TestCheckResourceAttr(name, "urls.#", "1"),
					resource.TestCheckResourceAttr(name, "configurations.#", "1"),
				),
			},
		},
	})
}

func TestCloudflareZoneLockdown_Import(t *testing.T) {
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := acctest.RandString(10)
	name := "cloudflare_zone_lockdown." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWAFRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneLockdownConfig(rnd, zone, "false", "this is notes", rnd+"."+zone+"/*", "ip", "198.51.100.4"),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zone),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testCloudflareZoneLockdownConfig(resourceID, zone, paused, description, url, target, value string) string {
	return fmt.Sprintf(`
				resource "cloudflare_zone_lockdown" "%[1]s" {
					zone = "%[2]s"
					paused = "%[3]s"
					description = "%[4]s"
					urls = ["%[5]s"]
					configurations = [
						{
							target = "%[6]s"
							value = "%[7]s"
						}
					]
				}`, resourceID, zone, paused, description, url, target, value)
}
