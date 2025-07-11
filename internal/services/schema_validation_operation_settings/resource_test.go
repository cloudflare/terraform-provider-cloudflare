package schema_validation_operation_settings_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflarePerOperationSetting(t *testing.T) {
	rndResourceName := utils.GenerateRandomResourceName()

	// resourceID is resourceIdentifier . resourceName
	resourceID := "cloudflare_schema_validation_operation_settings." + rndResourceName
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		IsUnitTest:               false,
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// this sets up a schema and a block per-operation mitigation action (note: we're not testing log as it might require further permissions)
			{
				Config: testAccCloudflareSchemaValidationWithOperationMitigationAction(rndResourceName, zoneID, "block"),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttrWith(resourceID, "operation_id", func(value string) error {
						_, err := uuid.ParseUUID(value)
						return err
					}),
					resource.TestCheckResourceAttr(resourceID, "mitigation_action", "block"),
				),
			},
			// update to skip this operation
			{
				Config: testAccCloudflareSchemaValidationWithOperationMitigationAction(rndResourceName, zoneID, "none"),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttrWith(resourceID, "operation_id", func(value string) error {
						_, err := uuid.ParseUUID(value)
						return err
					}),
					resource.TestCheckResourceAttr(resourceID, "mitigation_action", "none"),
				),
			},
		},
	})
}

func testAccCloudflareSchemaValidationWithOperationMitigationAction(resourceName string, zone string, action string) string {
	return acctest.LoadTestCase("schema.tf", resourceName, zone, action)
}
