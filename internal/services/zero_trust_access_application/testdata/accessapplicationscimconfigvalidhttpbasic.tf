
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
  enabled = true
	remote_uri = "scim.com"
	idp_uid = cloudflare_zero_trust_access_identity_provider.%[1]s.id
	deactivate_on_delete = true
	authentication = {
		scheme =  "httpbasic"
		user = "test"
		password = "12345"
}
	mappings =[ {
		schema = "urn:ietf:params:scim:schemas:core:2.0:User"
		enabled = true
		filter = "title pr or userType eq \"Intern\""
		transform_jsonata = "$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"
		operations =[ {
			create = true
			update = true
			delete = true
		}]
	}]
  }
}
