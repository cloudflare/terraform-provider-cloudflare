
	resource "cloudflare_api_token" "%[1]s" {
		name = "%[1]s"

		policy = [{
			effect = "allow"
			permission_groups = [{
		    id = "%[2]s"
			}]
			resources = { "com.cloudflare.api.account.zone.*" = "*" }
		}]

		condition = {
      request_ip = {
				in = ["192.0.2.1/32"]
			}
		}
	}
