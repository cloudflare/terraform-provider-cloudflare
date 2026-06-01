resource "cloudflare_google_tag_gateway" "%[2]s" {
  zone_id          = "%[1]s"
  enabled          = false
  endpoint         = "/newendpoint"
  hide_original_ip = true
  measurement_id   = "G-YYYYYYYYYY"
  set_up_tag       = true
}
