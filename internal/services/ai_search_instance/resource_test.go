package ai_search_instance_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ai_search"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_ai_search_instance", &resource.Sweeper{
		Name: "cloudflare_ai_search_instance",
		F:    testSweepCloudflareAISearchInstances,
	})
}

func testSweepCloudflareAISearchInstances(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping AI Search instance sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	instances, err := client.AISearch.Instances.List(ctx, ai_search.InstanceListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch AI Search instances: %s", err))
		return fmt.Errorf("failed to fetch AI Search instances: %w", err)
	}

	if len(instances.Result) == 0 {
		tflog.Info(ctx, "No AI Search instances to sweep")
		return nil
	}

	for _, instance := range instances.Result {
		if !utils.ShouldSweepResource(instance.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting AI Search instance: %s (account: %s)", instance.ID, accountID))
		_, err := client.AISearch.Instances.Delete(ctx, instance.ID, ai_search.InstanceDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete AI Search instance %s: %s", instance.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted AI Search instance: %s", instance.ID))
	}

	return nil
}

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
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{
					// no_refresh fields are not returned by the API on Read
					"chunk",
					"summarization",
					"summarization_model",
					"system_prompt_aisearch",
					"system_prompt_index_summarization",
					"system_prompt_rewrite_query",
					"index_method.keyword",
					"index_method.vector",
				},
			},
		},
	})
}

func testAccCheckCloudflareAISearchInstanceDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_ai_search_instance" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.AISearch.Instances.Read(
			context.Background(),
			rs.Primary.ID,
			ai_search.InstanceReadParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("AI Search instance still exists")
		}
	}

	return nil
}
