package zone_setting_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zones"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareZoneSetting_OnOff(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigOnOff(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("http3")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fmt.Sprintf("%s/http3", zoneID),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_HTTP3(t *testing.T) {

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigHTTP3(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("http3")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fmt.Sprintf("%s/http3", zoneID),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_Number(t *testing.T) {

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigNumber(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("browser_cache_ttl")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.Int64Exact(30)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fmt.Sprintf("%s/browser_cache_ttl", zoneID),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_NEL(t *testing.T) {

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigNEL(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("nel")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value").AtMapKey("enabled"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fmt.Sprintf("%s/nel", zoneID),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_SSLRecommender(t *testing.T) {
	t.Skip("pending fixing schema to be nested")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigSSLRecommender(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("ssl_recommender")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fmt.Sprintf("%s/ssl_recommender", zoneID),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_HSTS(t *testing.T) {

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigHSTS(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("security_header")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("include_subdomains"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("max_age"), knownvalue.Int64Exact(30)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("nosniff"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("preload"), knownvalue.Bool(false)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fmt.Sprintf("%s/security_header", zoneID),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_MinTLSVersion(t *testing.T) {

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigMinTLSVersion(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("min_tls_version")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("1.2")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fmt.Sprintf("%s/min_tls_version", zoneID),
			},
		},
	})
}

func TestAccCloudflareZoneSetting_Ciphers(t *testing.T) {

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneSettingConfigCiphersEmpty(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("ciphers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.ListSizeExact(0)),
				},
			},
			// This will cause the panic due to #6363
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fmt.Sprintf("%s/ciphers", zoneID),
			},
			{
				Config: testCloudflareZoneSettingConfigCiphers(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("ciphers")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.StringExact("ECDHE-ECDSA-AES128-GCM-SHA256"),
						knownvalue.StringExact("ECDHE-ECDSA-CHACHA20-POLY1305"),
					})),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fmt.Sprintf("%s/ciphers", zoneID),
			},
		},
	})
}

// Regression test for https://github.com/cloudflare/terraform-provider-cloudflare/issues/5795
// where certain zone settings have inconsistent "editable" values between plan and apply
func TestAccCloudflareZoneSetting_EditableInconsistency(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	// Test the problematic settings that have editable inconsistency issues
	problematicSettings := []struct {
		settingID string
		value     string
	}{
		{"advanced_ddos", "on"},
		{"http2", "on"},
		{"long_lived_grpc", "on"},
		{"origin_error_page_pass_thru", "on"},
		{"prefetch_preload", "on"},
		{"proxy_read_timeout", "300"},
		{"response_buffering", "on"},
		{"sort_query_string_for_cache", "on"},
		{"true_client_ip_header", "on"},
	}

	for _, setting := range problematicSettings {
		t.Run(setting.settingID, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := fmt.Sprintf("cloudflare_zone_setting.%s", rnd)

			valueCheck := knownvalue.StringExact(setting.value)

			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { acctest.TestAccPreCheck(t) },
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				CheckDestroy:             testAccCheckCloudflareZoneSettingDestroy,
				Steps: []resource.TestStep{
					{
						Config: testCloudflareZoneSettingEditableInconsistency(rnd, zoneID, setting.settingID, setting.value),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact(setting.settingID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), valueCheck),
						},
					},
				},
			})
		})
	}
}

func testCloudflareZoneSettingConfigOnOff(resourceID, zoneID string) string {
	return acctest.LoadTestCase("on_off.tf", resourceID, zoneID)
}

func testCloudflareZoneSettingConfigHTTP3(resourceID, zoneID string) string {
	return acctest.LoadTestCase("http3.tf", resourceID, zoneID)
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

func testCloudflareZoneSettingConfigCiphersEmpty(resourceID, zoneID string) string {
	return acctest.LoadTestCase("ciphers_empty.tf", resourceID, zoneID)
}

func testCloudflareZoneSettingEditableInconsistency(resourceID, zoneID, settingID, value string) string {
	return acctest.LoadTestCase("editable_inconsistency.tf", resourceID, zoneID, settingID, value)
}

func testAccCheckCloudflareZoneSettingDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zone_setting" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		settingID := rs.Primary.Attributes["setting_id"]

		// Zone settings cannot be destroyed, they always exist with default values
		// We verify that we can still fetch the setting after "deletion"
		_, err := client.Zones.Settings.Get(context.Background(), settingID, zones.SettingGetParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			return fmt.Errorf("error fetching zone setting after deletion: %v", err)
		}
	}

	return nil
}
