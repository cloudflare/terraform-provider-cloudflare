package ruleset_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareRulesetsProviderDataSource_PreventZoneIdAndAccountIdConflicts(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testCloudflareRulesetsProviderDataSourceConfigConflictingFields(rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("Exactly one of these attributes must be configured")),
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
	rnd := utils.GenerateRandomResourceName()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testCloudflareRulesetsProviderDataSourceRequireOneOfZoneAccountID(rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("Exactly one of these attributes must be configured")),
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
