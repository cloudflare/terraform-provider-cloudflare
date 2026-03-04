package ai_search_token_test

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
	resource.AddTestSweepers("cloudflare_ai_search_token", &resource.Sweeper{
		Name: "cloudflare_ai_search_token",
		F:    testSweepCloudflareAISearchTokens,
	})
}

func testSweepCloudflareAISearchTokens(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping AI Search token sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	tokens, err := client.AISearch.Tokens.List(ctx, ai_search.TokenListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch AI Search tokens: %s", err))
		return fmt.Errorf("failed to fetch AI Search tokens: %w", err)
	}

	if len(tokens.Result) == 0 {
		tflog.Info(ctx, "No AI Search tokens to sweep")
		return nil
	}

	for _, token := range tokens.Result {
		if !utils.ShouldSweepResource(token.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting AI Search token: %s (account: %s)", token.Name, accountID))
		_, err := client.AISearch.Tokens.Delete(ctx, token.ID, ai_search.TokenDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete AI Search token %s: %s", token.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted AI Search token: %s", token.ID))
	}

	return nil
}

func TestAccCloudflareAISearchToken_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_ai_search_token." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAISearchTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("aisearchtokenbasic.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"cf_api_key"},
			},
		},
	})
}

func testAccCheckCloudflareAISearchTokenDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_ai_search_token" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.AISearch.Tokens.Read(
			context.Background(),
			rs.Primary.ID,
			ai_search.TokenReadParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("AI Search token still exists")
		}
	}

	return nil
}
