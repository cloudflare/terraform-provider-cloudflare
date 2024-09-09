
	resource "cloudflare_record" "%[1]s" {
		zone_id  = "%[2]s"
		type     = "MX"
		name     = "%[1]s"
		value    = "."
		priority = 0
	  }
	