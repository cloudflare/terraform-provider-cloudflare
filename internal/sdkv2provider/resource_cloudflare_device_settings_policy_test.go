package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareDeviceSettingsPolicy_Create(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd, defaultRnd := generateRandomResourceName(), generateRandomResourceName()
	name, defaultName := fmt.Sprintf("cloudflare_device_settings_policy.%s", rnd), fmt.Sprintf("cloudflare_device_settings_policy.%s", defaultRnd)
	precedence := uint64(10)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareDeviceSettingsPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDeviceSettingsPolicy(rnd, accountID, precedence),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "allow_mode_switch", "true"),
					resource.TestCheckResourceAttr(name, "allow_updates", "true"),
					resource.TestCheckResourceAttr(name, "allowed_to_leave", "true"),
					resource.TestCheckResourceAttr(name, "auto_connect", "0"),
					resource.TestCheckResourceAttr(name, "captive_portal", "5"),
					resource.TestCheckResourceAttr(name, "default", "false"),
					resource.TestCheckResourceAttr(name, "disable_auto_fallback", "true"),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "match", "identity.email == \"foo@example.com\""),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "precedence", fmt.Sprintf("%d", precedence)),
					resource.TestCheckResourceAttr(name, "service_mode_v2_mode", "warp"),
					resource.TestCheckResourceAttr(name, "service_mode_v2_port", "0"),
					resource.TestCheckResourceAttr(name, "support_url", "https://cloudflare.com"),
					resource.TestCheckResourceAttr(name, "switch_locked", "true"),
					resource.TestCheckResourceAttr(name, "exclude_office_ips", "true"),
				),
			},
			{
				Config: testAccCloudflareDefaultDeviceSettingsPolicy(defaultRnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(defaultName, "id", accountID),
					resource.TestCheckResourceAttr(defaultName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(defaultName, "allow_mode_switch", "true"),
					resource.TestCheckResourceAttr(defaultName, "allow_updates", "true"),
					resource.TestCheckResourceAttr(defaultName, "allowed_to_leave", "true"),
					resource.TestCheckResourceAttr(defaultName, "auto_connect", "0"),
					resource.TestCheckResourceAttr(defaultName, "captive_portal", "5"),
					resource.TestCheckResourceAttr(defaultName, "default", "true"),
					resource.TestCheckResourceAttr(defaultName, "disable_auto_fallback", "true"),
					resource.TestCheckResourceAttr(defaultName, "enabled", "true"),
					resource.TestCheckResourceAttr(defaultName, "name", defaultRnd),
					resource.TestCheckResourceAttr(defaultName, "service_mode_v2_mode", "warp"),
					resource.TestCheckResourceAttr(defaultName, "service_mode_v2_port", "0"),
					resource.TestCheckResourceAttr(defaultName, "support_url", "https://cloudflare.com"),
					resource.TestCheckResourceAttr(defaultName, "switch_locked", "true"),
					resource.TestCheckResourceAttr(defaultName, "exclude_office_ips", "true"),
				),
			},
			{
				Config:      testAccCloudflareInvalidDefaultDeviceSettingsPolicy(rnd, accountID),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("match cannot be set for default policies")),
			},
		},
	})
}

func testAccCloudflareDeviceSettingsPolicy(rnd, accountID string, precedence uint64) string {
	return fmt.Sprintf(`
resource "cloudflare_device_settings_policy" "%[1]s" {
	account_id                = "%[2]s"
	allow_mode_switch         = true
	allow_updates             = true
	allowed_to_leave          = true
	auto_connect              = 0
	captive_portal            = 5
	disable_auto_fallback     = true
	enabled                   = true
	match                     = "identity.email == \"foo@example.com\""
	name                      = "%[1]s"
	precedence                = %[3]d
	support_url               = "https://cloudflare.com"
	switch_locked             = true
	exclude_office_ips		  = true
}
`, rnd, accountID, precedence)
}

func testAccCloudflareDefaultDeviceSettingsPolicy(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_settings_policy" "%[1]s" {
	account_id                = "%[2]s"
	default                   = true
	name                      = "%[1]s"
	allow_mode_switch         = true
	allow_updates             = true
	allowed_to_leave          = true
	auto_connect              = 0
	captive_portal            = 5
	disable_auto_fallback     = true
	enabled                   = true
	support_url               = "https://cloudflare.com"
	switch_locked             = true
	exclude_office_ips		  = true
}
`, rnd, accountID)
}

// invalid configuration - not allowed to set match for default policies.
func testAccCloudflareInvalidDefaultDeviceSettingsPolicy(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_settings_policy" "%[1]s" {
	account_id                = "%[2]s"
	default                   = true
	name                      = "%[1]s"
	allow_mode_switch         = true
	allow_updates             = true
	allowed_to_leave          = true
	auto_connect              = 0
	captive_portal            = 5
	disable_auto_fallback     = true
	support_url               = "https://cloudflare.com"
	switch_locked             = true
	match                     = "identity.email == \"foo@example.com\""
	exclude_office_ips		  = true
}
`, rnd, accountID)
}

func testAccCheckCloudflareDeviceSettingsPolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_device_settings_policy" {
			continue
		}

		_, policyID := parseDevicePolicyID(rs.Primary.ID)

		// cannot delete the default device setting policy
		if policyID == "" {
			return nil
		}

		_, err := client.GetDeviceSettingsPolicy(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], policyID)
		if err == nil {
			return fmt.Errorf("Device Posture Integration still exists")
		}
	}

	return nil
}
