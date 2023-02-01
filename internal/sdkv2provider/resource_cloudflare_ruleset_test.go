package sdkv2provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pkg/errors"
)

func init() {
	resource.AddTestSweepers("cloudflare_ruleset", &resource.Sweeper{
		Name: "cloudflare_ruleset",
		F:    testSweepCloudflareRuleset,
	})
}

func testSweepCloudflareRuleset(r string) error {
	ctx := context.Background()
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	// Clean up the account level rulesets
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	accountRulesets, accountRulesetsErr := client.ListAccountRulesets(context.Background(), accountID)
	if accountRulesetsErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Account Rulesets: %s", accountRulesetsErr))
	}

	if len(accountRulesets) == 0 {
		log.Print("[DEBUG] No Cloudflare Account Rulesets to sweep")
		return nil
	}

	for _, ruleset := range accountRulesets {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Account Ruleset ID: %s", ruleset.ID))
		client.DeleteAccountRuleset(context.Background(), accountID, ruleset.ID)
	}

	// .. and zone level rulesets
	zone := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zone == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	zoneRulesets, zoneRulesetsErr := client.ListZoneRulesets(context.Background(), zoneID)
	if zoneRulesetsErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Zone Rulesets: %s", zoneRulesetsErr))
	}

	if len(zoneRulesets) == 0 {
		log.Print("[DEBUG] No Cloudflare Zone Rulesets to sweep")
		return nil
	}

	for _, ruleset := range zoneRulesets {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Zone Ruleset ID: %s", ruleset.ID))
		client.DeleteZoneRuleset(context.Background(), zoneID, ruleset.ID)
	}

	return nil
}

