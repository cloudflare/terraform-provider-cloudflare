resource "cloudflare_access_mutual_tls_certificate" "my_cert" {
  zone_id              = "0da42c8d2132a9ddaf714f9e7c920711"
  name                 = "My Root Cert"
  certificate          = var.ca_pem
  associated_hostnames = ["staging.example.com"]
}
