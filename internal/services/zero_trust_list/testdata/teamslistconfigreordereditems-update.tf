
resource "cloudflare_zero_trust_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "My description update"
	type        = "SERIAL"
	items = [{ value = "csdf-5678"},{ value = "asdf-1234"}, { value = "bsdf-5678"}]
}
