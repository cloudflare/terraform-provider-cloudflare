package google_tag_gateway_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareGoogleTagGatewayDataSource_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_google_tag_gateway." + rnd
	dataSourceName := "data.cloudflare_google_tag_gateway." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGoogleTagGatewayDataSourceBasicConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, consts.ZoneIDSchemaKey, resourceName, consts.ZoneIDSchemaKey),
					resource.TestCheckResourceAttrPair(dataSourceName, "enabled", resourceName, "enabled"),
					resource.TestCheckResourceAttrPair(dataSourceName, "endpoint", resourceName, "endpoint"),
					resource.TestCheckResourceAttrPair(dataSourceName, "hide_original_ip", resourceName, "hide_original_ip"),
					resource.TestCheckResourceAttrPair(dataSourceName, "measurement_id", resourceName, "measurement_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "set_up_tag", resourceName, "set_up_tag"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						dataSourceName,
						tfjsonpath.New("zone_id"),
						knownvalue.StringExact(zoneID),
					),
					statecheck.ExpectKnownValue(
						dataSourceName,
						tfjsonpath.New("enabled"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						dataSourceName,
						tfjsonpath.New("endpoint"),
						knownvalue.StringExact("/gtm"),
					),
					statecheck.ExpectKnownValue(
						dataSourceName,
						tfjsonpath.New("hide_original_ip"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						dataSourceName,
						tfjsonpath.New("measurement_id"),
						knownvalue.StringExact("GTM-XXXXXXX"),
					),
					statecheck.ExpectKnownValue(
						dataSourceName,
						tfjsonpath.New("set_up_tag"),
						knownvalue.Bool(true),
					),
				},
			},
		},
	})
}

func testAccGoogleTagGatewayDataSourceBasicConfig(zoneID, name string) string {
	return acctest.LoadTestCase("datasource_basic.tf", zoneID, name)
}
