
resource "cloudflare_zero_trust_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "My description"
	type        = "SERIAL"
	items = [{ value = "asdf-5678"}, { value = "asdf-1234"}]
}
