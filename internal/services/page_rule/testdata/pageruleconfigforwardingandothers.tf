
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions =[ {
		disable_security = true
		forwarding_url =[ {
			url = "http://%s/forward"
			status_code = 301
		}]
	}]
}