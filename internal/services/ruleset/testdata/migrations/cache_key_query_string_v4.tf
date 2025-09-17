resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Cache settings test %[2]s"
  phase   = "http_request_cache_settings"
  kind    = "zone"

  rules {
    action = "set_cache_settings"
    action_parameters {
      cache = true
      cache_key {
        custom_key {
          query_string {
            include = ["param1", "param2", "param3"]
          }
          user {
            device_type = true
            lang        = true
          }
        }
        ignore_query_strings_order = true
      }
      edge_ttl {
        default = 3600
        mode    = "override_origin"
      }
      serve_stale {
        disable_stale_while_updating = true
      }
    }
    description = "Cache with specific query params"
    enabled     = true
    expression  = "http.request.uri.path matches \"^/api/\""
  }

  rules {
    action = "set_cache_settings"
    action_parameters {
      cache = true
      cache_key {
        custom_key {
          query_string {
            include = ["*"]
          }
        }
      }
    }
    description = "Cache with all query params"
    enabled     = false
    expression  = "http.request.uri.path eq \"/test\""
  }
}