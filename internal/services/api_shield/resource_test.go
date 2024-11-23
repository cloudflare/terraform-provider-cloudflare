package api_shield_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAPIShieldBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_api_shield." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShield(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.#", "1"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.%", "2"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.type", "header"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.name", "my-example-header"),
				),
			},
			{
				Config: testAccCloudflareAPIShieldDeleteAuthCharacteristics(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.#", "0"),
					resource.TestCheckResourceAttr(resourceID, "success", "true"),
				),
			},
		},
	})
}

func testAccCloudflareAPIShield(resourceName, zone string) string {
	return acctest.LoadTestCase("apishield.tf", resourceName, zone)
}

func testAccCloudflareAPIShieldDeleteAuthCharacteristics(resourceName, zone string) string {
	return acctest.LoadTestCase("apishieldeleteauthcharacteristics.tf", resourceName, zone)
}
