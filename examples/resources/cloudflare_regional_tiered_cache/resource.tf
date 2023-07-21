resource "cloudflare_regional_tiered_cache" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  value   = "on"
}
