
	resource "cloudflare_zero_trust_tunnel_warp_connector" "%[2]s" {
		account_id    = "%[1]s"
		name          = "%[2]s"
	}