func TestAccCloudflareRuleset_WAFBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetCustomWAFBasic(rnd, "my basic WAF ruleset", zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my basic WAF ruleset"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_custom"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "challenge"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "(ip.geoip.country eq \"GB\" or ip.geoip.country eq \"FR\") or cf.threat_score > 0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" ruleset rule description"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRuleset(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetManagedWAF(rnd, "my basic managed WAF ruleset", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my basic managed WAF ruleset"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "Execute Cloudflare Managed Ruleset on my zone-level phase ruleset"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetOWASP(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetManagedWAFOWASP(rnd, "Cloudflare OWASP managed ruleset", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Cloudflare OWASP managed ruleset"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "4814384a9e5d4991b9815dcfc25d2f1f"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "Execute Cloudflare Managed OWASP Ruleset on my zone-level phase ruleset"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetOWASPBlockXSSWithAnomalyOver60(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetManagedWAFOWASPBlockXSSAndAnomalyOver60(rnd, "Cloudflare OWASP managed ruleset blocking all XSS and anomaly scores over 60", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Cloudflare OWASP managed ruleset blocking all XSS and anomaly scores over 60"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.0.category", "xss"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.0.action", "block"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.0.status", "enabled"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.0.id", "4814384a9e5d4991b9815dcfc25d2f1f"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.0.overrides.0.rules.0.id", "6179ae15870a4bb7b2d480d4843b323c"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.0.overrides.0.rules.0.action", "block"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.0.overrides.0.rules.0.score_threshold", "60"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "zone"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetOWASPOnlyPL1(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetManagedWAFOWASPOnlyPL1(rnd, "Cloudflare OWASP managed ruleset only setting PL1", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Cloudflare OWASP managed ruleset only setting PL1"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "4814384a9e5d4991b9815dcfc25d2f1f"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.0.category", "paranoia-level-2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.0.status", "disabled"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.1.category", "paranoia-level-3"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.1.status", "disabled"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.2.category", "paranoia-level-4"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.2.status", "disabled"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.id", "6179ae15870a4bb7b2d480d4843b323c"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.action", "block"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.score_threshold", "60"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.status", "enabled"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "zone"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetDeployMultiple(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetManagedWAFDeployMultiple(rnd, "enable all Cloudflare managed rulesets", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "enable all Cloudflare managed rulesets"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "3"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "4814384a9e5d4991b9815dcfc25d2f1f"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.0.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.2.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.action_parameters.0.id", "c2e184081120413c86c3ab7e14069605"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.description", "zone deployment test"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetDeployMultipleWithSkip(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetManagedWAFDeployMultipleWithSkip(rnd, "enable all Cloudflare managed rulesets with a skip first", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "enable all Cloudflare managed rulesets with a skip first"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "4"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.ruleset", "current"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", fmt.Sprintf(`(http.host eq "%s" and http.request.method eq "GET")`, zoneName)),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "not this zone"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.0.id", "4814384a9e5d4991b9815dcfc25d2f1f"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.2.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.action_parameters.0.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.3.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.action_parameters.0.id", "c2e184081120413c86c3ab7e14069605"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.description", "zone deployment test"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetDeployMultipleWithTopSkipAndLastSkip(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetManagedWAFDeployMultipleWithTopSkipAndLastSkip(rnd, "enable all Cloudflare managed rulesets with a skip first", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "enable all Cloudflare managed rulesets with a skip first"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "5"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.ruleset", "current"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", fmt.Sprintf(`(http.host eq "%s" and http.request.uri.path contains "/app/")`, zoneName)),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "not this path"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.0.id", "4814384a9e5d4991b9815dcfc25d2f1f"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.2.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.action_parameters.0.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.3.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.action_parameters.0.id", "c2e184081120413c86c3ab7e14069605"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.4.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.4.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.4.action_parameters.0.ruleset", "current"),
					resource.TestCheckResourceAttr(resourceName, "rules.4.expression", fmt.Sprintf(`(http.host eq "%s" and http.request.uri.path contains "/httpbin/")`, zoneName)),
					resource.TestCheckResourceAttr(resourceName, "rules.4.description", "not this path either"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_SkipPhaseAndProducts(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetSkipPhaseAndProducts(rnd, "skip phases and product", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "skip phases and product"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "3"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.ruleset", "current"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", fmt.Sprintf(`http.host eq "%s"`, zoneName)),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "not this zone"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.0.phases.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "rules.1.action_parameters.0.phases.*", "http_ratelimit"),
					resource.TestCheckTypeSetElemAttr(resourceName, "rules.1.action_parameters.0.phases.*", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.2.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.action_parameters.0.products.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "rules.2.action_parameters.0.products.*", "zoneLockdown"),
					resource.TestCheckTypeSetElemAttr(resourceName, "rules.2.action_parameters.0.products.*", "uaBlock"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetWithCategoryAndRuleBasedOverrides(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetManagedWAFWithCategoryBasedOverrides(rnd, "my managed WAF ruleset with overrides", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my managed WAF ruleset with overrides"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "overrides to only enable wordpress rules to block"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.0.category", "wordpress"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.0.action", "block"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.1.category", "joomla"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.1.action", "block"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.id", "e3a567afc347477d9702d9047e97d760"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.status", "disabled"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetWithIDBasedOverrides(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetManagedWAFWithIDBasedOverrides(rnd, "my managed WAF ruleset with overrides", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my managed WAF ruleset with overrides"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "make 5de7edfa648c4d6891dc3e7f84534ffa and e3a567afc347477d9702d9047e97d760 log only"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "efb7b8c949ac4650a09736fc376e9aee"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.id", "5de7edfa648c4d6891dc3e7f84534ffa"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.action", "log"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.1.id", "e3a567afc347477d9702d9047e97d760"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.1.action", "log"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_MagicTransitUpdateWithHigherPriority(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	skipMagicTransitTestForNonConfiguredDefaultZone(t)

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_ruleset.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetMagicTransitSingle(rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", fmt.Sprintf("%s magic transit ruleset description", rnd)),
					resource.TestCheckResourceAttr(name, "rules.#", "1"),
					resource.TestCheckResourceAttr(name, "rules.0.action", "allow"),
					resource.TestCheckResourceAttr(name, "rules.0.description", "Allow TCP Ephemeral Ports"),
					resource.TestCheckResourceAttr(name, "rules.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "rules.0.expression", "tcp.dstport in { 32768..65535 }"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetMagicTransitMultiple(rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "rules.#", "2"),
					resource.TestCheckResourceAttr(name, "rules.0.action", "block"),
					resource.TestCheckResourceAttr(name, "rules.0.description", "Block UDP Ephemeral Ports"),
					resource.TestCheckResourceAttr(name, "rules.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "rules.0.expression", "udp.dstport in { 32768..65535 }"),
					resource.TestCheckResourceAttr(name, "rules.1.action", "allow"),
					resource.TestCheckResourceAttr(name, "rules.1.description", "Allow TCP Ephemeral Ports"),
					resource.TestCheckResourceAttr(name, "rules.1.enabled", "true"),
					resource.TestCheckResourceAttr(name, "rules.1.expression", "tcp.dstport in { 32768..65535 }"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetWithPayloadLogging(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetManagedWAFPayloadLogging(rnd, "my managed WAF ruleset with payload logging", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my managed WAF ruleset with payload logging"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.matched_data.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.matched_data.0.public_key", "not_a_real_public_key"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_RateLimit(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetRateLimit(rnd, "example HTTP rate limit", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example HTTP rate limit"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_ratelimit"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "block"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.response.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.response.0.status_code", "418"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.response.0.content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.response.0.content", "test content"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "(http.request.uri.path matches \"^/api/\")"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example http rate limit"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.characteristics.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.period", "60"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.requests_per_period", "100"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.mitigation_timeout", "60"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.requests_to_origin", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_RateLimitScorePerPeriod(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetRateLimitScorePerPeriod(rnd, "example HTTP rate limit by header score", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example HTTP rate limit by header score"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_ratelimit"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "block"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.response.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.response.0.status_code", "418"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.response.0.content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.response.0.content", "test content"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "(http.request.uri.path matches \"^/api/\")"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example http rate limit"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.characteristics.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.period", "60"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.score_per_period", "400"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.score_response_header_name", "my-score"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.mitigation_timeout", "60"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.requests_to_origin", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_PreserveRuleIDs(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	var adminRuleID, loginRuleID, adminRuleCopyID string

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// Create a ruleset with two rules (one for /admin, one for /login) and get their IDs.
				Config: testAccCheckCloudflareRulesetTwoCustomRules(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.id", getValue(&adminRuleID)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.id", getValue(&loginRuleID)),
				),
			},
			{
				// Reverse the order of rules. The IDs should remain the same, just in reverse order.
				Config: testAccCheckCloudflareRulesetTwoCustomRulesReversed(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.id", equalsValue(&loginRuleID)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.id", equalsValue(&adminRuleID)),
				),
			},
			{
				// Revert to the original version.
				Config: testAccCheckCloudflareRulesetTwoCustomRules(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.id", equalsValue(&adminRuleID)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.id", equalsValue(&loginRuleID)),
				),
			},
			{
				// Append a copy of the admin rule. The first two IDs should not change.
				Config: testAccCheckCloudflareRulesetThreeCustomRules(rnd, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.id", equalsValue(&adminRuleID)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.id", equalsValue(&loginRuleID)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.2.id", getValue(&adminRuleCopyID)),
				),
			},
			{
				// Disable the login rule. Its ID will change, but the admin rule IDs should remain the same.
				Config: testAccCheckCloudflareRulesetThreeCustomRules(rnd, zoneID, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.id", equalsValue(&adminRuleID)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.id", notEqualsValue(&loginRuleID)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.2.id", equalsValue(&adminRuleCopyID)),
				),
			},
		},
	})
}

func getValue(result *string) func(string) error {
	return func(value string) error {
		*result = value
		return nil
	}
}

func equalsValue(expected *string) func(string) error {
	return func(value string) error {
		if value != *expected {
			return fmt.Errorf("expected '%s' got '%s'", *expected, value)
		}
		return nil
	}
}

func notEqualsValue(expected *string) func(string) error {
	return func(value string) error {
		if value == *expected {
			return fmt.Errorf("expected != '%s'", *expected)
		}
		return nil
	}
}

func TestAccCloudflareRuleset_CustomErrors(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetCustomErrors(rnd, "example HTTP custom error response", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example HTTP custom error response"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_custom_errors"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "serve_error"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.content", "my example error page"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.status_code", "530"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "(http.request.uri.path matches \"^/api/\")"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example http custom error response"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_RequestOrigin(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetOrigin(rnd, "example HTTP request origin", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example HTTP request origin"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_origin"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "route"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.host_header", rnd+".terraform.cfapi.net"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.origin.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.origin.0.host", rnd+".terraform.cfapi.net"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.origin.0.port", "80"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.sni.0.value", rnd+".terraform.cfapi.net"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "(http.request.uri.path matches \"^/api/\")"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example http request origin"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_TransformationRuleURIPath(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetTransformationRuleURIPath(rnd, "transform rule for URI path", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "transform rule for URI path"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_transform"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "rewrite"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.uri.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.uri.0.path.0.value", "/static-rewrite"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_TransformationRuleURIQuery(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetTransformationRuleURIQuery(rnd, "transform rule for URI query", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "transform rule for URI query"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_transform"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "rewrite"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.uri.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.uri.0.query.0.value", "a=b"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_TransformHTTPResponseHeaders(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetExposedCredentialCheck(rnd, "example exposed credential check", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example exposed credential check"),
					resource.TestCheckResourceAttr(resourceName, "description", "This ruleset includes a rule checking for exposed credentials."),
					resource.TestCheckResourceAttr(resourceName, "kind", "custom"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_custom"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "log"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "http.request.method == \"POST\" && http.request.uri == \"/login.php\""),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example exposed credential check"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.exposed_credential_check.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.exposed_credential_check.0.username_expression", "url_decode(http.request.body.form[\"username\"][0])"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.exposed_credential_check.0.password_expression", "url_decode(http.request.body.form[\"password\"][0])"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_TransformationRuleURIPathAndQueryCombination(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetTransformationRuleURIPathAndQueryCombination(rnd, "uri combination of path and query", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "uri combination of path and query"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_transform"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "rewrite"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.uri.0.path.0.value", "/path/to/url"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.uri.0.query.0.expression", "concat(\"requestUrl=\", http.request.full_uri)"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example for combining URI action parameters for path and query"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.enabled", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_TransformationRuleRequestHeaders(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetTransformationRuleRequestHeaders(rnd, "transform rule for headers", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "transform rule for headers"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_late_transform"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "rewrite"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.#", "3"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.0.name", "example1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.0.value", "my-http-header-value1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.0.operation", "set"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.1.name", "example2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.1.operation", "set"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.1.expression", "cf.zone.name"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.2.name", "example3"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.2.operation", "remove"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_TransformationRuleResponseHeaders(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetTransformationRuleResponseHeaders(rnd, "transform rule for headers", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "transform rule for headers"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_response_headers_transform"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "rewrite"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.#", "3"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.0.name", "example1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.0.value", "my-http-header-value1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.0.operation", "set"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.1.name", "example2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.1.operation", "set"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.1.expression", "cf.zone.name"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.2.name", "example3"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.headers.2.operation", "remove"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_ActionParametersMultipleSkips(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetActionParametersMultipleSkips(rnd, "multiple skips for managed rulesets", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "multiple skips for managed rulesets"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "3"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.rulesets.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "(cf.zone.name eq \"domain.xyz\" and http.request.uri.query contains \"skip=rulesets\")"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "skip Cloudflare Manage ruleset"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.enabled", "true"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.0.rules.efb7b8c949ac4650a09736fc376e9aee", "5de7edfa648c4d6891dc3e7f84534ffa,e3a567afc347477d9702d9047e97d760"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.expression", "(cf.zone.name eq \"domain.xyz\" and http.request.uri.query contains \"skip=rules\")"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.description", "skip Wordpress rule and SQLi rule"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.enabled", "true"),

					resource.TestCheckResourceAttr(resourceName, "rules.2.action", "execute"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_ActionParametersOverridesAction(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesActionEnabled(rnd, "Overrides Cf Managed rules in Log", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Overrides Cf Managed rules in Log"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "Execute all rules in Cloudflare Managed Ruleset in log mode on my zone-level phase entry point ruleset"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.version", "latest"),
					resource.TestCheckNoResourceAttr(resourceName, "rules.0.action_parameters.0.enabled"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.action", "log"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.status", "enabled"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_ActionParametersHTTPDDoSOverride(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetActionParametersHTTPDDosOverride(rnd, "multiple skips for managed rulesets", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "multiple skips for managed rulesets"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "ddos_l7"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "4d21379b4f9f4bb088e0729962c8b3cf"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.id", "fdfdac75430c4c47a959592f0aa5e68a"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.sensitivity_level", "low"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "override HTTP DDoS ruleset rule"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.enabled", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_ActionParametersOverrideAllRulesetRules(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverrideSensitivityForAllRulesetRules(rnd, "overriding all ruleset rules sensitivity", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "overriding all ruleset rules sensitivity"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "ddos_l7"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "4d21379b4f9f4bb088e0729962c8b3cf"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.action", "log"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.sensitivity_level", "low"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "override HTTP DDoS ruleset rule"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.enabled", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_AccountLevelCustomWAFRule(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetAccountLevelCustomWAFRule(rnd, "account level custom rulesets", accountID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName+"_account_custom_firewall", "kind", "custom"),
					resource.TestCheckResourceAttr(resourceName+"_account_custom_firewall", "phase", "http_request_firewall_custom"),
					resource.TestCheckResourceAttr(resourceName+"_account_custom_firewall", "name", "Custom Ruleset for my account"),
					resource.TestCheckResourceAttr(resourceName+"_account_custom_firewall", "rules.0.action", "block"),

					resource.TestCheckResourceAttr(resourceName+"_account_custom_firewall_root", "kind", "root"),
					resource.TestCheckResourceAttr(resourceName+"_account_custom_firewall_root", "phase", "http_request_firewall_custom"),
					resource.TestCheckResourceAttr(resourceName+"_account_custom_firewall_root", "name", "Firewall Custom root"),
					resource.TestCheckResourceAttr(resourceName+"_account_custom_firewall_root", "rules.0.action", "execute"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_ExposedCredentialCheck(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetExposedCredentialCheck(rnd, "example exposed credential check", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example exposed credential check"),
					resource.TestCheckResourceAttr(resourceName, "description", "This ruleset includes a rule checking for exposed credentials."),
					resource.TestCheckResourceAttr(resourceName, "kind", "custom"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_custom"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "log"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "http.request.method == \"POST\" && http.request.uri == \"/login.php\""),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example exposed credential check"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.exposed_credential_check.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.exposed_credential_check.0.username_expression", "url_decode(http.request.body.form[\"username\"][0])"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.exposed_credential_check.0.password_expression", "url_decode(http.request.body.form[\"password\"][0])"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_Logging(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetDisableLoggingForSkipAction(rnd, "example disable logging for skip rule", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example disable logging for skip rule"),
					resource.TestCheckResourceAttr(resourceName, "description", "This ruleset includes a skip rule whose logging is disabled."),
					resource.TestCheckResourceAttr(resourceName, "kind", "root"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example disabled logging"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.logging.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.logging.0.status", "disabled"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_ConditionallySetActionParameterVersion(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetConditionallySetActionParameterVersion_ExecuteAlone(rnd, accountID, zoneName),
			},
			{
				Config: testAccCloudflareRulesetConditionallySetActionParameterVersion_ExecuteThenSkip(rnd, accountID, zoneName),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetWithActionManagedChallenge(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetManagedWAFWithCategoryBasedOverridesActionManagedChallenge(rnd, "my managed WAF ruleset with overrides", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my managed WAF ruleset with overrides"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "overrides to only enable wordpress rules to managed_challenge"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.0.category", "wordpress"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.0.action", "managed_challenge"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.0.status", "enabled"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.id", "e3a567afc347477d9702d9047e97d760"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.status", "enabled"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.action", "managed_challenge"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetManagedWAFWithActionManagedChallenge(rnd, "my basic managed WAF ruleset with action managed_challenge", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my basic managed WAF ruleset with action managed_challenge"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.action", "managed_challenge"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.status", ""),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "overrides change action to managed_challenge on the Cloudflare Manage Ruleset"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_LogCustomField(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetLogCustomField(rnd, "my basic log custom field ruleset", zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my basic log custom field ruleset"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_log_custom_fields"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "log_custom_field"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" log custom fields rule"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.request_fields.0", "content-type"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.request_fields.1", "host"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.request_fields.2", "x-forwarded-for"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.response_fields.0", "allow"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.response_fields.1", "content-type"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.response_fields.2", "server"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cookie_fields.0", "__cfruid"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cookie_fields.1", "__ga"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cookie_fields.2", "accountNumber"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_ActionParametersOverridesThrashingStatus(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.status", ""),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName, `status = "disabled"`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.status", "disabled"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.status", ""),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName, `status = "enabled"`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.status", "enabled"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName, `status = "disabled"`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.status", "disabled"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName, `status = "enabled"`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.status", "enabled"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.status", ""),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettings(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsAllEnabled(rnd, "my basic cache settings ruleset", zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my basic cache settings ruleset"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" set cache settings rule"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.mode", "override_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.default", "60"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.0.value", "50"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.0.status_code", "200"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.1.value", "30"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.1.status_code_range.0.from", "201"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.1.status_code_range.0.to", "300"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.browser_ttl.0.mode", "respect_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.serve_stale.0.disable_stale_while_updating", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.respect_strong_etags", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.ignore_query_strings_order", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.cache_deception_armor", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.query_string.0.exclude.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.query_string.0.exclude.0", "*"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.header.0.include.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.header.0.include.0", "habc"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.header.0.include.1", "hdef"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.header.0.check_presence.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.header.0.check_presence.0", "habc_t"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.header.0.check_presence.1", "hdef_t"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.header.0.exclude_origin", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.cookie.0.include.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.cookie.0.include.0", "cabc"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.cookie.0.include.1", "cdef"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.cookie.0.check_presence.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.cookie.0.check_presence.0", "cabc_t"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.cookie.0.check_presence.1", "cdef_t"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.user.0.device_type", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.user.0.geo", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.host.0.resolved", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.origin_error_page_passthru", "false"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsOptionalsEmpty(rnd, "my basic cache settings ruleset", zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my basic cache settings ruleset"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" set cache settings rule"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.mode", "override_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.default", "60"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.browser_ttl.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.browser_ttl.0.mode", "respect_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.cache_by_device_type", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.ignore_query_strings_order", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.cache_deception_armor", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.serve_stale.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.serve_stale.0.disable_stale_while_updating", "false"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsCustomKeyEmpty(rnd, "my basic cache settings ruleset", zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my basic cache settings ruleset"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" set cache settings rule"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.browser_ttl.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.serve_stale.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.query_string.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.query_string.0.include.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.query_string.0.exclude.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.header.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.header.0.include.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.header.0.check_presence.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.cookie.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.cookie.0.include.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.cookie.0.check_presence.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.user.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.user.0.device_type", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.user.0.lang", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.user.0.geo", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.host.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache_key.0.custom_key.0.host.0.resolved", "false"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsEdgeTTLRespectOrigin(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" set cache settings rule"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.status_code_ttl.0.value", "5"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.edge_ttl.0.mode", "respect_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.cache", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_Config(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetConfigAllEnabled(rnd, "my basic config ruleset", zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my basic config ruleset"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_config_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_config"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" set config rule"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.automatic_https_rewrites", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.autominify.0.html", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.autominify.0.css", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.autominify.0.js", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.bic", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.disable_apps", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.disable_zaraz", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.disable_railgun", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.email_obfuscation", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.mirage", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.opportunistic_encryption", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.polish", "off"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.rocket_loader", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.security_level", "off"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.server_side_excludes", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.ssl", "off"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.sxg", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.hotlink_protection", "true"),
				),
			},
		},
	})
}
func TestAccCloudflareRuleset_Redirect(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	accountId := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetRedirectFromList(rnd, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "kind", "root"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_redirect"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "redirect"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.from_list.0.name", "redirect_list_"+rnd),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.from_list.0.key", "http.request.full_uri"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_DynamicRedirect(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetRedirectFromValue(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_dynamic_redirect"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "redirect"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.from_value.0.status_code", "301"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.from_value.0.target_url.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.from_value.0.target_url.0.expression", ""),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.from_value.0.target_url.0.value", "some_host.com"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.from_value.0.preserve_query_string", "true"),
				),
			},
		},
	})
}

func testAccCheckCloudflareRulesetMagicTransitSingle(rnd, name, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    account_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s magic transit ruleset description"
    kind        = "root"
    phase       = "magic_transit"

    rules {
      action = "allow"
      expression = "tcp.dstport in { 32768..65535 }"
      description = "Allow TCP Ephemeral Ports"
    }
  }`, rnd, name, accountID)
}

func testAccCheckCloudflareRulesetMagicTransitMultiple(rnd, name, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    account_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s magic transit ruleset description"
    kind        = "root"
    phase       = "magic_transit"

    rules {
      action = "block"
      expression = "udp.dstport in { 32768..65535 }"
      description = "Block UDP Ephemeral Ports"
      status = "enabled"
    }

    rules {
      action = "allow"
      expression = "tcp.dstport in { 32768..65535 }"
      description = "Allow TCP Ephemeral Ports"
      status = "enabled"
    }
  }`, rnd, name, accountID)
}

func testAccCheckCloudflareRulesetCustomWAFBasic(rnd, name, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_custom"

    rules {
      action = "challenge"
      expression = "(ip.geoip.country eq \"GB\" or ip.geoip.country eq \"FR\") or cf.threat_score > 0"
      description = "%[1]s ruleset rule description"
    }
  }`, rnd, name, zoneID)
}

func testAccCheckCloudflareRulesetManagedWAF(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
      }
      expression = "true"
      description = "Execute Cloudflare Managed Ruleset on my zone-level phase ruleset"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFOWASP(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "execute"
      action_parameters {
        id = "4814384a9e5d4991b9815dcfc25d2f1f"
      }
      expression = "true"
      description = "Execute Cloudflare Managed OWASP Ruleset on my zone-level phase ruleset"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFOWASPBlockXSSAndAnomalyOver60(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    # enable all "XSS" rules
    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
        overrides {
          categories {
            category = "xss"
            action = "block"
            status = "enabled"
          }
        }
      }
      expression = "true"
      description = "zone"
      enabled = true
    }

    # set Anomaly Score for 60+ (low)
    rules {
      action = "execute"
      action_parameters {
        id = "4814384a9e5d4991b9815dcfc25d2f1f"
        overrides {
          rules {
            id = "6179ae15870a4bb7b2d480d4843b323c"
            action = "block"
            score_threshold = 60
          }
        }
      }
      expression = "true"
      description = "zone"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFOWASPOnlyPL1(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    # disable PL2, PL3 and PL4
    rules {
      action = "execute"
      action_parameters {
        id = "4814384a9e5d4991b9815dcfc25d2f1f"
        overrides {
          categories {
            category = "paranoia-level-2"
            status = "disabled"
          }

          categories {
            category = "paranoia-level-3"
            status = "disabled"
          }

          categories {
            category = "paranoia-level-4"
            status = "disabled"
          }

          rules {
            id = "6179ae15870a4bb7b2d480d4843b323c"
            action = "block"
            score_threshold = 60
            status = "enabled"
          }
        }
      }
      expression = "true"
      description = "zone"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFDeployMultiple(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "execute"
      action_parameters {
        id = "4814384a9e5d4991b9815dcfc25d2f1f"
      }
      expression = "true"
      description = "zone deployment test"
      enabled = true
    }

    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
      }
      expression = "true"
      description = "zone deployment test"
      enabled = true
    }

    rules {
      action = "execute"
      action_parameters {
        id = "c2e184081120413c86c3ab7e14069605"
      }
      expression = "true"
      description = "zone deployment test"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFDeployMultipleWithSkip(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "skip"
      action_parameters {
        ruleset = "current"
      }
      description = "not this zone"
      expression = "(http.host eq \"%[4]s\" and http.request.method eq \"GET\")"
      enabled = true
	  logging {
		status = "enabled"
	  }
    }

    rules {
      action = "execute"
      action_parameters {
        id = "4814384a9e5d4991b9815dcfc25d2f1f"
      }
      expression = "true"
      description = "zone deployment test"
      enabled = true
    }

    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
      }
      expression = "true"
      description = "zone deployment test"
      enabled = true
    }

    rules {
      action = "execute"
      action_parameters {
        id = "c2e184081120413c86c3ab7e14069605"
      }
      expression = "true"
      description = "zone deployment test"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFDeployMultipleWithTopSkipAndLastSkip(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "skip"
      action_parameters {
        ruleset = "current"
      }
      description = "not this path"
      expression = "(http.host eq \"%[4]s\" and http.request.uri.path contains \"/app/\")"
      enabled = true
	  logging {
		status = "enabled"
	  }
    }

    rules {
      action = "execute"
      action_parameters {
        id = "4814384a9e5d4991b9815dcfc25d2f1f"
      }
      expression = "true"
      description = "zone deployment test"
      enabled = true
    }

    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
      }
      expression = "true"
      description = "zone deployment test"
      enabled = true
    }

    rules {
      action = "execute"
      action_parameters {
        id = "c2e184081120413c86c3ab7e14069605"
      }
      expression = "true"
      description = "zone deployment test"
      enabled = true
    }

    rules {
      action = "skip"
      action_parameters {
        ruleset = "current"
      }
      description = "not this path either"
      expression = "(http.host eq \"%[4]s\" and http.request.uri.path contains \"/httpbin/\")"
      enabled = true
	  logging {
		status = "enabled"
	  }
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetSkipPhaseAndProducts(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "skip"
      action_parameters {
        ruleset = "current"
      }
      description = "not this zone"
      expression = "http.host eq \"%[4]s\""
      enabled = true
	  logging {
		status = "enabled"
	  }
    }

    rules {
      action = "skip"
      action_parameters {
        phases = ["http_ratelimit", "http_request_firewall_managed"]
      }
      expression = "http.request.uri.path contains \"/skip-phase/\""
      description = ""
      enabled = true
	  logging {
		status = "enabled"
	  }
    }

    rules {
      action = "skip"
      action_parameters {
        products = ["zoneLockdown", "uaBlock"]
      }
      expression = "http.request.uri.path contains \"/skip-products/\""
      description = ""
      enabled = true
	  logging {
		status = "enabled"
	  }
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFWithCategoryBasedOverrides(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
        overrides {
          categories {
            category = "wordpress"
            action = "block"
            status = "enabled"
          }

          categories {
            category = "joomla"
            action = "block"
            status = "enabled"
          }

			rules {
				id = "e3a567afc347477d9702d9047e97d760"
				status = "disabled"
			}
        }
      }

      expression = "true"
      description = "overrides to only enable wordpress rules to block"
      enabled = false
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFWithIDBasedOverrides(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
        overrides {
          rules {
            id = "5de7edfa648c4d6891dc3e7f84534ffa"
            action = "log"
            status = "enabled"
          }

          rules {
            id = "e3a567afc347477d9702d9047e97d760"
            action = "log"
            status = "enabled"
          }
        }
      }

      expression = "true"
      description = "make 5de7edfa648c4d6891dc3e7f84534ffa and e3a567afc347477d9702d9047e97d760 log only"
      enabled = false
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetTransformationRuleURIPath(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_transform"

    rules {
      action = "rewrite"
      action_parameters {
        uri {
          path {
            value = "/static-rewrite"
          }
        }
      }

      expression = "(http.host eq \"%[4]s\")"
      description = "URI transformation path example"
      enabled = false
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetTransformationRuleURIQuery(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_transform"

    rules {
      action = "rewrite"
      action_parameters {
        uri {
          query {
            value = "a=b"
          }
        }
      }

      expression = "(http.host eq \"%[4]s\")"
      description = "URI transformation query example"
      enabled = false
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetTransformationRuleRequestHeaders(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_late_transform"

    rules {
      action = "rewrite"
      action_parameters {
        headers {
          name      = "example1"
          operation = "set"
          value     = "my-http-header-value1"
        }

        headers {
          name       = "example2"
          operation  = "set"
          expression = "cf.zone.name"
        }

        headers {
          name      = "example3"
          operation = "remove"
        }
      }

      expression = "true"
      description = "example header transformation rule"
      enabled = false
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetTransformationRuleResponseHeaders(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_response_headers_transform"

    rules {
      action = "rewrite"
      action_parameters {
        headers {
          name      = "example1"
          operation = "set"
          value     = "my-http-header-value1"
        }

        headers {
          name       = "example2"
          operation  = "set"
          expression = "cf.zone.name"
        }

        headers {
          name      = "example3"
          operation = "remove"
        }
      }

      expression = "true"
      description = "example header transformation rule"
      enabled = false
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFPayloadLogging(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"
    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
        matched_data {
          public_key = "not_a_real_public_key"
        }
      }
      expression = "true"
      description = "example using WAF payload logging"
      enabled = false
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetCustomErrors(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_custom_errors"

    rules {
      action = "serve_error"
      action_parameters {
        content = "my example error page"
        content_type = "text/plain"
        status_code = "530"
      }
      expression = "(http.request.uri.path matches \"^/api/\")"
      description = "example http custom error response"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetOrigin(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_origin"

    rules {
      action = "route"
      action_parameters {
        host_header = "%[1]s.terraform.cfapi.net"
        origin {
          host = "%[1]s.terraform.cfapi.net"
          port = 80
        }
        sni {
          value = "%[1]s.terraform.cfapi.net"
        }
      }
      expression = "(http.request.uri.path matches \"^/api/\")"
      description = "example http request origin"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetRateLimit(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_ratelimit"

    rules {
      action = "block"
      action_parameters {
        response {
          status_code = 418
          content_type = "text/plain"
          content = "test content"
        }
      }
      ratelimit {
        characteristics = [
          "cf.colo.id",
          "ip.src"
        ]
        period = 60
        requests_per_period = 100
        mitigation_timeout = 60
        requests_to_origin = true
      }
      expression = "(http.request.uri.path matches \"^/api/\")"
      description = "example http rate limit"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetRateLimitScorePerPeriod(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_ratelimit"

    rules {
      action = "block"
      action_parameters {
        response {
          status_code = 418
          content_type = "text/plain"
          content = "test content"
        }
      }
      ratelimit {
        characteristics = [
          "cf.colo.id",
          "ip.src"
        ]
        period = 60
        score_per_period = 400
		score_response_header_name = "my-score"
        mitigation_timeout = 60
        requests_to_origin = true
      }
      expression = "(http.request.uri.path matches \"^/api/\")"
      description = "example http rate limit"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetTwoCustomRules(rnd, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id      = "%[2]s"
    name         = "Terraform provider test"
    description  = "%[1]s ruleset description"
    kind         = "zone"
    phase        = "http_request_firewall_custom"
    rules {
      action     = "log"
      enabled    = true
      expression = "(http.request.uri.path eq \"/admin\")"
    }
    rules {
      action     = "challenge"
      enabled    = true
      expression = "(http.request.uri.path eq \"/login\")"
    }
  }`, rnd, zoneID)
}

func testAccCheckCloudflareRulesetTwoCustomRulesReversed(rnd, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id      = "%[2]s"
    name         = "Terraform provider test"
    description  = "%[1]s ruleset description"
    kind         = "zone"
    phase        = "http_request_firewall_custom"
    rules {
      action     = "challenge"
      enabled    = true
      expression = "(http.request.uri.path eq \"/login\")"
    }
    rules {
      action     = "log"
      enabled    = true
      expression = "(http.request.uri.path eq \"/admin\")"
    }
  }`, rnd, zoneID)
}

func testAccCheckCloudflareRulesetThreeCustomRules(rnd, zoneID string, enableLoginRule bool) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id      = "%[2]s"
    name         = "Terraform provider test"
    description  = "%[1]s ruleset description"
    kind         = "zone"
    phase        = "http_request_firewall_custom"
    rules {
      action     = "log"
      enabled    = true
      expression = "(http.request.uri.path eq \"/admin\")"
    }
    rules {
      action     = "challenge"
      enabled    = %[3]t
      expression = "(http.request.uri.path eq \"/login\")"
    }
    rules {
      action     = "log"
      enabled    = true
      expression = "(http.request.uri.path eq \"/admin\")"
    }
  }`, rnd, zoneID, enableLoginRule)
}

func testAccCheckCloudflareRulesetActionParametersOverridesActionEnabled(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
        version = "latest"
        overrides {
          action = "log"
          status = "enabled"
        }
      }
      expression = "true"
      description = "Execute all rules in Cloudflare Managed Ruleset in log mode on my zone-level phase entry point ruleset"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetActionParametersMultipleSkips(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "skip"
      action_parameters {
        rulesets = ["efb7b8c949ac4650a09736fc376e9aee"]
      }
      expression = "(cf.zone.name eq \"domain.xyz\" and http.request.uri.query contains \"skip=rulesets\")"
      description = "skip Cloudflare Manage ruleset"
      enabled = true
	  logging {
		status = "enabled"
	  }
    }

    rules {
      action = "skip"
      action_parameters {
        # efb7b8c949ac4650a09736fc376e9aee is the ruleset ID of the Cloudflare Managed rules
        rules = {
          "efb7b8c949ac4650a09736fc376e9aee" = "5de7edfa648c4d6891dc3e7f84534ffa,e3a567afc347477d9702d9047e97d760"
        }
      }
      expression = "(cf.zone.name eq \"domain.xyz\" and http.request.uri.query contains \"skip=rules\")"
      description = "skip Wordpress rule and SQLi rule"
      enabled = true
	  logging {
		status = "enabled"
	  }
    }

    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
        version = "latest"
        overrides {
          rules {
            id = "5de7edfa648c4d6891dc3e7f84534ffa"
            action = "block"
            status = "enabled"
          }
          rules {
            id = "75a0060762034a6cb663fd51a02344cb"
            action = "log"
            status = "enabled"
          }
          categories {
            category = "wordpress"
            action = "js_challenge"
            status = "enabled"
          }
        }
      }
      expression = "true"
      description = "Execute Cloudflare Managed Ruleset on my zone-level phase entry point ruleset"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetActionParametersHTTPDDosOverride(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "ddos_l7"

    rules {
      action = "execute"
      action_parameters {
        id = "4d21379b4f9f4bb088e0729962c8b3cf"
        overrides {
          rules {
            id = "fdfdac75430c4c47a959592f0aa5e68a" # requests with odd HTTP headers or URI path
            sensitivity_level = "low"
          }
        }
      }
      expression = "true"
      description = "override HTTP DDoS ruleset rule"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetAccountLevelCustomWAFRule(rnd, name, accountID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s_account_custom_firewall" {
    account_id  = "%[3]s"
    name        = "Custom Ruleset for my account"
    description = "example block rule"
    kind        = "custom"
    phase       = "http_request_firewall_custom"

    rules {
      action = "block"
      expression = "(http.host eq \"%[4]s\")"
      description = "SID"
      enabled = true
    }
  }

  resource "cloudflare_ruleset" "%[1]s_account_custom_firewall_root" {
    account_id  = "%[3]s"
    name        = "Firewall Custom root"
    description = ""
    kind        = "root"
    phase       = "http_request_firewall_custom"

    rules {
      action = "execute"
      action_parameters {
        id = cloudflare_ruleset.%[1]s_account_custom_firewall.id
      }
      expression = "(cf.zone.name eq \"example.com\")"
      description = ""
      enabled = true
    }
  }`, rnd, name, accountID, zoneName)
}

func testAccCheckCloudflareRulesetTransformationRuleURIPathAndQueryCombination(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_transform"

    rules {
      action = "rewrite"
      action_parameters {
        uri {
          query {
            expression = "concat(\"requestUrl=\", http.request.full_uri)"
          }
          path {
            value = "/path/to/url"
          }
        }
      }
      expression = "true"
      description = "example for combining URI action parameters for path and query"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetExposedCredentialCheck(rnd, name, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    account_id  = "%[3]s"
    name        = "%[2]s"
    description = "This ruleset includes a rule checking for exposed credentials."
    kind        = "custom"
    phase       = "http_request_firewall_custom"

    rules {
      action = "log"
      expression = "http.request.method == \"POST\" && http.request.uri == \"/login.php\""
      enabled = true
      description = "example exposed credential check"
      exposed_credential_check {
        username_expression = "url_decode(http.request.body.form[\"username\"][0])"
        password_expression = "url_decode(http.request.body.form[\"password\"][0])"
      }
    }
  }
`, rnd, name, accountID)
}

func testAccCheckCloudflareRulesetDisableLoggingForSkipAction(rnd, name, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    account_id  = "%[3]s"
    name        = "%[2]s"
    description = "This ruleset includes a skip rule whose logging is disabled."
    kind        = "root"
    phase       = "http_request_firewall_managed"

    rules {
      action = "skip"
      action_parameters {
        ruleset = "current"
      }
      expression = "true"
      enabled = true
      description = "example disabled logging"
      logging {
        status = "disabled"
      }
    }
  }
`, rnd, name, accountID)
}

func testAccCloudflareRulesetConditionallySetActionParameterVersion_ExecuteAlone(rnd, accountID, domain string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    account_id  = "%[2]s"
    name        = "%[1]s managed WAF"
    description = "%[1]s managed WAF ruleset description"
    kind        = "root"
    phase       = "http_request_firewall_managed"

    rules {
      action = "execute"
      action_parameters {
        id = "4814384a9e5d4991b9815dcfc25d2f1f"
        overrides {
          rules {
            id = "6179ae15870a4bb7b2d480d4843b323c"
            action = "block"
            score_threshold = 25
          }
          status = "enabled"
        }
        matched_data {
           public_key = "zpUlcpNtaNiSUN6LL6NiNz8XgIJZWWG3iSZDdPbMszM="
        }
      }
      expression  = "(cf.zone.name eq \"%[3]s\")"
      description = "Account OWASP %[3]s"
      enabled     = true
    }
  }
`, rnd, accountID, domain)
}

func testAccCloudflareRulesetConditionallySetActionParameterVersion_ExecuteThenSkip(rnd, accountID, domain string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    account_id  = "%[2]s"
    name        = "%[1]s managed WAF"
    description = "%[1]s managed WAF ruleset description"
    kind        = "root"
    phase       = "http_request_firewall_managed"

    rules {
      action = "skip"
      action_parameters {
        rules = {
          "4814384a9e5d4991b9815dcfc25d2f1f" = "a6be45d4905042b9964ff81dc12e41d2,fa54f3d75ed446e78c22b4ea57b90acf,ec42fac3279943388b6be5ee9182835e,37da7855d2f94f69865365d894a556a4,f2db062052cf453fbe9e93f058ecf7e7,1129dfb383bb42e48466488cf3b37cb1"
        }
      }
      expression = "(cf.zone.name eq \"%[3]s\")"
      description = "Account skip rules OWASP"
      enabled = true
	  logging {
		status = "enabled"
	  }
    }

    rules {
      action = "execute"
      action_parameters {
        id = "4814384a9e5d4991b9815dcfc25d2f1f"
        overrides {
          rules {
            id = "6179ae15870a4bb7b2d480d4843b323c"
            action = "block"
            score_threshold = 25
          }
          status = "enabled"
        }
        matched_data {
           public_key = "zpUlcpNtaNiSUN6LL6NiNz8XgIJZWWG3iSZDdPbMszM="
        }
      }
      expression  = "(cf.zone.name eq \"%[3]s\")"
      description = "Account OWASP %[3]s"
      enabled     = true
    }
  }
`, rnd, accountID, domain)
}

func testAccCheckCloudflareRulesetManagedWAFWithCategoryBasedOverridesActionManagedChallenge(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
        overrides {
        	categories {
            	category = "wordpress"
            	action = "managed_challenge"
            	status = "enabled"
        	}
			rules {
				id = "e3a567afc347477d9702d9047e97d760"
				action = "managed_challenge"
				status = "enabled"
			}
        }
      }

      expression = "true"
      description = "overrides to only enable wordpress rules to managed_challenge"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFWithActionManagedChallenge(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
        overrides {
			action = "managed_challenge"
        }
      }

      expression = "true"
      description = "overrides change action to managed_challenge on the Cloudflare Manage Ruleset"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetLogCustomField(rnd, name, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_log_custom_fields"

    rules {
      action = "log_custom_field"
      action_parameters {
        request_fields = [
          "content-type",
          "x-forwarded-for",
          "host"
        ]
        response_fields = [
          "server",
          "content-type",
          "allow"
        ]
        cookie_fields = [
          "__ga",
          "accountNumber",
          "__cfruid"
        ]
      }

      expression = "true"
      description = "%[1]s log custom fields rule"
      enabled = true
    }
  }`, rnd, name, zoneID)
}

func testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName, status string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id     = "%[2]s"
    name        = "thrashing overrides for managed rules"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_firewall_managed"

    rules {
      action = "execute"
      action_parameters {
        id = "efb7b8c949ac4650a09736fc376e9aee"
        version = "latest"
        overrides {
          action = "log"
          %[4]s
        }
      }
      expression = "true"
      description = "Execute all rules in Cloudflare Managed Ruleset in log mode on my zone-level phase entry point ruleset"
      enabled = true
    }
  }`, rnd, zoneID, zoneName, status)
}

func testAccCloudflareRulesetCacheSettingsAllEnabled(rnd, accountID, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
	zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_cache_settings"

    rules {
      action = "set_cache_settings"
      action_parameters {
		edge_ttl {
			mode = "override_origin"
			default = 60
			status_code_ttl {
				status_code = 200
				value = 50
			}
			status_code_ttl {
				status_code_range {
					from = 201
					to = 300
				}
				value = 30
			}
		}
		browser_ttl {
			mode = "respect_origin"
		}
		serve_stale {
			disable_stale_while_updating = true
		}
		respect_strong_etags = true
		cache_key {
			ignore_query_strings_order = false
			cache_deception_armor = true
			custom_key {
				query_string {
					exclude = ["*"]
				}
				header {
					include = ["habc", "hdef"]
					check_presence = ["habc_t", "hdef_t"]
					exclude_origin = true
				}
				cookie {
					include = ["cabc", "cdef"]
					check_presence = ["cabc_t", "cdef_t"]
				}
				user {
					device_type = true
					geo = false
				}
				host {
					resolved = true
				}
			}
		}
		origin_error_page_passthru = false
      }
	  expression = "true"
	  description = "%[1]s set cache settings rule"
	  enabled = true
    }
  }`, rnd, accountID, zoneID)
}

func testAccCloudflareRulesetCacheSettingsOptionalsEmpty(rnd, accountID, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
	zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_cache_settings"

    rules {
      action = "set_cache_settings"
      action_parameters {
		edge_ttl {
			mode = "override_origin"
			default = 60
		}
		browser_ttl {
			mode = "respect_origin"
		}
		serve_stale {}
		cache_key {}
      }
	  expression = "true"
	  description = "%[1]s set cache settings rule"
	  enabled = true
    }
  }`, rnd, accountID, zoneID)
}

func testAccCloudflareRulesetCacheSettingsCustomKeyEmpty(rnd, accountID, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
	zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_cache_settings"

    rules {
      action = "set_cache_settings"
      action_parameters {
		cache_key {
			custom_key {
				query_string {}
				header {}
				cookie {}
				user {}
				host {}
			}
		}
      }
	  expression = "true"
	  description = "%[1]s set cache settings rule"
	  enabled = true
    }
  }`, rnd, accountID, zoneID)
}

func testAccCloudflareRulesetCacheSettingsEdgeTTLRespectOrigin(rnd, zoneID string) string {
	return fmt.Sprintf(`
	resource "cloudflare_ruleset" "%[1]s" {
		zone_id     = "%[2]s"
		name        = "%[1]s"
		description = "%[1]s ruleset description"
		kind        = "zone"
		phase       = "http_request_cache_settings"

		rules {
			action = "set_cache_settings"
			action_parameters {
				edge_ttl {
					status_code_ttl {
						status_code_range {
							from = 201
							to = 300
						}
						value = 5
					}
					mode = "respect_origin"
				}
				cache = true
			}
			expression = "true"
			description = "%[1]s set cache settings rule"
			enabled = true
		}
	}`, rnd, zoneID)
}

func testAccCloudflareRulesetConfigAllEnabled(rnd, accountID, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
	zone_id     = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_config_settings"

    rules {
      action = "set_config"
      action_parameters {
		automatic_https_rewrites = true
		autominify {
			html = true
			css = true
			js = true
		}
		bic = true
		disable_apps = true
		disable_zaraz = true
		disable_railgun = true
		email_obfuscation = true
		mirage = true
		opportunistic_encryption = true
		polish = "off"
		rocket_loader = true
		security_level = "off"
		server_side_excludes = true
		ssl = "off"
		sxg = true
		hotlink_protection = true
      }
	  expression = "true"
	  description = "%[1]s set config rule"
	  enabled = true
    }
  }`, rnd, accountID, zoneID)
}

func testAccCloudflareRulesetRedirectFromList(rnd, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "list-%[1]s" {
    account_id = "%[2]s"
    name = "redirect_list_%[1]s"
    description = "%[1]s list description"
    kind = "redirect"
  }

  resource "cloudflare_ruleset" "%[1]s" {
	account_id  = "%[2]s"
    name        = "%[1]s"
    description = "%[1]s ruleset description"
    kind        = "root"
    phase       = "http_request_redirect"

    rules {
      action = "redirect"
      action_parameters {
        from_list {
      	  name = cloudflare_list.list-%[1]s.name
      	  key = "http.request.full_uri"
        }
      }
      expression = "http.request.full_uri in $redirect_list_%[1]s"
      description = "Apply redirects from redirect list"
      enabled = true
    }
  }`, rnd, accountID)
}

func testAccCloudflareRulesetRedirectFromValue(rnd, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
	zone_id     = "%[2]s"
    name        = "%[1]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_dynamic_redirect"

    rules {
      action = "redirect"
      action_parameters {
        from_value {
		  status_code = 301
		  target_url {
			value = "some_host.com"
		  }
		  preserve_query_string = true
        }
      }
      expression = "true"
      description = "Apply redirect from value"
      enabled = true
    }
  }`, rnd, zoneID)
}

func testAccCheckCloudflareRulesetActionParametersOverrideSensitivityForAllRulesetRules(rnd, name, zoneID, zoneName string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "ddos_l7"

    rules {
      action = "execute"
      action_parameters {
        id = "4d21379b4f9f4bb088e0729962c8b3cf"
        overrides {
		  action            = "log"
		  sensitivity_level = "low"
        }
      }
      expression = "true"
      description = "override HTTP DDoS ruleset rule"
      enabled = true
    }
  }`, rnd, name, zoneID, zoneName)
}
