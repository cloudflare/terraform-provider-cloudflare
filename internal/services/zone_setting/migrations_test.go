package zone_setting_test

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateZoneSettingMigrationFromV4Basic tests basic migration from v4 zone_settings_override to v5 zone_setting
func TestMigrateZoneSettingMigrationFromV4Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config using zone_settings_override with basic settings
	v4Config := fmt.Sprintf(`
resource "cloudflare_zone_settings_override" "%[1]s" {
  zone_id = "%[2]s"
  settings {
    http3 = "on"
  }
}`, rnd, zoneID)

	// Use MigrationV2TestStepAllowCreate to handle one-to-many transformation
	migrationSteps := acctest.MigrationV2TestStepAllowCreate(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		// Verify http3 setting
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_http3", rnd), tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_http3", rnd), tfjsonpath.New("setting_id"), knownvalue.StringExact("http3")),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_http3", rnd), tfjsonpath.New("value"), knownvalue.StringExact("on")),
	})
	
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			// Step 1: Create with v4 provider
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateZoneSettingMigrationFromV4WithSpecialSettings tests migration with special settings like 0rtt and security_header
func TestMigrateZoneSettingMigrationFromV4WithSpecialSettings(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config using zone_settings_override with special settings including zero_rtt -> 0rtt mapping
	v4Config := fmt.Sprintf(`
resource "cloudflare_zone_settings_override" "%[1]s" {
  zone_id = "%[2]s"
  settings {
    zero_rtt = "on"
    security_header {
      enabled = true
      max_age = 86400
      include_subdomains = true
      preload = false
      nosniff = false
    }
  }
}`, rnd, zoneID)

	// Use MigrationV2TestStepAllowCreate to handle one-to-many transformation
	migrationSteps := acctest.MigrationV2TestStepAllowCreate(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		// Verify zero_rtt -> 0rtt mapping
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_zero_rtt", rnd), tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_zero_rtt", rnd), tfjsonpath.New("setting_id"), knownvalue.StringExact("0rtt")),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_zero_rtt", rnd), tfjsonpath.New("value"), knownvalue.StringExact("on")),
		// Verify security_header nested block transformation
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_security_header", rnd), tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_security_header", rnd), tfjsonpath.New("setting_id"), knownvalue.StringExact("security_header")),
		// Check that security_header value is an object with the nested strict_transport_security attributes
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_security_header", rnd), tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("enabled"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_security_header", rnd), tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("max_age"), knownvalue.NumberExact(big.NewFloat(86400))),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_security_header", rnd), tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("include_subdomains"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_security_header", rnd), tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("preload"), knownvalue.Bool(false)),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_security_header", rnd), tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("nosniff"), knownvalue.Bool(false)),
	})
	
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			// Step 1: Create with v4 provider
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateZoneSettingMigrationFromV4WithNEL is skipped because NEL is a read-only setting
// that cannot be set via the API. This test would fail even with the v4 provider.
func TestMigrateZoneSettingMigrationFromV4WithNEL(t *testing.T) {
	t.Skip("NEL is a read-only setting and cannot be tested via the provider")
}

// TestMigrateZoneSettingMigrationFromV4Complex tests migration with multiple settings including variables
func TestMigrateZoneSettingMigrationFromV4Complex(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config using zone_settings_override with variables and complex settings
	v4Config := fmt.Sprintf(`
variable "my_zone_id" {
  default = "%[2]s"
}

variable "enable_http3" {
  default = "on"
}

resource "cloudflare_zone_settings_override" "%[1]s" {
  zone_id = var.my_zone_id
  settings {
    http3           = var.enable_http3
    min_tls_version = "1.2"
    always_use_https = "on"
    browser_check    = "on"
  }
}`, rnd, zoneID)

	// Use MigrationV2TestStepAllowCreate to handle one-to-many transformation
	migrationSteps := acctest.MigrationV2TestStepAllowCreate(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		// Verify http3 setting preserves variable reference
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_http3", rnd), tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_http3", rnd), tfjsonpath.New("setting_id"), knownvalue.StringExact("http3")),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_http3", rnd), tfjsonpath.New("value"), knownvalue.StringExact("on")),
		// Verify other settings
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_min_tls_version", rnd), tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_min_tls_version", rnd), tfjsonpath.New("setting_id"), knownvalue.StringExact("min_tls_version")),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_min_tls_version", rnd), tfjsonpath.New("value"), knownvalue.StringExact("1.2")),
		// Verify always_use_https setting
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_always_use_https", rnd), tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_always_use_https", rnd), tfjsonpath.New("setting_id"), knownvalue.StringExact("always_use_https")),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_always_use_https", rnd), tfjsonpath.New("value"), knownvalue.StringExact("on")),
		// Verify browser_check setting
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_browser_check", rnd), tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_browser_check", rnd), tfjsonpath.New("setting_id"), knownvalue.StringExact("browser_check")),
		statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zone_setting.%s_browser_check", rnd), tfjsonpath.New("value"), knownvalue.StringExact("on")),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			// Step 1: Create with v4 provider
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}