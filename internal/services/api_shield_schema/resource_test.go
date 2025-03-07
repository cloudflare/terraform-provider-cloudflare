package api_shield_schema_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAPIShieldSchema_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_api_shield_schema.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAPIShieldSchemaBasic(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
				),
			},
		},
	})
}

func testAccCheckCloudflareAPIShieldSchemaBasic(zoneID, name string) string {
	return acctest.LoadTestCase("basic.tf", zoneID, name)
}
