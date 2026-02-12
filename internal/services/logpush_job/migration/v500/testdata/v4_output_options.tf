resource "cloudflare_logpush_job" "%[1]s" {
  account_id       = "%[2]s"
  dataset          = "audit_logs"
  destination_conf = "https://logpush-receiver.sd.cfplat.com"
  enabled          = true
  name             = "%[3]s"

  output_options {
    batch_prefix     = "batch-start-"
    batch_suffix     = "-batch-end"
    cve20214428      = true
    field_delimiter  = "|"
    field_names      = ["ClientIP", "EdgeStartTimestamp", "ClientRequestMethod"]
    output_type      = "ndjson"
    record_delimiter = "\\n"
    record_prefix    = "["
    record_suffix    = "]\\n"
    record_template  = "{{.ClientIP}}"
    sample_rate      = 0.5
    timestamp_format = "rfc3339"
  }
}
