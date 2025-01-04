
resource "cloudflare_zone_hold" "%s" {
	zone_id = "%s"
	hold = false
	hold_after = "%s"
}