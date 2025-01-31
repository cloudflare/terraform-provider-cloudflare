
  resource "cloudflare_byo_ip_prefix" "%[3]s" {
	  prefix_id = "%[1]s"
	  description = "%[2]s"
  }