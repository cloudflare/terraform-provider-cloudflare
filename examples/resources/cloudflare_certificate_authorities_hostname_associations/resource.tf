# Hostname associations for the active Cloudflare Managed CA.
resource "cloudflare_certificate_authorities_hostname_associations" "example_1" {
  zone_id   = "0da42c8d2132a9ddaf714f9e7c920711"
  hostnames = ["example.com"]
}

# Hostname associations for a specific mTLS certificate.
resource "cloudflare_certificate_authorities_hostname_associations" "example_2" {
  zone_id             = "0da42c8d2132a9ddaf714f9e7c920711"
  mtls_certificate_id = "1fc1e34f39e74dd591366239dc47c5a4"
  hostnames           = ["example.com"]
}
