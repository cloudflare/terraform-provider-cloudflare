package zero_trust_dex_rule_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareDexRule(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_dex_rule.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	match := `group_policy_id == \"bd2f8f7b-d0b0-4d24-ba2b-bdfdf244d9f9\"`
	matchEscaped := "group_policy_id == \"bd2f8f7b-d0b0-4d24-ba2b-bdfdf244d9f9\""

	sharedChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
		resource.TestCheckResourceAttr(name, "name", rnd),
		resource.TestCheckResourceAttr(name, "match", matchEscaped),
	}
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create rule
			{
				Config: testAccCloudflareDexRule(accountID, rnd, match, "description"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionCreate),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(sharedChecks...),
			},
			// Update rule
			{
				Config: testAccCloudflareDexRule(accountID, rnd, match, "description updated"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("description updated")),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					append(
						sharedChecks,
						resource.TestCheckResourceAttr(name, "description", "description updated"),
					)...,
				),
			},
			// import resource
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCloudflareDexRule(accountID, rnd, match, description string) string {
	return acctest.LoadTestCase("dexrulestest.tf", rnd, accountID, match, description)
}
