resource "cloudflare_queue" "test_queue" {
  account_id = "%s"
  name       = "%s"
}

resource "cloudflare_queue" "dlq1" {
  account_id = "%s"
  name       = "%s"
}

resource "cloudflare_queue_consumer" "%s" {
  account_id        = "%s"
  queue_id          = cloudflare_queue.test_queue.id
  type              = "http_pull"
  dead_letter_queue = cloudflare_queue.dlq1.name
}
