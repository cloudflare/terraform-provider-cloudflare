resource "cloudflare_custom_hostname" "%[2]s" {
	zone_id = "%[1]s"
	hostname = "%[2]s.%[3]s"
	ssl = {
		method = "txt"
		type = "dv"
		wildcard = true
	}
	custom_metadata = {
		"customer_id" = 12345
		"redirect_to_https" = true
		"security_tag" = "low"
	}
}
