package cloudflare

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/terraform"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudflareZonesMatchName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZonesConfigMatchName(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZonesDataSourceID("data.cloudflare_zones.examples_domains"),
					resource.TestCheckResourceAttr("data.cloudflare_zones.examples_domains", "zones.#", "2"),
				),
			},
		},
	})
}

func TestAccCloudflareZonesMatchPaused(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZonesConfigMatchPaused(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZonesDataSourceID("data.cloudflare_zones.examples_domains"),
					resource.TestCheckResourceAttr("data.cloudflare_zones.examples_domains", "zones.#", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareZonesMatchStatus(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZonesConfigMatchStatus(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZonesDataSourceID("data.cloudflare_zones.examples_domains"),
					testAccCheckCloudflareZonesReturned("data.cloudflare_zones.examples_domains", "zones.#", func(i int) bool {
						return i >= 1 && i <= 2
					}),
				),
			},
		},
	})
}

func testAccCheckCloudflareZonesDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		all := s.RootModule().Resources
		rs, ok := all[n]
		if !ok {
			return fmt.Errorf("Can't find zones data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Snapshot zones source ID not set")
		}
		return nil
	}
}

func testAccCheckCloudflareZonesReturned(n string, a string, check func(int) bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		all := s.RootModule().Resources
		rs, ok := all[n]
		if !ok {
			return fmt.Errorf("Can't find zones data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Snapshot zones source ID not set")
		}

		count, _ := strconv.Atoi(rs.Primary.Attributes[a])
		if !check(count) {
			return fmt.Errorf("Error evaluating %s.%s actual count: %d", n, a, count)
		}
		return nil
	}
}

func testAccCloudflareZonesConfigMatchName() string {
	return fmt.Sprintf(`
data "cloudflare_zones" "examples_domains" {
  filter {
    zone   = "baa.*"
    paused = "${cloudflare_zone.foo_net.paused}" // true
  }
}

%s
`, testZones)
}

func testAccCloudflareZonesConfigMatchPaused() string {
	return fmt.Sprintf(`
data "cloudflare_zones" "examples_domains" {
  filter {
    zone   = "baa.*"
    paused = "${cloudflare_zone.baa_com.paused}" // false
  }
}

%s
`, testZones)
}

func testAccCloudflareZonesConfigMatchStatus() string {
	return fmt.Sprintf(`
data "cloudflare_zones" "examples_domains" {
  filter {
    status = "active"
    paused = "${cloudflare_zone.baa_com.paused}" // false
  }
}

%s
`, testZones)
}

const testZones = `resource "cloudflare_zone" "baa_com" {
  zone       = "baa.com"
  paused     = false
  jump_start = false
}

resource "cloudflare_zone" "baa_org" {
  zone       = "baa.org"
  paused     = true
  jump_start = false
}

resource "cloudflare_zone" "baa_net" {
  zone       = "baa.net"
  paused     = true
  jump_start = false
}

resource "cloudflare_zone" "foo_net" {
  zone       = "foo.net"
  paused     = true
  jump_start = false
  depends_on = ["cloudflare_zone.baa_net", "cloudflare_zone.baa_org", "cloudflare_zone.baa_com"]
}`
