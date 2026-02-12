resource "cloudflare_logpush_job" "%s" {
  account_id      = "%s"
  dataset         = "audit_logs"
  destination_conf = "s3://terraform-provider-cloudflare-test-bucket/logpush?region=us-west-2"
  enabled         = true
  name            = "%s"

  output_options = {
    cve_2021_44228   = false
    field_delimiter  = ","
    record_prefix    = "{"
    record_suffix    = "}\n"
    timestamp_format = "unixnano"
    sample_rate      = 1.0
  }
}
