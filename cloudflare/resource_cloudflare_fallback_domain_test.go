package cloudflare

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareFallbackDomain_no_restore_on_delete(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_fallback_domain.%s", rnd)
	default_domain_count := len(default_domains())

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers: testAccProviders,
		CheckDestroy: testAccCheckCloudflareFallbackDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "false", "false", "example domain", "example.com", "2.2.2.2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "domains.#", "1"),
					resource.TestCheckResourceAttr(name, "domains.0.description", "example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.suffix", "example.com"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server.0", "2.2.2.2"),
				),
			},
			{
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "true", "false", "second example domain", "example_two.com", "1.1.1.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "domains.#", fmt.Sprintf("%d", default_domain_count + 1)),
					resource.TestCheckResourceAttr(name, fmt.Sprintf("domains.%d.description", default_domain_count), "second example domain"),
					resource.TestCheckResourceAttr(name, fmt.Sprintf("domains.%d.suffix", default_domain_count), "example_two.com"),
					resource.TestCheckResourceAttr(name, fmt.Sprintf("domains.%d.dns_server.0", default_domain_count), "1.1.1.1"),
				),
			},
		},
	})
}

func TestAccCloudflareFallbackDomain_with_restore_on_delete(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_fallback_domain.%s", rnd)
	default_domain_count := len(default_domains())

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers: testAccProviders,
		CheckDestroy: testAccCheckCloudflareFallbackDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "false", "true", "example domain", "example.com", "2.2.2.2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "domains.#", "1"),
					resource.TestCheckResourceAttr(name, "domains.0.description", "example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.suffix", "example.com"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server.0", "2.2.2.2"),
				),
			},
			{
				// test only changing `include_default_domains`
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "true", "true", "example domain", "example.com", "2.2.2.2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "domains.#", fmt.Sprintf("%d", default_domain_count + 1)),
					resource.TestCheckResourceAttr(name, fmt.Sprintf("domains.%d.description", default_domain_count), "example domain"),
					resource.TestCheckResourceAttr(name, fmt.Sprintf("domains.%d.suffix", default_domain_count), "example.com"),
					resource.TestCheckResourceAttr(name, fmt.Sprintf("domains.%d.dns_server.0", default_domain_count), "2.2.2.2"),
				),
			},
		},
	})
}

func testAccCloudflareFallbackDomain(rnd, accountID string, include_default string, restore_on_delete string, description string, suffix string, dns_server string) string {
	return fmt.Sprintf(`
resource "cloudflare_fallback_domain" "%[1]s" {
  account_id                        = "%[2]s"
  include_default_domains           = %[3]s
  restore_default_domains_on_delete = %[4]s
  domains {
    description = "%[5]s"
    suffix      = "%[6]s"
    dns_server  = ["%[7]s"]
  }
}
`, rnd, accountID, include_default, restore_on_delete, description, suffix, dns_server)
}

func testAccCheckCloudflareFallbackDomainDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_fallback_domain" {
			continue
		}

		result, _ := client.ListFallbackDomains(context.Background(), rs.Primary.ID)
		if rs.Primary.Attributes["restore_default_domains_on_delete"] == "true" {
			default_domain_count := len(default_domains())
			if len(result) != default_domain_count {
				return fmt.Errorf("Deleted Fallback Domains resource has %d domains instead of expected %d default entries.", len(result), default_domain_count)
			}
		} else {
			if len(result) > 0 {
				return fmt.Errorf("Deleted Fallback Domains resource has %d domains, but should be empty.", len(result))
			}
		}
	}

	return nil
}
