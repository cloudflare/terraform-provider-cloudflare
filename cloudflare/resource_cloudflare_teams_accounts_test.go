package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudflareTeamsAccountConfigurationBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_teams_account.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsAccountConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "tls_decrypt_enabled", "true"),
					resource.TestCheckResourceAttr(name, "activity_log_enabled", "true"),
					resource.TestCheckResourceAttr(name, "block_page.0.name", rnd),
					resource.TestCheckResourceAttr(name, "block_page.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "block_page.0.footer_text", "hello"),
					resource.TestCheckResourceAttr(name, "block_page.0.header_text", "hello"),
					resource.TestCheckResourceAttr(name, "block_page.0.background_color", "#000000"),
					resource.TestCheckResourceAttr(name, "block_page.0.logo_path", "https://google.com"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsAccountConfigBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_teams_account" "%[1]s" {
  account_id  = "%[2]s"
  tls_decrypt_enabled = true
  activity_log_enabled = true
  block_page {
	name="%[1]s"
	enabled=true
    footer_text="hello"
    header_text="hello"
    logo_path="https://google.com"
    background_color="#000000"
  }
  antivirus {
	enabled_download_phase = true
	enabled_upload_phase = false
	fail_closed = true
	}
}
`, rnd, accountID)
}
