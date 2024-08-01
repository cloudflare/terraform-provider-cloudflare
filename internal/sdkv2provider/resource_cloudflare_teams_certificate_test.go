package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareTeamsCertificate_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTeamsCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsCertificateManagedBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "binding_status", "inactive"),
					resource.TestCheckResourceAttr(name, "gateway_managed", "true"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsCertificateManagedBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_gateway_certificate" "%[1]s" {
	account_id      = "%[2]s"
	gateway_managed = true
}
`, rnd, accountID)
}

func testAccCheckCloudflareTeamsCertificateDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_gateway_certificate" {
			continue
		}

		accountID = rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.TeamsCertificate(context.Background(), accountID, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Teams Certificate still exists")
		}
	}

	return nil
}
