
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions =[ {
		cache_key_fields =[ {
			header =[ {
				exclude = ["origin"]
			}]
			host =[ {}]
			query_string =[ {}]
			user =[ {
				device_type = true
				geo = true
			}]
		}]
	}]
}