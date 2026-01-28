package authenticated_origin_pulls_settings_test

import (
	"context"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_authenticated_origin_pulls_settings", &resource.Sweeper{
		Name: "cloudflare_authenticated_origin_pulls_settings",
		F:    testSweepCloudflareAuthenticatedOriginPullsSettings,
	})
}

func testSweepCloudflareAuthenticatedOriginPullsSettings(r string) error {
	ctx := context.Background()
	// Authenticated Origin Pulls Settings is a zone-level configuration setting.
	// It's a singleton setting per zone, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Authenticated Origin Pulls Settings doesn't require sweeping (zone setting)")
	return nil
}

func testAccAuthenticatedOriginPullsSettingsEnabled(zoneID, rnd string) string {
	return acctest.LoadTestCase("authenticatedoriginpullssettings.tf", rnd, zoneID)
}

func testAccAuthenticatedOriginPullsSettingsDisabled(zoneID, rnd string) string {
	return acctest.LoadTestCase("authenticatedoriginpullssettings_disabled.tf", rnd, zoneID)
}

func TestAccAuthenticatedOriginPullsSettings_FullLifecycle(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_authenticated_origin_pulls_settings." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with enabled = true
			{
				Config: testAccAuthenticatedOriginPullsSettingsEnabled(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				},
			},
			// Step 2: Drift check - same config, expect empty plan
			{
				Config: testAccAuthenticatedOriginPullsSettingsEnabled(zoneID, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 3: Update to enabled = false
			{
				Config: testAccAuthenticatedOriginPullsSettingsDisabled(zoneID, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
				},
			},
			// Step 4: Drift check after update
			{
				Config: testAccAuthenticatedOriginPullsSettingsDisabled(zoneID, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 5: Update back to enabled = true (confirms toggle works both ways)
			{
				Config: testAccAuthenticatedOriginPullsSettingsEnabled(zoneID, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				},
			},
		},
	})
}
