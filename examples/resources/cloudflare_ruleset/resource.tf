# Magic Transit
resource "cloudflare_ruleset" "magic_transit_example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "account magic transit"
  description = "example magic transit ruleset description"
  kind        = "root"
  phase       = "magic_transit"

  rules {
    action      = "allow"
    expression  = "tcp.dstport in { 32768..65535 }"
    description = "Allow TCP Ephemeral Ports"
  }
}

# Zone-level WAF Managed Ruleset
resource "cloudflare_ruleset" "zone_level_managed_waf" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "managed WAF"
  description = "managed WAF ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_managed"

  rules {
    action = "execute"
    action_parameters {
      id = "efb7b8c949ac4650a09736fc376e9aee"
    }
    expression  = "(http.host eq \"example.host.com\")"
    description = "Execute Cloudflare Managed Ruleset on my zone-level phase entry point ruleset"
    enabled     = true
  }
}

# Zone-level WAF with tag-based overrides
resource "cloudflare_ruleset" "zone_level_managed_waf_with_category_based_overrides" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "managed WAF with tag-based overrides"
  description = "managed WAF with tag-based overrides ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_managed"

  rules {
    action = "execute"
    action_parameters {
      id = "efb7b8c949ac4650a09736fc376e9aee"
      overrides {
        categories {
          category = "wordpress"
          action   = "block"
          enabled  = true
        }

        categories {
          category = "joomla"
          action   = "block"
          enabled  = true
        }
      }
    }

    expression  = "(http.host eq \"example.host.com\")"
    description = "overrides to only enable wordpress rules to block"
    enabled     = false
  }
}

# Rewrite the URI path component to a static path
resource "cloudflare_ruleset" "transform_uri_rule_path" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "transform rule for URI path"
  description = "change the URI path to a new static path"
  kind        = "zone"
  phase       = "http_request_transform"

  rules {
    action = "rewrite"
    action_parameters {
      uri {
        path {
          value = "/my-new-route"
        }
      }
    }

    expression  = "(http.host eq \"example.com\" and http.request.uri.path eq \"/old-path\")"
    description = "example URI path transform rule"
    enabled     = true
  }
}

# Rewrite the URI query component to a static query
resource "cloudflare_ruleset" "transform_uri_rule_query" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "transform rule for URI query parameter"
  description = "change the URI query to a new static query"
  kind        = "zone"
  phase       = "http_request_transform"

  rules {
    action = "rewrite"
    action_parameters {
      uri {
        query {
          value = "old=new_again"
        }
      }
    }

    expression  = "(http.host eq \"example.host.com\")"
    description = "URI transformation query example"
    enabled     = true
  }
}

# Rewrite HTTP headers to a modified values
resource "cloudflare_ruleset" "transform_uri_http_headers" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "transform rule for HTTP headers"
  description = "modify HTTP headers before reaching origin"
  kind        = "zone"
  phase       = "http_request_late_transform"

  rules {
    action = "rewrite"
    action_parameters {
      headers {
        name      = "example-http-header-1"
        operation = "set"
        value     = "my-http-header-value-1"
      }

      headers {
        name       = "example-http-header-2"
        operation  = "set"
        expression = "cf.zone.name"
      }

      headers {
        name      = "example-http-header-3-to-remove"
        operation = "remove"
      }
    }

    expression  = "(http.host eq \"example.host.com\")"
    description = "example request header transform rule"
    enabled     = false
  }
}

# HTTP rate limit for an API route
resource "cloudflare_ruleset" "rate_limiting_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "restrict API requests count"
  description = "apply HTTP rate limiting for a route"
  kind        = "zone"
  phase       = "http_ratelimit"

  rules {
    action = "block"
    ratelimit {
      characteristics = [
        "cf.colo.id",
        "ip.src"
      ]
      period              = 60
      requests_per_period = 100
      mitigation_timeout  = 600
    }

    expression  = "(http.request.uri.path matches \"^/api/\")"
    description = "rate limit for API"
    enabled     = true
  }
}

# Change origin for an API route
resource "cloudflare_ruleset" "http_origin_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "Change to some origin"
  description = "Change origin for a route"
  kind        = "zone"
  phase       = "http_request_origin"

  rules {
    action = "route"
    action_parameters {
      host_header = "some.host"
      origin {
        host = "some.host"
        port = 80
      }
    }
    expression  = "(http.request.uri.path matches \"^/api/\")"
    description = "change origin to some.host"
    enabled     = true
  }
}

