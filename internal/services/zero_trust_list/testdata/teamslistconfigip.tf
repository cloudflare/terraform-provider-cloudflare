resource "cloudflare_zero_trust_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "IP list for testing"
	type        = "IP"
	items = [
		{ value = "192.168.1.1" },
		{ value = "10.0.0.0/8", description = "Private network" },
		{ value = "203.0.113.0/24", description = "Test network" }
	]
}