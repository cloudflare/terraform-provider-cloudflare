package zero_trust_dlp_data_tag_category_test

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

// TestAccCloudflareZeroTrustDLPDataTagCategory_Basic verifies basic CRUD for
// a data tag category: apply, read, import, destroy.
func TestAccCloudflareZeroTrustDLPDataTagCategory_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_zero_trust_dlp_data_tag_category.%s", rnd)

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
						knownvalue.StringExact("Acceptance test data tag category")),
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
			// Import: `<account_id>/<category_id>`.
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// TestAccCloudflareZeroTrustDLPDataTagCategory_Update verifies in-place update.
func TestAccCloudflareZeroTrustDLPDataTagCategory_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_zero_trust_dlp_data_tag_category.%s", rnd)

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
