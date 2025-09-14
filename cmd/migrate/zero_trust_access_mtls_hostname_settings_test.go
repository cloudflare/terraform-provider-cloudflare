package main

import (
	"testing"
)

func TestZeroTrustAccessMTLSHostnameSettingsTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "transform single settings block",
			Config: `
resource "cloudflare_zero_trust_access_mtls_hostname_settings" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  settings {
    hostname = "example.com"
    client_certificate_forwarding = true
    china_network = false
  }
}`,
			Expected: []string{`
resource "cloudflare_zero_trust_access_mtls_hostname_settings" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  settings = [{
    china_network                 = false
    client_certificate_forwarding = true
    hostname                      = "example.com"
  }]
}`},
		},
		{
			Name: "transform multiple settings blocks",
			Config: `
resource "cloudflare_zero_trust_access_mtls_hostname_settings" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  settings {
    hostname = "app1.example.com"
    client_certificate_forwarding = true
    china_network = false
  }
  settings {
    hostname = "app2.example.com"
    client_certificate_forwarding = false
    china_network = false
  }
}`,
			Expected: []string{`
resource "cloudflare_zero_trust_access_mtls_hostname_settings" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  settings = [{
    china_network                 = false
    client_certificate_forwarding = true
    hostname                      = "app1.example.com"
  }, {
    china_network                 = false
    client_certificate_forwarding = false
    hostname                      = "app2.example.com"
  }]
}`},
		},
		{
			Name: "transform with boolean defaults",
			Config: `
resource "cloudflare_zero_trust_access_mtls_hostname_settings" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  settings {
    hostname = "example.com"
  }
}`,
			Expected: []string{`
resource "cloudflare_zero_trust_access_mtls_hostname_settings" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  settings = [{
    china_network                 = false
    client_certificate_forwarding = false
    hostname                      = "example.com"
  }]
}`},
		},
		{
			Name: "leave other resources unchanged",
			Config: `
resource "cloudflare_zero_trust_access_policy" "example" {
  name = "example"
  include {
    everyone = true
  }
}`,
			Expected: []string{`
resource "cloudflare_zero_trust_access_policy" "example" {
  name = "example"
  include {
    everyone = true
  }
}`},
		},
		{
			Name: "dynamic settings blocks converted to for expression",
			Config: `
locals {
  mtls_domains = ["app1.example.com", "app2.example.com"]
}

resource "cloudflare_zero_trust_access_mtls_hostname_settings" "test" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  
  dynamic "settings" {
    for_each = local.mtls_domains
    content {
      hostname = settings.value
      client_certificate_forwarding = true
      china_network = false
    }
  }
}`,
			Expected: []string{`
locals {
  mtls_domains = ["app1.example.com", "app2.example.com"]
}

resource "cloudflare_zero_trust_access_mtls_hostname_settings" "test" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  settings = [for value in local.mtls_domains : {
    hostname                      = value
    china_network                 = false
    client_certificate_forwarding = true
  }]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileWithoutImports)
}

func TestAccessMutualTLSHostnameSettingsTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "transform cloudflare_access_mutual_tls_hostname_settings to v5 format",
			Config: `
resource "cloudflare_access_mutual_tls_hostname_settings" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  settings {
    hostname = "example.com"
    client_certificate_forwarding = true
    china_network = false
  }
}`,
			Expected: []string{`
resource "cloudflare_access_mutual_tls_hostname_settings" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  settings = [{
    china_network                 = false
    client_certificate_forwarding = true
    hostname                      = "example.com"
  }]
}`},
		},
		{
			Name: "transform multiple settings blocks for old resource name",
			Config: `
resource "cloudflare_access_mutual_tls_hostname_settings" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  settings {
    hostname = "app1.example.com"
    client_certificate_forwarding = true
    china_network = false
  }
  settings {
    hostname = "app2.example.com"
    client_certificate_forwarding = false
    china_network = false
  }
}`,
			Expected: []string{`
resource "cloudflare_access_mutual_tls_hostname_settings" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  settings = [{
    china_network                 = false
    client_certificate_forwarding = true
    hostname                      = "app1.example.com"
  }, {
    china_network                 = false
    client_certificate_forwarding = false
    hostname                      = "app2.example.com"
  }]
}`},
		},
		{
			Name: "handle dynamic blocks for old resource name",
			Config: `
resource "cloudflare_access_mutual_tls_hostname_settings" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  
  dynamic "settings" {
    for_each = local.domains
    content {
      hostname = settings.value
      client_certificate_forwarding = true
      china_network = false
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_access_mutual_tls_hostname_settings" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  settings = [for value in local.domains : {
    hostname                      = value
    china_network                 = false
    client_certificate_forwarding = true
  }]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileWithoutImports)
}
