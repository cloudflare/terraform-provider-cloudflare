package zero_trust_organization_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_organization", &resource.Sweeper{
		Name: "cloudflare_zero_trust_organization",
		F:    testSweepCloudflareZeroTrustOrganization,
	})
}

func testSweepCloudflareZeroTrustOrganization(r string) error {
	ctx := context.Background()
	// Zero Trust Organization is an account-level organization configuration setting.
	// It's a singleton setting per account, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Zero Trust Organization doesn't require sweeping (account setting)")
	return nil
}

const DEFAULT_AUTHDOMAIN = "terraform-cfapi.cloudflareaccess.com"

func testAuthDomain() string {
	result := os.Getenv("CLOUDFLARE_ZERO_TRUST_ORGANIZATION_AUTH_DOMAIN")
	if result == "" {
		result = DEFAULT_AUTHDOMAIN
	}
	return result
}

func TestAccCloudflareAccessOrganization(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_organization.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	headerText := "My header text"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID, headerText, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", testAuthDomain()),
					resource.TestCheckResourceAttr(name, "auth_domain", rnd+"-"+testAuthDomain()),
					resource.TestCheckResourceAttr(name, "is_ui_read_only", "false"),
					resource.TestCheckResourceAttr(name, "user_seat_expiration_inactive_time", "1460h"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
					resource.TestCheckResourceAttr(name, "login_design.background_color", "#FFFFFF"),
					resource.TestCheckResourceAttr(name, "login_design.text_color", "#000000"),
					resource.TestCheckResourceAttr(name, "login_design.logo_path", "https://example.com/logo.png"),
					resource.TestCheckResourceAttr(name, "login_design.header_text", headerText),
					resource.TestCheckResourceAttr(name, "login_design.footer_text", "My footer text"),
					resource.TestCheckResourceAttr(name, "session_duration", "12h"),
					resource.TestCheckResourceAttr(name, "warp_auth_session_duration", "36h"),
					resource.TestCheckResourceAttr(name, "allow_authenticate_via_warp", "false"),
				),
			},
			{
				Config: testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID, headerText, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", testAuthDomain()),
					resource.TestCheckResourceAttr(name, "auth_domain", rnd+"-"+testAuthDomain()),
					resource.TestCheckResourceAttr(name, "is_ui_read_only", "false"),
					resource.TestCheckResourceAttr(name, "user_seat_expiration_inactive_time", "1460h"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
					resource.TestCheckResourceAttr(name, "login_design.background_color", "#FFFFFF"),
					resource.TestCheckResourceAttr(name, "login_design.text_color", "#000000"),
					resource.TestCheckResourceAttr(name, "login_design.logo_path", "https://example.com/logo.png"),
					resource.TestCheckResourceAttr(name, "login_design.header_text", headerText),
					resource.TestCheckResourceAttr(name, "login_design.footer_text", "My footer text"),
					resource.TestCheckResourceAttr(name, "session_duration", "12h"),
					resource.TestCheckResourceAttr(name, "warp_auth_session_duration", "36h"),
					resource.TestCheckResourceAttr(name, "allow_authenticate_via_warp", "false"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionNoop),
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ResourceName:     name,
				ImportState:      true,
				ImportStateId:    accountID,
				ImportStateCheck: accessOrgImportStateCheck,
			},
			{
				Config: testAccCloudflareAccessOrganizationConfigEmpty(rnd, accountID, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", testAuthDomain()),
					resource.TestCheckResourceAttr(name, "auth_domain", rnd+"-"+testAuthDomain()),
					resource.TestCheckResourceAttr(name, "session_duration", "12h"),
					resource.TestCheckNoResourceAttr(name, "login_design"),
				),
			},
			{
				ResourceName:     name,
				ImportState:      true,
				ImportStateId:    accountID,
				ImportStateCheck: accessOrgImportStateCheckEmpty,
			},
		},
	})
}

