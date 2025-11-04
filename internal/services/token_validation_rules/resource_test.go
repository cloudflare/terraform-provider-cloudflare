package token_validation_rules_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareTokenValidationRules(t *testing.T) {
	rndResourceName := utils.GenerateRandomResourceName()

	// resourceName is resourceIdentifier . resourceName
	resourceName := "cloudflare_token_validation_rules." + rndResourceName
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		IsUnitTest:               false,
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// create a new rule but keep it disabled
			{
				Config: testAccCloudflareTokenRules(rndResourceName, zoneID, "title", "description", "block", false),
				Check: func(s *terraform.State) error {
					tokenConfigID := s.RootModule().Resources["cloudflare_token_validation_config."+rndResourceName].Primary.ID
					operationID := s.RootModule().Resources["cloudflare_api_shield_operation."+rndResourceName].Primary.ID

					return resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
						resource.TestCheckResourceAttr(resourceName, "title", "title"),
						resource.TestCheckResourceAttr(resourceName, "description", "description"),
						resource.TestCheckResourceAttr(resourceName, "action", "block"),
						resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
						resource.TestCheckResourceAttr(resourceName, "expression", fmt.Sprintf("(is_jwt_valid(\"%s\"))", tokenConfigID)),
						resource.TestCheckResourceAttr(resourceName, "selector.include.0.host.0", "example.com"),
						resource.TestCheckResourceAttr(resourceName, "selector.exclude.0.operation_ids.0", operationID),
						resource.TestCheckResourceAttrSet(resourceName, "created_at"),
						resource.TestCheckResourceAttrSet(resourceName, "last_updated"),
					)(s)
				},
			},

			// enable the rule
			{
				Config: testAccCloudflareTokenRules(rndResourceName, zoneID, "title", "description", "block", true),
				Check: func(s *terraform.State) error {
					tokenConfigID := s.RootModule().Resources["cloudflare_token_validation_config."+rndResourceName].Primary.ID
					operationID := s.RootModule().Resources["cloudflare_api_shield_operation."+rndResourceName].Primary.ID

					return resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
						resource.TestCheckResourceAttr(resourceName, "title", "title"),
						resource.TestCheckResourceAttr(resourceName, "description", "description"),
						resource.TestCheckResourceAttr(resourceName, "action", "block"),
						resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
						resource.TestCheckResourceAttr(resourceName, "expression", fmt.Sprintf("(is_jwt_valid(\"%s\"))", tokenConfigID)),
						resource.TestCheckResourceAttr(resourceName, "selector.include.0.host.0", "example.com"),
						resource.TestCheckResourceAttr(resourceName, "selector.exclude.0.operation_ids.0", operationID),
						resource.TestCheckResourceAttrSet(resourceName, "created_at"),
						resource.TestCheckResourceAttrSet(resourceName, "last_updated"),
					)(s)
				},
			},
			// deletes are implicitly tested

			// ensure import works
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					rs, ok := state.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("not found: %s", resourceName)
					}
					return fmt.Sprintf("%s/%s", zoneID, rs.Primary.ID), nil
				},
			},
		},
	})
}

func testAccCloudflareTokenRules(resourceName, zone string, title string, description string, action string, enabled bool) string {
	return acctest.LoadTestCase("rules.tf", resourceName, zone, title, description, action, fmt.Sprintf("%v", enabled))
}

func checkHasField(name string) resource.CheckResourceAttrWithFunc {
	return func(value string) error {
		if len(value) > 0 {
			return nil
		}
		return fmt.Errorf("%s is empty", name)
	}
}
