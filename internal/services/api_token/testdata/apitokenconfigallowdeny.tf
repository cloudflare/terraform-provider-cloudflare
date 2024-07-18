
		resource "cloudflare_api_token" "%[1]s" {
		  name = "%[1]s"

		  policy = [{
			effect = "allow"
			permission_groups = [
			  "%[2]s",
			]
			resources = {
			  "com.cloudflare.api.account.zone.*" = "*"
			}
		  }]
		  %[3]s
		}
		