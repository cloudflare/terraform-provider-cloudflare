
resource "cloudflare_page_rule" "%[3]s" {
	zone_id = "%[1]s"
	target = "%[2]s/updated"
	actions = [{
		browser_check = "on"
		ssl = "strict"
		rocket_loader = "on"
	}]
}