package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareFallbackDomain(t *testing.T) {
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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "example domain", "example.com", "['2.2.2.2']"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "domains.0.description", "example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.suffix", "example.com"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server", "['2.2.2.2']"),
				),
			},
			{
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "second example domain", "example_two.com", "['1.1.1.1']"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "domains.0.description", "second example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.host", "example_two.com"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server", "['1.1.1.1']"),
				),
			},
		},
	})
}

func testAccCloudflareFallbackDomain(rnd, accountID string, description string, suffix string, dns_server string) string {
	return fmt.Sprintf(`
resource "cloudflare_split_tunnel" "%[1]s" {
  account_id = "%[2]s"
  domains {
    description = "%[3]s"
    host = "%[4]s"
		dns_server = "%[5]s"
  }
}
`, rnd, accountID, description, suffix, dns_server)
}
