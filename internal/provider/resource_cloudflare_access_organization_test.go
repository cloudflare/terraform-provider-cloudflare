package provider

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareAccessOrganization(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_organization.%s", rnd)

	updatedName := fmt.Sprintf("%s-updated", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheck(t)
			testAccessAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessOrganizationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "auth_domain", fmt.Sprintf("%s.cloudflareaccess.com", rnd)),
					resource.TestCheckResourceAttr(name, "is_ui_read_only", "true"),
					resource.TestCheckResourceAttr(name, "login_design.#", "1"),
					resource.TestCheckResourceAttr(name, "login_design.0.background_color", "#FFFFFF"),
					resource.TestCheckResourceAttr(name, "login_design.0.text_color", "#000000"),
					resource.TestCheckResourceAttr(name, "login_design.0.logo_path", "https://example.com/logo.png"),
					resource.TestCheckResourceAttr(name, "login_design.0.header_text", "My header text"),
					resource.TestCheckResourceAttr(name, "login_design.0.footer_text", "My footer text"),
				),
			},
			{
				Config: testAccCloudflareAccessOrganizationConfigBasicUpdated(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", updatedName),
					resource.TestCheckResourceAttr(name, "auth_domain", fmt.Sprintf("%s.cloudflareaccess.com", rnd)),
					resource.TestCheckResourceAttr(name, "is_ui_read_only", "false"),
					resource.TestCheckResourceAttr(name, "login_design.#", "1"),
					resource.TestCheckResourceAttr(name, "login_design.0.background_color", "#FFFFFF"),
					resource.TestCheckResourceAttr(name, "login_design.0.text_color", "#000000"),
					resource.TestCheckResourceAttr(name, "login_design.0.logo_path", "https://example.com/logo.png"),
					resource.TestCheckResourceAttr(name, "login_design.0.header_text", "My header text"),
					resource.TestCheckResourceAttr(name, "login_design.0.footer_text", "My footer text"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessOrganizationDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_access_organization" {
			continue
		}

		var notFoundError *cloudflare.NotFoundError
		if rs.Primary.Attributes["zone_id"] != "" {
			_, _, err := client.ZoneLevelAccessOrganization(context.Background(), rs.Primary.Attributes["zone_id"])
			if !errors.As(err, &notFoundError) {
				return fmt.Errorf("AccessOrganization still exists")
			}
		}

		if rs.Primary.Attributes["account_id"] != "" {
			_, _, err := client.AccessOrganization(context.Background(), rs.Primary.Attributes["account_id"])
			if !errors.As(err, &notFoundError) {
				return fmt.Errorf("AccessOrganization still exists")
			}
		}

	}

	return nil
}

func testAccCloudflareAccessOrganizationConfigBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
		resource "cloudflare_access_organization" "%[1]s" {
			account_id                = "%[2]s"
			name                      = "%[1]s"
			auth_domain               = "%[1]s.cloudflareaccess.com"
			is_ui_read_only           = true

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

func testAccCloudflareAccessOrganizationConfigBasicUpdated(rnd, accountID string) string {
	return fmt.Sprintf(`
		resource "cloudflare_access_organization" "%[1]s" {
			account_id                = "%[2]s"
			name                      = "%[1]s-updated"
			auth_domain               = "%[1]s.cloudflareaccess.com"
			is_ui_read_only           = false

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
