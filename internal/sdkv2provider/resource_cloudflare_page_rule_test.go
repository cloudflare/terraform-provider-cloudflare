package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
)

func init() {
	resource.AddTestSweepers("cloudflare_page_rule", &resource.Sweeper{
		Name: "cloudflare_page_rule",
		F:    testSweepCloudflarePageRules,
	})
}

func testSweepCloudflarePageRules(r string) error {
	ctx := context.Background()
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	altZoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")

	if zoneID == "" || altZoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID and CLOUDFLARE_ALT_ZONE_ID must be set for cloudflare_page_rule sweepers")
	}

	pageRules, err := client.ListPageRules(context.Background(), zoneID)
	if err != nil {
		return fmt.Errorf("error listing page rules: %w", err)
	}

	for _, pageRule := range pageRules {
		err := client.DeletePageRule(context.Background(), zoneID, pageRule.ID)
		if err != nil {
			return fmt.Errorf("error deleting page rule %s: %w", pageRule.ID, err)
		}
	}

	altPageRules, err := client.ListPageRules(context.Background(), altZoneID)
	if err != nil {
		return fmt.Errorf("error listing page rules: %w", err)
	}

	for _, pageRule := range altPageRules {
		err := client.DeletePageRule(context.Background(), altZoneID, pageRule.ID)
		if err != nil {
			return fmt.Errorf("error deleting page rule %s: %w", pageRule.ID, err)
		}
	}

	return nil
}

func TestAccCloudflarePageRule_Basic(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigBasic(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					testAccCheckCloudflarePageRuleAttributesBasic(&pageRule),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "target", fmt.Sprintf("%s/", target)),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_FullySpecified(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigFullySpecified(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "target", fmt.Sprintf("%s/", target)),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_ForwardingOnly(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigForwardingOnly(zoneID, target, rnd, rnd+"."+domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "target", fmt.Sprintf("%s/", target)),
					resource.TestCheckResourceAttr(resourceName, "actions.0.forwarding_url.0.url", fmt.Sprintf("http://%s/forward", rnd+"."+domain)),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_ForwardingAndOthers(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigForwardingAndOthers(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "target", target),
					resource.TestCheckResourceAttr(resourceName, "target", fmt.Sprintf("%s/", target)),
				),

				ExpectError: regexp.MustCompile("\"forwarding_url\" cannot be set with any other actions"),
			},
		},
	})
}

func TestAccCloudflarePageRule_DisableZaraz(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigDisableZaraz(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "target", fmt.Sprintf("%s/", target)),
					resource.TestCheckResourceAttr(resourceName, "actions.0.disable_zaraz", "true"),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_Updated(t *testing.T) {
	var before, after cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigBasic(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &before),
				),
			},
			{
				Config: testAccCheckCloudflarePageRuleConfigNewValue(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &after),
					testAccCheckCloudflarePageRuleAttributesUpdated(&after),
					testAccCheckCloudflarePageRuleIDUnchanged(&before, &after),
					resource.TestCheckResourceAttr(resourceName, "target", fmt.Sprintf("%s/updated", target)),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_CreateAfterManualDestroy(t *testing.T) {
	var before, after cloudflare.PageRule
	var initialID string
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigBasic(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &before),
					testAccManuallyDeletePageRule(resourceName, &initialID),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudflarePageRuleConfigNewValue(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &after),
					testAccCheckCloudflarePageRuleRecreated(&before, &after),
					resource.TestCheckResourceAttr(resourceName, "target", fmt.Sprintf("%s/updated", target)),
					resource.TestCheckResourceAttr(resourceName, "actions.0.browser_check", "on"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.rocket_loader", "on"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.ssl", "strict"),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_UpdatingZoneForcesNewResource(t *testing.T) {
	var before, after cloudflare.PageRule
	oldZoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	newZoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	oldTarget := fmt.Sprintf("%s.%s", rnd, os.Getenv("CLOUDFLARE_DOMAIN"))
	newTarget := fmt.Sprintf("%s.%s", rnd, os.Getenv("CLOUDFLARE_ALT_DOMAIN"))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAltDomain(t)
			testAccPreCheckAltZoneID(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigBasic(oldZoneID, oldTarget, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &before),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, oldZoneID),
				),
			},
			{
				Config: testAccCheckCloudflarePageRuleConfigBasic(newZoneID, newTarget, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &after),
					testAccCheckCloudflarePageRuleRecreated(&before, &after),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, newZoneID),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_MinifyAction(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigMinify(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, "actions.0.minify.0.css", "on"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.minify.0.js", "off"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.minify.0.html", "on"),
				),
			},
		},
	})
}

