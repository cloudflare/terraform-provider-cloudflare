resource "cloudflare_argo" "example" {
  zone_id        = "0da42c8d2132a9ddaf714f9e7c920711"
  tiered_caching = "on"
  smart_routing  = "on"
}
