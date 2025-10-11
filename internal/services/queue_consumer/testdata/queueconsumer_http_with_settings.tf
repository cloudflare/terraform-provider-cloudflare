resource "cloudflare_queue" "test_queue" {
  account_id = "%s"
  name       = "%s"
}

resource "cloudflare_queue_consumer" "%s" {
  account_id = "%s"
  queue_id   = cloudflare_queue.test_queue.id
  type       = "http_pull"

  settings {
    batch_size        = 10
    max_retries       = 3
    max_wait_time_ms  = 5000
  }
}
