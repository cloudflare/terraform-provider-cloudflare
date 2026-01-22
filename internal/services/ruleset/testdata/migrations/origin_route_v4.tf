resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Origin Route Ruleset %[2]s"
  phase   = "http_request_origin"
  kind    = "zone"

  rules {
    expression = "http.request.uri.path contains \"/api\""
    action     = "route"
    description = "Route API requests with port override"

    action_parameters {
      origin {
        port = 8443
      }
    }
  }

  rules {
    expression = "http.request.uri.path contains \"/secure\""
    action     = "route"
    description = "Route with port 8080"

    action_parameters {
      origin {
        port = 8080
      }
    }
  }

  rules {
    expression = "http.request.uri.path contains \"/cdn\""
    action     = "route"
    description = "Route with standard HTTPS port"

    action_parameters {
      origin {
        port = 443
      }
    }
  }
}
