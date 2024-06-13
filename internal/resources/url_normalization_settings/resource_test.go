package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareURLNormalizationSettings_CreateThenUpdate(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_url_normalization_settings.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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
		CheckDestroy: testAccCheckCloudflareURLNormalizationSettingsDestroy,
	})
}

func testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, _type, scope, name string) string {
	return fmt.Sprintf(`
				resource "cloudflare_url_normalization_settings" "%[4]s" {
					zone_id = "%[1]s"
					type = "%[2]s"
					scope = "%[3]s"
				}`, zoneID, _type, scope, name)
}

func testAccCheckCloudflareURLNormalizationSettingsDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_url_normalization_settings" {
			continue
		}

		settings, err := client.URLNormalizationSettings(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]))
		if err != nil {
			return err
		}

		if settings.Type != "cloudflare" {
			return fmt.Errorf("expected Type to be reset to cloudflare, got: %s", settings.Type)
		}

		if settings.Scope != "incoming" {
			return fmt.Errorf("expected Scope to be reset to both, got: %s", settings.Scope)
		}
	}

	return nil
}
