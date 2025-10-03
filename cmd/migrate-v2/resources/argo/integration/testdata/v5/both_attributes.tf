resource "cloudflare_argo_smart_routing" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  value   = "on"
}

resource "cloudflare_tiered_cache" "example_tiered" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  value   = "on"
}

moved {
  from = cloudflare_argo.example
  to   = cloudflare_argo_smart_routing.example
}

moved {
  from = cloudflare_argo.example
  to   = cloudflare_tiered_cache.example_tiered
}