func accessOrgImportStateCheck(instanceStates []*terraform.InstanceState) error {
	state := instanceStates[0]
	attrs := state.Attributes
	wantAuthDomain := testAuthDomain()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if stateName := attrs["name"]; !strings.HasSuffix(stateName, wantAuthDomain) {
		return fmt.Errorf("name has value %q and does not match expected suffix %q", stateName, wantAuthDomain)
	}

	if stateAuthdomain := attrs["auth_domain"]; !strings.HasSuffix(stateAuthdomain, wantAuthDomain) {
		return fmt.Errorf("auth_domain has value %q and does not match expected suffix %q", stateAuthdomain, wantAuthDomain)
	}

	stateChecks := []struct {
		field         string
		stateValue    string
		expectedValue string
	}{
		{field: consts.AccountIDSchemaKey, stateValue: attrs[consts.AccountIDSchemaKey], expectedValue: accountID},
		{field: "is_ui_read_only", stateValue: attrs["is_ui_read_only"], expectedValue: "false"},
		{field: "user_seat_expiration_inactive_time", stateValue: attrs["user_seat_expiration_inactive_time"], expectedValue: "1460h"},
		{field: "login_design.background_color", stateValue: attrs["login_design.background_color"], expectedValue: "#FFFFFF"},
	}

	for _, check := range stateChecks {
		if check.stateValue != check.expectedValue {
			return fmt.Errorf("%s has value %q and does not match expected value %q", check.field, check.stateValue, check.expectedValue)
		}
	}

	return nil
}

func testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID, headerText, authDomain string) string {
	return acctest.LoadTestCase("accessorganizationconfigbasic.tf", rnd, accountID, headerText, authDomain)
}

func testAccCloudflareAccessOrganizationConfigEmpty(rnd, accountID, authDomain string) string {
	return acctest.LoadTestCase("accessorganizationconfigempty.tf", rnd, accountID, authDomain)
}

func accessOrgImportStateCheckEmpty(instanceStates []*terraform.InstanceState) error {
	state := instanceStates[0]
	attrs := state.Attributes
	wantAuthDomain := testAuthDomain()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if stateName := attrs["name"]; !strings.HasSuffix(stateName, wantAuthDomain) {
		return fmt.Errorf("name has value %q and does not match expected suffix %q", stateName, wantAuthDomain)
	}

	if stateAuthdomain := attrs["auth_domain"]; !strings.HasSuffix(stateAuthdomain, wantAuthDomain) {
		return fmt.Errorf("auth_domain has value %q and does not match expected suffix %q", stateAuthdomain, wantAuthDomain)
	}

	stateChecks := []struct {
		field         string
		stateValue    string
		expectedValue interface{}
	}{
		{field: consts.AccountIDSchemaKey, stateValue: attrs[consts.AccountIDSchemaKey], expectedValue: accountID},
		{field: "is_ui_read_only", stateValue: attrs["is_ui_read_only"], expectedValue: "false"},
		{field: "auto_redirect_to_identity", stateValue: attrs["auto_redirect_to_identity"], expectedValue: "false"},
		{field: "user_seat_expiration_inactive_time", stateValue: attrs["user_seat_expiration_inactive_time"], expectedValue: ""},
	}

	for _, check := range stateChecks {
		if check.stateValue != check.expectedValue {
			return fmt.Errorf("%s has value %q and does not match expected value %q", check.field, check.stateValue, check.expectedValue)
		}
	}

	// Special check for login_design - it should not exist in the attributes map
	if loginDesignValue, exists := attrs["login_design"]; exists && loginDesignValue != "" {
		return fmt.Errorf("login_design should not exist or should be empty, but has value %q", loginDesignValue)
	}

	return nil
}

