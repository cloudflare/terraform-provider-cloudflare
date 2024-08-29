package zero_trust_gateway_policy_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareTeamsRule_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_policy.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12302"),
					resource.TestCheckResourceAttr(name, "action", "block"),
					resource.TestCheckResourceAttr(name, "filters.0", "dns"),
					resource.TestCheckResourceAttr(name, "traffic", "any(dns.domains[*] == \"example.com\")"),
					resource.TestCheckResourceAttr(name, "rule_settings.#", "1"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.block_page_enabled", "true"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.block_page_reason", "cuz"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.insecure_disable_dnssec_validation", "false"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.egress.0.ipv4", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.egress.0.ipv6", "2001:db8::/32"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.untrusted_cert.0.action", "error"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.payload_log.0.enabled", "true"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsRuleConfigBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfigbasic.tf", rnd, accountID)
}

func TestAccCloudflareTeamsRule_NoSettings(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_policy.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigBasic(rnd, accountID),
			},
			{
				Config: testAccCloudflareTeamsRuleConfigNoSettings(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12302"),
					resource.TestCheckResourceAttr(name, "action", "block"),
					resource.TestCheckResourceAttr(name, "filters.0", "dns"),
					resource.TestCheckResourceAttr(name, "traffic", "any(dns.domains[*] == \"example.com\")"),
					resource.TestCheckResourceAttr(name, "rule_settings.#", "0"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsRuleConfigNoSettings(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfignosettings.tf", rnd, accountID)
}

func testAccCheckCloudflareTeamsRuleDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_gateway_policy" {
			continue
		}

		_, err := client.TeamsRule(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("teams rule still exists")
		}
	}

	return nil
}
