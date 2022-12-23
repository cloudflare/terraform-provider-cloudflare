resource "cloudflare_tiered_cache" "example" {
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
  cache_type = "smart"
}
