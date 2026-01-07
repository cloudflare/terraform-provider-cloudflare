package content_scanning_test

import (
	"fmt"
	"os"
	"regexp"
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

func TestAccCloudflareContentScanning_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_content_scanning.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create + Read
			{
				Config: testAccCloudflareContentScanningEnabled(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("enabled")),
					statecheck.ExpectKnownOutputValue("modified", knownvalue.NotNull()),
				},
			},
			// Step 2: Update + Read
			{
				Config: testAccCloudflareContentScanningDisabled(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("disabled")),
				},
			},
		},
	})
}

func TestAccCloudflareContentScanning_InvalidValue(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareContentScanningInvalidValue(zoneID, rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("Invalid Attribute Value Match")),
			},
		},
	})
}

func TestAccCloudflareContentScanning_StateConsistency(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_content_scanning.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareContentScanningEnabled(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("enabled")),
				},
			},
			{
				Config: testAccCloudflareContentScanningEnabled(zoneID, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("enabled")),
				},
			},
		},
	})
}

// Helper functions to load test case configurations
func testAccCloudflareContentScanningEnabled(zoneID, name string) string {
	return acctest.LoadTestCase("enabled.tf", zoneID, name)
}

func testAccCloudflareContentScanningDisabled(zoneID, name string) string {
	return acctest.LoadTestCase("disabled.tf", zoneID, name)
}

func testAccCloudflareContentScanningInvalidValue(zoneID, name string) string {
	return acctest.LoadTestCase("invalid_value.tf", zoneID, name)
}
