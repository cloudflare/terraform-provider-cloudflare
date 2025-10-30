package zero_trust_dlp_custom_profile_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

const (
	EnvTfAcc = "TF_ACC"
)

// setupDLPCustomProfileTest handles common test setup and TF_ACC environment check
func setupDLPCustomProfileTest(t *testing.T) (string, string) {
	if os.Getenv(EnvTfAcc) == "" {
		t.Skip(fmt.Sprintf(
			"Acceptance tests skipped unless env '%s' set",
			EnvTfAcc))
		return "", ""
	}

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	return rnd, accountID
}

func TestAccCloudflareZeroTrustDlpCustomProfile_Basic(t *testing.T) {
	rnd, accountID := setupDLPCustomProfileTest(t)
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_custom_profile.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDlpCustomProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("basic.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Test custom DLP profile")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_match_count"), knownvalue.Int64Exact(5)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ocr_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_context_enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("custom")),
				},
			},
			{
				Config: acctest.LoadTestCase("update.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Updated test custom DLP profile")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_match_count"), knownvalue.Int64Exact(10)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("ocr_enabled"), knownvalue.Bool(false)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_context_enabled"), knownvalue.Bool(true)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Updated test custom DLP profile")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_match_count"), knownvalue.Int64Exact(10)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ocr_enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_context_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("custom")),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"open_access", "created_at", "updated_at", "type", "context_awareness", "entries", "ai_context_enabled"},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDlpCustomProfile_MinimalRequired(t *testing.T) {
	rnd, accountID := setupDLPCustomProfileTest(t)
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_custom_profile.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDlpCustomProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("minimal.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.Null()),
					// Test default values
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_context_enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_match_count"), knownvalue.Int64Exact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("confidence_threshold"), knownvalue.StringExact("low")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ocr_enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("custom")),
					// Test computed attributes are not null
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: acctest.LoadTestCase("max_attributes.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-max")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Profile with all optional attributes set")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_match_count"), knownvalue.Int64Exact(1000)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("confidence_threshold"), knownvalue.StringExact("high")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("ocr_enabled"), knownvalue.Bool(true)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_context_enabled"), knownvalue.Bool(true)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-max")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Profile with all optional attributes set")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_match_count"), knownvalue.Int64Exact(1000)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("confidence_threshold"), knownvalue.StringExact("high")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ocr_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_context_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("custom")),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"open_access", "created_at", "updated_at", "type", "context_awareness", "entries", "ai_context_enabled"},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDlpCustomProfile_AllSharedEntryTypes(t *testing.T) {
	rnd, accountID := setupDLPCustomProfileTest(t)
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_custom_profile.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDlpCustomProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("all_entry_types.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-all-types")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Profile without shared entries for basic testing")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_match_count"), knownvalue.Int64Exact(10)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ocr_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_context_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("custom")),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"open_access", "created_at", "updated_at", "type", "context_awareness", "entries", "ai_context_enabled"},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDlpCustomProfile_DeprecatedAttributes(t *testing.T) {
	rnd, accountID := setupDLPCustomProfileTest(t)
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_custom_profile.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDlpCustomProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("custom_entries.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Test with custom entries")),
					// Test deprecated custom entries
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("entries"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"enabled":  knownvalue.Bool(true),
							"entry_id": knownvalue.Null(),
							"name":     knownvalue.StringExact("Credit Card Pattern"),
							"pattern": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"regex":      knownvalue.StringExact("\\d{4}-\\d{4}-\\d{4}-\\d{4}"),
								"validation": knownvalue.StringExact("luhn"),
							}),
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"enabled":  knownvalue.Bool(false),
							"entry_id": knownvalue.Null(),
							"name":     knownvalue.StringExact("SSN Pattern"),
							"pattern": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"regex":      knownvalue.StringExact("\\d{3}-\\d{2}-\\d{4}"),
								"validation": knownvalue.Null(),
							}),
						}),
					})),
					// Test deprecated context awareness
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("context_awareness"), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"enabled": knownvalue.Bool(true),
						"skip": knownvalue.ObjectExact(map[string]knownvalue.Check{
							"files": knownvalue.Bool(false),
						}),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"open_access", "created_at", "updated_at", "type", "context_awareness", "entries", "ai_context_enabled"},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDlpCustomProfile_BoundaryValues(t *testing.T) {
	rnd, accountID := setupDLPCustomProfileTest(t)
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_custom_profile.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDlpCustomProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("boundary_values.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-boundary")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Testing boundary values")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_match_count"), knownvalue.Int64Exact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("confidence_threshold"), knownvalue.StringExact("low")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ocr_enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_context_enabled"), knownvalue.Bool(false)),
				},
			},
			{
				Config: acctest.LoadTestCase("confidence_medium.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-medium")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("confidence_threshold"), knownvalue.StringExact("medium")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-medium")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("confidence_threshold"), knownvalue.StringExact("medium")),
					// Description should be null when not set
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.Null()),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"open_access", "created_at", "updated_at", "type", "context_awareness", "entries", "ai_context_enabled"},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDlpCustomProfile_SharedEntries(t *testing.T) {
	rnd, accountID := setupDLPCustomProfileTest(t)
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_custom_profile.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDlpCustomProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("shared_entries.tf", rnd, accountID, "true"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Test with shared entries")),
					// Test deprecated custom entries
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("shared_entries"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"enabled":  knownvalue.Bool(true),
							"entry_id": knownvalue.StringExact("56a8c060-01bb-4f89-ba1e-3ad42770a342"),
							"entry_type": knownvalue.StringExact("predefined"),
						}),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"open_access", "created_at", "updated_at", "type", "context_awareness", "entries", "shared_entries"},
			},
		},
	})
}

func testAccCheckCloudflareZeroTrustDlpCustomProfileDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_dlp_custom_profile" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.ZeroTrust.DLP.Profiles.Custom.Get(
			context.Background(),
			rs.Primary.ID,
			zero_trust.DLPProfileCustomGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("zero trust DLP custom profile still exists")
		}
	}

	return nil
}
