
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions = [{
		edge_cache_ttl = 10
	}]
}