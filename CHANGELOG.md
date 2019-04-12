## 1.13.0 (April 12, 2019)

**Improvements**

* **New Resource:** `cloudflare_logpush_job` ([#287](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/287))
* `cloudflare_zone_settings` - Remove option to toggle `always_on_ddos` ([#253](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/253))
* `cloudflare_page_rule` - Update documentation to clarify "0" usage
* `cloudflare_zones` - Return zone ID and zone name ([#275](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/275))
* `cloudflare_load_balancer` - Add `enabled` field ([#208](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/208))
* `cloudflare_record` - validators: Allow PTR DNS records ([#283](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/283))

**Fixes:**

* `cloudflare_custom_pages` - Use correct casing for `zone_id` lookups
* `cloudflare_rate_limit` - Make `correlate` optional and not flap in state management ([#271](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/271))
* `cloudflare_spectrum_application` - Fixed integration tests to work ([#275](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/275))
* `cloudflare_page_rule` - Better track field changes in `actions` resource. ([#107](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/107))

## 1.12.0 (March 07, 2019)

**Improvements:**

* provider: Enable request/response logging ([#212](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/212))
* resource/cloudflare_load_balancer_monitor: Add validation for `port` ([#213](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/213))
* resource/cloudflare_load_balancer_monitor: Add `allow_insecure` and `follow_redirects` ([#205](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/205))
* resource/cloudflare_page_rule: Updated available actions documentation to match what is available ([#228](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/228))
* provider: Swap to using go modules for dependency management ([#230](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/230))
* provider: Minimum Go version for development is now 1.11 ([#230](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/230))

**Fixes:**

* resource/cloudflare_record: Read `data` back from API correctly ([#217](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/217))
* resource/cloudflare_rate_limit: Read `correlate` back from API correctly ([#204](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/204))
* resource/cloudflare_load_balancer_monitor: Fix incorrect type cast for `port` ([#213](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/213))
* resource/cloudflare_load_balancer: Make `steering_policy` computed to avoid spurious diffs ([#214](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/214))
* resource/cloudflare_load_balancer: Read `session_affinity` back from API to make import work & detects drifts ([#214](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/214))

## 1.11.0 (January 11, 2019)

**Improvements:**
* **New Resource:** `cloudflare_spectrum_app` ([#156](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/156))
* **New Data Source:** `cloudflare_zones` ([#168](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/168))
* `cloudflare_load_balancer_monitor` - Add optional `port` parameter ([#179](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/179))
* `cloudflare_page_rule` - Improved documentation for `priority` attribute ([#182](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/182)], missing `explicit_cache_control` [[#185](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/185))
* `cloudflare_rate_limit` - Add `challenge` and `js_challenge` rate-limit modes ([#172](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/172))

**Fixes:**
* `cloudflare_page_rule` - Page rule `zone` attribute change to trigger new resource ([#183](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/183))

## 1.10.0 (December 18, 2018)

**Improvements:**
* `cloudflare_zone_settings_override` - Add `opportunistic_onion` zone setting support ([#170](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/170))
* `cloudflare_zone` - Add ability to set zone plan ([#160](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/160))

**Fixes:**
* `cloudflare_zone` - Allow zones to be properly imported ([#157](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/157))
* `cloudflare_access_policy` - Match access_policy argument requisites with reality ([#158](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/158))
* `cloudflare_filter` - Allow `zone_id` to set `zone` and vice versa ([#162](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/162))
* `cloudflare_firewall_rule` - Allow `zone_id` to set `zone` and vice versa ([#174](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/174))
* `cloudflare_access_rule` - Ensure `zone` and `zone_id` are always set ([#175](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/175))
* Minor documentation fixes

## 1.9.0 (November 15, 2018)

**Improvements:**
* **New Resource:** `cloudflare_access_application` ([#145](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/145))
* **New Resource:** `cloudflare_access_policy` ([#145](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/145))
* `cloudflare_load_balancer` - Add steering policy support ([#147](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/147))
* `cloudflare_load_balancer` - Support `session_affinity` ([#153](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/153))
* `cloudflare_load_balancer_pool` - Support `weight` ([#153](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/153))

**Fixes:**
* `cloudflare_record` - Compare name without the zone name ([#151](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/151))
* Minor documentation fixes ([#149](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/149)] [[#152](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/152))

## 1.8.0 (November 05, 2018)

**Improvements:**
* **New Resource:** `cloudflare_zone` ([#58](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/58))
* **New Resource:** `cloudflare_custom_pages` ([#132](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/132))
* `cloudflare_zone_settings_override` - Allow setting SSL level to Strict (SSL-Only Origin Pull) ([#122](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/122))
* Update provider usage/build docs and how to update a dependency ([#138](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/138))
* Improve `Building The Provider` instructions ([#143](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/143))
* `cloudflare_access_rule` - Make importable for all rule types ([#141](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/141))
* `cloudflare_load_balancer_pool` - Implement `Update` ([#140](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/140))

**Fixes:**
* `cloudflare_rate_limit` - Documentation fixes for markdown where \_ALL\_ is italicized ([#125](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/125))
* `cloudflare_worker_route` - Correctly set `multi_script` on Enterprise worker imports ([#124](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/124))
* `account_member` - Ignore role ID ordering ([#128](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/128))
* `cloudflare_rate_limit` - Origin traffic isn't default anymore ([#130](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/130))
* `cloudflare_rate_limit` - Update rate limit validation to allow `1` ([#129](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/129))
* `cloudflare_record` - Add validation to ensure TTL is not set while `proxied` is true ([#127](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/127))
* Updated code for provider version in User-Agent
* `cloudflare_zone_lockdown` - Fix import of zone lockdowns ([#135](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/135))

## 1.7.0 (October 09, 2018)

**Improvements:**
* **New Resource:** `cloudflare_account_member` ([#78](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/78))

## 1.6.0 (October 05, 2018)

**Improvements:**
* **New Resource:** `cloudflare_filter`
* **New Resource:** `cloudflare_firewall_rule`

## 1.5.0 (September 21, 2018)

**Improvements:**
* **New Resource:** `cloudflare_zone_lockdown` ([#115](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/115))

**Fixes:**
* Send User-Agent header with name and version when contacting API
* `cloudflare_page_rule` - Fix page rule polish (off, lossless or lossy) ([#116](https://github.com/terraform-providers/terraform-provider-cloudflare/issues/116))

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
