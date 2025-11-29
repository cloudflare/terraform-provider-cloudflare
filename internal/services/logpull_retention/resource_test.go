package logpull_retention_test

import (
	"context"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
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
	resource.AddTestSweepers("cloudflare_logpull_retention", &resource.Sweeper{
		Name: "cloudflare_logpull_retention",
		F:    testSweepCloudflareLogpullRetention,
	})
}

func testSweepCloudflareLogpullRetention(r string) error {
	ctx := context.Background()
	// Logpull Retention is a zone-level configuration setting (enabled/disabled).
	// It's a singleton setting per zone, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Logpull Retention doesn't require sweeping (zone setting)")
	return nil
}

func TestAccLogpullRetention_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Logpull
	// service is throwing authentication errors despite it being marked as
	// available.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_logpull_retention." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Since prior state is not guaranteed, no plancheck.ExpectResourceAction() on the first step.
			// However, it has extra step to ensure the update cases (set to true -> false -> true).
			// Set flag to true.
			{
				Config: testLogpullRetentionSetConfig(rnd, zoneID, true),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// No plancheck.ExpectResourceAction().
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("flag"), knownvalue.Bool(true)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("flag"), knownvalue.Bool(true)),
				},
			},
			// Set flag to false.
			{
				Config: testLogpullRetentionSetConfig(rnd, zoneID, false),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("flag"), knownvalue.Bool(false)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("flag"), knownvalue.Bool(false)),
				},
			},
			// Set flag to true.
			{
				Config: testLogpullRetentionSetConfig(rnd, zoneID, true),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("flag"), knownvalue.Bool(true)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("flag"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testLogpullRetentionSetConfig(id, zoneID string, enabled bool) string {
	return acctest.LoadTestCase("logpullretentionsetconfig.tf", id, zoneID, enabled)
}
