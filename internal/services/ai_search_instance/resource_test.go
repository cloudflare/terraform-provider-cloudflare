package ai_search_instance_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/ai_search"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const defaultNamespace = "default"

// no_refresh fields that the API does not return on Read.
var importStateVerifyIgnore = []string{
	"chunk",
	"summarization",
	"summarization_model",
	"system_prompt_aisearch",
	"system_prompt_index_summarization",
	"system_prompt_rewrite_query",
}

// Appended to an instance config; reads the instance back through both data sources.
const dataSourcesConfig = `
data "cloudflare_ai_search_instance" "%[1]s" {
  account_id = "%[2]s"
  namespace  = %[3]s
  id         = cloudflare_ai_search_instance.%[1]s.id
}

data "cloudflare_ai_search_instances" "%[1]s" {
  account_id = "%[2]s"
  namespace  = %[3]s
  depends_on = [cloudflare_ai_search_instance.%[1]s]
}
`

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_ai_search_instance", &resource.Sweeper{
		Name: "cloudflare_ai_search_instance",
		F:    testSweepCloudflareAISearchInstances,
	})
}

// testClient reads credentials and CLOUDFLARE_BASE_URL from the environment, so token-only credentials work.
func testClient() *cloudflare.Client {
	return cloudflare.NewClient()
}

// testSweepCloudflareAISearchInstances deletes test instances in every namespace, including "default".
func testSweepCloudflareAISearchInstances(r string) error {
	ctx := context.Background()
	client := testClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return nil
	}

	namespaces := []string{defaultNamespace}
	nsPager := client.AISearch.Namespaces.ListAutoPaging(ctx, ai_search.NamespaceListParams{AccountID: cloudflare.F(accountID)})
	for nsPager.Next() {
		if name := nsPager.Current().Name; name != defaultNamespace {
			namespaces = append(namespaces, name)
		}
	}
	if err := nsPager.Err(); err != nil {
		return fmt.Errorf("failed to fetch AI Search namespaces: %w", err)
	}

	for _, ns := range namespaces {
		pager := client.AISearch.Namespaces.Instances.ListAutoPaging(ctx, ns, ai_search.NamespaceInstanceListParams{AccountID: cloudflare.F(accountID)})
		var ids []string
		for pager.Next() {
			if id := pager.Current().ID; utils.ShouldSweepResource(id) {
				ids = append(ids, id)
			}
		}
		for _, id := range ids {
			_, _ = client.AISearch.Namespaces.Instances.Delete(ctx, ns, id, ai_search.NamespaceInstanceDeleteParams{AccountID: cloudflare.F(accountID)})
		}
	}
	return nil
}

// R2-backed instance in the default namespace.
func TestAccCloudflareAISearchInstance_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_ai_search_instance." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAISearchInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("aisearchinstancebasic.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("r2")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("token_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("source"), knownvalue.StringExact(rnd)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/%s/", accountID, defaultNamespace),
				ImportStateVerifyIgnore: importStateVerifyIgnore,
			},
		},
	})
}

// Instance without `namespace` lands in "default", converges, and imports with the three-part ID.
func TestAccCloudflareAISearchInstance_DefaultNamespace(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_ai_search_instance." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	config := acctest.LoadTestCase("aisearchinstancedefaultnamespace.tf", rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAISearchInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("namespace"), knownvalue.StringExact(defaultNamespace)),
				},
			},
			{
				Config:           config,
				ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/%s/", accountID, defaultNamespace),
				ImportStateVerifyIgnore: importStateVerifyIgnore,
			},
		},
	})
}

// Instance in a custom namespace: converges, imports, updates in place, and is replaced when moved to another namespace.
func TestAccCloudflareAISearchInstance_CustomNamespace(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_ai_search_instance." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	nsA, nsB := rnd+"-a", rnd+"-b"
	config := func(ns string, maxNumResults int) string {
		return acctest.LoadTestCase("aisearchinstancecustomnamespace.tf", rnd, accountID, nsA, nsB, ns, maxNumResults)
	}
	emptyPlan := resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAISearchInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: config("a", 10),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("namespace"), knownvalue.StringExact(nsA)),
				},
			},
			{Config: config("a", 10), ConfigPlanChecks: emptyPlan},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/%s/", accountID, nsA),
				ImportStateVerifyIgnore: importStateVerifyIgnore,
			},
			{
				Config: config("a", 20),
				ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
				}},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("max_num_results"), knownvalue.Int64Exact(20)),
				},
			},
			{
				Config: config("b", 20),
				ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionReplace),
				}},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("namespace"), knownvalue.StringExact(nsB)),
				},
			},
			{Config: config("b", 20), ConfigPlanChecks: emptyPlan},
		},
	})
}

