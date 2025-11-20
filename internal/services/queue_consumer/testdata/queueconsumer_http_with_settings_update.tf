resource "cloudflare_queue" "test_queue" {
  account_id = "%s"
  queue_name = "%s"
}

resource "cloudflare_queue_consumer" "%s" {
  account_id = "%s"
  queue_id   = cloudflare_queue.test_queue.id
  type       = "http_pull"

  settings = {
    batch_size        = 20
    max_retries       = 5
  }
}
