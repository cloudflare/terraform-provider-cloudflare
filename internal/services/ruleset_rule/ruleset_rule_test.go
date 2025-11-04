package ruleset_rule_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/rulesets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	accountID = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID    = os.Getenv("CLOUDFLARE_ZONE_ID")
	domain    = os.Getenv("CLOUDFLARE_DOMAIN")

	configVariables = config.Variables{
		"account_id": config.StringVariable(accountID),
		"zone_id":    config.StringVariable(zoneID),
		"domain":     config.StringVariable(domain),
	}
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_ruleset_rule", &resource.Sweeper{
		Name: "cloudflare_ruleset_rule",
		F: func(region string) error {
			ctx := context.Background()

			client := acctest.SharedClient()

			type ruleset struct {
				rulesets.RulesetListParams
				rulesets.RulesetListResponse
			}

			var entrypointRulesets, customRulesets []ruleset

			for _, params := range []rulesets.RulesetListParams{
				{AccountID: cloudflare.F(accountID)},
				{ZoneID: cloudflare.F(zoneID)},
			} {
				iter := client.Rulesets.ListAutoPaging(ctx, params)

				for iter.Next() {
					switch iter.Current().Kind {
					case rulesets.KindManaged:
					case rulesets.KindCustom:
						customRulesets = append(customRulesets, ruleset{params, iter.Current()})
					case rulesets.KindRoot:
						entrypointRulesets = append(entrypointRulesets, ruleset{params, iter.Current()})
					case rulesets.KindZone:
						entrypointRulesets = append(entrypointRulesets, ruleset{params, iter.Current()})
					default:
						return fmt.Errorf("unknown ruleset kind %q", iter.Current().Kind)
					}
				}

				if err := iter.Err(); err != nil {
					return fmt.Errorf("failed to list rulesets: %w", err)
				}
			}

			for _, ruleset := range append(entrypointRulesets, customRulesets...) {
				if err := client.Rulesets.Delete(ctx, ruleset.ID, rulesets.RulesetDeleteParams{
					AccountID: ruleset.AccountID,
					ZoneID:    ruleset.ZoneID,
				}); err != nil {
					return fmt.Errorf("failed to delete ruleset %q: %w", ruleset.ID, err)
				}
			}

			return nil
		},
	})
}

func TestAccCloudflareRulesetRule_Description(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigFile:      config.TestNameFile("1.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.block_external_traffic",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset_rule.allow_rancher",
							plancheck.ResourceActionCreate,
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset_rule.allow_rancher",
						tfjsonpath.New("description"),
						knownvalue.StringExact(""),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset_rule.allow_rancher",
						tfjsonpath.New("description"),
						knownvalue.StringExact(""),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset_rule.allow_rancher",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset_rule.allow_rancher",
							tfjsonpath.New("description"),
							knownvalue.StringExact("My rule description"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset_rule.allow_rancher",
						tfjsonpath.New("description"),
						knownvalue.StringExact("My rule description"),
					),
				},
			},
		},
	})
}
