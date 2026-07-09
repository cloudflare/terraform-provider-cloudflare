package origin_tls_compliance_modes_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_origin_tls_compliance_modes", &resource.Sweeper{
		Name: "cloudflare_origin_tls_compliance_modes",
		F:    testSweepCloudflareOriginTLSComplianceModes,
	})
}

func testSweepCloudflareOriginTLSComplianceModes(r string) error {
	ctx := context.Background()
	// Origin TLS Compliance Modes is a zone-level setting that constrains
	// the key-exchange algorithms Cloudflare may use when establishing TLS
	// connections to the zone's origin. It's a singleton setting per zone,
	// not something that accumulates. No sweeping required.
	tflog.Info(ctx, "Origin TLS Compliance Modes doesn't require sweeping (zone setting)")
	return nil
}

func TestAccCloudflareOriginTLSComplianceModes_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_origin_tls_compliance_modes.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with single mode
			{
				Config: testAccCheckCloudflareOriginTLSComplianceModesFipsOnly(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", zoneID),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "value.#", "1"),
					resource.TestCheckResourceAttr(name, "value.0", "fips"),
					resource.TestCheckResourceAttrSet(name, "modified_on"),
				),
			},
			// Step 2: Idempotency
			{
				Config:             testAccCheckCloudflareOriginTLSComplianceModesFipsOnly(zoneID, rnd),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Step 3: Update to multiple modes
			{
				Config: testAccCheckCloudflareOriginTLSComplianceModesFipsAndPqh(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", zoneID),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "value.#", "2"),
					resource.TestCheckResourceAttr(name, "value.0", "fips"),
					resource.TestCheckResourceAttr(name, "value.1", "pqh"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("value"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.StringExact("fips"),
								knownvalue.StringExact("pqh"),
							}),
						),
					},
				},
			},
			// Step 4: Import
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCloudflareOriginTLSComplianceModesFipsOnly(zoneID, name string) string {
	return acctest.LoadTestCase("fips_only.tf", zoneID, name)
}

func testAccCheckCloudflareOriginTLSComplianceModesFipsAndPqh(zoneID, name string) string {
	return acctest.LoadTestCase("fips_and_pqh.tf", zoneID, name)
}
