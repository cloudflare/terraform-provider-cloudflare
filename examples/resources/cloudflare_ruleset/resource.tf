# Magic Transit
resource "cloudflare_ruleset" "magic_transit_example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "My example Magic Transit ruleset"
  description = "My example Magic Transit ruleset description"
  phase       = "magic_transit"
  kind        = "root"

  rules {
    ref         = "allow_tcp_ephemeral_ports"
    description = "Allow TCP Ephemeral Ports"
    expression  = "tcp.dstport in { 32768..65535 }"
    action      = "allow"
  }
}

# Zone-level WAF Managed Ruleset
resource "cloudflare_ruleset" "zone_level_managed_waf" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example managed WAF ruleset"
  description = "My example managed WAF ruleset description"
  phase       = "http_request_firewall_managed"
  kind        = "zone"

  rules {
    ref         = "execute_managed_ruleset"
    description = "Execute Cloudflare Managed Ruleset on my zone-level phase entry point ruleset"
    expression  = "(http.host eq \"example.host.com\")"
    action      = "execute"
    action_parameters {
      id = "efb7b8c949ac4650a09736fc376e9aee"
    }
  }
}

# Zone-level WAF with tag-based overrides
resource "cloudflare_ruleset" "zone_level_managed_waf_with_category_based_overrides" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example managed WAF ruleset with tag-based overrides"
  description = "My example managed WAF ruleset with tag-based overrides ruleset description"
  phase       = "http_request_firewall_managed"
  kind        = "zone"

  rules {
    ref         = "execute_managed_ruleset"
    description = "Execute Cloudflare Managed Ruleset with overrides to change Wordpress rules to block"
    expression  = "(http.host eq \"example.host.com\")"
    action      = "execute"
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
    enabled = false
  }
}

# Rewrite the URI path component to a static path
resource "cloudflare_ruleset" "transform_uri_rule_path" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example transform ruleset"
  description = "My example transform ruleset description"
  phase       = "http_request_transform"
  kind        = "zone"

  rules {
    ref         = "transform_old_path"
    description = "Transform old path"
    expression  = "(http.host eq \"example.com\" and http.request.uri.path eq \"/old-path\")"
    action      = "rewrite"
    action_parameters {
      uri {
        path {
          value = "/my-new-route"
        }
      }
    }
  }
}

# Rewrite the URI query component to a static query
resource "cloudflare_ruleset" "transform_uri_rule_query" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example transform ruleset"
  description = "My example transform ruleset description"
  phase       = "http_request_transform"
  kind        = "zone"

  rules {
    ref         = "transform_uri_query_parameter"
    description = "Transform URI query parameter"
    expression  = "(http.host eq \"example.host.com\")"
    action      = "rewrite"
    action_parameters {
      uri {
        query {
          value = "old=new_again"
        }
      }
    }
  }
}

# Rewrite HTTP headers to a modified values
resource "cloudflare_ruleset" "transform_uri_http_headers" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example transform ruleset"
  description = "My example transform ruleset description"
  phase       = "http_request_transform"
  kind        = "zone"

  rules {
    ref         = "transform_request_headers"
    description = "Transform request headers"
    expression  = "(http.host eq \"example.host.com\")"
    action      = "rewrite"
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
  }
}

# HTTP rate limit for an API route
resource "cloudflare_ruleset" "rate_limiting_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example rate limit ruleset"
  description = "My example rate limit ruleset description"
  phase       = "http_ratelimit"
  kind        = "zone"

  rules {
    ref         = "rate_limit_api_requests"
    description = "Rate limit API requests"
    expression  = "(http.request.uri.path matches \"^/api/\")"
    action      = "block"
    ratelimit {
      characteristics = [
        "cf.colo.id",
        "ip.src"
      ]
      period              = 60
      requests_per_period = 100
      mitigation_timeout  = 600
    }
  }
}

