
resource "cloudflare_page_rule" "%[3]s" {
  zone_id = "%[1]s"
  target  = "%[3]s"
  actions = {
    cache_key_fields = {
      cookie = {
        check_presence = ["cookie_presence"]
        include        = ["cookie_include"]
      }
      header = {
        check_presence = ["header_presence"]
        include        = ["header_include"]
      }
      host = {
        resolved = true
      }
      query_string = {
        exclude = ["*"] 
      }
      user = {}
    }
  }
}
