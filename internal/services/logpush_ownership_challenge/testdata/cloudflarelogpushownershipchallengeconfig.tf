
		resource "cloudflare_logpush_ownership_challenge" "%[1]s" {
		  zone_id = "%[2]s"
		  destination_conf = "%[3]s"
		}
		