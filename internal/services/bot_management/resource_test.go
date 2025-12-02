package bot_management_test

import (
	"context"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_bot_management", &resource.Sweeper{
		Name: "cloudflare_bot_management",
		F:    testSweepCloudflareBotManagement,
	})
}

func testSweepCloudflareBotManagement(r string) error {
	ctx := context.Background()
	// Bot Management is a zone-level configuration for bot detection settings.
	// It's a singleton configuration per zone, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Bot Management doesn't require sweeping (zone settings)")
	return nil
}

func TestAccCloudflareBotManagement_SBFM(t *testing.T) {
	t.Skip("needs SBFM entitlements to run")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")

	sbfmConfig := cloudflare.BotManagement{
		EnableJS:                     cloudflare.BoolPtr(true),
		SBFMDefinitelyAutomated:      cloudflare.StringPtr("managed_challenge"),
		SBFMLikelyAutomated:          cloudflare.StringPtr("block"),
		SBFMVerifiedBots:             cloudflare.StringPtr("allow"),
		SBFMStaticResourceProtection: cloudflare.BoolPtr(false),
		OptimizeWordpress:            cloudflare.BoolPtr(true),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementSBFM(rnd, zoneID, sbfmConfig),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_definitely_automated"), knownvalue.StringExact("managed_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_likely_automated"), knownvalue.StringExact("block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_verified_bots"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_static_resource_protection"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("optimize_wordpress"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareBotManagement_Unentitled(t *testing.T) {
	t.Skip("Test expects entitlement error but test zone entitlements allow the configuration")

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	bmEntConfig := cloudflare.BotManagement{
		EnableJS:             cloudflare.BoolPtr(true),
		SuppressSessionScore: cloudflare.BoolPtr(false),
		AutoUpdateModel:      cloudflare.BoolPtr(false),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testCloudflareBotManagementEntSubscription(rnd, zoneID, bmEntConfig),
				ExpectError: regexp.MustCompile("zone not entitled to disable"),
			},
		},
	})
}

func TestAccCloudflareBotManagement_EnableJS(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementEnableJS(rnd, zoneID, false),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(false)),
				},
			},
			{
				Config: testCloudflareBotManagementEnableJS(rnd, zoneID, true),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ai_bots_protection",
					"crawler_protection",
					"stale_zone_configuration",
				},
			},
		},
	})
}

func TestAccCloudflareBotManagement_SuppressSessionScore(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementSuppressSessionScore(rnd, zoneID, false),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("suppress_session_score"), knownvalue.Bool(false)),
				},
			},
			{
				Config: testCloudflareBotManagementSuppressSessionScore(rnd, zoneID, true),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("suppress_session_score"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ai_bots_protection",
					"crawler_protection",
					"stale_zone_configuration",
				},
			},
		},
	})
}

func TestAccCloudflareBotManagement_AutoUpdateModel_Unentitled(t *testing.T) {
	t.Skip("needs SBFM entitlements to run")
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testCloudflareBotManagementAutoUpdateModel(rnd, zoneID, false),
				ExpectError: regexp.MustCompile("zone not entitled to disable"),
			},
		},
	})
}

func TestAccCloudflareBotManagement_AIBotsProtection(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementAIBotsProtection(rnd, zoneID, "block"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_bots_protection"), knownvalue.StringExact("block")),
				},
			},
			{
				Config: testCloudflareBotManagementAIBotsProtection(rnd, zoneID, "disabled"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_bots_protection"), knownvalue.StringExact("disabled")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareBotManagement_StateConsistency(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementStateConsistency(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fight_mode"), knownvalue.Bool(false)),
				},
			},
			{
				Config: testCloudflareBotManagementStateConsistency(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fight_mode"), knownvalue.Bool(false)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ai_bots_protection",
					"crawler_protection",
					"stale_zone_configuration",
					"fight_mode",
				},
			},
		},
	})
}

