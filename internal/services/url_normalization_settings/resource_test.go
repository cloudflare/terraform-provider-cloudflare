package url_normalization_settings_test

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go/v4/url_normalization"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"log"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pkg/errors"
)

func init() {
	resource.AddTestSweepers("cloudflare_url_normalization_settings", &resource.Sweeper{
		Name: "cloudflare_url_normalization_settings",
		F:    testSweepCloudflareURLNormalizationSettings,
	})
}

func testSweepCloudflareURLNormalizationSettings(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	// Clean up the account level rulesets
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	settings, err := client.URLNormalization.Get(context.Background(), url_normalization.URLNormalizationGetParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare url normalization settings: %s", err))
	}

	if settings == nil {
		log.Print("[DEBUG] No Cloudflare url normalization settings to sweep")
		return nil
	}

	err = client.URLNormalization.Delete(context.Background(), url_normalization.URLNormalizationDeleteParams{
		ZoneID: cloudflare.F(zoneID),
	})

	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to delete Cloudflare url normalization settings: %s", err))
	}

	return nil
}

func TestAccCloudflareURLNormalizationSettings_CreateThenUpdate(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_url_normalization_settings.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "cloudflare", "incoming", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "type", "cloudflare"),
					resource.TestCheckResourceAttr(name, "scope", "incoming"),
				),
			},
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "cloudflare", "both", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "type", "cloudflare"),
					resource.TestCheckResourceAttr(name, "scope", "both"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, _type, scope, name string) string {
	return acctest.LoadTestCase("urlnormalizationsettingsconfig.tf", zoneID, _type, scope, name)
}
