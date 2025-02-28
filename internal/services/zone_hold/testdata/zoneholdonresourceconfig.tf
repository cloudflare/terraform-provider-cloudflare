
resource "cloudflare_zone_hold" "%s" {
	zone_id = "%s"
	hold = true
}