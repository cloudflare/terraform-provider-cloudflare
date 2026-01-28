package leaked_credential_check_rule_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/leaked_credential_checks"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func testAccCheckCloudflareLeakedCredentialCheckRuleDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_leaked_credential_check_rule" {
			continue
		}

		res := new(http.Response)
		_, err := client.LeakedCredentialChecks.Detections.Get(
			context.Background(),
			rs.Primary.ID,
			leaked_credential_checks.DetectionGetParams{
				ZoneID: cloudflare.F(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
			},
			option.WithResponseBodyInto(&res),
		)

		// Check HTTP status code for 404
		if res != nil && res.StatusCode == 404 {
			continue // Resource is deleted, as expected
		}

		if err == nil {
			return fmt.Errorf("leaked credential check rule %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

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
		CheckDestroy:             testAccCheckCloudflareLeakedCredentialCheckRuleDestroy,
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
		CheckDestroy:             testAccCheckCloudflareLeakedCredentialCheckRuleDestroy,
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

func TestAccCloudflareLeakedCredentialsCheckRule_Import(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_leaked_credential_check_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLeakedCredentialCheckRuleDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create the resource
			{
				Config: testAccCloudflareLeakedCredentialsCheckEnabled(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"username\")")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("password"), knownvalue.StringExact("lookup_json_string(http.request.body.raw, \"pass\")")),
				},
			},
			// Step 2: Import the resource using zone_id/detection_id format
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: generateLeakedCredentialCheckRuleImportStateId(resourceName),
			},
		},
	})
}

// generateLeakedCredentialCheckRuleImportStateId generates the import ID in the format zone_id/detection_id
func generateLeakedCredentialCheckRuleImportStateId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		detectionID := rs.Primary.ID

		return fmt.Sprintf("%s/%s", zoneID, detectionID), nil
	}
}

// Helper functions to load test case configurations
func testAccCloudflareLeakedCredentialsCheckEnabled(zoneID, name string) string {
	return acctest.LoadTestCase("enabled.tf", zoneID, name)
}

func testAccCloudflareLeakedCredentialsCheckModified(zoneID, name string) string {
	return acctest.LoadTestCase("modified.tf", zoneID, name)
}
