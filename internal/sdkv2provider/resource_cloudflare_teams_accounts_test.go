package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareTeamsAccountConfigurationBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_teams_account.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsAccountBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "tls_decrypt_enabled", "true"),
					resource.TestCheckResourceAttr(name, "activity_log_enabled", "true"),
					resource.TestCheckResourceAttr(name, "fips.0.tls", "true"),
					resource.TestCheckResourceAttr(name, "block_page.0.name", rnd),
					resource.TestCheckResourceAttr(name, "block_page.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "block_page.0.footer_text", "hello"),
					resource.TestCheckResourceAttr(name, "block_page.0.header_text", "hello"),
					resource.TestCheckResourceAttr(name, "block_page.0.mailto_subject", "hello"),
					resource.TestCheckResourceAttr(name, "block_page.0.mailto_address", "test@cloudflare.com"),
					resource.TestCheckResourceAttr(name, "block_page.0.background_color", "#000000"),
					resource.TestCheckResourceAttr(name, "block_page.0.logo_path", "https://example.com"),
					resource.TestCheckResourceAttr(name, "logging.0.redact_pii", "true"),
					resource.TestCheckResourceAttr(name, "logging.0.settings_by_rule_type.0.dns.0.log_all", "false"),
					resource.TestCheckResourceAttr(name, "logging.0.settings_by_rule_type.0.dns.0.log_blocks", "true"),
					resource.TestCheckResourceAttr(name, "logging.0.settings_by_rule_type.0.http.0.log_all", "true"),
					resource.TestCheckResourceAttr(name, "logging.0.settings_by_rule_type.0.http.0.log_blocks", "true"),
					resource.TestCheckResourceAttr(name, "logging.0.settings_by_rule_type.0.l4.0.log_all", "false"),
					resource.TestCheckResourceAttr(name, "logging.0.settings_by_rule_type.0.l4.0.log_blocks", "true"),
					resource.TestCheckResourceAttr(name, "proxy.0.tcp", "true"),
					resource.TestCheckResourceAttr(name, "proxy.0.udp", "false"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsAccountBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_teams_account" "%[1]s" {
  account_id = "%[2]s"
  tls_decrypt_enabled = true
  activity_log_enabled = true
  block_page {
    name = "%[1]s"
    enabled = true
    footer_text = "hello"
    header_text = "hello"
    logo_path = "https://example.com"
    background_color = "#000000"
	mailto_subject = "hello"
	mailto_address = "test@cloudflare.com"
  }
  fips {
    tls = true
  }
  antivirus {
    enabled_download_phase = true
    enabled_upload_phase = false
    fail_closed = true
  }
  proxy {
    tcp = true
    udp = false
  }
  logging {
    redact_pii = true
    settings_by_rule_type {
      dns {
        log_all = false
        log_blocks = true
      }
      http {
        log_all = true
        log_blocks = true
      }
      l4 {
        log_all = false
        log_blocks = true
      }
    }
  }
}
`, rnd, accountID)
}
