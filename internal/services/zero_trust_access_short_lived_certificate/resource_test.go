package zero_trust_access_short_lived_certificate_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareAccessCACertificate_AccountLevel(t *testing.T) {
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_short_lived_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessCACertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessCACertificateBasic(rnd, domain, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(name, "app_id"),
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
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_short_lived_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessCACertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessCACertificateBasic(rnd, domain, cloudflare.ZoneIdentifier(zoneID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttrSet(name, "app_id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
					resource.TestCheckResourceAttrSet(name, "public_key"),
				),
			},
		},
	})
}

func testAccCloudflareAccessCACertificateBasic(resourceName, domain string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("accesscacertificatebasic.tf", resourceName, domain, identifier.Type, identifier.Identifier)
}

func testAccCheckCloudflareAccessCACertificateDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_access_short_lived_certificate" {
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
