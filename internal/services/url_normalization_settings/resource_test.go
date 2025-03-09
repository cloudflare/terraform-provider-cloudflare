package url_normalization_settings_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

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
		},
	})
}

func testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, _type, scope, name string) string {
	return acctest.LoadTestCase("urlnormalizationsettingsconfig.tf", zoneID, _type, scope, name)
}
