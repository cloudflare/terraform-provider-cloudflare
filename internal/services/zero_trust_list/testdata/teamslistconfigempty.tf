resource "cloudflare_zero_trust_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "Empty list for testing"
	type        = "DOMAIN"
}