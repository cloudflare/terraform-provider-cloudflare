# Magic Transit
resource "cloudflare_ruleset" "magic_transit_example" {
  account_id  = "d41d8cd98f00b204e9800998ecf8427e"
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
  zone_id     = "cb029e245cfdd66dc8d2e570d5dd3322"
  name        = "managed WAF"
  description = "managed WAF ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_managed"

  rules {
    action = "execute"
    action_parameters {
      id = "efb7b8c949ac4650a09736fc376e9aee"
    }
    expression  = "true"
    description = "Execute Cloudflare Managed Ruleset on my zone-level phase entry point ruleset"
    enabled     = true
  }
}

# Zone-level WAF with tag-based overrides
resource "cloudflare_ruleset" "zone_level_managed_waf_with_category_based_overrides" {
  zone_id     = "cb029e245cfdd66dc8d2e570d5dd3322"
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

    expression  = "true"
    description = "overrides to only enable wordpress rules to block"
    enabled     = false
  }
}

# Rewrite the URI path component to a static path
resource "cloudflare_ruleset" "transform_uri_rule_path" {
  zone_id     = "cb029e245cfdd66dc8d2e570d5dd3322"
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
  zone_id     = "cb029e245cfdd66dc8d2e570d5dd3322"
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

    expression  = "true"
    description = "URI transformation query example"
    enabled     = true
  }
}

# Rewrite HTTP headers to a modified values
resource "cloudflare_ruleset" "transform_uri_http_headers" {
  zone_id     = "cb029e245cfdd66dc8d2e570d5dd3322"
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

    expression  = "true"
    description = "example request header transform rule"
    enabled     = false
  }
}

# HTTP rate limit for an API route
resource "cloudflare_ruleset" "rate_limiting_example" {
  zone_id     = "cb029e245cfdd66dc8d2e570d5dd3322"
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
  zone_id     = "cb029e245cfdd66dc8d2e570d5dd3322"
  name        = "Change to some origin"
  description = "Change origin for a route"
  kind        = "zone"
  phase       = "http_request_origin"

  rules {
    action = "route"
    action_parameters {
      host_header = "some.host"
      origin = {
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
  zone_id     = "cb029e245cfdd66dc8d2e570d5dd3322"
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

    expression  = "true"
    description = "log custom fields rule"
    enabled     = true
  }
}
