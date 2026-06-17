package zone_auto_origin_tls_kex_test

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
	resource.AddTestSweepers("cloudflare_zone_auto_origin_tls_kex", &resource.Sweeper{
		Name: "cloudflare_zone_auto_origin_tls_kex",
		F:    testSweepCloudflareZoneAutoOriginTLSKex,
	})
}

func testSweepCloudflareZoneAutoOriginTLSKex(r string) error {
	ctx := context.Background()
	// Auto-Origin TLS KEX is a zone-level feature toggle (on/off).
	// It's a singleton setting per zone, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Zone Auto-Origin TLS KEX doesn't require sweeping (zone setting)")
	return nil
}

func TestAccCloudflareZoneAutoOriginTLSKex_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_auto_origin_tls_kex.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Enable
			{
				Config: testAccCheckCloudflareZoneAutoOriginTLSKexEnable(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", zoneID),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttrSet(name, "modified_on"),
				),
			},
			// Step 2: Idempotency
			{
				Config:             testAccCheckCloudflareZoneAutoOriginTLSKexEnable(zoneID, rnd),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Step 3: Disable
			{
				Config: testAccCheckCloudflareZoneAutoOriginTLSKexDisable(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", zoneID),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("enabled"),
							knownvalue.Bool(false),
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

func testAccCheckCloudflareZoneAutoOriginTLSKexEnable(zoneID, name string) string {
	return acctest.LoadTestCase("enable.tf", zoneID, name)
}

func testAccCheckCloudflareZoneAutoOriginTLSKexDisable(zoneID, name string) string {
	return acctest.LoadTestCase("disable.tf", zoneID, name)
}
