package schema_validation_settings_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareSchemaValidationZoneSettings(t *testing.T) {
	rndResourceName := utils.GenerateRandomResourceName()

	// resourceID is resourceIdentifier . resourceName
	resourceID := "cloudflare_schema_validation_settings." + rndResourceName
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		IsUnitTest:               false,
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// this sets up a schema and a zone-level block mitigation action
			{
				Config: testAccCloudflareSchemaValidationSettings(rndResourceName, zoneID, "block", nil),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "validation_default_mitigation_action", "block"),
					resource.TestCheckNoResourceAttr(resourceID, "validation_override_mitigation_action"),
				),
			},
			// change the default
			{
				Config: testAccCloudflareSchemaValidationSettings(rndResourceName, zoneID, "none", nil),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "validation_default_mitigation_action", "none"),
					resource.TestCheckNoResourceAttr(resourceID, "validation_override_mitigation_action"),
				),
			},
			// add an override and change back default to block
			{
				Config: testAccCloudflareSchemaValidationSettings(rndResourceName, zoneID, "block", acctest.PtrTo("none")),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "validation_default_mitigation_action", "block"),
					resource.TestCheckResourceAttr(resourceID, "validation_override_mitigation_action", "none"),
				),
			},
			// clear the override
			{
				Config: testAccCloudflareSchemaValidationSettings(rndResourceName, zoneID, "block", acctest.PtrTo("null")),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "validation_default_mitigation_action", "block"),
					resource.TestCheckNoResourceAttr(resourceID, "validation_override_mitigation_action"),
				),
			},
		},
	})
}

func testAccCloudflareSchemaValidationSettings(resourceName string, zone string, validation_default_mitigation_action string, validation_override_mitigation_action *string) string {
	block := fmt.Sprintf(`  validation_default_mitigation_action = "%s"`, validation_default_mitigation_action)
	if validation_override_mitigation_action != nil {
		override := *validation_override_mitigation_action
		if override != "null" {
			override = fmt.Sprintf(`"%s"`, override)
		}
		block += fmt.Sprintf("\n  validation_override_mitigation_action = %s", override)
	}
	return acctest.LoadTestCase("schema.tf", resourceName, zone, block)
}
