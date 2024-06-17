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

func TestAccCloudflareAPIShieldSchemaValidationSettings_Create(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceID := "cloudflare_api_shield_schema_validation_settings." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAPIShieldSchemaValidationSettingsDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldSchemaValidationSettingsDefaultMitigationSet(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "validation_default_mitigation_action", "log"),
					// default
					resource.TestCheckResourceAttr(resourceID, "validation_override_mitigation_action", ""),
				),
			},
			{
				Config: testAccCloudflareAPIShieldSchemaValidationSettingsAllMitigationsSet(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "validation_default_mitigation_action", "block"),
					resource.TestCheckResourceAttr(resourceID, "validation_override_mitigation_action", "none"),
				),
			},
		},
	})
}

func testAccCheckAPIShieldSchemaValidationSettingsDelete(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_api_shield_schema_validation_settings" {
			continue
		}

		result, err := client.GetAPIShieldSchemaValidationSettings(
			context.Background(),
			cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
		)
		if err != nil {
			return fmt.Errorf("encountered error getting schema validation settings: %w", err)
		}

		if result.DefaultMitigationAction != cloudflareAPIShieldSchemaValidationSettingsDefault().DefaultMitigationAction {
			return fmt.Errorf("expected validation_default_mitigation_action to be 'none' but got: %s", result.DefaultMitigationAction)
		}

		if result.OverrideMitigationAction != nil {
			return fmt.Errorf("expected validation_override_mitigation_action to be nil")
		}
	}

	return nil
}

func testAccCloudflareAPIShieldSchemaValidationSettingsDefaultMitigationSet(resourceName, zone string) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_shield_schema_validation_settings" "%[1]s" {
		zone_id = "%[2]s"
		validation_default_mitigation_action = "log"
	}
`, resourceName, zone)
}

func testAccCloudflareAPIShieldSchemaValidationSettingsAllMitigationsSet(resourceName, zone string) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_shield_schema_validation_settings" "%[1]s" {
		zone_id = "%[2]s"
		validation_default_mitigation_action = "block"
		validation_override_mitigation_action = "none"
	}
`, resourceName, zone)
}
