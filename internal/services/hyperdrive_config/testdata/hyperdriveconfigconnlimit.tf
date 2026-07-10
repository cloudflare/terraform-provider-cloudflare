resource "cloudflare_hyperdrive_config" "%[1]s" {
	account_id              = "%[2]s"
	name                    = "%[3]s"
	origin_connection_limit = %[4]d
	origin = {
		password = "%[5]s"
		database = "%[6]s"
		host     = "%[7]s"
		port     = %[8]d
		scheme   = "%[9]s"
		user     = "%[10]s"
	}
	caching = {
		disabled = %[11]t
	}
}
