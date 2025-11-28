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
