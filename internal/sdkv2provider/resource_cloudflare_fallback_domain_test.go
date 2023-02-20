package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
)

func TestAccCloudflareFallbackDomain_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_fallback_domain.%s", rnd)

	resource.Test(t, resource.TestCase{
		CheckDestroy: testAccCheckCloudflareFallbackDomainDestroy,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
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

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_fallback_domain.%s", rnd)

	resource.Test(t, resource.TestCase{
		CheckDestroy: testAccCheckCloudflareFallbackDomainDestroy,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
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
			{
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "third example domain", "example.net", "1.1.1.1"),
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

func TestAccCloudflareFallbackDomain_WithAttachedPolicy(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_fallback_domain.%s", rnd)

	resource.Test(t, resource.TestCase{
		CheckDestroy: testAccCheckCloudflareFallbackDomainDestroy,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "third example domain", "example.net", "1.1.1.1"),
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

func testAccCloudflareDefaultFallbackDomain(rnd, accountID string, description string, suffix string, dns_server string) string {
	return fmt.Sprintf(`
resource "cloudflare_fallback_domain" "%[1]s" {
  account_id = "%[2]s"
  domains {
    description = "%[3]s"
    suffix      = "%[4]s"
    dns_server  = ["%[5]s"]
  }
}
`, rnd, accountID, description, suffix, dns_server)
}

func testAccCloudflareFallbackDomain(rnd, accountID, description string, suffix string, dns_server string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_settings_policy" "%[1]s" {
	account_id                = "%[2]s"
	allow_mode_switch         = true
	allow_updates             = true
	allowed_to_leave          = true
	auto_connect              = 0
	captive_portal            = 5
	disable_auto_fallback     = true
	enabled                   = true
	match                     = "identity.email == \"foo@example.com\""
	name                      = "%[1]s"
	precedence                = 10
	support_url               = "support_url"
	switch_locked             = true
	exclude_office_ips		  = false
}

resource "cloudflare_fallback_domain" "%[1]s" {
  account_id = "%[2]s"
  domains {
    description = "%[3]s"
    suffix      = "%[4]s"
    dns_server  = ["%[5]s"]
  }
	policy_id = "${cloudflare_device_settings_policy.%[1]s.id}"
}
`, rnd, accountID, description, suffix, dns_server)
}

func testAccCheckCloudflareFallbackDomainDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_fallback_domain" {
			continue
		}

		accountID, policyID := parseDevicePolicyID(rs.Primary.ID)

		if policyID == "" {
			// Deletion of fallback domains should result in a reset to the default fallback domains.
			// Default device settings policies, and their fallback domains, cannot be deleted - only reset to default.
			result, _ := client.ListFallbackDomains(context.Background(), rs.Primary.ID)
			if len(result) == 0 {
				return errors.New("deleted Fallback Domain resource has does not include default domains")
			}
		} else {
			// For fallback domains on a non-default device settings policy, only need to check for the deletion of the policy.
			_, err := client.GetDeviceSettingsPolicy(context.Background(), accountID, policyID)
			if err == nil {
				return fmt.Errorf("device settings policy still exists")
			}
		}
	}

	return nil
}
