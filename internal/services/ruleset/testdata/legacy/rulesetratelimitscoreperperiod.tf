
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_ratelimit"

  rules = [{
    action = "block"
    action_parameters = {
      response = {
        status_code  = 418
        content_type = "text/plain"
        content      = "test content"
      }
    }
    ratelimit = {
      characteristics = [
        "cf.colo.id",
        "ip.src"
      ]
      period                     = 60
      score_per_period           = 400
      score_response_header_name = "my-score"
      mitigation_timeout         = 60
      requests_to_origin         = true
    }
    expression  = "(http.request.uri.path matches \"^/api/\")"
    description = "example http rate limit"
    enabled     = true
  }]
}
