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
											"categories":        knownvalue.ListExact([]knownvalue.Check{}),
											"enabled":           knownvalue.Null(),
											"rules":             knownvalue.ListExact([]knownvalue.Check{}),
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
										"categories":        knownvalue.ListExact([]knownvalue.Check{}),
										"enabled":           knownvalue.Null(),
										"rules":             knownvalue.ListExact([]knownvalue.Check{}),
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
											"categories":        knownvalue.ListExact([]knownvalue.Check{}),
											"enabled":           knownvalue.Bool(false),
											"rules":             knownvalue.ListExact([]knownvalue.Check{}),
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
										"categories":        knownvalue.ListExact([]knownvalue.Check{}),
										"enabled":           knownvalue.Bool(false),
										"rules":             knownvalue.ListExact([]knownvalue.Check{}),
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
				},
			},
		},
	})
}
