---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_ruleset"
sidebar_current: "docs-cloudflare-resource-ruleset"
description: |-
  Provides a resource which manages Cloudflare Rulesets.
---

# cloudflare_ruleset

The Cloudflare Ruleset Engine allows you to create and deploy rules and rulesets.
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

# Zone level managed WAF
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
    description = "Execute Cloudflare Managed Ruleset on my zone-level phase ruleset"
    enabled = true
  }
}

# Zone level WAF with category based overrides
resource "cloudflare_ruleset" "zone_level_managed_waf_with_category_based_overrides" {
  zone_id     = "cb029e245cfdd66dc8d2e570d5dd3322"
  name        = "managed WAF with category based overrides"
  description = "managed WAF with category based overrides ruleset description"
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
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional) The ID of the account where the Ruleset is being created. Conflicts with `"zone_id"`.
* `description` - (Required) Brief summary of the Ruleset and the intended use.
* `kind` - (Required) Type of Ruleset to create. Valid values are `"custom"`, `"managed"`, `"root"`, `"schema"` or `"zone"`.
* `name` - (Required) Name for the Ruleset.
* `phase` - (Required) Point in the request/response lifecycle where the Ruleset should be created. Valid values are `"ddos_l4"`, `"ddos_l7"`, `"http_request_firewall_custom"`, `"http_request_firewall_managed"`, `"http_request_late_transform"`, `"http_request_main"`, `"http_request_sanitize"`, `"http_request_transform"`, `"http_response_firewall_managed"`, `"magic_transit"` or `"http_ratelimit"`.
* `rules` - (Required) List of rules to apply to the Ruleset. (see [below for nested schema](#nestedblock--rules))
* `shareable_entitlement_name` - (Optional) Name of entitlement that is shareable between entities.
* `zone_id` - (Optional) The ID of the zone where the Ruleset is being created. Conflicts with `"account_id"`.

<a id="nestedblock--rules"></a>
**Nested schema for `rules`**

* `action_parameters` - (Required) List of parameters for actions to apply to the Ruleset Rule. (see [below for nested schema](#nestedblock--action-parameters))
* `action` - (Required) Action to take with Ruleset Rule. Valid values are `"block"`, `"challenge"`, `"ddos_dynamic"`, `"execute"`, `"force_connection_close"`, `"js_challenge"`, `"log"`, `"rewrite"`, `"score"` or  `"skip"`.
* `description` - (Optional) Brief summary of the Ruleset Rule and the intended use.
* `enabled` - (Optional) Whether the Ruleset is active.
* `expression` - (Required) Firewall Rules expression language based on Wireshark display filters. See [documentation](https://developers.cloudflare.com/firewall/cf-firewall-language) for all available fields, operators and functions.
* `id` - (Read only) Unique Rule identifier.
* `ref` - (Read only) Rule reference.
* `version`- (Read only) Version of the Ruleset that is to be deployed.

<a id="nestedblock--action-parameters"></a>
**Nested schema for `action_parameters`**

* `id` - (Optional) Action parameter identifier to target.
* `increment` - (Optional)
* `overrides` - (Optional) List of override configurations to apply to the Ruleset Rule. (see [below for nested schema](#nestedblock--action-parameters-overrides))
* `products` - (Optional) Products to target with the actions. Valid values are `"bic"`, `"hot"`, `"ratelimit"`, `"securityLevel"`, `"uablock"`, `"waf"` or `"zonelockdown"`.
* `ruleset` - (Optional) Which Ruleset to target. Valid value is `"current"`.
* `uri` - (Optional) List of URI properties to configure for the Ruleset Rule when performing URL rewrite transformation rule. (see [below for nested schema](#nestedblock--action-parameters-uri))
* `version` - (Optional)

<a id="nestedblock--action-parameters-uri"></a>
**Nested schema for `uri`**

* `path` - (Optional) Path configuration for URL rewriting. (see [below for nested schema](#nestedblock--action-parameters-uri-shared))
* `query` - (Optional) Query parameter configuration for URL rewriting. (see [below for nested schema](#nestedblock--action-parameters-uri-shared))

<a id="nestedblock--action-parameters-uri-shared"></a>
**Nested schema for `path`/`query`**

* `expression` - (Optional) Input and output pattern of the expression to apply to the URI path or query parameter component.
* `value` - (Optional) Static value to rewrite the value to.

<a id="nestedblock--action-parameters-overrides"></a>
**Nested schema for `overrides`**

* `categories` - (Optional) List of category based overrides. (see [below for nested schema](#nestedblock--action-parameters-overrides-categories))
* `enabled` - (Optional) Whether the Ruleset Rule override is active.
* `rules` - (Optional) List of rule based overrides. . (see [below for nested schema](#nestedblock--action-parameters-overrides-rules))

<a id="nestedblock--action-parameters-overrides-categories"></a>
**Nested schema for `categories`**

* `category` - (Optional) Category name to apply the Ruleset Rule override to.
* `action` - (Optional) Action to take with Ruleset Rule override. Valid values are `"block"`, `"challenge"`, `"ddos_dynamic"`, `"execute"`, `"force_connection_close"`, `"js_challenge"`, `"log"`, `"rewrite"`, `"score"` or  `"skip"`.
* `enabled` - (Optional) Whether the Ruleset Rule override is active.

<a id="nestedblock--action-parameters-overrides-rules"></a>
**Nested schema for `rules`**

* `id` - (Optional) Rule ID to apply the override to.
* `action` - (Optional) Action to take with Ruleset Rule override. Valid values are `"block"`, `"challenge"`, `"ddos_dynamic"`, `"execute"`, `"force_connection_close"`, `"js_challenge"`, `"log"`, `"rewrite"`, `"score"` or  `"skip"`.
* `enabled` - (Optional) Whether the Ruleset Rule override is active.
* `score_threshold` - (Optional) Anomaly score thresold to apply to Ruleset Rule override. Only applicable for modsecurity based Rulesets.

## Import

Importing Rulesets is not currently supported.
