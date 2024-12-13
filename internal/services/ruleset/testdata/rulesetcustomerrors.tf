
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_custom_errors"

  rules = [{
    action = "serve_error"
    action_parameters = {
      content      = "my example error page"
      content_type = "text/plain"
      status_code  = "530"
    }
    expression  = "(http.request.uri.path matches \"^/api/\")"
    description = "example http custom error response"
    enabled     = true
  }]
}