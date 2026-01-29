resource "cloudflare_page_rule" "%s" {
  zone_id  = "%s"
  target   = "%s.example.com/*"
  status   = "active"
  
  actions = {
    cache_key_fields = {
      host = {
        resolved = true
      }
      query_string = {
        exclude = ["utm_source"]
      }
      user = {
        device_type = true
        geo         = false
        lang        = false
      }
    }
  }
}
