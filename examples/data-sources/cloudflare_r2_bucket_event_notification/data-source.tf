data "cloudflare_r2_bucket_event_notification" "example_r2_bucket_event_notification" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  bucket_name = "example-bucket"
  queue_id = "queue_id"
}
