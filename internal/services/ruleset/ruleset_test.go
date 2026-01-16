package ruleset_test

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
	resource.AddTestSweepers("cloudflare_ruleset", &resource.Sweeper{
		Name: "cloudflare_ruleset",
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

func TestAccCloudflareRulesets(t *testing.T) {
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
							"cloudflare_ruleset.my_first_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_second_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_third_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_fourth_ruleset",
							plancheck.ResourceActionCreate,
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"data.cloudflare_rulesets.my_account_rulesets",
						tfjsonpath.New("rulesets"),
						knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"id":           knownvalue.StringExact("3b64149bfa6e4220bbbc2bd6db589552"),
								"version":      knownvalue.NotNull(),
								"name":         knownvalue.StringExact("Cloudflare L3/4 DDoS Ruleset"),
								"description":  knownvalue.StringExact("This is the Managed Cloudflare L3/4 DDoS Ruleset. Cloudflare routinely adds signatures to address new attack vectors. Additional configuration allows you to customize the sensitivity of each rule and the performed mitigation action."),
								"phase":        knownvalue.StringExact("ddos_l4"),
								"kind":         knownvalue.StringExact("managed"),
								"last_updated": knownvalue.NotNull(),
							}),
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"id":           knownvalue.StringExact("77454fe2d30c4220b5701f6fdfb893ba"),
								"version":      knownvalue.NotNull(),
								"name":         knownvalue.StringExact("Cloudflare Managed Free Ruleset"),
								"description":  knownvalue.StringExact("Created by the Cloudflare security team, this ruleset is designed to provide protection for free zones"),
								"phase":        knownvalue.StringExact("http_request_firewall_managed"),
								"kind":         knownvalue.StringExact("managed"),
								"last_updated": knownvalue.NotNull(),
							}),
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"id":           knownvalue.NotNull(),
								"version":      knownvalue.StringExact("1"),
								"name":         knownvalue.StringExact("My first ruleset"),
								"description":  knownvalue.StringExact(""),
								"phase":        knownvalue.StringExact("http_request_firewall_custom"),
								"kind":         knownvalue.StringExact("root"),
								"last_updated": knownvalue.NotNull(),
							}),
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"id":           knownvalue.NotNull(),
								"version":      knownvalue.StringExact("1"),
								"name":         knownvalue.StringExact("My second ruleset"),
								"description":  knownvalue.StringExact(""),
								"phase":        knownvalue.StringExact("http_request_firewall_managed"),
								"kind":         knownvalue.StringExact("root"),
								"last_updated": knownvalue.NotNull(),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_rulesets.my_zone_rulesets",
						tfjsonpath.New("rulesets"),
						knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"id":           knownvalue.StringExact("70339d97bdb34195bbf054b1ebe81f76"),
								"version":      knownvalue.NotNull(),
								"name":         knownvalue.StringExact("Cloudflare Normalization Ruleset"),
								"description":  knownvalue.StringExact("Created by the Cloudflare security team, this ruleset provides normalization on the URL path"),
								"phase":        knownvalue.StringExact("http_request_sanitize"),
								"kind":         knownvalue.StringExact("managed"),
								"last_updated": knownvalue.NotNull(),
							}),
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"id":           knownvalue.StringExact("77454fe2d30c4220b5701f6fdfb893ba"),
								"version":      knownvalue.NotNull(),
								"name":         knownvalue.StringExact("Cloudflare Managed Free Ruleset"),
								"description":  knownvalue.StringExact("Created by the Cloudflare security team, this ruleset is designed to provide protection for free zones"),
								"phase":        knownvalue.StringExact("http_request_firewall_managed"),
								"kind":         knownvalue.StringExact("managed"),
								"last_updated": knownvalue.NotNull(),
							}),
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"id":           knownvalue.NotNull(),
								"version":      knownvalue.StringExact("1"),
								"name":         knownvalue.StringExact("My third ruleset"),
								"description":  knownvalue.StringExact(""),
								"phase":        knownvalue.StringExact("http_request_firewall_custom"),
								"kind":         knownvalue.StringExact("zone"),
								"last_updated": knownvalue.NotNull(),
							}),
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"id":           knownvalue.NotNull(),
								"version":      knownvalue.StringExact("1"),
								"name":         knownvalue.StringExact("My fourth ruleset"),
								"description":  knownvalue.StringExact(""),
								"phase":        knownvalue.StringExact("http_request_firewall_managed"),
								"kind":         knownvalue.StringExact("zone"),
								"last_updated": knownvalue.NotNull(),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_Kind(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("kind"),
							knownvalue.StringExact("root"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("kind"),
						knownvalue.StringExact("root"),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("kind"),
						knownvalue.StringExact("root"),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionReplace,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("kind"),
							knownvalue.StringExact("custom"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("kind"),
						knownvalue.StringExact("custom"),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("kind"),
						knownvalue.StringExact("custom"),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_Name(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("name"),
							knownvalue.StringExact("My ruleset"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("name"),
						knownvalue.StringExact("My ruleset"),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("name"),
						knownvalue.StringExact("My ruleset"),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionReplace,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("name"),
							knownvalue.StringExact("My updated ruleset"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("name"),
						knownvalue.StringExact("My updated ruleset"),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("name"),
						knownvalue.StringExact("My updated ruleset"),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_Phase(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("phase"),
							knownvalue.StringExact("http_request_firewall_custom"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("phase"),
						knownvalue.StringExact("http_request_firewall_custom"),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("phase"),
						knownvalue.StringExact("http_request_firewall_custom"),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionReplace,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("phase"),
							knownvalue.StringExact("http_request_firewall_managed"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("phase"),
						knownvalue.StringExact("http_request_firewall_managed"),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("phase"),
						knownvalue.StringExact("http_request_firewall_managed"),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_Description(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("description"),
							knownvalue.StringExact(""),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("description"),
						knownvalue.StringExact(""),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
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
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("description"),
						knownvalue.StringExact(""),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("description"),
							knownvalue.StringExact("My ruleset description"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("description"),
						knownvalue.StringExact("My ruleset description"),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("description"),
						knownvalue.StringExact("My ruleset description"),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_Rules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{}),
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
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_RulesActionParameters(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action_parameters": knownvalue.NotNull(),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action_parameters": knownvalue.NotNull(),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action_parameters": knownvalue.NotNull(),
							}),
						}),
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
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action_parameters": knownvalue.NotNull(),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action_parameters": knownvalue.NotNull(),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_RulesDescription(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"description": knownvalue.StringExact(""),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"description": knownvalue.StringExact(""),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"description": knownvalue.StringExact(""),
							}),
						}),
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
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"description": knownvalue.StringExact(""),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"description": knownvalue.StringExact(""),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"description": knownvalue.StringExact("My rule description"),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"description": knownvalue.StringExact("My rule description"),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"description": knownvalue.StringExact("My rule description"),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_RulesEnabled(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"enabled": knownvalue.Bool(true),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"enabled": knownvalue.Bool(true),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"enabled": knownvalue.Bool(true),
							}),
						}),
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
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"enabled": knownvalue.Bool(true),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"enabled": knownvalue.Bool(true),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"enabled": knownvalue.Bool(false),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"enabled": knownvalue.Bool(false),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"enabled": knownvalue.Bool(false),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_RulesExposedCredentialCheck(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"exposed_credential_check": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"username_expression": knownvalue.StringExact("url_decode(http.request.body.form[\"username\"][0])"),
										"password_expression": knownvalue.StringExact("url_decode(http.request.body.form[\"password\"][0])"),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"exposed_credential_check": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"username_expression": knownvalue.StringExact("url_decode(http.request.body.form[\"username\"][0])"),
									"password_expression": knownvalue.StringExact("url_decode(http.request.body.form[\"password\"][0])"),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"exposed_credential_check": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"username_expression": knownvalue.StringExact("url_decode(http.request.body.form[\"username\"][0])"),
									"password_expression": knownvalue.StringExact("url_decode(http.request.body.form[\"password\"][0])"),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"exposed_credential_check": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"username_expression": knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"username\")"),
										"password_expression": knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"password\")"),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"exposed_credential_check": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"username_expression": knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"username\")"),
									"password_expression": knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"password\")"),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"exposed_credential_check": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"username_expression": knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"username\")"),
									"password_expression": knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"password\")"),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_RulesLogging(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.NotNull(),
							}),
						),
						plancheck.ExpectUnknownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("logging"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"enabled": knownvalue.Bool(true),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"enabled": knownvalue.Bool(true),
								}),
							}),
						}),
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
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"enabled": knownvalue.Bool(true),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"enabled": knownvalue.Bool(true),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"enabled": knownvalue.Bool(false),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"enabled": knownvalue.Bool(false),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"enabled": knownvalue.Bool(false),
								}),
							}),
						}),
					),
				},
			},
		},
	})

	t.Run("modify", func(t *testing.T) {
		t.Run("action", func(t *testing.T) {
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
									"cloudflare_ruleset.my_ruleset",
									plancheck.ResourceActionCreate,
								),
								plancheck.ExpectKnownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules"),
									knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectPartial(map[string]knownvalue.Check{
											"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
											"action":     knownvalue.StringExact("skip"),
										}),
									}),
								),
								plancheck.ExpectUnknownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"),
								),
								plancheck.ExpectUnknownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("logging"),
								),
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"action":     knownvalue.StringExact("skip"),
										"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"enabled": knownvalue.Bool(true),
										}),
									}),
								}),
							),
							statecheck.ExpectKnownValue(
								"data.cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"action":     knownvalue.StringExact("skip"),
										"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"enabled": knownvalue.Bool(true),
										}),
									}),
								}),
							),
						},
					},
					{
						ConfigFile:      config.TestNameFile("2.tf"),
						ConfigVariables: configVariables,
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								plancheck.ExpectResourceAction(
									"cloudflare_ruleset.my_ruleset",
									plancheck.ResourceActionUpdate,
								),
								plancheck.ExpectKnownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules"),
									knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectPartial(map[string]knownvalue.Check{
											"id":         knownvalue.NotNull(),
											"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
											"action":     knownvalue.StringExact("block"),
										}),
									}),
								),
								plancheck.ExpectUnknownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("logging"),
								),
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"action":     knownvalue.StringExact("block"),
										"logging":    knownvalue.Null(),
									}),
								}),
							),
							statecheck.ExpectKnownValue(
								"data.cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"action":     knownvalue.StringExact("block"),
										"logging":    knownvalue.Null(),
									}),
								}),
							),
						},
					},
				},
			})
		})

		t.Run("expression", func(t *testing.T) {
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
									"cloudflare_ruleset.my_ruleset",
									plancheck.ResourceActionCreate,
								),
								plancheck.ExpectKnownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules"),
									knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectPartial(map[string]knownvalue.Check{
											"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
											"action":     knownvalue.StringExact("skip"),
										}),
									}),
								),
								plancheck.ExpectUnknownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"),
								),
								plancheck.ExpectUnknownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("logging"),
								),
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"action":     knownvalue.StringExact("skip"),
										"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"enabled": knownvalue.Bool(true),
										}),
									}),
								}),
							),
							statecheck.ExpectKnownValue(
								"data.cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"action":     knownvalue.StringExact("skip"),
										"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"enabled": knownvalue.Bool(true),
										}),
									}),
								}),
							),
						},
					},
					{
						ConfigFile:      config.TestNameFile("2.tf"),
						ConfigVariables: configVariables,
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								plancheck.ExpectResourceAction(
									"cloudflare_ruleset.my_ruleset",
									plancheck.ResourceActionUpdate,
								),
								plancheck.ExpectKnownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules"),
									knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectPartial(map[string]knownvalue.Check{
											"id":         knownvalue.NotNull(),
											"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
											"action":     knownvalue.StringExact("skip"),
											"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"enabled": knownvalue.Bool(true),
											}),
										}),
									}),
								),
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"action":     knownvalue.StringExact("skip"),
										"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"enabled": knownvalue.Bool(true),
										}),
									}),
								}),
							),
							statecheck.ExpectKnownValue(
								"data.cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"action":     knownvalue.StringExact("skip"),
										"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"enabled": knownvalue.Bool(true),
										}),
									}),
								}),
							),
						},
					},
				},
			})
		})

		t.Run("id", func(t *testing.T) {
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
									"cloudflare_ruleset.my_ruleset",
									plancheck.ResourceActionCreate,
								),
								plancheck.ExpectKnownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules"),
									knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectPartial(map[string]knownvalue.Check{
											"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
											"action":     knownvalue.StringExact("skip"),
										}),
									}),
								),
								plancheck.ExpectUnknownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"),
								),
								plancheck.ExpectUnknownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("logging"),
								),
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"action":     knownvalue.StringExact("skip"),
										"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"enabled": knownvalue.Bool(true),
										}),
									}),
								}),
							),
							statecheck.ExpectKnownValue(
								"data.cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"action":     knownvalue.StringExact("skip"),
										"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"enabled": knownvalue.Bool(true),
										}),
									}),
								}),
							),
						},
					},
					{
						ConfigFile:      config.TestNameFile("2.tf"),
						ConfigVariables: configVariables,
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								plancheck.ExpectResourceAction(
									"cloudflare_ruleset.my_ruleset",
									plancheck.ResourceActionUpdate,
								),
								plancheck.ExpectKnownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules"),
									knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectPartial(map[string]knownvalue.Check{
											"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
											"action":     knownvalue.StringExact("skip"),
										}),
									}),
								),
								plancheck.ExpectUnknownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"),
								),
								plancheck.ExpectUnknownValue(
									"cloudflare_ruleset.my_ruleset",
									tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("logging"),
								),
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"action":     knownvalue.StringExact("skip"),
										"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"enabled": knownvalue.Bool(true),
										}),
									}),
								}),
							),
							statecheck.ExpectKnownValue(
								"data.cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"action":     knownvalue.StringExact("skip"),
										"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"enabled": knownvalue.Bool(true),
										}),
									}),
								}),
							),
						},
					},
				},
			})
		})
	})
}

