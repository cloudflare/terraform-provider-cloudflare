resource "cloudflare_origin_ca_certificate" "%s" {
  csr          = <<EOT
-----BEGIN CERTIFICATE REQUEST-----
MIICvDCCAaQCAQAwdzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWEx
FjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xGTAXBgNVBAoMEEV4YW1wbGUgQ29tcGFu
eTEgMB4GA1UEAwwXZXhhbXBsZS5jbG91ZGZsYXJlLmNvbTCCASIwDQYJKoZIhvcN
AQEBBQADggEPADCCAQoCggEBAKlwDWtJRLmvJc1BQPY6qDfKqH0Vl71jGZYs3t5b
-----END CERTIFICATE REQUEST-----
EOT
  request_type = "origin-rsa"
  hostnames    = ["example.cloudflare.com"]

  # requested_validity will default to 5475
}
