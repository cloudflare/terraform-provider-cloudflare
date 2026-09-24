package zone_lockdown_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zone_lockdown", &resource.Sweeper{
		Name: "cloudflare_zone_lockdown",
		F:    testSweepCloudflareZoneLockdowns,
	})
}

func testSweepCloudflareZoneLockdowns(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping zone lockdowns sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	lockdowns, _, err := client.ListZoneLockdowns(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.LockdownListParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch zone lockdowns: %s", err))
		return fmt.Errorf("failed to fetch zone lockdowns: %w", err)
	}

	if len(lockdowns) == 0 {
		tflog.Info(ctx, "No zone lockdowns to sweep")
		return nil
	}

	for _, lockdown := range lockdowns {
		// Use standard filtering helper on the description field
		if !utils.ShouldSweepResource(lockdown.Description) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting zone lockdown: %s (zone: %s)", lockdown.ID, zoneID))
		_, err := client.DeleteZoneLockdown(ctx, cloudflare.ZoneIdentifier(zoneID), lockdown.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete zone lockdown %s: %s", lockdown.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted zone lockdown: %s", lockdown.ID))
	}

	return nil
}

func TestAccCloudflareZoneLockdown(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_lockdown." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneLockdownConfig(rnd, zoneID, "false", "1", "this is notes", rnd+"."+zoneName+"/*", "ip", "198.51.100.4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "urls.#", "1"),
					resource.TestCheckResourceAttr(name, "configurations.#", "1"),
				),
			},
			// Mutate only `urls`, which is the one configurable attribute that is
			// not RequiresReplace, so this exercises Update in-place rather than a
			// destroy/create. This is the path that surfaces the `created_on`
			// plan-vs-apply inconsistency.
			{
				Config: testCloudflareZoneLockdownConfig(rnd, zoneID, "false", "1", "this is notes", rnd+"."+zoneName+"/updated/*", "ip", "198.51.100.4"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "urls.#", "1"),
					// `urls` is a set, so elements are hash-indexed rather than
					// positional; `urls.0` is not a valid key.
					resource.TestCheckTypeSetElemAttr(name, "urls.*", rnd+"."+zoneName+"/updated/*"),
					resource.TestCheckResourceAttr(name, "configurations.#", "1"),
					resource.TestCheckResourceAttrSet(name, "created_on"),
				),
			},
		},
	})
}

// Regression test for ESCALATION-10665 / upstream GH #7097.
//
// `urls` used to be modelled as an ordered list while the API returns the array
// in a server-assigned order. Apply left state in config order, refresh
// overwrote it with the API order, and the next plan diffed the two — a
// permanent in-place update on every apply. That update is also the path that
// surfaced the `created_on` post-apply consistency error.
//
// `urls` is now a set, so membership alone determines equality. Both steps
// below fail if it is ever changed back to a list.
func TestAccCloudflareZoneLockdown_URLsOrderInsensitive(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_lockdown." + rnd

	alpha := rnd + "." + zoneName + "/alpha/*"
	bravo := rnd + "." + zoneName + "/bravo/*"
	charlie := rnd + "." + zoneName + "/charlie/*"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneLockdownMultiURLConfig(rnd, zoneID, alpha, bravo, charlie),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "urls.#", "3"),
					resource.TestCheckTypeSetElemAttr(name, "urls.*", alpha),
					resource.TestCheckTypeSetElemAttr(name, "urls.*", bravo),
					resource.TestCheckTypeSetElemAttr(name, "urls.*", charlie),
				),
			},
			// Refresh against the live API and re-plan. This is the step that
			// caught the original drift: Read adopts whatever order the API
			// returns, which for a list disagreed with the stored order.
			{
				RefreshState: true,
				RefreshPlanChecks: resource.RefreshPlanChecks{
					PostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Same membership, different order in config. For a set this is not a
			// change at all, so the plan must be empty. As a list it planned an
			// in-place update.
			{
				Config:   testCloudflareZoneLockdownMultiURLConfig(rnd, zoneID, charlie, alpha, bravo),
				PlanOnly: true,
			},
		},
	})
}

// test creating a config with only the required fields.
func TestAccCloudflareZoneLockdown_OnlyRequired(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_lockdown." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneLockdownConfig(rnd, zoneID, "false", "1", "this is notes", rnd+"."+zoneName+"/*", "ip", "198.51.100.4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "urls.#", "1"),
					resource.TestCheckResourceAttr(name, "configurations.#", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareZoneLockdown_Import(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	// name := "cloudflare_zone_lockdown." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneLockdownConfig(rnd, zoneID, "false", "1", "this is notes", rnd+"."+zoneName+"/*", "ip", "198.51.100.4"),
			},
			// {
			// 	ResourceName:        name,
			// 	ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
			// 	ImportState:         true,
			// 	ImportStateVerify:   true,
			// },
		},
	})
}

func testCloudflareZoneLockdownConfig(resourceID, zoneID, paused, priority, description, url, target, value string) string {
	return acctest.LoadTestCase("cloudflarezonelockdownconfig.tf", resourceID, zoneID, paused, priority, description, url, target, value)
}

func testCloudflareZoneLockdownMultiURLConfig(resourceID, zoneID, url1, url2, url3 string) string {
	return acctest.LoadTestCase("cloudflarezonelockdownmultiurl.tf", resourceID, zoneID, url1, url2, url3)
}

func TestAccUpgradeZoneLockdown_FromPublishedV5(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()

	config := testCloudflareZoneLockdownConfig(rnd, zoneID, "false", "1", "this is notes", rnd+"."+zoneName+"/*", "ip", "198.51.100.4")

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.16.0",
					},
				},
				Config: config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}
