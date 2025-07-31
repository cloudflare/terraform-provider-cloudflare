resource "cloudflare_logpush_job" "%s" {
  account_id       = "%s"
  dataset          = "%s"
  destination_conf = "%s"
  output_options = {
    output_type      = "%s"
    sample_rate      = %f
    timestamp_format = "%s"
  }
}
