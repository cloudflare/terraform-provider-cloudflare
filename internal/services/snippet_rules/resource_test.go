package snippet_rules_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/snippets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func init() {
	resource.AddTestSweepers("cloudflare_snippet_rules", &resource.Sweeper{
		Name: "cloudflare_snippet_rules",
		F:    testSweepCloudflareSnippetRules,
	})
}

func testSweepCloudflareSnippetRules(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	if zoneID == "" {
		// Skip sweeping if no zone ID is set
		return nil
	}

	// Update to empty rules list to delete all rules
	_, err := client.Snippets.Rules.Update(ctx, snippets.RuleUpdateParams{
		ZoneID: cloudflare.F(zoneID),
		Rules:  cloudflare.F([]snippets.RuleUpdateParamsRule{}),
	})
	if err != nil {
		return fmt.Errorf("failed to delete snippet rules: %w", err)
	}

	return nil
}

func TestAccCloudflareSnippetRules_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_snippet_rules." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareSnippetRulesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSnippetRulesConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("snippet_name"), knownvalue.StringExact("rules_set_snippet")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("http.request.uri.path contains \"/test\"")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Test snippet rule")),
				},
			},
			{
				Config: testAccCloudflareSnippetRulesConfigUpdate(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("snippet_name"), knownvalue.StringExact("rules_set_snippet")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("http.request.uri.path contains \"/updated\"")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Updated test snippet rule")),
				},
			},
			// NOTE: we don't test importing b/c resource import is not implemented
		},
	})
}

func testAccCheckCloudflareSnippetRulesDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_snippet_rules" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]

		rules, err := client.Snippets.Rules.List(context.Background(), snippets.RuleListParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			return fmt.Errorf("error getting snippet rules: %w", err)
		}

		// Check if any rules still exist
		if rules != nil && len(rules.Result) > 0 {
			return fmt.Errorf("snippet rules still exist")
		}
	}

	return nil
}

func testAccCloudflareSnippetRulesConfig(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, zoneID)
}

func testAccCloudflareSnippetRulesConfigUpdate(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic_update.tf", rnd, zoneID)
}
