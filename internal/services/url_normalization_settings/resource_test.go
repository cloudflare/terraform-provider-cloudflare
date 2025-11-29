package url_normalization_settings_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6/url_normalization"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	cloudflare "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_url_normalization_settings", &resource.Sweeper{
		Name: "cloudflare_url_normalization_settings",
		F:    testSweepCloudflareURLNormalizationSettings,
	})
}

func testSweepCloudflareURLNormalizationSettings(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping URL normalization settings sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	settings, err := client.URLNormalization.Get(ctx, url_normalization.URLNormalizationGetParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch URL normalization settings: %s", err))
		return err
	}

	if settings == nil {
		tflog.Info(ctx, "No URL normalization settings to sweep")
		return nil
	}

	tflog.Info(ctx, fmt.Sprintf("Deleting URL normalization settings (zone: %s)", zoneID))
	err = client.URLNormalization.Delete(ctx, url_normalization.URLNormalizationDeleteParams{
		ZoneID: cloudflare.F(zoneID),
	})

	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to delete URL normalization settings: %s", err))
		return err
	}

	tflog.Info(ctx, "Deleted URL normalization settings")
	return nil
}

func TestAccCloudflareURLNormalizationSettings_CreateThenUpdate(t *testing.T) {

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_url_normalization_settings.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "cloudflare", "incoming", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("cloudflare")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("incoming")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     zoneID,
			},
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "cloudflare", "both", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("cloudflare")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("both")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     zoneID,
			},
		},
	})
}

func TestAccCloudflareURLNormalizationSettings_AllCombinations(t *testing.T) {

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_url_normalization_settings.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,

		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "cloudflare", "incoming", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("cloudflare")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("incoming")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     zoneID,
			},
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "cloudflare", "both", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("cloudflare")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("both")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     zoneID,
			},
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "rfc3986", "incoming", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("rfc3986")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("incoming")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     zoneID,
			},
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "rfc3986", "both", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("rfc3986")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("both")),
				},
			},
		},
	})
}

func TestAccCloudflareURLNormalizationSettings_InvalidValues(t *testing.T) {

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "invalid_type", "incoming", rnd),
				ExpectError: regexp.MustCompile(`value must be one of: \["cloudflare" "rfc3986"\]`),
			},
			{
				Config:      testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "cloudflare", "invalid_scope", rnd),
				ExpectError: regexp.MustCompile(`value must be one of: \["incoming"|"none"|"both"\]`),
			},
		},
	})
}

func TestAccCloudflareURLNormalizationSettings_CreateThenDelete(t *testing.T) {

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_url_normalization_settings.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "cloudflare", "incoming", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("cloudflare")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("incoming")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     zoneID,
			},
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfigEmpty(zoneID, rnd),
			},
		},
	})
}

func TestAccCloudflareURLNormalizationSettings_NoneScope(t *testing.T) {

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_url_normalization_settings.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "cloudflare", "none", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("cloudflare")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("none")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     zoneID,
			},
		},
	})
}

func testAccCheckCloudflareURLNormalizationSettingsConfigEmpty(zoneID, name string) string {
	return acctest.LoadTestCase("urlnormalizationsettingsconfig_empty.tf", zoneID, name)
}

func testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, _type, scope, name string) string {
	return acctest.LoadTestCase("urlnormalizationsettingsconfig.tf", zoneID, _type, scope, name)
}
