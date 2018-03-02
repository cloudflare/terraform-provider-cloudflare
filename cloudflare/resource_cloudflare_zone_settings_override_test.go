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
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_zone_settings_override.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareZoneSettingsOverrideConfigEmpty(zoneName),
				Check: resource.ComposeTestCheckFunc(
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
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_zone_settings_override.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareZoneSettingsOverrideConfigNormal(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareZoneSettings(name),
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

func testAccCheckCloudFlareZoneSettings(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Zone ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundZone, err := client.ZoneSettings(rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundZone.Result == nil || len(foundZone.Result) == 0 {
			return fmt.Errorf("Zone settings not found")
		}

		foundSettings := map[string]interface{}{}
		for _, zs := range foundZone.Result {
			if zs.ID == "brotli" && zs.Value == "on" ||
				zs.ID == "challenge_ttl" && zs.Value == float64(2700) ||
				zs.ID == "security_level" && zs.Value == "high" {
				foundSettings[zs.ID] = zs.Value
			} else if zs.ID == "brotli" || zs.ID == "challenge_ttl" || zs.ID == "security_level" {
				return fmt.Errorf("unexpected value for %q at API: %#v", zs.ID, zs.Value)
			}
		}
		if len(foundSettings) != 3 {
			return fmt.Errorf("expected to find 3 attributes matching the expected values but only got %d: %#v", len(foundSettings), foundSettings)
		}

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
