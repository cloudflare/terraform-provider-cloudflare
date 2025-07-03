package schema_validation_schemas_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareSchemaValidationSchemas(t *testing.T) {
	rndResourceName := utils.GenerateRandomResourceName()

	// resourceID is resourceIdentifier . resourceName
	resourceID := "cloudflare_schema_validation_schemas." + rndResourceName
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		IsUnitTest:               false,
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// upload a schema but have it disabled
			{
				Config: testAccCloudflareSchemaValidation(rndResourceName, zoneID, false),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttrWith(resourceID, "schema_id", func(value string) error {
						_, err := uuid.ParseUUID(value)
						return err
					}),
					resource.TestCheckResourceAttr(resourceID, "name", rndResourceName+".yaml"),
					resource.TestCheckResourceAttr(resourceID, "kind", "openapi_v3"),
					resource.TestCheckResourceAttrWith(resourceID, "created_at", func(value string) error {
						_, err := time.Parse(time.RFC3339, value)
						return err
					}),
					resource.TestCheckResourceAttr(resourceID, "validation_enabled", "false"),
					resource.TestCheckResourceAttrWith(resourceID, "source", func(value string) error {
						if len(value) > 0 {
							return nil
						}
						return fmt.Errorf("value is empty but should not be")
					}),
				),
			},
			// enable validation
			{
				Config: testAccCloudflareSchemaValidation(rndResourceName, zoneID, true),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttrWith(resourceID, "schema_id", func(value string) error {
						_, err := uuid.ParseUUID(value)
						return err
					}),
					resource.TestCheckResourceAttr(resourceID, "name", rndResourceName+".yaml"),
					resource.TestCheckResourceAttr(resourceID, "kind", "openapi_v3"),
					resource.TestCheckResourceAttrWith(resourceID, "created_at", func(value string) error {
						_, err := time.Parse(time.RFC3339, value)
						return err
					}),
					resource.TestCheckResourceAttr(resourceID, "validation_enabled", "true"),
					resource.TestCheckResourceAttrWith(resourceID, "source", func(value string) error {
						if len(value) > 0 {
							return nil
						}
						return fmt.Errorf("value is empty but should not be")
					}),
				),
			},
			// plan the same thing again, this should be a noop
			{
				Config: testAccCloudflareSchemaValidation(rndResourceName, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttrWith(resourceID, "schema_id", func(value string) error {
						_, err := uuid.ParseUUID(value)
						return err
					}),
					resource.TestCheckResourceAttr(resourceID, "name", rndResourceName+".yaml"),
					resource.TestCheckResourceAttr(resourceID, "kind", "openapi_v3"),
					resource.TestCheckResourceAttrWith(resourceID, "created_at", func(value string) error {
						_, err := time.Parse(time.RFC3339, value)
						return err
					}),
					resource.TestCheckResourceAttr(resourceID, "validation_enabled", "true"),
					resource.TestCheckResourceAttrWith(resourceID, "source", func(value string) error {
						if len(value) > 0 {
							return nil
						}
						return fmt.Errorf("value is empty but should not be")
					}),
				),
				PlanOnly: true,
			},
			// now import the previously planned thing, this should again result in no change
			{
				ResourceName: resourceID,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					rs, ok := state.RootModule().Resources[resourceID]
					if !ok {
						return "", fmt.Errorf("not found: %s", resourceID)
					}
					return fmt.Sprintf("%s/%s", zoneID, rs.Primary.ID), nil
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudflareSchemaValidation(resourceName, zone string, validation_enabled bool) string {
	return acctest.LoadTestCase("schema.tf", resourceName, zone, fmt.Sprintf("%v", validation_enabled))
}
