package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCloudFlareZone_Basic(t *testing.T) {
	var zone cloudflare.Zone
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_zone_settings_override.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareZoneSettingsOverrideConfigEmpty(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareZoneExists(name, &zone),
					resource.TestCheckResourceAttrSet(
						name, "settings.0.brotli"),
					resource.TestCheckResourceAttrSet(
						name, "settings.0.challenge_ttl"),
					resource.TestCheckResourceAttrSet(
						name, "settings.0.security_level"),
				),
			},
		},
	})
}

func TestAccCloudFlareZone_Overrides(t *testing.T) {
	var zone cloudflare.Zone
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_zone_settings_override.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareZoneSettingsOverrideConfigNormal(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareZoneExists(name, &zone),
					resource.TestCheckResourceAttr(
						name, "settings.0.brotli", "on"),
					resource.TestCheckResourceAttr(
						name, "settings.0.challenge_ttl", "2700"),
					resource.TestCheckResourceAttr(
						name, "settings.0.security_level", "high"),
				),
			},
		},
	})
}

func testAccCheckCloudFlareZoneExists(n string, zone *cloudflare.Zone) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Zone ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundZone, err := client.ZoneDetails(rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundZone.ID != rs.Primary.ID {
			return fmt.Errorf("Zone not found")
		}

		*zone = foundZone

		return nil
	}
}

func testAccCheckCloudFlareZoneSettingsOverrideConfigEmpty(zone string) string {
	return fmt.Sprintf(`
resource "cloudflare_zone_settings_override" "test" {
	name = "%s"
}`, zone)
}

func testAccCheckCloudFlareZoneSettingsOverrideConfigNormal(zone string) string {
	return fmt.Sprintf(`
resource "cloudflare_zone_settings_override" "test" {
	name = "%s"
	settings {
		brotli = "on",
		challenge_ttl = 2700
		security_level = "high"
		minify {
			css = "on"
			js = "off"
			html = "off"
		}
		security_header {
			enabled = true
		}
	}
}`, zone)
}
