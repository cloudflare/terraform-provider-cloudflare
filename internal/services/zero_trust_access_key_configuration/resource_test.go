package zero_trust_access_key_configuration_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_access_key_configuration", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_key_configuration",
		F:    testSweepCloudflareZeroTrustAccessKeyConfiguration,
	})
}

func testSweepCloudflareZeroTrustAccessKeyConfiguration(r string) error {
	ctx := context.Background()
	// Access Key Configuration is an account-level key rotation setting.
	// It's a singleton setting per account, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Zero Trust Access Key Configuration doesn't require sweeping (account setting)")
	return nil
}

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
