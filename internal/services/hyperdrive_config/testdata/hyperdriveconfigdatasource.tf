resource "cloudflare_hyperdrive_config" "%[1]s" {
	account_id = "%[2]s"
	name       = "%[3]s"
	origin     = {
		password = "%[4]s"
		database = "%[5]s"
		host     = "%[6]s"
		port     = %[7]d
		scheme   = "%[8]s"
		user     = "%[9]s"
	}
	caching = {
		disabled = %[10]t
	}
}

data "cloudflare_hyperdrive_config" "%[1]s" {
	account_id   = "%[2]s"
	hyperdrive_id = cloudflare_hyperdrive_config.%[1]s.id
}
