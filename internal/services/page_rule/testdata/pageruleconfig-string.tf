resource "cloudflare_page_rule" "%[1]s" {
	zone_id = "%[2]s"
	target  = "%[3]s"
	actions = {
	    %[4]s = "%[5]s"
	}
}
