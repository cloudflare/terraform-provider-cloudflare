resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
		effect = "allow"
		permission_groups = [{
		  id = "%[2]s"
		}]
		resources = {
		  "com.cloudflare.api.account.zone.*" = "*"
		}
  }]
}
