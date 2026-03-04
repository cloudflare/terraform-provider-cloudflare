resource "cloudflare_queue" "%[1]s" {
  account_id = "%[2]s"
  queue_name = "tf-acc-test-queue-settings-%[1]s"
}
