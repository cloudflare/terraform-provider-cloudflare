
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[3]s"
	actions =[ {
    cache_key_fields =[ {
      host =[ {
        resolved = true
      }]
      query_string =[ {
        ignore = true
      }]
      user =[ {
        device_type = true
        geo         = false
        lang        = false
      }]
    }]
  }]
}