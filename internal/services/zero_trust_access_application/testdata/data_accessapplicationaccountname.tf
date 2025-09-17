
	resource "cloudflare_zero_trust_access_application" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		domain = "%[1]s.%[3]s"
		type = "self_hosted"
	}

	data "cloudflare_zero_trust_access_application" "%[1]s" {
  	account_id = "%[2]s"
    filter = {
     	name = "%[1]s"
  	}
		depends_on = [cloudflare_zero_trust_access_application.%[1]s]
	}
