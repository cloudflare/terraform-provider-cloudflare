package zero_trust_dlp_sensitivity_level_order_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestAccCloudflareZeroTrustDLPSensitivityLevelOrder_Basic verifies the
// fundamental singleton lifecycle:
//   - Create: the order resource has no POST endpoint; on first apply Terraform
//     calls PUT /level_order with the desired list. State reflects that list.
//   - Read: GET /level_order returns the same list; state stays stable.
//   - Destroy: removing the resource just drops state (no API call).
//
// The test also covers Import, which is the only way to bring an existing order
// into Terraform state since the singleton has no "create" step.
func TestAccCloudflareZeroTrustDLPSensitivityLevelOrder_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	orderResourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_level_order.%s", rnd)
	publicLevelResource := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_level.%s_public", rnd)
	internalLevelResource := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_level.%s_internal", rnd)
	confidentialLevelResource := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_level.%s_confidential", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("basic.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(orderResourceName,
						tfjsonpath.New("account_id"),
						knownvalue.StringExact(accountID)),
					// The order resource's id is its sensitivity_group_id (per
					// Stainless config: id_property: sensitivity_group_id).
					statecheck.CompareValuePairs(
						orderResourceName, tfjsonpath.New("id"),
						orderResourceName, tfjsonpath.New("sensitivity_group_id"),
						compare.ValuesSame()),
					// level_ids must reflect the desired order with the
					// resolved level UUIDs.
					statecheck.CompareValuePairs(
						orderResourceName, tfjsonpath.New("level_ids").AtSliceIndex(0),
						publicLevelResource, tfjsonpath.New("id"),
						compare.ValuesSame()),
					statecheck.CompareValuePairs(
						orderResourceName, tfjsonpath.New("level_ids").AtSliceIndex(1),
						internalLevelResource, tfjsonpath.New("id"),
						compare.ValuesSame()),
					statecheck.CompareValuePairs(
						orderResourceName, tfjsonpath.New("level_ids").AtSliceIndex(2),
						confidentialLevelResource, tfjsonpath.New("id"),
						compare.ValuesSame()),
				},
			},
			// Re-apply: no changes expected.
			{
				Config: acctest.LoadTestCase("basic.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Import: the import ID is "<account_id>/<sensitivity_group_id>".
			{
				ResourceName:                         orderResourceName,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateIdPrefix:                  fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIdentifierAttribute: "sensitivity_group_id",
			},
		},
	})
}

// TestAccCloudflareZeroTrustDLPSensitivityLevelOrder_Reorder verifies that
// updating `level_ids` to a new permutation produces a single PUT (no
// recreation, no other field changes).
func TestAccCloudflareZeroTrustDLPSensitivityLevelOrder_Reorder(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	orderResourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_level_order.%s", rnd)
	publicLevelResource := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_level.%s_public", rnd)
	internalLevelResource := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_level.%s_internal", rnd)
	confidentialLevelResource := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_level.%s_confidential", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("basic.tf", rnd, accountID),
			},
			// Swap Internal and Confidential. Only level_ids should change in
			// the plan; the resource should not be recreated.
			{
				Config: acctest.LoadTestCase("reordered.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(orderResourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.CompareValuePairs(
						orderResourceName, tfjsonpath.New("level_ids").AtSliceIndex(0),
						publicLevelResource, tfjsonpath.New("id"),
						compare.ValuesSame()),
					statecheck.CompareValuePairs(
						orderResourceName, tfjsonpath.New("level_ids").AtSliceIndex(1),
						confidentialLevelResource, tfjsonpath.New("id"),
						compare.ValuesSame()),
					statecheck.CompareValuePairs(
						orderResourceName, tfjsonpath.New("level_ids").AtSliceIndex(2),
						internalLevelResource, tfjsonpath.New("id"),
						compare.ValuesSame()),
				},
			},
		},
	})
}

// TestAccCloudflareZeroTrustDLPSensitivityLevelOrder_RemoveFromConfig verifies
// that removing the order resource from HCL does NOT call any API (the singleton
// has no DELETE endpoint; its lifecycle is owned by the parent group). Terraform
// just drops the resource from state.
func TestAccCloudflareZeroTrustDLPSensitivityLevelOrder_RemoveFromConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: apply with the order resource present.
			{
				Config: acctest.LoadTestCase("basic.tf", rnd, accountID),
			},
			// Step 2: same group + levels but no order resource. The Destroy
			// implementation is a no-op (no API call). Terraform should drop the
			// order resource from state without error and without affecting the
			// remaining resources.
			{
				Config: acctest.LoadTestCase("without_order.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_level_order.%s", rnd),
							plancheck.ResourceActionDestroy,
						),
					},
				},
			},
		},
	})
}
