resource "cloudflare_access_mutual_tls_hostname_settings" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  settings {
    hostname                      = "example.com"
    client_certificate_forwarding = true
    china_network                 = false
  }
}
