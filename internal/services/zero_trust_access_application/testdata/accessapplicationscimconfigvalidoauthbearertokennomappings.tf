
resource "cloudflare_zero_trust_access_identity_provider" "%[1]s" {
	account_id = "%[2]s"
	name       = "%[1]s"
	type       = "azureAD"
	config = {
  client_id      = "test"
		client_secret  = "test"
		directory_id   = "directory"
		support_groups = true
}
	scim_config = {
  enabled                  = true
		group_member_deprovision = true
		seat_deprovision         = true
		user_deprovision         = true
}
}

resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  type             = "self_hosted"
  session_duration = "24h"
  domain = "%[1]s.%[3]s"
  scim_config = {
  enabled = false
	remote_uri = "scim2.com"
	idp_uid = cloudflare_zero_trust_access_identity_provider.%[1]s.id
	deactivate_on_delete = false
	authentication = {
		scheme =  "oauthbearertoken"
		token = "12345"
}
  }
}
