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
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareOriginCACertificateDataSource(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	csr, err := generateCSR(zoneName)
	if err != nil {
		t.Errorf("unable to generate CSR: %v", err)
		return
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareOriginCACertificateDataSource_Basic(rnd, zoneName, csr),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_origin_ca_certificate."+rnd, "request_type", "origin-rsa"),
					resource.TestCheckResourceAttr("data.cloudflare_origin_ca_certificate."+rnd, "hostnames.#", "2"),
					resource.TestCheckResourceAttrSet("data.cloudflare_origin_ca_certificate."+rnd, "certificate"),
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

func testAccCheckCloudflareOriginCACertificateDataSource_Basic(rnd, zoneName, csr string) string {
	return fmt.Sprintf(`
resource "cloudflare_origin_ca_certificate" "%[1]s" {
	csr                = <<EOT
%[3]sEOT
	hostnames          = [ "%[2]s", "*.%[2]s" ]
	request_type       = "origin-rsa"
	requested_validity = 7
}

data "cloudflare_origin_ca_certificate" "%[1]s" {
	id = cloudflare_origin_ca_certificate.%[1]s.id
}
`, rnd, zoneName, csr)
}
