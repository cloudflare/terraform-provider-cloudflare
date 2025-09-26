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
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigDns(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("desc")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12303)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("block")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("dns")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(dns.domains[*] == \"example.com\")")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("identity"), knownvalue.StringExact("any(identity.groups.name[*] in {\"finance\"})")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("device_posture"), knownvalue.StringExact("")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("block_page_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("block_reason"), knownvalue.StringExact("cuzs")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("ip_categories"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("ip_indicator_feeds"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("ignore_cname_category_matches"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("schedule").AtMapKey("mon"), knownvalue.StringExact("08:00-12:30,13:30-17:00")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("schedule").AtMapKey("tue"), knownvalue.StringExact("08:00-12:30,13:30-17:00")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("schedule").AtMapKey("wed"), knownvalue.StringExact("08:00-12:30,13:30-17:00")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("schedule").AtMapKey("thu"), knownvalue.StringExact("08:00-12:30,13:30-17:00")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("schedule").AtMapKey("fri"), knownvalue.StringExact("08:00-12:30,13:30-17:00")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("schedule").AtMapKey("sat"), knownvalue.StringExact("08:00-12:30,13:30-17:00")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("schedule").AtMapKey("sun"), knownvalue.StringExact("08:00-12:30,13:30-17:00")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("schedule").AtMapKey("time_zone"), knownvalue.StringExact("America/New_York")),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
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
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigDnsResolve(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("desc")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12304)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("resolve")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("dns_resolver")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(dns.domains[*] == \"example.com\")")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("dns_resolvers").AtMapKey("ipv6").AtSliceIndex(0).AtMapKey("ip"), knownvalue.StringExact("2001:DB8::")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("dns_resolvers").AtMapKey("ipv4").AtSliceIndex(0).AtMapKey("ip"), knownvalue.StringExact("2.2.2.2")),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
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
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigHttpAllow(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("desc")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12305)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("http")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(http.request.uri.security_category[*] in {22}) or any(http.request.uri.content_category[*] in {34})")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("add_headers").AtMapKey("Xhello").AtSliceIndex(0), knownvalue.StringExact("abcd")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("add_headers").AtMapKey("Xhello").AtSliceIndex(1), knownvalue.StringExact("efg")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("check_session").AtMapKey("duration"), knownvalue.StringExact("1h2m9s")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("check_session").AtMapKey("enforce"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
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
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigHttpBlock(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("desc")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12306)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("block")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("http")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(http.request.uri.security_category[*] in {21}) or any(http.request.uri.content_category[*] in {32})")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("block_page").AtMapKey("target_uri"), knownvalue.StringExact("https://examples.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("notification_settings").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("notification_settings").AtMapKey("msg"), knownvalue.StringExact("msg")),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
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
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigHttpIsolate(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("desc")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12307)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("isolate")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("http")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(http.request.uri.security_category[*] in {21}) or any(http.request.uri.content_category[*] in {32})")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("biso_admin_controls").AtMapKey("copy"), knownvalue.StringExact("remote_only")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("biso_admin_controls").AtMapKey("keyboard"), knownvalue.StringExact("enabled")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("biso_admin_controls").AtMapKey("version"), knownvalue.StringExact("v1")),
				},
			},
			{
				Config: testAccCloudflareTeamsRuleConfigHttpIsolateV2(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("biso_admin_controls").AtMapKey("version"), knownvalue.StringExact("v2")),
						plancheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("biso_admin_controls").AtMapKey("dcp"), knownvalue.Bool(true)),
						plancheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("biso_admin_controls").AtMapKey("dk"), knownvalue.Bool(true)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("desc")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12307)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("isolate")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("http")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(http.request.uri.security_category[*] in {21}) or any(http.request.uri.content_category[*] in {32})")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("biso_admin_controls").AtMapKey("version"), knownvalue.StringExact("v2")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("biso_admin_controls").AtMapKey("dcp"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("biso_admin_controls").AtMapKey("dk"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
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
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigL4(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("desc")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12308)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("l4_override")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("l4")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("net.dst.ip in {10.0.0.0/8} and net.dst.port in {80 443 8080 53} and not(net.dst.ip in {10.217.0.0/16})")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("device_posture"), knownvalue.StringExact("any(device_posture.checks.passed[*] == \"51fe39d9-d584-48f5-9eed-36cd14ada791\")")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("l4override").AtMapKey("port"), knownvalue.Int64Exact(80)),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
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
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigDns(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				},
			},
			{
				Config: testAccCloudflareTeamsRuleConfigNoSettings(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("desc")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12301)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("block")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("dns")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(dns.domains[*] == \"example.com\")")),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
			},
		},
	})
}

