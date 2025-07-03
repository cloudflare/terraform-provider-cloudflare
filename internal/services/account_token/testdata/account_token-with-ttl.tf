resource "cloudflare_account_token" "%[1]s" {
	name = "%[1]s"
  	account_id = "%[2]s"

	policies = [{
		effect = "allow"
		permission_groups = [{ id = "%[3]s" }]
		resources = {
			"com.cloudflare.api.account.%[2]s" = "*"
		}
	}]

	not_before = "2018-07-01T05:20:00Z"
	expires_on = "2032-01-01T00:00:00Z"
}
