package zero_trust_device_custom_profile_local_domain_fallback_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"golang.org/x/exp/rand"
)

func TestAccCloudflareFallbackDomain_CustomWithAttachedPolicy(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_custom_profile_local_domain_fallback.%s", rnd)

	randomPrecedence := rand.Intn(10)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "third example domain", "example.net", "1.1.1.1", randomPrecedence),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "domains.#", "1"),
					resource.TestCheckResourceAttr(name, "domains.0.description", "third example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.suffix", "example.net"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server.0", "1.1.1.1"),
					resource.TestCheckResourceAttrSet(name, "policy_id"),
				),
			},
		},
	})
}

func testAccCloudflareFallbackDomain(rnd, accountID, description string, suffix string, dns_server string, precedence int) string {
	return acctest.LoadTestCase("fallbackdomain.tf", rnd, accountID, description, suffix, dns_server, precedence)
}
