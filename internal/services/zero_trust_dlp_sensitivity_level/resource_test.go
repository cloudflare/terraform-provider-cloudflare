package zero_trust_dlp_sensitivity_level_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestAccCloudflareZeroTrustDLPSensitivityLevel_Basic verifies basic CRUD for
// a sensitivity level nested under a sensitivity group.
func TestAccCloudflareZeroTrustDLPSensitivityLevel_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_level.%s", rnd)
	groupName := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("basic.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name,
						tfjsonpath.New("account_id"),
						knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name,
						tfjsonpath.New("name"),
						knownvalue.StringExact(fmt.Sprintf("tf-acc-%s", rnd))),
					statecheck.ExpectKnownValue(name,
						tfjsonpath.New("description"),
						knownvalue.StringExact("Acceptance test sensitivity level")),
					statecheck.ExpectKnownValue(name,
						tfjsonpath.New("sensitivity_group_id"),
						knownvalue.NotNull()),
				},
			},
			// Re-apply: no drift.
			{
				Config: acctest.LoadTestCase("basic.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Import: `<account_id>/<sensitivity_group_id>/<sensitivity_level_id>`.
			// The import requires both path segments since the level is nested.
			{
				ResourceName: name,
				ImportState:  true,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					group, ok := state.RootModule().Resources[groupName]
					if !ok {
						return "", fmt.Errorf("group resource %s not found in state", groupName)
					}
					level, ok := state.RootModule().Resources[name]
					if !ok {
						return "", fmt.Errorf("level resource %s not found in state", name)
					}
					return fmt.Sprintf("%s/%s/%s",
						accountID,
						group.Primary.ID,
						level.Primary.ID), nil
				},
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccCloudflareZeroTrustDLPSensitivityLevel_Update verifies in-place update
// of name and description on a sensitivity level.
func TestAccCloudflareZeroTrustDLPSensitivityLevel_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_level.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("basic.tf", rnd, accountID),
			},
			{
				Config: acctest.LoadTestCase("updated.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name,
						tfjsonpath.New("name"),
						knownvalue.StringExact(fmt.Sprintf("tf-acc-%s-renamed", rnd))),
					statecheck.ExpectKnownValue(name,
						tfjsonpath.New("description"),
						knownvalue.StringExact("Updated description")),
				},
			},
		},
	})
}
