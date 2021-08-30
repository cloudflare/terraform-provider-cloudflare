---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_ruleset"
sidebar_current: "docs-cloudflare-resource-ruleset"
description: |-
  Provides a resource which manages Cloudflare rulesets.
---

# cloudflare_ruleset

The [Cloudflare Ruleset Engine](https://developers.cloudflare.com/firewall/cf-rulesets)
allows you to create and deploy rules and rulesets.
The engine syntax, inspired by the Wireshark Display Filter language, is the
same syntax used in custom Firewall Rules. Cloudflare uses the Ruleset Engine
in different products, allowing you to configure several products using the same
basic syntax.

## Example Usage

```hcl
# Magic Transit
resource "cloudflare_ruleset" "magic_transit_example" {
  account_id  = "d41d8cd98f00b204e9800998ecf8427e"
  name        = "account magic transit"
  description = "example magic transit ruleset description"
  kind        = "root"
  phase       = "magic_transit"

  rules {
    action = "allow"
    expression = "tcp.dstport in { 32768..65535 }"
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
    expression = "true"
    description = "Execute Cloudflare Managed Ruleset on my zone-level phase entry point ruleset"
    enabled = true
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
          action = "block"
          enabled = true
        }

        categories {
          category = "joomla"
          action = "block"
          enabled = true
        }
      }
    }

    expression = "true"
    description = "overrides to only enable wordpress rules to block"
    enabled = false
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

    expression = "(http.host eq \"example.com\" and http.uri.path eq \"/old-path\")"
    description = "example URI path transform rule"
    enabled = true
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

    expression = "true"
    description = "URI transformation query example"
    enabled = true
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

      expression = "true"
      description = "example request header transform rule"
      enabled = false
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional) The ID of the account where the ruleset is being created. Conflicts with `"zone_id"`.
* `description` - (Required) Brief summary of the ruleset and its intended use.
* `kind` - (Required) Type of Ruleset to create. Valid values are `"custom"`, `"managed"`, `"root"`, `"schema"` or `"zone"`.
* `name` - (Required) Name of the ruleset.
* `phase` - (Required) Point in the request/response lifecycle where the ruleset will be created. Valid values are `"ddos_l4"`, `"ddos_l7"`, `"http_request_firewall_custom"`, `"http_request_firewall_managed"`, `"http_request_late_transform"`, `"http_request_main"`, `"http_request_sanitize"`, `"http_request_transform"`, `"http_response_firewall_managed"`, `"magic_transit"`, or `"http_ratelimit"`.
* `rules` - (Required) List of rules to apply to the ruleset (refer to the [nested schema](#nestedblock--rules)).
* `shareable_entitlement_name` - (Optional) Name of entitlement that is shareable between entities.
* `zone_id` - (Optional) The ID of the zone where the ruleset is being created. Conflicts with `"account_id"`.

<a id="nestedblock--rules"></a>
**Nested schema for `rules`**

* `action_parameters` - (Required) List of parameters that configure the behavior of the ruleset rule action (refer to the [nested schema](#nestedblock--action-parameters)).
* `action` - (Required) Action to perform in the ruleset rule. Valid values are `"block"`, `"challenge"`, `"ddos_dynamic"`, `"execute"`, `"force_connection_close"`, `"js_challenge"`, `"log"`, `"rewrite"`, `"score"`, or  `"skip"`.
* `description` - (Optional) Brief summary of the ruleset rule and its intended use.
* `enabled` - (Optional) Whether the rule is active.
* `expression` - (Required) Criteria for an HTTP request to trigger the ruleset rule action. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.
* `id` - (Read only) Unique rule identifier.
* `ref` - (Read only) Rule reference.
* `version`- (Read only) Version of the ruleset to deploy.

<a id="nestedblock--action-parameters"></a>
**Nested schema for `action_parameters`**

* `id` - (Optional) Identifier of the action parameter to modify.
* `increment` - (Optional)
* `overrides` - (Optional) List of override configurations to apply to the ruleset (refer to the [nested schema](#nestedblock--action-parameters-overrides)).
* `products` - (Optional) Products to target with the actions. Valid values are `"bic"`, `"hot"`, `"ratelimit"`, `"securityLevel"`, `"uablock"`, `"waf"` or `"zonelockdown"`.
* `ruleset` - (Optional) Which ruleset to target. Valid value is `"current"`.
* `uri` - (Optional) List of URI properties to configure for the ruleset rule when performing URL rewrite transformations (refer to the [nested schema](#nestedblock--action-parameters-uri)).
* `headers` - (Optional) List of HTTP header modifications to perform in the ruleset rule (refer to the [nested schema](#nestedblock--action-parameters-headers)).
* `version` - (Optional)

<a id="nestedblock--action-parameters-uri"></a>
**Nested schema for `uri`**

* `path` - (Optional) URI path configuration when performing a URL rewrite (refer to the [nested schema](#nestedblock--action-parameters-uri-shared)).
* `query` - (Optional) Query string configuration when performing a URL rewrite (refer to the [nested schema](#nestedblock--action-parameters-uri-shared)).

<a id="nestedblock--action-parameters-headers"></a>
**Nested schema for `headers`**

* `name` - (Optional) Name of the HTTP request header to target.
* `operation` - (Optional) Action to perform on the HTTP request header. Valid values are `"set"` or `"remove"`.
* `expression` - (Optional) Use a value dynamically determined by the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions. Conflicts with `value`.
* `value` - (Optional) Static value to provide as the HTTP request header value. Conflicts with `expression`.

<a id="nestedblock--action-parameters-uri-shared"></a>
**Nested schema for `path`/`query`**

* `expression` - (Optional) Expression that defines the updated (dynamic) value of the URI path or query string component. Conflicts with `value`.
* `value` - (Optional) Static string value of the updated URI path or query string component. Conflicts with `expression`.

<a id="nestedblock--action-parameters-overrides"></a>
**Nested schema for `overrides`**

* `categories` - (Optional) List of tag-based overrides (refer to the [nested schema](#nestedblock--action-parameters-overrides-categories)).
* `enabled` - (Optional) Defines if the current ruleset-level override enables or disables the ruleset.
* `rules` - (Optional) List of rule-based overrides (refer to the [nested schema](#nestedblock--action-parameters-overrides-rules)).

<a id="nestedblock--action-parameters-overrides-categories"></a>
**Nested schema for `categories`**

* `category` - (Optional) Tag name to apply the ruleset rule override to.
* `action` - (Optional) Action to perform in the tag-level override. Valid values are `"block"`, `"challenge"`, `"ddos_dynamic"`, `"execute"`, `"force_connection_close"`, `"js_challenge"`, `"log"`, `"rewrite"`, `"score"`, or  `"skip"`.
* `enabled` - (Optional) Defines if the current tag-level override enables or disables the ruleset rules with the specified tag.

<a id="nestedblock--action-parameters-overrides-rules"></a>
**Nested schema for `rules`**

* `id` - (Optional) Rule ID to apply the override to.
* `action` - (Optional) Action to perform in the rule-level override. Valid values are `"block"`, `"challenge"`, `"ddos_dynamic"`, `"execute"`, `"force_connection_close"`, `"js_challenge"`, `"log"`, `"rewrite"`, `"score"`, or  `"skip"`.
* `enabled` - (Optional) Defines if the current rule-level override enables or disables the rule.
* `score_threshold` - (Optional) Anomaly score threshold to apply in the ruleset rule override. Only applicable to modsecurity-based rulesets.

## Import

Currently, you cannot import rulesets.