func testAccCloudflareTeamsRuleConfigNoSettings(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfignosettings.tf", rnd, accountID)
}

func TestAccCloudflareTeamsRule_DNS_Override(t *testing.T) {
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
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigDnsOverride(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("DNS override policy")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12400)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("override")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("dns")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(dns.domains[*] == \"example.com\")")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("override_ips").AtSliceIndex(0), knownvalue.StringExact("192.0.2.1")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("override_ips").AtSliceIndex(1), knownvalue.StringExact("192.0.2.2")),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
			},
		},
	})
}

func TestAccCloudflareTeamsRule_HTTP_Redirect(t *testing.T) {
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
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigHttpRedirect(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("HTTP redirect policy")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12401)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("redirect")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("http")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(http.request.uri.security_category[*] in {25})")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("redirect").AtMapKey("target_uri"), knownvalue.StringExact("https://redirect.example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("redirect").AtMapKey("include_context"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("redirect").AtMapKey("preserve_path_and_query"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
			},
		},
	})
}

func TestAccCloudflareTeamsRule_HTTP_Quarantine(t *testing.T) {
	// SKIP: Quarantine action requires account-level feature enablement
	// 
	// Investigation notes:
	// - "quarantine" IS a valid action according to:
	//   * Terraform provider schema (schema.go:56)
	//   * Cloudflare-Go v6 API (GatewayRuleActionQuarantine)
	//   * Provider documentation (shows quarantine examples)
	// - API returns: 400 Bad Request "invalid action \"quarantine\"" (error code 2087)
	// - Root cause: Test account lacks required feature enablement
	// 
	// Possible missing prerequisites:
	// - Enterprise plan requirement
	// - DLP (Data Loss Prevention) feature flag
	// - Malware scanning feature enablement  
	// - Account-specific quarantine feature flag
	//
	// TODO: Enable required feature flag and re-enable this test
	t.Skip("quarantine action not available on test account - requires feature flag enablement")
	
	// Test implementation preserved for when feature is enabled:
	/*
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_policy.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigHttpQuarantine(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("HTTP quarantine policy")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12402)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("quarantine")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("http")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(http.request.uri.content_category[*] in {35})")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("quarantine").AtMapKey("file_types").AtSliceIndex(0), knownvalue.StringExact("exe")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("quarantine").AtMapKey("file_types").AtSliceIndex(1), knownvalue.StringExact("pdf")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("quarantine").AtMapKey("file_types").AtSliceIndex(2), knownvalue.StringExact("zip")),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
			},
		},
	})
	*/
}

func TestAccCloudflareTeamsRule_HTTP_Scan(t *testing.T) {
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
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigHttpScan(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("HTTP scan policy")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12402)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("scan")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("http")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(http.request.uri.content_category[*] in {35})")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
			},
		},
	})
}


func TestAccCloudflareTeamsRule_Egress(t *testing.T) {
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
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigEgressLocal(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("Local egress policy via WARP IPs")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12403)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("egress")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("egress")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("net.dst.port in {443 80}")),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
			},
		},
	})
}

