resource "cloudflare_queue" "test_queue" {
  account_id = "%s"
  name       = "%s"
}

resource "cloudflare_queue_consumer" "%s" {
  account_id = "%s"
  queue_id   = cloudflare_queue.test_queue.id
  type       = "http_pull"
}