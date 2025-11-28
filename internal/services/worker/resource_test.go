package worker_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const resourcePrefix = "tfacctest-"

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_worker", &resource.Sweeper{
		Name: "cloudflare_worker",
		F:    testSweepCloudflareWorkers,
	})
}

func testSweepCloudflareWorkers(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping workers sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	list, err := client.Workers.Beta.Workers.List(ctx, workers.BetaWorkerListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list workers: %s", err))
		return fmt.Errorf("failed to list workers: %w", err)
	}

	hasWorkers := false
	for list != nil && len(list.Result) > 0 {
		for _, worker := range list.Result {
			hasWorkers = true
			// Use standard filtering helper
			if !utils.ShouldSweepResource(worker.ID) {
				continue
			}

			tflog.Info(ctx, fmt.Sprintf("Deleting worker: %s (account: %s)", worker.ID, accountID))
			_, err := client.Workers.Beta.Workers.Delete(ctx, worker.ID, workers.BetaWorkerDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete worker %s: %s", worker.ID, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted worker: %s", worker.ID))
		}

		list, err = list.GetNextPage()
		if err != nil {
			break
		}
	}

	if !hasWorkers {
		tflog.Info(ctx, "No workers to sweep")
	}

	return nil
}

func TestAccCloudflareWorker_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_worker." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkerConfig(resourceName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(resourceName)),
					// Verify computed attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("logpush"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("tags"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("observability"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"enabled":            knownvalue.Bool(false),
						"head_sampling_rate": knownvalue.Float64Exact(1),
						"logs": knownvalue.ObjectExact(map[string]knownvalue.Check{
							"enabled":            knownvalue.Bool(false),
							"head_sampling_rate": knownvalue.Float64Exact(1),
							"invocation_logs":    knownvalue.Bool(true),
						}),
					})),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("subdomain"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"enabled":          knownvalue.Bool(false),
						"previews_enabled": knownvalue.Bool(false),
					})),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("tail_consumers"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("updated_on"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccCloudflareWorkerConfigUpdate(resourceName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(resourceName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("tags"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("environment=production")})),
					// Verify computed attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("logpush"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("observability"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"enabled":            knownvalue.Bool(false),
						"head_sampling_rate": knownvalue.Float64Exact(1),
						"logs": knownvalue.ObjectExact(map[string]knownvalue.Check{
							"enabled":            knownvalue.Bool(false),
							"head_sampling_rate": knownvalue.Float64Exact(1),
							"invocation_logs":    knownvalue.Bool(true),
						}),
					})),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("subdomain"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"enabled":          knownvalue.Bool(false),
						"previews_enabled": knownvalue.Bool(false),
					})),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("tail_consumers"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("updated_on"), knownvalue.NotNull()),
				},
			},
			// { // This fails due to having no target config for the import
			// 	// Import using an import block and assert there are no changes
			// 	// reported in the import plan.
			// 	Config:              testAccCloudflareWorkerConfigUpdate(rnd, accountID),
			// 	ResourceName:        name,
			// 	ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			// 	ImportState:         true,
			// 	ImportStateKind:     resource.ImportBlockWithID,
			// 	ImportPlanChecks: resource.ImportPlanChecks{
			// 		PreApply: []plancheck.PlanCheck{
			// 			plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
			// 		},
			// 	},
			// },
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCloudflareWorkerConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID)
}

func testAccCloudflareWorkerConfigUpdate(rnd, accountID string) string {
	return acctest.LoadTestCase("basic_update.tf", rnd, accountID)
}
