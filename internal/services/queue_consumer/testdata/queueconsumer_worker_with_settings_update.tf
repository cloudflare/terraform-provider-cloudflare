resource "cloudflare_queue" "test_queue" {
  account_id = "%s"
  name       = "%s"
}

resource "cloudflare_queue_consumer" "%s" {
  account_id  = "%s"
  queue_id    = cloudflare_queue.test_queue.id
  type        = "worker"
  script_name = "test-worker"

  settings {
    batch_size        = 20
    max_retries       = 5
    max_wait_time_ms  = 8000
  }
}
