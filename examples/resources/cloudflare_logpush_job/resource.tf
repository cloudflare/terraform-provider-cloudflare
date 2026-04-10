resource "cloudflare_logpush_job" "example_logpush_job" {
  destination_conf = "s3://mybucket/logs?region=us-west-2"
  zone_id = "zone_id"
  dataset = "gateway_dns"
  enabled = false
  filter = "{\"where\":{\"and\":[{\"key\":\"ClientRequestPath\",\"operator\":\"contains\",\"value\":\"/static\"},{\"key\":\"ClientRequestHost\",\"operator\":\"eq\",\"value\":\"example.com\"}]}}"
  frequency = "high"
  kind = ""
  logpull_options = "fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"
  max_upload_bytes = 5000000
  max_upload_interval_seconds = 30
  max_upload_records = 1000
  name = "example.com"
  output_options = {
    batch_prefix = ""
    batch_suffix = ""
    cve_2021_44228 = false
    field_delimiter = ","
    field_names = ["Datetime", "DstIP", "SrcIP"]
    output_type = "ndjson"
    record_delimiter = ""
    record_prefix = "{"
    record_suffix = <<EOT
    }

    EOT
    record_template = "record_template"
    sample_rate = 1
    timestamp_format = "unixnano"
  }
  ownership_challenge = "00000000000000000000"
}
