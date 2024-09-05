package zero_trust_access_key_configuration_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAccessKeysConfiguration_WithKeyRotationIntervalDaysSet(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_key_configuration.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
	return acctest.LoadTestCase("accesskeysconfigurationwithkeyrotationintervaldays.tf", rnd, accountID, days)
}
