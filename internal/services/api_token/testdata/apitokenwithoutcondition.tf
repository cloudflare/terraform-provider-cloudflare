resource "cloudflare_api_token" "%[1]s" {
	name = "%[2]s"
	status = "active"

	policies = [{
		effect = "allow"
		permission_groups = [{ id = "%[3]s" }]
		resources = { "com.cloudflare.api.account.zone.*" = "*" }
	}]
}