func TestAccCloudflareRuleset_RulesRatelimit(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"ratelimit": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"characteristics": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.StringExact("cf.colo.id"),
											knownvalue.StringExact("ip.src"),
										}),
										"period":                     knownvalue.Int64Exact(60),
										"counting_expression":        knownvalue.Null(),
										"requests_per_period":        knownvalue.Int64Exact(10),
										"requests_to_origin":         knownvalue.Bool(false),
										"score_per_period":           knownvalue.Null(),
										"score_response_header_name": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"ratelimit": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"characteristics": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("cf.colo.id"),
										knownvalue.StringExact("ip.src"),
									}),
									"period":                     knownvalue.Int64Exact(60),
									"counting_expression":        knownvalue.Null(),
									"mitigation_timeout":         knownvalue.Int64Exact(0),
									"requests_per_period":        knownvalue.Int64Exact(10),
									"requests_to_origin":         knownvalue.Bool(false),
									"score_per_period":           knownvalue.Null(),
									"score_response_header_name": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"ratelimit": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"characteristics": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("cf.colo.id"),
										knownvalue.StringExact("ip.src"),
									}),
									"period":                     knownvalue.Int64Exact(60),
									"counting_expression":        knownvalue.Null(),
									"mitigation_timeout":         knownvalue.Int64Exact(0),
									"requests_per_period":        knownvalue.Int64Exact(10),
									"requests_to_origin":         knownvalue.Bool(false),
									"score_per_period":           knownvalue.Null(),
									"score_response_header_name": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"ratelimit": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"characteristics": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.StringExact("cf.colo.id"),
											knownvalue.StringExact("ip.src"),
										}),
										"period":                     knownvalue.Int64Exact(60),
										"counting_expression":        knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"mitigation_timeout":         knownvalue.Int64Exact(300),
										"requests_per_period":        knownvalue.Int64Exact(100),
										"requests_to_origin":         knownvalue.Bool(false),
										"score_per_period":           knownvalue.Null(),
										"score_response_header_name": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"ratelimit": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"characteristics": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("cf.colo.id"),
										knownvalue.StringExact("ip.src"),
									}),
									"period":                     knownvalue.Int64Exact(60),
									"counting_expression":        knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"mitigation_timeout":         knownvalue.Int64Exact(300),
									"requests_per_period":        knownvalue.Int64Exact(100),
									"requests_to_origin":         knownvalue.Bool(false),
									"score_per_period":           knownvalue.Null(),
									"score_response_header_name": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"ratelimit": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"characteristics": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("cf.colo.id"),
										knownvalue.StringExact("ip.src"),
									}),
									"period":                     knownvalue.Int64Exact(60),
									"counting_expression":        knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"mitigation_timeout":         knownvalue.Int64Exact(300),
									"requests_per_period":        knownvalue.Int64Exact(100),
									"requests_to_origin":         knownvalue.Bool(false),
									"score_per_period":           knownvalue.Null(),
									"score_response_header_name": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"ratelimit": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"characteristics": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.StringExact("cf.colo.id"),
											knownvalue.StringExact("ip.src"),
										}),
										"period":                     knownvalue.Int64Exact(60),
										"counting_expression":        knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"mitigation_timeout":         knownvalue.Int64Exact(600),
										"requests_per_period":        knownvalue.Null(),
										"requests_to_origin":         knownvalue.Bool(true),
										"score_per_period":           knownvalue.Int64Exact(400),
										"score_response_header_name": knownvalue.StringExact("my-score"),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"ratelimit": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"characteristics": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("cf.colo.id"),
										knownvalue.StringExact("ip.src"),
									}),
									"period":                     knownvalue.Int64Exact(60),
									"counting_expression":        knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"mitigation_timeout":         knownvalue.Int64Exact(600),
									"requests_per_period":        knownvalue.Null(),
									"requests_to_origin":         knownvalue.Bool(true),
									"score_per_period":           knownvalue.Int64Exact(400),
									"score_response_header_name": knownvalue.StringExact("my-score"),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"ratelimit": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"characteristics": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("cf.colo.id"),
										knownvalue.StringExact("ip.src"),
									}),
									"period":                     knownvalue.Int64Exact(60),
									"counting_expression":        knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"mitigation_timeout":         knownvalue.Int64Exact(600),
									"requests_per_period":        knownvalue.Null(),
									"requests_to_origin":         knownvalue.Bool(true),
									"score_per_period":           knownvalue.Int64Exact(400),
									"score_response_header_name": knownvalue.StringExact("my-score"),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_RulesRef(t *testing.T) {
	t.Run("add", func(t *testing.T) {
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
								"cloudflare_ruleset.my_ruleset",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"ref":        knownvalue.StringExact("one"),
									}),
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"ref":        knownvalue.StringExact("two"),
									}),
								}),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("id"),
							),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
						statecheck.ExpectKnownValue(
							"data.cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
					},
				},
				{
					ConfigFile:      config.TestNameFile("2.tf"),
					ConfigVariables: configVariables,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"cloudflare_ruleset.my_ruleset",
								plancheck.ResourceActionUpdate,
							),
							plancheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"ref":        knownvalue.StringExact("one"),
									}),
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 3.3.3.3"),
										"ref":        knownvalue.StringExact("three"),
									}),
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"ref":        knownvalue.StringExact("two"),
									}),
								}),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("id"),
							),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 3.3.3.3"),
									"ref":        knownvalue.StringExact("three"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
						statecheck.ExpectKnownValue(
							"data.cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 3.3.3.3"),
									"ref":        knownvalue.StringExact("three"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
					},
				},
			},
		})
	})

	t.Run("append", func(t *testing.T) {
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
								"cloudflare_ruleset.my_ruleset",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"ref":        knownvalue.StringExact("one"),
									}),
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"ref":        knownvalue.StringExact("two"),
									}),
								}),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("id"),
							),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
						statecheck.ExpectKnownValue(
							"data.cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
					},
				},
				{
					ConfigFile:      config.TestNameFile("2.tf"),
					ConfigVariables: configVariables,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"cloudflare_ruleset.my_ruleset",
								plancheck.ResourceActionUpdate,
							),
							plancheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"ref":        knownvalue.StringExact("one"),
									}),
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"ref":        knownvalue.StringExact("two"),
									}),
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 3.3.3.3"),
										"ref":        knownvalue.StringExact("three"),
									}),
								}),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("id"),
							),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 3.3.3.3"),
									"ref":        knownvalue.StringExact("three"),
								}),
							}),
						),
						statecheck.ExpectKnownValue(
							"data.cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 3.3.3.3"),
									"ref":        knownvalue.StringExact("three"),
								}),
							}),
						),
					},
				},
			},
		})
	})

	t.Run("modify", func(t *testing.T) {
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
								"cloudflare_ruleset.my_ruleset",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"ref":        knownvalue.StringExact("one"),
									}),
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"ref":        knownvalue.StringExact("two"),
									}),
								}),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("id"),
							),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
						statecheck.ExpectKnownValue(
							"data.cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
					},
				},
				{
					ConfigFile:      config.TestNameFile("2.tf"),
					ConfigVariables: configVariables,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"cloudflare_ruleset.my_ruleset",
								plancheck.ResourceActionUpdate,
							),
							plancheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"ref":        knownvalue.StringExact("one"),
									}),
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 3.3.3.3"),
										"ref":        knownvalue.StringExact("two"),
									}),
								}),
							),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 3.3.3.3"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
						statecheck.ExpectKnownValue(
							"data.cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 3.3.3.3"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
					},
				},
			},
		})
	})

	t.Run("remove", func(t *testing.T) {
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
								"cloudflare_ruleset.my_ruleset",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"ref":        knownvalue.StringExact("one"),
									}),
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"ref":        knownvalue.StringExact("two"),
									}),
								}),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("id"),
							),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
						statecheck.ExpectKnownValue(
							"data.cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
					},
				},
				{
					ConfigFile:      config.TestNameFile("2.tf"),
					ConfigVariables: configVariables,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"cloudflare_ruleset.my_ruleset",
								plancheck.ResourceActionUpdate,
							),
							plancheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"ref":        knownvalue.StringExact("two"),
									}),
								}),
							),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
						statecheck.ExpectKnownValue(
							"data.cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
					},
				},
			},
		})
	})

	t.Run("reverse", func(t *testing.T) {
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
								"cloudflare_ruleset.my_ruleset",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"ref":        knownvalue.StringExact("one"),
									}),
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"ref":        knownvalue.StringExact("two"),
									}),
								}),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("id"),
							),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
						statecheck.ExpectKnownValue(
							"data.cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
					},
				},
				{
					ConfigFile:      config.TestNameFile("2.tf"),
					ConfigVariables: configVariables,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"cloudflare_ruleset.my_ruleset",
								plancheck.ResourceActionUpdate,
							),
							plancheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"ref":        knownvalue.StringExact("two"),
									}),
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"ref":        knownvalue.StringExact("one"),
									}),
								}),
							),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
							}),
						),
						statecheck.ExpectKnownValue(
							"data.cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
							}),
						),
					},
				},
			},
		})
	})

	t.Run("truncate", func(t *testing.T) {
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
								"cloudflare_ruleset.my_ruleset",
								plancheck.ResourceActionCreate,
							),
							plancheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"ref":        knownvalue.StringExact("one"),
									}),
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
										"ref":        knownvalue.StringExact("two"),
									}),
								}),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"),
							),
							plancheck.ExpectUnknownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("id"),
							),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
						statecheck.ExpectKnownValue(
							"data.cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 2.2.2.2"),
									"ref":        knownvalue.StringExact("two"),
								}),
							}),
						),
					},
				},
				{
					ConfigFile:      config.TestNameFile("2.tf"),
					ConfigVariables: configVariables,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"cloudflare_ruleset.my_ruleset",
								plancheck.ResourceActionUpdate,
							),
							plancheck.ExpectKnownValue(
								"cloudflare_ruleset.my_ruleset",
								tfjsonpath.New("rules"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":         knownvalue.NotNull(),
										"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
										"ref":        knownvalue.StringExact("one"),
									}),
								}),
							),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
							}),
						),
						statecheck.ExpectKnownValue(
							"data.cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":         knownvalue.NotNull(),
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
							}),
						),
					},
				},
			},
		})
	})
}

