resource "cloudflare_logpush_job" "%s" {
  account_id      = "%s"
  dataset         = "http_requests"
  destination_conf = "s3://terraform-provider-cloudflare-test-bucket/logpush?region=us-west-2"
  enabled         = true
  name            = "%s"

  output_options {
    batch_prefix      = "batch-"
    batch_suffix      = "-end"
    cve20214428       = true
    field_delimiter   = ","
    field_names       = ["ClientIP", "ClientRequestHost", "ClientRequestMethod"]
    output_type       = "ndjson"
    record_delimiter  = "\n"
    record_prefix     = "{"
    record_suffix     = "}\n"
    sample_rate       = 0.5
    timestamp_format  = "rfc3339"
  }
}
