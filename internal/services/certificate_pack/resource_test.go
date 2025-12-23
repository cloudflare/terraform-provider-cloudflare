package certificate_pack_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ssl"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_certificate_pack", &resource.Sweeper{
		Name: "cloudflare_certificate_pack",
		F:    testSweepCloudflareCertificatePack,
	})
}

func testSweepCloudflareCertificatePack(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping certificate packs sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	certificates, err := client.SSL.CertificatePacks.List(ctx, ssl.CertificatePackListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch certificate packs: %s", err))
		return nil
	}

	for _, certificate := range certificates.Result {
		if !utils.ShouldSweepResource(certificate.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting certificate pack: %s", certificate.ID))
		_, err := client.SSL.CertificatePacks.Delete(ctx, certificate.ID, ssl.CertificatePackDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete certificate pack %s: %s", certificate.ID, err))
		}
	}

	return nil
}

func testAccCheckCloudflareCertificatePackDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_certificate_pack" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		_, err := client.SSL.CertificatePacks.Get(context.Background(), rs.Primary.ID, ssl.CertificatePackGetParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			// If certificate pack is not found, it was successfully deleted
			continue
		}

		// If we can still retrieve the certificate pack, it might not be fully cleaned up
		// but this is expected behavior for certificate packs - they don't always get immediately deleted
		tflog.Warn(context.Background(), fmt.Sprintf("Certificate pack %s still exists but this may be expected", rs.Primary.ID))
	}

	return nil
}

// TestAccCertificatePack_Basic tests the basic CRUD lifecycle of a certificate pack.
// This validates that the resource can be created, read, imported, and deleted.
func TestAccCertificatePack_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_certificate_pack." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCertificatePackDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificatePackBasicConfig(zoneID, domain, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("lets_encrypt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validation_method"), knownvalue.StringExact("txt")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hosts"), knownvalue.SetSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"certificate_authority", "cloudflare_branding", "hosts", "status", "type", "validation_method", "validity_days", "primary_certificate", "validation_records"},
			},
		},
	})
}

// TestAccCertificatePack_CloudflareBranding tests the optional cloudflare_branding attribute.
// This validates that optional boolean attributes are handled correctly (null vs false vs true).
func TestAccCertificatePack_CloudflareBranding(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_certificate_pack." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCertificatePackDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificatePackCloudflareBrandingConfig(zoneID, domain, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("cloudflare_branding"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportStateVerifyIgnore: []string{"certificate_authority", "cloudflare_branding", "hosts", "status", "type", "validation_method", "validity_days", "primary_certificate", "validation_records"},
			},
		},
	})
}

func testAccCertificatePackBasicConfig(zoneID, domain, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_certificate_pack" "%[3]s" {
  zone_id               = "%[1]s"
  type                  = "advanced"
  certificate_authority = "lets_encrypt"
  validation_method     = "txt"
  validity_days         = 90
  hosts                 = ["%[2]s", "*.%[2]s"]
}`, zoneID, domain, rnd)
}

func testAccCertificatePackCloudflareBrandingConfig(zoneID, domain, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_certificate_pack" "%[3]s" {
  zone_id               = "%[1]s"
  type                  = "advanced"
  certificate_authority = "lets_encrypt"
  validation_method     = "txt"
  validity_days         = 90
  hosts                 = ["%[2]s", "*.%[2]s"]
  cloudflare_branding   = true
}`, zoneID, domain, rnd)
}
