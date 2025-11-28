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

func TestMain(m *testing.M) {
	resource.TestMain(m)
}


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
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping authenticated origin pulls certificates sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	perZoneCertificates, certsErr := client.ListPerZoneAuthenticatedOriginPullsCertificates(ctx, zoneID)

	if certsErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch per-zone authenticated origin pull certificates: %s", certsErr))
		return certsErr
	}

	if len(perZoneCertificates) == 0 {
		tflog.Info(ctx, "No per-zone authenticated origin pull certificates to sweep")
	} else {
		for _, certificate := range perZoneCertificates {
			// Use standard filtering helper
			if !utils.ShouldSweepResource(certificate.ID) {
				continue
			}

			tflog.Info(ctx, fmt.Sprintf("Deleting per-zone authenticated origin pull certificate: %s (zone: %s)", certificate.ID, zoneID))
			_, err := client.DeletePerZoneAuthenticatedOriginPullsCertificate(ctx, zoneID, certificate.ID)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete per-zone authenticated origin pull certificate %s: %s", certificate.ID, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted per-zone authenticated origin pull certificate: %s", certificate.ID))
		}
	}

	perHostnameCertificates, certsErr := client.ListPerHostnameAuthenticatedOriginPullsCertificates(ctx, zoneID)
	if certsErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch per-hostname authenticated origin pull certificates: %s", certsErr))
		return certsErr
	}

	if len(perHostnameCertificates) == 0 {
		tflog.Info(ctx, "No per-hostname authenticated origin pull certificates to sweep")
		return nil
	}

	for _, certificate := range perHostnameCertificates {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(certificate.CertID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting per-hostname authenticated origin pull certificate: %s (zone: %s)", certificate.CertID, zoneID))
		_, err := client.DeletePerHostnameAuthenticatedOriginPullsCertificate(ctx, zoneID, certificate.CertID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete per-hostname authenticated origin pull certificate %s: %s", certificate.CertID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted per-hostname authenticated origin pull certificate: %s", certificate.CertID))
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
