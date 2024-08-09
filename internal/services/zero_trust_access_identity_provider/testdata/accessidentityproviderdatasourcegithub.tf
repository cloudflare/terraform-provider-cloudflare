
	resource "cloudflare_access_identity_provider" "%[1]s" {
		account_id = "%[2]s"
	  name = "%[1]s"
	  type = "github"
	  config = {
  client_id = "test"
		client_secret = "secret"
}
	}

	data "cloudflare_access_identity_provider" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		depends_on = [cloudflare_access_identity_provider.%[1]s]
	}
	