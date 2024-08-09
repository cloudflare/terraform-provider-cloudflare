
	  resource "cloudflare_workers_for_platforms_dispatch_namespace" "%[1]s" {
		account_id = "%[2]s"
		name       = "%[1]s"
	  }

	  resource "cloudflare_workers_script" "script_%[1]s" {
		account_id          = "%[2]s"
		name                = "script_%[1]s"
		content             = %[3]s
		module              = true
		compatibility_date  = "%[4]s"
		dispatch_namespace  = "%[1]s"
		tags                = %[5]q

		depends_on = [cloudflare_workers_for_platforms_dispatch_namespace.%[1]s]
	  }