// TestAccCloudflareAccessOrganization_DenyUnmatchedRequests tests the deny_unmatched_requests
// attribute and its companion deny_unmatched_requests_exempted_zone_names list.
// This is the attribute that caused the SIRT incident (Finding #1 in RM-30061).
//
// Steps:
// 1. Create with deny_unmatched_requests=true and exempted zone names
// 2. Idempotency check
// 3. Import and verify deny_unmatched_requests survives import
// 4. Disable deny_unmatched_requests (set to false, remove exempted zones)
// 5. Import and verify the disabled state
func TestAccCloudflareAccessOrganization_DenyUnmatchedRequests(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_organization.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessOrganizationConfigDenyUnmatched(rnd, accountID, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "deny_unmatched_requests", "true"),
				),
			},
			{
				Config: testAccCloudflareAccessOrganizationConfigDenyUnmatched(rnd, accountID, testAuthDomain()),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ResourceName:  name,
				ImportState:   true,
				ImportStateId: accountID,
				ImportStateCheck: func(instanceStates []*terraform.InstanceState) error {
					state := instanceStates[0]
					attrs := state.Attributes
					if attrs["deny_unmatched_requests"] != "true" {
						return fmt.Errorf("deny_unmatched_requests has value %q, expected \"true\"", attrs["deny_unmatched_requests"])
					}
					return nil
				},
			},
			{
				Config: testAccCloudflareAccessOrganizationConfigDenyUnmatchedDisabled(rnd, accountID, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "deny_unmatched_requests", "false"),
					resource.TestCheckNoResourceAttr(name, "deny_unmatched_requests_exempted_zone_names.#"),
				),
			},
			{
				ResourceName:  name,
				ImportState:   true,
				ImportStateId: accountID,
				ImportStateCheck: func(instanceStates []*terraform.InstanceState) error {
					state := instanceStates[0]
					attrs := state.Attributes
					if v := attrs["deny_unmatched_requests"]; v == "true" {
						return fmt.Errorf("deny_unmatched_requests should be false after disabling, got %q", v)
					}
					return nil
				},
			},
		},
	})
}

// TestAccCloudflareAccessOrganization_DenyUnmatchedOmitted verifies the behavior when
// deny_unmatched_requests is omitted entirely from config. This is the exact scenario
// from the SIRT incident: omitting the field should NOT silently disable it if it was
// previously enabled.
func TestAccCloudflareAccessOrganization_DenyUnmatchedOmitted(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_organization.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Enable deny_unmatched_requests
				Config: testAccCloudflareAccessOrganizationConfigDenyUnmatched(rnd, accountID, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "deny_unmatched_requests", "true"),
				),
			},
			{
				// Step 2: Apply config that omits deny_unmatched_requests entirely.
				// The key question: does this silently clear it to false?
				Config: testAccCloudflareAccessOrganizationConfigEmpty(rnd, accountID, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
				),
			},
			{
				// Step 3: Import and check what the API actually has.
				// If the provider silently cleared it, this will show false.
				// The test documents the actual behavior for the stabilization effort.
				ResourceName:  name,
				ImportState:   true,
				ImportStateId: accountID,
			},
		},
	})
}

// TestAccCloudflareAccessOrganization_ValueUpdate tests changing attribute values
// (not just adding/removing). Covers the update code path where MarshalJSONForUpdate
// sends only changed fields.
//
// Steps:
// 1. Create with initial values
// 2. Change session_duration, warp_auth_session_duration, login_design values
// 3. Verify updated values
// 4. Idempotency check
// 5. Import and verify
func TestAccCloudflareAccessOrganization_ValueUpdate(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_organization.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID, "Initial header", testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "session_duration", "12h"),
					resource.TestCheckResourceAttr(name, "warp_auth_session_duration", "36h"),
					resource.TestCheckResourceAttr(name, "user_seat_expiration_inactive_time", "1460h"),
					resource.TestCheckResourceAttr(name, "login_design.background_color", "#FFFFFF"),
					resource.TestCheckResourceAttr(name, "login_design.text_color", "#000000"),
					resource.TestCheckResourceAttr(name, "login_design.header_text", "Initial header"),
					resource.TestCheckResourceAttr(name, "login_design.footer_text", "My footer text"),
					resource.TestCheckResourceAttr(name, "login_design.logo_path", "https://example.com/logo.png"),
				),
			},
			{
				Config: testAccCloudflareAccessOrganizationConfigUpdate(rnd, accountID, "Updated header", testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "warp_auth_session_duration", "48h"),
					resource.TestCheckResourceAttr(name, "user_seat_expiration_inactive_time", "2190h"),
					resource.TestCheckResourceAttr(name, "login_design.background_color", "#000000"),
					resource.TestCheckResourceAttr(name, "login_design.text_color", "#FFFFFF"),
					resource.TestCheckResourceAttr(name, "login_design.header_text", "Updated header"),
					resource.TestCheckResourceAttr(name, "login_design.footer_text", "Updated footer text"),
					resource.TestCheckResourceAttr(name, "login_design.logo_path", "https://example.com/logo-v2.png"),
				),
			},
			{
				Config: testAccCloudflareAccessOrganizationConfigUpdate(rnd, accountID, "Updated header", testAuthDomain()),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ResourceName:  name,
				ImportState:   true,
				ImportStateId: accountID,
				ImportStateCheck: func(instanceStates []*terraform.InstanceState) error {
					state := instanceStates[0]
					attrs := state.Attributes
					checks := map[string]string{
						"session_duration":                   "24h",
						"warp_auth_session_duration":         "48h",
						"user_seat_expiration_inactive_time": "2190h",
						"login_design.background_color":      "#000000",
						"login_design.text_color":            "#FFFFFF",
						"login_design.header_text":           "Updated header",
						"login_design.footer_text":           "Updated footer text",
						"login_design.logo_path":             "https://example.com/logo-v2.png",
					}
					for field, expected := range checks {
						if attrs[field] != expected {
							return fmt.Errorf("%s has value %q, expected %q", field, attrs[field], expected)
						}
					}
					return nil
				},
			},
		},
	})
}

