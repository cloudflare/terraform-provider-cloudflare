package zero_trust_device_default_profile_test

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
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_device_default_profile", &resource.Sweeper{
		Name: "cloudflare_zero_trust_device_default_profile",
		F:    testSweepCloudflareZeroTrustDeviceDefaultProfile,
	})
}

// testSweepCloudflareZeroTrustDeviceDefaultProfile is a no-op sweeper for the default device profile.
//
// The default device profile is a singleton configuration per account - there's only one default
// profile per account. Tests modify the existing default profile rather than creating new resources.
// Since nothing accumulates, no sweeping is required.
//
// This sweeper is registered to maintain consistency with other resources, but performs no actions.
//
// API behavior:
// - Create operation uses Policies.Default.Edit() (modifies existing profile)
// - Delete operation is a no-op (can't delete the default profile)
// - Resource represents account-level configuration, not creatable/deletable resources
//
// Run with: go test ./internal/services/zero_trust_device_default_profile/ -v -sweep=all
// (No cleanup will be performed)
func testSweepCloudflareZeroTrustDeviceDefaultProfile(r string) error {
	ctx := context.Background()
	tflog.Info(ctx, "Zero Trust Device Default Profile doesn't require sweeping (singleton account setting)")
	return nil
}

func TestAccCloudflareZeroTrustDeviceDefaultProfile_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_default_profile.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccCloudflareZeroTrustDeviceDefaultProfileBasic(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_updates"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_to_leave"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(60)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(300)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("switch_locked"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"exclude", "include", "fallback_domains", "default", "gateway_unique_id", "service_mode_v2"},
				ImportStateId:           accountID,
			},
			{
				Config:   testAccCloudflareZeroTrustDeviceDefaultProfileBasic(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareZeroTrustDeviceDefaultProfile_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_default_profile.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccCloudflareZeroTrustDeviceDefaultProfileBasic(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(60)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
				},
			},
			// Update
			{
				Config: testAccCloudflareZeroTrustDeviceDefaultProfileUpdated(accountID, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(false)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(120)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(120)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("support_url"), knownvalue.StringExact("https://updated-support.example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"exclude", "include", "fallback_domains", "default", "gateway_unique_id", "service_mode_v2"},
				ImportStateId:           accountID,
			},
			{
				Config:   testAccCloudflareZeroTrustDeviceDefaultProfileUpdated(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareZeroTrustDeviceDefaultProfile_Complete(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_default_profile.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustDeviceDefaultProfileComplete(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(60)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("disable_auto_fallback"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude_office_ips"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tunnel_protocol"), knownvalue.StringExact("wireguard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("support_url"), knownvalue.StringExact("https://support.example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("lan_allow_minutes"), knownvalue.Float64Exact(30)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("lan_allow_subnet_size"), knownvalue.Float64Exact(24)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"exclude", "include", "fallback_domains", "default", "gateway_unique_id", "service_mode_v2", "lan_allow_minutes", "lan_allow_subnet_size"},
				ImportStateId:           accountID,
			},
			{
				Config:   testAccCloudflareZeroTrustDeviceDefaultProfileComplete(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareZeroTrustDeviceDefaultProfile_WithExclude(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_default_profile.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustDeviceDefaultProfileWithExclude(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(60)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(2)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"exclude", "include", "fallback_domains", "default", "gateway_unique_id", "service_mode_v2"},
				ImportStateId:           accountID,
			},
			{
				Config:   testAccCloudflareZeroTrustDeviceDefaultProfileWithExclude(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareZeroTrustDeviceDefaultProfile_WithInclude(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_default_profile.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustDeviceDefaultProfileWithInclude(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(60)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"exclude", "include", "fallback_domains", "default", "gateway_unique_id", "service_mode_v2"},
				ImportStateId:           accountID,
			},
			{
				Config:   testAccCloudflareZeroTrustDeviceDefaultProfileWithInclude(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func testAccCloudflareZeroTrustDeviceDefaultProfileBasic(accountID, rnd string) string {
	return acctest.LoadTestCase("devicedefaultprofilebasic.tf", rnd, accountID)
}

func testAccCloudflareZeroTrustDeviceDefaultProfileComplete(accountID, rnd string) string {
	return acctest.LoadTestCase("devicedefaultprofilecomplete.tf", rnd, accountID)
}

func testAccCloudflareZeroTrustDeviceDefaultProfileUpdated(accountID, rnd string) string {
	return acctest.LoadTestCase("devicedefaultprofileupdated.tf", rnd, accountID)
}

func testAccCloudflareZeroTrustDeviceDefaultProfileWithExclude(accountID, rnd string) string {
	return acctest.LoadTestCase("devicedefaultprofilewithexclude.tf", rnd, accountID)
}

func testAccCloudflareZeroTrustDeviceDefaultProfileWithInclude(accountID, rnd string) string {
	return acctest.LoadTestCase("devicedefaultprofilewithinclude.tf", rnd, accountID)
}