func TestAccCloudflareRuleset_LastUpdated(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectUnknownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("last_updated"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("last_updated"),
						knownvalue.NotNull(),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("last_updated"),
						knownvalue.NotNull(),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectUnknownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("last_updated"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("last_updated"),
						knownvalue.NotNull(),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("last_updated"),
						knownvalue.NotNull(),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_Version(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectUnknownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("version"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("version"),
						knownvalue.StringExact("1"),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("version"),
						knownvalue.StringExact("1"),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectUnknownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("version"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("version"),
						knownvalue.StringExact("2"),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("version"),
						knownvalue.StringExact("2"),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_BlockRules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("block"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"response": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("block"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"response": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("block"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"response": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("block"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"response": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"status_code":  knownvalue.Int64Exact(403),
											"content":      knownvalue.StringExact("Access denied"),
											"content_type": knownvalue.StringExact("text/plain"),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("block"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"response": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"status_code":  knownvalue.Int64Exact(403),
										"content":      knownvalue.StringExact("Access denied"),
										"content_type": knownvalue.StringExact("text/plain"),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("block"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"response": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"status_code":  knownvalue.Int64Exact(403),
										"content":      knownvalue.StringExact("Access denied"),
										"content_type": knownvalue.StringExact("text/plain"),
									}),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_ChallengeRules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("challenge"),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("challenge"),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("challenge"),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("js_challenge"),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("js_challenge"),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("js_challenge"),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("managed_challenge"),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("managed_challenge"),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("managed_challenge"),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_CompressResponseRules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("compress_response"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"algorithms": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"name": knownvalue.StringExact("auto"),
											}),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("compress_response"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"algorithms": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name": knownvalue.StringExact("auto"),
										}),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("compress_response"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"algorithms": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name": knownvalue.StringExact("auto"),
										}),
									}),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("compress_response"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"algorithms": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"name": knownvalue.StringExact("brotli"),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"name": knownvalue.StringExact("gzip"),
											}),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("compress_response"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"algorithms": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name": knownvalue.StringExact("brotli"),
										}),
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name": knownvalue.StringExact("gzip"),
										}),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("compress_response"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"algorithms": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name": knownvalue.StringExact("brotli"),
										}),
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name": knownvalue.StringExact("gzip"),
										}),
									}),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_ExecuteRules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("execute"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":           knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
										"matched_data": knownvalue.Null(),
										"overrides":    knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("execute"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":           knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
									"matched_data": knownvalue.Null(),
									"overrides":    knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("execute"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":           knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
									"matched_data": knownvalue.Null(),
									"overrides":    knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("execute"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id": knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
										"matched_data": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"public_key": knownvalue.StringExact("iGqBmyIUxuWt1rvxoAharN9FUXneUBxA/Y19PyyrEG0="),
										}),
										"overrides": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"action":            knownvalue.StringExact("log"),
											"categories":        knownvalue.Null(),
											"enabled":           knownvalue.Null(),
											"rules":             knownvalue.Null(),
											"sensitivity_level": knownvalue.Null(),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("execute"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id": knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
									"matched_data": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"public_key": knownvalue.StringExact("iGqBmyIUxuWt1rvxoAharN9FUXneUBxA/Y19PyyrEG0="),
									}),
									"overrides": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"action":            knownvalue.StringExact("log"),
										"categories":        knownvalue.Null(),
										"enabled":           knownvalue.Null(),
										"rules":             knownvalue.Null(),
										"sensitivity_level": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("execute"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id": knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
									"matched_data": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"public_key": knownvalue.StringExact("iGqBmyIUxuWt1rvxoAharN9FUXneUBxA/Y19PyyrEG0="),
									}),
									"overrides": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"action":            knownvalue.StringExact("log"),
										"categories":        knownvalue.Null(),
										"enabled":           knownvalue.Null(),
										"rules":             knownvalue.Null(),
										"sensitivity_level": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("execute"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id": knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
										"matched_data": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"public_key": knownvalue.StringExact("iGqBmyIUxuWt1rvxoAharN9FUXneUBxA/Y19PyyrEG0="),
										}),
										"overrides": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"action":            knownvalue.Null(),
											"categories":        knownvalue.Null(),
											"enabled":           knownvalue.Bool(false),
											"rules":             knownvalue.Null(),
											"sensitivity_level": knownvalue.Null(),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("execute"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id": knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
									"matched_data": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"public_key": knownvalue.StringExact("iGqBmyIUxuWt1rvxoAharN9FUXneUBxA/Y19PyyrEG0="),
									}),
									"overrides": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"action":            knownvalue.Null(),
										"categories":        knownvalue.Null(),
										"enabled":           knownvalue.Bool(false),
										"rules":             knownvalue.Null(),
										"sensitivity_level": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("execute"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id": knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
									"matched_data": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"public_key": knownvalue.StringExact("iGqBmyIUxuWt1rvxoAharN9FUXneUBxA/Y19PyyrEG0="),
									}),
									"overrides": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"action":            knownvalue.Null(),
										"categories":        knownvalue.Null(),
										"enabled":           knownvalue.Bool(false),
										"rules":             knownvalue.Null(),
										"sensitivity_level": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("4.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("execute"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":           knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
										"matched_data": knownvalue.Null(),
										"overrides": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"action": knownvalue.StringExact("log"),
											"categories": knownvalue.ListExact([]knownvalue.Check{
												knownvalue.ObjectExact(map[string]knownvalue.Check{
													"category":          knownvalue.StringExact("language-java"),
													"action":            knownvalue.StringExact("block"),
													"enabled":           knownvalue.Null(),
													"sensitivity_level": knownvalue.Null(),
												}),
												knownvalue.ObjectExact(map[string]knownvalue.Check{
													"category":          knownvalue.StringExact("language-php"),
													"action":            knownvalue.Null(),
													"enabled":           knownvalue.Bool(false),
													"sensitivity_level": knownvalue.Null(),
												}),
												knownvalue.ObjectExact(map[string]knownvalue.Check{
													"category":          knownvalue.StringExact("language-shell"),
													"action":            knownvalue.StringExact("block"),
													"enabled":           knownvalue.Bool(true),
													"sensitivity_level": knownvalue.Null(),
												}),
											}),
											"enabled": knownvalue.Bool(true),
											"rules": knownvalue.ListExact([]knownvalue.Check{
												knownvalue.ObjectExact(map[string]knownvalue.Check{
													"id":                knownvalue.StringExact("04116d14d7524986ba314d11c8a41e11"),
													"action":            knownvalue.StringExact("block"),
													"enabled":           knownvalue.Null(),
													"score_threshold":   knownvalue.Null(),
													"sensitivity_level": knownvalue.Null(),
												}),
												knownvalue.ObjectExact(map[string]knownvalue.Check{
													"id":                knownvalue.StringExact("55b58c71f653446fa0942cf7700f8c8e"),
													"action":            knownvalue.Null(),
													"enabled":           knownvalue.Bool(false),
													"score_threshold":   knownvalue.Null(),
													"sensitivity_level": knownvalue.Null(),
												}),
												knownvalue.ObjectExact(map[string]knownvalue.Check{
													"id":                knownvalue.StringExact("7683285d70b14023ac407b67eccbb280"),
													"action":            knownvalue.StringExact("block"),
													"enabled":           knownvalue.Bool(true),
													"score_threshold":   knownvalue.Int64Exact(40),
													"sensitivity_level": knownvalue.Null(),
												}),
											}),
											"sensitivity_level": knownvalue.Null(),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("execute"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":           knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
									"matched_data": knownvalue.Null(),
									"overrides": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"action": knownvalue.StringExact("log"),
										"categories": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"category":          knownvalue.StringExact("language-java"),
												"action":            knownvalue.StringExact("block"),
												"enabled":           knownvalue.Null(),
												"sensitivity_level": knownvalue.Null(),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"category":          knownvalue.StringExact("language-php"),
												"action":            knownvalue.Null(),
												"enabled":           knownvalue.Bool(false),
												"sensitivity_level": knownvalue.Null(),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"category":          knownvalue.StringExact("language-shell"),
												"action":            knownvalue.StringExact("block"),
												"enabled":           knownvalue.Bool(true),
												"sensitivity_level": knownvalue.Null(),
											}),
										}),
										"enabled": knownvalue.Bool(true),
										"rules": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"id":                knownvalue.StringExact("04116d14d7524986ba314d11c8a41e11"),
												"action":            knownvalue.StringExact("block"),
												"enabled":           knownvalue.Null(),
												"score_threshold":   knownvalue.Null(),
												"sensitivity_level": knownvalue.Null(),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"id":                knownvalue.StringExact("55b58c71f653446fa0942cf7700f8c8e"),
												"action":            knownvalue.Null(),
												"enabled":           knownvalue.Bool(false),
												"score_threshold":   knownvalue.Null(),
												"sensitivity_level": knownvalue.Null(),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"id":                knownvalue.StringExact("7683285d70b14023ac407b67eccbb280"),
												"action":            knownvalue.StringExact("block"),
												"enabled":           knownvalue.Bool(true),
												"score_threshold":   knownvalue.Int64Exact(40),
												"sensitivity_level": knownvalue.Null(),
											}),
										}),
										"sensitivity_level": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("execute"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":           knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
									"matched_data": knownvalue.Null(),
									"overrides": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"action": knownvalue.StringExact("log"),
										"categories": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"category":          knownvalue.StringExact("language-java"),
												"action":            knownvalue.StringExact("block"),
												"enabled":           knownvalue.Null(),
												"sensitivity_level": knownvalue.Null(),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"category":          knownvalue.StringExact("language-php"),
												"action":            knownvalue.Null(),
												"enabled":           knownvalue.Bool(false),
												"sensitivity_level": knownvalue.Null(),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"category":          knownvalue.StringExact("language-shell"),
												"action":            knownvalue.StringExact("block"),
												"enabled":           knownvalue.Bool(true),
												"sensitivity_level": knownvalue.Null(),
											}),
										}),
										"enabled": knownvalue.Bool(true),
										"rules": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"id":                knownvalue.StringExact("04116d14d7524986ba314d11c8a41e11"),
												"action":            knownvalue.StringExact("block"),
												"enabled":           knownvalue.Null(),
												"score_threshold":   knownvalue.Null(),
												"sensitivity_level": knownvalue.Null(),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"id":                knownvalue.StringExact("55b58c71f653446fa0942cf7700f8c8e"),
												"action":            knownvalue.Null(),
												"enabled":           knownvalue.Bool(false),
												"score_threshold":   knownvalue.Null(),
												"sensitivity_level": knownvalue.Null(),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"id":                knownvalue.StringExact("7683285d70b14023ac407b67eccbb280"),
												"action":            knownvalue.StringExact("block"),
												"enabled":           knownvalue.Bool(true),
												"score_threshold":   knownvalue.Int64Exact(40),
												"sensitivity_level": knownvalue.Null(),
											}),
										}),
										"sensitivity_level": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("5.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionReplace,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("execute"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"id":           knownvalue.StringExact("4d21379b4f9f4bb088e0729962c8b3cf"),
										"matched_data": knownvalue.Null(),
										"overrides": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"action": knownvalue.Null(),
											"categories": knownvalue.ListExact([]knownvalue.Check{
												knownvalue.ObjectExact(map[string]knownvalue.Check{
													"category":          knownvalue.StringExact("botnets"),
													"action":            knownvalue.Null(),
													"enabled":           knownvalue.Null(),
													"sensitivity_level": knownvalue.StringExact("medium"),
												}),
											}),
											"enabled": knownvalue.Null(),
											"rules": knownvalue.ListExact([]knownvalue.Check{
												knownvalue.ObjectExact(map[string]knownvalue.Check{
													"id":                knownvalue.StringExact("8fc7efb08f984ced8d61b34b254da96a"),
													"action":            knownvalue.Null(),
													"enabled":           knownvalue.Null(),
													"score_threshold":   knownvalue.Null(),
													"sensitivity_level": knownvalue.StringExact("low"),
												}),
											}),
											"sensitivity_level": knownvalue.StringExact("eoff"),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("execute"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":           knownvalue.StringExact("4d21379b4f9f4bb088e0729962c8b3cf"),
									"matched_data": knownvalue.Null(),
									"overrides": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"action": knownvalue.Null(),
										"categories": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"category":          knownvalue.StringExact("botnets"),
												"action":            knownvalue.Null(),
												"enabled":           knownvalue.Null(),
												"sensitivity_level": knownvalue.StringExact("medium"),
											}),
										}),
										"enabled": knownvalue.Null(),
										"rules": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"id":                knownvalue.StringExact("8fc7efb08f984ced8d61b34b254da96a"),
												"action":            knownvalue.Null(),
												"enabled":           knownvalue.Null(),
												"score_threshold":   knownvalue.Null(),
												"sensitivity_level": knownvalue.StringExact("low"),
											}),
										}),
										"sensitivity_level": knownvalue.StringExact("eoff"),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("execute"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"id":           knownvalue.StringExact("4d21379b4f9f4bb088e0729962c8b3cf"),
									"matched_data": knownvalue.Null(),
									"overrides": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"action": knownvalue.Null(),
										"categories": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"category":          knownvalue.StringExact("botnets"),
												"action":            knownvalue.Null(),
												"enabled":           knownvalue.Null(),
												"sensitivity_level": knownvalue.StringExact("medium"),
											}),
										}),
										"enabled": knownvalue.Null(),
										"rules": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"id":                knownvalue.StringExact("8fc7efb08f984ced8d61b34b254da96a"),
												"action":            knownvalue.Null(),
												"enabled":           knownvalue.Null(),
												"score_threshold":   knownvalue.Null(),
												"sensitivity_level": knownvalue.StringExact("low"),
											}),
										}),
										"sensitivity_level": knownvalue.StringExact("eoff"),
									}),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_LogRules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("log"),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("log"),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("log"),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_LogCustomFieldRules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("log_custom_field"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"cookie_fields": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"name": knownvalue.StringExact("__cfruid"),
											}),
										}),
										"raw_response_fields":        knownvalue.Null(),
										"request_fields":             knownvalue.Null(),
										"response_fields":            knownvalue.Null(),
										"transformed_request_fields": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("log_custom_field"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"cookie_fields": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name": knownvalue.StringExact("__cfruid"),
										}),
									}),
									"raw_response_fields":        knownvalue.Null(),
									"request_fields":             knownvalue.Null(),
									"response_fields":            knownvalue.Null(),
									"transformed_request_fields": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("log_custom_field"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"cookie_fields": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name": knownvalue.StringExact("__cfruid"),
										}),
									}),
									"raw_response_fields":        knownvalue.Null(),
									"request_fields":             knownvalue.Null(),
									"response_fields":            knownvalue.Null(),
									"transformed_request_fields": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("log_custom_field"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"cookie_fields": knownvalue.Null(),
										"raw_response_fields": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"name":                knownvalue.StringExact("allow"),
												"preserve_duplicates": knownvalue.Bool(false),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"name":                knownvalue.StringExact("content-type"),
												"preserve_duplicates": knownvalue.Bool(false),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"name":                knownvalue.StringExact("server"),
												"preserve_duplicates": knownvalue.Bool(true),
											}),
										}),
										"request_fields": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"name": knownvalue.StringExact("content-type"),
											}),
										}),
										"response_fields": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"name":                knownvalue.StringExact("access-control-allow-origin"),
												"preserve_duplicates": knownvalue.Bool(false),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"name":                knownvalue.StringExact("connection"),
												"preserve_duplicates": knownvalue.Bool(false),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"name":                knownvalue.StringExact("set-cookie"),
												"preserve_duplicates": knownvalue.Bool(true),
											}),
										}),
										"transformed_request_fields": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"name": knownvalue.StringExact("host"),
											}),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("log_custom_field"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"cookie_fields": knownvalue.Null(),
									"raw_response_fields": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name":                knownvalue.StringExact("allow"),
											"preserve_duplicates": knownvalue.Bool(false),
										}),
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name":                knownvalue.StringExact("content-type"),
											"preserve_duplicates": knownvalue.Bool(false),
										}),
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name":                knownvalue.StringExact("server"),
											"preserve_duplicates": knownvalue.Bool(true),
										}),
									}),
									"request_fields": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name": knownvalue.StringExact("content-type"),
										}),
									}),
									"response_fields": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name":                knownvalue.StringExact("access-control-allow-origin"),
											"preserve_duplicates": knownvalue.Bool(false),
										}),
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name":                knownvalue.StringExact("connection"),
											"preserve_duplicates": knownvalue.Bool(false),
										}),
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name":                knownvalue.StringExact("set-cookie"),
											"preserve_duplicates": knownvalue.Bool(true),
										}),
									}),
									"transformed_request_fields": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name": knownvalue.StringExact("host"),
										}),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("log_custom_field"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"cookie_fields": knownvalue.Null(),
									"raw_response_fields": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name":                knownvalue.StringExact("allow"),
											"preserve_duplicates": knownvalue.Bool(false),
										}),
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name":                knownvalue.StringExact("content-type"),
											"preserve_duplicates": knownvalue.Bool(false),
										}),
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name":                knownvalue.StringExact("server"),
											"preserve_duplicates": knownvalue.Bool(true),
										}),
									}),
									"request_fields": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name": knownvalue.StringExact("content-type"),
										}),
									}),
									"response_fields": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name":                knownvalue.StringExact("access-control-allow-origin"),
											"preserve_duplicates": knownvalue.Bool(false),
										}),
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name":                knownvalue.StringExact("connection"),
											"preserve_duplicates": knownvalue.Bool(false),
										}),
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name":                knownvalue.StringExact("set-cookie"),
											"preserve_duplicates": knownvalue.Bool(true),
										}),
									}),
									"transformed_request_fields": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"name": knownvalue.StringExact("host"),
										}),
									}),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_RedirectRules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("redirect"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"from_list": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"key":  knownvalue.StringExact("http.request.full_uri"),
											"name": knownvalue.StringExact("my_list"),
										}),
										"from_value": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("redirect"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"from_list": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"key":  knownvalue.StringExact("http.request.full_uri"),
										"name": knownvalue.StringExact("my_list"),
									}),
									"from_value": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("redirect"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"from_list": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"key":  knownvalue.StringExact("http.request.full_uri"),
										"name": knownvalue.StringExact("my_list"),
									}),
									"from_value": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionReplace,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("redirect"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"from_list": knownvalue.Null(),
										"from_value": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"preserve_query_string": knownvalue.Bool(false),
											"status_code":           knownvalue.Null(),
											"target_url": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"value":      knownvalue.StringExact("https://example.com"),
												"expression": knownvalue.Null(),
											}),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("redirect"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"from_list": knownvalue.Null(),
									"from_value": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"preserve_query_string": knownvalue.Bool(false),
										"status_code":           knownvalue.Null(),
										"target_url": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.StringExact("https://example.com"),
											"expression": knownvalue.Null(),
										}),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("redirect"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"from_list": knownvalue.Null(),
									"from_value": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"preserve_query_string": knownvalue.Bool(false),
										"status_code":           knownvalue.Null(),
										"target_url": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.StringExact("https://example.com"),
											"expression": knownvalue.Null(),
										}),
									}),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("redirect"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"from_list": knownvalue.Null(),
									"from_value": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"preserve_query_string": knownvalue.Bool(false),
										"status_code":           knownvalue.Null(),
										"target_url": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.StringExact("https://example.com"),
											"expression": knownvalue.Null(),
										}),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("redirect"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"from_list": knownvalue.Null(),
									"from_value": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"preserve_query_string": knownvalue.Bool(false),
										"status_code":           knownvalue.Null(),
										"target_url": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.StringExact("https://example.com"),
											"expression": knownvalue.Null(),
										}),
									}),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("4.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("redirect"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"from_list": knownvalue.Null(),
										"from_value": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"preserve_query_string": knownvalue.Bool(true),
											"status_code":           knownvalue.Int64Exact(301),
											"target_url": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"value":      knownvalue.Null(),
												"expression": knownvalue.StringExact("concat(\"https://m.example.com\", http.request.uri.path)"),
											}),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("redirect"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"from_list": knownvalue.Null(),
									"from_value": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"preserve_query_string": knownvalue.Bool(true),
										"status_code":           knownvalue.Int64Exact(301),
										"target_url": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.Null(),
											"expression": knownvalue.StringExact("concat(\"https://m.example.com\", http.request.uri.path)"),
										}),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("redirect"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"from_list": knownvalue.Null(),
									"from_value": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"preserve_query_string": knownvalue.Bool(true),
										"status_code":           knownvalue.Int64Exact(301),
										"target_url": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.Null(),
											"expression": knownvalue.StringExact("concat(\"https://m.example.com\", http.request.uri.path)"),
										}),
									}),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_RewriteRules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("rewrite"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"headers": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"my-first-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"operation":  knownvalue.StringExact("set"),
												"value":      knownvalue.StringExact("my-first-header-value"),
												"expression": knownvalue.Null(),
											}),
											"my-second-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"operation":  knownvalue.StringExact("set"),
												"value":      knownvalue.Null(),
												"expression": knownvalue.StringExact("ip.src"),
											}),
											"my-third-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"operation":  knownvalue.StringExact("remove"),
												"value":      knownvalue.Null(),
												"expression": knownvalue.Null(),
											}),
										}),
										"uri": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"my-first-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("set"),
											"value":      knownvalue.StringExact("my-first-header-value"),
											"expression": knownvalue.Null(),
										}),
										"my-second-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("set"),
											"value":      knownvalue.Null(),
											"expression": knownvalue.StringExact("ip.src"),
										}),
										"my-third-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("remove"),
											"value":      knownvalue.Null(),
											"expression": knownvalue.Null(),
										}),
									}),
									"uri": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"my-first-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("set"),
											"value":      knownvalue.StringExact("my-first-header-value"),
											"expression": knownvalue.Null(),
										}),
										"my-second-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("set"),
											"value":      knownvalue.Null(),
											"expression": knownvalue.StringExact("ip.src"),
										}),
										"my-third-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("remove"),
											"value":      knownvalue.Null(),
											"expression": knownvalue.Null(),
										}),
									}),
									"uri": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionReplace,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("rewrite"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"headers": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"my-first-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"operation":  knownvalue.StringExact("add"),
												"value":      knownvalue.StringExact("my-first-header-value"),
												"expression": knownvalue.Null(),
											}),
											"my-second-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"operation":  knownvalue.StringExact("add"),
												"value":      knownvalue.Null(),
												"expression": knownvalue.StringExact("http.host"),
											}),
											"my-third-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"operation":  knownvalue.StringExact("set"),
												"value":      knownvalue.StringExact("my-third-header-value"),
												"expression": knownvalue.Null(),
											}),
											"my-fourth-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"operation":  knownvalue.StringExact("set"),
												"value":      knownvalue.Null(),
												"expression": knownvalue.StringExact("ip.src"),
											}),
											"my-fifth-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"operation":  knownvalue.StringExact("remove"),
												"value":      knownvalue.Null(),
												"expression": knownvalue.Null(),
											}),
										}),
										"uri": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"my-first-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("add"),
											"value":      knownvalue.StringExact("my-first-header-value"),
											"expression": knownvalue.Null(),
										}),
										"my-second-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("add"),
											"value":      knownvalue.Null(),
											"expression": knownvalue.StringExact("http.host"),
										}),
										"my-third-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("set"),
											"value":      knownvalue.StringExact("my-third-header-value"),
											"expression": knownvalue.Null(),
										}),
										"my-fourth-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("set"),
											"value":      knownvalue.Null(),
											"expression": knownvalue.StringExact("ip.src"),
										}),
										"my-fifth-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("remove"),
											"value":      knownvalue.Null(),
											"expression": knownvalue.Null(),
										}),
									}),
									"uri": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"my-first-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("add"),
											"value":      knownvalue.StringExact("my-first-header-value"),
											"expression": knownvalue.Null(),
										}),
										"my-second-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("add"),
											"value":      knownvalue.Null(),
											"expression": knownvalue.StringExact("http.host"),
										}),
										"my-third-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("set"),
											"value":      knownvalue.StringExact("my-third-header-value"),
											"expression": knownvalue.Null(),
										}),
										"my-fourth-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("set"),
											"value":      knownvalue.Null(),
											"expression": knownvalue.StringExact("ip.src"),
										}),
										"my-fifth-header": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("remove"),
											"value":      knownvalue.Null(),
											"expression": knownvalue.Null(),
										}),
									}),
									"uri": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionReplace,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("rewrite"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"headers": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"Exposed-Credential-Check": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"operation":  knownvalue.StringExact("set"),
												"value":      knownvalue.StringExact("1"),
												"expression": knownvalue.Null(),
											}),
										}),
										"uri": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"Exposed-Credential-Check": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("set"),
											"value":      knownvalue.StringExact("1"),
											"expression": knownvalue.Null(),
										}),
									}),
									"uri": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"Exposed-Credential-Check": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"operation":  knownvalue.StringExact("set"),
											"value":      knownvalue.StringExact("1"),
											"expression": knownvalue.Null(),
										}),
									}),
									"uri": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("4.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionReplace,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("rewrite"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"headers": knownvalue.Null(),
										"uri": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"path": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"value":      knownvalue.StringExact("/foo"),
												"expression": knownvalue.Null(),
											}),
											"query": knownvalue.Null(),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.Null(),
									"uri": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"path": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.StringExact("/foo"),
											"expression": knownvalue.Null(),
										}),
										"query": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.Null(),
									"uri": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"path": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.StringExact("/foo"),
											"expression": knownvalue.Null(),
										}),
										"query":  knownvalue.Null(),
										"origin": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("5.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("rewrite"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"headers": knownvalue.Null(),
										"uri": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"path": knownvalue.Null(),
											"query": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"value":      knownvalue.StringExact(""),
												"expression": knownvalue.Null(),
											}),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.Null(),
									"uri": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"path": knownvalue.Null(),
										"query": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.StringExact(""),
											"expression": knownvalue.Null(),
										}),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.Null(),
									"uri": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"path": knownvalue.Null(),
										"query": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.StringExact(""),
											"expression": knownvalue.Null(),
										}),
										"origin": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("6.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("rewrite"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"headers": knownvalue.Null(),
										"uri": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"path": knownvalue.Null(),
											"query": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"value":      knownvalue.StringExact("foo=bar"),
												"expression": knownvalue.Null(),
											}),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.Null(),
									"uri": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"path": knownvalue.Null(),
										"query": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.StringExact("foo=bar"),
											"expression": knownvalue.Null(),
										}),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.Null(),
									"uri": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"path": knownvalue.Null(),
										"query": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.StringExact("foo=bar"),
											"expression": knownvalue.Null(),
										}),
										"origin": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("7.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("rewrite"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"headers": knownvalue.Null(),
										"uri": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"path": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"value":      knownvalue.Null(),
												"expression": knownvalue.StringExact("regex_replace(http.request.uri.path, \"/foo$\", \"/bar\")"),
											}),
											"query": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"value":      knownvalue.Null(),
												"expression": knownvalue.StringExact("regex_replace(http.request.uri.query, \"foo=bar\", \"\")"),
											}),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.Null(),
									"uri": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"path": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.Null(),
											"expression": knownvalue.StringExact("regex_replace(http.request.uri.path, \"/foo$\", \"/bar\")"),
										}),
										"query": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.Null(),
											"expression": knownvalue.StringExact("regex_replace(http.request.uri.query, \"foo=bar\", \"\")"),
										}),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("rewrite"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"headers": knownvalue.Null(),
									"uri": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"path": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.Null(),
											"expression": knownvalue.StringExact("regex_replace(http.request.uri.path, \"/foo$\", \"/bar\")"),
										}),
										"query": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value":      knownvalue.Null(),
											"expression": knownvalue.StringExact("regex_replace(http.request.uri.query, \"foo=bar\", \"\")"),
										}),
										"origin": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_RouteRules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("route"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"host_header": knownvalue.StringExact(domain),
										"origin":      knownvalue.Null(),
										"sni":         knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("route"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"host_header": knownvalue.StringExact(domain),
									"origin":      knownvalue.Null(),
									"sni":         knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("route"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"host_header": knownvalue.StringExact(domain),
									"origin":      knownvalue.Null(),
									"sni":         knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("route"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"host_header": knownvalue.Null(),
										"origin": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"host": knownvalue.StringExact(domain),
											"port": knownvalue.Null(),
										}),
										"sni": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("route"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"host_header": knownvalue.Null(),
									"origin": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"host": knownvalue.StringExact(domain),
										"port": knownvalue.Null(),
									}),
									"sni": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("route"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"host_header": knownvalue.Null(),
									"origin": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"host": knownvalue.StringExact(domain),
										"port": knownvalue.Null(),
									}),
									"sni": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("route"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"host_header": knownvalue.Null(),
										"origin": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"host": knownvalue.Null(),
											"port": knownvalue.Int64Exact(80),
										}),
										"sni": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("route"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"host_header": knownvalue.Null(),
									"origin": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"host": knownvalue.Null(),
										"port": knownvalue.Int64Exact(80),
									}),
									"sni": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("route"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"host_header": knownvalue.Null(),
									"origin": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"host": knownvalue.Null(),
										"port": knownvalue.Int64Exact(80),
									}),
									"sni": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("4.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("route"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"host_header": knownvalue.Null(),
										"origin":      knownvalue.Null(),
										"sni": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"value": knownvalue.StringExact(domain),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("route"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"host_header": knownvalue.Null(),
									"origin":      knownvalue.Null(),
									"sni": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"value": knownvalue.StringExact(domain),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("route"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"host_header": knownvalue.Null(),
									"origin":      knownvalue.Null(),
									"sni": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"value": knownvalue.StringExact(domain),
									}),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_ServeErrorRules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("serve_error"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"asset_name":   knownvalue.Null(),
										"content":      knownvalue.StringExact("1xxx error occurred"),
										"content_type": knownvalue.StringExact("text/plain"),
										"status_code":  knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("serve_error"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"asset_name":   knownvalue.Null(),
									"content":      knownvalue.StringExact("1xxx error occurred"),
									"content_type": knownvalue.StringExact("text/plain"),
									"status_code":  knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("serve_error"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"asset_name":   knownvalue.Null(),
									"content":      knownvalue.StringExact("1xxx error occurred"),
									"content_type": knownvalue.StringExact("text/plain"),
									"status_code":  knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("serve_error"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"asset_name":   knownvalue.StringExact("my_asset"),
										"content":      knownvalue.Null(),
										"content_type": knownvalue.StringExact("text/html"),
										"status_code":  knownvalue.Int64Exact(500),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("serve_error"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"asset_name":   knownvalue.StringExact("my_asset"),
									"content":      knownvalue.Null(),
									"content_type": knownvalue.StringExact("text/html"),
									"status_code":  knownvalue.Int64Exact(500),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("serve_error"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"asset_name":   knownvalue.StringExact("my_asset"),
									"content":      knownvalue.Null(),
									"content_type": knownvalue.StringExact("text/html"),
									"status_code":  knownvalue.Int64Exact(500),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_SetCacheSettingsRules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_cache_settings"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"additional_cacheable_ports": knownvalue.Null(),
										"browser_ttl":                knownvalue.Null(),
										"cache":                      knownvalue.Bool(false),
										"cache_key":                  knownvalue.Null(),
										"cache_reserve":              knownvalue.Null(),
										"edge_ttl":                   knownvalue.Null(),
										"origin_cache_control":       knownvalue.Null(),
										"origin_error_page_passthru": knownvalue.Null(),
										"read_timeout":               knownvalue.Null(),
										"respect_strong_etags":       knownvalue.Null(),
										"serve_stale":                knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl":                knownvalue.Null(),
									"cache":                      knownvalue.Bool(false),
									"cache_key":                  knownvalue.Null(),
									"cache_reserve":              knownvalue.Null(),
									"edge_ttl":                   knownvalue.Null(),
									"origin_cache_control":       knownvalue.Null(),
									"origin_error_page_passthru": knownvalue.Null(),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Null(),
									"serve_stale":                knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl":                knownvalue.Null(),
									"cache":                      knownvalue.Bool(false),
									"cache_key":                  knownvalue.Null(),
									"cache_reserve":              knownvalue.Null(),
									"edge_ttl":                   knownvalue.Null(),
									"origin_cache_control":       knownvalue.Null(),
									"origin_error_page_passthru": knownvalue.Null(),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Null(),
									"serve_stale":                knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_cache_settings"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"additional_cacheable_ports": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.Int64Exact(8080),
										}),
										"browser_ttl": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"mode":    knownvalue.StringExact("respect_origin"),
											"default": knownvalue.Null(),
										}),
										"cache": knownvalue.Bool(true),
										"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cache_by_device_type":       knownvalue.Null(),
											"cache_deception_armor":      knownvalue.Null(),
											"custom_key":                 knownvalue.Null(),
											"ignore_query_strings_order": knownvalue.Null(),
										}),
										"cache_reserve": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"eligible":          knownvalue.Bool(false),
											"minimum_file_size": knownvalue.Null(),
										}),
										"edge_ttl": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"mode":            knownvalue.StringExact("respect_origin"),
											"default":         knownvalue.Null(),
											"status_code_ttl": knownvalue.Null(),
										}),
										"origin_cache_control":       knownvalue.Bool(false),
										"origin_error_page_passthru": knownvalue.Bool(false),
										"read_timeout":               knownvalue.Int64Exact(900),
										"respect_strong_etags":       knownvalue.Bool(false),
										"serve_stale": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"disable_stale_while_updating": knownvalue.Null(),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.Int64Exact(8080),
									}),
									"browser_ttl": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"mode":    knownvalue.StringExact("respect_origin"),
										"default": knownvalue.Null(),
									}),
									"cache": knownvalue.Bool(true),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":       knownvalue.Null(),
										"cache_deception_armor":      knownvalue.Null(),
										"custom_key":                 knownvalue.Null(),
										"ignore_query_strings_order": knownvalue.Null(),
									}),
									"cache_reserve": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"eligible":          knownvalue.Bool(false),
										"minimum_file_size": knownvalue.Null(),
									}),
									"edge_ttl": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"mode":            knownvalue.StringExact("respect_origin"),
										"default":         knownvalue.Null(),
										"status_code_ttl": knownvalue.Null(),
									}),
									"origin_cache_control":       knownvalue.Bool(false),
									"origin_error_page_passthru": knownvalue.Bool(false),
									"read_timeout":               knownvalue.Int64Exact(900),
									"respect_strong_etags":       knownvalue.Bool(false),
									"serve_stale": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"disable_stale_while_updating": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.Int64Exact(8080),
									}),
									"browser_ttl": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"mode":    knownvalue.StringExact("respect_origin"),
										"default": knownvalue.Null(),
									}),
									"cache": knownvalue.Bool(true),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":       knownvalue.Null(),
										"cache_deception_armor":      knownvalue.Null(),
										"custom_key":                 knownvalue.Null(),
										"ignore_query_strings_order": knownvalue.Null(),
									}),
									"cache_reserve": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"eligible":          knownvalue.Bool(false),
										"minimum_file_size": knownvalue.Null(),
									}),
									"edge_ttl": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"mode":            knownvalue.StringExact("respect_origin"),
										"default":         knownvalue.Null(),
										"status_code_ttl": knownvalue.Null(),
									}),
									"origin_cache_control":       knownvalue.Bool(false),
									"origin_error_page_passthru": knownvalue.Bool(false),
									"read_timeout":               knownvalue.Int64Exact(900),
									"respect_strong_etags":       knownvalue.Bool(false),
									"serve_stale": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"disable_stale_while_updating": knownvalue.Null(),
									}),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_cache_settings"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"additional_cacheable_ports": knownvalue.Null(),
										"browser_ttl": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"mode":    knownvalue.StringExact("override_origin"),
											"default": knownvalue.Int64Exact(60),
										}),
										"cache": knownvalue.Null(),
										"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cache_by_device_type":  knownvalue.Bool(false),
											"cache_deception_armor": knownvalue.Bool(false),
											"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"cookie":       knownvalue.Null(),
												"header":       knownvalue.Null(),
												"host":         knownvalue.Null(),
												"query_string": knownvalue.Null(),
												"user":         knownvalue.Null(),
											}),
											"ignore_query_strings_order": knownvalue.Bool(false),
										}),
										"cache_reserve": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"eligible":          knownvalue.Bool(true),
											"minimum_file_size": knownvalue.Int64Exact(1024),
										}),
										"edge_ttl": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"mode":    knownvalue.StringExact("override_origin"),
											"default": knownvalue.Int64Exact(60),
											"status_code_ttl": knownvalue.ListExact([]knownvalue.Check{
												knownvalue.ObjectExact(map[string]knownvalue.Check{
													"status_code_range": knownvalue.ObjectExact(map[string]knownvalue.Check{
														"from": knownvalue.Int64Exact(500),
														"to":   knownvalue.Null(),
													}),
													"status_code": knownvalue.Null(),
													"value":       knownvalue.Int64Exact(-1),
												}),
												knownvalue.ObjectExact(map[string]knownvalue.Check{
													"status_code_range": knownvalue.ObjectExact(map[string]knownvalue.Check{
														"from": knownvalue.Null(),
														"to":   knownvalue.Int64Exact(199),
													}),
													"status_code": knownvalue.Null(),
													"value":       knownvalue.Int64Exact(0),
												}),
												knownvalue.ObjectExact(map[string]knownvalue.Check{
													"status_code_range": knownvalue.ObjectExact(map[string]knownvalue.Check{
														"from": knownvalue.Int64Exact(200),
														"to":   knownvalue.Int64Exact(399),
													}),
													"status_code": knownvalue.Null(),
													"value":       knownvalue.Int64Exact(1),
												}),
												knownvalue.ObjectExact(map[string]knownvalue.Check{
													"status_code_range": knownvalue.Null(),
													"status_code":       knownvalue.Int64Exact(400),
													"value":             knownvalue.Int64Exact(2),
												}),
											}),
										}),
										"origin_cache_control":       knownvalue.Bool(true),
										"origin_error_page_passthru": knownvalue.Bool(true),
										"read_timeout":               knownvalue.Null(),
										"respect_strong_etags":       knownvalue.Bool(true),
										"serve_stale": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"disable_stale_while_updating": knownvalue.Bool(false),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"mode":    knownvalue.StringExact("override_origin"),
										"default": knownvalue.Int64Exact(60),
									}),
									"cache": knownvalue.Null(),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":  knownvalue.Bool(false),
										"cache_deception_armor": knownvalue.Bool(false),
										"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cookie":       knownvalue.Null(),
											"header":       knownvalue.Null(),
											"host":         knownvalue.Null(),
											"query_string": knownvalue.Null(),
											"user":         knownvalue.Null(),
										}),
										"ignore_query_strings_order": knownvalue.Bool(false),
									}),
									"cache_reserve": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"eligible":          knownvalue.Bool(true),
										"minimum_file_size": knownvalue.Int64Exact(1024),
									}),
									"edge_ttl": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"mode":    knownvalue.StringExact("override_origin"),
										"default": knownvalue.Int64Exact(60),
										"status_code_ttl": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"status_code_range": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"from": knownvalue.Int64Exact(500),
													"to":   knownvalue.Null(),
												}),
												"status_code": knownvalue.Null(),
												"value":       knownvalue.Int64Exact(-1),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"status_code_range": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"from": knownvalue.Null(),
													"to":   knownvalue.Int64Exact(199),
												}),
												"status_code": knownvalue.Null(),
												"value":       knownvalue.Int64Exact(0),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"status_code_range": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"from": knownvalue.Int64Exact(200),
													"to":   knownvalue.Int64Exact(399),
												}),
												"status_code": knownvalue.Null(),
												"value":       knownvalue.Int64Exact(1),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"status_code_range": knownvalue.Null(),
												"status_code":       knownvalue.Int64Exact(400),
												"value":             knownvalue.Int64Exact(2),
											}),
										}),
									}),
									"origin_cache_control":       knownvalue.Bool(true),
									"origin_error_page_passthru": knownvalue.Bool(true),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Bool(true),
									"serve_stale": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"disable_stale_while_updating": knownvalue.Bool(false),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"mode":    knownvalue.StringExact("override_origin"),
										"default": knownvalue.Int64Exact(60),
									}),
									"cache": knownvalue.Null(),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":  knownvalue.Bool(false),
										"cache_deception_armor": knownvalue.Bool(false),
										"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cookie":       knownvalue.Null(),
											"header":       knownvalue.Null(),
											"host":         knownvalue.Null(),
											"query_string": knownvalue.Null(),
											"user":         knownvalue.Null(),
										}),
										"ignore_query_strings_order": knownvalue.Bool(false),
									}),
									"cache_reserve": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"eligible":          knownvalue.Bool(true),
										"minimum_file_size": knownvalue.Int64Exact(1024),
									}),
									"edge_ttl": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"mode":    knownvalue.StringExact("override_origin"),
										"default": knownvalue.Int64Exact(60),
										"status_code_ttl": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"status_code_range": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"from": knownvalue.Int64Exact(500),
													"to":   knownvalue.Null(),
												}),
												"status_code": knownvalue.Null(),
												"value":       knownvalue.Int64Exact(-1),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"status_code_range": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"from": knownvalue.Null(),
													"to":   knownvalue.Int64Exact(199),
												}),
												"status_code": knownvalue.Null(),
												"value":       knownvalue.Int64Exact(0),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"status_code_range": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"from": knownvalue.Int64Exact(200),
													"to":   knownvalue.Int64Exact(399),
												}),
												"status_code": knownvalue.Null(),
												"value":       knownvalue.Int64Exact(1),
											}),
											knownvalue.ObjectExact(map[string]knownvalue.Check{
												"status_code_range": knownvalue.Null(),
												"status_code":       knownvalue.Int64Exact(400),
												"value":             knownvalue.Int64Exact(2),
											}),
										}),
									}),
									"origin_cache_control":       knownvalue.Bool(true),
									"origin_error_page_passthru": knownvalue.Bool(true),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Bool(true),
									"serve_stale": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"disable_stale_while_updating": knownvalue.Bool(false),
									}),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("4.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_cache_settings"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"additional_cacheable_ports": knownvalue.Null(),
										"browser_ttl":                knownvalue.Null(),
										"cache":                      knownvalue.Null(),
										"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cache_by_device_type":  knownvalue.Null(),
											"cache_deception_armor": knownvalue.Bool(true),
											"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"cookie": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"check_presence": knownvalue.Null(),
													"include":        knownvalue.Null(),
												}),
												"header": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"check_presence": knownvalue.Null(),
													"contains":       knownvalue.Null(),
													"exclude_origin": knownvalue.Null(),
													"include":        knownvalue.Null(),
												}),
												"host": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"resolved": knownvalue.Null(),
												}),
												"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"include": knownvalue.Null(),
													"exclude": knownvalue.Null(),
												}),
												"user": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"device_type": knownvalue.Null(),
													"geo":         knownvalue.Null(),
													"lang":        knownvalue.Null(),
												}),
											}),
											"ignore_query_strings_order": knownvalue.Bool(true),
										}),
										"cache_reserve":              knownvalue.Null(),
										"edge_ttl":                   knownvalue.Null(),
										"origin_cache_control":       knownvalue.Null(),
										"origin_error_page_passthru": knownvalue.Null(),
										"read_timeout":               knownvalue.Null(),
										"respect_strong_etags":       knownvalue.Null(),
										"serve_stale": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"disable_stale_while_updating": knownvalue.Bool(true),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl":                knownvalue.Null(),
									"cache":                      knownvalue.Null(),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":  knownvalue.Null(),
										"cache_deception_armor": knownvalue.Bool(true),
										"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cookie": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"check_presence": knownvalue.Null(),
												"include":        knownvalue.Null(),
											}),
											"header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"check_presence": knownvalue.Null(),
												"contains":       knownvalue.Null(),
												"exclude_origin": knownvalue.Null(),
												"include":        knownvalue.Null(),
											}),
											"host": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"resolved": knownvalue.Null(),
											}),
											"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"include": knownvalue.Null(),
												"exclude": knownvalue.Null(),
											}),
											"user": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"device_type": knownvalue.Null(),
												"geo":         knownvalue.Null(),
												"lang":        knownvalue.Null(),
											}),
										}),
										"ignore_query_strings_order": knownvalue.Bool(true),
									}),
									"cache_reserve":              knownvalue.Null(),
									"edge_ttl":                   knownvalue.Null(),
									"origin_cache_control":       knownvalue.Null(),
									"origin_error_page_passthru": knownvalue.Null(),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Null(),
									"serve_stale": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"disable_stale_while_updating": knownvalue.Bool(true),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl":                knownvalue.Null(),
									"cache":                      knownvalue.Null(),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":  knownvalue.Null(),
										"cache_deception_armor": knownvalue.Bool(true),
										"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cookie": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"check_presence": knownvalue.Null(),
												"include":        knownvalue.Null(),
											}),
											"header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"check_presence": knownvalue.Null(),
												"contains":       knownvalue.Null(),
												"exclude_origin": knownvalue.Null(),
												"include":        knownvalue.Null(),
											}),
											"host": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"resolved": knownvalue.Null(),
											}),
											"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"include": knownvalue.Null(),
												"exclude": knownvalue.Null(),
											}),
											"user": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"device_type": knownvalue.Null(),
												"geo":         knownvalue.Null(),
												"lang":        knownvalue.Null(),
											}),
										}),
										"ignore_query_strings_order": knownvalue.Bool(true),
									}),
									"cache_reserve":              knownvalue.Null(),
									"edge_ttl":                   knownvalue.Null(),
									"origin_cache_control":       knownvalue.Null(),
									"origin_error_page_passthru": knownvalue.Null(),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Null(),
									"serve_stale": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"disable_stale_while_updating": knownvalue.Bool(true),
									}),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("5.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_cache_settings"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"additional_cacheable_ports": knownvalue.Null(),
										"browser_ttl":                knownvalue.Null(),
										"cache":                      knownvalue.Null(),
										"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cache_by_device_type":  knownvalue.Null(),
											"cache_deception_armor": knownvalue.Null(),
											"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"cookie": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"check_presence": knownvalue.ListExact([]knownvalue.Check{
														knownvalue.StringExact("myCookie1"),
													}),
													"include": knownvalue.ListExact([]knownvalue.Check{
														knownvalue.StringExact("myCookie2"),
													}),
												}),
												"header": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"check_presence": knownvalue.ListExact([]knownvalue.Check{
														knownvalue.StringExact("my-header-1"),
													}),
													"contains": knownvalue.MapExact(map[string]knownvalue.Check{
														"my-header": knownvalue.ListExact([]knownvalue.Check{
															knownvalue.StringExact("my-header-value"),
														}),
													}),
													"exclude_origin": knownvalue.Bool(false),
													"include": knownvalue.ListExact([]knownvalue.Check{
														knownvalue.StringExact("my-header-2"),
													}),
												}),
												"host": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"resolved": knownvalue.Bool(false),
												}),
												"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"include": knownvalue.ObjectExact(map[string]knownvalue.Check{
														"list": knownvalue.ListExact([]knownvalue.Check{
															knownvalue.StringExact("foo"),
														}),
														"all": knownvalue.Null(),
													}),
													"exclude": knownvalue.Null(),
												}),
												"user": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"device_type": knownvalue.Bool(false),
													"geo":         knownvalue.Bool(false),
													"lang":        knownvalue.Bool(false),
												}),
											}),
											"ignore_query_strings_order": knownvalue.Null(),
										}),
										"cache_reserve":              knownvalue.Null(),
										"edge_ttl":                   knownvalue.Null(),
										"origin_cache_control":       knownvalue.Null(),
										"origin_error_page_passthru": knownvalue.Null(),
										"read_timeout":               knownvalue.Null(),
										"respect_strong_etags":       knownvalue.Null(),
										"serve_stale":                knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl":                knownvalue.Null(),
									"cache":                      knownvalue.Null(),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":  knownvalue.Null(),
										"cache_deception_armor": knownvalue.Null(),
										"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cookie": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"check_presence": knownvalue.ListExact([]knownvalue.Check{
													knownvalue.StringExact("myCookie1"),
												}),
												"include": knownvalue.ListExact([]knownvalue.Check{
													knownvalue.StringExact("myCookie2"),
												}),
											}),
											"header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"check_presence": knownvalue.ListExact([]knownvalue.Check{
													knownvalue.StringExact("my-header-1"),
												}),
												"contains": knownvalue.MapExact(map[string]knownvalue.Check{
													"my-header": knownvalue.ListExact([]knownvalue.Check{
														knownvalue.StringExact("my-header-value"),
													}),
												}),
												"exclude_origin": knownvalue.Bool(false),
												"include": knownvalue.ListExact([]knownvalue.Check{
													knownvalue.StringExact("my-header-2"),
												}),
											}),
											"host": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"resolved": knownvalue.Bool(false),
											}),
											"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"include": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"list": knownvalue.ListExact([]knownvalue.Check{
														knownvalue.StringExact("foo"),
													}),
													"all": knownvalue.Null(),
												}),
												"exclude": knownvalue.Null(),
											}),
											"user": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"device_type": knownvalue.Bool(false),
												"geo":         knownvalue.Bool(false),
												"lang":        knownvalue.Bool(false),
											}),
										}),
										"ignore_query_strings_order": knownvalue.Null(),
									}),
									"cache_reserve":              knownvalue.Null(),
									"edge_ttl":                   knownvalue.Null(),
									"origin_cache_control":       knownvalue.Null(),
									"origin_error_page_passthru": knownvalue.Null(),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Null(),
									"serve_stale":                knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl":                knownvalue.Null(),
									"cache":                      knownvalue.Null(),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":  knownvalue.Null(),
										"cache_deception_armor": knownvalue.Null(),
										"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cookie": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"check_presence": knownvalue.ListExact([]knownvalue.Check{
													knownvalue.StringExact("myCookie1"),
												}),
												"include": knownvalue.ListExact([]knownvalue.Check{
													knownvalue.StringExact("myCookie2"),
												}),
											}),
											"header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"check_presence": knownvalue.ListExact([]knownvalue.Check{
													knownvalue.StringExact("my-header-1"),
												}),
												"contains": knownvalue.MapExact(map[string]knownvalue.Check{
													"my-header": knownvalue.ListExact([]knownvalue.Check{
														knownvalue.StringExact("my-header-value"),
													}),
												}),
												"exclude_origin": knownvalue.Bool(false),
												"include": knownvalue.ListExact([]knownvalue.Check{
													knownvalue.StringExact("my-header-2"),
												}),
											}),
											"host": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"resolved": knownvalue.Bool(false),
											}),
											"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"include": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"list": knownvalue.ListExact([]knownvalue.Check{
														knownvalue.StringExact("foo"),
													}),
													"all": knownvalue.Null(),
												}),
												"exclude": knownvalue.Null(),
											}),
											"user": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"device_type": knownvalue.Bool(false),
												"geo":         knownvalue.Bool(false),
												"lang":        knownvalue.Bool(false),
											}),
										}),
										"ignore_query_strings_order": knownvalue.Null(),
									}),
									"cache_reserve":              knownvalue.Null(),
									"edge_ttl":                   knownvalue.Null(),
									"origin_cache_control":       knownvalue.Null(),
									"origin_error_page_passthru": knownvalue.Null(),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Null(),
									"serve_stale":                knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("6.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_cache_settings"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"additional_cacheable_ports": knownvalue.Null(),
										"browser_ttl":                knownvalue.Null(),
										"cache":                      knownvalue.Null(),
										"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cache_by_device_type":  knownvalue.Null(),
											"cache_deception_armor": knownvalue.Null(),
											"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"cookie": knownvalue.Null(),
												"header": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"check_presence": knownvalue.Null(),
													"contains":       knownvalue.Null(),
													"exclude_origin": knownvalue.Bool(true),
													"include":        knownvalue.Null(),
												}),
												"host": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"resolved": knownvalue.Bool(true),
												}),
												"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"include": knownvalue.ObjectExact(map[string]knownvalue.Check{
														"list": knownvalue.Null(),
														"all":  knownvalue.Bool(true),
													}),
													"exclude": knownvalue.Null(),
												}),
												"user": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"device_type": knownvalue.Bool(true),
													"geo":         knownvalue.Bool(true),
													"lang":        knownvalue.Bool(true),
												}),
											}),
											"ignore_query_strings_order": knownvalue.Null(),
										}),
										"cache_reserve":              knownvalue.Null(),
										"edge_ttl":                   knownvalue.Null(),
										"origin_cache_control":       knownvalue.Null(),
										"origin_error_page_passthru": knownvalue.Null(),
										"read_timeout":               knownvalue.Null(),
										"respect_strong_etags":       knownvalue.Null(),
										"serve_stale":                knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl":                knownvalue.Null(),
									"cache":                      knownvalue.Null(),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":  knownvalue.Null(),
										"cache_deception_armor": knownvalue.Null(),
										"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cookie": knownvalue.Null(),
											"header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"check_presence": knownvalue.Null(),
												"contains":       knownvalue.Null(),
												"exclude_origin": knownvalue.Bool(true),
												"include":        knownvalue.Null(),
											}),
											"host": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"resolved": knownvalue.Bool(true),
											}),
											"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"include": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"list": knownvalue.Null(),
													"all":  knownvalue.Bool(true),
												}),
												"exclude": knownvalue.Null(),
											}),
											"user": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"device_type": knownvalue.Bool(true),
												"geo":         knownvalue.Bool(true),
												"lang":        knownvalue.Bool(true),
											}),
										}),
										"ignore_query_strings_order": knownvalue.Null(),
									}),
									"cache_reserve":              knownvalue.Null(),
									"edge_ttl":                   knownvalue.Null(),
									"origin_cache_control":       knownvalue.Null(),
									"origin_error_page_passthru": knownvalue.Null(),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Null(),
									"serve_stale":                knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl":                knownvalue.Null(),
									"cache":                      knownvalue.Null(),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":  knownvalue.Null(),
										"cache_deception_armor": knownvalue.Null(),
										"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cookie": knownvalue.Null(),
											"header": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"check_presence": knownvalue.Null(),
												"contains":       knownvalue.Null(),
												"exclude_origin": knownvalue.Bool(true),
												"include":        knownvalue.Null(),
											}),
											"host": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"resolved": knownvalue.Bool(true),
											}),
											"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"include": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"list": knownvalue.Null(),
													"all":  knownvalue.Bool(true),
												}),
												"exclude": knownvalue.Null(),
											}),
											"user": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"device_type": knownvalue.Bool(true),
												"geo":         knownvalue.Bool(true),
												"lang":        knownvalue.Bool(true),
											}),
										}),
										"ignore_query_strings_order": knownvalue.Null(),
									}),
									"cache_reserve":              knownvalue.Null(),
									"edge_ttl":                   knownvalue.Null(),
									"origin_cache_control":       knownvalue.Null(),
									"origin_error_page_passthru": knownvalue.Null(),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Null(),
									"serve_stale":                knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("7.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_cache_settings"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"additional_cacheable_ports": knownvalue.Null(),
										"browser_ttl":                knownvalue.Null(),
										"cache":                      knownvalue.Null(),
										"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cache_by_device_type":  knownvalue.Bool(true),
											"cache_deception_armor": knownvalue.Null(),
											"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"cookie": knownvalue.Null(),
												"header": knownvalue.Null(),
												"host":   knownvalue.Null(),
												"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"include": knownvalue.Null(),
													"exclude": knownvalue.ObjectExact(map[string]knownvalue.Check{
														"list": knownvalue.ListExact([]knownvalue.Check{
															knownvalue.StringExact("foo"),
														}),
														"all": knownvalue.Null(),
													}),
												}),
												"user": knownvalue.Null(),
											}),
											"ignore_query_strings_order": knownvalue.Null(),
										}),
										"cache_reserve":              knownvalue.Null(),
										"edge_ttl":                   knownvalue.Null(),
										"origin_cache_control":       knownvalue.Null(),
										"origin_error_page_passthru": knownvalue.Null(),
										"read_timeout":               knownvalue.Null(),
										"respect_strong_etags":       knownvalue.Null(),
										"serve_stale":                knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl":                knownvalue.Null(),
									"cache":                      knownvalue.Null(),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":  knownvalue.Bool(true),
										"cache_deception_armor": knownvalue.Null(),
										"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cookie": knownvalue.Null(),
											"header": knownvalue.Null(),
											"host":   knownvalue.Null(),
											"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"include": knownvalue.Null(),
												"exclude": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"list": knownvalue.ListExact([]knownvalue.Check{
														knownvalue.StringExact("foo"),
													}),
													"all": knownvalue.Null(),
												}),
											}),
											"user": knownvalue.Null(),
										}),
										"ignore_query_strings_order": knownvalue.Null(),
									}),
									"cache_reserve":              knownvalue.Null(),
									"edge_ttl":                   knownvalue.Null(),
									"origin_cache_control":       knownvalue.Null(),
									"origin_error_page_passthru": knownvalue.Null(),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Null(),
									"serve_stale":                knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl":                knownvalue.Null(),
									"cache":                      knownvalue.Null(),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":  knownvalue.Bool(true),
										"cache_deception_armor": knownvalue.Null(),
										"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cookie": knownvalue.Null(),
											"header": knownvalue.Null(),
											"host":   knownvalue.Null(),
											"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"include": knownvalue.Null(),
												"exclude": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"list": knownvalue.ListExact([]knownvalue.Check{
														knownvalue.StringExact("foo"),
													}),
													"all": knownvalue.Null(),
												}),
											}),
											"user": knownvalue.Null(),
										}),
										"ignore_query_strings_order": knownvalue.Null(),
									}),
									"cache_reserve":              knownvalue.Null(),
									"edge_ttl":                   knownvalue.Null(),
									"origin_cache_control":       knownvalue.Null(),
									"origin_error_page_passthru": knownvalue.Null(),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Null(),
									"serve_stale":                knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("8.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_cache_settings"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"additional_cacheable_ports": knownvalue.Null(),
										"browser_ttl":                knownvalue.Null(),
										"cache":                      knownvalue.Null(),
										"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cache_by_device_type":  knownvalue.Null(),
											"cache_deception_armor": knownvalue.Null(),
											"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"cookie": knownvalue.Null(),
												"header": knownvalue.Null(),
												"host":   knownvalue.Null(),
												"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"include": knownvalue.Null(),
													"exclude": knownvalue.ObjectExact(map[string]knownvalue.Check{
														"list": knownvalue.Null(),
														"all":  knownvalue.Bool(true),
													}),
												}),
												"user": knownvalue.Null(),
											}),
											"ignore_query_strings_order": knownvalue.Null(),
										}),
										"cache_reserve":              knownvalue.Null(),
										"edge_ttl":                   knownvalue.Null(),
										"origin_cache_control":       knownvalue.Null(),
										"origin_error_page_passthru": knownvalue.Null(),
										"read_timeout":               knownvalue.Null(),
										"respect_strong_etags":       knownvalue.Null(),
										"serve_stale":                knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl":                knownvalue.Null(),
									"cache":                      knownvalue.Null(),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":  knownvalue.Null(),
										"cache_deception_armor": knownvalue.Null(),
										"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cookie": knownvalue.Null(),
											"header": knownvalue.Null(),
											"host":   knownvalue.Null(),
											"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"include": knownvalue.Null(),
												"exclude": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"list": knownvalue.Null(),
													"all":  knownvalue.Bool(true),
												}),
											}),
											"user": knownvalue.Null(),
										}),
										"ignore_query_strings_order": knownvalue.Null(),
									}),
									"cache_reserve":              knownvalue.Null(),
									"edge_ttl":                   knownvalue.Null(),
									"origin_cache_control":       knownvalue.Null(),
									"origin_error_page_passthru": knownvalue.Null(),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Null(),
									"serve_stale":                knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_cache_settings"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"additional_cacheable_ports": knownvalue.Null(),
									"browser_ttl":                knownvalue.Null(),
									"cache":                      knownvalue.Null(),
									"cache_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"cache_by_device_type":  knownvalue.Null(),
										"cache_deception_armor": knownvalue.Null(),
										"custom_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"cookie": knownvalue.Null(),
											"header": knownvalue.Null(),
											"host":   knownvalue.Null(),
											"query_string": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"include": knownvalue.Null(),
												"exclude": knownvalue.ObjectExact(map[string]knownvalue.Check{
													"list": knownvalue.Null(),
													"all":  knownvalue.Bool(true),
												}),
											}),
											"user": knownvalue.Null(),
										}),
										"ignore_query_strings_order": knownvalue.Null(),
									}),
									"cache_reserve":              knownvalue.Null(),
									"edge_ttl":                   knownvalue.Null(),
									"origin_cache_control":       knownvalue.Null(),
									"origin_error_page_passthru": knownvalue.Null(),
									"read_timeout":               knownvalue.Null(),
									"respect_strong_etags":       knownvalue.Null(),
									"serve_stale":                knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_SetConfigRules(t *testing.T) {
	t.Skip(`FIXME: skip test due to feature deprecations; mirage and disable_apps are deprecated, but have suddenly EOL'd`)
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_config"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"automatic_https_rewrites": knownvalue.Bool(false),
										"autominify":               knownvalue.Null(),
										"bic":                      knownvalue.Null(),
										"disable_apps":             knownvalue.Null(),
										"disable_rum":              knownvalue.Null(),
										"disable_zaraz":            knownvalue.Null(),
										"email_obfuscation":        knownvalue.Null(),
										"fonts":                    knownvalue.Null(),
										"hotlink_protection":       knownvalue.Null(),
										"mirage":                   knownvalue.Null(),
										"opportunistic_encryption": knownvalue.Null(),
										"polish":                   knownvalue.Null(),
										"rocket_loader":            knownvalue.Null(),
										"security_level":           knownvalue.Null(),
										"server_side_excludes":     knownvalue.Null(),
										"ssl":                      knownvalue.Null(),
										"sxg":                      knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"automatic_https_rewrites": knownvalue.Bool(false),
									"autominify":               knownvalue.Null(),
									"bic":                      knownvalue.Null(),
									"disable_apps":             knownvalue.Null(),
									"disable_rum":              knownvalue.Null(),
									"disable_zaraz":            knownvalue.Null(),
									"email_obfuscation":        knownvalue.Null(),
									"fonts":                    knownvalue.Null(),
									"hotlink_protection":       knownvalue.Null(),
									"mirage":                   knownvalue.Null(),
									"opportunistic_encryption": knownvalue.Null(),
									"polish":                   knownvalue.Null(),
									"rocket_loader":            knownvalue.Null(),
									"security_level":           knownvalue.Null(),
									"server_side_excludes":     knownvalue.Null(),
									"ssl":                      knownvalue.Null(),
									"sxg":                      knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"automatic_https_rewrites": knownvalue.Bool(false),
									"autominify":               knownvalue.Null(),
									"bic":                      knownvalue.Null(),
									"disable_apps":             knownvalue.Null(),
									"disable_rum":              knownvalue.Null(),
									"disable_zaraz":            knownvalue.Null(),
									"email_obfuscation":        knownvalue.Null(),
									"fonts":                    knownvalue.Null(),
									"hotlink_protection":       knownvalue.Null(),
									"mirage":                   knownvalue.Null(),
									"opportunistic_encryption": knownvalue.Null(),
									"polish":                   knownvalue.Null(),
									"rocket_loader":            knownvalue.Null(),
									"security_level":           knownvalue.Null(),
									"server_side_excludes":     knownvalue.Null(),
									"ssl":                      knownvalue.Null(),
									"sxg":                      knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_config"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"automatic_https_rewrites": knownvalue.Bool(true),
										"autominify": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"css":  knownvalue.Bool(false),
											"html": knownvalue.Bool(false),
											"js":   knownvalue.Bool(false),
										}),
										"bic":                      knownvalue.Bool(false),
										"disable_apps":             knownvalue.Bool(true),
										"disable_rum":              knownvalue.Bool(true),
										"disable_zaraz":            knownvalue.Bool(true),
										"email_obfuscation":        knownvalue.Bool(false),
										"fonts":                    knownvalue.Bool(false),
										"hotlink_protection":       knownvalue.Bool(false),
										"mirage":                   knownvalue.Bool(false),
										"opportunistic_encryption": knownvalue.Bool(false),
										"polish":                   knownvalue.StringExact("off"),
										"rocket_loader":            knownvalue.Bool(false),
										"security_level":           knownvalue.StringExact("off"),
										"server_side_excludes":     knownvalue.Bool(false),
										"ssl":                      knownvalue.StringExact("off"),
										"sxg":                      knownvalue.Bool(false),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"automatic_https_rewrites": knownvalue.Bool(true),
									"autominify": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"css":  knownvalue.Bool(false),
										"html": knownvalue.Bool(false),
										"js":   knownvalue.Bool(false),
									}),
									"bic":                      knownvalue.Bool(false),
									"disable_apps":             knownvalue.Bool(true),
									"disable_rum":              knownvalue.Bool(true),
									"disable_zaraz":            knownvalue.Bool(true),
									"email_obfuscation":        knownvalue.Bool(false),
									"fonts":                    knownvalue.Bool(false),
									"hotlink_protection":       knownvalue.Bool(false),
									"mirage":                   knownvalue.Bool(false),
									"opportunistic_encryption": knownvalue.Bool(false),
									"polish":                   knownvalue.StringExact("off"),
									"rocket_loader":            knownvalue.Bool(false),
									"security_level":           knownvalue.StringExact("off"),
									"server_side_excludes":     knownvalue.Bool(false),
									"ssl":                      knownvalue.StringExact("off"),
									"sxg":                      knownvalue.Bool(false),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"automatic_https_rewrites": knownvalue.Bool(true),
									"autominify": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"css":  knownvalue.Bool(false),
										"html": knownvalue.Bool(false),
										"js":   knownvalue.Bool(false),
									}),
									"bic":                      knownvalue.Bool(false),
									"disable_apps":             knownvalue.Bool(true),
									"disable_rum":              knownvalue.Bool(true),
									"disable_zaraz":            knownvalue.Bool(true),
									"email_obfuscation":        knownvalue.Bool(false),
									"fonts":                    knownvalue.Bool(false),
									"hotlink_protection":       knownvalue.Bool(false),
									"mirage":                   knownvalue.Bool(false),
									"opportunistic_encryption": knownvalue.Bool(false),
									"polish":                   knownvalue.StringExact("off"),
									"rocket_loader":            knownvalue.Bool(false),
									"security_level":           knownvalue.StringExact("off"),
									"server_side_excludes":     knownvalue.Bool(false),
									"ssl":                      knownvalue.StringExact("off"),
									"sxg":                      knownvalue.Bool(false),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_config"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"automatic_https_rewrites": knownvalue.Null(),
										"autominify": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"css":  knownvalue.Bool(false),
											"html": knownvalue.Bool(false),
											"js":   knownvalue.Bool(false),
										}),
										"bic":                      knownvalue.Bool(true),
										"disable_apps":             knownvalue.Null(),
										"disable_rum":              knownvalue.Null(),
										"disable_zaraz":            knownvalue.Null(),
										"email_obfuscation":        knownvalue.Bool(true),
										"fonts":                    knownvalue.Bool(true),
										"hotlink_protection":       knownvalue.Bool(true),
										"mirage":                   knownvalue.Bool(true),
										"opportunistic_encryption": knownvalue.Bool(true),
										"polish":                   knownvalue.Null(),
										"rocket_loader":            knownvalue.Bool(true),
										"security_level":           knownvalue.Null(),
										"server_side_excludes":     knownvalue.Bool(true),
										"ssl":                      knownvalue.Null(),
										"sxg":                      knownvalue.Bool(true),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"automatic_https_rewrites": knownvalue.Null(),
									"autominify": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"css":  knownvalue.Bool(false),
										"html": knownvalue.Bool(false),
										"js":   knownvalue.Bool(false),
									}),
									"bic":                      knownvalue.Bool(true),
									"disable_apps":             knownvalue.Null(),
									"disable_rum":              knownvalue.Null(),
									"disable_zaraz":            knownvalue.Null(),
									"email_obfuscation":        knownvalue.Bool(true),
									"fonts":                    knownvalue.Bool(true),
									"hotlink_protection":       knownvalue.Bool(true),
									"mirage":                   knownvalue.Bool(true),
									"opportunistic_encryption": knownvalue.Bool(true),
									"polish":                   knownvalue.Null(),
									"rocket_loader":            knownvalue.Bool(true),
									"security_level":           knownvalue.Null(),
									"server_side_excludes":     knownvalue.Bool(true),
									"ssl":                      knownvalue.Null(),
									"sxg":                      knownvalue.Bool(true),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"automatic_https_rewrites": knownvalue.Null(),
									"autominify": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"css":  knownvalue.Bool(false),
										"html": knownvalue.Bool(false),
										"js":   knownvalue.Bool(false),
									}),
									"bic":                      knownvalue.Bool(true),
									"disable_apps":             knownvalue.Null(),
									"disable_rum":              knownvalue.Null(),
									"disable_zaraz":            knownvalue.Null(),
									"email_obfuscation":        knownvalue.Bool(true),
									"fonts":                    knownvalue.Bool(true),
									"hotlink_protection":       knownvalue.Bool(true),
									"mirage":                   knownvalue.Bool(true),
									"opportunistic_encryption": knownvalue.Bool(true),
									"polish":                   knownvalue.Null(),
									"rocket_loader":            knownvalue.Bool(true),
									"security_level":           knownvalue.Null(),
									"server_side_excludes":     knownvalue.Bool(true),
									"ssl":                      knownvalue.Null(),
									"sxg":                      knownvalue.Bool(true),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("4.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_config"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"automatic_https_rewrites": knownvalue.Null(),
										"autominify": knownvalue.ObjectExact(map[string]knownvalue.Check{
											"css":  knownvalue.Bool(true),
											"html": knownvalue.Bool(true),
											"js":   knownvalue.Bool(true),
										}),
										"bic":                      knownvalue.Null(),
										"disable_apps":             knownvalue.Null(),
										"disable_rum":              knownvalue.Null(),
										"disable_zaraz":            knownvalue.Null(),
										"email_obfuscation":        knownvalue.Null(),
										"fonts":                    knownvalue.Null(),
										"hotlink_protection":       knownvalue.Null(),
										"mirage":                   knownvalue.Null(),
										"opportunistic_encryption": knownvalue.Null(),
										"polish":                   knownvalue.Null(),
										"rocket_loader":            knownvalue.Null(),
										"security_level":           knownvalue.Null(),
										"server_side_excludes":     knownvalue.Null(),
										"ssl":                      knownvalue.Null(),
										"sxg":                      knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"automatic_https_rewrites": knownvalue.Null(),
									"autominify": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"css":  knownvalue.Bool(true),
										"html": knownvalue.Bool(true),
										"js":   knownvalue.Bool(true),
									}),
									"bic":                      knownvalue.Null(),
									"disable_apps":             knownvalue.Null(),
									"disable_rum":              knownvalue.Null(),
									"disable_zaraz":            knownvalue.Null(),
									"email_obfuscation":        knownvalue.Null(),
									"fonts":                    knownvalue.Null(),
									"hotlink_protection":       knownvalue.Null(),
									"mirage":                   knownvalue.Null(),
									"opportunistic_encryption": knownvalue.Null(),
									"polish":                   knownvalue.Null(),
									"rocket_loader":            knownvalue.Null(),
									"security_level":           knownvalue.Null(),
									"server_side_excludes":     knownvalue.Null(),
									"ssl":                      knownvalue.Null(),
									"sxg":                      knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"automatic_https_rewrites": knownvalue.Null(),
									"autominify": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"css":  knownvalue.Bool(true),
										"html": knownvalue.Bool(true),
										"js":   knownvalue.Bool(true),
									}),
									"bic":                      knownvalue.Null(),
									"disable_apps":             knownvalue.Null(),
									"disable_rum":              knownvalue.Null(),
									"disable_zaraz":            knownvalue.Null(),
									"email_obfuscation":        knownvalue.Null(),
									"fonts":                    knownvalue.Null(),
									"hotlink_protection":       knownvalue.Null(),
									"mirage":                   knownvalue.Null(),
									"opportunistic_encryption": knownvalue.Null(),
									"polish":                   knownvalue.Null(),
									"rocket_loader":            knownvalue.Null(),
									"security_level":           knownvalue.Null(),
									"server_side_excludes":     knownvalue.Null(),
									"ssl":                      knownvalue.Null(),
									"sxg":                      knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("5.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_config"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"request_body_buffering":  knownvalue.StringExact("none"),
										"response_body_buffering": knownvalue.StringExact("none"),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"request_body_buffering":  knownvalue.StringExact("none"),
									"response_body_buffering": knownvalue.StringExact("none"),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"request_body_buffering":  knownvalue.StringExact("none"),
									"response_body_buffering": knownvalue.StringExact("none"),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("6.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_config"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"request_body_buffering":  knownvalue.StringExact("standard"),
										"response_body_buffering": knownvalue.StringExact("standard"),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"request_body_buffering":  knownvalue.StringExact("standard"),
									"response_body_buffering": knownvalue.StringExact("standard"),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"request_body_buffering":  knownvalue.StringExact("standard"),
									"response_body_buffering": knownvalue.StringExact("standard"),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("7.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("set_config"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"request_body_buffering":  knownvalue.StringExact("full"),
										"response_body_buffering": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"request_body_buffering":  knownvalue.StringExact("full"),
									"response_body_buffering": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("set_config"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"request_body_buffering":  knownvalue.StringExact("full"),
									"response_body_buffering": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareRuleset_SkipRules(t *testing.T) {
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
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("skip"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										// "phase":  knownvalue.StringExact("current"),
										"phases": knownvalue.Null(),
										"products": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.StringExact("bic"),
										}),
										"rules":    knownvalue.Null(),
										"ruleset":  knownvalue.Null(),
										"rulesets": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("skip"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									// "phase":  knownvalue.StringExact("current"),
									"phases": knownvalue.Null(),
									"products": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("bic"),
									}),
									"rules":    knownvalue.Null(),
									"ruleset":  knownvalue.Null(),
									"rulesets": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("skip"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									// "phase":  knownvalue.StringExact("current"),
									"phases": knownvalue.Null(),
									"products": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("bic"),
									}),
									"rules":    knownvalue.Null(),
									"ruleset":  knownvalue.Null(),
									"rulesets": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("2.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("skip"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										// "phase": knownvalue.Null(),
										"phases": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.StringExact("http_request_firewall_managed"),
										}),
										"products": knownvalue.Null(),
										"rules":    knownvalue.Null(),
										"ruleset":  knownvalue.StringExact("current"),
										"rulesets": knownvalue.Null(),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("skip"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									// "phase": knownvalue.Null(),
									"phases": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("http_request_firewall_managed"),
									}),
									"products": knownvalue.Null(),
									"rules":    knownvalue.Null(),
									"ruleset":  knownvalue.StringExact("current"),
									"rulesets": knownvalue.Null(),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("skip"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									// "phase": knownvalue.Null(),
									"phases": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("http_request_firewall_managed"),
									}),
									"products": knownvalue.Null(),
									"rules":    knownvalue.Null(),
									"ruleset":  knownvalue.StringExact("current"),
									"rulesets": knownvalue.Null(),
								}),
							}),
						}),
					),
				},
			},
			{
				ConfigFile:      config.TestNameFile("3.tf"),
				ConfigVariables: configVariables,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"cloudflare_ruleset.my_ruleset",
							plancheck.ResourceActionReplace,
						),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("rules"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"action": knownvalue.StringExact("skip"),
									"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										// "phase":    knownvalue.Null(),
										"phases":   knownvalue.Null(),
										"products": knownvalue.Null(),
										"rules": knownvalue.MapExact(map[string]knownvalue.Check{
											"4814384a9e5d4991b9815dcfc25d2f1f": knownvalue.ListExact([]knownvalue.Check{
												knownvalue.StringExact("04116d14d7524986ba314d11c8a41e11"),
											}),
										}),
										"ruleset": knownvalue.Null(),
										"rulesets": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
										}),
									}),
								}),
							}),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("skip"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									// "phase":    knownvalue.Null(),
									"phases":   knownvalue.Null(),
									"products": knownvalue.Null(),
									"rules": knownvalue.MapExact(map[string]knownvalue.Check{
										"4814384a9e5d4991b9815dcfc25d2f1f": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.StringExact("04116d14d7524986ba314d11c8a41e11"),
										}),
									}),
									"ruleset": knownvalue.Null(),
									"rulesets": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
									}),
								}),
							}),
						}),
					),
					statecheck.ExpectKnownValue(
						"data.cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("rules"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"action": knownvalue.StringExact("skip"),
								"action_parameters": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									// "phase":    knownvalue.Null(),
									"phases":   knownvalue.Null(),
									"products": knownvalue.Null(),
									"rules": knownvalue.MapExact(map[string]knownvalue.Check{
										"4814384a9e5d4991b9815dcfc25d2f1f": knownvalue.ListExact([]knownvalue.Check{
											knownvalue.StringExact("04116d14d7524986ba314d11c8a41e11"),
										}),
									}),
									"ruleset": knownvalue.Null(),
									"rulesets": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("4814384a9e5d4991b9815dcfc25d2f1f"),
									}),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}
