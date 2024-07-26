
	resource "cloudflare_worker_script" "%[2]s" {
		account_id = "%[4]s"
		name       = "%[1]s"
		content    = "%[5]s"
	}

	resource "cloudflare_worker_secret" "%[2]s" {
		account_id  = "%[4]s"
		script_name = cloudflare_worker_script.%[2]s.name
		name 		= "%[2]s"
		secret_text	= "%[3]s"
	}