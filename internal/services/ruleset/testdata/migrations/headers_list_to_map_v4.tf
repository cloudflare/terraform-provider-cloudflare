resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Headers transform test %[2]s"
  phase   = "http_response_headers_transform"
  kind    = "zone"

  rules {
    action = "rewrite"
    action_parameters {
      headers {
        name      = "X-Custom-Header"
        operation = "set"
        value     = "custom-value"
      }
    }
    description = "Set single header"
    enabled     = true
    expression  = "true"
  }

  rules {
    action = "rewrite"
    action_parameters {
      headers {
        name      = "Authorization"
        operation = "remove"
      }
    }
    description = "Remove auth header"
    enabled     = true
    expression  = "http.host eq \"example.com\""
  }
}