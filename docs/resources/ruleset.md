---
page_title: "cloudflare_ruleset Resource - Cloudflare"
subcategory: ""
description: |-
  The Cloudflare Ruleset Engine https://developers.cloudflare.com/firewall/cf-rulesets
  allows you to create and deploy rules and rulesets.
  The engine syntax, inspired by the Wireshark Display Filter language, is the
  same syntax used in custom Firewall Rules. Cloudflare uses the Ruleset Engine
  in different products, allowing you to configure several products using the same
  basic syntax.
---

# cloudflare_ruleset (Resource)

The [Cloudflare Ruleset Engine](https://developers.cloudflare.com/firewall/cf-rulesets)
allows you to create and deploy rules and rulesets.

The engine syntax, inspired by the Wireshark Display Filter language, is the
same syntax used in custom Firewall Rules. Cloudflare uses the Ruleset Engine
in different products, allowing you to configure several products using the same
basic syntax.

~> If you previously configured Rulesets using the dashboard,
you first need to delete them ([zone](https://api.cloudflare.com/#zone-rulesets-delete-zone-ruleset),
[account](https://api.cloudflare.com/#account-rulesets-delete-account-ruleset) documentation)
and clean up the resources before attempting to configure them with
Terraform. This is because Terraform will fail to apply if configuration
already exists to prevent blindly overwriting changes.

~> `enabled` has been immediately deprecated in favour of
`status`. You should swap over to ensure that your configuration doesn't
have inconsistent operations and inadvertently disable rulesets.

## Example Usage

```terraform
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
          status   = "enabled"
        }

        categories {
          category = "joomla"
          action   = "block"
          status   = "enabled"
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
  kind        = "root"
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `kind` (String) Type of Ruleset to create. Available values: `custom`, `managed`, `root`, `schema`, `zone`.
- `name` (String) Name of the ruleset. **Modifying this attribute will force creation of a new resource.**
- `phase` (String) Point in the request/response lifecycle where the ruleset will be created. Available values: `ddos_l4`, `ddos_l7`, `http_custom_errors`, `http_log_custom_fields`, `http_request_cache_settings`, `http_request_firewall_custom`, `http_request_firewall_managed`, `http_request_late_transform`, `http_request_late_transform_managed`, `http_request_main`, `http_request_origin`, `http_request_dynamic_redirect`, `http_request_redirect`, `http_request_sanitize`, `http_request_transform`, `http_response_firewall_managed`, `http_response_headers_transform`, `http_response_headers_transform_managed`, `magic_transit`, `http_ratelimit`, `http_request_sbfm`, `http_config_settings`.

### Optional

- `account_id` (String) The account identifier to target for the resource. Conflicts with `zone_id`.
- `description` (String) Brief summary of the ruleset and its intended use.
- `rules` (Block List) List of rules to apply to the ruleset. (see [below for nested schema](#nestedblock--rules))
- `shareable_entitlement_name` (String) Name of entitlement that is shareable between entities.
- `zone_id` (String) The zone identifier to target for the resource. Conflicts with `account_id`.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--rules"></a>
### Nested Schema for `rules`

Required:

- `expression` (String) Criteria for an HTTP request to trigger the ruleset rule action. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.

Optional:

- `action` (String) Action to perform in the ruleset rule. Available values: `block`, `challenge`, `ddos_dynamic`, `execute`, `force_connection_close`, `js_challenge`, `log`, `log_custom_field`, `managed_challenge`, `redirect`, `rewrite`, `route`, `score`, `set_cache_settings`, `set_config`, `serve_error`, `skip`.
- `action_parameters` (Block List, Max: 1) List of parameters that configure the behavior of the ruleset rule action. (see [below for nested schema](#nestedblock--rules--action_parameters))
- `description` (String) Brief summary of the ruleset rule and its intended use.
- `enabled` (Boolean) Whether the rule is active.
- `exposed_credential_check` (Block List, Max: 1) List of parameters that configure exposed credential checks. (see [below for nested schema](#nestedblock--rules--exposed_credential_check))
- `last_updated` (String) The most recent update to this rule.
- `logging` (Block List, Max: 1) List parameters to configure how the rule generates logs. (see [below for nested schema](#nestedblock--rules--logging))
- `ratelimit` (Block List, Max: 1) List of parameters that configure HTTP rate limiting behaviour. (see [below for nested schema](#nestedblock--rules--ratelimit))

Read-Only:

- `id` (String) Unique rule identifier.
- `ref` (String) Rule reference.
- `version` (String) Version of the ruleset to deploy.

<a id="nestedblock--rules--action_parameters"></a>
### Nested Schema for `rules.action_parameters`

Optional:

- `automatic_https_rewrites` (Boolean) Turn on or off Cloudflare Automatic HTTPS rewrites.
- `autominify` (Block List) Indicate which file extensions to minify automatically. (see [below for nested schema](#nestedblock--rules--action_parameters--autominify))
- `bic` (Boolean) Inspect the visitor's browser for headers commonly associated with spammers and certain bots.
- `browser_ttl` (Block List, Max: 1) List of browser TTL parameters to apply to the request. (see [below for nested schema](#nestedblock--rules--action_parameters--browser_ttl))
- `cache` (Boolean) Whether to cache if expression matches.
- `cache_key` (Block List, Max: 1) List of cache key parameters to apply to the request. (see [below for nested schema](#nestedblock--rules--action_parameters--cache_key))
- `content` (String) Content of the custom error response.
- `content_type` (String) Content-Type of the custom error response.
- `cookie_fields` (Set of String) List of cookie values to include as part of custom fields logging.
- `disable_apps` (Boolean) Turn off all active Cloudflare Apps.
- `disable_railgun` (Boolean) Turn off railgun feature of the Cloudflare Speed app.
- `disable_zaraz` (Boolean) Turn off zaraz feature.
- `edge_ttl` (Block List, Max: 1) List of edge TTL parameters to apply to the request. (see [below for nested schema](#nestedblock--rules--action_parameters--edge_ttl))
- `email_obfuscation` (Boolean) Turn on or off the Cloudflare Email Obfuscation feature of the Cloudflare Scrape Shield app.
- `from_list` (Block List, Max: 1) Use a list to lookup information for the action. (see [below for nested schema](#nestedblock--rules--action_parameters--from_list))
- `from_value` (Block List, Max: 1) Use a value to lookup information for the action. (see [below for nested schema](#nestedblock--rules--action_parameters--from_value))
- `headers` (Block List) List of HTTP header modifications to perform in the ruleset rule. (see [below for nested schema](#nestedblock--rules--action_parameters--headers))
- `host_header` (String) Host Header that request origin receives.
- `hotlink_protection` (Boolean) Turn on or off the hotlink protection feature.
- `id` (String) Identifier of the action parameter to modify.
- `increment` (Number)
- `matched_data` (Block List, Max: 1) List of properties to configure WAF payload logging. (see [below for nested schema](#nestedblock--rules--action_parameters--matched_data))
- `mirage` (Boolean) Turn on or off Cloudflare Mirage of the Cloudflare Speed app.
- `opportunistic_encryption` (Boolean) Turn on or off the Cloudflare Opportunistic Encryption feature of the Edge Certificates tab in the Cloudflare SSL/TLS app.
- `origin` (Block List, Max: 1) List of properties to change request origin. (see [below for nested schema](#nestedblock--rules--action_parameters--origin))
- `origin_error_page_passthru` (Boolean) Pass-through error page for origin.
- `overrides` (Block List, Max: 1) List of override configurations to apply to the ruleset. (see [below for nested schema](#nestedblock--rules--action_parameters--overrides))
- `phases` (Set of String) Point in the request/response lifecycle where the ruleset will be created. Available values: `ddos_l4`, `ddos_l7`, `http_custom_errors`, `http_log_custom_fields`, `http_request_cache_settings`, `http_request_firewall_custom`, `http_request_firewall_managed`, `http_request_late_transform`, `http_request_late_transform_managed`, `http_request_main`, `http_request_origin`, `http_request_dynamic_redirect`, `http_request_redirect`, `http_request_sanitize`, `http_request_transform`, `http_response_firewall_managed`, `http_response_headers_transform`, `http_response_headers_transform_managed`, `magic_transit`, `http_ratelimit`, `http_request_sbfm`, `http_config_settings`.
- `polish` (String) Apply options from the Polish feature of the Cloudflare Speed app.
- `products` (Set of String) Products to target with the actions. Available values: `bic`, `hot`, `ratelimit`, `securityLevel`, `uablock`, `waf`, `zonelockdown`.
- `request_fields` (Set of String) List of request headers to include as part of custom fields logging, in lowercase.
- `respect_strong_etags` (Boolean) Respect strong ETags.
- `response` (Block List) List of parameters that configure the response given to end users. (see [below for nested schema](#nestedblock--rules--action_parameters--response))
- `response_fields` (Set of String) List of response headers to include as part of custom fields logging, in lowercase.
- `rocket_loader` (Boolean) Turn on or off Cloudflare Rocket Loader in the Cloudflare Speed app.
- `rules` (Map of String) Map of managed WAF rule ID to comma-delimited string of ruleset rule IDs. Example: `rules = { "efb7b8c949ac4650a09736fc376e9aee" = "5de7edfa648c4d6891dc3e7f84534ffa,e3a567afc347477d9702d9047e97d760" }`.
- `ruleset` (String) Which ruleset ID to target.
- `rulesets` (Set of String) List of managed WAF rule IDs to target. Only valid when the `"action"` is set to skip.
- `security_level` (String) Control options for the Security Level feature from the Security app.
- `serve_stale` (Block List, Max: 1) List of serve stale parameters to apply to the request. (see [below for nested schema](#nestedblock--rules--action_parameters--serve_stale))
- `server_side_excludes` (Boolean) Turn on or off the Server Side Excludes feature of the Cloudflare Scrape Shield app.
- `sni` (Block List, Max: 1) List of properties to manange Server Name Indication. (see [below for nested schema](#nestedblock--rules--action_parameters--sni))
- `ssl` (String) Control options for the SSL feature of the Edge Certificates tab in the Cloudflare SSL/TLS app.
- `status_code` (Number) HTTP status code of the custom error response.
- `sxg` (Boolean) Turn on or off the SXG feature.
- `uri` (Block List, Max: 1) List of URI properties to configure for the ruleset rule when performing URL rewrite transformations. (see [below for nested schema](#nestedblock--rules--action_parameters--uri))
- `version` (String) Version of the ruleset to deploy.

<a id="nestedblock--rules--action_parameters--autominify"></a>
### Nested Schema for `rules.action_parameters.autominify`

Optional:

- `css` (Boolean) SSL minification.
- `html` (Boolean) HTML minification.
- `js` (Boolean) JS minification.


<a id="nestedblock--rules--action_parameters--browser_ttl"></a>
### Nested Schema for `rules.action_parameters.browser_ttl`

Required:

- `mode` (String) Mode of the browser TTL.

Optional:

- `default` (Number) Default browser TTL.


<a id="nestedblock--rules--action_parameters--cache_key"></a>
### Nested Schema for `rules.action_parameters.cache_key`

Optional:

- `cache_by_device_type` (Boolean) Cache by device type. Conflicts with "custom_key.user.device_type".
- `cache_deception_armor` (Boolean) Cache deception armor.
- `custom_key` (Block List, Max: 1) Custom key parameters for the request. (see [below for nested schema](#nestedblock--rules--action_parameters--cache_key--custom_key))
- `ignore_query_strings_order` (Boolean) Ignore query strings order.

<a id="nestedblock--rules--action_parameters--cache_key--custom_key"></a>
### Nested Schema for `rules.action_parameters.cache_key.custom_key`

Optional:

- `cookie` (Block List, Max: 1) Cookie parameters for the custom key. (see [below for nested schema](#nestedblock--rules--action_parameters--cache_key--custom_key--cookie))
- `header` (Block List, Max: 1) Header parameters for the custom key. (see [below for nested schema](#nestedblock--rules--action_parameters--cache_key--custom_key--header))
- `host` (Block List, Max: 1) Host parameters for the custom key. (see [below for nested schema](#nestedblock--rules--action_parameters--cache_key--custom_key--host))
- `query_string` (Block List, Max: 1) Query string parameters for the custom key. (see [below for nested schema](#nestedblock--rules--action_parameters--cache_key--custom_key--query_string))
- `user` (Block List, Max: 1) User parameters for the custom key. (see [below for nested schema](#nestedblock--rules--action_parameters--cache_key--custom_key--user))

<a id="nestedblock--rules--action_parameters--cache_key--custom_key--cookie"></a>
### Nested Schema for `rules.action_parameters.cache_key.custom_key.cookie`

Optional:

- `check_presence` (List of String) List of cookies to check for presence in the custom key.
- `include` (List of String) List of cookies to include in the custom key.


<a id="nestedblock--rules--action_parameters--cache_key--custom_key--header"></a>
### Nested Schema for `rules.action_parameters.cache_key.custom_key.header`

Optional:

- `check_presence` (List of String) List of headers to check for presence in the custom key.
- `exclude_origin` (Boolean) Exclude the origin header from the custom key.
- `include` (List of String) List of headers to include in the custom key.


<a id="nestedblock--rules--action_parameters--cache_key--custom_key--host"></a>
### Nested Schema for `rules.action_parameters.cache_key.custom_key.host`

Optional:

- `resolved` (Boolean) Resolve hostname to IP address.


<a id="nestedblock--rules--action_parameters--cache_key--custom_key--query_string"></a>
### Nested Schema for `rules.action_parameters.cache_key.custom_key.query_string`

Optional:

- `exclude` (List of String) List of query string parameters to exclude from the custom key. Conflicts with "include".
- `include` (List of String) List of query string parameters to include in the custom key. Conflicts with "exclude".


<a id="nestedblock--rules--action_parameters--cache_key--custom_key--user"></a>
### Nested Schema for `rules.action_parameters.cache_key.custom_key.user`

Optional:

- `device_type` (Boolean) Add device type to the custom key. Conflicts with "cache_key.cache_by_device_type".
- `geo` (Boolean) Add geo data to the custom key.
- `lang` (Boolean) Add language data to the custom key.




<a id="nestedblock--rules--action_parameters--edge_ttl"></a>
### Nested Schema for `rules.action_parameters.edge_ttl`

Required:

- `mode` (String) Mode of the edge TTL.

Optional:

- `default` (Number) Default edge TTL.
- `status_code_ttl` (Block List) Edge TTL for the status codes. (see [below for nested schema](#nestedblock--rules--action_parameters--edge_ttl--status_code_ttl))

<a id="nestedblock--rules--action_parameters--edge_ttl--status_code_ttl"></a>
### Nested Schema for `rules.action_parameters.edge_ttl.status_code_ttl`

Required:

- `value` (Number) Status code edge TTL value.

Optional:

- `status_code` (Number) Status code for which the edge TTL is applied. Conflicts with "status_code_range".
- `status_code_range` (Block List) Status code range for which the edge TTL is applied. Conflicts with "status_code". (see [below for nested schema](#nestedblock--rules--action_parameters--edge_ttl--status_code_ttl--status_code_range))

<a id="nestedblock--rules--action_parameters--edge_ttl--status_code_ttl--status_code_range"></a>
### Nested Schema for `rules.action_parameters.edge_ttl.status_code_ttl.status_code_range`

Optional:

- `from` (Number) From status code.
- `to` (Number) To status code.




<a id="nestedblock--rules--action_parameters--from_list"></a>
### Nested Schema for `rules.action_parameters.from_list`

Required:

- `key` (String) Expression to use for the list lookup.
- `name` (String) Name of the list.


<a id="nestedblock--rules--action_parameters--from_value"></a>
### Nested Schema for `rules.action_parameters.from_value`

Optional:

- `preserve_query_string` (Boolean) Preserve query string for redirect URL.
- `status_code` (Number) Status code for redirect.
- `target_url` (Block List, Max: 1) Target URL for redirect. (see [below for nested schema](#nestedblock--rules--action_parameters--from_value--target_url))

<a id="nestedblock--rules--action_parameters--from_value--target_url"></a>
### Nested Schema for `rules.action_parameters.from_value.target_url`

Optional:

- `expression` (String) Use a value dynamically determined by the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions. Conflicts with `"value"`.
- `value` (String) Static value to provide as the HTTP request header value. Conflicts with `"expression"`.



<a id="nestedblock--rules--action_parameters--headers"></a>
### Nested Schema for `rules.action_parameters.headers`

Optional:

- `expression` (String) Use a value dynamically determined by the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions. Conflicts with `"value"`.
- `name` (String) Name of the HTTP request header to target.
- `operation` (String) Action to perform on the HTTP request header. Available values: `remove`, `set`.
- `value` (String) Static value to provide as the HTTP request header value. Conflicts with `"expression"`.


<a id="nestedblock--rules--action_parameters--matched_data"></a>
### Nested Schema for `rules.action_parameters.matched_data`

Optional:

- `public_key` (String) Public key to use within WAF Ruleset payload logging to view the HTTP request parameters. You can generate a public key [using the `matched-data-cli` command-line tool](https://developers.cloudflare.com/waf/managed-rulesets/payload-logging/command-line/generate-key-pair) or [in the Cloudflare dashboard](https://developers.cloudflare.com/waf/managed-rulesets/payload-logging/configure).


<a id="nestedblock--rules--action_parameters--origin"></a>
### Nested Schema for `rules.action_parameters.origin`

Optional:

- `host` (String) Origin Hostname where request is sent.
- `port` (Number) Origin Port where request is sent.


<a id="nestedblock--rules--action_parameters--overrides"></a>
### Nested Schema for `rules.action_parameters.overrides`

Optional:

- `action` (String) Action to perform in the rule-level override. Available values: `block`, `challenge`, `ddos_dynamic`, `execute`, `force_connection_close`, `js_challenge`, `log`, `log_custom_field`, `managed_challenge`, `redirect`, `rewrite`, `route`, `score`, `set_cache_settings`, `set_config`, `serve_error`, `skip`.
- `categories` (Block List) List of tag-based overrides. (see [below for nested schema](#nestedblock--rules--action_parameters--overrides--categories))
- `enabled` (Boolean, Deprecated) Defines if the current ruleset-level override enables or disables the ruleset.
- `rules` (Block List) List of rule-based overrides. (see [below for nested schema](#nestedblock--rules--action_parameters--overrides--rules))
- `sensitivity_level` (String) Sensitivity level to override for all ruleset rules. Available values: `default`, `medium`, `low`, `eoff`.
- `status` (String) Defines if the current ruleset-level override enables or disables the ruleset. Available values: `enabled`, `disabled`. Defaults to `""`.

<a id="nestedblock--rules--action_parameters--overrides--categories"></a>
### Nested Schema for `rules.action_parameters.overrides.categories`

Optional:

- `action` (String) Action to perform in the tag-level override. Available values: `block`, `challenge`, `ddos_dynamic`, `execute`, `force_connection_close`, `js_challenge`, `log`, `log_custom_field`, `managed_challenge`, `redirect`, `rewrite`, `route`, `score`, `set_cache_settings`, `set_config`, `serve_error`, `skip`.
- `category` (String) Tag name to apply the ruleset rule override to.
- `enabled` (Boolean, Deprecated) Defines if the current tag-level override enables or disables the ruleset rules with the specified tag.
- `status` (String) Defines if the current tag-level override enables or disables the ruleset rules with the specified tag. Available values: `enabled`, `disabled`. Defaults to `""`.


<a id="nestedblock--rules--action_parameters--overrides--rules"></a>
### Nested Schema for `rules.action_parameters.overrides.rules`

Optional:

- `action` (String) Action to perform in the rule-level override. Available values: `block`, `challenge`, `ddos_dynamic`, `execute`, `force_connection_close`, `js_challenge`, `log`, `log_custom_field`, `managed_challenge`, `redirect`, `rewrite`, `route`, `score`, `set_cache_settings`, `set_config`, `serve_error`, `skip`.
- `enabled` (Boolean, Deprecated) Defines if the current rule-level override enables or disables the rule.
- `id` (String) Rule ID to apply the override to.
- `score_threshold` (Number) Anomaly score threshold to apply in the ruleset rule override. Only applicable to modsecurity-based rulesets.
- `sensitivity_level` (String) Sensitivity level for a ruleset rule override.
- `status` (String) Defines if the current rule-level override enables or disables the rule. Available values: `enabled`, `disabled`. Defaults to `""`.



<a id="nestedblock--rules--action_parameters--response"></a>
### Nested Schema for `rules.action_parameters.response`

Optional:

- `content` (String) Body content to include in the response.
- `content_type` (String) HTTP content type to send in the response.
- `status_code` (Number) HTTP status code to send in the response.


<a id="nestedblock--rules--action_parameters--serve_stale"></a>
### Nested Schema for `rules.action_parameters.serve_stale`

Optional:

- `disable_stale_while_updating` (Boolean) Disable stale while updating.


<a id="nestedblock--rules--action_parameters--sni"></a>
### Nested Schema for `rules.action_parameters.sni`

Optional:

- `value` (String) Value to define for SNI.


<a id="nestedblock--rules--action_parameters--uri"></a>
### Nested Schema for `rules.action_parameters.uri`

Optional:

- `origin` (Boolean)
- `path` (Block List, Max: 1) URI path configuration when performing a URL rewrite. (see [below for nested schema](#nestedblock--rules--action_parameters--uri--path))
- `query` (Block List, Max: 1) Query string configuration when performing a URL rewrite. (see [below for nested schema](#nestedblock--rules--action_parameters--uri--query))

<a id="nestedblock--rules--action_parameters--uri--path"></a>
### Nested Schema for `rules.action_parameters.uri.path`

Optional:

- `expression` (String) Expression that defines the updated (dynamic) value of the URI path or query string component. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.
- `value` (String) Static string value of the updated URI path or query string component.


<a id="nestedblock--rules--action_parameters--uri--query"></a>
### Nested Schema for `rules.action_parameters.uri.query`

Optional:

- `expression` (String) Expression that defines the updated (dynamic) value of the URI path or query string component. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.
- `value` (String) Static string value of the updated URI path or query string component.




<a id="nestedblock--rules--exposed_credential_check"></a>
### Nested Schema for `rules.exposed_credential_check`

Optional:

- `password_expression` (String) Firewall Rules expression language based on Wireshark display filters for where to check for the "password" value. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language).
- `username_expression` (String) Firewall Rules expression language based on Wireshark display filters for where to check for the "username" value. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language).


<a id="nestedblock--rules--logging"></a>
### Nested Schema for `rules.logging`

Optional:

- `enabled` (Boolean, Deprecated) Override the default logging behavior when a rule is matched.
- `status` (String) Override the default logging behavior when a rule is matched. Available values: `enabled`, `disabled`. Defaults to `""`.


<a id="nestedblock--rules--ratelimit"></a>
### Nested Schema for `rules.ratelimit`

Optional:

- `characteristics` (Set of String) List of parameters that define how Cloudflare tracks the request rate for this rule.
- `counting_expression` (String) Criteria for counting HTTP requests to trigger the Rate Limiting action. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.
- `mitigation_timeout` (Number) Once the request rate is reached, the Rate Limiting rule blocks further requests for the period of time defined in this field.
- `period` (Number) The period of time to consider (in seconds) when evaluating the request rate.
- `requests_per_period` (Number) The number of requests over the period of time that will trigger the Rate Limiting rule.
- `requests_to_origin` (Boolean) Whether to include requests to origin within the Rate Limiting count.
- `score_per_period` (Number) The maximum aggregate score over the period of time that will trigger Rate Limiting rule.
- `score_response_header_name` (String) Name of HTTP header in the response, set by the origin server, with the score for the current request.

## Import

Import is supported using the following syntax:

```shell
# Import an account scoped Ruleset configuration.
$ terraform import cloudflare_ruleset.example account/<account_id>/<ruleset_id>

# Import a zone scoped Ruleset configuration.
$ terraform import cloudflare_ruleset.example zone/<zone_id>/<ruleset_id>
```