func TestAccCloudflareBotManagement_FieldPermutations_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementBasicPermutation(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("suppress_session_score"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_bots_protection"), knownvalue.StringExact("block")),
				},
			},
			{
				Config: testCloudflareBotManagementUpdatedPermutation(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("suppress_session_score"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_bots_protection"), knownvalue.StringExact("disabled")),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareBotManagement_FieldPermutations_SBFM(t *testing.T) {
	t.Skip("needs SBFM entitlements to run")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementSBFMPermutation1(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_definitely_automated"), knownvalue.StringExact("block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_likely_automated"), knownvalue.StringExact("managed_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_verified_bots"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_static_resource_protection"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testCloudflareBotManagementSBFMPermutation2(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_definitely_automated"), knownvalue.StringExact("managed_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_likely_automated"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_verified_bots"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_static_resource_protection"), knownvalue.Bool(false)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareBotManagement_SuppressSessionScore_Issue5519(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementIssue5519(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testCloudflareBotManagementIssue5519(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccCloudflareBotManagement_Issue5519_LifecycleIgnoreChanges(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementIssue5519Lifecycle(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testCloudflareBotManagementIssue5519Lifecycle(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccCloudflareBotManagement_Issue5519_MinimalConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementMinimalConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				},
			},
			{
				Config: testCloudflareBotManagementMinimalConfig(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccCloudflareBotManagement_Issue5519_ExistingResourceDrift(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementIssue5519ExistingResourceConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_bots_protection"), knownvalue.StringExact("block")),
				},
			},
			{
				Config: testCloudflareBotManagementIssue5519ExistingResourceConfig(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareBotManagement_Issue5519_PlanMismatch(t *testing.T) {
	t.Skip("needs SBFM entitlements to run")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementIssue5519PlanMismatch(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_definitely_automated"), knownvalue.StringExact("block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sbfm_verified_bots"), knownvalue.StringExact("allow")),
				},
			},
			{
				Config: testCloudflareBotManagementIssue5519PlanMismatch(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccCloudflareBotManagement_Issue5519_AutoUpdateModelNull(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementIssue5519AutoUpdate(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testCloudflareBotManagementIssue5519AutoUpdate(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccCloudflareBotManagement_Issue5519_NullFieldDrift(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementIssue5519NullFields(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testCloudflareBotManagementIssue5519NullFields(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccCloudflareBotManagement_Issue5519_REPRODUCED(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementIssue5519Reproduce(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				},
			},
			{
				Config: testCloudflareBotManagementIssue5519Reproduce(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccCloudflareBotManagement_ComputedFields(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementComputedFields(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("using_latest_model"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_robots_txt_managed"), knownvalue.Bool(false)),
				},
			},
			{
				Config: testCloudflareBotManagementComputedFields(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccCloudflareBotManagement_AutoUpdateModelStateConsistency(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementAutoUpdateModelStateConsistency(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fight_mode"), knownvalue.Bool(false)),
				},
			},
			{
				Config: testCloudflareBotManagementAutoUpdateModelStateConsistency(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fight_mode"), knownvalue.Bool(false)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ai_bots_protection",
					"crawler_protection",
					"stale_zone_configuration",
					"fight_mode",
				},
			},
		},
	})
}

func TestAccCloudflareBotManagement_EnableJSAutoUpdateSuppression(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementEnableJSAutoUpdateSupression(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(false)),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("enable_js"),
							knownvalue.Bool(false),
						),
						plancheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("auto_update_model"),
							knownvalue.Bool(true),
						),
						plancheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("suppress_session_score"),
							knownvalue.Bool(false),
						),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ai_bots_protection",
					"crawler_protection",
					"stale_zone_configuration",
				},
			},
		},
	})
}

func testCloudflareBotManagementAIBotsProtection(resourceName, zoneID string, aiBotsProtection string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementaibotsprotection.tf", resourceName, zoneID, aiBotsProtection)
}

func testCloudflareBotManagementStateConsistency(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementstateconsistency.tf", resourceName, zoneID)
}

func testCloudflareBotManagementAutoUpdateModelStateConsistency(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementautoupdatemodelstateconsistency.tf", resourceName, zoneID)
}

func TestAccCloudflareBotManagement_AutoUpdateModelStateConsistency_UserZone(t *testing.T) {
	// Test against zones with different entitlement levels:
	// - CLOUDFLARE_ZONE_ID: Entitled zone that returns all fields in API responses
	// - CLOUDFLARE_ALT_ZONE_ID: Non-entitled zone that omits certain fields from API responses
	testCases := []struct {
		name   string
		zoneID string
	}{
		{"StandardZone", os.Getenv("CLOUDFLARE_ZONE_ID")},
		{"AltZone", os.Getenv("CLOUDFLARE_ALT_ZONE_ID")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_bot_management." + rnd

			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { acctest.TestAccPreCheck(t) },
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testCloudflareBotManagementAutoUpdateModelStateConsistencyUserZone(rnd, tc.zoneID),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(tc.zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fight_mode"), knownvalue.Bool(false)),
						},
					},
					{
						Config: testCloudflareBotManagementAutoUpdateModelStateConsistencyUserZone(rnd, tc.zoneID),
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								plancheck.ExpectEmptyPlan(),
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fight_mode"), knownvalue.Bool(false)),
						},
					},
				},
			})
		})
	}
}

func testCloudflareBotManagementAutoUpdateModelStateConsistencyUserZone(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementautoupdatemodelstateconsistencyuserzone.tf", resourceName, zoneID)
}

func TestAccCloudflareBotManagement_FightModeStateConsistency(t *testing.T) {
	// Test against zones with different entitlement levels:
	// - CLOUDFLARE_ZONE_ID: Entitled zone that returns all fields in API responses
	// - CLOUDFLARE_ALT_ZONE_ID: Non-entitled zone that omits certain fields from API responses
	testCases := []struct {
		name   string
		zoneID string
	}{
		{"StandardZone", os.Getenv("CLOUDFLARE_ZONE_ID")},
		{"AltZone", os.Getenv("CLOUDFLARE_ALT_ZONE_ID")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_bot_management." + rnd

			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { acctest.TestAccPreCheck(t) },
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testCloudflareBotManagementFightModeStateConsistency(rnd, tc.zoneID),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(tc.zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fight_mode"), knownvalue.Bool(false)),
						},
					},
					{
						Config: testCloudflareBotManagementFightModeStateConsistency(rnd, tc.zoneID),
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								plancheck.ExpectEmptyPlan(),
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fight_mode"), knownvalue.Bool(false)),
						},
					},
				},
			})
		})
	}
}

