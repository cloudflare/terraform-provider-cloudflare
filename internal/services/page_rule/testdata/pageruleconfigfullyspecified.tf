
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s"
	actions =[ {
		browser_check = "on"
		browser_cache_ttl = 0
		email_obfuscation = "on"
		ip_geolocation = "on"
		server_side_exclude = "on"
		disable_apps = true
		disable_performance = true
		disable_security = true
		cache_level = "bypass"
		security_level = "essentially_off"
		ssl = "flexible"
	}]
}