resource "cloudflare_zone_dnssec" "%s" {
	zone_id            = "%s"
	status             = "active"
	dnssec_use_nsec3   = true
}