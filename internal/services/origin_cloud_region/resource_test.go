package origin_cloud_region_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// Documentation IPs (RFC 5737 / RFC 3849) — safe to use in tests
const (
	testIPv4         = "192.0.2.1"
	testIPv4Other    = "192.0.2.2"
	testIPv6         = "2001:db8::1"
	testVendor       = "aws"
	testRegion       = "us-east-1"
	testRegionUpdate = "us-west-2"
)

// Basic CRUD: create → read → update region → delete
func TestAccCloudflareOriginCloudRegion_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_origin_cloud_region.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccCloudflareOriginCloudRegionBasic(rnd, zoneID, testIPv4, testVendor, testRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "origin_ip", testIPv4),
					resource.TestCheckResourceAttr(name, "vendor", testVendor),
					resource.TestCheckResourceAttr(name, "region", testRegion),
					resource.TestCheckResourceAttrSet(name, "modified_on"),
				),
			},
			// Plan should be empty after apply (round-trip consistency)
			{
				Config:             testAccCloudflareOriginCloudRegionBasic(rnd, zoneID, testIPv4, testVendor, testRegion),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Update region (in-place; vendor stays)
			{
				Config: testAccCloudflareOriginCloudRegionBasic(rnd, zoneID, testIPv4, testVendor, testRegionUpdate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "origin_ip", testIPv4),
					resource.TestCheckResourceAttr(name, "vendor", testVendor),
					resource.TestCheckResourceAttr(name, "region", testRegionUpdate),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(name, tfjsonpath.New("region"), knownvalue.StringExact(testRegionUpdate)),
						plancheck.ExpectKnownValue(name, tfjsonpath.New("vendor"), knownvalue.StringExact(testVendor)),
					},
				},
			},
			// Import test — expects "<zone_id>/<origin_ip>" format
			{
				ResourceName:                         name,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "origin_ip",
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", zoneID, testIPv4), nil
				},
			},
		},
	})
}

// Changing origin_ip should force replace (destroy + recreate)
func TestAccCloudflareOriginCloudRegion_ChangeOriginIPForcesReplace(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_origin_cloud_region.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareOriginCloudRegionBasic(rnd, zoneID, testIPv4, testVendor, testRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "origin_ip", testIPv4),
				),
			},
			{
				Config: testAccCloudflareOriginCloudRegionBasic(rnd, zoneID, testIPv4Other, testVendor, testRegion),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "origin_ip", testIPv4Other),
				),
			},
		},
	})
}

// IPv6 origin
func TestAccCloudflareOriginCloudRegion_IPv6(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_origin_cloud_region.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareOriginCloudRegionBasic(rnd, zoneID, testIPv6, "gcp", "us-central1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "origin_ip", testIPv6),
					resource.TestCheckResourceAttr(name, "vendor", "gcp"),
				),
			},
			{
				Config:             testAccCloudflareOriginCloudRegionBasic(rnd, zoneID, testIPv6, "gcp", "us-central1"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// Invalid vendor should fail validation
func TestAccCloudflareOriginCloudRegion_InvalidVendor(t *testing.T) {
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
				Config:      testAccCloudflareOriginCloudRegionBasic(rnd, zoneID, testIPv4, "badvendor", testRegion),
				ExpectError: regexp.MustCompile("Invalid Attribute Value Match|Attribute vendor"),
			},
		},
	})
}

func testAccCloudflareOriginCloudRegionBasic(rnd, zoneID, originIP, vendor, region string) string {
	return acctest.LoadTestCase("basic.tf", rnd, zoneID, originIP, vendor, region)
}
