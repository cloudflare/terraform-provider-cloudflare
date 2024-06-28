resource "cloudflare_custom_ssl" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  custom_ssl_options {
    certificate      = "-----INSERT CERTIFICATE-----"
    private_key      = "-----INSERT PRIVATE KEY-----"
    bundle_method    = "ubiquitous"
    geo_restrictions = "us"
    type             = "legacy_custom"
  }
}