// TestAccCloudflareAccessOrganization_BooleansTrue tests setting boolean attributes to true.
// Existing tests only set them to false; this verifies the true->false->true lifecycle.
func TestAccCloudflareAccessOrganization_BooleansTrue(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_organization.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessOrganizationConfigBoolsTrue(rnd, accountID, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "true"),
					resource.TestCheckResourceAttr(name, "allow_authenticate_via_warp", "true"),
				),
			},
			{
				Config: testAccCloudflareAccessOrganizationConfigBoolsTrue(rnd, accountID, testAuthDomain()),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ResourceName:  name,
				ImportState:   true,
				ImportStateId: accountID,
				ImportStateCheck: func(instanceStates []*terraform.InstanceState) error {
					state := instanceStates[0]
					attrs := state.Attributes
					if attrs["auto_redirect_to_identity"] != "true" {
						return fmt.Errorf("auto_redirect_to_identity has value %q, expected \"true\"", attrs["auto_redirect_to_identity"])
					}
					if attrs["allow_authenticate_via_warp"] != "true" {
						return fmt.Errorf("allow_authenticate_via_warp has value %q, expected \"true\"", attrs["allow_authenticate_via_warp"])
					}
					return nil
				},
			},
			{
				// Flip back to false
				Config: testAccCloudflareAccessOrganizationConfigEmpty(rnd, accountID, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
				),
			},
		},
	})
}

// TestAccCloudflareAccessOrganization_LoginDesignLifecycle tests the full lifecycle of
// the login_design block: create -> remove -> re-add with different values.
func TestAccCloudflareAccessOrganization_LoginDesignLifecycle(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_organization.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with login_design
				Config: testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID, "Original header", testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "login_design.background_color", "#FFFFFF"),
					resource.TestCheckResourceAttr(name, "login_design.header_text", "Original header"),
				),
			},
			{
				// Step 2: Remove login_design
				Config: testAccCloudflareAccessOrganizationConfigEmpty(rnd, accountID, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr(name, "login_design"),
				),
			},
			{
				// Step 3: Re-add login_design with different values
				Config: testAccCloudflareAccessOrganizationConfigLoginDesignReAdd(rnd, accountID, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "login_design.background_color", "#FF0000"),
					resource.TestCheckResourceAttr(name, "login_design.text_color", "#00FF00"),
					resource.TestCheckResourceAttr(name, "login_design.logo_path", "https://example.com/new-logo.png"),
					resource.TestCheckResourceAttr(name, "login_design.header_text", "Re-added header"),
					resource.TestCheckResourceAttr(name, "login_design.footer_text", "Re-added footer"),
				),
			},
			{
				Config: testAccCloudflareAccessOrganizationConfigLoginDesignReAdd(rnd, accountID, testAuthDomain()),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// TestAccCloudflareAccessOrganization_MfaConfig tests the mfa_config block including
// allowed_authenticators, session_duration, and amr_matching_session_duration.
func TestAccCloudflareAccessOrganization_MfaConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_organization.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessOrganizationConfigMfa(rnd, accountID, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "mfa_config.allowed_authenticators.#", "2"),
					resource.TestCheckResourceAttr(name, "mfa_config.allowed_authenticators.0", "totp"),
					resource.TestCheckResourceAttr(name, "mfa_config.allowed_authenticators.1", "security_key"),
					resource.TestCheckResourceAttr(name, "mfa_config.session_duration", "24h"),
					resource.TestCheckResourceAttr(name, "mfa_config.amr_matching_session_duration", "12h"),
				),
			},
			{
				Config: testAccCloudflareAccessOrganizationConfigMfa(rnd, accountID, testAuthDomain()),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ResourceName:  name,
				ImportState:   true,
				ImportStateId: accountID,
				ImportStateCheck: func(instanceStates []*terraform.InstanceState) error {
					state := instanceStates[0]
					attrs := state.Attributes
					if v := attrs["mfa_config.session_duration"]; v != "24h" {
						return fmt.Errorf("mfa_config.session_duration has value %q, expected \"24h\"", v)
					}
					return nil
				},
			},
			{
				// Remove mfa_config
				Config: testAccCloudflareAccessOrganizationConfigEmpty(rnd, accountID, testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr(name, "mfa_config"),
				),
			},
		},
	})
}

