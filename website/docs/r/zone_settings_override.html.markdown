---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_zone_settings_override"
sidebar_current: "docs-cloudflare-resource-zone-settings-override"
description: |-
  Provides a resource which customizes Cloudflare zone settings.
---

# cloudflare_zone_settings_override

Provides a resource which customizes Cloudflare zone settings. Note that after destroying this resource Zone Settings will be reset to their initial values.

## Example Usage

```hcl
resource "cloudflare_zone_settings_override" "test" {
	zone_id = var.cloudflare_zone_id
	settings {
		brotli = "on"
		challenge_ttl = 2700
		security_level = "high"
		opportunistic_encryption = "on"
		automatic_https_rewrites = "on"
		mirage = "on"
		waf = "on"
		minify {
			css = "on"
			js = "off"
			html = "off"
		}
		security_header {
			enabled = true
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone ID to which apply settings.
* `settings` - (Optional) Settings overrides that will be applied to the zone. If a setting is not specified the existing setting will be used. For a full list of available settings see below.

The **settings** block supports settings that may be applied to the zone. These may be on/off values, unitary fields, string values, integers or nested objects.

### On/Off Values

These can be specified as "on" or "off" string. Similar to boolean values, but here the empty string also means to use the existing value. Attributes available:

* `always_online` (default: `on`)
* `always_use_https` (default: `off`)
* `automatic_https_rewrites` (default value depends on the zone's plan level)
* `brotli` (default: `off`)
* `browser_check` (default: `on`)
* `development_mode` (default: `off`)
* `email_obfuscation` (default: `on`)
* `hotlink_protection` (default: `off`)
* `http2` (default: `off`)
* `http3` (default: `off`)
* `ip_geolocation` (default: `on`)
* `ipv6` (default: `off`)
* `mirage` (default: `off`)
* `opportunistic_encryption` (default value depends on the zone's plan level)
* `opportunistic_onion` (default: `off`)
* `origin_error_page_pass_thru` (default: `off`)
* `prefetch_preload` (default: `off`)
* `privacy_pass` (default: `on`)
* `response_buffering` (default: `off`)
* `rocket_loader` (default: `off`)
* `server_side_exclude` (default: `on`)
* `sort_query_string_for_cache` (default: `off`)
* `tls_client_auth` (default: `on`)
* `true_client_ip_header` (default: `off`)
* `waf` (default: `off`)
* `webp` (default: `off`). Note that the value specified will be ignored unless `polish` is turned on (i.e. is "lossless" or "lossy")
* `websockets` (default: `off`)
* `zero_rtt` (default: `off`)

### String Values

* `cache_level`. Allowed values: "aggressive" (default) - delivers a different resource each time the query string changes, "basic" - delivers resources from cache when there is no query string, "simplified" - delivers the same resource to everyone independent of the query string.
* `cname_flattening`. Allowed values: "flatten_at_root" (default), "flatten_all", "flatten_none".
* `h2_prioritization`. Allowed values: "on", "off" (default), "custom".
* `image_resizing`. Allowed values: "on", "off" (default), "open".
* `min_tls_version`. Allowed values: "1.0" (default), "1.1", "1.2", "1.3".
* `polish`. Allowed values: "off" (default), "lossless", "lossy".
* `pseudo_ipv4`. Allowed values: "off" (default), "add_header", "overwrite_header".
* `security_level`. Allowed values: "off" (Enterprise only), "essentially_off", "low", "medium" (default), "high", "under_attack".
* `ssl`. Allowed values: "off" (default), "flexible", "full", "strict", "origin_pull".
* `tls_1_3`. Allowed values: "off" (default), "on", "zrt".

### Integer Values

* `browser_cache_ttl` (default: `14400`)
* `challenge_ttl` (default: `1800`)
* `edge_cache_ttl` (default: `7200`)
* `max_upload` (default: `100`)

### Nested Objects

* `minify`
* `mobile_redirect`
* `security_header`

The **minify** attribute supports the following fields:

* `css` (Required) "on"/"off"
* `html` (Required) "on"/"off"
* `js` (Required)"on"/"off"

The **mobile_redirect** attribute supports the following fields:

* `mobile_subdomain` (Required) String value
* `status` (Required) "on"/"off"
* `strip_uri` (Required) true/false

The **security_header** attribute supports the following fields:

* `enabled` (Optional) true/false
* `preload` (Optional) true/false
* `max_age` (Optional) Integer
* `include_subdomains` (Optional) true/false
* `nosniff` (Optional) true/false

## Attributes Reference

The following attributes are exported:

* `id` - The zone ID.
* `initial_settings` - Settings present in the zone at the time the resource is created. This will be used to restore the original settings when this resource is destroyed. Shares the same schema as the `settings` attribute (Above).
* `intial_settings_read_at` - Time when this resource was created and the `initial_settings` were set.
* `readonly_settings` - Which of the current `settings` are not able to be set by the user. Which settings these are is determined by plan level and user permissions.
* `zone_status`. A full zone implies that DNS is hosted with Cloudflare. A partial zone is typically a partner-hosted zone or a CNAME setup.
* `zone_type`. Status of the zone. Valid values: active, pending, initializing, moved, deleted, deactivated.
