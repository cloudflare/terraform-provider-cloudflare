
resource "cloudflare_page_rule" "%[3]s" {
  zone_id = "%[1]s"
  target  = "%[3]s"
  actions = {
    cache_key_fields = {
      host = {
        resolved = true
      }
      # `ignore` was a v4-only attribute, dropped in v5. `ignore = true` meant
      # "exclude every query string parameter", which v5 spells as ["*"].
      query_string = {
        exclude = ["*"]
      }
      user = {
        device_type = true
        geo         = false
        lang        = false
      }
    }
  }
}
