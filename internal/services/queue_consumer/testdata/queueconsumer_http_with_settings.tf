resource "cloudflare_queue" "test_queue" {
  account_id = "%[1]s"
  queue_name = "%[2]s"
}

resource "cloudflare_queue_consumer" "%[3]s" {
  account_id = "%[4]s"
  queue_id   = cloudflare_queue.test_queue.id
  type       = "http_pull"

  settings = {
    batch_size            = 10
    max_retries           = 3
    retry_delay           = 0
    visibility_timeout_ms = 30000
  }
  lifecycle {
    ignore_changes = [
      settings
    ]
  }
}
