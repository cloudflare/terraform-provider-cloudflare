resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Redirect From List Ruleset %[2]s"
  phase   = "http_request_dynamic_redirect"
  kind    = "zone"

  rules {
    expression = "http.request.uri.path contains \"/blocked\""
    action     = "redirect"
    description = "Redirect blocked paths"

    action_parameters {
      from_value {
        status_code = 302
        target_url {
          value = "https://example.com/blocked"
        }
        preserve_query_string = false
      }
    }
  }

  rules {
    expression = "http.request.uri.path contains \"/legacy\""
    action     = "redirect"
    description = "Redirect legacy paths"

    action_parameters {
      from_value {
        status_code = 301
        target_url {
          expression = "concat(\"https://example.com/new\", http.request.uri.path)"
        }
        preserve_query_string = true
      }
    }
  }

  rules {
    expression = "http.host eq \"redirect.example.com\""
    action     = "redirect"
    description = "Static redirect"

    action_parameters {
      from_value {
        status_code = 301
        target_url {
          value = "https://example.com/new-location"
        }
        preserve_query_string = true
      }
    }
  }
}
