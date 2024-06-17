package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAPIShieldOperationSchemaValidationSettings_Create(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceID := "cloudflare_api_shield_operation_schema_validation_settings." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	block := "block"
	none := "none"
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldOperationSchemaValidationSettingsMitigation(rnd, zoneID, &block),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "mitigation_action", "block"),
				),
			},
			{
				Config: testAccCloudflareAPIShieldOperationSchemaValidationSettingsMitigation(rnd, zoneID, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "mitigation_action", ""),
				),
			},
			{
				Config: testAccCloudflareAPIShieldOperationSchemaValidationSettingsMitigation(rnd, zoneID, &none),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "mitigation_action", "none"),
				),
			},
			{
				Config: testAccCloudflareAPIShieldOperationSchemaValidationSettingsNoMitigation(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "mitigation_action", ""),
				),
			},
		},
	})
}

func testAccCloudflareAPIShieldOperationSchemaValidationSettingsMitigation(resourceName, zone string, mitigation *string) string {
	action := "null"
	if mitigation != nil {
		action = fmt.Sprintf("\"%s\"", *mitigation)
	}

	return fmt.Sprintf(`
	resource "cloudflare_api_shield_operation" "terraform_test_acc_operation" {
		zone_id = "%[2]s"
		host = "foo.com"
		method = "GET"
        endpoint = "/api"
	}
	resource "cloudflare_api_shield_operation_schema_validation_settings" "%[1]s" {
		zone_id = "%[2]s"
		operation_id = cloudflare_api_shield_operation.terraform_test_acc_operation.id
		mitigation_action = %[3]s
	}
`, resourceName, zone, action)
}

func testAccCloudflareAPIShieldOperationSchemaValidationSettingsNoMitigation(resourceName, zone string) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_shield_operation" "terraform_test_acc_operation" {
		zone_id = "%[2]s"
		host = "foo.com"
		method = "GET"
        endpoint = "/api"
	}
	resource "cloudflare_api_shield_operation_schema_validation_settings" "%[1]s" {
		zone_id = "%[2]s"
		operation_id = cloudflare_api_shield_operation.terraform_test_acc_operation.id
	}
`, resourceName, zone)
}
