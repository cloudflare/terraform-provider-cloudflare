resource "cloudflare_keyless_certificate" "example" {
  zone_id       = "0da42c8d2132a9ddaf714f9e7c920711"
  bundle_method = "ubiquitous"
  name          = "example.com Keyless SSL"
  host          = "example.com"
  port          = 24008
  enabled       = true
  certificate   = "-----INSERT CERTIFICATE-----"
}