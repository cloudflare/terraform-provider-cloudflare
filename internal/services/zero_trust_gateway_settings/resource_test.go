package zero_trust_gateway_settings_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_gateway_settings", &resource.Sweeper{
		Name: "cloudflare_zero_trust_gateway_settings",
		F:    testSweepCloudflareZeroTrustGatewaySettings,
	})
}

func testSweepCloudflareZeroTrustGatewaySettings(r string) error {
	ctx := context.Background()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Warn(ctx, "CLOUDFLARE_ACCOUNT_ID not set, skipping gateway settings sweep")
		return nil
	}

	apiKey := os.Getenv("CLOUDFLARE_API_KEY")
	email := os.Getenv("CLOUDFLARE_EMAIL")
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")

	var client *cloudflare.Client
	var err error
	if apiToken != "" {
		client = cloudflare.NewClient(option.WithAPIToken(apiToken))
	} else {
		client = cloudflare.NewClient(
			option.WithAPIKey(apiKey),
			option.WithAPIEmail(email),
		)
	}
	if err != nil {
		return fmt.Errorf("error creating Cloudflare client: %w", err)
	}

	// Reset gateway settings to clean defaults so subsequent test runs start
	// from a known state. Tests that enable fips.tls, body_scanning deep
	// inspection, or tls_decrypt leave the account in a state that causes
	// error 2211 ("TLS decryption cannot be enabled without a certificate")
	// for any test that follows.
	tflog.Info(ctx, "Sweeping Zero Trust Gateway Settings — resetting to clean defaults")
	_, err = client.ZeroTrust.Gateway.Configurations.Update(ctx, zero_trust.GatewayConfigurationUpdateParams{
		AccountID: cloudflare.F(accountID),
		Settings: cloudflare.F(zero_trust.GatewayConfigurationSettingsParam{
			ActivityLog:       cloudflare.F(zero_trust.ActivityLogSettingsParam{Enabled: cloudflare.F(false)}),
			TLSDecrypt:        cloudflare.F(zero_trust.TLSSettingsParam{Enabled: cloudflare.F(false)}),
			ProtocolDetection: cloudflare.F(zero_trust.ProtocolDetectionParam{Enabled: cloudflare.F(false)}),
			BodyScanning:      cloudflare.F(zero_trust.BodyScanningSettingsParam{InspectionMode: cloudflare.F(zero_trust.BodyScanningSettingsInspectionModeShallow)}),
		}),
	})
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Failed to reset gateway settings: %s", err))
		// Non-fatal — log and continue so other sweepers still run
	} else {
		tflog.Info(ctx, "Zero Trust Gateway Settings reset to clean defaults")
	}
	return nil
}

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
	certID := os.Getenv("CLOUDFLARE_GATEWAY_CERTIFICATE_ID")
	if certID == "" {
		t.Skip("CLOUDFLARE_GATEWAY_CERTIFICATE_ID must be set for this acceptance test.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsAccountBasic(rnd, accountID, certID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "settings.tls_decrypt.enabled", "true"),
					resource.TestCheckResourceAttrSet(name, "settings.certificate.id"),
					resource.TestCheckResourceAttr(name, "settings.protocol_detection.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.activity_log.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.fips.tls", "true"),
					resource.TestCheckResourceAttr(name, "settings.block_page.name", rnd),
					resource.TestCheckResourceAttr(name, "settings.block_page.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.block_page.footer_text", "footer"),
					resource.TestCheckResourceAttr(name, "settings.block_page.header_text", "header"),
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
					resource.TestCheckResourceAttr(name, "settings.host_selector.enabled", "false"),
					resource.TestCheckResourceAttr(name, "settings.inspection.mode", "static"),
				),
			},
			{
				Config: testAccCloudflareTeamsAccountBasicMinimal1(rnd, accountID, certID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "settings.tls_decrypt.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.protocol_detection.enabled", "false"),
					resource.TestCheckResourceAttr(name, "settings.activity_log.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.antivirus.enabled_download_phase", "false"),
					resource.TestCheckResourceAttr(name, "settings.antivirus.enabled_upload_phase", "false"),
					resource.TestCheckResourceAttr(name, "settings.antivirus.fail_closed", "true"),
					resource.TestCheckResourceAttr(name, "settings.antivirus.notification_settings.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.antivirus.notification_settings.msg", "msg"),
					resource.TestCheckResourceAttr(name, "settings.antivirus.notification_settings.support_url", "https://hello.com/"),
					resource.TestCheckResourceAttr(name, "settings.extended_email_matching.enabled", "true"),
				),
			},
			{
				Config: testAccCloudflareTeamsAccountBasicMinimal2(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "settings.browser_isolation.url_browser_isolation_enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.browser_isolation.non_identity_enabled", "false"),
					resource.TestCheckResourceAttr(name, "settings.block_page.name", rnd),
					resource.TestCheckResourceAttr(name, "settings.block_page.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.block_page.footer_text", "footer"),
					resource.TestCheckResourceAttr(name, "settings.block_page.header_text", "header"),
					resource.TestCheckResourceAttr(name, "settings.block_page.mailto_subject", "hello"),
					resource.TestCheckResourceAttr(name, "settings.block_page.mailto_address", "test@cloudflare.com"),
					resource.TestCheckResourceAttr(name, "settings.block_page.background_color", "#000000"),
					resource.TestCheckResourceAttr(name, "settings.block_page.logo_path", "https://example.com"),
					resource.TestCheckResourceAttr(name, "settings.body_scanning.inspection_mode", "deep"),
					resource.TestCheckResourceAttr(name, "settings.extended_email_matching.enabled", "true"),
					resource.TestCheckResourceAttr(name, "settings.host_selector.enabled", "false"),
					resource.TestCheckResourceAttr(name, "settings.extended_email_matching.enabled", "true"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsAccountBasic(rnd, accountID, certID string) string {
	return acctest.LoadTestCase("teamsaccountbasic.tf", rnd, accountID, certID)
}

func testAccCloudflareTeamsAccountBasicMinimal1(rnd, accountID, certID string) string {
	return acctest.LoadTestCase("teamsaccountminimal1.tf", rnd, accountID, certID)
}

func testAccCloudflareTeamsAccountBasicMinimal2(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsaccountminimal2.tf", rnd, accountID)
}

func TestAccUpgradeZeroTrustGatewaySettings_FromPublishedV5(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	certID := os.Getenv("CLOUDFLARE_GATEWAY_CERTIFICATE_ID")
	if certID == "" {
		t.Skip("CLOUDFLARE_GATEWAY_CERTIFICATE_ID must be set for this acceptance test.")
	}

	config := testAccCloudflareTeamsAccountBasic(rnd, accountID, certID)

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
