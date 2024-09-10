
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_ratelimit"

    rules =[ {
      action = "block"
      action_parameters = {
    response = {
          status_code = 418
          content_type = "text/plain"
          content = "test content"
        }
  }
      ratelimit = {
    characteristics = [
          "cf.colo.id",
          "ip.src"
        ]
        period              = 60
        requests_per_period = 1000
        requests_to_origin  = false
        mitigation_timeout  = 0
  }
      expression = "(http.request.uri.path matches \"^/api/\")"
      description = "example http rate limit"
      enabled = true
    }]
  }
