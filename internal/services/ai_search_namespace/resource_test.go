package ai_search_namespace_test

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

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_ai_search_namespace", &resource.Sweeper{
		Name: "cloudflare_ai_search_namespace",
		F:    testSweepCloudflareAISearchNamespaces,
	})
}

// testClient reads credentials and CLOUDFLARE_BASE_URL from the environment, so token-only credentials work.
func testClient() *cloudflare.Client {
	return cloudflare.NewClient()
}

// testSweepCloudflareAISearchNamespaces empties and deletes test namespaces; "default" is never touched.
func testSweepCloudflareAISearchNamespaces(r string) error {
	ctx := context.Background()
	client := testClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return nil
	}

	var names []string
	pager := client.AISearch.Namespaces.ListAutoPaging(ctx, ai_search.NamespaceListParams{AccountID: cloudflare.F(accountID)})
	for pager.Next() {
		if name := pager.Current().Name; name != "default" && utils.ShouldSweepResource(name) {
			names = append(names, name)
		}
	}
	if err := pager.Err(); err != nil {
		return fmt.Errorf("failed to fetch AI Search namespaces: %w", err)
	}

	for _, name := range names {
		instances := client.AISearch.Namespaces.Instances.ListAutoPaging(ctx, name, ai_search.NamespaceInstanceListParams{AccountID: cloudflare.F(accountID)})
		var ids []string
		for instances.Next() {
			ids = append(ids, instances.Current().ID)
		}
		for _, id := range ids {
			_, _ = client.AISearch.Namespaces.Instances.Delete(ctx, name, id, ai_search.NamespaceInstanceDeleteParams{AccountID: cloudflare.F(accountID)})
		}
		// Instance deletes can take a few seconds to show, so the namespace may briefly still look non-empty.
		_ = retry.RetryContext(ctx, 30*time.Second, func() *retry.RetryError {
			_, err := client.AISearch.Namespaces.Delete(ctx, name, ai_search.NamespaceDeleteParams{AccountID: cloudflare.F(accountID)})
			if err == nil || isNotFound(err) {
				return nil
			}
			return retry.RetryableError(err)
		})
	}
	return nil
}

func testAccAISearchNamespaceConfig(rnd, accountID, description string) string {
	return fmt.Sprintf(`
resource "cloudflare_ai_search_namespace" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = %[3]s
}
`, rnd, accountID, description)
}

// Namespace create, id == name, convergence, in-place update, and import with `<account_id>/<name>`.
func TestAccCloudflareAISearchNamespace_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_ai_search_namespace." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	created := testAccAISearchNamespaceConfig(rnd, accountID, `"created by terraform acceptance tests"`)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAISearchNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: created,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				},
			},
			{
				Config:           created,
				ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}},
			},
			{
				Config: testAccAISearchNamespaceConfig(rnd, accountID, `"updated by terraform acceptance tests"`),
				ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
				}},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("updated by terraform acceptance tests")),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// Namespace created by the newest published provider (no `id` attribute) is not replaced and converges.
func TestAccUpgradeAISearchNamespace_FromPublishedV5(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_ai_search_namespace." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	config := testAccAISearchNamespaceConfig(rnd, accountID, "null")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: testAccCheckCloudflareAISearchNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {Source: "cloudflare/cloudflare", VersionConstraint: "5.28.1"},
				},
				Config: config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply:             []plancheck.PlanCheck{expectNoReplace(resourceName)},
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(rnd)),
				},
			},
		},
	})
}

func testAccCheckCloudflareAISearchNamespaceDestroy(s *terraform.State) error {
	client := testClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_ai_search_namespace" {
			continue
		}
		name := rs.Primary.Attributes["name"]
		// The API can serve stale reads for a few seconds after a delete, so poll until 404.
		err := retry.RetryContext(context.Background(), 60*time.Second, func() *retry.RetryError {
			_, err := client.AISearch.Namespaces.Read(context.Background(), name, ai_search.NamespaceReadParams{
				AccountID: cloudflare.F(rs.Primary.Attributes[consts.AccountIDSchemaKey]),
			})
			switch {
			case err == nil:
				return retry.RetryableError(fmt.Errorf("AI Search namespace %s still exists", name))
			case isNotFound(err):
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

func isNotFound(err error) bool {
	var apiErr *cloudflare.Error
	return errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound
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