func TestTranformForwardingURL(t *testing.T) {
	key, val, err := transformFromCloudflarePageRuleAction(&cloudflare.PageRuleAction{
		ID: "forwarding_url",
		Value: map[string]interface{}{
			"url":         "http://test.com/forward",
			"status_code": 302,
		},
	})
	if err != nil {
		t.Fatalf("Unexpected error transforming page rule action: %s", err)
	}

	if key != "forwarding_url" {
		t.Fatalf("Unexpected key transforming page rule action. Expected \"forwarding_url\", got \"%s\"", key)
	}

	// the transformed value for a forwarding_url should be [{url: "", "status_code": 302}] (single item slice where the
	// element in the slice is a map)
	if sl, isSlice := val.([]interface{}); !isSlice {
		t.Fatalf("Unexpected value type from transforming page rule action. Expected slice, got %s", reflect.TypeOf(val).Kind())
	} else if len(sl) != 1 {
		t.Fatalf("Unexpected slice length after transforming page rule action. Expected 1, got %d", len(sl))
	} else if _, isMap := sl[0].(map[string]interface{}); !isMap {
		t.Fatalf("Unexpected type in slice after tranforming page rule action. Expected map[string]interface{}, got %s", reflect.TypeOf(sl[0]).Kind())
	}
}

// This test ensures there is no crash while encountering a nil query_string section, which may happen when updating
// existing Page Rule that didn't have this value set previously.
func TestCacheKeyFieldsNilValue(t *testing.T) {
	pageRuleAction, err := transformToCloudflarePageRuleAction(
		context.Background(),
		"cache_key_fields",
		[]interface{}{
			map[string]interface{}{
				"cookie": []interface{}{
					map[string]interface{}{
						"include":        schema.NewSet(schema.HashString, []interface{}{}),
						"check_presence": schema.NewSet(schema.HashString, []interface{}{"next-i18next"}),
					},
				},
				"header": []interface{}{
					map[string]interface{}{
						"check_presence": schema.NewSet(schema.HashString, []interface{}{}),
						"exclude":        schema.NewSet(schema.HashString, []interface{}{}),
						"include":        schema.NewSet(schema.HashString, []interface{}{"x-forwarded-host"}),
					},
				},
				"host": []interface{}{
					map[string]interface{}{
						"resolved": false,
					},
				},
				"query_string": []interface{}{
					interface{}(nil),
				},
				"user": []interface{}{
					map[string]interface{}{
						"device_type": true,
						"geo":         true,
						"lang":        true,
					},
				},
			},
		},
		nil,
	)

	if err != nil {
		t.Fatalf("Unexpected error transforming page rule action: %s", err)
	}

	if !reflect.DeepEqual(pageRuleAction.Value.(map[string]interface{})["query_string"], map[string]interface{}{"include": "*"}) {
		t.Fatalf("Unexpected transformToCloudflarePageRuleAction result, expected %#v, got %#v", map[string]interface{}{"include": "*"}, pageRuleAction.Value.(map[string]interface{})["query_string"])
	}
}

func TestAccCloudflarePageRule_CreatesBrowserCacheTTLIntegerValues(t *testing.T) {
	var pageRule cloudflare.PageRule
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	testAccRunResourceTestSteps(t, []resource.TestStep{
		{
			Config: buildPageRuleConfig(rnd, zoneID, "browser_cache_ttl = 1", target),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
				testAccCheckCloudflarePageRuleHasAction(&pageRule, "browser_cache_ttl", float64(1)),
				resource.TestCheckResourceAttr(resourceName, "actions.0.browser_cache_ttl", "1"),
			),
		},
	})
}

