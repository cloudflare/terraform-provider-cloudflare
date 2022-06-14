resource "cloudflare_access_mutual_tls_certificate" "my_cert" {
  zone_id              = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name                 = "My Root Cert"
  certificate          = var.ca_pem
  associated_hostnames = ["staging.example.com"]
}
