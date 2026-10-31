package page_shield_settings_test

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
)

func init() {
	resource.AddTestSweepers("cloudflare_page_shield_settings", &resource.Sweeper{
		Name: "cloudflare_page_shield_settings",
		F:    testSweepCloudflarePageShieldSettings,
	})
}

func testSweepCloudflarePageShieldSettings(r string) error {
	ctx := context.Background()
	// Page Shield Settings is a zone-level configuration setting.
	// It's a singleton setting per zone, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Page Shield Settings doesn't require sweeping (zone setting)")
	return nil
}

func TestAccCloudflarePageShieldSettings_Create(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_page_shield_settings.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflarePageShieldSettings(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
		},
	})
}

func testAccCloudflarePageShieldSettings(resourceName, zoneID string) string {
	return acctest.LoadTestCase("pageshieldsettings.tf", resourceName, zoneID)
}