func TestAccCloudflarePageRule_CreatesBrowserCacheTTLThatRespectsExistingHeaders(t *testing.T) {
	var pageRule cloudflare.PageRule
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	testAccRunResourceTestSteps(t, []resource.TestStep{
		{
			Config: buildPageRuleConfig(rnd, zoneID, "browser_cache_ttl = 0", target),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
				resource.TestCheckResourceAttr(resourceName, "actions.0.browser_cache_ttl", "0"),
				testAccCheckCloudflarePageRuleHasAction(&pageRule, "browser_cache_ttl", float64(0)),
			),
		},
	})
}

func TestAccCloudflarePageRule_UpdatesBrowserCacheTTLToSameValue(t *testing.T) {
	var pageRule cloudflare.PageRule
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	testAccRunResourceTestSteps(t, []resource.TestStep{
		{
			Config: buildPageRuleConfig(rnd, zoneID, "browser_cache_ttl = 1", target),
		},
		{
			Config: buildPageRuleConfig(rnd, zoneID, `browser_cache_ttl = 1
			browser_check = "on"`, target),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
				testAccCheckCloudflarePageRuleHasAction(&pageRule, "browser_cache_ttl", float64(1)),
				resource.TestCheckResourceAttr(resourceName, "actions.0.browser_cache_ttl", "1"),
			),
		},
	})
}

func TestAccCloudflarePageRule_UpdatesBrowserCacheTTLThatRespectsExistingHeaders(t *testing.T) {
	var pageRule cloudflare.PageRule
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	testAccRunResourceTestSteps(t, []resource.TestStep{
		{
			Config: buildPageRuleConfig(rnd, zoneID, "browser_cache_ttl = 1", target),
		},
		{
			Config: buildPageRuleConfig(rnd, zoneID, "browser_cache_ttl = 0", target),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
				testAccCheckCloudflarePageRuleHasAction(&pageRule, "browser_cache_ttl", float64(0)),
				resource.TestCheckResourceAttr(resourceName, "actions.0.browser_cache_ttl", "0"),
			),
		},
	})
}

func TestAccCloudflarePageRule_DeletesBrowserCacheTTLThatRespectsExistingHeaders(t *testing.T) {
	var pageRule cloudflare.PageRule
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	testAccRunResourceTestSteps(t, []resource.TestStep{
		{
			Config: buildPageRuleConfig(rnd, zoneID, "browser_cache_ttl = 0", target),
		},
		{
			Config: buildPageRuleConfig(rnd, zoneID, `browser_check = "on"`, target),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
				resource.TestCheckResourceAttr(resourceName, "actions.0.browser_cache_ttl", ""),
			),
		},
	})
}

