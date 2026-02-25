package zero_trust_device_default_profile_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
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

// testSweepCloudflareZeroTrustDeviceDefaultProfile resets the include/exclude split tunnel fields
// on the account's default device profile.
//
// The default device profile is a singleton per account — it cannot be deleted, only patched.
// Acceptance tests (particularly WithSplitTunnel variants) write include/exclude data to the
// default profile and may leave it behind if a test fails. This sweeper resets those fields to
// empty so subsequent test runs start from a clean baseline.
//
// API behavior:
// - Create/Update uses Policies.Default.Edit() (PATCH on existing singleton profile)
// - Delete is a no-op (the default profile cannot be removed)
//
// Run with: go test ./internal/services/zero_trust_device_default_profile/ -v -sweep=all
func testSweepCloudflareZeroTrustDeviceDefaultProfile(r string) error {
	ctx := context.Background()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return nil
	}

	client := acctest.SharedClient()

	// Reset include and exclude to empty slices, clearing any leftover split tunnel state
	// from previous test runs. Both cannot be set in the same request, so we clear include
	// first (setting it to []) then clear exclude separately.
	_, err := client.ZeroTrust.Devices.Policies.Default.Edit(ctx, zero_trust.DevicePolicyDefaultEditParams{
		AccountID: cloudflare.F(accountID),
		Include:   cloudflare.F([]zero_trust.SplitTunnelIncludeUnionParam{}),
	})
	if err != nil {
		fmt.Printf("failed to reset include on default device profile: %v\n", err)
	}

	_, err = client.ZeroTrust.Devices.Policies.Default.Edit(ctx, zero_trust.DevicePolicyDefaultEditParams{
		AccountID: cloudflare.F(accountID),
		Exclude:   cloudflare.F([]zero_trust.SplitTunnelExcludeUnionParam{}),
	})
	if err != nil {
		fmt.Printf("failed to reset exclude on default device profile: %v\n", err)
	}

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

func TestAccUpgradeZeroTrustDeviceDefaultProfile_FromPublishedV5(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	config := testAccCloudflareZeroTrustDeviceDefaultProfileBasic(accountID, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck_AccountID(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.16.0",
					},
				},
				Config: config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDeviceDefaultProfile_ServiceModeProxy(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_default_profile.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustDeviceDefaultProfileServiceModeProxy(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service_mode_v2").AtMapKey("mode"), knownvalue.StringExact("proxy")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service_mode_v2").AtMapKey("port"), knownvalue.Float64Exact(8080)),
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
				Config:   testAccCloudflareZeroTrustDeviceDefaultProfileServiceModeProxy(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareZeroTrustDeviceDefaultProfile_UnsetOptionals(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_default_profile.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustDeviceDefaultProfileWithOptionals(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_updates"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_to_leave"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(60)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(300)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("disable_auto_fallback"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude_office_ips"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("register_interface_ip_with_dns"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sccm_vpn_boundary_support"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("switch_locked"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tunnel_protocol"), knownvalue.StringExact("wireguard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("support_url"), knownvalue.StringExact("https://support.example.com")),
				},
			},
			{
				Config: testAccCloudflareZeroTrustDeviceDefaultProfileWithoutOptionals(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_updates"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_to_leave"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(180)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("disable_auto_fallback"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude_office_ips"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("register_interface_ip_with_dns"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("sccm_vpn_boundary_support"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("switch_locked"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tunnel_protocol"), knownvalue.StringExact("")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("support_url"), knownvalue.StringExact("")),
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
				Config:   testAccCloudflareZeroTrustDeviceDefaultProfileWithoutOptionals(accountID, rnd),
				PlanOnly: true,
			},
		},
	})
}

func testAccCloudflareZeroTrustDeviceDefaultProfileServiceModeProxy(accountID, rnd string) string {
	return acctest.LoadTestCase("devicedefaultprofileservicemodeproxy.tf", rnd, accountID)
}

func testAccCloudflareZeroTrustDeviceDefaultProfileWithOptionals(accountID, rnd string) string {
	return acctest.LoadTestCase("devicedefaultprofilewithoptionals.tf", rnd, accountID)
}

func testAccCloudflareZeroTrustDeviceDefaultProfileWithoutOptionals(accountID, rnd string) string {
	return acctest.LoadTestCase("devicedefaultprofilewithoutoptionals.tf", rnd, accountID)
}
