package main

import (
	"testing"
)

func TestZeroTrustAccessMTLSCertificateTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "v4 certificate resource no transformation needed",
			Config: `
resource "cloudflare_access_mutual_tls_certificate" "example" {
  account_id           = "test-account-id"
  name                 = "test-cert"
  certificate          = file("cert.pem")
  associated_hostnames = ["example.com", "test.example.com"]
}`,
			Expected: []string{`
resource "cloudflare_access_mutual_tls_certificate" "example" {
  account_id           = "test-account-id"
  name                 = "test-cert"
  certificate          = file("cert.pem")
  associated_hostnames = ["example.com", "test.example.com"]
}`},
		},
		{
			Name: "v5 certificate resource remains unchanged",
			Config: `
resource "cloudflare_zero_trust_access_mtls_certificate" "example" {
  account_id           = "test-account-id"
  name                 = "test-cert"
  certificate          = file("cert.pem")
  associated_hostnames = ["example.com"]
}`,
			Expected: []string{`
resource "cloudflare_zero_trust_access_mtls_certificate" "example" {
  account_id           = "test-account-id"
  name                 = "test-cert"
  certificate          = file("cert.pem")
  associated_hostnames = ["example.com"]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}