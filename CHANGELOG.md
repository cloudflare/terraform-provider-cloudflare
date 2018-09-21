## 1.5.0 (Unreleased)

**Improvements:**
* **New Resource:** `cloudflare_zone_lockdown` [GH-115]

**Fixes:**
* Send User-Agent header with name and version when contacting API
* `cloudflare_page_rule` - Fix page rule polish (off, lossless or lossy) [GH-116]

## 1.4.0 (September 11, 2018)

**Improvements:**
* **New Resource:** `cloudflare_worker_route` ([#110](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/110))
* **New Resource:** `cloudflare_worker_script` ([#110](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/110))

## 1.3.0 (September 04, 2018)

**Improvements:**
* **New Resource:** `cloudflare_access_rule` ([#64](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/64))

**Fixes:**
* `cloudflare_zone_settings_override` -  Change Zone Settings Override to use GetOkExists ([#107](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/107))

## 1.2.0 (August 13, 2018)

**Improvements:**
* **New Resource:** `cloudflare_waf_rule` ([#98](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/98))
* `cloudflare_zone_settings_override` - Add `off` as Security Level setting ([#99](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/99))
* `resource_cloudflare_rate_limit` - Add nat support ([#96](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/96))
* `resource_cloudflare_zone_settings_override` - Add `zrt` as a value for the `tls_1_3` setting ([#106](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/106))
* Minor documentation improvements

**Fixes:**
* `cloudflare_record` - Setting a DNS record's `proxied` flag to false stopped working ([#103](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/103))

## 1.1.0 (July 25, 2018)

FIXES:

* `cloudflare_ip_ranges` - IPv6 CIDR blocks should return IPv6 addresses ([#51](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/51))
* `cloudflare_zone_settings_override` - Allow `0` for `browser_cache_ttl` ([#71](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/71))
* `cloudflare_page_rule` - `forwarding_urls` in page rules are lists ([#79](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/79))
* `cloudflare_page_rule` - The API supports `active` and `disabled`, not `paused` ([#84](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/84))

IMPROVEMENTS:
* `cloudflare_zone_settings_override` - Add support for `min_tls_version` ([#72](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/72))
* `cloudflare_page_rule` - Add support for more settings: `bypass_cache_on_cookie`, `cache_by_device_type`, `cache_deception_armor`, `cache_on_cookie`, `host_header_override`, `polish`, `explicit_cache_control`, `origin_error_page_pass_thru`, `sort_query_string_for_cache`, `resolve_override`, `respect_strong_etag`, `response_buffering`, `true_client_ip_header`, `mirage`, `disable_railgun`, `cache_key`, `waf`, `rocket_loader`, `cname_flattening` ([#68](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/68)], [[#81](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/81)], [[#85](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/85))
* `cloudflare_page_rule` - Add `off` setting to `security_level` ([#81](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/81))
* `cloudflare_record` - DNS Record improvements ([#97](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/97))
* Various documentation improvements

## 1.0.0 (April 06, 2018)

BACKWARDS INCOMPATIBILITIES / NOTES:

* resource/cloudflare_record: Changing `name` or `domain` now force a recreation
  of the record ([#29](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/29))

FEATURES:

* **New Resource:** `cloudflare_rate_limit` ([#30](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/30))
* **New Resource:** `cloudflare_page_rule` ([#38](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/38))
* **New Resource:** `cloudflare_load_balancer` ([#40](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/40))
* **New Resource:** `cloudflare_load_balancer_pool` ([#40](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/40))
* **New Resource:** `cloudflare_zone_settings_override` ([#41](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/41))
* **New Resource:** `cloudflare_load_balancer_monitor` ([#42](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/42))
* **New Data Source:** `cloudflare_ip_ranges` ([#28](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/28))

IMPROVEMENTS:

* resource/cloudflare_record: Validate `TXT` records ([#14](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/14))
* resource/cloudflare_record: Add `data` input to suppport SRV, LOC records
  ([#29](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/29))
* resource/cloudflare_record: Add computed attributes `created_on`,
  `modified_on`, `proxiable`, and `metadata` to records ([#29](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/29))
* resource/cloudflare_record: Support import of existing records ([#36](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/36))
* New Provider configuration options for API rate limiting ([#43](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/43))
* New Provider configuration options for using Organizations ([#40](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/40))

## 0.1.0 (June 20, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider
  Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
