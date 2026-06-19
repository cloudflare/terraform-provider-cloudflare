package zero_trust_dlp_data_class_test

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

// TestAccCloudflareZeroTrustDLPDataClass_Basic verifies basic CRUD for a data
// class, including its references to a data_tag and a sensitivity_level.
func TestAccCloudflareZeroTrustDLPDataClass_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_zero_trust_dlp_data_class.%s", rnd)

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
						knownvalue.StringExact("Acceptance test data class")),
					statecheck.ExpectKnownValue(name,
						tfjsonpath.New("expression"),
						knownvalue.StringExact("dlp_match(dlp.entries[\"906fcb91-2eb5-4534-8f86-f95214b651eb\"])")),
					statecheck.ExpectKnownValue(name,
						tfjsonpath.New("data_tags"),
						knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(name,
						tfjsonpath.New("sensitivity_levels"),
						knownvalue.ListSizeExact(1)),
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
			// Import: `<account_id>/<data_class_id>`.
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// TestAccCloudflareZeroTrustDLPDataClass_Update verifies in-place update of
// name, description, and expression on a data class.
func TestAccCloudflareZeroTrustDLPDataClass_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	name := fmt.Sprintf("cloudflare_zero_trust_dlp_data_class.%s", rnd)

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
					statecheck.ExpectKnownValue(name,
						tfjsonpath.New("expression"),
						knownvalue.StringExact("dlp_match(dlp.entries[\"5bce03be-0b03-434c-9f56-7b42512e0ff6\"])")),
				},
			},
		},
	})
}
