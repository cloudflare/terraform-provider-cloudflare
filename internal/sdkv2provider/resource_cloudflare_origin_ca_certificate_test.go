package sdkv2provider

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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareOriginCACertificate_Basic(t *testing.T) {
	var cert cloudflare.OriginCACertificate
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	name := "cloudflare_origin_ca_certificate." + rnd

	csr, err := generateCSR(zoneName)
	if err != nil {
		t.Errorf("unable to generate CSR: %v", err)
		return
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareOriginCACertificateDestroy,
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
			{
				ResourceName: name,
				ImportState:  true,
			},
		},
	})
}

func TestCalculateRequestedValidityFromCertificate(t *testing.T) {
	testCases := []struct {
		NotBefore time.Time
		NotAfter  time.Time
		expected  int
	}{
		{
			NotBefore: time.Date(2022, 1, 18, 10, 48, 0, 0, time.UTC),
			NotAfter:  time.Date(2023, 1, 18, 10, 48, 0, 0, time.UTC),
			expected:  365,
		},
		{
			NotBefore: time.Date(2022, 1, 18, 10, 48, 0, 0, time.UTC),
			NotAfter:  time.Date(2022, 1, 25, 10, 48, 0, 0, time.UTC),
			expected:  7,
		},
		// The following test cases demonstrate some possible edge cases
		{
			NotBefore: time.Date(2022, 1, 18, 10, 48, 0, 0, time.UTC),
			NotAfter:  time.Date(2037, 1, 15, 10, 48, 0, 0, time.UTC),
			expected:  5475,
		},
		{
			NotBefore: time.Date(2021, 1, 18, 10, 48, 0, 0, time.UTC),
			NotAfter:  time.Date(2022, 1, 17, 10, 48, 0, 0, time.UTC),
			expected:  365,
		},
		{
			NotBefore: time.Time{},
			NotAfter:  time.Time{},
			expected:  7,
		},
	}

	for i, testCase := range testCases {
		cert := &x509.Certificate{
			NotBefore: testCase.NotBefore,
			NotAfter:  testCase.NotAfter,
		}
		days := calculateRequestedValidityFromCertificate(cert)
		if days != testCase.expected {
			t.Errorf("expected %d days got %d for %d", testCase.expected, days, i)
		}
	}
}

func testAccCheckCloudflareOriginCACertificateDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

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

		client := testAccProvider.Meta().(*cloudflare.API)
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
	return fmt.Sprintf(`
resource "cloudflare_origin_ca_certificate" "%[1]s" {
	csr                = <<EOT
%[3]sEOT
	hostnames          = [ "%[2]s", "*.%[2]s" ]
	request_type       = "origin-rsa"
	requested_validity = 7
}
`, name, zoneName, csr)
}
