
	resource "cloudflare_dns_record" "%[1]s" {
		zone_id  = "%[2]s"
		type     = "MX"
		name     = "%[1]s"
		content    = "."
		priority = 0
		ttl = 300
	  }
