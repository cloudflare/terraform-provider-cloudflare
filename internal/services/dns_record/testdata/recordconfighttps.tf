
resource "cloudflare_dns_record" "%[2]s" {
	zone_id = "%[1]s"
	name = "%[2]s"
	type = "HTTPS"
	data = {
  priority = "1"
		target   = "."
		value    = "alpn=\"h2\""
}
	ttl = 300
}
