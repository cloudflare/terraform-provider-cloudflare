---
page_title: "cloudflare_ruleset Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_ruleset (Resource)



## Example Usage

```terraform
# Magic Transit
resource "cloudflare_ruleset" "magic_transit_example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "account magic transit"
  description = "example magic transit ruleset description"
  kind        = "root"
  phase       = "magic_transit"

  rules = [{
    action      = "allow"
    expression  = "tcp.dstport in { 32768..65535 }"
    description = "Allow TCP Ephemeral Ports"
  }]
}

# Zone-level WAF Managed Ruleset
resource "cloudflare_ruleset" "zone_level_managed_waf" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "managed WAF"
  description = "managed WAF ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_managed"

  rules = [{
    action = "execute"
    action_parameters = {
    id = "efb7b8c949ac4650a09736fc376e9aee"
  }
    expression  = "(http.host eq \"example.host.com\")"
    description = "Execute Cloudflare Managed Ruleset on my zone-level phase entry point ruleset"
    enabled     = true
  }]
}

# Zone-level WAF with tag-based overrides
resource "cloudflare_ruleset" "zone_level_managed_waf_with_category_based_overrides" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "managed WAF with tag-based overrides"
  description = "managed WAF with tag-based overrides ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_managed"

  rules = [{
    action = "execute"
    action_parameters = {
    id = "efb7b8c949ac4650a09736fc376e9aee"
      overrides = { categories = [{
          category = "wordpress"
          action   = "block"
          enabled  = true
          },
          {
            category = "joomla"
            action   = "block"
            enabled  = true
        }] }
  }

    expression  = "(http.host eq \"example.host.com\")"
    description = "overrides to only enable wordpress rules to block"
    enabled     = false
  }]
}

# Rewrite the URI path component to a static path
resource "cloudflare_ruleset" "transform_uri_rule_path" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "transform rule for URI path"
  description = "change the URI path to a new static path"
  kind        = "zone"
  phase       = "http_request_transform"

  rules = [{
    action = "rewrite"
    action_parameters = {
    uri = {
    path = {
    value = "/my-new-route"
  }
  }
  }

    expression  = "(http.host eq \"example.com\" and http.request.uri.path eq \"/old-path\")"
    description = "example URI path transform rule"
    enabled     = true
  }]
}

# Rewrite the URI query component to a static query
resource "cloudflare_ruleset" "transform_uri_rule_query" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "transform rule for URI query parameter"
  description = "change the URI query to a new static query"
  kind        = "zone"
  phase       = "http_request_transform"

  rules = [{
    action = "rewrite"
    action_parameters = {
    uri = {
    query = {
    value = "old=new_again"
  }
  }
  }

    expression  = "(http.host eq \"example.host.com\")"
    description = "URI transformation query example"
    enabled     = true
  }]
}

# Rewrite HTTP headers to a modified values
resource "cloudflare_ruleset" "transform_uri_http_headers" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "transform rule for HTTP headers"
  description = "modify HTTP headers before reaching origin"
  kind        = "zone"
  phase       = "http_request_late_transform"

  rules = [{
    action = "rewrite"
    action_parameters = {
    headers = [{
        name      = "example-http-header-1"
        operation = "set"
        value     = "my-http-header-value-1"
        },
        {
          name       = "example-http-header-2"
          operation  = "set"
          expression = "cf.zone.name"
        },
        {
          name      = "example-http-header-3-to-remove"
          operation = "remove"
      }]
  }

    expression  = "(http.host eq \"example.host.com\")"
    description = "example request header transform rule"
    enabled     = false
  }]
}

# HTTP rate limit for an API route
resource "cloudflare_ruleset" "rate_limiting_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "restrict API requests count"
  description = "apply HTTP rate limiting for a route"
  kind        = "zone"
  phase       = "http_ratelimit"

  rules = [{
    action = "block"
    ratelimit = {
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
  }]
}

# Change origin for an API route
resource "cloudflare_ruleset" "http_origin_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "Change to some origin"
  description = "Change origin for a route"
  kind        = "zone"
  phase       = "http_request_origin"

  rules = [{
    action = "route"
    action_parameters = {
    host_header = "some.host"
      origin = {
    host = "some.host"
        port = 80
  }
  }
    expression  = "(http.request.uri.path matches \"^/api/\")"
    description = "change origin to some.host"
    enabled     = true
  }]
}

# Custom fields logging
resource "cloudflare_ruleset" "custom_fields_logging_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "log custom fields"
  description = "add custom fields to logging"
  kind        = "zone"
  phase       = "http_log_custom_fields"

  rules = [{
    action = "log_custom_field"
    action_parameters = {
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
  }]
}

# Custom cache keys + settings
resource "cloudflare_ruleset" "cache_settings_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "set cache settings"
  description = "set cache settings for the request"
  kind        = "zone"
  phase       = "http_request_cache_settings"

  rules = [{
    action = "set_cache_settings"
    action_parameters = {
    edge_ttl = {
    mode    = "override_origin"
        default = 60
        status_code_ttl = [{
          status_code = 200
          value       = 50
          },
          {
            status_code_range = [{
              from = 201
              to   = 300
            }]
            value = 30
        }]
  }
      browser_ttl = {
    mode = "respect_origin"
  }
      serve_stale = {
    disable_stale_while_updating = true
  }
      respect_strong_etags = true
      cache_key = {
    ignore_query_strings_order = false
        cache_deception_armor      = true
        custom_key = {
    query_string = {
    exclude = ["*"]
  }
          header = {
    include        = ["habc", "hdef"]
            check_presence = ["habc_t", "hdef_t"]
            exclude_origin = true
  }
          cookie = {
    include        = ["cabc", "cdef"]
            check_presence = ["cabc_t", "cdef_t"]
  }
          user = {
    device_type = true
            geo         = false
  }
          host = {
    resolved = true
  }
  }
  }
      origin_error_page_passthru = false
  }
    expression  = "(http.host eq \"example.host.com\")"
    description = "set cache settings rule"
    enabled     = true
  }]
}

# Redirects based on a List resource
resource "cloudflare_ruleset" "redirect_from_list_example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "redirects"
  description = "Redirect ruleset"
  kind        = "root"
  phase       = "http_request_redirect"

  rules = [{
    action = "redirect"
    action_parameters = {
    from_list = {
    name = "redirect_list"
        key  = "http.request.full_uri"
  }
  }
    expression  = "http.request.full_uri in $redirect_list"
    description = "Apply redirects from redirect_list"
    enabled     = true
  }]
}

# Dynamic Redirects from value resource
resource "cloudflare_ruleset" "redirect_from_value_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "redirects"
  description = "Redirect ruleset"
  kind        = "zone"
  phase       = "http_request_dynamic_redirect"

  rules = [{
    action = "redirect"
    action_parameters = {
    from_value = {
    status_code = 301
        target_url = {
    value = "some_host.com"
  }
        preserve_query_string = true
  }
  }
    expression  = "(http.request.uri.path matches \"^/api/\")"
    description = "Apply redirect from value"
    enabled     = true
  }]
}

# Serve some custom error response
resource "cloudflare_ruleset" "http_custom_error_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "Serve some error response"
  description = "Serve some error response"
  kind        = "zone"
  phase       = "http_custom_errors"
  rules = [{
    action = "serve_error"
    action_parameters = {
    content      = "some error html"
      content_type = "text/html"
      status_code  = "530"
  }
    expression  = "(http.request.uri.path matches \"^/api/\")"
    description = "serve some error response"
    enabled     = true
  }]
}

# Set Configuration Rules for an API route
resource "cloudflare_ruleset" "http_config_rules_example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "set config rules"
  description = "set config rules for request"
  kind        = "zone"
  phase       = "http_config_settings"

  rules = [{
    action = "set_config"
    action_parameters = {
    email_obfuscation = true
      bic               = true
  }
    expression  = "(http.request.uri.path matches \"^/api/\")"
    description = "set config rules for matching request"
    enabled     = true
  }]
}

# Set compress algorithm for response.
resource "cloudflare_ruleset" "response_compress_brotli_html" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "Brotli response compression for HTML"
  description = "Response compression ruleset"
  kind        = "zone"
  phase       = "http_response_compression"

  rules = [{
    action = "compress_response"
    action_parameters = {
    algorithms = [{
        name = "brotli"
        },
        {
          name = "auto"
      }]
  }
    expression  = "http.response.content_type.media_type == \"text/html\""
    description = "Prefer brotli compression for HTML"
    enabled     = true
  }]
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `kind` (String) The kind of the ruleset.
- `name` (String) The human-readable name of the ruleset.
- `phase` (String) The phase of the ruleset.
- `rules` (Attributes List) The list of rules in the ruleset. (see [below for nested schema](#nestedatt--rules))

### Optional

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `description` (String) An informative description of the ruleset.
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

### Read-Only

- `id` (String) The unique ID of the ruleset.
- `last_updated` (String) The timestamp of when the ruleset was last modified.
- `version` (String) The version of the ruleset.

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

- `action` (String) The action to perform when the rule matches.
- `action_parameters` (Attributes) The parameters configuring the rule's action. (see [below for nested schema](#nestedatt--rules--action_parameters))
- `description` (String) An informative description of the rule.
- `enabled` (Boolean) Whether the rule should be executed.
- `expression` (String) The expression defining which traffic will match the rule.
- `logging` (Attributes) An object configuring the rule's logging behavior. (see [below for nested schema](#nestedatt--rules--logging))
- `ref` (String) The reference of the rule (the rule ID by default).

Read-Only:

- `categories` (List of String) The categories of the rule.
- `id` (String) The unique ID of the rule.
- `last_updated` (String) The timestamp of when the rule was last modified.
- `version` (String) The version of the rule.

<a id="nestedatt--rules--action_parameters"></a>
### Nested Schema for `rules.action_parameters`

Optional:

- `additional_cacheable_ports` (List of Number) List of additional ports that caching can be enabled on.
- `algorithms` (Attributes List) Custom order for compression algorithms. (see [below for nested schema](#nestedatt--rules--action_parameters--algorithms))
- `automatic_https_rewrites` (Boolean) Turn on or off Automatic HTTPS Rewrites.
- `autominify` (Attributes) Select which file extensions to minify automatically. (see [below for nested schema](#nestedatt--rules--action_parameters--autominify))
- `bic` (Boolean) Turn on or off Browser Integrity Check.
- `browser_ttl` (Attributes) Specify how long client browsers should cache the response. Cloudflare cache purge will not purge content cached on client browsers, so high browser TTLs may lead to stale content. (see [below for nested schema](#nestedatt--rules--action_parameters--browser_ttl))
- `cache` (Boolean) Mark whether the request’s response from origin is eligible for caching. Caching itself will still depend on the cache-control header and your other caching configurations.
- `cache_key` (Attributes) Define which components of the request are included or excluded from the cache key Cloudflare uses to store the response in cache. (see [below for nested schema](#nestedatt--rules--action_parameters--cache_key))
- `cache_reserve` (Attributes) Mark whether the request's response from origin is eligible for  Cache Reserve (requires a Cache Reserve add-on plan). (see [below for nested schema](#nestedatt--rules--action_parameters--cache_reserve))
- `content` (String) Error response content.
- `content_type` (String) Content-type header to set with the response.
- `cookie_fields` (Attributes List) The cookie fields to log. (see [below for nested schema](#nestedatt--rules--action_parameters--cookie_fields))
- `disable_apps` (Boolean) Turn off all active Cloudflare Apps.
- `disable_rum` (Boolean) Turn off Real User Monitoring (RUM).
- `disable_zaraz` (Boolean) Turn off Zaraz.
- `edge_ttl` (Attributes) TTL (Time to Live) specifies the maximum time to cache a resource in the Cloudflare edge network. (see [below for nested schema](#nestedatt--rules--action_parameters--edge_ttl))
- `email_obfuscation` (Boolean) Turn on or off Email Obfuscation.
- `fonts` (Boolean) Turn on or off Cloudflare Fonts.
- `from_list` (Attributes) Serve a redirect based on a bulk list lookup. (see [below for nested schema](#nestedatt--rules--action_parameters--from_list))
- `from_value` (Attributes) Serve a redirect based on the request properties. (see [below for nested schema](#nestedatt--rules--action_parameters--from_value))
- `headers` (Attributes Map) Map of request headers to modify. (see [below for nested schema](#nestedatt--rules--action_parameters--headers))
- `host_header` (String) Rewrite the HTTP Host header.
- `hotlink_protection` (Boolean) Turn on or off the Hotlink Protection.
- `id` (String) The ID of the ruleset to execute.
- `increment` (Number) Increment contains the delta to change the score and can be either positive or negative.
- `matched_data` (Attributes) The configuration to use for matched data logging. (see [below for nested schema](#nestedatt--rules--action_parameters--matched_data))
- `mirage` (Boolean) Turn on or off Mirage.
- `opportunistic_encryption` (Boolean) Turn on or off Opportunistic Encryption.
- `origin` (Attributes) Override the IP/TCP destination. (see [below for nested schema](#nestedatt--rules--action_parameters--origin))
- `origin_cache_control` (Boolean) When enabled, Cloudflare will aim to strictly adhere to RFC 7234.
- `origin_error_page_passthru` (Boolean) Generate Cloudflare error pages from issues sent from the origin server. When on, error pages will trigger for issues from the origin
- `overrides` (Attributes) A set of overrides to apply to the target ruleset. (see [below for nested schema](#nestedatt--rules--action_parameters--overrides))
- `phases` (List of String) A list of phases to skip the execution of. This option is incompatible with the ruleset and rulesets options.
- `polish` (String) Configure the Polish level.
- `products` (List of String) A list of legacy security products to skip the execution of.
- `read_timeout` (Number) Define a timeout value between two successive read operations to your origin server. Historically, the timeout value between two read options from Cloudflare to an origin server is 100 seconds. If you are attempting to reduce HTTP 524 errors because of timeouts from an origin server, try increasing this timeout value.
- `request_fields` (Attributes List) The request fields to log. (see [below for nested schema](#nestedatt--rules--action_parameters--request_fields))
- `respect_strong_etags` (Boolean) Specify whether or not Cloudflare should respect strong ETag (entity tag) headers. When off, Cloudflare converts strong ETag headers to weak ETag headers.
- `response` (Attributes) The response to show when the block is applied. (see [below for nested schema](#nestedatt--rules--action_parameters--response))
- `response_fields` (Attributes List) The response fields to log. (see [below for nested schema](#nestedatt--rules--action_parameters--response_fields))
- `rocket_loader` (Boolean) Turn on or off Rocket Loader
- `rules` (Map of List of String) A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the execution of. This option is incompatible with the ruleset option.
- `ruleset` (String) A ruleset to skip the execution of. This option is incompatible with the rulesets, rules and phases options.
- `rulesets` (List of String) A list of ruleset IDs to skip the execution of. This option is incompatible with the ruleset and phases options.
- `security_level` (String) Configure the Security Level.
- `serve_stale` (Attributes) Define if Cloudflare should serve stale content while getting the latest content from the origin. If on, Cloudflare will not serve stale content while getting the latest content from the origin. (see [below for nested schema](#nestedatt--rules--action_parameters--serve_stale))
- `server_side_excludes` (Boolean) Turn on or off Server Side Excludes.
- `sni` (Attributes) Override the Server Name Indication (SNI). (see [below for nested schema](#nestedatt--rules--action_parameters--sni))
- `ssl` (String) Configure the SSL level.
- `status_code` (Number) The status code to use for the error.
- `sxg` (Boolean) Turn on or off Signed Exchanges (SXG).
- `uri` (Attributes) URI to rewrite the request to. (see [below for nested schema](#nestedatt--rules--action_parameters--uri))

<a id="nestedatt--rules--action_parameters--algorithms"></a>
### Nested Schema for `rules.action_parameters.algorithms`

Optional:

- `name` (String) Name of compression algorithm to enable.


<a id="nestedatt--rules--action_parameters--autominify"></a>
### Nested Schema for `rules.action_parameters.autominify`

Optional:

- `css` (Boolean) Minify CSS files.
- `html` (Boolean) Minify HTML files.
- `js` (Boolean) Minify JS files.


<a id="nestedatt--rules--action_parameters--browser_ttl"></a>
### Nested Schema for `rules.action_parameters.browser_ttl`

Required:

- `mode` (String) Determines which browser ttl mode to use.

Optional:

- `default` (Number) The TTL (in seconds) if you choose override_origin mode.


<a id="nestedatt--rules--action_parameters--cache_key"></a>
### Nested Schema for `rules.action_parameters.cache_key`

Optional:

- `cache_by_device_type` (Boolean) Separate cached content based on the visitor’s device type
- `cache_deception_armor` (Boolean) Protect from web cache deception attacks while allowing static assets to be cached
- `custom_key` (Attributes) Customize which components of the request are included or excluded from the cache key. (see [below for nested schema](#nestedatt--rules--action_parameters--cache_key--custom_key))
- `ignore_query_strings_order` (Boolean) Treat requests with the same query parameters the same, regardless of the order those query parameters are in. A value of true ignores the query strings' order.

<a id="nestedatt--rules--action_parameters--cache_key--custom_key"></a>
### Nested Schema for `rules.action_parameters.cache_key.ignore_query_strings_order`

Optional:

- `cookie` (Attributes) The cookies to include in building the cache key. (see [below for nested schema](#nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--cookie))
- `header` (Attributes) The header names and values to include in building the cache key. (see [below for nested schema](#nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--header))
- `host` (Attributes) Whether to use the original host or the resolved host in the cache key. (see [below for nested schema](#nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--host))
- `query_string` (Attributes) Use the presence or absence of parameters in the query string to build the cache key. (see [below for nested schema](#nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--query_string))
- `user` (Attributes) Characteristics of the request user agent used in building the cache key. (see [below for nested schema](#nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--user))

<a id="nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--cookie"></a>
### Nested Schema for `rules.action_parameters.cache_key.ignore_query_strings_order.cookie`

Optional:

- `check_presence` (List of String) Checks for the presence of these cookie names. The presence of these cookies is used in building the cache key.
- `include` (List of String) Include these cookies' names and their values.


<a id="nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--header"></a>
### Nested Schema for `rules.action_parameters.cache_key.ignore_query_strings_order.header`

Optional:

- `check_presence` (List of String) Checks for the presence of these header names. The presence of these headers is used in building the cache key.
- `contains` (Map of List of String) For each header name and list of values combination, check if the request header contains any of the values provided. The presence of the request header and whether any of the values provided are contained in the request header value is used in building the cache key.
- `exclude_origin` (Boolean) Whether or not to include the origin header. A value of true will exclude the origin header in the cache key.
- `include` (List of String) Include these headers' names and their values.


<a id="nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--host"></a>
### Nested Schema for `rules.action_parameters.cache_key.ignore_query_strings_order.host`

Optional:

- `resolved` (Boolean) Use the resolved host in the cache key. A value of true will use the resolved host, while a value or false will use the original host.


<a id="nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--query_string"></a>
### Nested Schema for `rules.action_parameters.cache_key.ignore_query_strings_order.query_string`

Optional:

- `exclude` (Attributes) build the cache key using all query string parameters EXCECPT these excluded parameters (see [below for nested schema](#nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--query_string--exclude))
- `include` (Attributes) build the cache key using a list of query string parameters that ARE in the request. (see [below for nested schema](#nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--query_string--include))

<a id="nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--query_string--exclude"></a>
### Nested Schema for `rules.action_parameters.cache_key.ignore_query_strings_order.query_string.include`

Optional:

- `all` (Boolean) Exclude all query string parameters from use in building the cache key.
- `list` (List of String) A list of query string parameters NOT used to build the cache key. All parameters present in the request but missing in this list will be used to build the cache key.


<a id="nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--query_string--include"></a>
### Nested Schema for `rules.action_parameters.cache_key.ignore_query_strings_order.query_string.include`

Optional:

- `all` (Boolean) Use all query string parameters in the cache key.
- `list` (List of String) A list of query string parameters used to build the cache key.



<a id="nestedatt--rules--action_parameters--cache_key--ignore_query_strings_order--user"></a>
### Nested Schema for `rules.action_parameters.cache_key.ignore_query_strings_order.user`

Optional:

- `device_type` (Boolean) Use the user agent's device type in the cache key.
- `geo` (Boolean) Use the user agents's country in the cache key.
- `lang` (Boolean) Use the user agent's language in the cache key.




<a id="nestedatt--rules--action_parameters--cache_reserve"></a>
### Nested Schema for `rules.action_parameters.cache_reserve`

Required:

- `eligible` (Boolean) Determines whether cache reserve is enabled. If this is true and a request meets eligibility criteria, Cloudflare will write the resource to cache reserve.

Optional:

- `minimum_file_size` (Number) The minimum file size eligible for store in cache reserve.


<a id="nestedatt--rules--action_parameters--cookie_fields"></a>
### Nested Schema for `rules.action_parameters.cookie_fields`

Required:

- `name` (String) The name of the field.


<a id="nestedatt--rules--action_parameters--edge_ttl"></a>
### Nested Schema for `rules.action_parameters.edge_ttl`

Required:

- `mode` (String) edge ttl options

Optional:

- `default` (Number) The TTL (in seconds) if you choose override_origin mode.
- `status_code_ttl` (Attributes List) List of single status codes, or status code ranges to apply the selected mode (see [below for nested schema](#nestedatt--rules--action_parameters--edge_ttl--status_code_ttl))

<a id="nestedatt--rules--action_parameters--edge_ttl--status_code_ttl"></a>
### Nested Schema for `rules.action_parameters.edge_ttl.status_code_ttl`

Required:

- `value` (Number) Time to cache a response (in seconds). A value of 0 is equivalent to setting the Cache-Control header with the value "no-cache". A value of -1 is equivalent to setting Cache-Control header with the value of "no-store".

Optional:

- `status_code` (Number) Set the ttl for responses with this specific status code
- `status_code_range` (Attributes) The range of status codes used to apply the selected mode. (see [below for nested schema](#nestedatt--rules--action_parameters--edge_ttl--status_code_ttl--status_code_range))

<a id="nestedatt--rules--action_parameters--edge_ttl--status_code_ttl--status_code_range"></a>
### Nested Schema for `rules.action_parameters.edge_ttl.status_code_ttl.status_code_range`

Optional:

- `from` (Number) response status code lower bound
- `to` (Number) response status code upper bound




<a id="nestedatt--rules--action_parameters--from_list"></a>
### Nested Schema for `rules.action_parameters.from_list`

Optional:

- `key` (String) Expression that evaluates to the list lookup key.
- `name` (String) The name of the list to match against.


<a id="nestedatt--rules--action_parameters--from_value"></a>
### Nested Schema for `rules.action_parameters.from_value`

Optional:

- `preserve_query_string` (Boolean) Keep the query string of the original request.
- `status_code` (Number) The status code to be used for the redirect.
- `target_url` (Attributes) The URL to redirect the request to. (see [below for nested schema](#nestedatt--rules--action_parameters--from_value--target_url))

<a id="nestedatt--rules--action_parameters--from_value--target_url"></a>
### Nested Schema for `rules.action_parameters.from_value.target_url`

Optional:

- `expression` (String) An expression to evaluate to get the URL to redirect the request to.
- `value` (String) The URL to redirect the request to.



<a id="nestedatt--rules--action_parameters--headers"></a>
### Nested Schema for `rules.action_parameters.headers`

Optional:

- `expression` (String) Expression for the header value.
- `operation` (String)
- `value` (String) Static value for the header.


<a id="nestedatt--rules--action_parameters--matched_data"></a>
### Nested Schema for `rules.action_parameters.matched_data`

Required:

- `public_key` (String) The public key to encrypt matched data logs with.


<a id="nestedatt--rules--action_parameters--origin"></a>
### Nested Schema for `rules.action_parameters.origin`

Optional:

- `host` (String) Override the resolved hostname.
- `port` (Number) Override the destination port.


<a id="nestedatt--rules--action_parameters--overrides"></a>
### Nested Schema for `rules.action_parameters.overrides`

Optional:

- `action` (String) An action to override all rules with. This option has lower precedence than rule and category overrides.
- `categories` (Attributes List) A list of category-level overrides. This option has the second-highest precedence after rule-level overrides. (see [below for nested schema](#nestedatt--rules--action_parameters--overrides--categories))
- `enabled` (Boolean) Whether to enable execution of all rules. This option has lower precedence than rule and category overrides.
- `rules` (Attributes List) A list of rule-level overrides. This option has the highest precedence. (see [below for nested schema](#nestedatt--rules--action_parameters--overrides--rules))
- `sensitivity_level` (String) A sensitivity level to set for all rules. This option has lower precedence than rule and category overrides and is only applicable for DDoS phases.

<a id="nestedatt--rules--action_parameters--overrides--categories"></a>
### Nested Schema for `rules.action_parameters.overrides.sensitivity_level`

Required:

- `category` (String) The name of the category to override.

Optional:

- `action` (String) The action to override rules in the category with.
- `enabled` (Boolean) Whether to enable execution of rules in the category.
- `sensitivity_level` (String) The sensitivity level to use for rules in the category.


<a id="nestedatt--rules--action_parameters--overrides--rules"></a>
### Nested Schema for `rules.action_parameters.overrides.sensitivity_level`

Required:

- `id` (String) The ID of the rule to override.

Optional:

- `action` (String) The action to override the rule with.
- `enabled` (Boolean) Whether to enable execution of the rule.
- `score_threshold` (Number) The score threshold to use for the rule.
- `sensitivity_level` (String) The sensitivity level to use for the rule.



<a id="nestedatt--rules--action_parameters--request_fields"></a>
### Nested Schema for `rules.action_parameters.request_fields`

Required:

- `name` (String) The name of the field.


<a id="nestedatt--rules--action_parameters--response"></a>
### Nested Schema for `rules.action_parameters.response`

Required:

- `content` (String) The content to return.
- `content_type` (String) The type of the content to return.
- `status_code` (Number) The status code to return.


<a id="nestedatt--rules--action_parameters--response_fields"></a>
### Nested Schema for `rules.action_parameters.response_fields`

Required:

- `name` (String) The name of the field.


<a id="nestedatt--rules--action_parameters--serve_stale"></a>
### Nested Schema for `rules.action_parameters.serve_stale`

Required:

- `disable_stale_while_updating` (Boolean) Defines whether Cloudflare should serve stale content while updating. If true, Cloudflare will not serve stale content while getting the latest content from the origin.


<a id="nestedatt--rules--action_parameters--sni"></a>
### Nested Schema for `rules.action_parameters.sni`

Required:

- `value` (String) The SNI override.


<a id="nestedatt--rules--action_parameters--uri"></a>
### Nested Schema for `rules.action_parameters.uri`

Optional:

- `path` (Attributes) Path portion rewrite. (see [below for nested schema](#nestedatt--rules--action_parameters--uri--path))
- `query` (Attributes) Query portion rewrite. (see [below for nested schema](#nestedatt--rules--action_parameters--uri--query))

<a id="nestedatt--rules--action_parameters--uri--path"></a>
### Nested Schema for `rules.action_parameters.uri.query`

Optional:

- `expression` (String) Expression to evaluate for the replacement value.
- `value` (String) Predefined replacement value.


<a id="nestedatt--rules--action_parameters--uri--query"></a>
### Nested Schema for `rules.action_parameters.uri.query`

Optional:

- `expression` (String) Expression to evaluate for the replacement value.
- `value` (String) Predefined replacement value.




<a id="nestedatt--rules--logging"></a>
### Nested Schema for `rules.logging`

Required:

- `enabled` (Boolean) Whether to generate a log when the rule matches.

## Import

Import is supported using the following syntax:

```shell
# Import an account scoped Ruleset configuration.
$ terraform import cloudflare_ruleset.example account/<account_id>/<ruleset_id>

# Import a zone scoped Ruleset configuration.
$ terraform import cloudflare_ruleset.example zone/<zone_id>/<ruleset_id>
```
