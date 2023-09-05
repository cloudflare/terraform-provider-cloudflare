package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_hostname_tls_setting", &resource.Sweeper{
		Name: "cloudflare_hostname_tls_setting",
		F:    testSweepCloudflareHostnameTLSSettings,
	})
}

func testSweepCloudflareHostnameTLSSettings(r string) error {
	ctx := context.Background()
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneIDrc := cloudflare.ZoneIdentifier(zoneID)
	params := cloudflare.ListHostnameTLSSettingsParams{
		Setting: "min_tls_version",
	}
	settings, resultInfo, err := client.ListHostnameTLSSettings(ctx, zoneIDrc, params)
	if err != nil {
		return err
	}
	if resultInfo.Count == 0 {
		tflog.Debug(ctx, "no hostname tls settings to sweep")
		return nil
	}

	for _, setting := range settings {
		deleteParams := cloudflare.DeleteHostnameTLSSettingParams{
			Setting:  "min_tls_version",
			Hostname: setting.Hostname,
		}
		_, err := client.DeleteHostnameTLSSetting(ctx, zoneIDrc, deleteParams)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("failed to delete hostname tls setting for hostname (%s) in zone ID: %s", setting.Hostname, zoneID))
		}
	}

	return nil
}

func TestAccCloudflareHostnameTLSSetting_Basic(t *testing.T) {
	t.Parallel()
	var htlss cloudflare.HostnameTLSSetting
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	hostname := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_hostname_tls_setting." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHostnameTLSSettingConfig(zoneID, fmt.Sprintf("%s.%s", rnd, hostname), "min_tls_version", "1.2", rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareHostnameTLSSettingExists(resourceName, &htlss),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
		CheckDestroy: testAccCheckCloudflareHostnameTLSSettingDestroy,
	})
}

func testAccCheckCloudflareHostnameTLSSettingConfig(zoneID, hostname, setting, value, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_hostname_tls_setting" "%[5]s" {
	zone_id	= "%[1]s"
	hostname = "%[2]s"
	setting = "%[3]s"
	value = "%[4]s"
}
`, zoneID, hostname, setting, value, rnd)
}

func testAccCheckCloudflareHostnameTLSSettingExists(name string, htlss *cloudflare.HostnameTLSSetting) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		zoneIDrc := cloudflare.ZoneIdentifier(zoneID)
		params := cloudflare.ListHostnameTLSSettingsParams{
			Setting:  "min_tls_version",
			Hostname: []string{rs.Primary.Attributes["hostname"]},
		}
		settings, resultInfo, err := client.ListHostnameTLSSettings(context.Background(), zoneIDrc, params)
		if err != nil {
			return err
		}
		if resultInfo.Count != 1 {
			return fmt.Errorf("incorrect number of settings returned (%d), should be 1", resultInfo.Count)
		}
		if settings[0].Hostname != rs.Primary.Attributes["hostname"] {
			return fmt.Errorf("incorrect hostname returned")
		}
		*htlss = settings[0]
		return nil
	}
}

func testAccCheckCloudflareHostnameTLSSettingDestroy(s *terraform.State) error {
	// sleep in order to allow htlss to enter active state before being deleted
	// lintignore: R018
	time.Sleep(time.Second)
	client := testAccProvider.Meta().(*cloudflare.API)
	for _, rs := range s.RootModule().Resources {
		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		zoneIDrc := cloudflare.ZoneIdentifier(zoneID)
		_, err := client.DeleteHostnameTLSSetting(context.Background(), zoneIDrc, cloudflare.DeleteHostnameTLSSettingParams{Hostname: rs.Primary.ID, Setting: "min_tls_version"})
		if err == nil {
			return fmt.Errorf("error deleting hostname tls setting in zone %q: %w", zoneID, err)
		}
	}
	return nil
}
