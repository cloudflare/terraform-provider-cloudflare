package google_tag_gateway_test

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

func TestAccCloudflareGoogleTagGateway_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_google_tag_gateway.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGoogleTagGatewayBasicConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "endpoint", "/gtm"),
					resource.TestCheckResourceAttr(name, "hide_original_ip", "true"),
					resource.TestCheckResourceAttr(name, "measurement_id", "GTM-XXXXXXX"),
					resource.TestCheckResourceAttr(name, "set_up_tag", "true"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("zone_id"),
						knownvalue.StringExact(zoneID),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("enabled"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("endpoint"),
						knownvalue.StringExact("/gtm"),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("hide_original_ip"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("measurement_id"),
						knownvalue.StringExact("GTM-XXXXXXX"),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("set_up_tag"),
						knownvalue.Bool(true),
					),
				},
			},
			// Refresh plan validation - ensure no changes on re-apply
			{
				Config: testAccGoogleTagGatewayBasicConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "endpoint", "/gtm"),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Update step - change multiple fields
			{
				Config: testAccGoogleTagGatewayUpdateConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "endpoint", "/metrics"),
					resource.TestCheckResourceAttr(name, "hide_original_ip", "false"),
					resource.TestCheckResourceAttr(name, "measurement_id", "G-XXXXXXXXXX"),
					resource.TestCheckResourceAttr(name, "set_up_tag", "false"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						// Verify only the expected fields are changing
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("enabled"),
							knownvalue.Bool(false),
						),
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("endpoint"),
							knownvalue.StringExact("/metrics"),
						),
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("hide_original_ip"),
							knownvalue.Bool(false),
						),
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("measurement_id"),
							knownvalue.StringExact("G-XXXXXXXXXX"),
						),
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("set_up_tag"),
							knownvalue.Bool(false),
						),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("enabled"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("endpoint"),
						knownvalue.StringExact("/metrics"),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("hide_original_ip"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("measurement_id"),
						knownvalue.StringExact("G-XXXXXXXXXX"),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("set_up_tag"),
						knownvalue.Bool(false),
					),
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

func TestAccCloudflareGoogleTagGateway_Minimal(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_google_tag_gateway.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGoogleTagGatewayMinimalConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "endpoint", "/analytics"),
					resource.TestCheckResourceAttr(name, "hide_original_ip", "false"),
					resource.TestCheckResourceAttr(name, "measurement_id", "GTM-ABCDEFG"),
					// set_up_tag is optional, verify it's not set
					resource.TestCheckNoResourceAttr(name, "set_up_tag"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareGoogleTagGateway_UpdateIndividualFields(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_google_tag_gateway.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with initial config
			{
				Config: testAccGoogleTagGatewayBasicConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "endpoint", "/gtm"),
					resource.TestCheckResourceAttr(name, "hide_original_ip", "true"),
				),
			},
			// Step 2: Update only the enabled field
			{
				Config: testAccGoogleTagGatewayUpdateEnabledOnlyConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "endpoint", "/gtm"),
					resource.TestCheckResourceAttr(name, "hide_original_ip", "true"),
					resource.TestCheckResourceAttr(name, "measurement_id", "GTM-XXXXXXX"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						// Verify only the enabled field changes
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("enabled"),
							knownvalue.Bool(false),
						),
					},
				},
			},
			// Step 3: Update only the endpoint field
			{
				Config: testAccGoogleTagGatewayUpdateEndpointOnlyConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "endpoint", "/newendpoint"),
					resource.TestCheckResourceAttr(name, "hide_original_ip", "true"),
					resource.TestCheckResourceAttr(name, "measurement_id", "GTM-XXXXXXX"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						// Verify only the endpoint field changes
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("endpoint"),
							knownvalue.StringExact("/newendpoint"),
						),
					},
				},
			},
			// Step 4: Update only the measurement_id field
			{
				Config: testAccGoogleTagGatewayUpdateMeasurementIDOnlyConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "endpoint", "/newendpoint"),
					resource.TestCheckResourceAttr(name, "hide_original_ip", "true"),
					resource.TestCheckResourceAttr(name, "measurement_id", "G-YYYYYYYYYY"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						// Verify only the measurement_id field changes
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("measurement_id"),
							knownvalue.StringExact("G-YYYYYYYYYY"),
						),
					},
				},
			},
		},
	})
}

func TestAccCloudflareGoogleTagGateway_RefreshPlanEmpty(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_google_tag_gateway.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGoogleTagGatewayBasicConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
			// Validate that refresh produces an empty plan (no drift)
			{
				Config:             testAccGoogleTagGatewayBasicConfig(zoneID, rnd),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				Config: testAccGoogleTagGatewayUpdateConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "false"),
				),
			},
			// Validate that refresh produces an empty plan after update
			{
				Config:             testAccGoogleTagGatewayUpdateConfig(zoneID, rnd),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func testAccGoogleTagGatewayBasicConfig(zoneID, name string) string {
	return acctest.LoadTestCase("basic.tf", zoneID, name)
}

func testAccGoogleTagGatewayMinimalConfig(zoneID, name string) string {
	return acctest.LoadTestCase("minimal.tf", zoneID, name)
}

func testAccGoogleTagGatewayUpdateConfig(zoneID, name string) string {
	return acctest.LoadTestCase("update.tf", zoneID, name)
}

func testAccGoogleTagGatewayUpdateEnabledOnlyConfig(zoneID, name string) string {
	return acctest.LoadTestCase("update_enabled_only.tf", zoneID, name)
}

func testAccGoogleTagGatewayUpdateEndpointOnlyConfig(zoneID, name string) string {
	return acctest.LoadTestCase("update_endpoint_only.tf", zoneID, name)
}

func testAccGoogleTagGatewayUpdateMeasurementIDOnlyConfig(zoneID, name string) string {
	return acctest.LoadTestCase("update_measurement_id_only.tf", zoneID, name)
}
