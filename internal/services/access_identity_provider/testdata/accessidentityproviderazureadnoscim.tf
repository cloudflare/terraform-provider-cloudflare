
resource "cloudflare_access_identity_provider" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	type       = "azureAD"
	config = {
  client_id      = "test"
		client_secret  = "test"
		directory_id   = "directory"
		support_groups = true
}
}