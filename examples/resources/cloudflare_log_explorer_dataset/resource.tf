resource "cloudflare_log_explorer_dataset" "example_log_explorer_dataset" {
  dataset               = "http_requests"
  zone_id               = "zone_id"
  filter_attack_traffic = false
  fields = [
    {
      enabled = true
      name    = "edgeendtimestamp"
    },
    {
      enabled = true
      name    = "clientip"
    }
  ]
}
