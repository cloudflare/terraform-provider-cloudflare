resource "cloudflare_pipeline_sink" "example_pipeline_sink" {
  account_id = "0123105f4ecef8ad9ca31a8372d0c353"
  name = "my_sink"
  type = "r2"
  config = {
    account_id = "account_id"
    bucket = "bucket"
    credentials = {
      access_key_id = "access_key_id"
      secret_access_key = "secret_access_key"
    }
    file_naming = {
      prefix = "prefix"
      strategy = "serial"
      suffix = "suffix"
    }
    jurisdiction = "jurisdiction"
    partitioning = {
      time_pattern = "year=%Y/month=%m/day=%d/hour=%H"
    }
    path = "path"
    rolling_policy = {
      file_size_bytes = 0
      inactivity_seconds = 1
      interval_seconds = 1
    }
  }
  format = {
    type = "json"
    decimal_encoding = "number"
    timestamp_format = "rfc3339"
    unstructured = true
  }
  schema = {
    fields = [{
      type = "int32"
      metadata_key = "metadata_key"
      name = "name"
      required = true
      sql_name = "sql_name"
    }]
    format = {
      type = "json"
      decimal_encoding = "number"
      timestamp_format = "rfc3339"
      unstructured = true
    }
    inferred = true
  }
}
