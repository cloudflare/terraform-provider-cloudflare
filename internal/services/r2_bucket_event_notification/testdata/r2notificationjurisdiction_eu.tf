resource "cloudflare_queue" "%[1]s" {
  account_id = "%[2]s"
  queue_name = "%[1]s-queue-eu"
}

resource "cloudflare_r2_bucket" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  jurisdiction = "eu"
}

resource "cloudflare_r2_bucket_event_notification" "%[1]s" {
  account_id   = "%[2]s"
  bucket_name  = cloudflare_r2_bucket.%[1]s.name
  queue_id     = cloudflare_queue.%[1]s.queue_id
  jurisdiction = "eu"

  rules = [{
    actions     = ["PutObject", "DeleteObject"]
    description = "EU jurisdiction event notifications"
    prefix      = "gdpr/"
  }]
}