func testCloudflareBotManagementFightModeStateConsistency(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementfightmodeconsistency.tf", resourceName, zoneID)
}

func testCloudflareBotManagementBasicPermutation(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementbasicpermutation.tf", resourceName, zoneID)
}

func testCloudflareBotManagementUpdatedPermutation(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementupdatedpermutation.tf", resourceName, zoneID)
}

func testCloudflareBotManagementSBFMPermutation1(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementsbfmpermutation1.tf", resourceName, zoneID)
}

func testCloudflareBotManagementSBFMPermutation2(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementsbfmpermutation2.tf", resourceName, zoneID)
}

func testCloudflareBotManagementComputedFields(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementcomputedfields.tf", resourceName, zoneID)
}

func testCloudflareBotManagementIssue5519(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementissue5519.tf", resourceName, zoneID)
}

func testCloudflareBotManagementIssue5519Lifecycle(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementissue5519lifecycle.tf", resourceName, zoneID)
}

func testCloudflareBotManagementMinimalConfig(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementminimal.tf", resourceName, zoneID)
}

func testCloudflareBotManagementIssue5519ExistingResourceConfig(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementissue5519existing.tf", resourceName, zoneID)
}

func testCloudflareBotManagementIssue5519Exact(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementissue5519exact.tf", resourceName, zoneID)
}

func testCloudflareBotManagementIssue5519PlanMismatch(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementissue5519planmismatch.tf", resourceName, zoneID)
}

func testCloudflareBotManagementIssue5519AutoUpdate(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementissue5519autoupdate.tf", resourceName, zoneID)
}

func testCloudflareBotManagementIssue5519NullFields(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementissue5519nullfields.tf", resourceName, zoneID)
}

func testCloudflareBotManagementIssue5519Reproduce(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementissue5519reproduce.tf", resourceName, zoneID)
}

func testCloudflareBotManagementSBFM(resourceName, rnd string, bm cloudflare.BotManagement) string {
	return acctest.LoadTestCase("cloudflarebotmanagementsbfm.tf", resourceName, rnd,
		*bm.EnableJS, *bm.SBFMDefinitelyAutomated,
		*bm.SBFMLikelyAutomated, *bm.SBFMVerifiedBots,
		*bm.SBFMStaticResourceProtection, *bm.OptimizeWordpress)
}

func testCloudflareBotManagementEntSubscription(resourceName, zoneID string, bm cloudflare.BotManagement) string {
	return acctest.LoadTestCase("cloudflarebotmanagemententsubscription.tf", resourceName, zoneID, *bm.EnableJS, *bm.SuppressSessionScore, false)
}

func testCloudflareBotManagementEnableJS(resourceName, zoneID string, enableJS bool) string {
	return acctest.LoadTestCase("cloudflarebotmanagementenablejs.tf", resourceName, zoneID, enableJS)
}

func testCloudflareBotManagementSuppressSessionScore(resourceName, zoneID string, suppressSessionScore bool) string {
	return acctest.LoadTestCase("cloudflarebotmanagementsuppresssessionscore.tf", resourceName, zoneID, suppressSessionScore)
}

func testCloudflareBotManagementAutoUpdateModel(resourceName, zoneID string, autoUpdateModel bool) string {
	return acctest.LoadTestCase("cloudflarebotmanagementautoupdatemodel.tf", resourceName, zoneID, autoUpdateModel)
}

func testCloudflareBotManagementEnableJSAutoUpdateSupression(resourceName, zoneID string) string {
	return acctest.LoadTestCase("cloudflarebotmanagementenablejsautoupdatesupression.tf", resourceName, zoneID)
}
