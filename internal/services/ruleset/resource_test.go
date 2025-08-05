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
	"zone_id": config.StringVariable(os.Getenv("CLOUDFLARE_ZONE_ID")),
}

func TestAccCloudflareRuleset_RuleRef(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { acctest.TestAccPreCheck(t) },
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile:      config.TestNameFile("before.tf"),
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
					ConfigFile:      config.TestNameFile("after.tf"),
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
					ConfigFile:      config.TestNameFile("before.tf"),
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
					ConfigFile:      config.TestNameFile("after.tf"),
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
					ConfigFile:      config.TestNameFile("before.tf"),
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
					ConfigFile:      config.TestNameFile("after.tf"),
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
					ConfigFile:      config.TestNameFile("before.tf"),
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
					ConfigFile:      config.TestNameFile("after.tf"),
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
					ConfigFile:      config.TestNameFile("before.tf"),
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
					ConfigFile:      config.TestNameFile("after.tf"),
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
					ConfigFile:      config.TestNameFile("before.tf"),
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
					ConfigFile:      config.TestNameFile("after.tf"),
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
