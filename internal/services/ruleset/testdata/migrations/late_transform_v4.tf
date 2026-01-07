resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Late Transform Ruleset %[2]s"
  phase   = "http_request_late_transform"
  kind    = "zone"

  rules {
    expression = "http.request.uri.path contains \"/api\""
    action     = "rewrite"
    description = "Rewrite request headers"

    action_parameters {
      headers {
        name      = "X-API-Key"
        operation = "set"
        value     = "api-key-value"
      }

      headers {
        name      = "X-Custom-Header"
        operation = "set"
        value     = "custom-value"
      }

      headers {
        name      = "X-Remove-Header"
        operation = "remove"
      }
    }
  }

  rules {
    expression = "http.host eq \"example.com\""
    action     = "rewrite"
    description = "Modify headers for domain"

    action_parameters {
      headers {
        name      = "X-Domain"
        operation = "set"
        value     = "example.com"
      }
    }
  }

  rules {
    expression = "http.request.uri.path contains \"/secure\""
    action     = "rewrite"
    description = "Add security headers"

    action_parameters {
      headers {
        name      = "X-Secure-Path"
        operation = "set"
        value     = "true"
      }
    }
  }
}
