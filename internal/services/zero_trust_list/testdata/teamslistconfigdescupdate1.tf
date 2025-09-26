resource "cloudflare_zero_trust_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "Testing description updates"
	type        = "SERIAL"
	items = [
		{ value = "device-001" },
		{ value = "device-002", description = "Original description" },
		{ value = "device-003", description = "Will be updated" }
	]
}