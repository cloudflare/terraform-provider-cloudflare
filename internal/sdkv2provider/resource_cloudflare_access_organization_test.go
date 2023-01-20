package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
				Config:           testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID),
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
		{field: "user_seat_expiration_inactive_time", stateValue: attrs["user_seat_expiration_inactive_time"], expectedValue: "1460h"},
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
