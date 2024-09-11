
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_transform"

  rules = [{
    action = "rewrite"
    action_parameters = {
      uri = {
        query = {
          expression = "concat(\"requestUrl=\", http.request.full_uri)"
        }
        path = {
          value = "/path/to/url"
        }
      }
    }
    expression  = "true"
    description = "example for combining URI action parameters for path and query"
    enabled     = true
  }]
}