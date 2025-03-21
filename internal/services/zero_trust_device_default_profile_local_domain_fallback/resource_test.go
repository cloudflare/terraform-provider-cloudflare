package zero_trust_device_default_profile_local_domain_fallback_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var (
	accountID = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
)

func TestAccCloudflareFallbackDomain_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_default_profile_local_domain_fallback.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDefaultFallbackDomain(rnd, accountID, "example domain", "example.com", "1.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "domains.#", "1"),
					resource.TestCheckResourceAttr(name, "domains.0.description", "example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.suffix", "example.com"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server.0", "1.0.0.1"),
					resource.TestCheckNoResourceAttr(name, "policy_id"),
				),
			},
		},
	})
}

func TestAccCloudflareFallbackDomain_DefaultPolicy(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_default_profile_local_domain_fallback.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDefaultFallbackDomain(rnd, accountID, "second example domain", "example.net", "1.1.1.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "domains.#", "1"),
					resource.TestCheckResourceAttr(name, "domains.0.description", "second example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.suffix", "example.net"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server.0", "1.1.1.1"),
					resource.TestCheckNoResourceAttr(name, "policy_id"),
				),
			},
		},
	})
}

func testAccCloudflareDefaultFallbackDomain(rnd, accountID string, description string, suffix string, dns_server string) string {
	return acctest.LoadTestCase("defaultfallbackdomain.tf", rnd, accountID, description, suffix, dns_server)
}
