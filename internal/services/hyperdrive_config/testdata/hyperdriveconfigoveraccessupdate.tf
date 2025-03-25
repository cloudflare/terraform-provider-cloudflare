  resource "cloudflare_hyperdrive_config" "%[1]s" {
		account_id = "%[3]s"
		name       = "%[4]s"
		origin     = {
			password             = "%[5]s"
			database             = "%[6]s"
			host                 = "%[7]s"
			scheme               = "%[8]s"
			user                 = "%[9]s"
			access_client_id     = "%[10]s"
			access_client_secret = "%[11]s"
		}
		caching = {
			disabled               = %[12]t
		}
	}
