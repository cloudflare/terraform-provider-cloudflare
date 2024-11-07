
	resource "cloudflare_api_token" "%[1]s" {
		name = "%[2]s"

		policy = [{
			effect = "allow"
			permission_groups = [{ id = "%[3]s" }]
			resources = { "com.cloudflare.api.account.zone.*" = "*" }
		}]
	}
