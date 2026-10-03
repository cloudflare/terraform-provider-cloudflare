resource "cloudflare_zone_tracing" "example_zone_tracing" {
  zone_id = "zone_id"
  destinations = ["x"]
  enabled = true
  forward_context = true
  persist = true
  propagation_policy = "accept"
  sampling_ratio = 0
}
