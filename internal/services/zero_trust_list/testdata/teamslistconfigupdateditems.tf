resource "cloudflare_zero_trust_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "Updated list for testing"
	type        = "DOMAIN"
	items = [
		{ value = "updated-example.com", description = "Updated domain" },
		{ value = "new-test.com", description = "New domain" }
	]
}