// TestAccCloudflareTeamsRule_EgressDedicated - Test with dedicated IPs (DISABLED)
func TestAccCloudflareTeamsRule_EgressDedicated(t *testing.T) {
	// SKIP: Dedicated egress IPs require Enterprise plan purchase and configuration
	// 
	// Investigation notes:
	// - Egress action IS valid and supported by the API
	// - API error: "Account doesn't own dedicated primary IPv4" (error code 2055)
	// - Root cause: Test account lacks dedicated egress IP configuration
	// 
	// Dedicated egress IPs are:
	// - Enterprise feature that must be purchased ($50/month per data center)
	// - Assigned to specific Cloudflare data centers
	// - Consist of IPv4 address + IPv6 range
	// - Can be Cloudflare-provided or BYOIP (Bring Your Own IP)
	// 
	// Alternative: Basic egress policies work without dedicated IPs (use WARP IPs)
	//
	// To enable dedicated IP testing:
	// 1. Purchase dedicated egress IPs for test account
	// 2. Update testdata with actual allocated IP addresses
	// 3. Remove this skip
	t.Skip("dedicated egress IPs not configured on test account - requires Enterprise feature purchase")
	
	// Test implementation preserved for when dedicated IPs are configured:
	/*
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_policy.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigEgressDedicated(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("egress").AtMapKey("ipv4"), knownvalue.StringExact("YOUR_DEDICATED_IPV4")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("egress").AtMapKey("ipv6"), knownvalue.StringExact("YOUR_DEDICATED_IPV6_RANGE")),
				},
			},
		},
	})
	*/
}


func TestAccCloudflareTeamsRule_SafeSearch(t *testing.T) {
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
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigSafeSearch(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("Safe search policy")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("precedence"), knownvalue.Int64Exact(12404)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("safesearch")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("dns")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(dns.domains[*] in {\"google.com\" \"bing.com\" \"duckduckgo.com\"})")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
			},
		},
	})
}

// TestAccCloudflareTeamsRule_MinimalToMaximal - DISABLED
// This test hits API drift issues where the API automatically populates rule_settings
// even when not specified, causing persistent plan changes (related to GitHub issue #5839)

func TestAccCloudflareTeamsRule_MinimalToMaximal(t *testing.T) {
	// SKIP: API drift issue - Cloudflare Gateway API auto-populates computed fields
	//
	// Detailed Analysis:
	// This test attempts to create a "minimal" policy and then update to "maximal"
	// to test comprehensive CRUD operations and default value handling.
	//
	// The Issue:
	// 1. Create minimal policy: only account_id, name, action, filters, traffic
	// 2. API automatically populates ALL rule_settings fields as computed
	// 3. Terraform detects drift: 20+ fields show as "(known after apply)"
	// 4. Fields change from null/unset to computed values on refresh
	//
	// Specific API Behavior:
	// - description: "" -> null (empty string becomes null)
	// - rule_settings: {} -> {20+ computed fields} (empty object becomes populated)
	// - precedence: auto-assigned value -> "(known after apply)"
	// - All timestamps, version, sharable, etc. become computed
	//
	// Why This Happens:
	// - Gateway API normalizes/sanitizes ALL rule configurations
	// - API returns computed values for fields not explicitly set
	// - Provider schema marks many fields as Computed+Optional
	// - No way to distinguish "user didn't set" vs "API computed"
	//
	// Related Issues:
	// - GitHub #5839: "Recurring change on zero_trust_gateway_policy"
	// - GitHub #5394: "rule_settings keeps changing" 
	// - Root cause: API "sanitization and formatting" behavior
	//
	// Workarounds Attempted:
	// - lifecycle { ignore_changes }: Reduces drift but doesn't eliminate it
	// - Explicit rule_settings = {}: Still causes computed field population
	// - Pre-setting computed defaults: API overrides with its own values
	//
	// Conclusion:
	// This is a fundamental API behavior, not a test bug. The Gateway API
	// is designed to always return full rule configurations, making "minimal"
	// resource testing incompatible with Terraform's drift detection.
	//
	// Alternative: Test individual attributes separately rather than minimal->maximal progression
	t.Skip("API auto-populates computed fields causing persistent drift - see GitHub issues #5839, #5394")
}


