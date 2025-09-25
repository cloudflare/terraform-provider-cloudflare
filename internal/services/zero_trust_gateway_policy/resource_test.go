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

func testAccCloudflareTeamsRuleConfigDns(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfigdns.tf", rnd, accountID)
}

func testAccCloudflareTeamsRuleConfigDnsResolve(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfigdns-resolve.tf", rnd, accountID)
}

func testAccCloudflareTeamsRuleConfigHttpAllow(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfighttp.tf", rnd, accountID)
}

func testAccCloudflareTeamsRuleConfigHttpBlock(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfighttp-block.tf", rnd, accountID)
}

func testAccCloudflareTeamsRuleConfigHttpIsolate(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfighttp-isolate.tf", rnd, accountID)
}

func testAccCloudflareTeamsRuleConfigHttpIsolateV2(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfighttp-isolate-v2.tf", rnd, accountID)
}

func testAccCloudflareTeamsRuleConfigL4(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfigl4.tf", rnd, accountID)
}

func TestAccCloudflareTeamsRule_Dns(t *testing.T) {
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
				Config: testAccCloudflareTeamsRuleConfigDns(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12303"),
					resource.TestCheckResourceAttr(name, "action", "block"),
					resource.TestCheckResourceAttr(name, "filters.0", "dns"),
					resource.TestCheckResourceAttr(name, "traffic", "any(dns.domains[*] == \"example.com\")"),
					resource.TestCheckResourceAttr(name, "identity", "any(identity.groups.name[*] in {\"finance\"})"),
					resource.TestCheckResourceAttr(name, "device_posture", ""),
					resource.TestCheckResourceAttr(name, "rule_settings.block_page_enabled", "true"),
					resource.TestCheckResourceAttr(name, "rule_settings.block_reason", "cuzs"),
					resource.TestCheckNoResourceAttr(name, "rule_settings.insecure_disable_dnssec_validation"),
					resource.TestCheckResourceAttr(name, "rule_settings.ip_categories", "true"),
					resource.TestCheckResourceAttr(name, "rule_settings.ip_indicator_feeds", "true"),
					resource.TestCheckResourceAttr(name, "rule_settings.ignore_cname_category_matches", "true"),
					resource.TestCheckResourceAttr(name, "schedule.mon", "08:00-12:30,13:30-17:00"),
					resource.TestCheckResourceAttr(name, "schedule.tue", "08:00-12:30,13:30-17:00"),
					resource.TestCheckResourceAttr(name, "schedule.wed", "08:00-12:30,13:30-17:00"),
					resource.TestCheckResourceAttr(name, "schedule.thu", "08:00-12:30,13:30-17:00"),
					resource.TestCheckResourceAttr(name, "schedule.fri", "08:00-12:30,13:30-17:00"),
					resource.TestCheckResourceAttr(name, "schedule.sat", "08:00-12:30,13:30-17:00"),
					resource.TestCheckResourceAttr(name, "schedule.sun", "08:00-12:30,13:30-17:00"),
					resource.TestCheckResourceAttr(name, "schedule.time_zone", "America/New_York"),
				),
			},
		},
	})
}

func TestAccCloudflareTeamsRule_DNS_Resolve(t *testing.T) {
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
				Config: testAccCloudflareTeamsRuleConfigDnsResolve(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12304"),
					resource.TestCheckResourceAttr(name, "action", "resolve"),
					resource.TestCheckResourceAttr(name, "filters.0", "dns_resolver"),
					resource.TestCheckResourceAttr(name, "traffic", "any(dns.domains[*] == \"example.com\")"),
					resource.TestCheckResourceAttr(name, "rule_settings.dns_resolvers.ipv6.0.ip", "2001:DB8::"),
					resource.TestCheckResourceAttr(name, "rule_settings.dns_resolvers.ipv4.0.ip", "2.2.2.2"),
				),
			},
		},
	})
}

