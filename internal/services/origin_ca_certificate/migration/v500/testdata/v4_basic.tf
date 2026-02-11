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
  hostnames    = ["example.cloudflare.com", "*.example.cloudflare.com"]

  # Optional: requested_validity as Int in v4 (will convert to Float64 in v5)
  requested_validity = 365

  # Optional: min_days_for_renewal exists in v4 but removed in v5 (will be dropped)
  min_days_for_renewal = 7
}
