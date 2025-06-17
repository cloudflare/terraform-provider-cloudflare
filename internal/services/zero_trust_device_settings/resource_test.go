package zero_trust_device_settings_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareZTDeviceSettings(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_settings.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsDeviceSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZTDeviceSettings(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "disable_for_time", "20"),
					resource.TestCheckResourceAttr(name, "gateway_proxy_enabled", "true"),
					resource.TestCheckResourceAttr(name, "gateway_udp_proxy_enabled", "true"),
					resource.TestCheckResourceAttr(name, "root_certificate_installation_enabled", "true"),
					resource.TestCheckResourceAttr(name, "use_zt_virtual_ip", "true"),
				),
			},
			{
				Config: testAccCloudflareZTDeviceSomeSettings(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "disable_for_time", "0"),
					resource.TestCheckResourceAttr(name, "gateway_proxy_enabled", "true"),
					resource.TestCheckResourceAttr(name, "gateway_udp_proxy_enabled", "false"),
					resource.TestCheckResourceAttr(name, "root_certificate_installation_enabled", "false"),
					resource.TestCheckResourceAttr(name, "use_zt_virtual_ip", "true"),
				),
			},
		},
	})
}

func testAccCloudflareZTDeviceSettings(rnd, accountID string) string {
	return acctest.LoadTestCase("zt_device_settings_init.tf", rnd, accountID)
}

func testAccCloudflareZTDeviceSomeSettings(rnd, accountID string) string {
	return acctest.LoadTestCase("zt_device_settings_update.tf", rnd, accountID)
}

func testAccCheckCloudflareTeamsDeviceSettingsDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_device_settings" {
			continue
		}

		settings, err := client.TeamsAccountDeviceConfiguration(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey])
		if err != nil || settings.GatewayProxyEnabled {
			return fmt.Errorf("teams device settings still exists")
		}
	}

	return nil
}
