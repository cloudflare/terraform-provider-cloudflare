
resource "cloudflare_page_rule" "%[3]s" {
  zone_id = "%[1]s"
  target  = "%[2]s"
  actions = {
    cache_ttl_by_status = {}
  }
}
