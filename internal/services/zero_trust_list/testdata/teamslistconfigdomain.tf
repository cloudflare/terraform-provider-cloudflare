resource "cloudflare_zero_trust_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "Domain list for testing"
	type        = "DOMAIN"
	items = [
		{ value = "example.com" },
		{ value = "test.com", description = "Test domain" },
		{ value = "cloudflare.com", description = "Main domain" }
	]
}