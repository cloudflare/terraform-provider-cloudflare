package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudflareRuleset_WAFBasic(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetCustomWAFBasic(rnd, "my basic WAF ruleset", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my basic WAF ruleset"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "custom"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_custom"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "log"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" ruleset rule description"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRuleset(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.0.enabled", "true"),

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
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.0.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.1.category", "paranoia-level-3"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.1.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.2.category", "paranoia-level-4"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.categories.2.enabled", "false"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.id", "6179ae15870a4bb7b2d480d4843b323c"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.action", "block"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.score_threshold", "60"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.0.rules.0.enabled", "true"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "zone"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetDeployMultiple(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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

func TestAccCloudflareRuleset_WAFManagedRulesetWithCategoryBasedOverrides(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetWithIDBasedOverrides(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.0.overrides.#", "1"),
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
	skipMagicTransitTestForNonConfiguredDefaultZone(t)

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_ruleset.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckAccount(t) },
		Providers: testAccProviders,
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
					resource.TestCheckResourceAttr(
						name, "name", rnd),
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
      enabled = true
    }

    rules {
      action = "allow"
      expression = "tcp.dstport in { 32768..65535 }"
      description = "Allow TCP Ephemeral Ports"
      enabled = true
    }
  }`, rnd, name, accountID)
}

func testAccCheckCloudflareRulesetCustomWAFBasic(rnd, name, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_ruleset" "%[1]s" {
    account_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "custom"
    phase       = "http_request_firewall_custom"

    rules {
      action = "log"
      expression = "true"
      description = "%[1]s ruleset rule description"
    }
  }`, rnd, name, accountID)
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
            enabled = true
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
            enabled = false
          }

          categories {
            category = "paranoia-level-3"
            enabled = false
          }

          categories {
            category = "paranoia-level-4"
            enabled = false
          }

          rules {
            id = "6179ae15870a4bb7b2d480d4843b323c"
            action = "block"
            score_threshold = 60
            enabled = true
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
            enabled = true
          }

          categories {
            category = "joomla"
            action = "block"
            enabled = true
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
            enabled = true
          }

          rules {
            id = "e3a567afc347477d9702d9047e97d760"
            action = "log"
            enabled = true
          }
        }
      }

      expression = "true"
      description = "make 5de7edfa648c4d6891dc3e7f84534ffa and e3a567afc347477d9702d9047e97d760 log only"
      enabled = false
    }
  }`, rnd, name, zoneID, zoneName)
}
