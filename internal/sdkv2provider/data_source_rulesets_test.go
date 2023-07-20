package sdkv2provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareRulesetsProviderDataSource_PreventZoneIdAndAccountIdConflicts(t *testing.T) {
	rnd := generateRandomResourceName()
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testCloudflareRulesetsProviderDataSourceConfigConflictingFields(rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("only one of `account_id,zone_id` can be specified")),
			},
		},
	})
}

func testCloudflareRulesetsProviderDataSourceConfigConflictingFields(rnd string) string {
	return fmt.Sprintf(`
data "cloudflare_rulesets" "%[1]s" {
  account_id = "123abc"
  zone_id = "456def"
}
`, rnd)
}

func TestAccCloudflareRulesetsProviderDataSource_RequireOneOfZoneAccountID(t *testing.T) {
	rnd := generateRandomResourceName()
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testCloudflareRulesetsProviderDataSourceRequireOneOfZoneAccountID(rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("one of `account_id,zone_id` must be specified")),
			},
		},
	})
}

func testCloudflareRulesetsProviderDataSourceRequireOneOfZoneAccountID(rnd string) string {
	return fmt.Sprintf(`
data "cloudflare_rulesets" "%[1]s" {
}
`, rnd)
}

func TestAccCloudflareRulesetsProviderDataSource_FetchOWASPRulesetByName(t *testing.T) {
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	name := fmt.Sprintf("data.cloudflare_rulesets.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareRulesetsProviderDataSourceFetchOWASPRulesetByName(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRulesetsDataSourceID(name),
					resource.TestCheckResourceAttr(name, "rulesets.0.name", "Cloudflare OWASP Core Ruleset"),
				),
			},
		},
	})
}

func testCloudflareRulesetsProviderDataSourceFetchOWASPRulesetByName(rnd string, zoneID string) string {
	return fmt.Sprintf(`
data "cloudflare_rulesets" "%[1]s" {
	zone_id = "%[2]s"

	filter {
		name = ".*OWASP.*"
	}
}
`, rnd, zoneID)
}

func TestAccCloudflareRulesetsProviderDataSource_FetchOWASPRulesetByID(t *testing.T) {
	rnd := generateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	name := fmt.Sprintf("data.cloudflare_rulesets.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareRulesetsProviderDataSourceFetchOWASPRulesetByID(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRulesetsDataSourceID(name),
					resource.TestCheckResourceAttr(name, "rulesets.0.name", "Cloudflare OWASP Core Ruleset"),
				),
			},
		},
	})
}

func testCloudflareRulesetsProviderDataSourceFetchOWASPRulesetByID(rnd string, zoneID string) string {
	return fmt.Sprintf(`
data "cloudflare_rulesets" "%[1]s" {
	zone_id = "%[2]s"

	filter {
		id = "4814384a9e5d4991b9815dcfc25d2f1f"
	}
}
`, rnd, zoneID)
}

func testAccCheckCloudflareRulesetsDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		all := s.RootModule().Resources
		rs, ok := all[n]
		if !ok {
			return fmt.Errorf("can't find Rulesets data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Snapshot Rulesets source ID not set")
		}
		return nil
	}
}
