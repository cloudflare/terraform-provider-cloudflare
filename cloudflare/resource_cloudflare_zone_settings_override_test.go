package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"reflect"
	"strings"
)

func TestAccCloudflareZoneSettingsOverride_Empty(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_zone_settings_override.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareZoneSettingsOverrideConfigEmpty(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZoneSettingsEmpty(name),
				),
			},
		},
	})
}

func TestAccCloudflareZoneSettingsOverride_Full(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_zone_settings_override.test"

	initialSettings := make(map[string]interface{})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareZoneSettingsOverrideConfigEmpty(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccGetInitialZoneSettings(t, zoneName, initialSettings),
				),
			},
			{
				Config: testAccCheckCloudflareZoneSettingsOverrideConfigNormal(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZoneSettings(name),
					resource.TestCheckResourceAttr(
						name, "settings.0.brotli", "on"),
					resource.TestCheckResourceAttr(
						name, "settings.0.challenge_ttl", "2700"),
					resource.TestCheckResourceAttr(
						name, "settings.0.security_level", "high"),
				),
			},
		},
		CheckDestroy: testAccCheckInitialZoneSettings(zoneName, initialSettings),
	})
}

func TestAccCloudflareZoneSettingsOverride_RemoveAttributes(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_zone_settings_override.test"

	initialSettings := make(map[string]interface{})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareZoneSettingsOverrideConfigEmpty(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccGetInitialZoneSettings(t, zoneName, initialSettings),
				),
			},
			{
				Config: testAccCheckCloudflareZoneSettingsOverrideConfigNormal(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZoneSettings(name),
				),
			},
			{
				Config: testAccCheckCloudflareZoneSettingsOverrideConfigEmpty(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZoneSettings(name),
				),
			},
		},
		CheckDestroy: testAccCheckInitialZoneSettings(zoneName, initialSettings),
	})
}

func testAccCheckCloudflareZoneSettingsEmpty(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Zone ID is set")
		}

		for k, v := range rs.Primary.Attributes {
			if strings.Contains(k, "initial_settings") && k != "initial_settings_read_at" && !strings.Contains(k, "#") {
				currentSettingKey := strings.TrimPrefix(k, "initial_")

				currentVal := rs.Primary.Attributes[currentSettingKey]
				if currentVal != "" && currentVal != "0" {
					return fmt.Errorf("Current setting for %q: %q is not equal to initial setting for %q: %q",
						currentSettingKey, rs.Primary.Attributes[currentSettingKey], k, v)
				}
			}
		}
		return nil
	}
}

func testAccCheckCloudflareZoneSettings(n string) resource.TestCheckFunc {
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

func testAccGetInitialZoneSettings(t *testing.T, zoneName string, settings map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)

		zoneID, err := client.ZoneIDByName(zoneName)
		if err != nil {
			t.Fatalf("couldn't find zone %q: %s ", zoneName, err.Error())
		}
		foundZone, err := client.ZoneSettings(zoneID)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if foundZone.Result == nil || len(foundZone.Result) == 0 {
			t.Fatalf("Zone settings not found")
		}

		for _, zs := range foundZone.Result {
			settings[zs.ID] = zs.Value
		}
		return nil
	}
}

func testAccCheckInitialZoneSettings(zoneName string, initialSettings map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)

		zoneID, err := client.ZoneIDByName(zoneName)
		if err != nil {
			return fmt.Errorf("couldn't find zone %q : %s", zoneName, err.Error())
		}
		foundZone, err := client.ZoneSettings(zoneID)
		if err != nil {
			return err
		}

		if foundZone.Result == nil || len(foundZone.Result) == 0 {
			return fmt.Errorf("Zone settings not found")
		}

		for _, zs := range foundZone.Result {
			if !reflect.DeepEqual(zs.Value, initialSettings[zs.ID]) {
				return fmt.Errorf("Final setting for %q: %+v not equal to initial setting: %+v", zs.ID, zs.Value, initialSettings[zs.ID])
			}
		}
		return nil
	}
}

func testAccCheckCloudflareZoneSettingsOverrideConfigEmpty(zone string) string {
	return fmt.Sprintf(`
resource "cloudflare_zone_settings_override" "test" {
	name = "%s"
}`, zone)
}

func testAccCheckCloudflareZoneSettingsOverrideConfigNormal(zone string) string {
	return fmt.Sprintf(`
resource "cloudflare_zone_settings_override" "test" {
	name = "%s"
	settings {
		brotli = "on",
		challenge_ttl = 2700
		security_level = "high"
		opportunistic_encryption = "on"
		automatic_https_rewrites = "on"
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
