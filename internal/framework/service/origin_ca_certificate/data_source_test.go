package origin_ca_certificate_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareUserDataSource(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	csr, err := generateCSR(zoneName)
	if err != nil {
		t.Errorf("unable to generate CSR: %v", err)
		return
	}

	name := fmt.Sprintf("data.cloudflare_origin_ca_certificate.%s", zoneName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareOriginCACertificateDataSource_Basic(zoneName, csr),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "request_type", "origin-rsa"),
					resource.TestCheckResourceAttr(name, "requested_validity", "7"),
					resource.TestCheckResourceAttr(name, "hostnames.#", "1"),
					resource.TestCheckResourceAttr(name, "hostnames.0", "example.com"),
				),
			},
		},
	})
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

func testAccCheckCloudflareOriginCACertificateDataSource_Basic(zoneName string, csr string) string {
	return fmt.Sprintf(`
resource "cloudflare_origin_ca_certificate" "%[1]s" {
	csr                = <<EOT
%[2]sEOT
	hostnames          = [ "%[1]s", "*.%[1]s" ]
	request_type       = "origin-rsa"
	requested_validity = 7
}

data "cloudflare_origin_ca_root_certificate" "%[1]s" {
	id = cloudflare_origin_ca_certificate.%[1]s.id
}
`, zoneName, csr)
}
