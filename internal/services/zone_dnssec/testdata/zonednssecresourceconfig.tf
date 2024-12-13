resource "cloudflare_zone_dnssec" "%s" {
	zone_id = "%s"
	status = "active"
}