func TestAccCloudflarePageRule_EdgeCacheTTLNotClobbered(t *testing.T) {
	var before, after cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigWithEdgeCacheTtl(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &before),
					resource.TestCheckResourceAttr(resourceName, "actions.0.edge_cache_ttl", "10"),
				),
			},
			{
				Config: testAccCheckCloudflarePageRuleConfigWithEdgeCacheTtlAndAlwaysOnline(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &after),
					resource.TestCheckResourceAttr(resourceName, "actions.0.edge_cache_ttl", "10"),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_CacheKeyFieldsBasic(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigCacheKeyFields(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.cookie.0.check_presence.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.cookie.0.include.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.0.check_presence.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.0.include.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.host.0.resolved", "true"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.query_string.0.exclude.#", "1"),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_CacheKeyFieldsIgnoreQueryStringOrdering(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	pageRuleTarget := fmt.Sprintf("%s.%s", rnd, domain)
	resourceName := fmt.Sprintf("cloudflare_page_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigCacheKeyFieldsWithUnorderedEntries(zoneID, rnd, pageRuleTarget),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.cookie.0.check_presence.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.cookie.0.include.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.0.check_presence.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.0.include.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.host.0.resolved", "true"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.query_string.0.include.#", "7"),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_CacheKeyFieldsExcludeAllQueryString(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	pageRuleTarget := fmt.Sprintf("%s.%s", rnd, domain)
	resourceName := fmt.Sprintf("cloudflare_page_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigCacheKeyFieldsIgnoreAllQueryString(zoneID, rnd, pageRuleTarget),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.cookie.0.check_presence.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.cookie.0.include.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.0.check_presence.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.0.include.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.host.0.resolved", "true"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.query_string.0.ignore", "true"),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_CacheKeyFieldsExcludeMultipleValuesQueryString(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	pageRuleTarget := fmt.Sprintf("%s.%s", rnd, domain)
	resourceName := fmt.Sprintf("cloudflare_page_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigCacheKeyFieldsExcludeMultipleValuesQueryString(zoneID, rnd, pageRuleTarget),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.cookie.0.check_presence.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.cookie.0.include.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.0.check_presence.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.0.include.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.host.0.resolved", "true"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.query_string.0.exclude.#", "2"),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_CacheKeyFieldsNoQueryStringValuesDefined(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigCacheKeyFieldsNoQueryStringValuesDefined(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.0.exclude.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.host.0.resolved", "false"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.user.0.device_type", "true"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.user.0.geo", "true"),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_CacheKeyFieldsIncludeAllQueryStringValues(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigCacheKeyFieldsIncludeAllQueryStringValues(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.0.exclude.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.host.0.resolved", "false"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.user.0.device_type", "true"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.user.0.geo", "true"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.query_string.0.ignore", "false"),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_CacheKeyFieldsIncludeMultipleValuesQueryString(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	pageRuleTarget := fmt.Sprintf("%s.%s", rnd, domain)
	resourceName := fmt.Sprintf("cloudflare_page_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigCacheKeyFieldsIncludeMultipleValuesQueryString(zoneID, rnd, pageRuleTarget),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.cookie.0.check_presence.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.cookie.0.include.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.0.check_presence.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.0.include.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.host.0.resolved", "true"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.query_string.0.include.#", "2"),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_EmptyCookie(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	pageRuleTarget := fmt.Sprintf("%s.%s", rnd, domain)
	resourceName := fmt.Sprintf("cloudflare_page_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleEmtpyCookie(zoneID, rnd, pageRuleTarget),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.host.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.query_string.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.user.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.cookie.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.header.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.host.0.resolved", "true"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.query_string.0.include.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.user.0.device_type", "true"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.user.0.geo", "false"),
					resource.TestCheckResourceAttr(resourceName, "actions.0.cache_key_fields.0.user.0.lang", "false"),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_CacheTTLByStatus(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigCacheTTLByStatus(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
				),
			},
		},
	})
}

func testAccCheckCloudflarePageRuleRecreated(before, after *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			return fmt.Errorf("expected change of PageRule Ids, but both were %v", before.ID)
		}
		return nil
	}
}

func testAccCheckCloudflarePageRuleIDUnchanged(before, after *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID != after.ID {
			return fmt.Errorf("ID should not change suring in place update, but got change %s -> %s", before.ID, after.ID)
		}
		return nil
	}
}

func testAccCheckCloudflarePageRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_page_rule" {
			continue
		}

		_, err := client.PageRule(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("PageRule still exists")
		}
	}

	return nil
}

func testAccCheckCloudflarePageRuleAttributesBasic(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// check the api only has attributes we set non-empty values for
		// this covers on/off attribute types and setting enum-type strings

		actionMap := pageRuleActionsToMap(pageRule.Actions)
		if val, ok := actionMap["ssl"]; ok {
			if _, ok := val.(string); !ok || val != "flexible" {
				return fmt.Errorf("'ssl' not specified correctly at api, found: %q", val)
			}
		} else {
			return fmt.Errorf("'ssl' not specified at api")
		}

		if len(pageRule.Actions) != 1 {
			return fmt.Errorf("api should only have attributes we set non-empty (%d) but got %d: %#v", 2, len(pageRule.Actions), pageRule.Actions)
		}

		return nil
	}
}

