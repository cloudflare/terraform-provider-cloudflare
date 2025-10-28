resource "cloudflare_queue" "%[1]s" {
	account_id = "%[2]s"
	queue_name = "%[3]s"

	settings = {
		message_retention_period = 30
	}
}
