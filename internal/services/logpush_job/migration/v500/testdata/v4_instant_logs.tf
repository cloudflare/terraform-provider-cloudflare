resource "cloudflare_logpush_job" "%s" {
  zone_id          = "%s"
  dataset          = "http_requests"
  destination_conf = "s3://terraform-provider-cloudflare-test-bucket/logpush?region=us-west-2"
  enabled          = true
  kind             = "instant-logs"
}
