
resource "cloudflare_zero_trust_access_identity_provider" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	type       = "azureAD"
	config = {
  client_id      = "test2"
		client_secret  = "test2"
		directory_id   = "directory"
		support_groups = true
}
	scim_config = {
  enabled                  = true
		group_member_deprovision = true
		seat_deprovision         = false
		user_deprovision         = true
}
}