# Change origin for an API route
resource "cloudflare_ruleset" "http_origin_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example origin ruleset"
  description = "My example origin ruleset description"
  phase       = "http_request_origin"
  kind        = "zone"

  rules {
    ref         = "change_origin"
    description = "Change to some.host"
    expression  = "(http.request.uri.path matches \"^/api/\")"
    action      = "route"
    action_parameters {
      host_header = "some.host"
      origin {
        host = "some.host"
        port = 80
      }
    }
  }
}

# Custom fields logging
resource "cloudflare_ruleset" "custom_fields_logging_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example log custom field ruleset"
  description = "My example log custom field ruleset description"
  phase       = "http_log_custom_fields"
  kind        = "zone"

  rules {
    ref         = "log_custom_fields"
    description = "Log custom fields"
    expression  = "(http.host eq \"example.host.com\")"
    action      = "log_custom_field"
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
  }
}

# Custom cache keys and settings
resource "cloudflare_ruleset" "cache_settings_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example cache settings ruleset"
  description = "My example cache settings ruleset description"
  phase       = "http_request_cache_settings"
  kind        = "zone"

  rules {
    ref         = "cache_settings"
    description = "Set cache settings rule"
    expression  = "(http.host eq \"example.host.com\")"
    action      = "set_cache_settings"
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
            contains = {
              "accept"          = ["image/webp", "image/png"]
              "accept-encoding" = ["br", "zstd"]
              "some-header"     = ["some-value", "some-other-value"]
            }
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
      cache_reserve = {
        eligible          = true
        minimum_file_size = 100000
      }
      origin_error_page_passthru = false
    }
  }
}

# Redirects based on a List resource
resource "cloudflare_ruleset" "redirect_from_list_example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "My example redirect ruleset"
  description = "My example redirect ruleset description"
  phase       = "http_request_redirect"
  kind        = "root"

  rules {
    ref         = "redirects_from_list"
    description = "Apply redirects from redirect_list"
    expression  = "http.request.full_uri in $redirect_list"
    action      = "redirect"
    action_parameters {
      from_list {
        name = "redirect_list"
        key  = "http.request.full_uri"
      }
    }
  }
}

# Dynamic Redirects from value resource
resource "cloudflare_ruleset" "redirect_from_value_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example dynamic redirect ruleset"
  description = "My example dynamic redirect ruleset description"
  phase       = "http_request_dynamic_redirect"
  kind        = "zone"

  rules {
    ref         = "redirect_from_value"
    description = "Apply redirect from value"
    expression  = "(http.request.uri.path matches \"^/api/\")"
    action      = "redirect"
    action_parameters {
      from_value {
        status_code = 301
        target_url {
          value = "some_host.com"
        }
        preserve_query_string = true
      }
    }
  }
}

# Serve some custom error response
resource "cloudflare_ruleset" "http_custom_error_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example custom errors ruleset"
  description = "My example custom errors ruleset description"
  phase       = "http_custom_errors"
  kind        = "zone"

  rules {
    ref         = "serve_some_error_response"
    description = "Serve some error response"
    expression  = "(http.request.uri.path matches \"^/api/\")"
    action      = "serve_error"
    action_parameters {
      content      = "some error html"
      content_type = "text/html"
      status_code  = "530"
    }
  }
}

# Set Configuration Rules for an API route
resource "cloudflare_ruleset" "http_config_rules_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example config settings ruleset"
  description = "My example config settings ruleset description"
  phase       = "http_config_settings"
  kind        = "zone"

  rules {
    ref         = "set_config_settings"
    description = "Set config settings"
    expression  = "(http.request.uri.path matches \"^/api/\")"
    action      = "set_config"
    action_parameters {
      email_obfuscation = true
      bic               = true
    }
  }
}

# Set compress algorithm for response
resource "cloudflare_ruleset" "response_compress_brotli_html" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "My example response compression ruleset"
  description = "My example response compression description"
  phase       = "http_response_compression"
  kind        = "zone"

  rules {
    ref         = "prefer_brotli_for_html"
    description = "Prefer Brotli compression for HTML"
    expression  = "http.response.content_type.media_type == \"text/html\""
    action      = "compress_response"
    action_parameters {
      algorithms {
        name = "brotli"
      }
      algorithms {
        name = "auto"
      }
    }
  }
}
