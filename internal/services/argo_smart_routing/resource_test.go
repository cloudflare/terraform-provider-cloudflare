package argo_smart_routing_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestAccCloudflareArgoSmartRouting_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_argo_smart_routing.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareArgoSmartRoutingEnable(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", zoneID),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "value", "on"),
				),
			},
			{
				Config: testAccCheckCloudflareArgoSmartRoutingEnable(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", zoneID),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "value", "on"),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				Config: testAccCheckCloudflareArgoSmartRoutingDisable(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", zoneID),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "value", "off"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("value"),
							knownvalue.StringExact("off"),
						),
					},
				},
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareArgoSmartRouting_InvalidValue(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareArgoSmartRoutingInvalidValue(zoneID, rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("Invalid Attribute Value Match")),
			},
		},
	})
}

func testAccCheckCloudflareArgoSmartRoutingEnable(zoneID, name string) string {
	return acctest.LoadTestCase("enable.tf", zoneID, name)
}

func testAccCheckCloudflareArgoSmartRoutingDisable(zoneID, name string) string {
	return acctest.LoadTestCase("disable.tf", zoneID, name)
}

func testAccCheckCloudflareArgoSmartRoutingInvalidValue(zoneID, name string) string {
	return acctest.LoadTestCase("invalid_value.tf", zoneID, name)
}
