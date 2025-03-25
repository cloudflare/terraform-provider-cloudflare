
  resource "cloudflare_hyperdrive_config" "%[1]s" {
		account_id = "%[2]s"
		name       = "%[3]s"
		origin     = {
			password             = "%[4]s"
			database             = "%[5]s"
			host                 = "%[6]s"
			scheme               = "%[7]s"
			user                 = "%[8]s"
			access_client_id     = "%[9]s"
			access_client_secret = "%[10]s"
		}
		caching = {
			disabled               = %[11]t
		}
	}
