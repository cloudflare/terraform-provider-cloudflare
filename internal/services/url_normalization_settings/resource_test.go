package url_normalization_settings_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/url_normalization"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareURLNormalizationSettings_CreateThenUpdate(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_url_normalization_settings.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareURLNormalizationSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "cloudflare", "incoming", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
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
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_url_normalization_settings.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareURLNormalizationSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, "cloudflare", "incoming", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
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

func testAccCheckCloudflareURLNormalizationSettingsDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_url_normalization_settings" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		_, err := client.URLNormalization.Get(context.Background(), url_normalization.URLNormalizationGetParams{ZoneID: cloudflare.String(zoneID)})
		// URL normalization settings always exist (they have defaults), so we can't really verify destruction
		// The best we can do is ensure no error occurred when fetching them
		if err != nil {
			return fmt.Errorf("error fetching URL normalization settings: %v", err)
		}
	}

	return nil
}

func TestAccCloudflareURLNormalizationSettings_InvalidValues(t *testing.T) {
	t.Parallel()
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
				ExpectError: regexp.MustCompile(`value must be one of: \["incoming" "both"\]`),
			},
		},
	})
}

func testAccCheckCloudflareURLNormalizationSettingsConfig(zoneID, _type, scope, name string) string {
	return acctest.LoadTestCase("urlnormalizationsettingsconfig.tf", zoneID, _type, scope, name)
}
