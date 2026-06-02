resource "cloudflare_google_tag_gateway" "%[2]s" {
  zone_id          = "%[1]s"
  enabled          = false
  endpoint         = "/metrics"
  hide_original_ip = false
  measurement_id   = "G-XXXXXXXXXX"
  set_up_tag       = false
}
