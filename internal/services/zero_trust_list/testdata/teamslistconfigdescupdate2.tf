resource "cloudflare_zero_trust_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "Testing description updates"
	type        = "SERIAL"
	items = [
		{ value = "device-001", description = "Added description" },
		{ value = "device-002" },
		{ value = "device-003", description = "Updated description" },
		{ value = "device-004", description = "New item with description" }
	]
}