
		resource "cloudflare_zero_trust_tunnel_cloudflared" "%[1]s" {
		  account_id = "%[2]s"
		  name       = "%[1]s"
		  tunnel_secret     = "%[3]s"
		}

		resource "cloudflare_zero_trust_tunnel_cloudflared_config" "%[1]s" {
		  account_id         = "%[2]s"
		  tunnel_id          = cloudflare_zero_trust_tunnel_cloudflared.%[1]s.id

		  config = {
			ingress = [{
				service = "https://10.0.0.1:8081"
			  }]
		  }
		}
