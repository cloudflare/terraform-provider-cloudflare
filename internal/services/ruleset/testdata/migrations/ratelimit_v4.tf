resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Rate Limit Ruleset %[2]s"
  phase   = "http_ratelimit"
  kind    = "zone"

  rules {
    expression = "true"
    action     = "block"
    description = "Basic rate limit with characteristics"

    ratelimit {
      characteristics = ["cf.colo.id", "ip.src"]
      period          = 60
      requests_per_period = 100
      mitigation_timeout = 600
    }
  }

  rules {
    expression = "http.request.uri.path contains \"/api\""
    action     = "challenge"
    description = "API rate limit with counting expression"

    ratelimit {
      characteristics = ["cf.colo.id", "ip.src"]
      period          = 10
      requests_per_period = 5
      counting_expression = "http.request.method eq \"POST\""
      mitigation_timeout = 0
    }
  }

  rules {
    expression = "http.host eq \"example.com\""
    action     = "block"
    description = "Rate limit with mitigation timeout"

    ratelimit {
      characteristics = ["cf.colo.id", "cf.unique_visitor_id"]
      period          = 300
      requests_per_period = 1000
      mitigation_timeout = 1800
    }
  }
}
