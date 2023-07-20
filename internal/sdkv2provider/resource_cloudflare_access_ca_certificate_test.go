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

func TestAccCloudflareAccessCACertificate_AccountLevel(t *testing.T) {
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_ca_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessCACertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessCACertificateBasic(rnd, domain, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(name, "application_id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
					resource.TestCheckResourceAttrSet(name, "public_key"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessCACertificate_ZoneLevel(t *testing.T) {
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_ca_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareAccessCACertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessCACertificateBasic(rnd, domain, cloudflare.ZoneIdentifier(zoneID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttrSet(name, "application_id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
					resource.TestCheckResourceAttrSet(name, "public_key"),
				),
			},
		},
	})
}

func testAccCloudflareAccessCACertificateBasic(resourceName, domain string, identifier *cloudflare.ResourceContainer) string {
	return fmt.Sprintf(`
resource "cloudflare_access_application" "%[1]s" {
	name     = "%[1]s"
	%[3]s_id = "%[4]s"
	domain   = "%[1]s.%[2]s"
}

resource "cloudflare_access_ca_certificate" "%[1]s" {
  %[3]s_id       = "%[4]s"
  application_id = cloudflare_access_application.%[1]s.id
}`, resourceName, domain, identifier.Type, identifier.Identifier)
}

func testAccCheckCloudflareAccessCACertificateDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_access_ca_certificate" {
			continue
		}

		_, err := client.GetAccessCACertificate(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Access CA certificate still exists")
		}

		_, err = client.GetAccessCACertificate(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Access CA certificate still exists")
		}
	}

	return nil
}
