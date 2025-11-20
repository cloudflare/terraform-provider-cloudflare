resource "cloudflare_queue" "%[1]s" {
  account_id = "%[2]s"
  queue_name = "%[1]s-queue-fedramp"
}

resource "cloudflare_r2_bucket" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  jurisdiction = "fedramp"
}

resource "cloudflare_r2_bucket_event_notification" "%[1]s" {
  account_id   = "%[2]s"
  bucket_name  = cloudflare_r2_bucket.%[1]s.name
  queue_id     = cloudflare_queue.%[1]s.queue_id
  jurisdiction = "fedramp"

  rules = [{
    actions     = ["PutObject", "DeleteObject"]
    description = "FedRAMP jurisdiction event notifications"
    prefix      = "gdpr/"
  }]
}
