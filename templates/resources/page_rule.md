---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_page_rule"
description: Provides a Cloudflare page rule resource.
---

# cloudflare_page_rule

Provides a Cloudflare page rule resource.

## Example Usage

```hcl
# Add a page rule to the domain
resource "cloudflare_page_rule" "foobar" {
  zone_id = var.cloudflare_zone_id
  target = "sub.${var.cloudflare_zone}/page"
  priority = 1

  actions {
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

- `zone_id` - (Required) The DNS zone ID to which the page rule should be added.
- `target` - (Required) The URL pattern to target with the page rule.
- `actions` - (Required) The actions taken by the page rule, options given below.
- `priority` - (Optional) The priority of the page rule among others for this target, the higher the number the higher the priority as per [API documentation](https://api.cloudflare.com/#page-rules-for-a-zone-create-page-rule).
- `status` - (Optional) Whether the page rule is active or disabled.

Action blocks support the following:

- `always_use_https` - (Optional) Boolean of whether this action is enabled. Default: false.
- `automatic_https_rewrites` - (Optional) Whether this action is `"on"` or `"off"`.
- `browser_cache_ttl` - (Optional) The Time To Live for the browser cache. `0` means 'Respect Existing Headers'
- `browser_check` - (Optional) Whether this action is `"on"` or `"off"`.
- `bypass_cache_on_cookie` - (Optional) String value of cookie name to conditionally bypass cache the page.
- `cache_by_device_type` - (Optional) Whether this action is `"on"` or `"off"`.
- `cache_deception_armor` - (Optional) Whether this action is `"on"` or `"off"`.
- `cache_key_fields` - (Optional) Controls how Cloudflare creates Cache Keys used to identify files in cache. [See below](#cache-key-fields) for full description.
- `cache_level` - (Optional) Whether to set the cache level to `"bypass"`, `"basic"`, `"simplified"`, `"aggressive"`, or `"cache_everything"`.
- `cache_on_cookie` - (Optional) String value of cookie name to conditionally cache the page.
- `cache_ttl_by_status` - (Optional) Set cache TTL based on the response status from the origin web server. Can be specified multiple times. [See below](#cache-ttl-by-status) for full description.
- `disable_apps` - (Optional) Boolean of whether this action is enabled. Default: false.
- `disable_performance` - (Optional) Boolean of whether this action is enabled. Default: false.
- `disable_railgun` - (Optional) Boolean of whether this action is enabled. Default: false.
- `disable_security` - (Optional) Boolean of whether this action is enabled. Default: false.
- `disable_zaraz` - (Optional) Boolean of whether this action is enabled. Default: false.
- `edge_cache_ttl` - (Optional) The Time To Live for the edge cache.
- `email_obfuscation` - (Optional) Whether this action is `"on"` or `"off"`.
- `explicit_cache_control` - (Optional) Whether origin Cache-Control action is `"on"` or `"off"`.
- `forwarding_url` - (Optional) The URL to forward to, and with what status. See below.
- `host_header_override` - (Optional) Value of the Host header to send.
- `ip_geolocation` - (Optional) Whether this action is `"on"` or `"off"`.
- `minify` - (Optional) The configuration for HTML, CSS and JS minification. See below for full list of options.
- `mirage` - (Optional) Whether this action is `"on"` or `"off"`.
- `opportunistic_encryption` - (Optional) Whether this action is `"on"` or `"off"`.
- `origin_error_page_pass_thru` - (Optional) Whether this action is `"on"` or `"off"`.
- `polish` - (Optional) Whether this action is `"off"`, `"lossless"` or `"lossy"`.
- `resolve_override` - (Optional) Overridden origin server name.
- `respect_strong_etag` - (Optional) Whether this action is `"on"` or `"off"`.
- `response_buffering` - (Optional) Whether this action is `"on"` or `"off"`.
- `rocket_loader` - (Optional) Whether to set the rocket loader to `"on"`, `"off"`.
- `security_level` - (Optional) Whether to set the security level to `"off"`, `"essentially_off"`, `"low"`, `"medium"`, `"high"`, or `"under_attack"`.
- `server_side_exclude` - (Optional) Whether this action is `"on"` or `"off"`.
- `smart_errors` - (Optional) Whether this action is `"on"` or `"off"`.
- `sort_query_string_for_cache` - (Optional) Whether this action is `"on"` or `"off"`.
- `ssl` - (Optional) Whether to set the SSL mode to `"off"`, `"flexible"`, `"full"`, `"strict"`, or `"origin_pull"`.
- `true_client_ip_header` - (Optional) Whether this action is `"on"` or `"off"`.
- `waf` - (Optional) Whether this action is `"on"` or `"off"`.

Forwarding URL actions support the following:

- `url` - (Required) The URL to which the page rule should forward.
- `status_code` - (Required) The status code to use for the redirection.

Minify actions support the following:

- `html` - (Required) Whether HTML should be minified. Valid values are `"on"` or `"off"`.
- `css` - (Required) Whether CSS should be minified. Valid values are `"on"` or `"off"`.
- `js` - (Required) Whether Javascript should be minified. Valid values are `"on"` or `"off"`.

### Cache Key Fields

-> This setting is available to Enterprise customers only.

A Cache Key is an identifier that Cloudflare uses for a file in a cache. The Cache Key Template defines this identifier for a given HTTP request.

For detailed description of use cases and semantics for the particular setting please refer to [Cloudflare Support article](https://support.cloudflare.com/hc/en-us/articles/115004290387-Creating-Cache-Keys).

Example:

```hcl
# Cache JavaScript files:
# - ignore CORS Origin header (one copy regardless of requesting Host)
# - ignore API key query string
# - include browser language preference (e.g. string translations)
resource "cloudflare_page_rule" "foobar" {
  zone_id = var.cloudflare_zone_id
  target = "embed.${var.cloudflare_zone}/*.js"
  priority = 1

  actions {
    cache_key_fields {
      header {
        exclude = ["origin"]
      }
      query_string {
        exclude = ["api_token"]
      }
      user {
        lang = true
      }
      cookie {}
      host {}
    }
  }
}
```

- `cookie` - (Optional) Controls what cookies go into Cache Key:
  - `check_presence` - (Optional, Array) Check for presence of specified cookies, without including their actual values.
  - `include` - (Optional, Array) Use values of specified cookies in Cache Key.
- `header` - (Optional) Controls what HTTP headers go into Cache Key:
  - `check_presence` - (Optional, Array) Check for presence of specified HTTP headers, without including their actual values.
  - `exclude` - (Optional, Array) Exclude these HTTP headers from Cache Key. Currently, only the `Origin` header can be excluded.
  - `include` - (Optional, Array) Use values of specified HTTP headers in Cache Key. Please refer to [Support article](https://support.cloudflare.com/hc/en-us/articles/115004290387-Creating-Cache-Keys) for the list of HTTP headers that cannot be included. The `Origin` header is always included unless explicitly excluded.
- `host` - (Required, but allowed to be empty) Controls which Host header goes into Cache Key:
  - `resolved` - (Optional, Boolean) `false` (default) - includes the Host header in the HTTP request sent to the origin; `true` - includes the Host header that was resolved to get the origin IP for the request (e.g. changed with Resolve Override Page Rule).
- `query_string` - (Required, but allowed to be empty) Controls which URL query string parameters go into the Cache Key.
  - `exclude` - (Optional, Array) Exclude these query string parameters from Cache Key.
  - `include` - (Optional, Array) Only use values of specified query string parameters in Cache Key.
  - `ignore` - (Optional, Boolean) `false` (default) - all query string parameters are used for Cache Key, unless explicitly excluded; `true` - all query string parameters are ignored; value should be `false` if any of `exclude` or `include` is non-empty.
- `user` - (Required, but allowed to be empty) Controls which end user-related features go into the Cache Key.
  - `device_type` - (Optional, Boolean) `true` - classifies a request as “mobile”, “desktop”, or “tablet” based on the User Agent; defaults to `false`.
  - `geo` - (Optional, Boolean) `true` - includes the client’s country, derived from the IP address; defaults to `false`.
  - `lang` - (Optional, Boolean) `true` - includes the first language code contained in the `Accept-Language` header sent by the client; defaults to `false`.

Example:

```hcl
# Unrealistic example with all features used
resource "cloudflare_page_rule" "foobar" {
  zone_id = var.cloudflare_zone_id
  target = "${var.cloudflare_zone}/app/*"
  priority = 1

  actions {
    cache_key_fields {
      cookie {
        check_presence = ["wordpress_test_cookie"]
      }
      header {
        check_presence = ["header_present"]
        exclude = ["origin"]
        include = ["api-key", "dnt"]
      }
      host {
        resolved = true
      }
      query_string {
        ignore = true
      }
      user {
        device_type = false
        geo = true
        lang = true
      }
    }
  }
}
```

### Cache TTL by status

-> This setting is available to Enterprise customers only.

Set cache TTL based on the response status from the origin web server. Cache TTL (time-to-live) refers to the duration a resource lives in the Cloudflare network before it is marked as stale or discarded from cache. Status codes are returned by a resource’s origin.

For detailed description please refer to [Cloudflare Support article](https://support.cloudflare.com/hc/en-us/articles/360043842472-Configuring-cache-TTL-by-status-code).

- `codes` - (Required) A HTTP code (e.g. `404`) or range of codes (e.g. `400-499`)
- `ttl` - (Required) Duration a resource lives in the Cloudflare cache.
  - positive number - cache for specified duration in seconds
  - `0` - sets `no-cache`, saved to cache, but expired immediately (revalidate from origin every time)
  - `-1` - sets `no-store`, never save to cache

Example:

```hcl
resource "cloudflare_page_rule" "test" {
	zone_id = var.cloudflare_zone_id
	target = "${var.cloudflare_zone}/app/*"
	priority = 1

	actions {
		cache_ttl_by_status {
			codes = "200-299"
			ttl = 300
		}
		cache_ttl_by_status {
			codes = "300-399"
			ttl = 60
		}
		cache_ttl_by_status {
			codes = "400-403"
			ttl = -1
		}
		cache_ttl_by_status {
			codes = "404"
			ttl = 30
		}
		cache_ttl_by_status {
			codes = "405-499"
			ttl = -1
		}
		cache_ttl_by_status {
			codes = "500-599"
			ttl = 0
		}
	}
}
```

## Attributes Reference

The following attributes are exported:

- `id` - The page rule ID.
- `target` - The URL pattern targeted by the page rule.
- `actions` - The actions applied by the page rule.
- `priority` - The priority of the page rule.
- `status` - Whether the page rule is active or disabled.

## Import

Page rules can be imported using a composite ID formed of zone ID and page rule ID, e.g.

```
$ terraform import cloudflare_page_rule.default d41d8cd98f00b204e9800998ecf8427e/ch8374ftwdghsif43
```
