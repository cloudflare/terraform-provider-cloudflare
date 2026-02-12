resource "cloudflare_logpush_job" "%[1]s" {
  account_id       = "%[2]s"
  dataset          = "audit_logs"
  destination_conf = "https://logpush-receiver.sd.cfplat.com"
  enabled          = true
  name             = "%[3]s"

  output_options = {
    cve_2021_44228   = false
    field_delimiter  = ","
    record_prefix    = "{"
    record_suffix    = "}\n"
    timestamp_format = "unixnano"
    sample_rate      = 1.0
  }
}
