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
	name = "${var.domain}"
	settings {
		brotli = "on",
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

* `name` - (Required) The domain name of the zone.
* `settings` - (Optional) Settings overrides that will be applied to the zone. If a setting is not specified the existing setting will be used. For a full list of available settings see below.

The **settings** block supports settings that may be applied to the zone. These may be on/off values, unitary fields, string values, integers or nested objects.

### On/Off Values

These can be specified as "on" or "off" string. Similar to boolean values, but here the empty string also means to use the existing value. Attributes available:

* `advanced_ddos`
* `always_online`
* `brotli`
* `browser_check`
* `cache_level`
* `development_mode`
* `origin_error_page_pass_thru`
* `sort_query_string_for_cache`
* `email_obfuscation`
* `hotlink_protection`
* `ip_geolocation`
* `ipv6`
* `websockets`
* `mirage`
* `opportunistic_encryption`
* `prefetch_preload`
* `privacy_pass`
* `response_buffering`
* `server_side_exclude`
* `tls_client_auth`
* `true_client_ip_header`
* `waf`
* `tls_1_2_only`
* `tls_1_3`
* `automatic_https_rewrites`
* `http2`
* `sha1_support`
* `always_use_https`. In some cases setting this might give the error `HTTP status 400: content "{\"success\":false,\"errors\":[{\"code\":1016,\"message\":\"An unknown error has occurred\"}],\"messages\":[],\"result\":null}"`. Regardless, the value is set correctly.
* `webp`. Note that the value specified will be ignored unless `polish` is turned on (i.e. is "lossless" or "lossy")

### String Values

* `cache_level`. Allowed values: "aggressive", "basic", "simplified".
* `polish`. Allowed values: "off", "lossless", "lossy".
* `rocket_loader`. Allowed values: "on", "off", "manual".
* `security_level`. Allowed values: "essentially_off", "low", "medium", "high", "under_attack".
* `ssl`. Allowed values: "off", "flexible", "full", "strict".
* `pseudo_ipv4`. Allowed values: "off", "add_header", "overwrite_header".
* `cname_flattening`.
* `min_tls_version`. Allowed values: "1.0", "1.1", "1.2", "1.3"

### Integer Values

* `browser_cache_ttl`
* `challenge_ttl`
* `max_upload`
* `edge_cache_ttl`

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
* `strip_uri` (Required) true/false
* `status` (Required) "on"/"off"

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