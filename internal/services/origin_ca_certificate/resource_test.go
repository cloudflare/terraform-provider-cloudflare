package origin_ca_certificate_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping origin CA certificates sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	// Origin CA certificates don't have a List endpoint - they can only be retrieved by ID
	// The certificates are revoked on delete, not removed from the API
	// We'll attempt to list using the SDK method if available
	certs, err := client.ListOriginCACertificates(ctx, cloudflare.ListOriginCertificatesParams{
		ZoneID: zoneID,
	})
	if err != nil {
		// If we can't list certificates (zone invalid, API error, etc), log and skip gracefully
		tflog.Warn(ctx, fmt.Sprintf("Failed to fetch origin CA certificates for zone %s: %s (skipping)", zoneID, err))
		return nil
	}

	if len(certs) == 0 {
		tflog.Info(ctx, "No origin CA certificates to sweep")
		return nil
	}

	for _, cert := range certs {
		// Skip already revoked certificates
		if cert.RevokedAt != (time.Time{}) {
			continue
		}

		// Use standard filtering helper on certificate ID
		if !utils.ShouldSweepResource(cert.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Revoking origin CA certificate: %s", cert.ID))
		_, err := client.RevokeOriginCACertificate(ctx, cert.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to revoke origin CA certificate %s: %s", cert.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Revoked origin CA certificate: %s", cert.ID))
	}

	return nil
}

func TestAccCloudflareOriginCACertificate_Basic(t *testing.T) {
	var cert cloudflare.OriginCACertificate
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_origin_ca_certificate." + rnd

	csr, err := generateCSR(zoneName)
	if err != nil {
		t.Errorf("unable to generate CSR: %v", err)
		return
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareOriginCACertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareOriginCACertificateConfigBasic(rnd, zoneName, csr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareOriginCACertificateExists(name, &cert),
					testAccCheckCloudflareOriginCACertificateAttributes(zoneName, &cert),
					resource.TestMatchResourceAttr(name, "id", regexp.MustCompile("^[0-9]+$")),
					resource.TestCheckResourceAttr(name, "csr", csr),
					resource.TestCheckResourceAttr(name, "request_type", "origin-rsa"),
					resource.TestCheckResourceAttr(name, "requested_validity", "7"),
				),
			},
			// {
			// 	ResourceName: name,
			// 	ImportState:  true,
			// },
		},
	})
}

// func TestCalculateRequestedValidityFromCertificate(t *testing.T) {
// 	testCases := []struct {
// 		NotBefore time.Time
// 		NotAfter  time.Time
// 		expected  int
// 	}{
// 		{
// 			NotBefore: time.Date(2022, 1, 18, 10, 48, 0, 0, time.UTC),
// 			NotAfter:  time.Date(2023, 1, 18, 10, 48, 0, 0, time.UTC),
// 			expected:  365,
// 		},
// 		{
// 			NotBefore: time.Date(2022, 1, 18, 10, 48, 0, 0, time.UTC),
// 			NotAfter:  time.Date(2022, 1, 25, 10, 48, 0, 0, time.UTC),
// 			expected:  7,
// 		},
// 		// The following test cases demonstrate some possible edge cases
// 		{
// 			NotBefore: time.Date(2022, 1, 18, 10, 48, 0, 0, time.UTC),
// 			NotAfter:  time.Date(2037, 1, 15, 10, 48, 0, 0, time.UTC),
// 			expected:  5475,
// 		},
// 		{
// 			NotBefore: time.Date(2021, 1, 18, 10, 48, 0, 0, time.UTC),
// 			NotAfter:  time.Date(2022, 1, 17, 10, 48, 0, 0, time.UTC),
// 			expected:  365,
// 		},
// 		{
// 			NotBefore: time.Time{},
// 			NotAfter:  time.Time{},
// 			expected:  7,
// 		},
// 	}

// 	for i, testCase := range testCases {
// 		cert := &x509.Certificate{
// 			NotBefore: testCase.NotBefore,
// 			NotAfter:  testCase.NotAfter,
// 		}
// 		days := calculateRequestedValidityFromCertificate(cert)
// 		if days != testCase.expected {
// 			t.Errorf("expected %d days got %d for %d", testCase.expected, days, i)
// 		}
// 	}
// }

func testAccCheckCloudflareOriginCACertificateDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_origin_ca_certificate" {
			continue
		}

		cert, err := client.GetOriginCACertificate(context.Background(), rs.Primary.ID)
		if err == nil && cert.RevokedAt == (time.Time{}) {
			return fmt.Errorf("Origin CA Certificate still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckCloudflareOriginCACertificateExists(name string, cert *cloudflare.OriginCACertificate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Origin CA Certificate ID is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundOriginCACertificate, err := client.GetOriginCACertificate(context.Background(), rs.Primary.ID)
		if err != nil {
			return err
		}

		*cert = *foundOriginCACertificate
		return nil
	}
}

func testAccCheckCloudflareOriginCACertificateAttributes(zone string, cert *cloudflare.OriginCACertificate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		actual := schema.NewSet(schema.HashString, []interface{}{})
		for _, h := range cert.Hostnames {
			actual.Add(h)
		}
		expected := schema.NewSet(schema.HashString, []interface{}{zone, fmt.Sprintf("*.%s", zone)})
		if actual.Difference(expected).Len() > 0 {
			return fmt.Errorf("incorrect hostnames: expected %v, got %v", expected, actual)
		}

		block, _ := pem.Decode([]byte(cert.Certificate))
		if block == nil {
			return fmt.Errorf("bad certificate: %s", cert.Certificate)
		}

		_, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return err
		}

		if !cert.ExpiresOn.After(time.Now()) {
			return fmt.Errorf("expiration date of new cert is in the past: %s", cert.ExpiresOn.Format(time.RFC3339))
		}

		return nil
	}
}

func generateCSR(zone string) (string, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	template := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName: zone,
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, key)
	if err != nil {
		return "", err
	}

	csrPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
	return string(csrPem), nil
}

func testAccCheckCloudflareOriginCACertificateConfigBasic(name string, zoneName, csr string) string {
	return acctest.LoadTestCase("origincacertificateconfigbasic.tf", name, zoneName, csr)
}
