package cloudflare

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/terraform"

	"github.com/hashicorp/terraform/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_zones", &resource.Sweeper{
		Name: "cloudflare_zones",
		F:    testSweepCloudflareZones,
	})
}

func testSweepCloudflareZones(r string) error {
	client, clientErr := sharedClient()
	if clientErr != nil {
		log.Fatalf("[ERROR] Failed to create Cloudflare client: %s", clientErr)
	}

	zones, zoneErr := client.ListZones("baa.com", "baa.net", "baa.org", "foo.net")
	if zoneErr != nil {
		log.Fatalf("[ERROR] Failed to fetch Cloudflare zones: %s", zoneErr)
	}

	if len(zones) == 0 {
		log.Print("[DEBUG] No Cloudflare Zones to sweep")
		return nil
	}

	for _, zone := range zones {
		log.Printf("[INFO] Deleting Cloudflare Zone ID: %s", zone.ID)
		_, err := client.DeleteZone(zone.ID)

		if err != nil {
			log.Printf("[ERROR] Failed to delete Cloudflare Zone (%s): %s", zone.Name, err)
		}
	}

	return nil
}

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
					resource.TestCheckResourceAttrSet("data.cloudflare_zones.examples_domains", "zones.#"),
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
						return i > 0
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
    name   = "baa.*"
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
    name   = "baa.*"
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
