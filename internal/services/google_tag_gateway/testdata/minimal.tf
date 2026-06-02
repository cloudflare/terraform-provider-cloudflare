resource "cloudflare_google_tag_gateway" "%[2]s" {
  zone_id          = "%[1]s"
  enabled          = true
  endpoint         = "/analytics"
  hide_original_ip = false
  measurement_id   = "GTM-ABCDEFG"
}
