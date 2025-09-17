
resource "cloudflare_workers_for_platforms_dispatch_namespace" "%[2]s" {
	account_id = "%[4]s"
	name       = "%[1]s"
}

resource "cloudflare_workers_script" "%[2]s" {
	account_id  = "%[4]s"
	script_name = "%[1]s"
	metadata = {
		main_module = "main.js"
	}
}

resource "cloudflare_workers_secret" "%[2]s" {
	account_id  = "%[4]s"
	script_name = cloudflare_workers_script.%[2]s.script_name
	dispatch_namespace = cloudflare_workers_for_platforms_dispatch_namespace.%[2]s.name
	name 		= "%[2]s"
	text	= "%[3]s"
	type = "secret_text"
	depends_on = [
		cloudflare_workers_for_platforms_dispatch_namespace.%[2]s,
		cloudflare_workers_script.%[2]s
	]
}