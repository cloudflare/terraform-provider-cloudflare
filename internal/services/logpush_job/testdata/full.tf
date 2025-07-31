resource "cloudflare_logpush_job" "%s" {
  account_id                  = "%s"
  dataset                     = "%s"
  destination_conf            = "%s"
  enabled                     = %t
  name                        = "%s"
  filter                      = %q
  kind                        = "%s"
  max_upload_bytes            = %d
  max_upload_records          = %d
  max_upload_interval_seconds = %d
  frequency                   = "%s"
  logpull_options             = "%s"
  output_options = {
	  batch_prefix     = "%s"
	  batch_suffix     = "%s"
	  cve_2021_44228   = %t
	  field_delimiter  = "%s"
	  field_names      = %s
	  output_type      = "%s"
	  record_delimiter = "%s"
	  record_prefix    = "%s"
	  record_suffix    = "%s"
	  record_template  = "%s"
	  sample_rate      = %f
	  timestamp_format = "%s"
  }
}
