package keyless_certificate_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareKeylessSSL_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_keyless_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareKeylessCertificateDestroy,
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
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

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

	return acctest.LoadTestCase("keylesscertificate.tf", resourceName, zoneId, domain, cert)
}
