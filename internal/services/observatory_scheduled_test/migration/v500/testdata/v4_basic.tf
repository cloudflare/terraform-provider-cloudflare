resource "cloudflare_observatory_scheduled_test" "test" {
  zone_id   = "%s"
  url       = "%s"
  region    = "us-central1"
  frequency = "DAILY"
}
