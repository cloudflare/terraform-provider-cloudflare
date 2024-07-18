
  resource "cloudflare_authenticated_origin_pulls" "%[2]s" {
	  zone_id        = "%[1]s"
	  enabled = true
  }