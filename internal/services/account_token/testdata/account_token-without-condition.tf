resource "cloudflare_account_token" "%[1]s" {
	name = "%[3]s"
  	account_id = "%[2]s"

	policies = [{
		effect = "allow"
		permission_groups = [{ id = "%[4]s" }]
		resources = {
			"com.cloudflare.api.account.%[2]s" = "*"
		}
	}]
}