# Custom fields logging
resource "cloudflare_ruleset" "custom_fields_logging_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "log custom fields"
  description = "add custom fields to logging"
  kind        = "zone"
  phase       = "http_log_custom_fields"

  rules {
    action = "log_custom_field"
    action_parameters {
      request_fields = [
        "content-type",
        "x-forwarded-for",
        "host"
      ]
      response_fields = [
        "server",
        "content-type",
        "allow"
      ]
      cookie_fields = [
        "__ga",
        "accountNumber",
        "__cfruid"
      ]
    }

    expression  = "(http.host eq \"example.host.com\")"
    description = "log custom fields rule"
    enabled     = true
  }
}

# Custom cache keys + settings
resource "cloudflare_ruleset" "cache_settings_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "set cache settings"
  description = "set cache settings for the request"
  kind        = "zone"
  phase       = "http_request_cache_settings"

  rules {
    action = "set_cache_settings"
    action_parameters {
      edge_ttl {
        mode    = "override_origin"
        default = 60
        status_code_ttl {
          status_code = 200
          value       = 50
        }
        status_code_ttl {
          status_code_range {
            from = 201
            to   = 300
          }
          value = 30
        }
      }
      browser_ttl {
        mode = "respect_origin"
      }
      serve_stale {
        disable_stale_while_updating = true
      }
      respect_strong_etags = true
      cache_key {
        ignore_query_strings_order = false
        cache_deception_armor      = true
        custom_key {
          query_string {
            exclude = ["*"]
          }
          header {
            include        = ["habc", "hdef"]
            check_presence = ["habc_t", "hdef_t"]
            exclude_origin = true
          }
          cookie {
            include        = ["cabc", "cdef"]
            check_presence = ["cabc_t", "cdef_t"]
          }
          user {
            device_type = true
            geo         = false
          }
          host {
            resolved = true
          }
        }
      }
      origin_error_page_passthru = false
    }
    expression  = "(http.host eq \"example.host.com\")"
    description = "set cache settings rule"
    enabled     = true
  }
}

# Redirects based on a List resource
resource "cloudflare_ruleset" "redirect_from_list_example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "redirects"
  description = "Redirect ruleset"
  kind        = "root"
  phase       = "http_request_redirect"

  rules {
    action = "redirect"
    action_parameters {
      from_list {
        name = "redirect_list"
        key  = "http.request.full_uri"
      }
    }
    expression  = "http.request.full_uri in $redirect_list"
    description = "Apply redirects from redirect_list"
    enabled     = true
  }
}

# Dynamic Redirects from value resource
resource "cloudflare_ruleset" "redirect_from_value_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "redirects"
  description = "Redirect ruleset"
  kind        = "zone"
  phase       = "http_request_dynamic_redirect"

  rules {
    action = "redirect"
    action_parameters {
      from_value {
        status_code = 301
        target_url {
          value = "some_host.com"
        }
        preserve_query_string = true
      }
    }
    expression  = "(http.request.uri.path matches \"^/api/\")"
    description = "Apply redirect from value"
    enabled     = true
  }
}

# Serve some custom error response
resource "cloudflare_ruleset" "http_custom_error_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "Serve some error response"
  description = "Serve some error response"
  kind        = "zone"
  phase       = "http_custom_errors"
  rules {
    action = "serve_error"
    action_parameters {
      content      = "some error html"
      content_type = "text/html"
      status_code  = "530"
    }
    expression  = "(http.request.uri.path matches \"^/api/\")"
    description = "serve some error response"
    enabled     = true
  }
}

# Set Configuration Rules for an API route
resource "cloudflare_ruleset" "http_config_rules_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "set config rules"
  description = "set config rules for request"
  kind        = "zone"
  phase       = "http_config_settings"

  rules {
    action = "set_config"
    action_parameters {
      email_obfuscation = true
      bic               = true
    }
    expression  = "(http.request.uri.path matches \"^/api/\")"
    description = "set config rules for matching request"
    enabled     = true
  }
}

# Set compress algorithm for response.
resource "cloudflare_ruleset" "response_compress_brotli_html" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "Brotli response compression for HTML"
  description = "Response compression ruleset"
  kind        = "zone"
  phase       = "http_response_compression"

  rules {
    action = "compress_response"
    action_parameters {
      algorithms {
        name = "brotli"
      }
      algorithms {
        name = "auto"
      }
    }
    expression  = "http.response.content_type.media_type == \"text/html\""
    description = "Prefer brotli compression for HTML"
    enabled     = true
  }
}
