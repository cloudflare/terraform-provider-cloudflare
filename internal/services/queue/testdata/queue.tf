
resource "cloudflare_queue" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
}