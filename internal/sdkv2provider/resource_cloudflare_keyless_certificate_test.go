package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareKeylessSSL_Basic(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_keyless_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareKeylessCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareKeylessCertificate(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "port", "24008"),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "host", domain),
				),
			},
		},
	})
}

func testAccCheckCloudflareKeylessCertificateDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_keyless_certificate" {
			continue
		}

		ctx := context.Background()
		err := retry.RetryContext(ctx, time.Second*10, func() *retry.RetryError {
			keylessCertificates, err := client.ListKeylessSSL(ctx, zoneID)
			if err != nil {
				return retry.NonRetryableError(fmt.Errorf("failed to fetch keyless certificate: %w", err))
			}

			for _, keylessCertificate := range keylessCertificates {
				if keylessCertificate.ID == rs.Primary.Attributes["id"] {
					return retry.RetryableError(fmt.Errorf("keyless certificate cleanup is processing"))
				}
			}

			return nil
		})
		if err != nil {
			return errors.New("failed to initiate retries for Keyless SSL deletion")
		}
	}

	return nil
}

func testAccCloudflareKeylessCertificate(resourceName, zoneId string, domain string) string {
	expiry := time.Now().Add(time.Hour * 730)
	cert, _, _ := utils.GenerateEphemeralCertAndKey([]string{domain}, expiry)

	return fmt.Sprintf(`
resource "cloudflare_keyless_certificate" "%[1]s" {
  zone_id       = "%[2]s"
  bundle_method = "force"
  name          = "%[1]s"
  host          = "%[3]s"
  port          = 24008
  certificate   = <<EOT
%[4]s
  EOT
}`, resourceName, zoneId, domain, cert)
}