func testAccCheckCloudflarePageRuleAttributesUpdated(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		actionMap := pageRuleActionsToMap(pageRule.Actions)

		if _, ok := actionMap["disable_apps"]; ok {
			return fmt.Errorf("'disable_apps' found at api, but we should have removed it")
		}

		if val, ok := actionMap["browser_check"]; ok {
			if _, ok := val.(string); !ok || val != "on" { // lots of booleans get mapped to on/off at api
				return fmt.Errorf("'browser_check' not specified correctly at api, found: '%v'", val)
			}
		} else {
			return fmt.Errorf("'browser_check' not specified at api")
		}

		if val, ok := actionMap["ssl"]; ok {
			if _, ok := val.(string); !ok || val != "strict" {
				return fmt.Errorf("'ssl' not specified correctly at api, found: %q", val)
			}
		} else {
			return fmt.Errorf("'ssl' not specified at api")
		}

		if val, ok := actionMap["rocket_loader"]; ok {
			if _, ok := val.(string); !ok || val != "on" {
				return fmt.Errorf("'rocket_loader' not specified correctly at api, found: %q", val)
			}
		} else {
			return fmt.Errorf("'rocket_loader' not specified at api")
		}

		if len(pageRule.Actions) != 3 {
			return fmt.Errorf("api should only have attributes we set non-empty (%d) but got %d: %#v", 4, len(pageRule.Actions), pageRule.Actions)
		}

		return nil
	}
}

func testAccCheckCloudflarePageRuleExists(n string, pageRule *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No PageRule ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundPageRule, err := client.PageRule(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundPageRule.ID != rs.Primary.ID {
			return fmt.Errorf("PageRule not found")
		}

		*pageRule = foundPageRule

		return nil
	}
}

