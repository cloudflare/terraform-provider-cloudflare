package ruleset_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var configVariables = config.Variables{
	"account_id": config.StringVariable(os.Getenv("CLOUDFLARE_ACCOUNT_ID")),
	"zone_id":    config.StringVariable(os.Getenv("CLOUDFLARE_ZONE_ID")),
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
							tfjsonpath.New("description"),
							knownvalue.StringExact("My updated ruleset description"),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"cloudflare_ruleset.my_ruleset",
						tfjsonpath.New("description"),
						knownvalue.StringExact("My updated ruleset description"),
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
									"description": knownvalue.StringExact("My updated rule description"),
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
								"description": knownvalue.StringExact("My updated rule description"),
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
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
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
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
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
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
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
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
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
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
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
									"expression": knownvalue.StringExact("ip.src eq 1.1.1.1"),
									"ref":        knownvalue.StringExact("one"),
								}),
								knownvalue.ObjectPartial(map[string]knownvalue.Check{
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
				},
			},
		})
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
				},
			},
		},
	})
}
