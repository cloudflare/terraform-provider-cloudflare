
	resource "cloudflare_api_token" "%[1]s" {
		name = "%[1]s"

		policy = [{
			effect = "allow"
			permission_groups = [ "%[2]s" ]
			resources = { "com.cloudflare.api.account.zone.*" = "*" }
		}]

		not_before = "2018-07-01T05:20:00Z"
		expires_on = "2032-01-01T00:00:00Z"
	}
