package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareAccessKeysConfiguration_WithKeyRotationIntervalDaysSet(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_keys_configuration.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessKeysConfigurationWithKeyRotationIntervalDays(rnd, accountID, 60),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "key_rotation_interval_days", "60"),
				),
			},
		},
	})
}

func testAccessKeysConfigurationWithKeyRotationIntervalDays(rnd, accountID string, days int) string {
	return fmt.Sprintf(`
resource "cloudflare_access_keys_configuration" "%[1]s" {
  account_id = "%[2]s"
  key_rotation_interval_days = "%[3]d"
}`, rnd, accountID, days)
}

func TestAccCloudflareAccessKeysConfiguration_WithoutKeyRotationIntervalDaysSet(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_keys_configuration.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessKeysConfigurationWithoutKeyRotationIntervalDays(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "key_rotation_interval_days"),
				),
			},
		},
	})
}

func testAccessKeysConfigurationWithoutKeyRotationIntervalDays(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_keys_configuration" "%[1]s" {
  account_id = "%[2]s"
}`, rnd, accountID)
}
