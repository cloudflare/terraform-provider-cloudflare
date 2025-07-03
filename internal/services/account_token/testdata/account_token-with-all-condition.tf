
resource "cloudflare_account_token" "%[1]s" {
	name = "%[1]s"
  	account_id = "%[2]s"

	policies = [{
		effect = "allow"
		permission_groups = [{
	    	id = "%[3]s"
		}]
		resources = {
			"com.cloudflare.api.account.%[2]s" = "*"
		}
	}]

	condition = {
		request_ip = {
				in     = ["192.0.2.1/32"]
				not_in = ["198.51.100.1/32"]
			}
	}
}