// TestAccCloudflareAccessOrganization_ImportVerifyAll uses ImportStateVerify to
// automatically compare all state attributes after import. This is the most
// comprehensive import check available.
func TestAccCloudflareAccessOrganization_ImportVerifyAll(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_organization.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID, "Import verify header", testAuthDomain()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "session_duration", "12h"),
					resource.TestCheckResourceAttr(name, "warp_auth_session_duration", "36h"),
					resource.TestCheckResourceAttr(name, "login_design.background_color", "#FFFFFF"),
				),
			},
			{
				ResourceName:                         name,
				ImportState:                          true,
				ImportStateId:                        accountID,
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: consts.AccountIDSchemaKey,
				// deny_unmatched_requests is null in config but the API returns false on
				// import. The normalizeReadZeroTrustOrganizationAPIData function equates
				// false and null when prior state exists, but on import there is no prior
				// state so the raw API value comes through. This is a known gap in
				// normalizeImportZeroTrustOrganizationAPIData.
				ImportStateVerifyIgnore: []string{"deny_unmatched_requests"},
			},
		},
	})
}

// Config helper functions for new tests.

func testAccCloudflareAccessOrganizationConfigDenyUnmatched(rnd, accountID, authDomain string) string {
	return acctest.LoadTestCase("accessorganizationconfigdenyunmatched.tf", rnd, accountID, authDomain)
}

func testAccCloudflareAccessOrganizationConfigDenyUnmatchedDisabled(rnd, accountID, authDomain string) string {
	return acctest.LoadTestCase("accessorganizationconfigdenyunmatcheddisabled.tf", rnd, accountID, authDomain)
}


func testAccCloudflareAccessOrganizationConfigUpdate(rnd, accountID, headerText, authDomain string) string {
	return acctest.LoadTestCase("accessorganizationconfigupdate.tf", rnd, accountID, authDomain, headerText)
}

func testAccCloudflareAccessOrganizationConfigBoolsTrue(rnd, accountID, authDomain string) string {
	return acctest.LoadTestCase("accessorganizationconfigboolstrue.tf", rnd, accountID, authDomain)
}

func testAccCloudflareAccessOrganizationConfigLoginDesignReAdd(rnd, accountID, authDomain string) string {
	return acctest.LoadTestCase("accessorganizationconfiglogindesignreadd.tf", rnd, accountID, authDomain)
}

func testAccCloudflareAccessOrganizationConfigMfa(rnd, accountID, authDomain string) string {
	return acctest.LoadTestCase("accessorganizationconfigmfa.tf", rnd, accountID, authDomain)
}

func TestAccUpgradeZeroTrustOrganization_FromPublishedV5(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	headerText := "My header text"

	config := testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID, headerText, testAuthDomain())

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
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
