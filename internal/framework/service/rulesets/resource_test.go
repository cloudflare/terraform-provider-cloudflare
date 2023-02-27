package rulesets_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareRuleset_RateLimit(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_ZONE_NAME")

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRulesetDestroy,
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
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("zone/%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCheckCloudflareRulesetDestroy(s *terraform.State) error {
	return nil
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
