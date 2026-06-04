resource "cloudflare_google_tag_gateway" "example_google_tag_gateway" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  enabled = true
  endpoint = "/metrics"
  hide_original_ip = true
  measurement_id = "GTM-P2F3N47Q"
  set_up_tag = true
}
