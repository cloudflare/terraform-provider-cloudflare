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
						plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
						plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionReplace),
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
						plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
						plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionReplace),
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
						plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
						plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionReplace),
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
						plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
						plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(
							"cloudflare_ruleset.my_ruleset",
							tfjsonpath.New("name"),
							knownvalue.StringExact("My updated ruleset description"),
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
						plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
						plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
						plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionUpdate),
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
								plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
								plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionUpdate),
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
								plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
								plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionUpdate),
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
								plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
								plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionUpdate),
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
							plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
							plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionUpdate),
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
							plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
							plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionUpdate),
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
							plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
							plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionUpdate),
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
							plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
							plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionUpdate),
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
							plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
							plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionUpdate),
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
							plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionCreate),
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
							plancheck.ExpectResourceAction("cloudflare_ruleset.my_ruleset", plancheck.ResourceActionUpdate),
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
