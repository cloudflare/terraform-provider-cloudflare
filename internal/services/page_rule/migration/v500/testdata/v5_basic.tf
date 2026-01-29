resource "cloudflare_page_rule" "%s" {
  zone_id  = "%s"
  target   = "%s.example.com/*"
  status   = "active"
  
  actions = {
    cache_level = "bypass"
  }
}
