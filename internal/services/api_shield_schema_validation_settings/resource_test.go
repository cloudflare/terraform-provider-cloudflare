package api_shield_schema_validation_settings_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAPIShieldSchemaValidationSettings_Create(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_api_shield_schema_validation_settings." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckAPIShieldSchemaValidationSettingsDelete,
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

// func testAccCheckAPIShieldSchemaValidationSettingsDelete(s *terraform.State) error {
// 	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
// 	if clientErr != nil {
// 		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
// 	}

// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "cloudflare_api_shield_schema_validation_settings" {
// 			continue
// 		}

// 		result, err := client.GetAPIShieldSchemaValidationSettings(
// 			context.Background(),
// 			cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
// 		)
// 		if err != nil {
// 			return fmt.Errorf("encountered error getting schema validation settings: %w", err)
// 		}

// 		if result.DefaultMitigationAction != cloudflareAPIShieldSchemaValidationSettingsDefault().DefaultMitigationAction {
// 			return fmt.Errorf("expected validation_default_mitigation_action to be 'none' but got: %s", result.DefaultMitigationAction)
// 		}

// 		if result.OverrideMitigationAction != nil {
// 			return fmt.Errorf("expected validation_override_mitigation_action to be nil")
// 		}
// 	}

// 	return nil
// }

func testAccCloudflareAPIShieldSchemaValidationSettingsDefaultMitigationSet(resourceName, zone string) string {
	return acctest.LoadTestCase("apishieldschemavalidationsettingsdefaultmitigationset.tf", resourceName, zone)
}

func testAccCloudflareAPIShieldSchemaValidationSettingsAllMitigationsSet(resourceName, zone string) string {
	return acctest.LoadTestCase("apishieldschemavalidationsettingsallmitigationsset.tf", resourceName, zone)
}
