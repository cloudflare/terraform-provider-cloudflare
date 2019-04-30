---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_page_rule"
sidebar_current: "docs-cloudflare-resource-page-rule"
description: |-
  Provides a Cloudflare page rule resource.
---

# cloudflare_page_rule

Provides a Cloudflare page rule resource.

## Example Usage

```hcl
# Add a page rule to the domain
resource "cloudflare_page_rule" "foobar" {
  zone = "${var.cloudflare_zone}"
  target = "sub.${var.cloudflare_zone}/page"
  priority = 1

  actions = {
    ssl = "flexible"
    email_obfuscation = "on"
    minify {
      html = "off"
      css  = "on"
      js   = "on"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `zone` - (Required) The DNS zone to which the page rule should be added.
* `target` - (Required) The URL pattern to target with the page rule.
* `actions` - (Required) The actions taken by the page rule, options given below.
* `priority` - (Optional) The priority of the page rule among others for this target, the higher the number the higher the priority as per [API documentation](https://api.cloudflare.com/#page-rules-for-a-zone-create-page-rule).
* `status` - (Optional) Whether the page rule is active or disabled.

Action blocks support the following:

* `always_online` - (Optional) Whether this action is `"on"` or `"off"`.
* `always_use_https` - (Optional) Boolean of whether this action is enabled. Default: false.
* `automatic_https_rewrites` - (Optional) Whether this action is `"on"` or `"off"`.
* `browser_cache_ttl` - (Optional) The Time To Live for the browser cache. `0` means 'Respect Existing Headers'
* `browser_check` - (Optional) Whether this action is `"on"` or `"off"`.
* `bypass_cache_on_cookie` - (Optional) String value of cookie name to conditionally bypass cache the page.
* `cache_by_device_type` - (Optional) Whether this action is `"on"` or `"off"`.
* `cache_deception_armor` - (Optional) Whether this action is `"on"` or `"off"`.
* `cache_level` - (Optional) Whether to set the cache level to `"bypass"`, `"basic"`, `"simplified"`, `"aggressive"`, or `"cache_everything"`.
* `cache_on_cookie` - (Optional) String value of cookie name to conditionally cache the page.
* `disable_apps` - (Optional) Boolean of whether this action is enabled. Default: false.
* `disable_performance` - (Optional) Boolean of whether this action is enabled. Default: false.
* `disable_railgun` - (Optional) Boolean of whether this action is enabled. Default: false.
* `disable_security` - (Optional) Boolean of whether this action is enabled. Default: false.
* `edge_cache_ttl` - (Optional) The Time To Live for the edge cache.
* `email_obfuscation` - (Optional) Whether this action is `"on"` or `"off"`.
* `explicit_cache_control` - (Optional) Whether origin Cache-Control action is `"on"` or `"off"`.
* `forwarding_url` - (Optional) The URL to forward to, and with what status. See below.
* `host_header_override` - (Optional) Value of the Host header to send.
* `ip_geolocation` - (Optional) Whether this action is `"on"` or `"off"`.
* `minify` - (Optional) The configuration for HTML, CSS and JS minification. See below for full list of options.
* `mirage` - (Optional) Whether this action is `"on"` or `"off"`.
* `opportunistic_encryption` - (Optional) Whether this action is `"on"` or `"off"`.
* `origin_error_page_pass_thru` - (Optional) Whether this action is `"on"` or `"off"`.
* `polish` - (Optional) Whether this action is `"off"`, `"lossless"` or `"lossy"`.
* `resolve_override` - (Optional) Overridden origin server name.
* `respect_strong_etag` - (Optional) Whether this action is `"on"` or `"off"`.
* `response_buffering` - (Optional) Whether this action is `"on"` or `"off"`.
* `rocket_loader` - (Optional) Whether to set the rocket loader to `"on"`, `"off"`.
* `security_level` - (Optional) Whether to set the security level to `"off"`, `"essentially_off"`, `"low"`, `"medium"`, `"high"`, or `"under_attack"`.
* `server_side_exclude` - (Optional) Whether this action is `"on"` or `"off"`.
* `smart_errors` - (Optional) Whether this action is `"on"` or `"off"`.
* `sort_query_string_for_cache` - (Optional) Whether this action is `"on"` or `"off"`.
* `ssl` - (Optional) Whether to set the SSL mode to `"off"`, `"flexible"`, `"full"`, or `"strict"`.
* `true_client_ip_header` - (Optional) Whether this action is `"on"` or `"off"`.
* `waf` - (Optional) Whether this action is `"on"` or `"off"`.

Forwarding URL actions support the following:

* `url` - (Required) The URL to which the page rule should forward.
* `status_code` - (Required) The status code to use for the redirection.

Minify actions support the following:

* `html` - (Required) Whether HTML should be minified. Valid values are `"on"` or `"off"`.
* `css` - (Required) Whether CSS should be minified. Valid values are `"on"` or `"off"`.
* `js` - (Required) Whether Javascript should be minified. Valid values are `"on"` or `"off"`.

## Attributes Reference

The following attributes are exported:

* `id` - The page rule ID.
* `zone_id` - The ID of the zone in which the page rule will be applied.
* `target` - The URL pattern targeted by the page rule.
* `actions` - The actions applied by the page rule.
* `priority` - The priority of the page rule.
* `status` - Whether the page rule is active or disabled.

## Import

Page rules can be imported using a composite ID formed of zone name and page rule ID, e.g.

```
$ terraform import cloudflare_page_rule.default example.com/ch8374ftwdghsif43
```
