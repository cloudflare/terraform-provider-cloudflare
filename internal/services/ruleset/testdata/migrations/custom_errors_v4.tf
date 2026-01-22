resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Custom Errors Ruleset %[2]s"
  phase   = "http_custom_errors"
  kind    = "zone"

  rules {
    expression = "true"
    action     = "serve_error"
    description = "Custom 404 error page"

    action_parameters {
      content      = "<html><body>Custom 404 Error</body></html>"
      content_type = "text/html"
      status_code  = 404
    }
  }

  rules {
    expression = "http.host eq \"example.com\""
    action     = "serve_error"
    description = "Custom 500 error page"

    action_parameters {
      content      = "<html><body>Custom 500 Error</body></html>"
      content_type = "text/html"
      status_code  = 500
    }
  }

  rules {
    expression = "http.request.uri.path contains \"/api\""
    action     = "serve_error"
    description = "Custom 403 error with JSON response"

    action_parameters {
      content      = "{\"error\": \"Forbidden\"}"
      content_type = "application/json"
      status_code  = 403
    }
  }
}
