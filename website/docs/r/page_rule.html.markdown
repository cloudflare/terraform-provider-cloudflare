---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_page rule"
sidebar_current: "docs-cloudflare-resource-page_rule"
description: |-
  Provides a Cloudflare page rule resource.
---

# cloudflare_page rule

Provides a Cloudflare page rule resource.

## Example Usage

```hcl
# Add a page rule to the domain
resource "cloudflare_page_rule" "foobar" {
  domain = "${var.cloudflare_domain}"
  target = "sub.${self.domain}/page"
  priority = 1

  actions = {
    ssl = "flexible",
    email_obfuscation = "on",
  }
}
```

## Argument Reference

The following arguments are supported:

* `zone` - (Required) The zone to which the page rule should be added.
* `target` - (Required) The URL pattern to target with the page rule.
* `actions` - (Required) The actions taken by the page rule, options given below.
* `priority` - (Optional) The priority of the page rule among others for this target.
* `status` - (Optional) Whether the page rule is active or disabled.

Action blocks support the following:

* `always_online` - (Optional) Whether this action is `"on"` or `"off"`.
* `automatic_https_rewrites` - (Optional) Whether this action is `"on"` or `"off"`.
* `browser_check` - (Optional) Whether this action is `"on"` or `"off"`.
* `email_obfuscation` - (Optional) Whether this action is `"on"` or `"off"`.
* `explicit_cache_control` - (Optional) Whether this action is `"on"` or `"off"`.
* `ip_geolocation` - (Optional) Whether this action is `"on"` or `"off"`.
* `opportunistic_encryption` - (Optional) Whether this action is `"on"` or `"off"`.
* `server_side_exclude` - (Optional) Whether this action is `"on"` or `"off"`.
* `smart_errors` - (Optional) Whether this action is `"on"` or `"off"`.
* `always_use_https` - (Optional) Boolean of whether this action is enabled. Default: false.
* `disable_apps` - (Optional) Boolean of whether this action is enabled. Default: false.
* `disable_performance` - (Optional) Boolean of whether this action is enabled. Default: false.
* `disable_railgun` - (Optional) Boolean of whether this action is enabled. Default: false.
* `disable_security` - (Optional) Boolean of whether this action is enabled. Default: false.
* `browser_cache_ttl` - (Optional) The Time To Live for the browser cache.
* `bypass_cache_on_cookie` - (Optional) Disable the cache on cookie regex match.
* `edge_cache_ttl` - (Optional) The Time To Live for the edge cache.
* `cache_level` - (Optional) Whether to set the cache level to `"byypass"`, `"basic"`, `"simplified"`, `"aggressive"`, or `"cache_everything"`.
* `forwarding_url` - (Optional) The URL to forward to, and with what status. See below.
* `host_header_override` - (Optional) The Host Header to override on the origin servers.
* `mirage` - (Optional) Whether this action is `"on"` or `"off"`.
* `origin_error_page_pass_thru` - (Optional) Whether this action is `"on"` or `"off"`.
* `resolve_override` - (Optional) Override the origin server with this host.
* `rocket_loader` - (Optional) Whether to set the rocket loader to `"off"`, `"manual"`, or `"automatic"`.
* `security_level` - (Optional) Whether to set the security level to `"off"`, `"essentially_off"`, `"low"`, `"medium"`, `"high"`, or `"under_attack"`.
* `sort_query_string_for_cache` - (Optional) Whether this action is `"on"` or `"off"`.
* `ssl` - (Optional) Whether to set the SSL mode to `"off"`, `"flexible"`, `"full"`, or `"strict"`.
* `true_client_ip_header` - (Optional) Whether this action is `"on"` or `"off"`.
* `waf` - (Optional) Whether this action is `"on"` or `"off"`.

Forwarding URL actions support the following:

* `url` - (Required) The URL to which the page rule should forward.
* `status_code` - (Required) The status code to use for the redirection.

## Attributes Reference

The following attributes are exported:

* `id` - The page rule ID.
* `zone_id` - The ID of the zone in which the page rule will be applied.
* `target` - The URL pattern targeted by the page rule.
* `actions` - The actions applied by the page rule.
* `priority` - The priority of the page rule.
* `status` - Whether the page rule is active or disabled.
