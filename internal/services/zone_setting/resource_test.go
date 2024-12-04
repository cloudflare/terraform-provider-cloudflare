package zone_setting_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareZoneSetting_OnOff(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_setting." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigOnOff(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_Boolean(t *testing.T) {
	t.Skip("pending finding a setting to test against")

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_setting." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigBoolean(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_Number(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_setting." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigNumber(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_NEL(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_setting." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigNEL(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_SSLRecommender(t *testing.T) {
	t.Skip("pending fixing schema to be nested")

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_setting." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigSSLRecommender(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_HSTS(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_setting." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigHSTS(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_MinTLSVersion(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_setting." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigMinTLSVersion(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_Ciphers(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_setting." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigCiphers(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}

func testCloudflareZoneSettingConfigOnOff(resourceID, zoneID string) string {
	return acctest.LoadTestCase("on_off.tf", resourceID, zoneID)
}

func testCloudflareZoneSettingConfigBoolean(resourceID, zoneID string) string {
	return acctest.LoadTestCase("boolean.tf", resourceID, zoneID)
}

func testCloudflareZoneSettingConfigNumber(resourceID, zoneID string) string {
	return acctest.LoadTestCase("number.tf", resourceID, zoneID)
}

func testCloudflareZoneSettingConfigNEL(resourceID, zoneID string) string {
	return acctest.LoadTestCase("nel.tf", resourceID, zoneID)
}

func testCloudflareZoneSettingConfigSSLRecommender(resourceID, zoneID string) string {
	return acctest.LoadTestCase("ssl_recommender.tf", resourceID, zoneID)
}

func testCloudflareZoneSettingConfigHSTS(resourceID, zoneID string) string {
	return acctest.LoadTestCase("hsts.tf", resourceID, zoneID)
}

func testCloudflareZoneSettingConfigMinTLSVersion(resourceID, zoneID string) string {
	return acctest.LoadTestCase("min_tls_version.tf", resourceID, zoneID)
}

func testCloudflareZoneSettingConfigCiphers(resourceID, zoneID string) string {
	return acctest.LoadTestCase("ciphers.tf", resourceID, zoneID)
}
