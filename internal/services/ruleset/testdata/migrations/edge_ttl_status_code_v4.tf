resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Edge TTL Status Code Ruleset %[2]s"
  phase   = "http_request_cache_settings"
  kind    = "zone"

  rules {
    expression = "http.request.uri.path contains \"/static\""
    action     = "set_cache_settings"
    description = "Cache static content with status code TTL"

    action_parameters {
      cache = true

      edge_ttl {
        mode    = "override_origin"
        default = 7200

        status_code_ttl {
          status_code = 200
          value       = 86400
        }

        status_code_ttl {
          status_code = 404
          value       = 300
        }
      }
    }
  }

  rules {
    expression = "http.request.uri.path contains \"/api\""
    action     = "set_cache_settings"
    description = "API with status code range"

    action_parameters {
      cache = true

      edge_ttl {
        mode    = "override_origin"
        default = 3600

        status_code_ttl {
          status_code_range {
            from = 200
            to   = 299
          }
          value = 3600
        }
      }
    }
  }

  rules {
    expression = "http.request.uri.path contains \"/images\""
    action     = "set_cache_settings"
    description = "Images with multiple status code TTLs"

    action_parameters {
      cache = true

      edge_ttl {
        mode    = "respect_origin"

        status_code_ttl {
          status_code = 200
          value       = 604800
        }

        status_code_ttl {
          status_code_range {
            from = 400
            to   = 499
          }
          value = 60
        }

        status_code_ttl {
          status_code = 500
          value       = 0
        }
      }
    }
  }
}