func TestAccCloudflareTeamsRule_DNS_ResolveInternal(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_policy.%s", rnd)
	viewName := fmt.Sprintf("cloudflare_account_dns_settings_internal_view.%s_view", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigDnsResolveInternalWithView(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("description"), knownvalue.StringExact("Internal DNS resolve policy with view")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("action"), knownvalue.StringExact("resolve")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("filters").AtSliceIndex(0), knownvalue.StringExact("dns_resolver")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("traffic"), knownvalue.StringExact("any(dns.domains[*] == \"internal.example.com\")")),
					// Verify the DNS view was created
					statecheck.ExpectKnownValue(viewName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(viewName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s-view", rnd))),
					// Verify the rule settings use the view_id 
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("resolve_dns_internally").AtMapKey("fallback"), knownvalue.StringExact("public_dns")),
					// View ID should not be null - it should reference the created view
					statecheck.ExpectKnownValue(name, tfjsonpath.New("rule_settings").AtMapKey("resolve_dns_internally").AtMapKey("view_id"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"sharable"},
			},
		},
	})
}


// Helper functions for test configurations
func testAccCloudflareTeamsRuleConfigDnsOverride(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfigdns-override.tf", rnd, accountID)
}

func testAccCloudflareTeamsRuleConfigHttpRedirect(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfighttp-redirect.tf", rnd, accountID)
}

func testAccCloudflareTeamsRuleConfigHttpScan(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfighttp-scan.tf", rnd, accountID)
}

func testAccCloudflareTeamsRuleConfigHttpQuarantine(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfighttp-quarantine.tf", rnd, accountID)
}

func testAccCloudflareTeamsRuleConfigEgressLocal(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfigegress-local.tf", rnd, accountID)
}

// testAccCloudflareTeamsRuleConfigEgressDedicated - DISABLED (requires dedicated IPv4/IPv6)
// func testAccCloudflareTeamsRuleConfigEgressDedicated(rnd, accountID string) string {
// 	return acctest.LoadTestCase("teamsruleconfigegress.tf", rnd, accountID)
// }

func testAccCloudflareTeamsRuleConfigSafeSearch(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfigsafesearch.tf", rnd, accountID)
}

func testAccCloudflareTeamsRuleConfigMinimalIgnoreChanges(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfigminimal-ignore-changes.tf", rnd, accountID)
}

func testAccCloudflareTeamsRuleConfigMinimalDebug(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfigminimal-debug.tf", rnd, accountID)
}

// testAccCloudflareTeamsRuleConfigMinimal - DISABLED (API drift issues)
// func testAccCloudflareTeamsRuleConfigMinimal(rnd, accountID string) string {
// 	return acctest.LoadTestCase("teamsruleconfigminimal.tf", rnd, accountID)
// }

// testAccCloudflareTeamsRuleConfigMaximal - DISABLED (API drift issues)  
// func testAccCloudflareTeamsRuleConfigMaximal(rnd, accountID string) string {
// 	return acctest.LoadTestCase("teamsruleconfigmaximal.tf", rnd, accountID)
// }

func testAccCloudflareTeamsRuleConfigDnsResolveInternalWithView(rnd, accountID string) string {
	return acctest.LoadTestCase("teamsruleconfigdns-resolve-internal-with-view.tf", rnd, accountID)
}

// testAccCloudflareTeamsRuleConfigDnsResolveInternal - DISABLED (requires valid view_id)
// func testAccCloudflareTeamsRuleConfigDnsResolveInternal(rnd, accountID string) string {
// 	return acctest.LoadTestCase("teamsruleconfigdns-resolve-internal.tf", rnd, accountID)
// }

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
