resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Redirect ruleset %[2]s"
  phase   = "http_request_dynamic_redirect"
  kind    = "zone"

  rules {
    action = "redirect"
    action_parameters {
      from_value {
        preserve_query_string = true
        status_code           = 308
        target_url {
          expression = "concat(\"https://example.com\", http.request.uri.path)"
        }
      }
    }
    description = "Dynamic redirect with expression"
    enabled     = true
    expression  = "http.host eq \"old.example.com\""
  }

  rules {
    action = "redirect"
    action_parameters {
      from_value {
        preserve_query_string = false
        status_code           = 302
        target_url {
          value = "https://example.com/new-path"
        }
      }
    }
    description = "Static redirect"
    enabled     = false
    expression  = "http.request.uri.path eq \"/old-path\""
  }
}