func TestAccCloudflareTeamsRule_HttpAllow(t *testing.T) {
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
				Config: testAccCloudflareTeamsRuleConfigHttpAllow(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12305"),
					resource.TestCheckResourceAttr(name, "action", "allow"),
					resource.TestCheckResourceAttr(name, "filters.0", "http"),
					resource.TestCheckResourceAttr(name, "traffic", "any(http.request.uri.security_category[*] in {22}) or any(http.request.uri.content_category[*] in {34})"),
					resource.TestCheckResourceAttr(name, "rule_settings.add_headers.Xhello.0", "abcd"),
					resource.TestCheckResourceAttr(name, "rule_settings.add_headers.Xhello.1", "efg"),
					resource.TestCheckResourceAttr(name, "rule_settings.untrusted_cert.action", "pass_through"),
					resource.TestCheckResourceAttr(name, "rule_settings.check_session.duration", "1h2m9s"),
					resource.TestCheckResourceAttr(name, "rule_settings.check_session.enforce", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareTeamsRule_HttpBlock(t *testing.T) {
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
				Config: testAccCloudflareTeamsRuleConfigHttpBlock(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12306"),
					resource.TestCheckResourceAttr(name, "action", "block"),
					resource.TestCheckResourceAttr(name, "filters.0", "http"),
					resource.TestCheckResourceAttr(name, "traffic", "any(http.request.uri.security_category[*] in {21}) or any(http.request.uri.content_category[*] in {32})"),
					resource.TestCheckResourceAttr(name, "rule_settings.block_page.target_uri", "https://examples.com"),
					resource.TestCheckResourceAttr(name, "rule_settings.notification_settings.enabled", "true"),
					resource.TestCheckResourceAttr(name, "rule_settings.notification_settings.msg", "msg"),
				),
			},
		},
	})
}

func TestAccCloudflareTeamsRule_HttpIsolate(t *testing.T) {
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
				Config: testAccCloudflareTeamsRuleConfigHttpIsolate(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12307"),
					resource.TestCheckResourceAttr(name, "action", "isolate"),
					resource.TestCheckResourceAttr(name, "filters.0", "http"),
					resource.TestCheckResourceAttr(name, "traffic", "any(http.request.uri.security_category[*] in {21}) or any(http.request.uri.content_category[*] in {32})"),
					resource.TestCheckResourceAttr(name, "rule_settings.biso_admin_controls.copy", "remote_only"),
					resource.TestCheckResourceAttr(name, "rule_settings.biso_admin_controls.keyboard", "enabled"),
					resource.TestCheckResourceAttr(name, "rule_settings.biso_admin_controls.version", "v1"),
				),
			},
			{
				Config: testAccCloudflareTeamsRuleConfigHttpIsolateV2(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12307"),
					resource.TestCheckResourceAttr(name, "action", "isolate"),
					resource.TestCheckResourceAttr(name, "filters.0", "http"),
					resource.TestCheckResourceAttr(name, "traffic", "any(http.request.uri.security_category[*] in {21}) or any(http.request.uri.content_category[*] in {32})"),
					resource.TestCheckResourceAttr(name, "rule_settings.biso_admin_controls.version", "v2"),
					resource.TestCheckResourceAttr(name, "rule_settings.biso_admin_controls.dcp", "true"),
					resource.TestCheckResourceAttr(name, "rule_settings.biso_admin_controls.dk", "true"),
					resource.TestCheckNoResourceAttr(name, "rule_settings.biso_admin_controls.dd"),
					resource.TestCheckNoResourceAttr(name, "rule_settings.biso_admin_controls.du"),
				),
			},
		},
	})
}

func TestAccCloudflareTeamsRule_L4(t *testing.T) {
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
				Config: testAccCloudflareTeamsRuleConfigL4(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12308"),
					resource.TestCheckResourceAttr(name, "action", "l4_override"),
					resource.TestCheckResourceAttr(name, "filters.0", "l4"),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "traffic", "net.dst.ip in {10.0.0.0/8} and net.dst.port in {80 443 8080 53} and not(net.dst.ip in {10.217.0.0/16})"),
					resource.TestCheckResourceAttr(name, "device_posture", "any(device_posture.checks.passed[*] == \"51fe39d9-d584-48f5-9eed-36cd14ada791\")"),
					resource.TestCheckResourceAttr(name, "rule_settings.l4override.port", "80"),
				),
			},
		},
	})
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
				Config: testAccCloudflareTeamsRuleConfigDns(rnd, accountID),
			},
			{
				Config: testAccCloudflareTeamsRuleConfigNoSettings(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12301"),
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
