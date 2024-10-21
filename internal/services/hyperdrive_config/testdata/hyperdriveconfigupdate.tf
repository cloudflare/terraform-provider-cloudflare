  resource "cloudflare_hyperdrive_config" "%[1]s" {
		id         = "%[2]s"
		account_id = "%[3]s"
		name       = "%[4]s"
		origin     = {
			password = "%[5]s"
			database = "%[6]s"
			host     = "%[7]s"
			port     = "%[8]s"
			scheme   = "%[9]s"
			user     = "%[10]s"
		}
		caching = {
			disabled = %[11]s
		}
	}
