package zero_trust_dlp_data_tag_test

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

// TestAccCloudflareZeroTrustDLPDataTag_Basic verifies basic CRUD for
// a data tag nested under a data tag category.
func TestAccCloudflareZeroTrustDLPDataTag_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_zero_trust_dlp_data_tag.%s", rnd)
	categoryName := fmt.Sprintf("cloudflare_zero_trust_dlp_data_tag_category.%s", rnd)

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
						knownvalue.StringExact("Acceptance test data tag")),
					statecheck.ExpectKnownValue(name,
						tfjsonpath.New("category_id"),
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
			// Import: `<account_id>/<category_id>/<tag_id>`.
			// The import requires both path segments since the tag is nested.
			{
				ResourceName: name,
				ImportState:  true,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					category, ok := state.RootModule().Resources[categoryName]
					if !ok {
						return "", fmt.Errorf("category resource %s not found in state", categoryName)
					}
					tag, ok := state.RootModule().Resources[name]
					if !ok {
						return "", fmt.Errorf("tag resource %s not found in state", name)
					}
					return fmt.Sprintf("%s/%s/%s",
						accountID,
						category.Primary.ID,
						tag.Primary.ID), nil
				},
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccCloudflareZeroTrustDLPDataTag_Update verifies in-place update
// of name and description on a data tag.
func TestAccCloudflareZeroTrustDLPDataTag_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_zero_trust_dlp_data_tag.%s", rnd)

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
