package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflarePageShield_Create(t *testing.T) {
	rnd := generateRandomResourceName()
	resourceID := "cloudflare_page_shield." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckPageShieldDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflarePageShieldDefaultEnabledSet(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "enabled", "true"),
				),
			},
			{
				Config: testAccCloudflarePageShieldAllFieldsSet(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceID, "use_cloudflare_reporting_endpoint", "true"),
					resource.TestCheckResourceAttr(resourceID, "use_connection_url_path", "true"),
				),
			},
		},
	})
}

func testAccCheckPageShieldDelete(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_page_shield" {
			continue
		}

		result, err := client.GetPageShieldSettings(
			context.Background(),
			cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
			cloudflare.GetPageShieldSettingsParams{},
		)
		if err != nil {
			return fmt.Errorf("encountered error getting page shield: %w", err)
		}

		if *result.PageShield.Enabled != false {
			return fmt.Errorf("expected enabled to be 'false' but got: %s", strconv.FormatBool(*result.PageShield.Enabled))
		}

		if *result.PageShield.UseCloudflareReportingEndpoint != true {
			return fmt.Errorf("expected use_cloudflare_reporting_endpoint to be 'true' but got: %s", strconv.FormatBool(*result.PageShield.UseCloudflareReportingEndpoint))
		}

		if *result.PageShield.UseConnectionURLPath != false {
			return fmt.Errorf("expected use_connection_url_path to be 'false' but got: %s", strconv.FormatBool(*result.PageShield.UseConnectionURLPath))
		}
	}

	return nil
}

func testAccCloudflarePageShieldDefaultEnabledSet(resourceName, zone string) string {
	return fmt.Sprintf(`
	resource "cloudflare_page_shield" "%[1]s" {
		zone_id = "%[2]s"
	}
`, resourceName, zone)
}

func testAccCloudflarePageShieldAllFieldsSet(resourceName, zone string) string {
	return fmt.Sprintf(`
	resource "cloudflare_page_shield" "%[1]s" {
		zone_id = "%[2]s"
        use_cloudflare_reporting_endpoint = true
        use_connection_url_path = true
	}
`, resourceName, zone)
}
