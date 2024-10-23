package zero_trust_gateway_settings_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareTeamsAccounts_ConfigurationBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_settings.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsAccountBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "settings.custom_certificate.enabled", "false"),
					resource.TestCheckResourceAttr(name, "settings.tls_decrypt.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.protocol_detection.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.activity_log.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.fips.tls", "true"),
					resource.TestCheckResourceAttr(name, "settings.block_page.name", rnd),
					resource.TestCheckResourceAttr(name, "settings.block_page.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.block_page.footer_text", "hello"),
					resource.TestCheckResourceAttr(name, "settings.block_page.header_text", "hello"),
					resource.TestCheckResourceAttr(name, "settings.block_page.mailto_subject", "hello"),
					resource.TestCheckResourceAttr(name, "settings.block_page.mailto_address", "test@cloudflare.com"),
					resource.TestCheckResourceAttr(name, "settings.block_page.background_color", "#000000"),
					resource.TestCheckResourceAttr(name, "settings.block_page.logo_path", "https://example.com"),
					resource.TestCheckResourceAttr(name, "settings.antivirus.enabled_download_phase", "true"),
					resource.TestCheckResourceAttr(name, "settings.antivirus.enabled_upload_phase", "false"),
					resource.TestCheckResourceAttr(name, "settings.antivirus.fail_closed", "true"),
					resource.TestCheckResourceAttr(name, "settings.antivirus.notification_settings.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.antivirus.notification_settings.msg", "msg"),
					resource.TestCheckResourceAttr(name, "settings.antivirus.notification_settings.support_url", "https://hello.com/"),
					resource.TestCheckResourceAttr(name, "settings.body_scanning.inspection_mode", "deep"),
					resource.TestCheckResourceAttr(name, "settings.browser_isolation.non_identity_enabled", "false"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsAccountBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsaccountbasic.tf", rnd, accountID)
}
