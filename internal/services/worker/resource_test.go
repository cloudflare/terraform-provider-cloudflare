package worker_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
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
		return nil
	}

	list, err := client.Workers.Beta.Workers.List(ctx, workers.BetaWorkerListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		return fmt.Errorf("failed to list workers: %w", err)
	}

	for list != nil && len(list.Result) > 0 {
		for _, worker := range list.Result {
			if !strings.HasPrefix(worker.ID, resourcePrefix) {
				continue
			}

			_, err := client.Workers.Beta.Workers.Delete(ctx, worker.ID, workers.BetaWorkerDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				continue
			}
		}

		list, err = list.GetNextPage()
		if err != nil {
			break
		}
	}

	return nil
}

func TestAccCloudflareWorker_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_worker." + rnd
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
