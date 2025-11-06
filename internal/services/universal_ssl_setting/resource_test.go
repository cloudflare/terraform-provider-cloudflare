package universal_ssl_setting_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testUniversalSSLSetting_Enable(rnd, zoneID string) string {
	return acctest.LoadTestCase("universalssl_enabled.tf", rnd, zoneID)
}

func testUniversalSSLSetting_Disable(rnd, zoneID string) string {
	return acctest.LoadTestCase("universalssl_disabled.tf", rnd, zoneID)
}

func TestAccCloudflareUniversalSSLSetting_Enable(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_universal_ssl_setting." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testUniversalSSLSetting_Enable(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareUniversalSSLSetting_Disable(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_universal_ssl_setting." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testUniversalSSLSetting_Disable(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareUniversalSSLSetting_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_universal_ssl_setting." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testUniversalSSLSetting_Enable(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
			{
				Config: testUniversalSSLSetting_Disable(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
				),
			},
			{
				Config: testUniversalSSLSetting_Enable(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
		},
	})
}