// Data sources with an explicit `namespace = "default"`; data sources have no namespace default.
func TestAccCloudflareAISearchInstanceDataSource_DefaultNamespace(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	config := acctest.LoadTestCase("aisearchinstancedefaultnamespace.tf", rnd, accountID) +
		fmt.Sprintf(dataSourcesConfig, rnd, accountID, `"default"`)
	testAccAISearchInstanceDataSources(t, rnd, config, defaultNamespace)
}

// Data sources with an explicit non-default `namespace`.
func TestAccCloudflareAISearchInstanceDataSource_CustomNamespace(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	nsA := rnd + "-a"
	config := acctest.LoadTestCase("aisearchinstancecustomnamespace.tf", rnd, accountID, nsA, rnd+"-b", "a", 10) +
		fmt.Sprintf(dataSourcesConfig, rnd, accountID, "cloudflare_ai_search_instance."+rnd+".namespace")
	testAccAISearchInstanceDataSources(t, rnd, config, nsA)
}

func testAccAISearchInstanceDataSources(t *testing.T, rnd, config, namespace string) {
	dataSourceName := "data.cloudflare_ai_search_instance." + rnd
	listDataSourceName := "data.cloudflare_ai_search_instances." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAISearchInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("id"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("namespace"), knownvalue.StringExact(namespace)),
					statecheck.ExpectKnownValue(listDataSourceName, tfjsonpath.New("namespace"), knownvalue.StringExact(namespace)),
				},
				Check: resource.TestCheckTypeSetElemNestedAttrs(listDataSourceName, "result.*", map[string]string{"id": rnd}),
			},
		},
	})
}

// Instance created by a published provider (legacy endpoint) is not replaced and converges.
// Pinned to 5.24.0: v5.25.0-v5.28.1 declare vectorize_name in the model but not the schema, so their Create fails.
func TestAccUpgradeAISearchInstance_FromPublishedV5(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_ai_search_instance." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	config := acctest.LoadTestCase("aisearchinstancedefaultnamespace.tf", rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: testAccCheckCloudflareAISearchInstanceDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.24.0"},
				},
				Config: config,
				// 5.24.0 itself drifts after refresh; strict checks run in step 2.
				ExpectNonEmptyPlan: true,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply:             []plancheck.PlanCheck{expectNoReplace(resourceName)},
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("namespace"), knownvalue.StringExact(defaultNamespace)),
				},
			},
		},
	})
}

func testAccCheckCloudflareAISearchInstanceDestroy(s *terraform.State) error {
	client := testClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_ai_search_instance" {
			continue
		}
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = defaultNamespace
		}
		// The API can serve stale reads for a few seconds after a delete, so poll until 404.
		err := retry.RetryContext(context.Background(), 60*time.Second, func() *retry.RetryError {
			_, err := client.AISearch.Namespaces.Instances.Read(context.Background(), namespace, rs.Primary.ID, ai_search.NamespaceInstanceReadParams{
				AccountID: cloudflare.F(rs.Primary.Attributes[consts.AccountIDSchemaKey]),
			})
			var apiErr *cloudflare.Error
			switch {
			case err == nil:
				return retry.RetryableError(fmt.Errorf("AI Search instance %s/%s still exists", namespace, rs.Primary.ID))
			case errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound:
				return nil
			default:
				return retry.NonRetryableError(err)
			}
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// expectNoReplace fails if the plan destroys or replaces the resource; updates and no-ops are allowed.
type expectNoReplace string

func (e expectNoReplace) CheckPlan(_ context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	for _, rc := range req.Plan.ResourceChanges {
		if rc.Address == string(e) && (rc.Change.Actions.Replace() || rc.Change.Actions.Delete()) {
			resp.Error = fmt.Errorf("%s: expected no destroy/replace, got %v", e, rc.Change.Actions)
		}
	}
}
