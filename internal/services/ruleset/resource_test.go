package ruleset_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_ruleset", &resource.Sweeper{
		Name: "cloudflare_ruleset",
		F: func(region string) error {
			client, err := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

			if err != nil {
				return fmt.Errorf("error establishing client: %w", err)
			}

			ctx := context.Background()
			accountRulesets, err := client.ListRulesets(ctx, cfv1.AccountIdentifier(accountID), cfv1.ListRulesetsParams{})
			if err != nil {
				return fmt.Errorf("failed to fetch rulesets: %w", err)
			}

			for _, ruleset := range accountRulesets {
				if ruleset.Kind != "managed" {
					err := client.DeleteRuleset(ctx, cfv1.AccountIdentifier(accountID), ruleset.ID)
					if err != nil {
						return fmt.Errorf("failed to delete ruleset %q: %w", ruleset.ID, err)
					}
				}
			}

			zoneRulesets, _ := client.ListRulesets(ctx, cfv1.ZoneIdentifier(zoneID), cfv1.ListRulesetsParams{})
			for _, ruleset := range zoneRulesets {
				if ruleset.Kind != "managed" {
					err := client.DeleteRuleset(ctx, cfv1.ZoneIdentifier(zoneID), ruleset.ID)
					if err != nil {
						return fmt.Errorf("failed to delete ruleset %q: %w", ruleset.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccCloudflareRuleset_WAFBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
			// {
			// 	ResourceName:        resourceName,
			// 	ImportStateIdPrefix: fmt.Sprintf("zone/%s/", zoneID),
			// 	ImportState:         true,
			// 	ImportStateVerify:   true,
			// },
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "Execute Cloudflare Managed Ruleset on my zone-level phase ruleset"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_WAFManagedRulesetWithoutDescription(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetManagedWAFWithoutDescription(rnd, "my basic managed WAF ruleset", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my basic managed WAF ruleset"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "efb7b8c949ac4650a09736fc376e9aee"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "4814384a9e5d4991b9815dcfc25d2f1f"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.0.category", "xss"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.0.action", "block"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.0.enabled", "true"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.id", "4814384a9e5d4991b9815dcfc25d2f1f"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.overrides.rules.0.id", "6179ae15870a4bb7b2d480d4843b323c"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.overrides.rules.0.action", "block"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.overrides.rules.0.score_threshold", "60"),

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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "4814384a9e5d4991b9815dcfc25d2f1f"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.0.category", "paranoia-level-2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.0.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.1.category", "paranoia-level-3"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.1.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.2.category", "paranoia-level-4"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.2.enabled", "false"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.id", "6179ae15870a4bb7b2d480d4843b323c"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.action", "block"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.score_threshold", "60"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.enabled", "true"),

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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "4814384a9e5d4991b9815dcfc25d2f1f"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.2.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.action_parameters.id", "c2e184081120413c86c3ab7e14069605"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.ruleset", "current"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", fmt.Sprintf(`(http.host eq "%s" and http.request.method eq "GET")`, zoneName)),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "not this zone"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.id", "4814384a9e5d4991b9815dcfc25d2f1f"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.2.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.action_parameters.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.3.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.action_parameters.id", "c2e184081120413c86c3ab7e14069605"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.ruleset", "current"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", fmt.Sprintf(`(http.host eq "%s" and http.request.uri.path contains "/app/")`, zoneName)),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "not this path"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.id", "4814384a9e5d4991b9815dcfc25d2f1f"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.2.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.action_parameters.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.3.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.action_parameters.id", "c2e184081120413c86c3ab7e14069605"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.3.description", "zone deployment test"),

					resource.TestCheckResourceAttr(resourceName, "rules.4.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.4.action_parameters.ruleset", "current"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.ruleset", "current"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", fmt.Sprintf(`http.host eq "%s"`, zoneName)),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "not this zone"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.phases.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "rules.1.action_parameters.phases.*", "http_ratelimit"),
					resource.TestCheckTypeSetElemAttr(resourceName, "rules.1.action_parameters.phases.*", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.2.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.action_parameters.products.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "rules.2.action_parameters.products.*", "zoneLockdown"),
					resource.TestCheckTypeSetElemAttr(resourceName, "rules.2.action_parameters.products.*", "uaBlock"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.0.category", "wordpress"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.0.action", "block"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.1.category", "joomla"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.1.action", "block"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.id", "e3a567afc347477d9702d9047e97d760"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.enabled", "false"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "efb7b8c949ac4650a09736fc376e9aee"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.id", "5de7edfa648c4d6891dc3e7f84534ffa"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.action", "log"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.1.id", "e3a567afc347477d9702d9047e97d760"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.1.action", "log"),
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

	acctest.TestAccSkipForDefaultAccount(t, "Default account is not configured for Magic Transit.")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_ruleset.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetMagicTransitSingle(rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", fmt.Sprintf("%s magic transit ruleset description", rnd)),
					resource.TestCheckResourceAttr(name, "rules.#", "1"),
					resource.TestCheckResourceAttr(name, "rules.0.action", "skip"),
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
					resource.TestCheckResourceAttr(name, "rules.1.action", "skip"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.matched_data.public_key", "bm90X2FfcmVhbF9wdWJsaWNfa2V5"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response.0.status_code", "418"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response.0.content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response.0.content", "test content"),
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
			// {
			// 	ResourceName:        resourceName,
			// 	ImportStateIdPrefix: fmt.Sprintf("zone/%s/", zoneID),
			// 	ImportState:         true,
			// 	ImportStateVerify:   true,
			// },
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

	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response.0.status_code", "418"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response.0.content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response.0.content", "test content"),
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

func TestAccCloudflareRuleset_RateLimitMitigationTimeoutOfZero(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetRateLimitWithMitigationTimeoutOfZero(rnd, "example HTTP rate limit", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example HTTP rate limit"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_ratelimit"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "block"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response.0.status_code", "418"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response.0.content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response.0.content", "test content"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "(http.request.uri.path matches \"^/api/\")"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example http rate limit"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.#", "1"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.characteristics.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.period", "60"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.requests_per_period", "1000"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.mitigation_timeout", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ratelimit.0.requests_to_origin", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_PreserveRuleRefs(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	var adminRuleRef, loginRuleRef, adminRuleCopyRef, adminRuleExplicitRef string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create a ruleset with two rules (one for /admin, one for
				// /login) and get their refs.
				Config: testAccCheckCloudflareRulesetTwoCustomRules(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", getValue(&adminRuleRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", getValue(&loginRuleRef)),
				),
			},
			{
				// Reverse the order of rules. The refs should remain the same,
				// just in reverse order.
				Config: testAccCheckCloudflareRulesetTwoCustomRulesReversed(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", equalsValue(&loginRuleRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", equalsValue(&adminRuleRef)),
				),
			},
			{
				// Revert to the original version.
				Config: testAccCheckCloudflareRulesetTwoCustomRules(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", equalsValue(&adminRuleRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", equalsValue(&loginRuleRef)),
				),
			},
			{
				// Append a copy of the admin rule. The first two refs should
				// not change.
				Config: testAccCheckCloudflareRulesetThreeCustomRules(rnd, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", equalsValue(&adminRuleRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", equalsValue(&loginRuleRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.2.ref", notEqualsValue(&adminRuleRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.2.ref", getValue(&adminRuleCopyRef)),
				),
			},
			{
				// Disable the login rule. Its ref will change, but the admin
				// rule refs should remain the same.
				Config: testAccCheckCloudflareRulesetThreeCustomRules(rnd, zoneID, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", equalsValue(&adminRuleRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", notEqualsValue(&loginRuleRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.2.ref", equalsValue(&adminRuleCopyRef)),
				),
			},
			{
				// Revert to the original version. The preserved admin rule ref
				// should stay the same, and the login rule ref should change.
				Config: testAccCheckCloudflareRulesetTwoCustomRules(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", equalsValue(&adminRuleRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", notEqualsValue(&loginRuleRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", getValue(&loginRuleRef)),
				),
			},
			{
				// Give the admin rule a ref.
				Config: testAccCheckCloudflareRulesetTwoCustomRulesWithRef(rnd, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.0.ref", "foo"),
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", getValue(&adminRuleExplicitRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", equalsValue(&loginRuleRef)),
				),
			},
			{
				// Disable the admin rule. Its ref should stay the same.
				Config: testAccCheckCloudflareRulesetTwoCustomRulesWithRef(rnd, zoneID, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", equalsValue(&adminRuleExplicitRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", equalsValue(&loginRuleRef)),
				),
			},
			{
				// Prepend a copy of the admin rule without an explicit ref. The
				// original rule should keep its explicit ref and the new rule
				// should get a new ref.
				Config: testAccCheckCloudflareRulesetThreeCustomRulesWithRef(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", notEqualsValue(&adminRuleRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", notEqualsValue(&adminRuleCopyRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", notEqualsValue(&adminRuleExplicitRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", equalsValue(&adminRuleExplicitRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.2.ref", equalsValue(&loginRuleRef)),
				),
			},
			{
				// Remove the prepended admin rule and re-enable the original
				// admin rule.
				Config: testAccCheckCloudflareRulesetTwoCustomRulesWithRef(rnd, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", equalsValue(&adminRuleExplicitRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", equalsValue(&loginRuleRef)),
				),
			},
			{
				// Revert to the original version. The refs should remain
				// exactly the same.
				Config: testAccCheckCloudflareRulesetTwoCustomRules(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", equalsValue(&adminRuleExplicitRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", equalsValue(&loginRuleRef)),
				),
			},
			{
				// Reverse the order of rules. The refs should remain the same,
				// just in reverse order.
				Config: testAccCheckCloudflareRulesetTwoCustomRulesReversed(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "rules.0.ref", equalsValue(&loginRuleRef)),
					resource.TestCheckResourceAttrWith(resourceName, "rules.1.ref", equalsValue(&adminRuleExplicitRef)),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.content", "my example error page"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.status_code", "530"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.host_header", rnd+"."+zoneName),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.origin.host", rnd+"."+zoneName),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.origin.port", "80"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.sni.value", rnd+"."+zoneName),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "(http.request.uri.path matches \"^/api/\")"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example http request origin"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_RequestOriginPortWithoutHost(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetOriginPortWithoutOrigin(rnd, "example HTTP request origin", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example HTTP request origin"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_origin"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "route"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.host_header", rnd+"."+zoneName),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.origin.port", "80"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.uri.path.value", "/static-rewrite"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.uri.query.value", "a=b"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.uri.path.value", "/path/to/url"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.uri.query.expression", "concat(\"requestUrl=\", http.request.full_uri)"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.headers.example1.value", "my-http-header-value1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.headers.example1.operation", "set"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.headers.example2.operation", "set"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.headers.example2.expression", "cf.zone.name"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.headers.example3.operation", "remove"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.headers.example1.value", "my-http-header-value1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.headers.example1.operation", "set"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.headers.example2.operation", "set"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.headers.example2.expression", "cf.zone.name"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.headers.example3.operation", "remove"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_ResponseCompression(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetResponseCompression(rnd, "my basic response compression ruleset", zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my basic response compression ruleset"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_response_compression"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "compress_response"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" compress response rule"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.algorithms.0.name", "brotli"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.algorithms.1.name", "default"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.algorithms.#", "2"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.rulesets.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "(cf.zone.name eq \"domain.xyz\" and http.request.uri.query contains \"skip=rulesets\")"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "skip Cloudflare Manage ruleset"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.enabled", "true"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.rules.efb7b8c949ac4650a09736fc376e9aee.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.rules.efb7b8c949ac4650a09736fc376e9aee.0", "5de7edfa648c4d6891dc3e7f84534ffa"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.action_parameters.rules.efb7b8c949ac4650a09736fc376e9aee.1", "e3a567afc347477d9702d9047e97d760"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckNoResourceAttr(resourceName, "rules.0.action_parameters.enabled"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.action", "log"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.enabled", "true"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "4d21379b4f9f4bb088e0729962c8b3cf"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.id", "d59a914a1e494067b613534f1fc1e601"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.sensitivity_level", "low"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "4d21379b4f9f4bb088e0729962c8b3cf"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.action", "log"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.sensitivity_level", "low"),
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

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
	acctest.TestAccSkipForDefaultZone(t, "Pending API documentation for username_expression and password_expression")
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the WAF
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.exposed_credential_check.username_expression", "url_decode(http.request.body.form[\"username\"][0])"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.exposed_credential_check.password_expression", "url_decode(http.request.body.form[\"password\"][0])"),
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

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "(cf.zone.plan eq \"ENT\")"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example disabled logging"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.logging.enabled", "false"),
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

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.0.category", "wordpress"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.0.action", "managed_challenge"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.categories.0.enabled", "true"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.id", "e3a567afc347477d9702d9047e97d760"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.rules.0.action", "managed_challenge"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetManagedWAFWithActionManagedChallenge(rnd, "my managed WAF ruleset with overrides", zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my managed WAF ruleset with overrides"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_firewall_managed"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "execute"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.id", "efb7b8c949ac4650a09736fc376e9aee"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.action", "managed_challenge"),
					resource.TestCheckNoResourceAttr(resourceName, "rules.0.action_parameters.overrides.enabled"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "overrides change action to managed_challenge on the Cloudflare Manage Ruleset"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_LogCustomField(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.request_fields.0.name", "content-type"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.request_fields.1.name", "x-forwarded-for"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.request_fields.2.name", "host"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response_fields.0.name", "server"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response_fields.1.name", "content-type"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.response_fields.2.name", "allow"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cookie_fields.0.name", "__ga"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cookie_fields.1.name", "accountNumber"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cookie_fields.2.name", "__cfruid"),
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

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatusWithoutEnabled(rnd, zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr(resourceName, "rules.0.action_parameters.overrides.enabled"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.enabled", "false"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatusWithoutEnabled(rnd, zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr(resourceName, "rules.0.action_parameters.overrides.enabled"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.enabled", "true"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.enabled", "false"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.overrides.enabled", "true"),
				),
			},
			{
				Config: testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatusWithoutEnabled(rnd, zoneID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr(resourceName, "rules.0.action_parameters.overrides.enabled"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsAllEnabled(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.additional_cacheable_ports.0", "8443"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.mode", "override_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.default", "60"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.0.value", "50"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.0.status_code", "200"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.1.value", "30"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.1.status_code_range.from", "201"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.1.status_code_range.to", "300"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.browser_ttl.mode", "respect_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.serve_stale.disable_stale_while_updating", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.read_timeout", "2000"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.respect_strong_etags", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.ignore_query_strings_order", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.cache_deception_armor", "true"),
					// resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.query_string.exclude.all", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.include.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.include.0", "habc"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.include.1", "hdef"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.check_presence.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.check_presence.0", "habc_t"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.check_presence.1", "hdef_t"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.exclude_origin", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.cookie.include.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.cookie.include.0", "cabc"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.cookie.include.1", "cdef"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.cookie.check_presence.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.cookie.check_presence.0", "cabc_t"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.cookie.check_presence.1", "cdef_t"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.user.device_type", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.user.geo", "false"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.host.resolved", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.origin_cache_control", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.origin_error_page_passthru", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsOptionalsEmpty(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.mode", "override_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.default", "60"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.browser_ttl.mode", "respect_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.0.custom_key.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.serve_stale.#", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsOnlyExludeOrigin(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsOnlyExludeOrigin(rnd, "my basic cache settings ruleset", zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "my basic cache settings ruleset"),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" set cache settings rule"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.include.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.check_presence.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.exclude_origin", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsEdgeTTLRespectOrigin(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.0.value", "5"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.mode", "respect_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsNoCacheForStatus(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsNoCacheForStatus(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" set cache settings rule"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.0.status_code_range.from", "400"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.0.status_code_range.to", "500"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.0.value", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsStatusRangeGreaterThan(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsStatusRangeGreaterThan(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" set cache settings rule"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.0.value", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.0.status_code_range.from", "105"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.1.value", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.1.status_code_range.from", "100"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.1.status_code_range.to", "101"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsStatusRangeLessThan(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsStatusRangeLessThan(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" set cache settings rule"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.0.value", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.0.status_code_range.to", "400"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.1.value", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.1.status_code_range.from", "500"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.status_code_ttl.1.status_code_range.to", "501"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsFalse(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsFalse(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "description", rnd+" ruleset description"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", rnd+" set cache settings rule"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_Config(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.automatic_https_rewrites", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.autominify.html", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.autominify.css", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.autominify.js", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.bic", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.disable_apps", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.disable_zaraz", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.email_obfuscation", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.mirage", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.opportunistic_encryption", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.polish", "off"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.rocket_loader", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.security_level", "off"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.server_side_excludes", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.ssl", "off"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.sxg", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.hotlink_protection", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.disable_rum", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.fonts", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_Redirect(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountId := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetRedirectFromList(rnd, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "kind", "root"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_redirect"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "redirect"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.from_list.name", "redirect_list_"+rnd),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.from_list.key", "http.request.full_uri"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_DynamicRedirect(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetRedirectFromValue(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_dynamic_redirect"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "redirect"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.from_value.status_code", "301"),
					resource.TestCheckNoResourceAttr(resourceName, "rules.0.action_parameters.from_value.target_url.expression"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.from_value.target_url.value", "some_host.com"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.from_value.preserve_query_string", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_DynamicRedirectWithoutPreservingQueryString(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetRedirectFromValueWithoutPreservingQueryString(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_dynamic_redirect"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "redirect"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.from_value.status_code", "301"),
					resource.TestCheckNoResourceAttr(resourceName, "rules.0.action_parameters.from_value.target_url.expression"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.from_value.target_url.value", "some_host.com"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_TransformationRuleURIStripOffQueryString(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetRewriteForEmptyQueryString(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_transform"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "rewrite"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.uri.query.value", ""),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_TransformationRuleURIStripOffPath(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetRewriteForEmptyPath(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_transform"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "rewrite"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.uri.path.value", "/"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_ConfigSingleFalseyValue(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetConfigSingleFalseyValue(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_config_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_config"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.bic", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsMissingEdgeTTLWithOverrideOrigin(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending porting schema validation")

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareRulesetCacheSettingsMissingDefaultEdgeTTLOverrideOrigin(rnd, zoneID),
				ExpectError: regexp.MustCompile("using mode 'override_origin' requires setting a default for ttl"),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsMissingBrowserTTLWithOverrideOrigin(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending porting schema validation")

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareRulesetCacheSettingsMissingDefaultBrowserTTLOverrideOrigin(rnd, zoneID),
				ExpectError: regexp.MustCompile("using mode 'override_origin' requires setting a default for ttl"),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsInvalidEdgeTTLWithOverrideOrigin(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending porting schema validation")

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareRulesetCacheSettingsInvalidDefaultEdgeTTLOverrideOrigin(rnd, zoneID),
				ExpectError: regexp.MustCompile("using mode 'override_origin' requires setting a default for ttl"),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsEdgeTTLWithBypassByDefault(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsBypassByDefaultEdge(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "description", "set cache settings for the request"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.mode", "bypass_by_default"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsInvalidEdgeTTLWithBypassByDefault(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending porting schema validation")

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareRulesetCacheSettingsBypassByDefaultEdgeInvalid(rnd, zoneID),
				ExpectError: regexp.MustCompile("cannot set default ttl when using mode 'bypass_by_default'"),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsBrowserTTLWithBypass(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsBypassBrowser(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "description", "set cache settings for the request"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.browser_ttl.mode", "bypass_by_default"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsInvalidBrowserTTLWithBypass(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending porting schema validation")

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareRulesetCacheSettingsBypassBrowserInvalid(rnd, zoneID),
				ExpectError: regexp.MustCompile("cannot set default ttl when using mode 'bypass'"),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsInvalidBrowserTTLWithOverrideOrigin(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending porting schema validation")

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareRulesetCacheSettingsInvalidDefaultBrowserTTLOverrideOrigin(rnd, zoneID),
				ExpectError: regexp.MustCompile("using mode 'override_origin' requires setting a default for ttl"),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsDefinedQueryStringExcludeKeys(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending updating service to match API documentation")

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsExplicitCustomKeyCacheKeysQueryStringsExclude(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "set cache settings for the request"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.mode", "override_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.default", "7200"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.query_string.exclude.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.query_string.exclude.0", "example"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsDefinedQueryStringIncludeKeys(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending updating service to match API documentation")

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsExplicitCustomKeyCacheKeysQueryStringsInclude(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "set cache settings for the request"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.mode", "override_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.default", "7200"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.query_string.include.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.query_string.include.0", "another_example"),
				),
			},
		},
	})
}

func TestAccCloudflareRuleset_CacheSettingsHandleDefaultHeaderExcludeOrigin(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_ruleset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRulesetCacheSettingsHandleDefaultHeaderExcludeOrigin(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "set cache settings for the request"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.mode", "override_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.default", "7200"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.check_presence.0", "x-forwarded-for"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.include.0", "x-test"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.include.1", "x-test2"),
					// resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.exclude_origin", "false"),
				),
			},
			{
				Config: testAccCloudflareRulesetCacheSettingsHandleHeaderExcludeOriginFalse(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "set cache settings for the request"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.mode", "override_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.default", "7200"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.check_presence.0", "x-forwarded-for"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.include.0", "x-test"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.include.1", "x-test2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.exclude_origin", "false"),
				),
			},
			{
				Config: testAccCloudflareRulesetCacheSettingsHandleHeaderExcludeOriginSet(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "set cache settings for the request"),
					resource.TestCheckResourceAttr(resourceName, "kind", "zone"),
					resource.TestCheckResourceAttr(resourceName, "phase", "http_request_cache_settings"),

					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", "set_cache_settings"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "example"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.mode", "override_origin"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.edge_ttl.default", "7200"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.check_presence.0", "x-forwarded-for"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.include.0", "x-test"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.include.1", "x-test2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action_parameters.cache_key.custom_key.header.exclude_origin", "true"),
				),
			},
		},
	})
}

func testAccCheckCloudflareRulesetMagicTransitSingle(rnd, name, accountID string) string {
	return acctest.LoadTestCase("rulesetmagictransitsingle.tf", rnd, name, accountID)
}

func testAccCheckCloudflareRulesetMagicTransitMultiple(rnd, name, accountID string) string {
	return acctest.LoadTestCase("rulesetmagictransitmultiple.tf", rnd, name, accountID)
}

func testAccCheckCloudflareRulesetCustomWAFBasic(rnd, name, zoneID string) string {
	return acctest.LoadTestCase("rulesetcustomwafbasic.tf", rnd, name, zoneID)
}

func testAccCheckCloudflareRulesetManagedWAF(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwaf.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFWithoutDescription(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwafwithoutdescription.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFOWASP(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwafowasp.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFOWASPBlockXSSAndAnomalyOver60(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwafowaspblockxssandanomalyover60.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFOWASPOnlyPL1(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwafowasponlypl1.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFDeployMultiple(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwafdeploymultiple.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFDeployMultipleWithSkip(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwafdeploymultiplewithskip.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFDeployMultipleWithTopSkipAndLastSkip(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwafdeploymultiplewithtopskipandlastskip.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetSkipPhaseAndProducts(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetskipphaseandproducts.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFWithCategoryBasedOverrides(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwafwithcategorybasedoverrides.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFWithIDBasedOverrides(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwafwithidbasedoverrides.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetTransformationRuleURIPath(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesettransformationruleuripath.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetTransformationRuleURIQuery(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesettransformationruleuriquery.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetTransformationRuleRequestHeaders(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesettransformationrulerequestheaders.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetTransformationRuleResponseHeaders(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesettransformationruleresponseheaders.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetResponseCompression(rnd, name, zoneID string) string {
	return acctest.LoadTestCase("rulesetresponsecompression.tf", rnd, name, zoneID)
}

func testAccCheckCloudflareRulesetManagedWAFPayloadLogging(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwafpayloadlogging.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetCustomErrors(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetcustomerrors.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetOrigin(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetorigin.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetOriginPortWithoutOrigin(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetoriginportwithoutorigin.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetRateLimit(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetratelimit.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetRateLimitScorePerPeriod(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetratelimitscoreperperiod.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetRateLimitWithMitigationTimeoutOfZero(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetratelimitwithmitigationtimeoutofzero.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetTwoCustomRules(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesettwocustomrules.tf", rnd, zoneID)
}

func testAccCheckCloudflareRulesetTwoCustomRulesReversed(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesettwocustomrulesreversed.tf", rnd, zoneID)
}

func testAccCheckCloudflareRulesetThreeCustomRules(rnd, zoneID string, enableLoginRule bool) string {
	return acctest.LoadTestCase("rulesetthreecustomrules.tf", rnd, zoneID, enableLoginRule)
}

func testAccCheckCloudflareRulesetTwoCustomRulesWithRef(rnd, zoneID string, enableAdminRule bool) string {
	return acctest.LoadTestCase("rulesettwocustomruleswithref.tf", rnd, zoneID, enableAdminRule)
}

func testAccCheckCloudflareRulesetThreeCustomRulesWithRef(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetthreecustomruleswithref.tf", rnd, zoneID)
}

func testAccCheckCloudflareRulesetActionParametersOverridesActionEnabled(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetactionparametersoverridesactionenabled.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetActionParametersMultipleSkips(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetactionparametersmultipleskips.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetActionParametersHTTPDDosOverride(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetactionparametershttpddosoverride.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetAccountLevelCustomWAFRule(rnd, name, accountID, zoneName string) string {
	return acctest.LoadTestCase("rulesetaccountlevelcustomwafrule.tf", rnd, name, accountID, zoneName)
}

func testAccCheckCloudflareRulesetTransformationRuleURIPathAndQueryCombination(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesettransformationruleuripathandquerycombination.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetExposedCredentialCheck(rnd, name, accountID string) string {
	return acctest.LoadTestCase("rulesetexposedcredentialcheck.tf", rnd, name, accountID)
}

func testAccCheckCloudflareRulesetDisableLoggingForSkipAction(rnd, name, accountID string) string {
	return acctest.LoadTestCase("rulesetdisableloggingforskipaction.tf", rnd, name, accountID)
}

func testAccCloudflareRulesetConditionallySetActionParameterVersion_ExecuteAlone(rnd, accountID, domain string) string {
	return acctest.LoadTestCase("executealone.tf", rnd, accountID, domain)
}

func testAccCloudflareRulesetConditionallySetActionParameterVersion_ExecuteThenSkip(rnd, accountID, domain string) string {
	return acctest.LoadTestCase("executethenskip.tf", rnd, accountID, domain)
}

func testAccCheckCloudflareRulesetManagedWAFWithCategoryBasedOverridesActionManagedChallenge(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwafwithcategorybasedoverridesactionmanagedchallenge.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetManagedWAFWithActionManagedChallenge(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetmanagedwafwithactionmanagedchallenge.tf", rnd, name, zoneID, zoneName)
}

func testAccCheckCloudflareRulesetLogCustomField(rnd, name, zoneID string) string {
	return acctest.LoadTestCase("rulesetlogcustomfield.tf", rnd, name, zoneID)
}

func testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatus(rnd, zoneID, zoneName string, status bool) string {
	return acctest.LoadTestCase("rulesetactionparametersoverridesthrashingstatus.tf", rnd, zoneID, zoneName, status)
}

func testAccCheckCloudflareRulesetActionParametersOverridesThrashingStatusWithoutEnabled(rnd, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetactionparametersoverridesthrashingstatuswithoutenabled.tf", rnd, zoneID, zoneName)
}

func testAccCloudflareRulesetCacheSettingsAllEnabled(rnd, accountID, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsallenabled.tf", rnd, accountID, zoneID)
}

func testAccCloudflareRulesetCacheSettingsOptionalsEmpty(rnd, accountID, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsoptionalsempty.tf", rnd, accountID, zoneID)
}

func testAccCloudflareRulesetCacheSettingsOnlyExludeOrigin(rnd, accountID, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsonlyexludeorigin.tf", rnd, accountID, zoneID)
}

func testAccCloudflareRulesetCacheSettingsMissingDefaultEdgeTTLOverrideOrigin(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsmissingdefaultedgettloverrideorigin.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsMissingDefaultBrowserTTLOverrideOrigin(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsmissingdefaultbrowserttloverrideorigin.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsInvalidDefaultEdgeTTLOverrideOrigin(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsinvaliddefaultedgettloverrideorigin.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsInvalidDefaultBrowserTTLOverrideOrigin(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsinvaliddefaultbrowserttloverrideorigin.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsEdgeTTLRespectOrigin(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsedgettlrespectorigin.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsNoCacheForStatus(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsnocacheforstatus.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsStatusRangeGreaterThan(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsstatusrangegreaterthan.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsStatusRangeLessThan(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsstatusrangelessthan.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsFalse(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsfalse.tf", rnd, zoneID)
}

func testAccCloudflareRulesetConfigAllEnabled(rnd, accountID, zoneID string) string {
	return acctest.LoadTestCase("rulesetconfigallenabled.tf", rnd, accountID, zoneID)
}

func testAccCloudflareRulesetRedirectFromList(rnd, accountID string) string {
	return acctest.LoadTestCase("rulesetredirectfromlist.tf", rnd, accountID)
}

func testAccCloudflareRulesetRedirectFromValue(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetredirectfromvalue.tf", rnd, zoneID)
}

func testAccCloudflareRulesetRedirectFromValueWithoutPreservingQueryString(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetredirectfromvaluewithoutpreservingquerystring.tf", rnd, zoneID)
}

func testAccCheckCloudflareRulesetActionParametersOverrideSensitivityForAllRulesetRules(rnd, name, zoneID, zoneName string) string {
	return acctest.LoadTestCase("rulesetactionparametersoverridesensitivityforallrulesetrules.tf", rnd, name, zoneID, zoneName)
}

func testAccCloudflareRulesetRewriteForEmptyQueryString(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetrewriteforemptyquerystring.tf", rnd, zoneID)
}

func testAccCloudflareRulesetRewriteForEmptyPath(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetrewriteforemptypath.tf", rnd, zoneID)
}

func testAccCloudflareRulesetConfigSingleFalseyValue(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetconfigsinglefalseyvalue.tf", rnd, zoneID)
}

func testAccCloudflareRulesetConfigConflictingCacheByDeviceConfigs(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetconfigconflictingcachebydeviceconfigs.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsExplicitCustomKeyCacheKeysQueryStringsExclude(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsexplicitcustomkeycachekeysquerystringsexclude.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsExplicitCustomKeyCacheKeysQueryStringsInclude(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsexplicitcustomkeycachekeysquerystringsinclude.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsHandleDefaultHeaderExcludeOrigin(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingshandledefaultheaderexcludeorigin.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsHandleHeaderExcludeOriginSet(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingshandleheaderexcludeoriginset.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsHandleHeaderExcludeOriginFalse(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingshandleheaderexcludeoriginfalse.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsBypassByDefaultEdge(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsbypassbydefaultedge.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsBypassByDefaultEdgeInvalid(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsbypassbydefaultedgeinvalid.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsBypassBrowser(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsbypassbrowser.tf", rnd, zoneID)
}

func testAccCloudflareRulesetCacheSettingsBypassBrowserInvalid(rnd, zoneID string) string {
	return acctest.LoadTestCase("rulesetcachesettingsbypassbrowserinvalid.tf", rnd, zoneID)
}

func testAccCheckCloudflareRulesetDestroy(s *terraform.State) error {
	return nil
}
