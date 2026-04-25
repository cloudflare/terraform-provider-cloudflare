package zero_trust_dlp_settings_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareZeroTrustDLPSettings_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_settings.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		// No CheckDestroy: this is a singleton settings resource.
		// DELETE resets to initial values; GET always returns a response.
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("dlp_settings_init.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_context_analysis"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ocr"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("payload_logging").AtMapKey("public_key"), knownvalue.StringExact("EmpOvSXw8BfbrGCi0fhGiD/3yXk2SiV1Nzg2lru3oj0=")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("payload_logging").AtMapKey("masking_level"), knownvalue.StringExact("partial")),
				},
			},
			{
				Config: acctest.LoadTestCase("dlp_settings_update.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_context_analysis"), knownvalue.Bool(false)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("ocr"), knownvalue.Bool(false)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("payload_logging").AtMapKey("masking_level"), knownvalue.StringExact("full")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_context_analysis"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ocr"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("payload_logging").AtMapKey("public_key"), knownvalue.StringExact("EmpOvSXw8BfbrGCi0fhGiD/3yXk2SiV1Nzg2lru3oj0=")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("payload_logging").AtMapKey("masking_level"), knownvalue.StringExact("full")),
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

// TestAccCloudflareZeroTrustDLPSettings_PartialConfig verifies that a user
// can create DLP settings with only a subset of computed_optional fields
// (e.g. only ocr=true). The apply succeeds and state reflects the correct
// values from the API.
func TestAccCloudflareZeroTrustDLPSettings_PartialConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_settings.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("dlp_settings_partial_ocr.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ocr"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_context_analysis"), knownvalue.Bool(false)),
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
