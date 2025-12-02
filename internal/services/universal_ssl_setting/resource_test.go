package universal_ssl_setting_test

import (
	"context"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_universal_ssl_setting", &resource.Sweeper{
		Name: "cloudflare_universal_ssl_setting",
		F:    testSweepCloudflareUniversalSSLSetting,
	})
}

func testSweepCloudflareUniversalSSLSetting(r string) error {
	ctx := context.Background()
	// Universal SSL Setting is a zone-level SSL configuration setting.
	// It's a singleton setting per zone, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Universal SSL Setting doesn't require sweeping (zone setting)")
	return nil
}

func testUniversalSSLSettingConfig(rnd, zoneID string, enabled bool) string {
	return acctest.LoadTestCase("universalssl.tf", rnd, zoneID, enabled)
}

func TestAccCloudflareUniversalSSLSetting_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_universal_ssl_setting." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create with enabled = true
				Config: testUniversalSSLSettingConfig(rnd, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
			{
				// Update to enabled = false
				Config: testUniversalSSLSettingConfig(rnd, zoneID, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
				),
			},
			{
				// Import test
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
