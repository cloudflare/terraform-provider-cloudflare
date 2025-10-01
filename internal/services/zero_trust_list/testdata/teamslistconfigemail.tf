resource "cloudflare_zero_trust_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "Email list for testing"
	type        = "EMAIL"
	items = [
		{ value = "test@example.com" },
		{ value = "admin@test.com", description = "Admin email" },
		{ value = "user@cloudflare.com", description = "User email" }
	]
}