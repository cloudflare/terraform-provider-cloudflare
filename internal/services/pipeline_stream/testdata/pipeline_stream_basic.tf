resource "cloudflare_pipeline_stream" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  format = {
    type = "json"
  }
  schema = {
    fields = [{
      name     = "value"
      type     = "json"
      required = true
    }]
  }
  http = {
    enabled        = true
    authentication = false
    cors           = {}
  }
  worker_binding = {
    enabled = false
  }
}
