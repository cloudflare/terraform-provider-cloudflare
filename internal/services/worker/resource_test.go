package worker_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareWorker_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_worker." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_worker." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkerConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("logpush"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tags"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("observability"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"enabled":            knownvalue.Bool(false),
						"head_sampling_rate": knownvalue.Float64Exact(1),
						"logs": knownvalue.ObjectExact(map[string]knownvalue.Check{
							"enabled":            knownvalue.Bool(false),
							"head_sampling_rate": knownvalue.Float64Exact(1),
							"invocation_logs":    knownvalue.Bool(true),
						}),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("subdomain"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"enabled":          knownvalue.Bool(false),
						"previews_enabled": knownvalue.Bool(false),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tail_consumers"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("updated_on"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccCloudflareWorkerConfigUpdate(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tags"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("environment=production")})),
					// Verify computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("logpush"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("observability"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"enabled":            knownvalue.Bool(false),
						"head_sampling_rate": knownvalue.Float64Exact(1),
						"logs": knownvalue.ObjectExact(map[string]knownvalue.Check{
							"enabled":            knownvalue.Bool(false),
							"head_sampling_rate": knownvalue.Float64Exact(1),
							"invocation_logs":    knownvalue.Bool(true),
						}),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("subdomain"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"enabled":          knownvalue.Bool(false),
						"previews_enabled": knownvalue.Bool(false),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tail_consumers"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("updated_on"), knownvalue.NotNull()),
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

func TestAccCloudflareWorker_ImportPlanNoDiff(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_worker." + rnd
	var workerID string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)

			// Create Worker outside of Terraform that we can import.
			ctx := context.Background()
			client := acctest.SharedClient()
			res := new(http.Response)
			env := worker.WorkerResultEnvelope{}
			_, err := client.Workers.Beta.Workers.New(
				ctx,
				workers.BetaWorkerNewParams{
					AccountID: cloudflare.F(accountID),
				},
				option.WithRequestBody(
					"application/json",
					strings.NewReader(fmt.Sprintf(`{"name": "%s", "tags": ["environment=production"]}`, rnd)),
				),
				option.WithResponseBodyInto(&res),
			)
			if err != nil {
				t.Fatalf("failed to create worker to import: %v", err)
			}
			bytes, _ := io.ReadAll(res.Body)
			err = apijson.Unmarshal(bytes, &env)
			if err != nil {
				t.Fatalf("failed to unmarshal worker to import: %v", err)
			}
			workerID = env.Result.ID.ValueString()
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Import using an import block and assert there are no changes
				// reported in the import plan.
				Config:          testAccCloudflareWorkerConfigUpdate(rnd, accountID),
				ResourceName:    resourceName,
				ImportState:     true,
				ImportStateKind: resource.ImportBlockWithID,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, workerID), nil
				},
				ImportPlanChecks: resource.ImportPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
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
