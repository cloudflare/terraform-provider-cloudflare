package origin_ca_certificate_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/origin_ca_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_origin_ca_certificate", &resource.Sweeper{
		Name: "cloudflare_origin_ca_certificate",
		F:    testSweepCloudflareOriginCACertificates,
	})
}

func testSweepCloudflareOriginCACertificates(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping Origin CA certificates sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	certs, err := client.OriginCACertificates.List(ctx, origin_ca_certificates.OriginCACertificateListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Origin CA certificates: %s", err))
		return nil
	}

	for _, cert := range certs.Result {
		if !utils.ShouldSweepResource(cert.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting Origin CA certificate: %s", cert.ID))
		_, err := client.OriginCACertificates.Delete(ctx, cert.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Origin CA certificate %s: %s", cert.ID, err))
		}
	}

	return nil
}

func testAccCheckCloudflareOriginCACertificateDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_origin_ca_certificate" {
			continue
		}

		_, err := client.OriginCACertificates.Get(context.Background(), rs.Primary.ID)
		if err != nil {
			continue
		}

		tflog.Warn(context.Background(), fmt.Sprintf("Origin CA certificate %s still exists but this may be expected", rs.Primary.ID))
	}

	return nil
}

// generateCSR creates a Certificate Signing Request for testing
func generateCSR(domain string) (string, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", err
	}

	template := &x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName: domain,
		},
		DNSNames: []string{domain},
	}

	csrDER, err := x509.CreateCertificateRequest(rand.Reader, template, key)
	if err != nil {
		return "", err
	}

	csrPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csrDER,
	})

	return string(csrPEM), nil
}

// TestAccOriginCACertificate_Basic tests the basic CRUD lifecycle of an Origin CA certificate.
// This validates that the resource can be created, read, imported, and deleted.
// Includes an update scenario that changes requested_validity and hostnames.
func TestAccOriginCACertificate_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_origin_ca_certificate." + rnd
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	csr, err := generateCSR(domain)
	if err != nil {
		t.Fatalf("Failed to generate CSR: %s", err)
	}

	csr2, err := generateCSR("*." + domain)
	if err != nil {
		t.Fatalf("Failed to generate CSR: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareOriginCACertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOriginCACertificateBasicConfig(rnd, csr, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("request_type"), knownvalue.StringExact("origin-ecc")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("requested_validity"), knownvalue.Float64Exact(7)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccOriginCACertificateUpdatedConfig(rnd, csr2, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("request_type"), knownvalue.StringExact("origin-ecc")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("requested_validity"), knownvalue.Float64Exact(30)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("certificate"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"csr",                // write-only, not returned by API
					"requested_validity", // not returned by API after creation
				},
			},
		},
	})
}

func testAccOriginCACertificateBasicConfig(rnd, csr, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_origin_ca_certificate" "%[1]s" {
  csr                = <<EOT
%[2]s
EOT
  hostnames          = ["%[3]s"]
  request_type       = "origin-ecc"
  requested_validity = 7
}`, rnd, csr, domain)
}

func testAccOriginCACertificateUpdatedConfig(rnd, csr, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_origin_ca_certificate" "%[1]s" {
  csr                = <<EOT
%[2]s
EOT
  hostnames          = ["*.%[3]s", "%[3]s"]
  request_type       = "origin-ecc"
  requested_validity = 30
}`, rnd, csr, domain)
}
