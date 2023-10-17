package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareAccessOrganization(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_organization.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", "terraform-cfapi.cloudflareaccess.com"),
					resource.TestCheckResourceAttr(name, "auth_domain", "terraform-cfapi.cloudflareaccess.com1"),
					resource.TestCheckResourceAttr(name, "is_ui_read_only", "false"),
					resource.TestCheckResourceAttr(name, "ui_read_only_toggle_reason", ""),
					resource.TestCheckResourceAttr(name, "user_seat_expiration_inactive_time", "1460h"),
					resource.TestCheckResourceAttr(name, "auto_redirect_to_identity", "false"),
					resource.TestCheckResourceAttr(name, "login_design.#", "1"),
					resource.TestCheckResourceAttr(name, "login_design.0.", "1"),
					resource.TestCheckResourceAttr(name, "login_design.0.background_color", "#FFFFFF"),
					resource.TestCheckResourceAttr(name, "login_design.0.text_color", "#000000"),
					resource.TestCheckResourceAttr(name, "login_design.0.logo_path", "https://example.com/logo.png"),
					resource.TestCheckResourceAttr(name, "login_design.0.header_text", "My header text"),
					resource.TestCheckResourceAttr(name, "login_design.0.footer_text", "My footer text"),
					resource.TestCheckResourceAttr(name, "session_duration", "12h"),
				),
				ResourceName:     name,
				ImportState:      true,
				ImportStateId:    accountID,
				ImportStateCheck: accessOrgImportStateCheck,
			},
		},
	})
}

func accessOrgImportStateCheck(instanceStates []*terraform.InstanceState) error {
	state := instanceStates[0]
	attrs := state.Attributes

	stateChecks := []struct {
		field         string
		stateValue    string
		expectedValue string
	}{
		{field: "ID", stateValue: state.ID, expectedValue: accountID},
		{field: consts.AccountIDSchemaKey, stateValue: attrs[consts.AccountIDSchemaKey], expectedValue: accountID},
		{field: "name", stateValue: attrs["name"], expectedValue: "terraform-cfapi.cloudflareaccess.com"},
		{field: "auth_domain", stateValue: attrs["auth_domain"], expectedValue: "terraform-cfapi.cloudflareaccess.com"},
		{field: "is_ui_read_only", stateValue: attrs["is_ui_read_only"], expectedValue: "false"},
		{field: "ui_read_only_toggle_reason", stateValue: attrs["ui_read_only_toggle_reason"], expectedValue: ""}, // UI read only is off so no message returned
		{field: "user_seat_expiration_inactive_time", stateValue: attrs["user_seat_expiration_inactive_time"], expectedValue: "1460h"},
		{field: "auto_redirect_to_identity", stateValue: attrs["auto_redirect_to_identity"], expectedValue: "false"},
		{field: "login_design.#", stateValue: attrs["login_design.#"], expectedValue: "1"},
	}

	for _, check := range stateChecks {
		if check.stateValue != check.expectedValue {
			return fmt.Errorf("%s has value %q and does not match expected value %q", check.field, check.stateValue, check.expectedValue)
		}
	}

	return nil
}

func testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
		resource "cloudflare_access_organization" "%[1]s" {
			account_id      = "%[2]s"
			name            = "terraform-cfapi.cloudflareaccess.com"
			auth_domain     = "terraform-cfapi.cloudflareaccess.com1"
			is_ui_read_only = false
			user_seat_expiration_inactive_time = "1460h"
			auto_redirect_to_identity = false
			session_duration = "12h"

			login_design {
				background_color = "#FFFFFF"
				text_color       = "#000000"
				logo_path        = "https://example.com/logo.png"
				header_text      = "My header text"
				footer_text      = "My footer text"
			}
		}
		`, rnd, accountID)
}
