resource "cloudflare_queue" "test_queue" {
  account_id = "%[1]s"
  queue_name = "%[2]s"
}

# This config has NO lifecycle ignore_changes to properly test the consumer_id bug
resource "cloudflare_queue_consumer" "%[3]s" {
  account_id = "%[4]s"
  queue_id   = cloudflare_queue.test_queue.id
  type       = "http_pull"
}
