package cloudflare

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
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
		CheckDestroy: testAccCheckCloudflareFallbackDomainDestroy,
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "example domain", "example.com", "1.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "domains.#", "1"),
					resource.TestCheckResourceAttr(name, "domains.0.description", "example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.suffix", "example.com"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server.0", "1.0.0.1"),
				),
			},
			{
				Config: testAccCloudflareFallbackDomain(rnd, accountID, "second example domain", "example.net", "1.1.1.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "domains.#", "1"),
					resource.TestCheckResourceAttr(name, "domains.0.description", "second example domain"),
					resource.TestCheckResourceAttr(name, "domains.0.suffix", "example.net"),
					resource.TestCheckResourceAttr(name, "domains.0.dns_server.0", "1.1.1.1"),
				),
			},
		},
	})
}

func testAccCloudflareFallbackDomain(rnd, accountID string, description string, suffix string, dns_server string) string {
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

func testAccCheckCloudflareFallbackDomainDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_fallback_domain" {
			continue
		}

		result, _ := client.ListFallbackDomains(context.Background(), rs.Primary.ID)
		if len(result) == 0 {
			return errors.New("deleted Fallback Domain resource has does not include default domains")
		}
	}

	return nil
}
