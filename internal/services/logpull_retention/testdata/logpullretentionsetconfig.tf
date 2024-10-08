
  resource "cloudflare_logpull_retention" "%[1]s" {
    zone_id = "%[2]s"
	  enabled = "%[3]s"
  }