resource "cloudflare_page_rule" "%s" {
  zone_id  = "%s"
  target   = "%s.example.com/*"

  actions {
    cache_level = "bypass"
  }
}
