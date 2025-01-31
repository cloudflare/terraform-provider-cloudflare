
		resource "cloudflare_filter" "%[1]s" {
		  zone_id = "%[2]s"
		  paused = "%[3]s"
		  description = "%[4]s"
		  expression = "%[5]s"
		}
		resource "cloudflare_firewall_rule" "%[1]s" {
		  zone_id = "%[2]s"
		  paused = "%[3]s"
		  description = "%[4]s"
		  filter_id = "${cloudflare_filter.%[1]s.id}"
		  action = "bypass"
		  priority = "2"
		  products = ["uaBlock", "waf"]
		}
		