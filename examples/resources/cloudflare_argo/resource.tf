resource "cloudflare_argo" "example" {
  zone_id        = "d41d8cd98f00b204e9800998ecf8427e"
  tiered_caching = "on"
  smart_routing  = "on"
}
