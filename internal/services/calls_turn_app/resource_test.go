package calls_turn_app_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testAccCallsTURNAppBasic(accID, name string) string {
	return acctest.LoadTestCase("calls_turn_app_basic.tf", accID, name)
}

func testAccCallsTURNAppUpdateName(accID, resourceName, name string) string {
	return acctest.LoadTestCase("calls_turn_app_update_name.tf", accID, resourceName, name)
}

func callsTURNImportIDFunc(resourceName string) resource.ImportStateIdFunc {
	return func(state *terraform.State) (string, error) {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes["account_id"], rs.Primary.Attributes["key_id"]), nil
	}
}

// TestAccCallsTURNApp_Lifecycle exercises the full CRUD + Import lifecycle.
//
// The critical regression being guarded here is ESCALATION-10887 / GH-7079:
// after a successful create, the API returns the resource identifier as `uid`
// in the JSON body but the provider uses `key_id` as the URL path parameter
// for Get/Update/Delete. Without the fix, `key_id` is never stored in state,
// so every subsequent plan fails with "missing required key_id parameter"
// before any HTTP request is attempted.
//
// Steps:
//  1. Create — key_id must be populated in state from the API's uid field
//  2. PlanOnly — must be a no-op (regression: used to fail with missing key_id)
//  3. Update name — in-place rename must succeed
//  4. Import — all attributes must survive a round-trip through the API
func TestAccCallsTURNApp_Lifecycle(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_calls_turn_app.%s", rnd)
	updatedName := rnd + "_updated"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Create — key_id must be populated from uid.
				Config: testAccCallsTURNAppBasic(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "uid"),
					// key_id must equal uid so subsequent Read/Update/Delete work.
					resource.TestCheckResourceAttrPair(resourceName, "key_id", resourceName, "uid"),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttrSet(resourceName, "created"),
				),
			},
			{
				// Step 2: PlanOnly — must produce an empty plan.
				// Before the fix this step failed:
				//   Error: failed to make http request
				//   missing required key_id parameter
				Config:   testAccCallsTURNAppBasic(accID, rnd),
				PlanOnly: true,
			},
			{
				// Step 3: Update name in-place.
				Config: testAccCallsTURNAppUpdateName(accID, rnd, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				),
			},
			{
				// Step 4: Import — key_id must be restored so the next plan is clean.
				// This resource has no top-level id attribute; key_id is used as the
				// import identifier for verification. key is excluded because it is
				// write-once (returned only on Create, marked no_refresh) and cannot
				// be round-tripped through import.
				ResourceName:                        resourceName,
				ImportStateIdFunc:                   callsTURNImportIDFunc(resourceName),
				ImportState:                         true,
				ImportStateVerify:                   true,
				ImportStateVerifyIdentifierAttribute: "key_id",
				ImportStateVerifyIgnore:             []string{"key"},
			},
		},
	})
}
