resource "cloudflare_page_rule" "%s" {
  zone_id  = "%s"
  target   = "%s.example.com/*"

  actions {
    cache_key_fields {
      host {
        resolved = true
      }
      query_string {
        exclude = ["utm_source"]
        ignore  = false
      }
      user {
        device_type = true
        geo         = false
        lang        = false
      }
    }
  }
}
