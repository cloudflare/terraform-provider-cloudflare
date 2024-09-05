resource "cloudflare_total_tls" "example" {
  zone_id               = "0da42c8d2132a9ddaf714f9e7c920711"
  enabled               = true
  certificate_authority = "lets_encrypt"
}
