package zero_trust_device_custom_profile_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/zero_trust"
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

func TestAccCloudflareZeroTrustDeviceCustomProfile_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_custom_profile.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDeviceCustomProfileDestroy,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccCloudflareZeroTrustDeviceCustomProfileBasic(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match"), knownvalue.StringExact("identity.email == \"test@example.com\"")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.Float64Exact(100)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Test custom device profile")),
				},
			},
			// Import
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareZeroTrustDeviceCustomProfile_Complete(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_custom_profile.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDeviceCustomProfileDestroy,
		Steps: []resource.TestStep{
			// Create with all optional fields
			{
				Config: testAccCloudflareZeroTrustDeviceCustomProfileComplete(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("match"), knownvalue.StringExact("identity.groups.name == \"Engineering\"")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.Float64Exact(50)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Complete custom device profile with all settings")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_updates"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_to_leave"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(60)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(300)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("disable_auto_fallback"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude_office_ips"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("switch_locked"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tunnel_protocol"), knownvalue.StringExact("wireguard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("support_url"), knownvalue.StringExact("https://support.example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("lan_allow_minutes"), knownvalue.Float64Exact(30)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("lan_allow_subnet_size"), knownvalue.Float64Exact(24)),
				},
			},
			// Update
			{
				Config: testAccCloudflareZeroTrustDeviceCustomProfileUpdated(accountID, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s-updated", rnd))),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.Float64Exact(75)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s-updated", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.Float64Exact(75)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Updated custom device profile")),
				},
			},
			// Import
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareZeroTrustDeviceCustomProfile_WithExclude(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_custom_profile.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDeviceCustomProfileDestroy,
		Steps: []resource.TestStep{
			// Create with exclude
			{
				Config: testAccCloudflareZeroTrustDeviceCustomProfileWithExclude(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(2)),
				},
			},
			// Update to include
			{
				Config: testAccCloudflareZeroTrustDeviceCustomProfileWithInclude(accountID, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)),
				},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDeviceCustomProfile_ServiceMode(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_custom_profile.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustDeviceCustomProfileDestroy,
		Steps: []resource.TestStep{
			// Create with service_mode_v2
			{
				Config: testAccCloudflareZeroTrustDeviceCustomProfileWithServiceMode(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service_mode_v2", "mode"), knownvalue.StringExact("proxy")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service_mode_v2", "port"), knownvalue.Float64Exact(3128)),
				},
			},
		},
	})
}

func testAccCheckCloudflareZeroTrustDeviceCustomProfileDestroy(s *terraform.State) error {
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_device_custom_profile" {
			continue
		}

		_, err := client.ZeroTrust.Devices.Policies.Custom.Get(
			context.Background(),
			zero_trust.DevicePolicyCustomGetParams{
				AccountID: cloudflare.F(accountID),
				PolicyID:  cloudflare.F(rs.Primary.ID),
			},
		)
		if err == nil {
			return fmt.Errorf("Zero Trust Device Custom Profile still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCloudflareZeroTrustDeviceCustomProfileBasic(accountID, rnd string) string {
	return acctest.LoadTestCase("devicecustomprofilebasic.tf", rnd, accountID)
}

func testAccCloudflareZeroTrustDeviceCustomProfileComplete(accountID, rnd string) string {
	return acctest.LoadTestCase("devicecustomprofilecomplete.tf", rnd, accountID)
}

func testAccCloudflareZeroTrustDeviceCustomProfileUpdated(accountID, rnd string) string {
	return acctest.LoadTestCase("devicecustomprofileupdated.tf", rnd, accountID)
}

func testAccCloudflareZeroTrustDeviceCustomProfileWithExclude(accountID, rnd string) string {
	return acctest.LoadTestCase("devicecustomprofilewithexclude.tf", rnd, accountID)
}

func testAccCloudflareZeroTrustDeviceCustomProfileWithInclude(accountID, rnd string) string {
	return acctest.LoadTestCase("devicecustomprofilewithinclude.tf", rnd, accountID)
}

func testAccCloudflareZeroTrustDeviceCustomProfileWithServiceMode(accountID, rnd string) string {
	return acctest.LoadTestCase("devicecustomprofilewithservicemode.tf", rnd, accountID)
}