func testAccManuallyDeletePageRule(name string, initialID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		*initialID = rs.Primary.ID
		err := client.DeletePageRule(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflarePageRuleConfigMinify(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions {
		minify {
			js = "off"
			css = "on"
			html = "on"
		}
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigBasic(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions {
		ssl = "flexible"
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigNewValue(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s/updated"
	actions {
		browser_check = "on"
		ssl = "strict"
		rocket_loader = "on"
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigFullySpecified(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions {
		browser_check = "on"
		browser_cache_ttl = 0
		email_obfuscation = "on"
		ip_geolocation = "on"
		server_side_exclude = "on"
		disable_apps = true
		disable_performance = true
		disable_security = true
		cache_level = "bypass"
		security_level = "essentially_off"
		ssl = "flexible"
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigForwardingOnly(zoneID, target, rnd, zoneName string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions {
		// on/off options cannot even be set to off without causing error
		forwarding_url {
			url = "http://%[4]s/forward"
			status_code = 301
		}
	}
}`, zoneID, target, rnd, zoneName)
}

func testAccCheckCloudflarePageRuleConfigForwardingAndOthers(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions {
		disable_security = true
		forwarding_url {
			url = "http://%s/forward"
			status_code = 301
		}
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigDisableZaraz(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions {
		disable_zaraz = true
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigWithEdgeCacheTtl(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions {
		ssl = "flexible"
		edge_cache_ttl = 10
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigWithEdgeCacheTtlAndAlwaysOnline(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions {
		edge_cache_ttl = 10
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigCacheKeyFields(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions {
		cache_key_fields {
			cookie {
				check_presence = ["cookie_presence"]
				include = ["cookie_include"]
			}
			header {
				check_presence = ["header_presence"]
				include = ["header_include"]
			}
			host {
				resolved = true
			}
			query_string {
				exclude = ["qs_exclude"]
			}
			user {}
		}
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigCacheKeyFieldsWithUnorderedEntries(zoneID, rnd, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[3]s"
	actions {
		cache_key_fields {
			cookie {
				check_presence = ["cookie_presence"]
				include = ["cookie_include"]
			}
			header {
				check_presence = ["header_presence"]
				include = ["header_include"]
			}
			host {
				resolved = true
			}
			query_string {
				include = [
          "test.anothertest",
          "test.regiontest",
          "test.devicetest",
          "test.testthis",
          "test.hello",
          "test.segmenttest",
          "test.usertype"
				]
			}
			user {}
		}
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigCacheKeyFieldsIgnoreAllQueryString(zoneID, rnd, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[3]s"
	actions {
		cache_key_fields {
			cookie {
				check_presence = ["cookie_presence"]
				include = ["cookie_include"]
			}
			header {
				check_presence = ["header_presence"]
				include = ["header_include"]
			}
			host {
				resolved = true
			}
			query_string {
				ignore = true
			}
			user {}
		}
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigCacheKeyFieldsExcludeMultipleValuesQueryString(zoneID, rnd, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[3]s"
	actions {
		cache_key_fields {
			cookie {
				check_presence = ["cookie_presence"]
				include = ["cookie_include"]
			}
			header {
				check_presence = ["header_presence"]
				include = ["header_include"]
			}
			host {
				resolved = true
			}
			query_string {
				exclude = ["query1", "query2"]
			}
			user {}
		}
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigCacheKeyFieldsNoQueryStringValuesDefined(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions {
		cache_key_fields {
			header {
				exclude = ["origin"]
			}
			host {}
			query_string {}
			user {
				device_type = true
				geo = true
			}
		}
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigCacheKeyFieldsIncludeAllQueryStringValues(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions {
		cache_key_fields {
			header {
				exclude = ["origin"]
			}
			host {}
			query_string {
				ignore = false
			}
			user {
				device_type = true
				geo = true
			}
		}
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigCacheKeyFieldsIncludeMultipleValuesQueryString(zoneID, rnd, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[3]s"
	actions {
		cache_key_fields {
			cookie {
				check_presence = ["cookie_presence"]
				include = ["cookie_include"]
			}
			header {
				check_presence = ["header_presence"]
				include = ["header_include"]
			}
			host {
				resolved = true
			}
			query_string {
				include = ["query1", "query2"]
			}
			user {}
		}
	}
}`, zoneID, target, rnd)
}

func testAccCheckCloudflarePageRuleConfigCacheTTLByStatus(zoneID, target, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions {
		cache_ttl_by_status {
			codes = "200-299"
			ttl = 300
		}
		cache_ttl_by_status {
			codes = "300-399"
			ttl = 60
		}
		cache_ttl_by_status {
			codes = "400-403"
			ttl = -1
		}
		cache_ttl_by_status {
			codes = "404"
			ttl = 30
		}
		cache_ttl_by_status {
			codes = "405-499"
			ttl = -1
		}
		cache_ttl_by_status {
			codes = "500-599"
			ttl = 0
		}
	}
}`, zoneID, target, rnd)
}
func buildPageRuleConfig(rnd, zoneID, actions, target string) string {
	return fmt.Sprintf(`
		resource "cloudflare_page_rule" "%[1]s" {
			zone_id = "%[2]s"
			target = "%[3]s"
			actions {
				%[4]s
			}
		}`,
		rnd,
		zoneID,
		target,
		actions,
	)
}

func testAccRunResourceTestSteps(t *testing.T, testSteps []resource.TestStep) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps:             testSteps,
	})
}

func testAccCheckCloudflarePageRuleHasAction(pageRule *cloudflare.PageRule, key string, value interface{}) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		for _, pageRuleAction := range pageRule.Actions {
			if pageRuleAction.ID == key && pageRuleAction.Value == value {
				return nil
			}
		}
		return fmt.Errorf("cloudflare page rule action not found %#v:%#v\nAction State\n%#v", key, value, pageRule.Actions)
	}
}

func testAccCheckCloudflarePageRuleEmtpyCookie(zoneID, rnd, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[3]s"
	actions {
    cache_key_fields {
      host {
        resolved = true
      }
      query_string {
        ignore = true
      }
      user {
        device_type = true
        geo         = false
        lang        = false
      }
    }
  }
}`, zoneID, target, rnd)
}
