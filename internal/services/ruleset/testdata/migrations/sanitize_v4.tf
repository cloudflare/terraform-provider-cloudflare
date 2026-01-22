resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Request Late Transform Ruleset %[2]s"
  phase   = "http_request_late_transform"
  kind    = "zone"

  rules {
    expression = "true"
    action     = "rewrite"
    description = "Remove multiple request headers"

    action_parameters {
      headers {
        name      = "X-Forwarded-For"
        operation = "remove"
      }

      headers {
        name      = "X-Real-IP"
        operation = "remove"
      }
    }
  }

  rules {
    expression = "http.request.uri.path contains \"/admin\""
    action     = "rewrite"
    description = "Remove sensitive headers from admin requests"

    action_parameters {
      headers {
        name      = "Authorization"
        operation = "remove"
      }
    }
  }

  rules {
    expression = "http.host eq \"example.com\""
    action     = "rewrite"
    description = "Set custom header"

    action_parameters {
      headers {
        name      = "X-Custom-Header"
        operation = "set"
        value     = "custom-value"
      }
    }
  }
}
