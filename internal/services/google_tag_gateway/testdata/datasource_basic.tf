resource "cloudflare_google_tag_gateway" "%[2]s" {
  zone_id          = "%[1]s"
  enabled          = true
  endpoint         = "/gtm"
  hide_original_ip = true
  measurement_id   = "GTM-XXXXXXX"
  set_up_tag       = true
}

data "cloudflare_google_tag_gateway" "%[2]s" {
  zone_id    = cloudflare_google_tag_gateway.%[2]s.zone_id
  depends_on = [cloudflare_google_tag_gateway.%[2]s]
}
