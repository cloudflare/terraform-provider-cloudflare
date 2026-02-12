resource "cloudflare_logpush_job" "%s" {
  account_id      = "%s"
  dataset         = "audit_logs"
  destination_conf = "s3://terraform-provider-cloudflare-test-bucket/logpush?region=us-west-2"
  enabled         = true
  name            = "%s"
  filter          = ""
  logpull_options = ""

  output_options {
    cve20214428       = false
    field_delimiter   = ","
    record_prefix     = "{"
    record_suffix     = "}\n"
    timestamp_format  = "unixnano"
    sample_rate       = 1.0
  }
}
