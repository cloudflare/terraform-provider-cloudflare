package leaked_credential_check_rule_test

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

func TestAccCloudflareLeakedCredentialsCheckRule_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_leaked_credential_check_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create + Read
			{
				Config: testAccCloudflareLeakedCredentialsCheckEnabled(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"username\")")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("password"), knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"pass\")")),
				},
			},
			// Step 2: Update + Read
			{
				Config: testAccCloudflareLeakedCredentialsCheckModified(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"username_modified\")")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("password"), knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"pass_modified\")")),
				},
			},
		},
	})
}

func TestAccCloudflareLeakedCredentialsCheckRule_StateConsistency(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_leaked_credential_check_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareLeakedCredentialsCheckEnabled(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"username\")")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("password"), knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"pass\")")),
				},
			},
			{
				Config: testAccCloudflareLeakedCredentialsCheckEnabled(zoneID, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"username\")")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("password"), knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"pass\")")),
				},
			},
		},
	})
}

// Helper functions to load test case configurations
func testAccCloudflareLeakedCredentialsCheckEnabled(zoneID, name string) string {
	return acctest.LoadTestCase("enabled.tf", zoneID, name)
}

func testAccCloudflareLeakedCredentialsCheckNotEnabled(zoneID, name string) string {
	return acctest.LoadTestCase("not_enabled.tf", zoneID, name)
}

func testAccCloudflareLeakedCredentialsCheckModified(zoneID, name string) string {
	return acctest.LoadTestCase("modified.tf", zoneID, name)
}
