package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
)

func TestAccCloudFlareZoneSettingsOverride_Empty(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_zone_settings_override.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareZoneSettingsOverrideConfigEmpty(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareZoneSettingsUnchanged(name),
				),
			},
		},
	})
}

func TestAccCloudFlareZoneSettingsOverride_Full(t *testing.T) {
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

func testAccCheckCloudFlareZoneSettingsUnchanged(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Zone ID is set")
		}

		for k, v := range rs.Primary.Attributes {
			if strings.Contains(k, "initial_settings") && k != "initial_settings_read_at" {
				currentSettingKey := strings.TrimPrefix(k, "initial_")

				if v != rs.Primary.Attributes[currentSettingKey] {
					return fmt.Errorf("Current setting for %q: %q is not equal to initial setting for %q: %q",
						currentSettingKey, rs.Primary.Attributes[currentSettingKey], k, v)
				}
			}
		}
		return nil
	}
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
		opportunistic_encryption = "on"
		automatic_https_rewrites = "on"
		always_use_https = "off"
		polish = "off"
		webp = "on"
		mirage = "on"
		waf = "on"
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
