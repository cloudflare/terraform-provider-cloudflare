resource "cloudflare_pipeline_stream" "example_pipeline_stream" {
  account_id = "0123105f4ecef8ad9ca31a8372d0c353"
  name = "my_stream"
  format = {
    type = "json"
    decimal_encoding = "number"
    timestamp_format = "rfc3339"
    unstructured = true
  }
  http = {
    authentication = false
    enabled = true
    cors = {
      origins = ["string"]
    }
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
  worker_binding = {
    enabled = true
  }
}
