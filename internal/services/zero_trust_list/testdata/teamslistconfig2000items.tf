
resource "cloudflare_zero_trust_list" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "Large list for no-change plan performance test"
	type        = "SERIAL"
	items       = [%[3]s]
}
