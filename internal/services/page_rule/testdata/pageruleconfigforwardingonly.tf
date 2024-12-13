
resource "cloudflare_page_rule" "%[3]s" {
  zone_id = "%[1]s"
  target  = "%[2]s"
  actions = {
    // on/off options cannot even be set to off without causing error
    forwarding_url = {
      url         = "http://%[4]s/forward"
      status_code = 301
    }
  }
}
