package zero_trust_dlp_sensitivity_group_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestAccCloudflareZeroTrustDLPSensitivityGroup_Basic verifies basic CRUD
// for a sensitivity group: apply, read, import, destroy.
func TestAccCloudflareZeroTrustDLPSensitivityGroup_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_group.%s", rnd)

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
						knownvalue.StringExact("Acceptance test sensitivity group")),
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
			// Import: `<account_id>/<sensitivity_group_id>`.
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// TestAccCloudflareZeroTrustDLPSensitivityGroup_Update verifies that updating
// name and description produces an in-place update (no recreation).
func TestAccCloudflareZeroTrustDLPSensitivityGroup_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_group.%s", rnd)

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
						knownvalue.StringExact("Updated description for acceptance test")),
				},
			},
		},
	})
}
