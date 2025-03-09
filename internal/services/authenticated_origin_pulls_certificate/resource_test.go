package authenticated_origin_pulls_certificate_test

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

func init() {
	resource.AddTestSweepers("cloudflare_authenticated_origin_pulls_certificate", &resource.Sweeper{
		Name: "cloudflare_authenticated_origin_pulls_certificate",
		F:    testSweepCloudflareAuthenticatdOriginPullsCertificates,
	})
}

func testSweepCloudflareAuthenticatdOriginPullsCertificates(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	perZoneCertificates, certsErr := client.ListPerZoneAuthenticatedOriginPullsCertificates(context.Background(), zoneID)

	if certsErr != nil {
		tflog.Error(ctx, fmt.Sprintf("failed to fetch per-zone authenticated origin pull certificates: %s", certsErr))
	}

	if len(perZoneCertificates) == 0 {
		tflog.Debug(ctx, "no authenticated origin pull certificates to sweep")
		return nil
	}

	for _, certificate := range perZoneCertificates {
		_, err := client.DeletePerZoneAuthenticatedOriginPullsCertificate(context.Background(), zoneID, certificate.ID)

		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("failed to delete per-zone authenticated origin pull certificate (%s) in zone ID: %s", certificate.ID, zoneID))
		}
	}

	perHostnameCertificates, certsErr := client.ListPerHostnameAuthenticatedOriginPullsCertificates(context.Background(), zoneID)
	if certsErr != nil {
		tflog.Error(ctx, fmt.Sprintf("failed to fetch per-hostname authenticated origin pull certificates: %s", certsErr))
	}

	if len(perHostnameCertificates) == 0 {
		tflog.Debug(ctx, "no authenticated origin pull certificates to sweep")
		return nil
	}

	for _, certificate := range perHostnameCertificates {
		_, err := client.DeletePerHostnameAuthenticatedOriginPullsCertificate(context.Background(), zoneID, certificate.CertID)

		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("failed to delete per-hostname authenticated origin pull certificate (%s) in zone ID: %s", certificate.CertID, zoneID))
		}
	}

	return nil
}

func TestAccCloudflareAuthenticatedOriginPullsCertificatePerZone(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending investigation into correct test setup for reproducibility.")

	var perZoneAOP cloudflare.PerZoneAuthenticatedOriginPullsCertificateDetails
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_authenticated_origin_pulls_certificate.%s", rnd)
	aopType := "per-zone"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAuthenticatedOriginPullsCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAuthenticatedOriginPullsCertificateConfig(zoneID, rnd, aopType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAuthenticatedOriginPullsCertificatePerZoneExists(name, &perZoneAOP),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "type", aopType),
				),
			},
		},
	})
}

func TestAccCloudflareAuthenticatedOriginPullsCertificatePerHostname(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending investigation into correct test setup for reproducibility.")

	var perZoneAOP cloudflare.PerHostnameAuthenticatedOriginPullsCertificateDetails
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_authenticated_origin_pulls_certificate.%s", rnd)
	aopType := "per-hostname"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAuthenticatedOriginPullsCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAuthenticatedOriginPullsCertificateConfig(zoneID, rnd, aopType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAuthenticatedOriginPullsCertificatePerHostnameExists(name, &perZoneAOP),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "type", aopType),
				),
			},
		},
	})
}

func testAccCheckCloudflareAuthenticatedOriginPullsCertificatePerZoneExists(n string, perZoneAOPCert *cloudflare.PerZoneAuthenticatedOriginPullsCertificateDetails) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No cert ID is set")
		}
		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundPerZoneAOPCert, err := client.GetPerZoneAuthenticatedOriginPullsCertificateDetails(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err != nil {
			return err
		}
		if foundPerZoneAOPCert.ID != rs.Primary.ID {
			return fmt.Errorf("cert not found")
		}
		*perZoneAOPCert = foundPerZoneAOPCert
		return nil
	}
}

func testAccCheckCloudflareAuthenticatedOriginPullsCertificatePerHostnameExists(n string, perHostnameAOPCert *cloudflare.PerHostnameAuthenticatedOriginPullsCertificateDetails) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No cert ID is set")
		}
		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundPerHostnameAOPCert, err := client.GetPerHostnameAuthenticatedOriginPullsCertificate(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err != nil {
			return err
		}
		if foundPerHostnameAOPCert.ID != rs.Primary.ID {
			return fmt.Errorf("cert not found")
		}
		*perHostnameAOPCert = foundPerHostnameAOPCert
		return nil
	}
}

func testAccCheckCloudflareAuthenticatedOriginPullsCertificateConfig(zoneID, name, aopType string) string {
	return acctest.LoadTestCase("authenticatedoriginpullscertificateconfig.tf", zoneID, name, aopType)
}

func testAccCheckCloudflareAuthenticatedOriginPullsCertificateDestroy(s *terraform.State) error {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Primary.Attributes["type"] == "per-zone" {
			_, err := client.DeletePerZoneAuthenticatedOriginPullsCertificate(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("error deleting Per-Zone AOP certificate on zone %q: %w", zoneID, err)
			}
		} else if rs.Primary.Attributes["type"] == "per-hostname" {
			_, err := client.DeletePerZoneAuthenticatedOriginPullsCertificate(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("error deleting Per-Zone AOP certificate on zone %q: %w", zoneID, err)
			}
		}
	}
	return nil
}
