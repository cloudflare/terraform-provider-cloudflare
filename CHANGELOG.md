## 3.4.0 (Unreleased)

## 3.3.0 (October 20th, 2021)

FEATURES:

* **New Data Source:** `cloudflare_account_roles` ([#1238](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1238))

ENHANCEMENTS:

* resource/cloudflare_access_application: add support for 'SameSite' and 'HttpOnly' cookie attributes ([#1241](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1241))
* resource/cloudflare_argo_tunnel: add `cname` as exported attribute ([#1259](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1259))
* resource/cloudflare_load_balancer_pool: add support for origin steering ([#1240](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1240))
* resource/cloudflare_ruleset: add support for 'Action' and 'Enabled' action_parameters > overrides attributes ([#1249](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1249))
* resource/cloudflare_zone_setting_override: add support for overriding `binary_ast` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
* resource/cloudflare_zone_setting_override: add support for overriding `filter_logs_to_cloudflare` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
* resource/cloudflare_zone_setting_override: add support for overriding `log_to_cloudflare` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
* resource/cloudflare_zone_setting_override: add support for overriding `orange_to_orange` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
* resource/cloudflare_zone_setting_override: add support for overriding `proxy_read_timeout` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
* resource/cloudflare_zone_setting_override: add support for overriding `visitor_ip` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))

BUG FIXES:

* resource/cloudflare_access_policy: handle empty `nil` values for building policies ([#1237](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1237))
* resource/cloudflare_ruleset: don't attempt to update "custom" rulesets using the phase entrypoint ([#1245](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1245))

## 3.2.0 (October 7th, 2021)

NOTES:

* provider: cloudflare-go has been upgraded to v0.25.0 ([#1236](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1236))

FEATURES:

* **New Data Source:** `cloudflare_zone` ([#1213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1213))
* **New Resource:** `cloudflare_split_tunnel` ([#1207](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1207))

ENHANCEMENTS:

* provider: add support for debugging via debuggers (like delve) ([#1217](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1217))
* resource/cloudflare_access_policy: add support for approval_required flag ([#1230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1230))

BUG FIXES:

* resource/cloudflare_account_member: handle role changes made in the dashboard ([#1202](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1202))
* resource/cloudflare_origin_ca_certificate: ignore `requested_validity` changes due to the value decreasing but still store it ([#1214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1214))
* resource/cloudflare_record: handle `Update`s for records with `data` blocks ([#1229](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1229))

## 3.1.0 (September 21st, 2021)

ENHANCEMENTS:

* resource/cloudflare_ruleset: add support for ddos_l7 configuration ([#1212](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1212))

## 3.0.1 (September 21st, 2021)

ENHANCEMENTS:

* resource/cloudflare_access_rule: add state migrator for 3.x ([#1211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1211))
* resource/cloudflare_custom_ssl: add state migrator for 3.x ([#1211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1211))
* resource/cloudflare_load_balancer: add state migrator for 3.x ([#1211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1211))
* resource/cloudflare_record: add state migrator for 3.x ([#1211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1211))

## 3.0.0 (September 20th, 2021)

[2.x to 3.x upgrade guide](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/guides/version-3-upgrade)

BREAKING CHANGES:

* resource/cloudflare_access_rule: `configuration` is now a `TypeList` instead of a `TypeMap` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_custom_ssl: `custom_ssl_options` is now a `TypeList` instead of `TypeMap` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_load_balancer: `fixed_response` is now a `TypeList` instead of a `TypeMap` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_load_balancer: fixed_response.status_code` is now a `TypeInt` instead of a `TypeString` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_record: `data` is now a `TypeList` instead of a `TypeMap` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))

NOTES:

* provider: Golang version has been upgraded to 1.17 ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* provider: HTTP user agent is now "terraform/:version terraform-plugin-sdk/:version terraform-provider-cloudflare/:version" ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* provider: Minimum Terraform core version is now 0.14 ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* provider: terraform-plugin-sdk has been upgraded to 2.x ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))

ENHANCEMENTS:

* resource/cloudflare_custom_hostname: `settings.ciphers` is now a `TypeSet` internally to handle suppress ordering changes. Schema representation remains the same ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_custom_hostname: `settings` is now `Optional`/`Computed` to reflect the stricter schema validation introduced in terraform-plugin-sdk v2 ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
* resource/cloudflare_custom_hostname: `status` is now `Computed` as the value isn't managed by an end user ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))

## 2.27.0 (September 20th, 2021)

NOTES:

* provider: Update to cloudflare-go v0.22.0 ([#1184](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1184))

FEATURES:

* **New Resource:** `cloudflare_access_keys_configuration` ([#1186](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1186))
* **New Resource:** `cloudflare_teams_account` ([#1173](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1173))
* **New Resource:** `cloudflare_teams_rule` ([#1173](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1173))

ENHANCEMENTS:

* resource/cloudflare_access_policy: add support for purpose justification and approvals ([#1199](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1199))
* resource/cloudflare_ruleset: add support for HTTP rate limiting ([#1179](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1179))
* resource/cloudflare_ruleset: add support for Transform Rules ([#1169](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1169))
* resource/cloudflare_ruleset: add support for WAF payload logging ([#1174](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1174))
* resource/cloudflare_ruleset: add support for more complex skip ruleset configurations ([#1201](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1201))

BUG FIXES:

* resource/cloudflare_ruleset: fix state handling for terraform-plugin-sdk v2 ([#1183](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1183))
* resource/cloudflare_zone_settings_override: remap `zero_rtt` => `0rtt` for resource delete ([#1175](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1175))

## 2.26.1 (August 30th, 2021)

**Fixes**

* `resource/cloudflare_ruleset`: Send a single payload for rules instead of many individual payloads to prevent overwriting previous rules ([#1171](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1171))
## 2.26.0 (August 27th, 2021)

* **New resource**: `cloudflare_notification_policy` ([#1138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1138))
* **New resource**: `cloudflare_notification_policy_webhooks` ([#1151](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1151))
* **New resource**: `cloudflare_ruleset` ([#1143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1143))
* **New resource**: `cloudflare_teams_location` ([#1154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1154))
* **New datasource**: `cloudflare_origin_ca_root_certificate` ([#1158](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1158))

**Improvements**

* `resource/cloudflare_waiting_room`: Add support for `json_response_enabled` as an argument ([#1122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1122))

## 2.25.0 (August 4th, 2021)

**Improvements**

* `resource/cloudflare_access_device_posture_rule`: Add support for `domain_joined`, `firewall`, `os_version`, and `disk_encryption` ([#1137](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1137))
* provider: bump `cloudflare-go` to v0.20.0 ([#1146](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1146))

## 2.24.0 (July 19th, 2021)

**Improvements**

* `resource/cloudflare_logpush_job`: Add support for `"nel_reports"` as a dataset ([#1122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1122))
* `resource/cloudflare_custom_hostname`: Allow SSL options to be optional when not required ([#1131](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1131))
* `resource/cloudflare_access_identity_provider`: Support optional Okta API token ([#1119](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1119))
* `resource/cloudflare_load_balancer_pool`: Add support for load shedding ([#1108](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1108))
* `resource/cloudflare_load_balancer_pool`: Add support for longitude and latitude ([#1093](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1093))

**Fixes**

* `resource/cloudflare_record`: Use correct `Import` method on resource ([#1116](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1116))
* `resource/cloudflare_worker_cron_trigger`: Account for deletion of scripts and force a refresh of triggers ([#1121](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1121))
* `resource/cloudflare_rate_limit`: Handle `origin_traffic` missing from API response ([#1125](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1125))
* `resource/cloudflare_record`: Support `allow_overwrite` for root records ([#1129](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1129))

## 2.23.0 (June 30th, 2021)

- **New resource**: `cloudflare_waiting_room` ([#1053](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1053))

**Improvements**

* `datasource/cloudflare_waf_rules`: Export `default_mode` as an attribute ([#1079](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1079))

**Fixes**

* `resource/cloudflare_access_application`: Revert removal of schema changes causing existing applications unable to re-apply ([#1118](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1118))

## 2.22.0 (June 25th, 2021)

- **New resource**: `cloudflare_static_route` ([#1098](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1098))

**Improvements**

* `resource/cloudflare_origin_ca`: Ignore decreasing `requested_validity` ([#1043](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1078))
* `resource/waf_override`: Allow `rules` to be optional ([#1090](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1090))
* `resource/cloudflare_zone`: Don't attempt to set free zone rate plans as that is already the default ([#1102](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1102))
* `resource/cloudflare_access_application`: Ability to set `type` for Applications ([#1076](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1076))
* `resource/cloudflare_zone_lockdown`: Update documentation to show examples of multiple configurations ([#1106](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1106))

## 2.21.0 (May 26th, 2021)

- **New resource**: `cloudflare_device_posture_rule` ([#1058](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1058))
- **New resource**: `cloudflare_teams_list` ([#1058](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1058))

**Improvements**

* provider: Update to terraform-plugin-sdk v1.17.1 ([#1035](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1035), [#1043](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1043))
* `resource/cloudflare_logpush_job`: Allow `ownership_challenge` to be optional to account for Datadog, Splunk or S3-Compatible endpoints ([#1048](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1048))
* `resource/cloudflare_access_group`: Add support for `login_method` ([#1066](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1066))
* `resource/cloudflare_load_balancer`: Add support for `promixity` based steering ([#1072](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1072))
* `resource/cloudflare_access_application`: Prevent bad CORS configuration when credentials and all origins are permitted ([#1073](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1073))
* `resource/cloudflare_access_service_tokens`: Allow configuration to manage automatic renewal when the threshold is crossed and Terraform operations are performed within the window ([#1057](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1057))
* `resource/cloudflare_load_balancer_pool`: Allow support for `Host` header settings ([#1042](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1042))

**Fixes**

* `resource/cloudflare_access_policy`: Allow empty slices in blocks when building policies ([#1034](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1034))
* `resource/cloudflare_load_balancer`: Fix `override` attributes `pop_pools` and `region_pools` referencing incorrect values causing a panic ([#1039](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1039))

## 2.20.0 (April 15th, 2021)

**New resource**: `cloudflare_access_ca_certificate` ([#995](https://github.com/cloudflare/terraform-provider-cloudflare/issues/995))

**Improvements**

* `resource/cloudflare_access_application`: Improve documentation for `Import` usage ([#1002](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1002))
* `resource/cloudflare_logpush_job`: Update documentation to reflect requirements for `destination_conf` to match across all uses ([#1024](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1024))
* `resource/cloudflare_custom_hostname_fallback`: Better handle service lag when updating existing resources by attempting retries ([#1014](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1014))
* `resource/cloudflare_waf_group`: Simplify error handling using inbuilt helpers ([#1015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1015))
* `resource/cloudflare_waf_rule`: Simplify error handling using inbuilt helpers ([#1015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1015))
* `resource/cloudflare_waf_package`: Simplify error handling using inbuilt helpers ([#1015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1015))
* `resource/cloudflare_access_group`: Add support for `login_method` ([#1018](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1018))
* provider: Update to cloudflare-go v0.16.0 ([#1018](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1018))
* provider: Update to terraform-plugin-sdk v1.16.1 ([#1003](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1003))
* `resource/cloudflare_load_balancer`: Add support for `rules` ([#1016](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1016))

## 2.19.2 (March 15th, 2021)

**Fixes**

* `resource/cloudflare_record`: Address regression from 2.19.1 by checking the API response instead of the schema output for `Priority` ([#992](https://github.com/cloudflare/terraform-provider-cloudflare/issues/992))

## 2.19.1 (March 11th, 2021)

**Fixes**

* `resource/cloudflare_record`: Update `Priority` handling for MX parked records ([#986](https://github.com/cloudflare/terraform-provider-cloudflare/issues/986))

## 2.19.0 (March 10th, 2021)

**Fixes**

* `resource/cloudflare_access_group`: Fix crash when constructing a GSuite group ([#940](https://github.com/cloudflare/terraform-provider-cloudflare/issues/940))
* `resource/cloudflare_access_policy`: Make `precedence` required ([#941](https://github.com/cloudflare/terraform-provider-cloudflare/issues/941))
* `resource/cloudflare_access_group`: Fix crash when constructing a SAML group ([#948](https://github.com/cloudflare/terraform-provider-cloudflare/issues/948))
* `resource/cloudflare_zone`: Update `Retry` logic to look at an available field for passing conditions ([#973](https://github.com/cloudflare/terraform-provider-cloudflare/issues/973))
* `resource/cloudflare_page_rule`: Allow ignoring/including all query string parameters for `cache_key_fields` ([#975](https://github.com/cloudflare/terraform-provider-cloudflare/issues/975))

**Improvements**

* `resource/cloudflare_access_policy`: Enable zone and account level resources to be imported ([#956](https://github.com/cloudflare/terraform-provider-cloudflare/issues/956))
* `resource/cloudflare_origin_ca_certificate`: Smoother import process with less recreation ([#955](https://github.com/cloudflare/terraform-provider-cloudflare/issues/955))
* provider: Update internals to match `cloudflare-go` 0.14 for better error handling and context aware methods ([#976](https://github.com/cloudflare/terraform-provider-cloudflare/issues/976))

## 2.18.0 (February 3rd, 2021)

* **New Resource:** `cloudflare_argo_tunnel` ([#905](https://github.com/cloudflare/terraform-provider-cloudflare/issues/905))
* **New Resource:** `cloudflare_worker_cron_trigger` ([#926](https://github.com/cloudflare/terraform-provider-cloudflare/issues/926))

**Fixes**

* `datasource/cloudflare_zones`: Pagination is now correctly handled internally and will return more than the single page of results ([cloudflare/cloudflare-go#534](https://github.com/cloudflare/cloudflare-go/pull/534)).
* `resource/cloudflare_access_policy`: Correctly handle transforming API responses to schema ([#917](https://github.com/cloudflare/terraform-provider-cloudflare/issues/917))
* `resource/cloudflare_access_group`: Correctly handle transforming API responses to schema ([#918](https://github.com/cloudflare/terraform-provider-cloudflare/issues/918))
* `resource/cloudflare_ip_list`: Ensure account ID is persisted during `Import` ([#916](https://github.com/cloudflare/terraform-provider-cloudflare/issues/916))

**Improvements**

* `resource/cloudflare_access_application`: Allow any `session_duration` that is `time.ParseDuration` compatible ([#910](https://github.com/cloudflare/terraform-provider-cloudflare/issues/910))
* `resource/cloudflare_rate_limit`: Add the ability to configure `match.response.headers` in rate limits ([#911](https://github.com/cloudflare/terraform-provider-cloudflare/issues/911))
* `resource/cloudflare_access_rule`: Validate IP masks within schema ([#921](https://github.com/cloudflare/terraform-provider-cloudflare/issues/921))

## 2.17.0 (January 5th, 2021)

* **New Resource:** `cloudflare_magic_firewall_ruleset` ([#884](https://github.com/cloudflare/terraform-provider-cloudflare/issues/884))

**Fixes**

* `resource/cloudfare_api_token`: Omitting `conditions` will no longer send empty arrays causing IP restriction issues and unusable tokens ([#902](https://github.com/cloudflare/terraform-provider-cloudflare/pull/902))

## 2.16.0 (January 5th, 2021)

**Improvements**

* `resource/cloudflare_access_application`: Add support for `custom_deny_message` and `custom_deny_url` values ([#895](https://github.com/cloudflare/terraform-provider-cloudflare/issues/895))
* `resource/cloudflare_load_balancer_monitor`: Add support for `probe_zone` for monitors ([#903](https://github.com/cloudflare/terraform-provider-cloudflare/issues/903))

## 2.15.0 (December 29th, 2020)

**Improvements**

* `resource/cloudflare_load_balancer`: Add support for `session_affinity_ttl` ([#882](https://github.com/cloudflare/terraform-provider-cloudflare/issues/882))
* `resource/cloudflare_load_balancer`: Add support for `session_affinity_attributes` ([#883](https://github.com/cloudflare/terraform-provider-cloudflare/issues/883))

**Fixes**

* `resource/cloudflare_page_rule`: Fixed crash during update when using custom cache key ([#894](https://github.com/cloudflare/terraform-provider-cloudflare/pull/894))

## 2.14.0 (November 26th, 2020)

* **New Resource:** `cloudflare_api_token` ([#862](https://github.com/cloudflare/terraform-provider-cloudflare/issues/862))
* **New Datasource:** `cloudflare_api_token_permission_groups` ([#862](https://github.com/cloudflare/terraform-provider-cloudflare/issues/862))
* **New Resource:** `cloudflare_zone_dnssec` ([#852](https://github.com/cloudflare/terraform-provider-cloudflare/issues/852))
* **New Datasource:** `cloudflare_zone_dnssec` ([#852](https://github.com/cloudflare/terraform-provider-cloudflare/issues/852))

**Improvements**

* `resource/cloudflare_record`: Add explicit fields for CAA records instead of relying on the map value ([#866](https://github.com/cloudflare/terraform-provider-cloudflare/issues/866))
* `resource/cloudflare_account_member`: Swap schema `role_ids` to `TypeSet` to better handle internal ordering changes ([#876](https://github.com/cloudflare/terraform-provider-cloudflare/issues/876))

**Fixes**

* `datasource/cloudflare_waf_groups`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))
* `datasource/cloudflare_waf_packages`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))
* `datasource/cloudflare_waf_rules`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))
* `datasource/cloudflare_zones`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))

## 2.13.2 (November 6th, 2020)

**Fixes**

* `resource/cloudflare_filter`: Remove schema based validation for filters ([#863](https://github.com/cloudflare/terraform-provider-cloudflare/issues/863))

## 2.13.1 (November 5th, 2020)

**Improvements**

* `resource/cloudflare_filter`: Pass missing credential error through to end user ([#860](https://github.com/cloudflare/terraform-provider-cloudflare/issues/860))

## 2.13.0 (November 5th, 2020)

**Improvements**

* `datasource/cloudflare_ip_ranges`: Add the ability to query `china_ipv4_cidr_blocks` and `china_ipv6_cidr_blocks` ([#833](https://github.com/cloudflare/terraform-provider-cloudflare/issues/833))
* `resource/cloudflare_filter`: Improve validation of expressions using the schema ([#848](https://github.com/cloudflare/terraform-provider-cloudflare/issues/848))

**Fixes**

* `resource/cloudflare_page_rule`: Set default for `cache_key_fields.host.resolved` to prevent panics ([#832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/832))
* `resource/cloudflare_authenticated_origin_pulls`: Fix off-by-one error check in `Import` ([#832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/859))
* `resource/cloudflare_authenticated_origin_pulls_certificate`: Fix off-by-one error check in `Import` ([#832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/859))

## 2.12.0 (October 22nd, 2020)

**Improvements**

* `resource/cloudflare_certificate_pack`: Swap internal representation of `hosts` to remove inconsistent ordering issues ([#800](https://github.com/cloudflare/terraform-provider-cloudflare/issues/800))
* `resource/cloudflare_logpush_job`: Handle deletion outside of Terraform ([#798](https://github.com/cloudflare/terraform-provider-cloudflare/issues/798))
* `resource/cloudflare_access_group`: Add support for `geo` conditionals ([#803](https://github.com/cloudflare/terraform-provider-cloudflare/issues/803))
* `resource/cloudflare_access_application`: Add support for `enable_binding_cookie` ([#802](https://github.com/cloudflare/terraform-provider-cloudflare/issues/802))
* `resource/cloudflare_waf_rule`: Improve documentation for `mode` ([#824](https://github.com/cloudflare/terraform-provider-cloudflare/issues/824))
* `datasource/cloudflare_waf_rule`: Improve documentation for `mode` ([#824](https://github.com/cloudflare/terraform-provider-cloudflare/issues/824))
* `resource/cloudflare_access_application`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))
* `resource/cloudflare_access_group`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))
* `resource/cloudflare_access_identity_provider`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))
* `resource/cloudflare_access_policy`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))

**Fixes**

* `resource/cloudflare_custom_hostname_fallback_origin`: Don't retry the "active" status of custom hostnames fallbacks ([#818](https://github.com/cloudflare/terraform-provider-cloudflare/issues/818))
* `resource/cloudflare_zone`: Remove `DiffSuppressFunc` causing `jump_start` issues ([#830](https://github.com/cloudflare/terraform-provider-cloudflare/issues/830))

## 2.11.0 (September 11th, 2020)

* **New Resource:** `cloudflare_certificate_pack` ([#778](https://github.com/cloudflare/terraform-provider-cloudflare/issues/778))

**Improvements**

* `resource/cloudflare_access_group`: Add support for `auth_method` ([#762](https://github.com/cloudflare/terraform-provider-cloudflare/issues/762))
* `resource/cloudflare_access_group`: De-duplicate blocks in groups by accepting lists instead ([#739](https://github.com/cloudflare/terraform-provider-cloudflare/issues/739))
* `resource/cloudflare_worker_script`: Adds support for `webassembly_binding` ([#780](https://github.com/cloudflare/terraform-provider-cloudflare/issues/780))
* `resource/cloudflare_healthcheck`: Retry hostname resolution errors when encountering "no such host" responses ([#789](https://github.com/cloudflare/terraform-provider-cloudflare/issues/789))
* `resource/cloudflare_access_application`: Better validation for allowed methods and origin combinations to prevent getting state into an unrecoverable state ([#793](https://github.com/cloudflare/terraform-provider-cloudflare/issues/793))

**Fixes**

* `resource/cloudflare_healthcheck`: Handle resource deletion outside of Terraform ([#787](https://github.com/cloudflare/terraform-provider-cloudflare/issues/787))
* `resource/cloudflare_custom_hostname`: Ensure `Import` sets hostname to prevent recreation ([#788](https://github.com/cloudflare/terraform-provider-cloudflare/issues/788))
* `resource/cloudflare_ip_list`: Handle resource deletion outside of Terraform ([#794](https://github.com/cloudflare/terraform-provider-cloudflare/issues/794))
* `resource/cloudflare_ip_list`: Remove `item`.`id` from schema ([#796](https://github.com/cloudflare/terraform-provider-cloudflare/issues/796))

## 2.10.1 (August 24th, 2020)

**Fixes**

* `resource/cloudflare_access_application`: Handle the `zone_id` => `account_id` move internally ([#724](https://github.com/cloudflare/terraform-provider-cloudflare/issues/724))

## 2.10.0 (August 24th, 2020)

* **New Resource:** `cloudflare_custom_hostname_origin_fallback` ([#757](https://github.com/cloudflare/terraform-provider-cloudflare/issues/757))
* **New Resource:** `cloudflare_authenticated_origin_pulls` ([#749](https://github.com/cloudflare/terraform-provider-cloudflare/issues/749))
* **New Resource:** `cloudflare_authenticated_origin_pulls_certificate` ([#749](https://github.com/cloudflare/terraform-provider-cloudflare/issues/749))
* **New Resource:** `cloudflare_ip_list` ([#766](https://github.com/cloudflare/terraform-provider-cloudflare/issues/766))

**Improvements**

* `resource/cloudflare_spectrum_application`: Add support for port ranges ([#745](https://github.com/cloudflare/terraform-provider-cloudflare/issues/745))
* `resource/cloudflare_custom_hostname`: Force creation of a new resource if the `zone_id` value changes ([#761](https://github.com/cloudflare/terraform-provider-cloudflare/issues/761))
* `resource/cloudflare_record`: Retry record creation/update if the response includes an "already exists" exception for handling race conditions ([#773](https://github.com/cloudflare/terraform-provider-cloudflare/issues/773))

**Fixes**

* `resource/cloudflare_firewall_rule`: Compare descriptions after converting unicode + HTML entities to prevent unnecessary diffs ([#758](https://github.com/cloudflare/terraform-provider-cloudflare/issues/758))
* `resource/cloudflare_filter`: Compare descriptions after converting unicode + HTML entities to prevent unnecessary diffs ([#758](https://github.com/cloudflare/terraform-provider-cloudflare/issues/758))

## 2.9.0 (July 30th, 2020)

* **New Resource:** `cloudflare_custom_hostname` (SSL for SaaS) ([#746](https://github.com/cloudflare/terraform-provider-cloudflare/issues/746))

**Improvements**

* `resource/access_application`: Add support for `allowed_idps` and restricting which Identity Providers are associated with an Application ([#734](https://github.com/cloudflare/terraform-provider-cloudflare/issues/734))
* `resource/access_application`: Add support for `auto_redirect_to_identity` ([#730](https://github.com/cloudflare/terraform-provider-cloudflare/issues/730))
* `resource/access_application`: Add CORS support ([#725](https://github.com/cloudflare/terraform-provider-cloudflare/issues/725))
* `resource/cloudflare_custom_ssl`: Allow `geo_restrictions` to be `nil` and not included in the request payload ([#714](https://github.com/cloudflare/terraform-provider-cloudflare/issues/714))
* `datasource/cloudflare_zones`: Filtering is now performed on the server side and the `name` parameter is no longer a regex. Instead, `name` is a string to match on and `match` is a regex. See the website documentation for more examples and updated references ([#708](https://github.com/cloudflare/terraform-provider-cloudflare/issues/708)) in order to make your code compatible with this release.


## 2.8.0 (June 22, 2020)

* **New Resource:** `cloudflare_waf_override` ([#691](https://github.com/cloudflare/terraform-provider-cloudflare/issues/691))

**Improvements**

* `resource/cloudflare_argo`: Allow `tiered_caching` and `smart_routing` to be toggled individually allowing for entitlement differences ([#703](https://github.com/cloudflare/terraform-provider-cloudflare/issues/703))
* `resource/cloudflare_page_rule`: Add support for `cache_ttl_by_status` ([#706](https://github.com/cloudflare/terraform-provider-cloudflare/issues/706))
* `resource/cloudflare_worker_script`: Add support for `plain_text` and `secret_text` bindings ([#710](https://github.com/cloudflare/terraform-provider-cloudflare/issues/710))

**Fixes**

* `resource/cloudflare_record`: Update `TestAccCloudflareRecord_LOC` test asserted value to use less precise floats and match the API responses ([#712](https://github.com/cloudflare/terraform-provider-cloudflare/issues/712))
* `resource/cloudflare_record`: Update `TestAccCloudflareRecord_Basic` test `metadata` attributes to match updated API payload ([#713](https://github.com/cloudflare/terraform-provider-cloudflare/issues/713))

## 2.7.0 (May 20, 2020)

* **New Resource:** `cloudflare_byo_ip_prefix` ([#671](https://github.com/cloudflare/terraform-provider-cloudflare/issues/671))
* **New Resource:** `cloudflare_logpull_retention` ([#678](https://github.com/cloudflare/terraform-provider-cloudflare/issues/678))
* **New Resource:** `cloudflare_healthcheck` ([#680](https://github.com/cloudflare/terraform-provider-cloudflare/issues/680))

**Improvements:**

* `resource/cloudflare_worker_route`: Improve documentation to mention using `account_id` for the underlying APIs ([#669](https://github.com/cloudflare/terraform-provider-cloudflare/issues/669))
* `resource/cloudflare_worker_script`: Improve documentation to mention using `account_id` for the underlying APIs ([#670](https://github.com/cloudflare/terraform-provider-cloudflare/issues/670))
* `resource/cloudflare_load_balancer_pool`: Improve documentation to mention `notification_email` accepts a comma delimited list of emails ([#687](https://github.com/cloudflare/terraform-provider-cloudflare/issues/687))
* `resource/cloudflare_page_rule`: Add support for `cache_key_fields` Page Rule action ([#662](https://github.com/cloudflare/terraform-provider-cloudflare/issues/662))

**Fixes:**
* `resource/cloudflare_zone_settings_override`: Fix regression where if you didn't have universal SSL settings defined, it would error when setting them ([#663](https://github.com/cloudflare/terraform-provider-cloudflare/issues/663))
* `resource/cloudflare_zone`: Handle changing zone rate plan from "free" to "enterprise" ([#668](https://github.com/cloudflare/terraform-provider-cloudflare/issues/668))
* `resource/cloudflare_record`: Update validation to allow PTR records ([9a8fd43](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9a8fd43))

## 2.6.0 (April 22, 2020)

**Improvements:**

* `resource/cloudflare_zone_settings_override`: Add `universal_ssl` to control enablement of Universal SSL on a zone ([#658](https://github.com/cloudflare/terraform-provider-cloudflare/issues/658))
* provider: API keys and API tokens are now validated to help differentiate incorrect usage before making API calls ([#661](https://github.com/cloudflare/terraform-provider-cloudflare/issues/661))
* `resource/cloudflare_logpush_job`: Add support for "firewall_events" dataset parameter ([#660](https://github.com/cloudflare/terraform-provider-cloudflare/issues/660))
* `resource/cloudflare_logpush_job`: Add support for "dataset" parameter ([#649](https://github.com/cloudflare/terraform-provider-cloudflare/issues/649))
* `resource/cloudflare_zone_settings_override`: Remove `edge_cache_ttl` ([#654](https://github.com/cloudflare/terraform-provider-cloudflare/issues/654))
* `resource/cloudflare_access_group`: Allow Access conditions for `include`/`require`/`exclude` to be used consistently between Access Groups and Access Policies ([#646](https://github.com/cloudflare/terraform-provider-cloudflare/issues/646))

**Fixes:**
* `resource/cloudflare_logpush_job`: fix for `strconv.Atoi: parsing ""` error while creating Logpush job

## 2.5.1 (April 03, 2020)

**Improvements:**

* `resource/cloudflare_zone_settings_override`: Update `image_resizing` options to include `"open"` ([#639](https://github.com/cloudflare/terraform-provider-cloudflare/issues/639))

**Fixes:**

* `resource/cloudflare_access_group`: Fixed misspelt Okta in JSON payload ([cloudflare/cloudflare-go#440](https://github.com/cloudflare/cloudflare-go/issues/440))

## 2.5.0 (March 27, 2020)

**Improvements:**

* `resource/cloudflare_access_policy`: Add support for `service_token` and `any_valid_service_token` ([#612](https://github.com/cloudflare/terraform-provider-cloudflare/issues/612))
* `resource/cloudflare_waf_group`: Handle WAF group deletions in the API responses ([#623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/623))
* `resource/cloudflare_waf_package`: Handle WAF package deletions in the API responses ([#623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/623))
* `resource/cloudflare_waf_rule`: Handle WAF rule deletions in the API responses ([#623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/623))
* `resource/cloudflare_access_policy`: Add support for `group` ([#626](https://github.com/cloudflare/terraform-provider-cloudflare/issues/626))
* `resource/cloudflare_firewall_rule`: Add support for bypassing specific `products` ([#630](https://github.com/cloudflare/terraform-provider-cloudflare/issues/630))
* `resource/cloudflare_spectrum_application`: Add support for `edge_ips`, `argo_smart_routing` and `edge_ip_connectivity` ([#631](https://github.com/cloudflare/terraform-provider-cloudflare/issues/631))
* `resource/cloudflare_access_group`: Add support for using external providers (`gsuite`, `github`, `azure`, `okta`, `saml`, `mTLS certificate`, `common name`
) ([#633](https://github.com/cloudflare/terraform-provider-cloudflare/issues/633))

## 2.4.1 (March 12, 2020)

**Improvements:**

* `resource/cloudflare_logpush_job`: Support `Import` on the resource ([#618](https://github.com/cloudflare/terraform-provider-cloudflare/issues/618))

**Fixes:**

* `resource/cloudflare_record`: Missing CAA in DNS validation ([#619](https://github.com/cloudflare/terraform-provider-cloudflare/issues/619))

## 2.4.0 (March 09, 2020)

* **New Resource:** `cloudflare_workers_kv` ([#595](https://github.com/cloudflare/terraform-provider-cloudflare/issues/595))
* **New Resource:** `cloudflare_access_identity_provider` ([#597](https://github.com/cloudflare/terraform-provider-cloudflare/issues/597))


**Improvements:**

* `resource/cloudflare_record`: Stricter validation for record types ([#610](https://github.com/cloudflare/terraform-provider-cloudflare/issues/610))
* `resource/logpush_job`: Add more verbose error handling ([#564](https://github.com/cloudflare/terraform-provider-cloudflare/issues/564))
* `resource/zone_settings_override`: Update documentation for `cache_level` values ([#606](https://github.com/cloudflare/terraform-provider-cloudflare/issues/606))
* `resource/access_application`: Add documentation for available attributes ([#587](https://github.com/cloudflare/terraform-provider-cloudflare/issues/587))
* `resource/cloudflare_firewall_rule`: Add support for bypassing security configuration rules by URL ([#568](https://github.com/cloudflare/terraform-provider-cloudflare/issues/568))
* `resource/cloudflare_record_migrate`: Use `zone_id` for state migration before attempting to use `domain` ([#566](https://github.com/cloudflare/terraform-provider-cloudflare/issues/566))
* `resource/cloudflare_load_balancer`: Update `session_affinity` validation to allow `"ip_cookie"` ([#573](https://github.com/cloudflare/terraform-provider-cloudflare/issues/573))
* `datasource/ip_ranges`: Update documentation to show 0.12 syntax ([#617](https://github.com/cloudflare/terraform-provider-cloudflare/issues/617))

**Fixes**

* `resource/zone_settings_override`: Handle individual zone settings within `Delete` operations ([#599](https://github.com/cloudflare/terraform-provider-cloudflare/issues/599))

## 2.3.0 (December 18, 2019)

* **New Resource:** `cloudflare_origin_ca_certificate` ([#547](https://github.com/cloudflare/terraform-provider-cloudflare/issues/547))

**Fixes:**

* `resource/cloudflare_zone_settings_override`: Renamed `0rtt` to `zero_rtt` to conform to HCL grammar requirements ([#557](https://github.com/cloudflare/terraform-provider-cloudflare/issues/557))

**Improvements:**

* `resource/cloudflare_access_rule`: Add `ip6` as valid option ([#560](https://github.com/cloudflare/terraform-provider-cloudflare/issues/560))
* `resource/cloudflare_spectrum_application`: Swap `proxy_protocol` to string field with supporting enum values instead ([#561](https://github.com/cloudflare/terraform-provider-cloudflare/issues/561))
* `resource/cloudflare_waf_rule`: Add `package_id` as valid option and export `group_id` ([#552](https://github.com/cloudflare/terraform-provider-cloudflare/issues/552))

## 2.2.0 (December 05, 2019)

* **New Resource:** `cloudflare_access_group` ([#510](https://github.com/cloudflare/terraform-provider-cloudflare/issues/510))
* **New Resource:** `cloudflare_workers_kv_namespace` ([#443](https://github.com/cloudflare/terraform-provider-cloudflare/issues/443))

**Improvements:**

* `resource/cloudflare_zone_settings_override`: Add `non_identity` to allowed `decision` schema ([#541](https://github.com/cloudflare/terraform-provider-cloudflare/issues/541))
* `resource/cloudflare_zone_settings_override`: Add support for `0rtt` and `http3` settings ([#542](https://github.com/cloudflare/terraform-provider-cloudflare/issues/542))
* `resource/cloudflare_load_balancer_monitor`: Allow empty string for `expected_body` ([#539](https://github.com/cloudflare/terraform-provider-cloudflare/issues/539))
* `resource/cloudflare_worker_script`: Add support for Worker KV Namespace Bindings ([#544](https://github.com/cloudflare/terraform-provider-cloudflare/issues/544))
* `data_source/waf_rules`, `resource/cloudflare_waf_rule`, Support allowed modes for WAF Rules ([#550](https://github.com/cloudflare/terraform-provider-cloudflare/issues/550))

**Fixes:**

* `resource/cloudflare_spectrum_application`: Spectrum origin_port is optional ([#549](https://github.com/cloudflare/terraform-provider-cloudflare/issues/549))

## 2.1.0 (November 07, 2019)

* **New datasource:** `cloudflare_waf_rules` ([#525](https://github.com/cloudflare/terraform-provider-cloudflare/issues/525))

**Improvements:**

* `resource/cloudflare_zone`: Expose `verification_key` for partial setups ([#532](https://github.com/cloudflare/terraform-provider-cloudflare/issues/532))
* `resource/cloudflare_worker_route`: Enable API Tokens support from upstream [cloudflare-go](https://github.com/cloudflare/cloudflare-go) release

## 2.0.1 (October 22, 2019)

* **New Resource:** `cloudflare_access_service_tokens` ([#521](https://github.com/cloudflare/terraform-provider-cloudflare/issues/521))
* **New Resource:** `cloudflare_waf_package` ([#475](https://github.com/cloudflare/terraform-provider-cloudflare/issues/475))
* **New Resource:** `cloudflare_waf_group` ([#476](https://github.com/cloudflare/terraform-provider-cloudflare/issues/476))
* **New datasource:** `cloudflare_waf_groups` ([#508](https://github.com/cloudflare/terraform-provider-cloudflare/issues/508))
* **New datasource:** `cloudflare_waf_packages` ([#509](https://github.com/cloudflare/terraform-provider-cloudflare/issues/509))

**Fixes:**

* `resource/cloudflare_page_rule`: Set `h2_prioritization` individually not via bulk endpoint ([#493](https://github.com/cloudflare/terraform-provider-cloudflare/issues/493))
* `resource/cloudflare_zone_settings_override`: Set `zone_id` to prevent unnecessary re-creation of resources ([#502](https://github.com/cloudflare/terraform-provider-cloudflare/issues/502))

**Improvements:**

* `resource/cloudflare_spectrum_application`: Add support for setting `traffic_type` ([#481](https://github.com/cloudflare/terraform-provider-cloudflare/issues/481))
* `resource/cloudflare_zone_settings_override`: Update documentation with default values ([#498](https://github.com/cloudflare/terraform-provider-cloudflare/issues/498))

**Internals:**

* Migrated to Terraform plugin SDK ([#489](https://github.com/cloudflare/terraform-provider-cloudflare/issues/489))

## 2.0.0 (September 30, 2019)

**Breaking changes:**
* `provider/cloudflare`:
 * renamed `token` to `api_key`
 * renamed `org_id` to `account_id`
 * removed `use_org_from_zone`, you need to explicitly specify `account_id`
* Environment variables:
 * renamed `CLOUDFLARE_TOKEN` to `CLOUDFLARE_API_TOKEN`
 * renamed `CLOUDFLARE_ORG_ID` to `CLOUDFLARE_ACCOUNT_ID`
 * removed `CLOUDFLARE_ORG_ZONE`, you need to explicitly specify `CLOUDFLARE_ACCOUNT_ID`
* Changed the following resources to require Zone ID:
 * `cloudflare_access_rule`
 * `cloudflare_filter`
 * `cloudflare_firewall_rule`
 * `cloudflare_load_balancer`
 * `cloudflare_page_rule`
 * `cloudflare_rate_limit`
 * `cloudflare_record`
 * `cloudflare_waf_rule`
 * `cloudflare_worker_route"`
 * `cloudflare_zone_lockdown`
 * `cloudflare_zone_settings_override`
* Workers single-script support removed

Please see [Version 2 Upgrade Guide](https://www.terraform.io/docs/providers/cloudflare/guides/version-2-upgrade.html) for details.

**Improvements:**

* `cloudflare/resource_cloudflare_argo`: Handle errors when fetching tiered caching + smart routing settings ([#477](https://github.com/cloudflare/terraform-provider-cloudflare/issues/477))
* Various documentation updates for 0.12 syntax

## 1.18.1 (August 29, 2019)

**Fixes:**

* `resource/cloudflare_load_balancer`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_page_rule`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_rate_limit`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_waf_rule`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_worker_route`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_worker_script`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_zone_lockdown`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))

## 1.18.0 (August 27, 2019)

**Fixes:**

* `resource/cloudflare_page_rule`: Fix a logic condition where setting `edge_cache_ttl` action but then not updating it in subsequent `apply` runs causes it to be blown away ([#453](https://github.com/cloudflare/terraform-provider-cloudflare/issues/453))

**Improvements:**

* provider: You can now use API tokens to authenticate instead of user email and key ([#450](https://github.com/cloudflare/terraform-provider-cloudflare/issues/450))
* `resource/cloudflare_zone_lockdown`: `priority` can now be set on the resource ([#445](https://github.com/cloudflare/terraform-provider-cloudflare/issues/445))
* `resource/cloudflare_custom_ssl`: Updated website documentation navigation to include link for resource ([#442)](https://github.com/cloudflare/terraform-provider-cloudflare/issues/442))

**Deprecations:**

* `resource/cloudflare_access_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_filter`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_firewall_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_load_balancer`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_page_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_rate_limit`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_waf_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_worker_route`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_worker_script`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_zone_lockdown`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))

## 1.17.1 (August 09, 2019)

**Fixes:**

* Partially revert [[#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421)] deprecation messages

## 1.17.0 (August 09, 2019)

**Removals:**

* `resource/cloudflare_zone_settings_override`: `sha1_support` has been removed due to Cloudflare no longer supporting SHA1 certificates or the API endpoint ([#415](https://github.com/cloudflare/terraform-provider-cloudflare/issues/415))

**Deprecations:**

* `resource/cloudflare_zone_settings_override`: `tls_1_2_only` has been superseded by using `min_tls_version` instead ([#405](https://github.com/cloudflare/terraform-provider-cloudflare/issues/405))
* `resource/cloudflare_access_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_filter`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_firewall_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_load_balancer`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_page_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_rate_limit`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_waf_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_worker_route`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_worker_script`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_zone_lockdown`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))

**Improvements:**

* **New Resource:** `cloudflare_custom_ssl` ([#418](https://github.com/cloudflare/terraform-provider-cloudflare/issues/418))
* `resource/cloudflare_filter`: Strip all surrounding whitespace from filter expressions to match API responses ([#361](https://github.com/cloudflare/terraform-provider-cloudflare/issues/361))
* `resource/cloudflare_zone`: Support unicode zone name values ([#412](https://github.com/cloudflare/terraform-provider-cloudflare/issues/412))
* `resource/cloudflare_page_rule`: Allow setting `origin_pull` for SSL ([#430](https://github.com/cloudflare/terraform-provider-cloudflare/issues/430))
* `resource/cloudflare_load_balancer_monitor`: Add TCP support for load balancer monitor ([#428](https://github.com/cloudflare/terraform-provider-cloudflare/issues/428))

**Fixes:**
* `resource/cloudflare_logpush_job`: Update documentation ([#395](https://github.com/cloudflare/terraform-provider-cloudflare/issues/395))
* `resource/cloudflare_zone_lockdown`: Fix: examples in documentation ([#407](https://github.com/cloudflare/terraform-provider-cloudflare/issues/407))
* `resource/cloudflare_page_rule`: Set nil on changed string-based Page Rule actions

## 1.16.1 (June 27, 2019)

**Fixes:**

* `resource/cloudflare_page_rule`: Fix regression in `browser_cache_ttl` where the value was sent as a string instead of an integer to the remote ([#390](https://github.com/cloudflare/terraform-provider-cloudflare/issues/390))

## 1.16.0 (June 20, 2019)

**Improvements:**

* `resource/cloudflare_zone_settings_override`: Add support for `h2_prioritization` and `image_resizing` ([#381](https://github.com/cloudflare/terraform-provider-cloudflare/issues/381))
* `resource/cloudflare_load_balancer_pool`: Update IP range for tests to not use reserved ranges ([#369](https://github.com/cloudflare/terraform-provider-cloudflare/issues/369))

**Fixes:**

* `resource/cloudflare_page_rule`: Fix issues with `browser_cache_ttl` defaults and when value is `0` (for Enterprise users)   ([#379](https://github.com/cloudflare/terraform-provider-cloudflare/issues/379))

## 1.15.0 (May 24, 2019)

* The provider is now compatible with Terraform v0.12, while retaining compatibility with prior versions. ([#309](https://github.com/cloudflare/terraform-provider-cloudflare/issues/309))

## 1.14.0 (May 15, 2019)

**Improvements:**

* **New Resource:** `cloudflare_argo` Manage Argo features ([#304](https://github.com/cloudflare/terraform-provider-cloudflare/issues/304))
* `cloudflare_zone`: Support management of partial zones ([#303](https://github.com/cloudflare/terraform-provider-cloudflare/issues/303))
* `cloudflare_rate_limit`: Update `modes` documentation ([#293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/212))
* `cloudflare_load_balancer`: Allow steering policy of "random" ([#329](https://github.com/cloudflare/terraform-provider-cloudflare/issues/329))

**Fixes:**

* `cloudflare_page_rule` - Allow setting `browser_cache_ttl` to 0 ([#293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/291))
* `cloudflare_page_rule` - Swap to completely replacing rules ([#338](https://github.com/cloudflare/terraform-provider-cloudflare/issues/338))

## 1.13.0 (April 12, 2019)

**Improvements**

* **New Resource:** `cloudflare_logpush_job` ([#287](https://github.com/cloudflare/terraform-provider-cloudflare/issues/287))
* `cloudflare_zone_settings` - Remove option to toggle `always_on_ddos` ([#253](https://github.com/cloudflare/terraform-provider-cloudflare/issues/253))
* `cloudflare_page_rule` - Update documentation to clarify "0" usage
* `cloudflare_zones` - Return zone ID and zone name ([#275](https://github.com/cloudflare/terraform-provider-cloudflare/issues/275))
* `cloudflare_load_balancer` - Add `enabled` field ([#208](https://github.com/cloudflare/terraform-provider-cloudflare/issues/208))
* `cloudflare_record` - validators: Allow PTR DNS records ([#283](https://github.com/cloudflare/terraform-provider-cloudflare/issues/283))

**Fixes:**

* `cloudflare_custom_pages` - Use correct casing for `zone_id` lookups
* `cloudflare_rate_limit` - Make `correlate` optional and not flap in state management ([#271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/271))
* `cloudflare_spectrum_application` - Fixed integration tests to work ([#275](https://github.com/cloudflare/terraform-provider-cloudflare/issues/275))
* `cloudflare_page_rule` - Better track field changes in `actions` resource. ([#107](https://github.com/cloudflare/terraform-provider-cloudflare/issues/107))

## 1.12.0 (March 07, 2019)

**Improvements:**

* provider: Enable request/response logging ([#212](https://github.com/cloudflare/terraform-provider-cloudflare/issues/212))
* resource/cloudflare_load_balancer_monitor: Add validation for `port` ([#213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/213))
* resource/cloudflare_load_balancer_monitor: Add `allow_insecure` and `follow_redirects` ([#205](https://github.com/cloudflare/terraform-provider-cloudflare/issues/205))
* resource/cloudflare_page_rule: Updated available actions documentation to match what is available ([#228](https://github.com/cloudflare/terraform-provider-cloudflare/issues/228))
* provider: Swap to using go modules for dependency management ([#230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/230))
* provider: Minimum Go version for development is now 1.11 ([#230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/230))

**Fixes:**

* resource/cloudflare_record: Read `data` back from API correctly ([#217](https://github.com/cloudflare/terraform-provider-cloudflare/issues/217))
* resource/cloudflare_rate_limit: Read `correlate` back from API correctly ([#204](https://github.com/cloudflare/terraform-provider-cloudflare/issues/204))
* resource/cloudflare_load_balancer_monitor: Fix incorrect type cast for `port` ([#213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/213))
* resource/cloudflare_load_balancer: Make `steering_policy` computed to avoid spurious diffs ([#214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/214))
* resource/cloudflare_load_balancer: Read `session_affinity` back from API to make import work & detects drifts ([#214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/214))

## 1.11.0 (January 11, 2019)

**Improvements:**
* **New Resource:** `cloudflare_spectrum_app` ([#156](https://github.com/cloudflare/terraform-provider-cloudflare/issues/156))
* **New Data Source:** `cloudflare_zones` ([#168](https://github.com/cloudflare/terraform-provider-cloudflare/issues/168))
* `cloudflare_load_balancer_monitor` - Add optional `port` parameter ([#179](https://github.com/cloudflare/terraform-provider-cloudflare/issues/179))
* `cloudflare_page_rule` - Improved documentation for `priority` attribute ([#182](https://github.com/cloudflare/terraform-provider-cloudflare/issues/182)], missing `explicit_cache_control` [[#185](https://github.com/cloudflare/terraform-provider-cloudflare/issues/185))
* `cloudflare_rate_limit` - Add `challenge` and `js_challenge` rate-limit modes ([#172](https://github.com/cloudflare/terraform-provider-cloudflare/issues/172))

**Fixes:**
* `cloudflare_page_rule` - Page rule `zone` attribute change to trigger new resource ([#183](https://github.com/cloudflare/terraform-provider-cloudflare/issues/183))

## 1.10.0 (December 18, 2018)

**Improvements:**
* `cloudflare_zone_settings_override` - Add `opportunistic_onion` zone setting support ([#170](https://github.com/cloudflare/terraform-provider-cloudflare/issues/170))
* `cloudflare_zone` - Add ability to set zone plan ([#160](https://github.com/cloudflare/terraform-provider-cloudflare/issues/160))

**Fixes:**
* `cloudflare_zone` - Allow zones to be properly imported ([#157](https://github.com/cloudflare/terraform-provider-cloudflare/issues/157))
* `cloudflare_access_policy` - Match access_policy argument requisites with reality ([#158](https://github.com/cloudflare/terraform-provider-cloudflare/issues/158))
* `cloudflare_filter` - Allow `zone_id` to set `zone` and vice versa ([#162](https://github.com/cloudflare/terraform-provider-cloudflare/issues/162))
* `cloudflare_firewall_rule` - Allow `zone_id` to set `zone` and vice versa ([#174](https://github.com/cloudflare/terraform-provider-cloudflare/issues/174))
* `cloudflare_access_rule` - Ensure `zone` and `zone_id` are always set ([#175](https://github.com/cloudflare/terraform-provider-cloudflare/issues/175))
* Minor documentation fixes

## 1.9.0 (November 15, 2018)

**Improvements:**
* **New Resource:** `cloudflare_access_application` ([#145](https://github.com/cloudflare/terraform-provider-cloudflare/issues/145))
* **New Resource:** `cloudflare_access_policy` ([#145](https://github.com/cloudflare/terraform-provider-cloudflare/issues/145))
* `cloudflare_load_balancer` - Add steering policy support ([#147](https://github.com/cloudflare/terraform-provider-cloudflare/issues/147))
* `cloudflare_load_balancer` - Support `session_affinity` ([#153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/153))
* `cloudflare_load_balancer_pool` - Support `weight` ([#153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/153))

**Fixes:**
* `cloudflare_record` - Compare name without the zone name ([#151](https://github.com/cloudflare/terraform-provider-cloudflare/issues/151))
* Minor documentation fixes ([#149](https://github.com/cloudflare/terraform-provider-cloudflare/issues/149)] [[#152](https://github.com/cloudflare/terraform-provider-cloudflare/issues/152))

## 1.8.0 (November 05, 2018)

**Improvements:**
* **New Resource:** `cloudflare_zone` ([#58](https://github.com/cloudflare/terraform-provider-cloudflare/issues/58))
* **New Resource:** `cloudflare_custom_pages` ([#132](https://github.com/cloudflare/terraform-provider-cloudflare/issues/132))
* `cloudflare_zone_settings_override` - Allow setting SSL level to Strict (SSL-Only Origin Pull) ([#122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/122))
* Update provider usage/build docs and how to update a dependency ([#138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/138))
* Improve `Building The Provider` instructions ([#143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/143))
* `cloudflare_access_rule` - Make importable for all rule types ([#141](https://github.com/cloudflare/terraform-provider-cloudflare/issues/141))
* `cloudflare_load_balancer_pool` - Implement `Update` ([#140](https://github.com/cloudflare/terraform-provider-cloudflare/issues/140))

**Fixes:**
* `cloudflare_rate_limit` - Documentation fixes for markdown where \_ALL\_ is italicized ([#125](https://github.com/cloudflare/terraform-provider-cloudflare/issues/125))
* `cloudflare_worker_route` - Correctly set `multi_script` on Enterprise worker imports ([#124](https://github.com/cloudflare/terraform-provider-cloudflare/issues/124))
* `account_member` - Ignore role ID ordering ([#128](https://github.com/cloudflare/terraform-provider-cloudflare/issues/128))
* `cloudflare_rate_limit` - Origin traffic isn't default anymore ([#130](https://github.com/cloudflare/terraform-provider-cloudflare/issues/130))
* `cloudflare_rate_limit` - Update rate limit validation to allow `1` ([#129](https://github.com/cloudflare/terraform-provider-cloudflare/issues/129))
* `cloudflare_record` - Add validation to ensure TTL is not set while `proxied` is true ([#127](https://github.com/cloudflare/terraform-provider-cloudflare/issues/127))
* Updated code for provider version in User-Agent
* `cloudflare_zone_lockdown` - Fix import of zone lockdowns ([#135](https://github.com/cloudflare/terraform-provider-cloudflare/issues/135))

## 1.7.0 (October 09, 2018)

**Improvements:**
* **New Resource:** `cloudflare_account_member` ([#78](https://github.com/cloudflare/terraform-provider-cloudflare/issues/78))

## 1.6.0 (October 05, 2018)

**Improvements:**
* **New Resource:** `cloudflare_filter`
* **New Resource:** `cloudflare_firewall_rule`

## 1.5.0 (September 21, 2018)

**Improvements:**
* **New Resource:** `cloudflare_zone_lockdown` ([#115](https://github.com/cloudflare/terraform-provider-cloudflare/issues/115))

**Fixes:**
* Send User-Agent header with name and version when contacting API
* `cloudflare_page_rule` - Fix page rule polish (off, lossless or lossy) ([#116](https://github.com/cloudflare/terraform-provider-cloudflare/issues/116))

## 1.4.0 (September 11, 2018)

**Improvements:**
* **New Resource:** `cloudflare_worker_route` ([#110](https://github.com/cloudflare/terraform-provider-cloudflare/issues/110))
* **New Resource:** `cloudflare_worker_script` ([#110](https://github.com/cloudflare/terraform-provider-cloudflare/issues/110))

## 1.3.0 (September 04, 2018)

**Improvements:**
* **New Resource:** `cloudflare_access_rule` ([#64](https://github.com/cloudflare/terraform-provider-cloudflare/issues/64))

**Fixes:**
* `cloudflare_zone_settings_override` -  Change Zone Settings Override to use GetOkExists ([#107](https://github.com/cloudflare/terraform-provider-cloudflare/issues/107))

## 1.2.0 (August 13, 2018)

**Improvements:**
* **New Resource:** `cloudflare_waf_rule` ([#98](https://github.com/cloudflare/terraform-provider-cloudflare/issues/98))
* `cloudflare_zone_settings_override` - Add `off` as Security Level setting ([#99](https://github.com/cloudflare/terraform-provider-cloudflare/issues/99))
* `resource_cloudflare_rate_limit` - Add nat support ([#96](https://github.com/cloudflare/terraform-provider-cloudflare/issues/96))
* `resource_cloudflare_zone_settings_override` - Add `zrt` as a value for the `tls_1_3` setting ([#106](https://github.com/cloudflare/terraform-provider-cloudflare/issues/106))
* Minor documentation improvements

**Fixes:**
* `cloudflare_record` - Setting a DNS record's `proxied` flag to false stopped working ([#103](https://github.com/cloudflare/terraform-provider-cloudflare/issues/103))

## 1.1.0 (July 25, 2018)

FIXES:

* `cloudflare_ip_ranges` - IPv6 CIDR blocks should return IPv6 addresses ([#51](https://github.com/cloudflare/terraform-provider-cloudflare/issues/51))
* `cloudflare_zone_settings_override` - Allow `0` for `browser_cache_ttl` ([#71](https://github.com/cloudflare/terraform-provider-cloudflare/issues/71))
* `cloudflare_page_rule` - `forwarding_urls` in page rules are lists ([#79](https://github.com/cloudflare/terraform-provider-cloudflare/issues/79))
* `cloudflare_page_rule` - The API supports `active` and `disabled`, not `paused` ([#84](https://github.com/cloudflare/terraform-provider-cloudflare/issues/84))

IMPROVEMENTS:
* `cloudflare_zone_settings_override` - Add support for `min_tls_version` ([#72](https://github.com/cloudflare/terraform-provider-cloudflare/issues/72))
* `cloudflare_page_rule` - Add support for more settings: `bypass_cache_on_cookie`, `cache_by_device_type`, `cache_deception_armor`, `cache_on_cookie`, `host_header_override`, `polish`, `explicit_cache_control`, `origin_error_page_pass_thru`, `sort_query_string_for_cache`, `resolve_override`, `respect_strong_etag`, `response_buffering`, `true_client_ip_header`, `mirage`, `disable_railgun`, `cache_key`, `waf`, `rocket_loader`, `cname_flattening` ([#68](https://github.com/cloudflare/terraform-provider-cloudflare/issues/68)], [[#81](https://github.com/cloudflare/terraform-provider-cloudflare/issues/81)], [[#85](https://github.com/cloudflare/terraform-provider-cloudflare/issues/85))
* `cloudflare_page_rule` - Add `off` setting to `security_level` ([#81](https://github.com/cloudflare/terraform-provider-cloudflare/issues/81))
* `cloudflare_record` - DNS Record improvements ([#97](https://github.com/cloudflare/terraform-provider-cloudflare/issues/97))
* Various documentation improvements

## 1.0.0 (April 06, 2018)

BACKWARDS INCOMPATIBILITIES / NOTES:

* resource/cloudflare_record: Changing `name` or `domain` now force a recreation
  of the record ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))

FEATURES:

* **New Resource:** `cloudflare_rate_limit` ([#30](https://github.com/cloudflare/terraform-provider-cloudflare/issues/30))
* **New Resource:** `cloudflare_page_rule` ([#38](https://github.com/cloudflare/terraform-provider-cloudflare/issues/38))
* **New Resource:** `cloudflare_load_balancer` ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))
* **New Resource:** `cloudflare_load_balancer_pool` ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))
* **New Resource:** `cloudflare_zone_settings_override` ([#41](https://github.com/cloudflare/terraform-provider-cloudflare/issues/41))
* **New Resource:** `cloudflare_load_balancer_monitor` ([#42](https://github.com/cloudflare/terraform-provider-cloudflare/issues/42))
* **New Data Source:** `cloudflare_ip_ranges` ([#28](https://github.com/cloudflare/terraform-provider-cloudflare/issues/28))

IMPROVEMENTS:

* resource/cloudflare_record: Validate `TXT` records ([#14](https://github.com/cloudflare/terraform-provider-cloudflare/issues/14))
* resource/cloudflare_record: Add `data` input to suppport SRV, LOC records
  ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))
* resource/cloudflare_record: Add computed attributes `created_on`,
  `modified_on`, `proxiable`, and `metadata` to records ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))
* resource/cloudflare_record: Support import of existing records ([#36](https://github.com/cloudflare/terraform-provider-cloudflare/issues/36))
* New Provider configuration options for API rate limiting ([#43](https://github.com/cloudflare/terraform-provider-cloudflare/issues/43))
* New Provider configuration options for using Organizations ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))

## 0.1.0 (June 20, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider
  Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
## 2.26.0 (August 27th, 2021)

* **New resource**: `cloudflare_notification_policy` ([#1138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1138))
* **New resource**: `cloudflare_notification_policy_webhooks` ([#1151](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1151))
* **New resource**: `cloudflare_ruleset` ([#1143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1143))
* **New resource**: `cloudflare_teams_location` ([#1154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1154))
* **New datasource**: `cloudflare_origin_ca_root_certificate` ([#1158](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1158))

**Improvements**

* `resource/cloudflare_waiting_room`: Add support for `json_response_enabled` as an argument ([#1122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1122))

## 2.25.0 (August 4th, 2021)

**Improvements**

* `resource/cloudflare_access_device_posture_rule`: Add support for `domain_joined`, `firewall`, `os_version`, and `disk_encryption` ([#1137](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1137))
* provider: bump `cloudflare-go` to v0.20.0 ([#1146](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1146))

## 2.24.0 (July 19th, 2021)

**Improvements**

* `resource/cloudflare_logpush_job`: Add support for `"nel_reports"` as a dataset ([#1122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1122))
* `resource/cloudflare_custom_hostname`: Allow SSL options to be optional when not required ([#1131](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1131))
* `resource/cloudflare_access_identity_provider`: Support optional Okta API token ([#1119](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1119))
* `resource/cloudflare_load_balancer_pool`: Add support for load shedding ([#1108](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1108))
* `resource/cloudflare_load_balancer_pool`: Add support for longitude and latitude ([#1093](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1093))

**Fixes**

* `resource/cloudflare_record`: Use correct `Import` method on resource ([#1116](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1116))
* `resource/cloudflare_worker_cron_trigger`: Account for deletion of scripts and force a refresh of triggers ([#1121](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1121))
* `resource/cloudflare_rate_limit`: Handle `origin_traffic` missing from API response ([#1125](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1125))
* `resource/cloudflare_record`: Support `allow_overwrite` for root records ([#1129](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1129))

## 2.23.0 (June 30th, 2021)

- **New resource**: `cloudflare_waiting_room` ([#1053](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1053))

**Improvements**

* `datasource/cloudflare_waf_rules`: Export `default_mode` as an attribute ([#1079](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1079))

**Fixes**

* `resource/cloudflare_access_application`: Revert removal of schema changes causing existing applications unable to re-apply ([#1118](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1118))

## 2.22.0 (June 25th, 2021)

- **New resource**: `cloudflare_static_route` ([#1098](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1098))

**Improvements**

* `resource/cloudflare_origin_ca`: Ignore decreasing `requested_validity` ([#1043](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1078))
* `resource/waf_override`: Allow `rules` to be optional ([#1090](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1090))
* `resource/cloudflare_zone`: Don't attempt to set free zone rate plans as that is already the default ([#1102](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1102))
* `resource/cloudflare_access_application`: Ability to set `type` for Applications ([#1076](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1076))
* `resource/cloudflare_zone_lockdown`: Update documentation to show examples of multiple configurations ([#1106](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1106))

## 2.21.0 (May 26th, 2021)

- **New resource**: `cloudflare_device_posture_rule` ([#1058](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1058))
- **New resource**: `cloudflare_teams_list` ([#1058](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1058))

**Improvements**

* provider: Update to terraform-plugin-sdk v1.17.1 ([#1035](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1035), [#1043](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1043))
* `resource/cloudflare_logpush_job`: Allow `ownership_challenge` to be optional to account for Datadog, Splunk or S3-Compatible endpoints ([#1048](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1048))
* `resource/cloudflare_access_group`: Add support for `login_method` ([#1066](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1066))
* `resource/cloudflare_load_balancer`: Add support for `promixity` based steering ([#1072](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1072))
* `resource/cloudflare_access_application`: Prevent bad CORS configuration when credentials and all origins are permitted ([#1073](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1073))
* `resource/cloudflare_access_service_tokens`: Allow configuration to manage automatic renewal when the threshold is crossed and Terraform operations are performed within the window ([#1057](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1057))
* `resource/cloudflare_load_balancer_pool`: Allow support for `Host` header settings ([#1042](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1042))

**Fixes**

* `resource/cloudflare_access_policy`: Allow empty slices in blocks when building policies ([#1034](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1034))
* `resource/cloudflare_load_balancer`: Fix `override` attributes `pop_pools` and `region_pools` referencing incorrect values causing a panic ([#1039](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1039))

## 2.20.0 (April 15th, 2021)

**New resource**: `cloudflare_access_ca_certificate` ([#995](https://github.com/cloudflare/terraform-provider-cloudflare/issues/995))

**Improvements**

* `resource/cloudflare_access_application`: Improve documentation for `Import` usage ([#1002](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1002))
* `resource/cloudflare_logpush_job`: Update documentation to reflect requirements for `destination_conf` to match across all uses ([#1024](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1024))
* `resource/cloudflare_custom_hostname_fallback`: Better handle service lag when updating existing resources by attempting retries ([#1014](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1014))
* `resource/cloudflare_waf_group`: Simplify error handling using inbuilt helpers ([#1015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1015))
* `resource/cloudflare_waf_rule`: Simplify error handling using inbuilt helpers ([#1015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1015))
* `resource/cloudflare_waf_package`: Simplify error handling using inbuilt helpers ([#1015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1015))
* `resource/cloudflare_access_group`: Add support for `login_method` ([#1018](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1018))
* provider: Update to cloudflare-go v0.16.0 ([#1018](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1018))
* provider: Update to terraform-plugin-sdk v1.16.1 ([#1003](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1003))
* `resource/cloudflare_load_balancer`: Add support for `rules` ([#1016](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1016))

## 2.19.2 (March 15th, 2021)

**Fixes**

* `resource/cloudflare_record`: Address regression from 2.19.1 by checking the API response instead of the schema output for `Priority` ([#992](https://github.com/cloudflare/terraform-provider-cloudflare/issues/992))

## 2.19.1 (March 11th, 2021)

**Fixes**

* `resource/cloudflare_record`: Update `Priority` handling for MX parked records ([#986](https://github.com/cloudflare/terraform-provider-cloudflare/issues/986))

## 2.19.0 (March 10th, 2021)

**Fixes**

* `resource/cloudflare_access_group`: Fix crash when constructing a GSuite group ([#940](https://github.com/cloudflare/terraform-provider-cloudflare/issues/940))
* `resource/cloudflare_access_policy`: Make `precedence` required ([#941](https://github.com/cloudflare/terraform-provider-cloudflare/issues/941))
* `resource/cloudflare_access_group`: Fix crash when constructing a SAML group ([#948](https://github.com/cloudflare/terraform-provider-cloudflare/issues/948))
* `resource/cloudflare_zone`: Update `Retry` logic to look at an available field for passing conditions ([#973](https://github.com/cloudflare/terraform-provider-cloudflare/issues/973))
* `resource/cloudflare_page_rule`: Allow ignoring/including all query string parameters for `cache_key_fields` ([#975](https://github.com/cloudflare/terraform-provider-cloudflare/issues/975))

**Improvements**

* `resource/cloudflare_access_policy`: Enable zone and account level resources to be imported ([#956](https://github.com/cloudflare/terraform-provider-cloudflare/issues/956))
* `resource/cloudflare_origin_ca_certificate`: Smoother import process with less recreation ([#955](https://github.com/cloudflare/terraform-provider-cloudflare/issues/955))
* provider: Update internals to match `cloudflare-go` 0.14 for better error handling and context aware methods ([#976](https://github.com/cloudflare/terraform-provider-cloudflare/issues/976))

## 2.18.0 (February 3rd, 2021)

* **New Resource:** `cloudflare_argo_tunnel` ([#905](https://github.com/cloudflare/terraform-provider-cloudflare/issues/905))
* **New Resource:** `cloudflare_worker_cron_trigger` ([#926](https://github.com/cloudflare/terraform-provider-cloudflare/issues/926))

**Fixes**

* `datasource/cloudflare_zones`: Pagination is now correctly handled internally and will return more than the single page of results ([cloudflare/cloudflare-go#534](https://github.com/cloudflare/cloudflare-go/pull/534)).
* `resource/cloudflare_access_policy`: Correctly handle transforming API responses to schema ([#917](https://github.com/cloudflare/terraform-provider-cloudflare/issues/917))
* `resource/cloudflare_access_group`: Correctly handle transforming API responses to schema ([#918](https://github.com/cloudflare/terraform-provider-cloudflare/issues/918))
* `resource/cloudflare_ip_list`: Ensure account ID is persisted during `Import` ([#916](https://github.com/cloudflare/terraform-provider-cloudflare/issues/916))

**Improvements**

* `resource/cloudflare_access_application`: Allow any `session_duration` that is `time.ParseDuration` compatible ([#910](https://github.com/cloudflare/terraform-provider-cloudflare/issues/910))
* `resource/cloudflare_rate_limit`: Add the ability to configure `match.response.headers` in rate limits ([#911](https://github.com/cloudflare/terraform-provider-cloudflare/issues/911))
* `resource/cloudflare_access_rule`: Validate IP masks within schema ([#921](https://github.com/cloudflare/terraform-provider-cloudflare/issues/921))

## 2.17.0 (January 5th, 2021)

* **New Resource:** `cloudflare_magic_firewall_ruleset` ([#884](https://github.com/cloudflare/terraform-provider-cloudflare/issues/884))

**Fixes**

* `resource/cloudfare_api_token`: Omitting `conditions` will no longer send empty arrays causing IP restriction issues and unusable tokens ([#902](https://github.com/cloudflare/terraform-provider-cloudflare/pull/902))

## 2.16.0 (January 5th, 2021)

**Improvements**

* `resource/cloudflare_access_application`: Add support for `custom_deny_message` and `custom_deny_url` values ([#895](https://github.com/cloudflare/terraform-provider-cloudflare/issues/895))
* `resource/cloudflare_load_balancer_monitor`: Add support for `probe_zone` for monitors ([#903](https://github.com/cloudflare/terraform-provider-cloudflare/issues/903))

## 2.15.0 (December 29th, 2020)

**Improvements**

* `resource/cloudflare_load_balancer`: Add support for `session_affinity_ttl` ([#882](https://github.com/cloudflare/terraform-provider-cloudflare/issues/882))
* `resource/cloudflare_load_balancer`: Add support for `session_affinity_attributes` ([#883](https://github.com/cloudflare/terraform-provider-cloudflare/issues/883))

**Fixes**

* `resource/cloudflare_page_rule`: Fixed crash during update when using custom cache key ([#894](https://github.com/cloudflare/terraform-provider-cloudflare/pull/894))

## 2.14.0 (November 26th, 2020)

* **New Resource:** `cloudflare_api_token` ([#862](https://github.com/cloudflare/terraform-provider-cloudflare/issues/862))
* **New Datasource:** `cloudflare_api_token_permission_groups` ([#862](https://github.com/cloudflare/terraform-provider-cloudflare/issues/862))
* **New Resource:** `cloudflare_zone_dnssec` ([#852](https://github.com/cloudflare/terraform-provider-cloudflare/issues/852))
* **New Datasource:** `cloudflare_zone_dnssec` ([#852](https://github.com/cloudflare/terraform-provider-cloudflare/issues/852))

**Improvements**

* `resource/cloudflare_record`: Add explicit fields for CAA records instead of relying on the map value ([#866](https://github.com/cloudflare/terraform-provider-cloudflare/issues/866))
* `resource/cloudflare_account_member`: Swap schema `role_ids` to `TypeSet` to better handle internal ordering changes ([#876](https://github.com/cloudflare/terraform-provider-cloudflare/issues/876))

**Fixes**

* `datasource/cloudflare_waf_groups`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))
* `datasource/cloudflare_waf_packages`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))
* `datasource/cloudflare_waf_rules`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))
* `datasource/cloudflare_zones`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))

## 2.13.2 (November 6th, 2020)

**Fixes**

* `resource/cloudflare_filter`: Remove schema based validation for filters ([#863](https://github.com/cloudflare/terraform-provider-cloudflare/issues/863))

## 2.13.1 (November 5th, 2020)

**Improvements**

* `resource/cloudflare_filter`: Pass missing credential error through to end user ([#860](https://github.com/cloudflare/terraform-provider-cloudflare/issues/860))

## 2.13.0 (November 5th, 2020)

**Improvements**

* `datasource/cloudflare_ip_ranges`: Add the ability to query `china_ipv4_cidr_blocks` and `china_ipv6_cidr_blocks` ([#833](https://github.com/cloudflare/terraform-provider-cloudflare/issues/833))
* `resource/cloudflare_filter`: Improve validation of expressions using the schema ([#848](https://github.com/cloudflare/terraform-provider-cloudflare/issues/848))

**Fixes**

* `resource/cloudflare_page_rule`: Set default for `cache_key_fields.host.resolved` to prevent panics ([#832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/832))
* `resource/cloudflare_authenticated_origin_pulls`: Fix off-by-one error check in `Import` ([#832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/859))
* `resource/cloudflare_authenticated_origin_pulls_certificate`: Fix off-by-one error check in `Import` ([#832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/859))

## 2.12.0 (October 22nd, 2020)

**Improvements**

* `resource/cloudflare_certificate_pack`: Swap internal representation of `hosts` to remove inconsistent ordering issues ([#800](https://github.com/cloudflare/terraform-provider-cloudflare/issues/800))
* `resource/cloudflare_logpush_job`: Handle deletion outside of Terraform ([#798](https://github.com/cloudflare/terraform-provider-cloudflare/issues/798))
* `resource/cloudflare_access_group`: Add support for `geo` conditionals ([#803](https://github.com/cloudflare/terraform-provider-cloudflare/issues/803))
* `resource/cloudflare_access_application`: Add support for `enable_binding_cookie` ([#802](https://github.com/cloudflare/terraform-provider-cloudflare/issues/802))
* `resource/cloudflare_waf_rule`: Improve documentation for `mode` ([#824](https://github.com/cloudflare/terraform-provider-cloudflare/issues/824))
* `datasource/cloudflare_waf_rule`: Improve documentation for `mode` ([#824](https://github.com/cloudflare/terraform-provider-cloudflare/issues/824))
* `resource/cloudflare_access_application`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))
* `resource/cloudflare_access_group`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))
* `resource/cloudflare_access_identity_provider`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))
* `resource/cloudflare_access_policy`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))

**Fixes**

* `resource/cloudflare_custom_hostname_fallback_origin`: Don't retry the "active" status of custom hostnames fallbacks ([#818](https://github.com/cloudflare/terraform-provider-cloudflare/issues/818))
* `resource/cloudflare_zone`: Remove `DiffSuppressFunc` causing `jump_start` issues ([#830](https://github.com/cloudflare/terraform-provider-cloudflare/issues/830))

## 2.11.0 (September 11th, 2020)

* **New Resource:** `cloudflare_certificate_pack` ([#778](https://github.com/cloudflare/terraform-provider-cloudflare/issues/778))

**Improvements**

* `resource/cloudflare_access_group`: Add support for `auth_method` ([#762](https://github.com/cloudflare/terraform-provider-cloudflare/issues/762))
* `resource/cloudflare_access_group`: De-duplicate blocks in groups by accepting lists instead ([#739](https://github.com/cloudflare/terraform-provider-cloudflare/issues/739))
* `resource/cloudflare_worker_script`: Adds support for `webassembly_binding` ([#780](https://github.com/cloudflare/terraform-provider-cloudflare/issues/780))
* `resource/cloudflare_healthcheck`: Retry hostname resolution errors when encountering "no such host" responses ([#789](https://github.com/cloudflare/terraform-provider-cloudflare/issues/789))
* `resource/cloudflare_access_application`: Better validation for allowed methods and origin combinations to prevent getting state into an unrecoverable state ([#793](https://github.com/cloudflare/terraform-provider-cloudflare/issues/793))

**Fixes**

* `resource/cloudflare_healthcheck`: Handle resource deletion outside of Terraform ([#787](https://github.com/cloudflare/terraform-provider-cloudflare/issues/787))
* `resource/cloudflare_custom_hostname`: Ensure `Import` sets hostname to prevent recreation ([#788](https://github.com/cloudflare/terraform-provider-cloudflare/issues/788))
* `resource/cloudflare_ip_list`: Handle resource deletion outside of Terraform ([#794](https://github.com/cloudflare/terraform-provider-cloudflare/issues/794))
* `resource/cloudflare_ip_list`: Remove `item`.`id` from schema ([#796](https://github.com/cloudflare/terraform-provider-cloudflare/issues/796))

## 2.10.1 (August 24th, 2020)

**Fixes**

* `resource/cloudflare_access_application`: Handle the `zone_id` => `account_id` move internally ([#724](https://github.com/cloudflare/terraform-provider-cloudflare/issues/724))

## 2.10.0 (August 24th, 2020)

* **New Resource:** `cloudflare_custom_hostname_origin_fallback` ([#757](https://github.com/cloudflare/terraform-provider-cloudflare/issues/757))
* **New Resource:** `cloudflare_authenticated_origin_pulls` ([#749](https://github.com/cloudflare/terraform-provider-cloudflare/issues/749))
* **New Resource:** `cloudflare_authenticated_origin_pulls_certificate` ([#749](https://github.com/cloudflare/terraform-provider-cloudflare/issues/749))
* **New Resource:** `cloudflare_ip_list` ([#766](https://github.com/cloudflare/terraform-provider-cloudflare/issues/766))

**Improvements**

* `resource/cloudflare_spectrum_application`: Add support for port ranges ([#745](https://github.com/cloudflare/terraform-provider-cloudflare/issues/745))
* `resource/cloudflare_custom_hostname`: Force creation of a new resource if the `zone_id` value changes ([#761](https://github.com/cloudflare/terraform-provider-cloudflare/issues/761))
* `resource/cloudflare_record`: Retry record creation/update if the response includes an "already exists" exception for handling race conditions ([#773](https://github.com/cloudflare/terraform-provider-cloudflare/issues/773))

**Fixes**

* `resource/cloudflare_firewall_rule`: Compare descriptions after converting unicode + HTML entities to prevent unnecessary diffs ([#758](https://github.com/cloudflare/terraform-provider-cloudflare/issues/758))
* `resource/cloudflare_filter`: Compare descriptions after converting unicode + HTML entities to prevent unnecessary diffs ([#758](https://github.com/cloudflare/terraform-provider-cloudflare/issues/758))

## 2.9.0 (July 30th, 2020)

* **New Resource:** `cloudflare_custom_hostname` (SSL for SaaS) ([#746](https://github.com/cloudflare/terraform-provider-cloudflare/issues/746))

**Improvements**

* `resource/access_application`: Add support for `allowed_idps` and restricting which Identity Providers are associated with an Application ([#734](https://github.com/cloudflare/terraform-provider-cloudflare/issues/734))
* `resource/access_application`: Add support for `auto_redirect_to_identity` ([#730](https://github.com/cloudflare/terraform-provider-cloudflare/issues/730))
* `resource/access_application`: Add CORS support ([#725](https://github.com/cloudflare/terraform-provider-cloudflare/issues/725))
* `resource/cloudflare_custom_ssl`: Allow `geo_restrictions` to be `nil` and not included in the request payload ([#714](https://github.com/cloudflare/terraform-provider-cloudflare/issues/714))
* `datasource/cloudflare_zones`: Filtering is now performed on the server side and the `name` parameter is no longer a regex. Instead, `name` is a string to match on and `match` is a regex. See the website documentation for more examples and updated references ([#708](https://github.com/cloudflare/terraform-provider-cloudflare/issues/708)) in order to make your code compatible with this release.


## 2.8.0 (June 22, 2020)

* **New Resource:** `cloudflare_waf_override` ([#691](https://github.com/cloudflare/terraform-provider-cloudflare/issues/691))

**Improvements**

* `resource/cloudflare_argo`: Allow `tiered_caching` and `smart_routing` to be toggled individually allowing for entitlement differences ([#703](https://github.com/cloudflare/terraform-provider-cloudflare/issues/703))
* `resource/cloudflare_page_rule`: Add support for `cache_ttl_by_status` ([#706](https://github.com/cloudflare/terraform-provider-cloudflare/issues/706))
* `resource/cloudflare_worker_script`: Add support for `plain_text` and `secret_text` bindings ([#710](https://github.com/cloudflare/terraform-provider-cloudflare/issues/710))

**Fixes**

* `resource/cloudflare_record`: Update `TestAccCloudflareRecord_LOC` test asserted value to use less precise floats and match the API responses ([#712](https://github.com/cloudflare/terraform-provider-cloudflare/issues/712))
* `resource/cloudflare_record`: Update `TestAccCloudflareRecord_Basic` test `metadata` attributes to match updated API payload ([#713](https://github.com/cloudflare/terraform-provider-cloudflare/issues/713))

## 2.7.0 (May 20, 2020)

* **New Resource:** `cloudflare_byo_ip_prefix` ([#671](https://github.com/cloudflare/terraform-provider-cloudflare/issues/671))
* **New Resource:** `cloudflare_logpull_retention` ([#678](https://github.com/cloudflare/terraform-provider-cloudflare/issues/678))
* **New Resource:** `cloudflare_healthcheck` ([#680](https://github.com/cloudflare/terraform-provider-cloudflare/issues/680))

**Improvements:**

* `resource/cloudflare_worker_route`: Improve documentation to mention using `account_id` for the underlying APIs ([#669](https://github.com/cloudflare/terraform-provider-cloudflare/issues/669))
* `resource/cloudflare_worker_script`: Improve documentation to mention using `account_id` for the underlying APIs ([#670](https://github.com/cloudflare/terraform-provider-cloudflare/issues/670))
* `resource/cloudflare_load_balancer_pool`: Improve documentation to mention `notification_email` accepts a comma delimited list of emails ([#687](https://github.com/cloudflare/terraform-provider-cloudflare/issues/687))
* `resource/cloudflare_page_rule`: Add support for `cache_key_fields` Page Rule action ([#662](https://github.com/cloudflare/terraform-provider-cloudflare/issues/662))

**Fixes:**
* `resource/cloudflare_zone_settings_override`: Fix regression where if you didn't have universal SSL settings defined, it would error when setting them ([#663](https://github.com/cloudflare/terraform-provider-cloudflare/issues/663))
* `resource/cloudflare_zone`: Handle changing zone rate plan from "free" to "enterprise" ([#668](https://github.com/cloudflare/terraform-provider-cloudflare/issues/668))
* `resource/cloudflare_record`: Update validation to allow PTR records ([9a8fd43](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9a8fd43))

## 2.6.0 (April 22, 2020)

**Improvements:**

* `resource/cloudflare_zone_settings_override`: Add `universal_ssl` to control enablement of Universal SSL on a zone ([#658](https://github.com/cloudflare/terraform-provider-cloudflare/issues/658))
* provider: API keys and API tokens are now validated to help differentiate incorrect usage before making API calls ([#661](https://github.com/cloudflare/terraform-provider-cloudflare/issues/661))
* `resource/cloudflare_logpush_job`: Add support for "firewall_events" dataset parameter ([#660](https://github.com/cloudflare/terraform-provider-cloudflare/issues/660))
* `resource/cloudflare_logpush_job`: Add support for "dataset" parameter ([#649](https://github.com/cloudflare/terraform-provider-cloudflare/issues/649))
* `resource/cloudflare_zone_settings_override`: Remove `edge_cache_ttl` ([#654](https://github.com/cloudflare/terraform-provider-cloudflare/issues/654))
* `resource/cloudflare_access_group`: Allow Access conditions for `include`/`require`/`exclude` to be used consistently between Access Groups and Access Policies ([#646](https://github.com/cloudflare/terraform-provider-cloudflare/issues/646))

**Fixes:**
* `resource/cloudflare_logpush_job`: fix for `strconv.Atoi: parsing ""` error while creating Logpush job

## 2.5.1 (April 03, 2020)

**Improvements:**

* `resource/cloudflare_zone_settings_override`: Update `image_resizing` options to include `"open"` ([#639](https://github.com/cloudflare/terraform-provider-cloudflare/issues/639))

**Fixes:**

* `resource/cloudflare_access_group`: Fixed misspelt Okta in JSON payload ([cloudflare/cloudflare-go#440](https://github.com/cloudflare/cloudflare-go/issues/440))

## 2.5.0 (March 27, 2020)

**Improvements:**

* `resource/cloudflare_access_policy`: Add support for `service_token` and `any_valid_service_token` ([#612](https://github.com/cloudflare/terraform-provider-cloudflare/issues/612))
* `resource/cloudflare_waf_group`: Handle WAF group deletions in the API responses ([#623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/623))
* `resource/cloudflare_waf_package`: Handle WAF package deletions in the API responses ([#623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/623))
* `resource/cloudflare_waf_rule`: Handle WAF rule deletions in the API responses ([#623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/623))
* `resource/cloudflare_access_policy`: Add support for `group` ([#626](https://github.com/cloudflare/terraform-provider-cloudflare/issues/626))
* `resource/cloudflare_firewall_rule`: Add support for bypassing specific `products` ([#630](https://github.com/cloudflare/terraform-provider-cloudflare/issues/630))
* `resource/cloudflare_spectrum_application`: Add support for `edge_ips`, `argo_smart_routing` and `edge_ip_connectivity` ([#631](https://github.com/cloudflare/terraform-provider-cloudflare/issues/631))
* `resource/cloudflare_access_group`: Add support for using external providers (`gsuite`, `github`, `azure`, `okta`, `saml`, `mTLS certificate`, `common name`
) ([#633](https://github.com/cloudflare/terraform-provider-cloudflare/issues/633))

## 2.4.1 (March 12, 2020)

**Improvements:**

* `resource/cloudflare_logpush_job`: Support `Import` on the resource ([#618](https://github.com/cloudflare/terraform-provider-cloudflare/issues/618))

**Fixes:**

* `resource/cloudflare_record`: Missing CAA in DNS validation ([#619](https://github.com/cloudflare/terraform-provider-cloudflare/issues/619))

## 2.4.0 (March 09, 2020)

* **New Resource:** `cloudflare_workers_kv` ([#595](https://github.com/cloudflare/terraform-provider-cloudflare/issues/595))
* **New Resource:** `cloudflare_access_identity_provider` ([#597](https://github.com/cloudflare/terraform-provider-cloudflare/issues/597))


**Improvements:**

* `resource/cloudflare_record`: Stricter validation for record types ([#610](https://github.com/cloudflare/terraform-provider-cloudflare/issues/610))
* `resource/logpush_job`: Add more verbose error handling ([#564](https://github.com/cloudflare/terraform-provider-cloudflare/issues/564))
* `resource/zone_settings_override`: Update documentation for `cache_level` values ([#606](https://github.com/cloudflare/terraform-provider-cloudflare/issues/606))
* `resource/access_application`: Add documentation for available attributes ([#587](https://github.com/cloudflare/terraform-provider-cloudflare/issues/587))
* `resource/cloudflare_firewall_rule`: Add support for bypassing security configuration rules by URL ([#568](https://github.com/cloudflare/terraform-provider-cloudflare/issues/568))
* `resource/cloudflare_record_migrate`: Use `zone_id` for state migration before attempting to use `domain` ([#566](https://github.com/cloudflare/terraform-provider-cloudflare/issues/566))
* `resource/cloudflare_load_balancer`: Update `session_affinity` validation to allow `"ip_cookie"` ([#573](https://github.com/cloudflare/terraform-provider-cloudflare/issues/573))
* `datasource/ip_ranges`: Update documentation to show 0.12 syntax ([#617](https://github.com/cloudflare/terraform-provider-cloudflare/issues/617))

**Fixes**

* `resource/zone_settings_override`: Handle individual zone settings within `Delete` operations ([#599](https://github.com/cloudflare/terraform-provider-cloudflare/issues/599))

## 2.3.0 (December 18, 2019)

* **New Resource:** `cloudflare_origin_ca_certificate` ([#547](https://github.com/cloudflare/terraform-provider-cloudflare/issues/547))

**Fixes:**

* `resource/cloudflare_zone_settings_override`: Renamed `0rtt` to `zero_rtt` to conform to HCL grammar requirements ([#557](https://github.com/cloudflare/terraform-provider-cloudflare/issues/557))

**Improvements:**

* `resource/cloudflare_access_rule`: Add `ip6` as valid option ([#560](https://github.com/cloudflare/terraform-provider-cloudflare/issues/560))
* `resource/cloudflare_spectrum_application`: Swap `proxy_protocol` to string field with supporting enum values instead ([#561](https://github.com/cloudflare/terraform-provider-cloudflare/issues/561))
* `resource/cloudflare_waf_rule`: Add `package_id` as valid option and export `group_id` ([#552](https://github.com/cloudflare/terraform-provider-cloudflare/issues/552))

## 2.2.0 (December 05, 2019)

* **New Resource:** `cloudflare_access_group` ([#510](https://github.com/cloudflare/terraform-provider-cloudflare/issues/510))
* **New Resource:** `cloudflare_workers_kv_namespace` ([#443](https://github.com/cloudflare/terraform-provider-cloudflare/issues/443))

**Improvements:**

* `resource/cloudflare_zone_settings_override`: Add `non_identity` to allowed `decision` schema ([#541](https://github.com/cloudflare/terraform-provider-cloudflare/issues/541))
* `resource/cloudflare_zone_settings_override`: Add support for `0rtt` and `http3` settings ([#542](https://github.com/cloudflare/terraform-provider-cloudflare/issues/542))
* `resource/cloudflare_load_balancer_monitor`: Allow empty string for `expected_body` ([#539](https://github.com/cloudflare/terraform-provider-cloudflare/issues/539))
* `resource/cloudflare_worker_script`: Add support for Worker KV Namespace Bindings ([#544](https://github.com/cloudflare/terraform-provider-cloudflare/issues/544))
* `data_source/waf_rules`, `resource/cloudflare_waf_rule`, Support allowed modes for WAF Rules ([#550](https://github.com/cloudflare/terraform-provider-cloudflare/issues/550))

**Fixes:**

* `resource/cloudflare_spectrum_application`: Spectrum origin_port is optional ([#549](https://github.com/cloudflare/terraform-provider-cloudflare/issues/549))

## 2.1.0 (November 07, 2019)

* **New datasource:** `cloudflare_waf_rules` ([#525](https://github.com/cloudflare/terraform-provider-cloudflare/issues/525))

**Improvements:**

* `resource/cloudflare_zone`: Expose `verification_key` for partial setups ([#532](https://github.com/cloudflare/terraform-provider-cloudflare/issues/532))
* `resource/cloudflare_worker_route`: Enable API Tokens support from upstream [cloudflare-go](https://github.com/cloudflare/cloudflare-go) release

## 2.0.1 (October 22, 2019)

* **New Resource:** `cloudflare_access_service_tokens` ([#521](https://github.com/cloudflare/terraform-provider-cloudflare/issues/521))
* **New Resource:** `cloudflare_waf_package` ([#475](https://github.com/cloudflare/terraform-provider-cloudflare/issues/475))
* **New Resource:** `cloudflare_waf_group` ([#476](https://github.com/cloudflare/terraform-provider-cloudflare/issues/476))
* **New datasource:** `cloudflare_waf_groups` ([#508](https://github.com/cloudflare/terraform-provider-cloudflare/issues/508))
* **New datasource:** `cloudflare_waf_packages` ([#509](https://github.com/cloudflare/terraform-provider-cloudflare/issues/509))

**Fixes:**

* `resource/cloudflare_page_rule`: Set `h2_prioritization` individually not via bulk endpoint ([#493](https://github.com/cloudflare/terraform-provider-cloudflare/issues/493))
* `resource/cloudflare_zone_settings_override`: Set `zone_id` to prevent unnecessary re-creation of resources ([#502](https://github.com/cloudflare/terraform-provider-cloudflare/issues/502))

**Improvements:**

* `resource/cloudflare_spectrum_application`: Add support for setting `traffic_type` ([#481](https://github.com/cloudflare/terraform-provider-cloudflare/issues/481))
* `resource/cloudflare_zone_settings_override`: Update documentation with default values ([#498](https://github.com/cloudflare/terraform-provider-cloudflare/issues/498))

**Internals:**

* Migrated to Terraform plugin SDK ([#489](https://github.com/cloudflare/terraform-provider-cloudflare/issues/489))

## 2.0.0 (September 30, 2019)

**Breaking changes:**
* `provider/cloudflare`:
 * renamed `token` to `api_key`
 * renamed `org_id` to `account_id`
 * removed `use_org_from_zone`, you need to explicitly specify `account_id`
* Environment variables:
 * renamed `CLOUDFLARE_TOKEN` to `CLOUDFLARE_API_TOKEN`
 * renamed `CLOUDFLARE_ORG_ID` to `CLOUDFLARE_ACCOUNT_ID`
 * removed `CLOUDFLARE_ORG_ZONE`, you need to explicitly specify `CLOUDFLARE_ACCOUNT_ID`
* Changed the following resources to require Zone ID:
 * `cloudflare_access_rule`
 * `cloudflare_filter`
 * `cloudflare_firewall_rule`
 * `cloudflare_load_balancer`
 * `cloudflare_page_rule`
 * `cloudflare_rate_limit`
 * `cloudflare_record`
 * `cloudflare_waf_rule`
 * `cloudflare_worker_route"`
 * `cloudflare_zone_lockdown`
 * `cloudflare_zone_settings_override`
* Workers single-script support removed

Please see [Version 2 Upgrade Guide](https://www.terraform.io/docs/providers/cloudflare/guides/version-2-upgrade.html) for details.

**Improvements:**

* `cloudflare/resource_cloudflare_argo`: Handle errors when fetching tiered caching + smart routing settings ([#477](https://github.com/cloudflare/terraform-provider-cloudflare/issues/477))
* Various documentation updates for 0.12 syntax

## 1.18.1 (August 29, 2019)

**Fixes:**

* `resource/cloudflare_load_balancer`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_page_rule`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_rate_limit`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_waf_rule`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_worker_route`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_worker_script`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_zone_lockdown`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))

## 1.18.0 (August 27, 2019)

**Fixes:**

* `resource/cloudflare_page_rule`: Fix a logic condition where setting `edge_cache_ttl` action but then not updating it in subsequent `apply` runs causes it to be blown away ([#453](https://github.com/cloudflare/terraform-provider-cloudflare/issues/453))

**Improvements:**

* provider: You can now use API tokens to authenticate instead of user email and key ([#450](https://github.com/cloudflare/terraform-provider-cloudflare/issues/450))
* `resource/cloudflare_zone_lockdown`: `priority` can now be set on the resource ([#445](https://github.com/cloudflare/terraform-provider-cloudflare/issues/445))
* `resource/cloudflare_custom_ssl`: Updated website documentation navigation to include link for resource ([#442)](https://github.com/cloudflare/terraform-provider-cloudflare/issues/442))

**Deprecations:**

* `resource/cloudflare_access_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_filter`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_firewall_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_load_balancer`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_page_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_rate_limit`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_waf_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_worker_route`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_worker_script`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_zone_lockdown`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))

## 1.17.1 (August 09, 2019)

**Fixes:**

* Partially revert [[#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421)] deprecation messages

## 1.17.0 (August 09, 2019)

**Removals:**

* `resource/cloudflare_zone_settings_override`: `sha1_support` has been removed due to Cloudflare no longer supporting SHA1 certificates or the API endpoint ([#415](https://github.com/cloudflare/terraform-provider-cloudflare/issues/415))

**Deprecations:**

* `resource/cloudflare_zone_settings_override`: `tls_1_2_only` has been superseded by using `min_tls_version` instead ([#405](https://github.com/cloudflare/terraform-provider-cloudflare/issues/405))
* `resource/cloudflare_access_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_filter`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_firewall_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_load_balancer`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_page_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_rate_limit`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_waf_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_worker_route`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_worker_script`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_zone_lockdown`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))

**Improvements:**

* **New Resource:** `cloudflare_custom_ssl` ([#418](https://github.com/cloudflare/terraform-provider-cloudflare/issues/418))
* `resource/cloudflare_filter`: Strip all surrounding whitespace from filter expressions to match API responses ([#361](https://github.com/cloudflare/terraform-provider-cloudflare/issues/361))
* `resource/cloudflare_zone`: Support unicode zone name values ([#412](https://github.com/cloudflare/terraform-provider-cloudflare/issues/412))
* `resource/cloudflare_page_rule`: Allow setting `origin_pull` for SSL ([#430](https://github.com/cloudflare/terraform-provider-cloudflare/issues/430))
* `resource/cloudflare_load_balancer_monitor`: Add TCP support for load balancer monitor ([#428](https://github.com/cloudflare/terraform-provider-cloudflare/issues/428))

**Fixes:**
* `resource/cloudflare_logpush_job`: Update documentation ([#395](https://github.com/cloudflare/terraform-provider-cloudflare/issues/395))
* `resource/cloudflare_zone_lockdown`: Fix: examples in documentation ([#407](https://github.com/cloudflare/terraform-provider-cloudflare/issues/407))
* `resource/cloudflare_page_rule`: Set nil on changed string-based Page Rule actions

## 1.16.1 (June 27, 2019)

**Fixes:**

* `resource/cloudflare_page_rule`: Fix regression in `browser_cache_ttl` where the value was sent as a string instead of an integer to the remote ([#390](https://github.com/cloudflare/terraform-provider-cloudflare/issues/390))

## 1.16.0 (June 20, 2019)

**Improvements:**

* `resource/cloudflare_zone_settings_override`: Add support for `h2_prioritization` and `image_resizing` ([#381](https://github.com/cloudflare/terraform-provider-cloudflare/issues/381))
* `resource/cloudflare_load_balancer_pool`: Update IP range for tests to not use reserved ranges ([#369](https://github.com/cloudflare/terraform-provider-cloudflare/issues/369))

**Fixes:**

* `resource/cloudflare_page_rule`: Fix issues with `browser_cache_ttl` defaults and when value is `0` (for Enterprise users)   ([#379](https://github.com/cloudflare/terraform-provider-cloudflare/issues/379))

## 1.15.0 (May 24, 2019)

* The provider is now compatible with Terraform v0.12, while retaining compatibility with prior versions. ([#309](https://github.com/cloudflare/terraform-provider-cloudflare/issues/309))

## 1.14.0 (May 15, 2019)

**Improvements:**

* **New Resource:** `cloudflare_argo` Manage Argo features ([#304](https://github.com/cloudflare/terraform-provider-cloudflare/issues/304))
* `cloudflare_zone`: Support management of partial zones ([#303](https://github.com/cloudflare/terraform-provider-cloudflare/issues/303))
* `cloudflare_rate_limit`: Update `modes` documentation ([#293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/212))
* `cloudflare_load_balancer`: Allow steering policy of "random" ([#329](https://github.com/cloudflare/terraform-provider-cloudflare/issues/329))

**Fixes:**

* `cloudflare_page_rule` - Allow setting `browser_cache_ttl` to 0 ([#293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/291))
* `cloudflare_page_rule` - Swap to completely replacing rules ([#338](https://github.com/cloudflare/terraform-provider-cloudflare/issues/338))

## 1.13.0 (April 12, 2019)

**Improvements**

* **New Resource:** `cloudflare_logpush_job` ([#287](https://github.com/cloudflare/terraform-provider-cloudflare/issues/287))
* `cloudflare_zone_settings` - Remove option to toggle `always_on_ddos` ([#253](https://github.com/cloudflare/terraform-provider-cloudflare/issues/253))
* `cloudflare_page_rule` - Update documentation to clarify "0" usage
* `cloudflare_zones` - Return zone ID and zone name ([#275](https://github.com/cloudflare/terraform-provider-cloudflare/issues/275))
* `cloudflare_load_balancer` - Add `enabled` field ([#208](https://github.com/cloudflare/terraform-provider-cloudflare/issues/208))
* `cloudflare_record` - validators: Allow PTR DNS records ([#283](https://github.com/cloudflare/terraform-provider-cloudflare/issues/283))

**Fixes:**

* `cloudflare_custom_pages` - Use correct casing for `zone_id` lookups
* `cloudflare_rate_limit` - Make `correlate` optional and not flap in state management ([#271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/271))
* `cloudflare_spectrum_application` - Fixed integration tests to work ([#275](https://github.com/cloudflare/terraform-provider-cloudflare/issues/275))
* `cloudflare_page_rule` - Better track field changes in `actions` resource. ([#107](https://github.com/cloudflare/terraform-provider-cloudflare/issues/107))

## 1.12.0 (March 07, 2019)

**Improvements:**

* provider: Enable request/response logging ([#212](https://github.com/cloudflare/terraform-provider-cloudflare/issues/212))
* resource/cloudflare_load_balancer_monitor: Add validation for `port` ([#213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/213))
* resource/cloudflare_load_balancer_monitor: Add `allow_insecure` and `follow_redirects` ([#205](https://github.com/cloudflare/terraform-provider-cloudflare/issues/205))
* resource/cloudflare_page_rule: Updated available actions documentation to match what is available ([#228](https://github.com/cloudflare/terraform-provider-cloudflare/issues/228))
* provider: Swap to using go modules for dependency management ([#230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/230))
* provider: Minimum Go version for development is now 1.11 ([#230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/230))

**Fixes:**

* resource/cloudflare_record: Read `data` back from API correctly ([#217](https://github.com/cloudflare/terraform-provider-cloudflare/issues/217))
* resource/cloudflare_rate_limit: Read `correlate` back from API correctly ([#204](https://github.com/cloudflare/terraform-provider-cloudflare/issues/204))
* resource/cloudflare_load_balancer_monitor: Fix incorrect type cast for `port` ([#213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/213))
* resource/cloudflare_load_balancer: Make `steering_policy` computed to avoid spurious diffs ([#214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/214))
* resource/cloudflare_load_balancer: Read `session_affinity` back from API to make import work & detects drifts ([#214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/214))

## 1.11.0 (January 11, 2019)

**Improvements:**
* **New Resource:** `cloudflare_spectrum_app` ([#156](https://github.com/cloudflare/terraform-provider-cloudflare/issues/156))
* **New Data Source:** `cloudflare_zones` ([#168](https://github.com/cloudflare/terraform-provider-cloudflare/issues/168))
* `cloudflare_load_balancer_monitor` - Add optional `port` parameter ([#179](https://github.com/cloudflare/terraform-provider-cloudflare/issues/179))
* `cloudflare_page_rule` - Improved documentation for `priority` attribute ([#182](https://github.com/cloudflare/terraform-provider-cloudflare/issues/182)], missing `explicit_cache_control` [[#185](https://github.com/cloudflare/terraform-provider-cloudflare/issues/185))
* `cloudflare_rate_limit` - Add `challenge` and `js_challenge` rate-limit modes ([#172](https://github.com/cloudflare/terraform-provider-cloudflare/issues/172))

**Fixes:**
* `cloudflare_page_rule` - Page rule `zone` attribute change to trigger new resource ([#183](https://github.com/cloudflare/terraform-provider-cloudflare/issues/183))

## 1.10.0 (December 18, 2018)

**Improvements:**
* `cloudflare_zone_settings_override` - Add `opportunistic_onion` zone setting support ([#170](https://github.com/cloudflare/terraform-provider-cloudflare/issues/170))
* `cloudflare_zone` - Add ability to set zone plan ([#160](https://github.com/cloudflare/terraform-provider-cloudflare/issues/160))

**Fixes:**
* `cloudflare_zone` - Allow zones to be properly imported ([#157](https://github.com/cloudflare/terraform-provider-cloudflare/issues/157))
* `cloudflare_access_policy` - Match access_policy argument requisites with reality ([#158](https://github.com/cloudflare/terraform-provider-cloudflare/issues/158))
* `cloudflare_filter` - Allow `zone_id` to set `zone` and vice versa ([#162](https://github.com/cloudflare/terraform-provider-cloudflare/issues/162))
* `cloudflare_firewall_rule` - Allow `zone_id` to set `zone` and vice versa ([#174](https://github.com/cloudflare/terraform-provider-cloudflare/issues/174))
* `cloudflare_access_rule` - Ensure `zone` and `zone_id` are always set ([#175](https://github.com/cloudflare/terraform-provider-cloudflare/issues/175))
* Minor documentation fixes

## 1.9.0 (November 15, 2018)

**Improvements:**
* **New Resource:** `cloudflare_access_application` ([#145](https://github.com/cloudflare/terraform-provider-cloudflare/issues/145))
* **New Resource:** `cloudflare_access_policy` ([#145](https://github.com/cloudflare/terraform-provider-cloudflare/issues/145))
* `cloudflare_load_balancer` - Add steering policy support ([#147](https://github.com/cloudflare/terraform-provider-cloudflare/issues/147))
* `cloudflare_load_balancer` - Support `session_affinity` ([#153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/153))
* `cloudflare_load_balancer_pool` - Support `weight` ([#153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/153))

**Fixes:**
* `cloudflare_record` - Compare name without the zone name ([#151](https://github.com/cloudflare/terraform-provider-cloudflare/issues/151))
* Minor documentation fixes ([#149](https://github.com/cloudflare/terraform-provider-cloudflare/issues/149)] [[#152](https://github.com/cloudflare/terraform-provider-cloudflare/issues/152))

## 1.8.0 (November 05, 2018)

**Improvements:**
* **New Resource:** `cloudflare_zone` ([#58](https://github.com/cloudflare/terraform-provider-cloudflare/issues/58))
* **New Resource:** `cloudflare_custom_pages` ([#132](https://github.com/cloudflare/terraform-provider-cloudflare/issues/132))
* `cloudflare_zone_settings_override` - Allow setting SSL level to Strict (SSL-Only Origin Pull) ([#122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/122))
* Update provider usage/build docs and how to update a dependency ([#138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/138))
* Improve `Building The Provider` instructions ([#143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/143))
* `cloudflare_access_rule` - Make importable for all rule types ([#141](https://github.com/cloudflare/terraform-provider-cloudflare/issues/141))
* `cloudflare_load_balancer_pool` - Implement `Update` ([#140](https://github.com/cloudflare/terraform-provider-cloudflare/issues/140))

**Fixes:**
* `cloudflare_rate_limit` - Documentation fixes for markdown where \_ALL\_ is italicized ([#125](https://github.com/cloudflare/terraform-provider-cloudflare/issues/125))
* `cloudflare_worker_route` - Correctly set `multi_script` on Enterprise worker imports ([#124](https://github.com/cloudflare/terraform-provider-cloudflare/issues/124))
* `account_member` - Ignore role ID ordering ([#128](https://github.com/cloudflare/terraform-provider-cloudflare/issues/128))
* `cloudflare_rate_limit` - Origin traffic isn't default anymore ([#130](https://github.com/cloudflare/terraform-provider-cloudflare/issues/130))
* `cloudflare_rate_limit` - Update rate limit validation to allow `1` ([#129](https://github.com/cloudflare/terraform-provider-cloudflare/issues/129))
* `cloudflare_record` - Add validation to ensure TTL is not set while `proxied` is true ([#127](https://github.com/cloudflare/terraform-provider-cloudflare/issues/127))
* Updated code for provider version in User-Agent
* `cloudflare_zone_lockdown` - Fix import of zone lockdowns ([#135](https://github.com/cloudflare/terraform-provider-cloudflare/issues/135))

## 1.7.0 (October 09, 2018)

**Improvements:**
* **New Resource:** `cloudflare_account_member` ([#78](https://github.com/cloudflare/terraform-provider-cloudflare/issues/78))

## 1.6.0 (October 05, 2018)

**Improvements:**
* **New Resource:** `cloudflare_filter`
* **New Resource:** `cloudflare_firewall_rule`

## 1.5.0 (September 21, 2018)

**Improvements:**
* **New Resource:** `cloudflare_zone_lockdown` ([#115](https://github.com/cloudflare/terraform-provider-cloudflare/issues/115))

**Fixes:**
* Send User-Agent header with name and version when contacting API
* `cloudflare_page_rule` - Fix page rule polish (off, lossless or lossy) ([#116](https://github.com/cloudflare/terraform-provider-cloudflare/issues/116))

## 1.4.0 (September 11, 2018)

**Improvements:**
* **New Resource:** `cloudflare_worker_route` ([#110](https://github.com/cloudflare/terraform-provider-cloudflare/issues/110))
* **New Resource:** `cloudflare_worker_script` ([#110](https://github.com/cloudflare/terraform-provider-cloudflare/issues/110))

## 1.3.0 (September 04, 2018)

**Improvements:**
* **New Resource:** `cloudflare_access_rule` ([#64](https://github.com/cloudflare/terraform-provider-cloudflare/issues/64))

**Fixes:**
* `cloudflare_zone_settings_override` -  Change Zone Settings Override to use GetOkExists ([#107](https://github.com/cloudflare/terraform-provider-cloudflare/issues/107))

## 1.2.0 (August 13, 2018)

**Improvements:**
* **New Resource:** `cloudflare_waf_rule` ([#98](https://github.com/cloudflare/terraform-provider-cloudflare/issues/98))
* `cloudflare_zone_settings_override` - Add `off` as Security Level setting ([#99](https://github.com/cloudflare/terraform-provider-cloudflare/issues/99))
* `resource_cloudflare_rate_limit` - Add nat support ([#96](https://github.com/cloudflare/terraform-provider-cloudflare/issues/96))
* `resource_cloudflare_zone_settings_override` - Add `zrt` as a value for the `tls_1_3` setting ([#106](https://github.com/cloudflare/terraform-provider-cloudflare/issues/106))
* Minor documentation improvements

**Fixes:**
* `cloudflare_record` - Setting a DNS record's `proxied` flag to false stopped working ([#103](https://github.com/cloudflare/terraform-provider-cloudflare/issues/103))

## 1.1.0 (July 25, 2018)

FIXES:

* `cloudflare_ip_ranges` - IPv6 CIDR blocks should return IPv6 addresses ([#51](https://github.com/cloudflare/terraform-provider-cloudflare/issues/51))
* `cloudflare_zone_settings_override` - Allow `0` for `browser_cache_ttl` ([#71](https://github.com/cloudflare/terraform-provider-cloudflare/issues/71))
* `cloudflare_page_rule` - `forwarding_urls` in page rules are lists ([#79](https://github.com/cloudflare/terraform-provider-cloudflare/issues/79))
* `cloudflare_page_rule` - The API supports `active` and `disabled`, not `paused` ([#84](https://github.com/cloudflare/terraform-provider-cloudflare/issues/84))

IMPROVEMENTS:
* `cloudflare_zone_settings_override` - Add support for `min_tls_version` ([#72](https://github.com/cloudflare/terraform-provider-cloudflare/issues/72))
* `cloudflare_page_rule` - Add support for more settings: `bypass_cache_on_cookie`, `cache_by_device_type`, `cache_deception_armor`, `cache_on_cookie`, `host_header_override`, `polish`, `explicit_cache_control`, `origin_error_page_pass_thru`, `sort_query_string_for_cache`, `resolve_override`, `respect_strong_etag`, `response_buffering`, `true_client_ip_header`, `mirage`, `disable_railgun`, `cache_key`, `waf`, `rocket_loader`, `cname_flattening` ([#68](https://github.com/cloudflare/terraform-provider-cloudflare/issues/68)], [[#81](https://github.com/cloudflare/terraform-provider-cloudflare/issues/81)], [[#85](https://github.com/cloudflare/terraform-provider-cloudflare/issues/85))
* `cloudflare_page_rule` - Add `off` setting to `security_level` ([#81](https://github.com/cloudflare/terraform-provider-cloudflare/issues/81))
* `cloudflare_record` - DNS Record improvements ([#97](https://github.com/cloudflare/terraform-provider-cloudflare/issues/97))
* Various documentation improvements

## 1.0.0 (April 06, 2018)

BACKWARDS INCOMPATIBILITIES / NOTES:

* resource/cloudflare_record: Changing `name` or `domain` now force a recreation
  of the record ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))

FEATURES:

* **New Resource:** `cloudflare_rate_limit` ([#30](https://github.com/cloudflare/terraform-provider-cloudflare/issues/30))
* **New Resource:** `cloudflare_page_rule` ([#38](https://github.com/cloudflare/terraform-provider-cloudflare/issues/38))
* **New Resource:** `cloudflare_load_balancer` ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))
* **New Resource:** `cloudflare_load_balancer_pool` ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))
* **New Resource:** `cloudflare_zone_settings_override` ([#41](https://github.com/cloudflare/terraform-provider-cloudflare/issues/41))
* **New Resource:** `cloudflare_load_balancer_monitor` ([#42](https://github.com/cloudflare/terraform-provider-cloudflare/issues/42))
* **New Data Source:** `cloudflare_ip_ranges` ([#28](https://github.com/cloudflare/terraform-provider-cloudflare/issues/28))

IMPROVEMENTS:

* resource/cloudflare_record: Validate `TXT` records ([#14](https://github.com/cloudflare/terraform-provider-cloudflare/issues/14))
* resource/cloudflare_record: Add `data` input to suppport SRV, LOC records
  ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))
* resource/cloudflare_record: Add computed attributes `created_on`,
  `modified_on`, `proxiable`, and `metadata` to records ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))
* resource/cloudflare_record: Support import of existing records ([#36](https://github.com/cloudflare/terraform-provider-cloudflare/issues/36))
* New Provider configuration options for API rate limiting ([#43](https://github.com/cloudflare/terraform-provider-cloudflare/issues/43))
* New Provider configuration options for using Organizations ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))

## 0.1.0 (June 20, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider
  Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
## 2.26.0 (August 27th, 2021)

* **New resource**: `cloudflare_notification_policy` ([#1138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1138))
* **New resource**: `cloudflare_notification_policy_webhooks` ([#1151](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1151))
* **New resource**: `cloudflare_ruleset` ([#1143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1143))
* **New resource**: `cloudflare_teams_location` ([#1154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1154))
* **New datasource**: `cloudflare_origin_ca_root_certificate` ([#1158](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1158))

**Improvements**

* `resource/cloudflare_waiting_room`: Add support for `json_response_enabled` as an argument ([#1122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1122))

## 2.25.0 (August 4th, 2021)

**Improvements**

* `resource/cloudflare_access_device_posture_rule`: Add support for `domain_joined`, `firewall`, `os_version`, and `disk_encryption` ([#1137](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1137))
* provider: bump `cloudflare-go` to v0.20.0 ([#1146](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1146))

## 2.24.0 (July 19th, 2021)

**Improvements**

* `resource/cloudflare_logpush_job`: Add support for `"nel_reports"` as a dataset ([#1122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1122))
* `resource/cloudflare_custom_hostname`: Allow SSL options to be optional when not required ([#1131](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1131))
* `resource/cloudflare_access_identity_provider`: Support optional Okta API token ([#1119](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1119))
* `resource/cloudflare_load_balancer_pool`: Add support for load shedding ([#1108](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1108))
* `resource/cloudflare_load_balancer_pool`: Add support for longitude and latitude ([#1093](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1093))

**Fixes**

* `resource/cloudflare_record`: Use correct `Import` method on resource ([#1116](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1116))
* `resource/cloudflare_worker_cron_trigger`: Account for deletion of scripts and force a refresh of triggers ([#1121](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1121))
* `resource/cloudflare_rate_limit`: Handle `origin_traffic` missing from API response ([#1125](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1125))
* `resource/cloudflare_record`: Support `allow_overwrite` for root records ([#1129](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1129))

## 2.23.0 (June 30th, 2021)

- **New resource**: `cloudflare_waiting_room` ([#1053](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1053))

**Improvements**

* `datasource/cloudflare_waf_rules`: Export `default_mode` as an attribute ([#1079](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1079))

**Fixes**

* `resource/cloudflare_access_application`: Revert removal of schema changes causing existing applications unable to re-apply ([#1118](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1118))

## 2.22.0 (June 25th, 2021)

- **New resource**: `cloudflare_static_route` ([#1098](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1098))

**Improvements**

* `resource/cloudflare_origin_ca`: Ignore decreasing `requested_validity` ([#1043](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1078))
* `resource/waf_override`: Allow `rules` to be optional ([#1090](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1090))
* `resource/cloudflare_zone`: Don't attempt to set free zone rate plans as that is already the default ([#1102](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1102))
* `resource/cloudflare_access_application`: Ability to set `type` for Applications ([#1076](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1076))
* `resource/cloudflare_zone_lockdown`: Update documentation to show examples of multiple configurations ([#1106](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1106))

## 2.21.0 (May 26th, 2021)

- **New resource**: `cloudflare_device_posture_rule` ([#1058](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1058))
- **New resource**: `cloudflare_teams_list` ([#1058](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1058))

**Improvements**

* provider: Update to terraform-plugin-sdk v1.17.1 ([#1035](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1035), [#1043](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1043))
* `resource/cloudflare_logpush_job`: Allow `ownership_challenge` to be optional to account for Datadog, Splunk or S3-Compatible endpoints ([#1048](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1048))
* `resource/cloudflare_access_group`: Add support for `login_method` ([#1066](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1066))
* `resource/cloudflare_load_balancer`: Add support for `promixity` based steering ([#1072](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1072))
* `resource/cloudflare_access_application`: Prevent bad CORS configuration when credentials and all origins are permitted ([#1073](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1073))
* `resource/cloudflare_access_service_tokens`: Allow configuration to manage automatic renewal when the threshold is crossed and Terraform operations are performed within the window ([#1057](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1057))
* `resource/cloudflare_load_balancer_pool`: Allow support for `Host` header settings ([#1042](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1042))

**Fixes**

* `resource/cloudflare_access_policy`: Allow empty slices in blocks when building policies ([#1034](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1034))
* `resource/cloudflare_load_balancer`: Fix `override` attributes `pop_pools` and `region_pools` referencing incorrect values causing a panic ([#1039](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1039))

## 2.20.0 (April 15th, 2021)

**New resource**: `cloudflare_access_ca_certificate` ([#995](https://github.com/cloudflare/terraform-provider-cloudflare/issues/995))

**Improvements**

* `resource/cloudflare_access_application`: Improve documentation for `Import` usage ([#1002](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1002))
* `resource/cloudflare_logpush_job`: Update documentation to reflect requirements for `destination_conf` to match across all uses ([#1024](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1024))
* `resource/cloudflare_custom_hostname_fallback`: Better handle service lag when updating existing resources by attempting retries ([#1014](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1014))
* `resource/cloudflare_waf_group`: Simplify error handling using inbuilt helpers ([#1015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1015))
* `resource/cloudflare_waf_rule`: Simplify error handling using inbuilt helpers ([#1015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1015))
* `resource/cloudflare_waf_package`: Simplify error handling using inbuilt helpers ([#1015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1015))
* `resource/cloudflare_access_group`: Add support for `login_method` ([#1018](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1018))
* provider: Update to cloudflare-go v0.16.0 ([#1018](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1018))
* provider: Update to terraform-plugin-sdk v1.16.1 ([#1003](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1003))
* `resource/cloudflare_load_balancer`: Add support for `rules` ([#1016](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1016))

## 2.19.2 (March 15th, 2021)

**Fixes**

* `resource/cloudflare_record`: Address regression from 2.19.1 by checking the API response instead of the schema output for `Priority` ([#992](https://github.com/cloudflare/terraform-provider-cloudflare/issues/992))

## 2.19.1 (March 11th, 2021)

**Fixes**

* `resource/cloudflare_record`: Update `Priority` handling for MX parked records ([#986](https://github.com/cloudflare/terraform-provider-cloudflare/issues/986))

## 2.19.0 (March 10th, 2021)

**Fixes**

* `resource/cloudflare_access_group`: Fix crash when constructing a GSuite group ([#940](https://github.com/cloudflare/terraform-provider-cloudflare/issues/940))
* `resource/cloudflare_access_policy`: Make `precedence` required ([#941](https://github.com/cloudflare/terraform-provider-cloudflare/issues/941))
* `resource/cloudflare_access_group`: Fix crash when constructing a SAML group ([#948](https://github.com/cloudflare/terraform-provider-cloudflare/issues/948))
* `resource/cloudflare_zone`: Update `Retry` logic to look at an available field for passing conditions ([#973](https://github.com/cloudflare/terraform-provider-cloudflare/issues/973))
* `resource/cloudflare_page_rule`: Allow ignoring/including all query string parameters for `cache_key_fields` ([#975](https://github.com/cloudflare/terraform-provider-cloudflare/issues/975))

**Improvements**

* `resource/cloudflare_access_policy`: Enable zone and account level resources to be imported ([#956](https://github.com/cloudflare/terraform-provider-cloudflare/issues/956))
* `resource/cloudflare_origin_ca_certificate`: Smoother import process with less recreation ([#955](https://github.com/cloudflare/terraform-provider-cloudflare/issues/955))
* provider: Update internals to match `cloudflare-go` 0.14 for better error handling and context aware methods ([#976](https://github.com/cloudflare/terraform-provider-cloudflare/issues/976))

## 2.18.0 (February 3rd, 2021)

* **New Resource:** `cloudflare_argo_tunnel` ([#905](https://github.com/cloudflare/terraform-provider-cloudflare/issues/905))
* **New Resource:** `cloudflare_worker_cron_trigger` ([#926](https://github.com/cloudflare/terraform-provider-cloudflare/issues/926))

**Fixes**

* `datasource/cloudflare_zones`: Pagination is now correctly handled internally and will return more than the single page of results ([cloudflare/cloudflare-go#534](https://github.com/cloudflare/cloudflare-go/pull/534)).
* `resource/cloudflare_access_policy`: Correctly handle transforming API responses to schema ([#917](https://github.com/cloudflare/terraform-provider-cloudflare/issues/917))
* `resource/cloudflare_access_group`: Correctly handle transforming API responses to schema ([#918](https://github.com/cloudflare/terraform-provider-cloudflare/issues/918))
* `resource/cloudflare_ip_list`: Ensure account ID is persisted during `Import` ([#916](https://github.com/cloudflare/terraform-provider-cloudflare/issues/916))

**Improvements**

* `resource/cloudflare_access_application`: Allow any `session_duration` that is `time.ParseDuration` compatible ([#910](https://github.com/cloudflare/terraform-provider-cloudflare/issues/910))
* `resource/cloudflare_rate_limit`: Add the ability to configure `match.response.headers` in rate limits ([#911](https://github.com/cloudflare/terraform-provider-cloudflare/issues/911))
* `resource/cloudflare_access_rule`: Validate IP masks within schema ([#921](https://github.com/cloudflare/terraform-provider-cloudflare/issues/921))

## 2.17.0 (January 5th, 2021)

* **New Resource:** `cloudflare_magic_firewall_ruleset` ([#884](https://github.com/cloudflare/terraform-provider-cloudflare/issues/884))

**Fixes**

* `resource/cloudfare_api_token`: Omitting `conditions` will no longer send empty arrays causing IP restriction issues and unusable tokens ([#902](https://github.com/cloudflare/terraform-provider-cloudflare/pull/902))

## 2.16.0 (January 5th, 2021)

**Improvements**

* `resource/cloudflare_access_application`: Add support for `custom_deny_message` and `custom_deny_url` values ([#895](https://github.com/cloudflare/terraform-provider-cloudflare/issues/895))
* `resource/cloudflare_load_balancer_monitor`: Add support for `probe_zone` for monitors ([#903](https://github.com/cloudflare/terraform-provider-cloudflare/issues/903))

## 2.15.0 (December 29th, 2020)

**Improvements**

* `resource/cloudflare_load_balancer`: Add support for `session_affinity_ttl` ([#882](https://github.com/cloudflare/terraform-provider-cloudflare/issues/882))
* `resource/cloudflare_load_balancer`: Add support for `session_affinity_attributes` ([#883](https://github.com/cloudflare/terraform-provider-cloudflare/issues/883))

**Fixes**

* `resource/cloudflare_page_rule`: Fixed crash during update when using custom cache key ([#894](https://github.com/cloudflare/terraform-provider-cloudflare/pull/894))

## 2.14.0 (November 26th, 2020)

* **New Resource:** `cloudflare_api_token` ([#862](https://github.com/cloudflare/terraform-provider-cloudflare/issues/862))
* **New Datasource:** `cloudflare_api_token_permission_groups` ([#862](https://github.com/cloudflare/terraform-provider-cloudflare/issues/862))
* **New Resource:** `cloudflare_zone_dnssec` ([#852](https://github.com/cloudflare/terraform-provider-cloudflare/issues/852))
* **New Datasource:** `cloudflare_zone_dnssec` ([#852](https://github.com/cloudflare/terraform-provider-cloudflare/issues/852))

**Improvements**

* `resource/cloudflare_record`: Add explicit fields for CAA records instead of relying on the map value ([#866](https://github.com/cloudflare/terraform-provider-cloudflare/issues/866))
* `resource/cloudflare_account_member`: Swap schema `role_ids` to `TypeSet` to better handle internal ordering changes ([#876](https://github.com/cloudflare/terraform-provider-cloudflare/issues/876))

**Fixes**

* `datasource/cloudflare_waf_groups`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))
* `datasource/cloudflare_waf_packages`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))
* `datasource/cloudflare_waf_rules`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))
* `datasource/cloudflare_zones`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))

## 2.13.2 (November 6th, 2020)

**Fixes**

* `resource/cloudflare_filter`: Remove schema based validation for filters ([#863](https://github.com/cloudflare/terraform-provider-cloudflare/issues/863))

## 2.13.1 (November 5th, 2020)

**Improvements**

* `resource/cloudflare_filter`: Pass missing credential error through to end user ([#860](https://github.com/cloudflare/terraform-provider-cloudflare/issues/860))

## 2.13.0 (November 5th, 2020)

**Improvements**

* `datasource/cloudflare_ip_ranges`: Add the ability to query `china_ipv4_cidr_blocks` and `china_ipv6_cidr_blocks` ([#833](https://github.com/cloudflare/terraform-provider-cloudflare/issues/833))
* `resource/cloudflare_filter`: Improve validation of expressions using the schema ([#848](https://github.com/cloudflare/terraform-provider-cloudflare/issues/848))

**Fixes**

* `resource/cloudflare_page_rule`: Set default for `cache_key_fields.host.resolved` to prevent panics ([#832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/832))
* `resource/cloudflare_authenticated_origin_pulls`: Fix off-by-one error check in `Import` ([#832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/859))
* `resource/cloudflare_authenticated_origin_pulls_certificate`: Fix off-by-one error check in `Import` ([#832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/859))

## 2.12.0 (October 22nd, 2020)

**Improvements**

* `resource/cloudflare_certificate_pack`: Swap internal representation of `hosts` to remove inconsistent ordering issues ([#800](https://github.com/cloudflare/terraform-provider-cloudflare/issues/800))
* `resource/cloudflare_logpush_job`: Handle deletion outside of Terraform ([#798](https://github.com/cloudflare/terraform-provider-cloudflare/issues/798))
* `resource/cloudflare_access_group`: Add support for `geo` conditionals ([#803](https://github.com/cloudflare/terraform-provider-cloudflare/issues/803))
* `resource/cloudflare_access_application`: Add support for `enable_binding_cookie` ([#802](https://github.com/cloudflare/terraform-provider-cloudflare/issues/802))
* `resource/cloudflare_waf_rule`: Improve documentation for `mode` ([#824](https://github.com/cloudflare/terraform-provider-cloudflare/issues/824))
* `datasource/cloudflare_waf_rule`: Improve documentation for `mode` ([#824](https://github.com/cloudflare/terraform-provider-cloudflare/issues/824))
* `resource/cloudflare_access_application`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))
* `resource/cloudflare_access_group`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))
* `resource/cloudflare_access_identity_provider`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))
* `resource/cloudflare_access_policy`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))

**Fixes**

* `resource/cloudflare_custom_hostname_fallback_origin`: Don't retry the "active" status of custom hostnames fallbacks ([#818](https://github.com/cloudflare/terraform-provider-cloudflare/issues/818))
* `resource/cloudflare_zone`: Remove `DiffSuppressFunc` causing `jump_start` issues ([#830](https://github.com/cloudflare/terraform-provider-cloudflare/issues/830))

## 2.11.0 (September 11th, 2020)

* **New Resource:** `cloudflare_certificate_pack` ([#778](https://github.com/cloudflare/terraform-provider-cloudflare/issues/778))

**Improvements**

* `resource/cloudflare_access_group`: Add support for `auth_method` ([#762](https://github.com/cloudflare/terraform-provider-cloudflare/issues/762))
* `resource/cloudflare_access_group`: De-duplicate blocks in groups by accepting lists instead ([#739](https://github.com/cloudflare/terraform-provider-cloudflare/issues/739))
* `resource/cloudflare_worker_script`: Adds support for `webassembly_binding` ([#780](https://github.com/cloudflare/terraform-provider-cloudflare/issues/780))
* `resource/cloudflare_healthcheck`: Retry hostname resolution errors when encountering "no such host" responses ([#789](https://github.com/cloudflare/terraform-provider-cloudflare/issues/789))
* `resource/cloudflare_access_application`: Better validation for allowed methods and origin combinations to prevent getting state into an unrecoverable state ([#793](https://github.com/cloudflare/terraform-provider-cloudflare/issues/793))

**Fixes**

* `resource/cloudflare_healthcheck`: Handle resource deletion outside of Terraform ([#787](https://github.com/cloudflare/terraform-provider-cloudflare/issues/787))
* `resource/cloudflare_custom_hostname`: Ensure `Import` sets hostname to prevent recreation ([#788](https://github.com/cloudflare/terraform-provider-cloudflare/issues/788))
* `resource/cloudflare_ip_list`: Handle resource deletion outside of Terraform ([#794](https://github.com/cloudflare/terraform-provider-cloudflare/issues/794))
* `resource/cloudflare_ip_list`: Remove `item`.`id` from schema ([#796](https://github.com/cloudflare/terraform-provider-cloudflare/issues/796))

## 2.10.1 (August 24th, 2020)

**Fixes**

* `resource/cloudflare_access_application`: Handle the `zone_id` => `account_id` move internally ([#724](https://github.com/cloudflare/terraform-provider-cloudflare/issues/724))

## 2.10.0 (August 24th, 2020)

* **New Resource:** `cloudflare_custom_hostname_origin_fallback` ([#757](https://github.com/cloudflare/terraform-provider-cloudflare/issues/757))
* **New Resource:** `cloudflare_authenticated_origin_pulls` ([#749](https://github.com/cloudflare/terraform-provider-cloudflare/issues/749))
* **New Resource:** `cloudflare_authenticated_origin_pulls_certificate` ([#749](https://github.com/cloudflare/terraform-provider-cloudflare/issues/749))
* **New Resource:** `cloudflare_ip_list` ([#766](https://github.com/cloudflare/terraform-provider-cloudflare/issues/766))

**Improvements**

* `resource/cloudflare_spectrum_application`: Add support for port ranges ([#745](https://github.com/cloudflare/terraform-provider-cloudflare/issues/745))
* `resource/cloudflare_custom_hostname`: Force creation of a new resource if the `zone_id` value changes ([#761](https://github.com/cloudflare/terraform-provider-cloudflare/issues/761))
* `resource/cloudflare_record`: Retry record creation/update if the response includes an "already exists" exception for handling race conditions ([#773](https://github.com/cloudflare/terraform-provider-cloudflare/issues/773))

**Fixes**

* `resource/cloudflare_firewall_rule`: Compare descriptions after converting unicode + HTML entities to prevent unnecessary diffs ([#758](https://github.com/cloudflare/terraform-provider-cloudflare/issues/758))
* `resource/cloudflare_filter`: Compare descriptions after converting unicode + HTML entities to prevent unnecessary diffs ([#758](https://github.com/cloudflare/terraform-provider-cloudflare/issues/758))

## 2.9.0 (July 30th, 2020)

* **New Resource:** `cloudflare_custom_hostname` (SSL for SaaS) ([#746](https://github.com/cloudflare/terraform-provider-cloudflare/issues/746))

**Improvements**

* `resource/access_application`: Add support for `allowed_idps` and restricting which Identity Providers are associated with an Application ([#734](https://github.com/cloudflare/terraform-provider-cloudflare/issues/734))
* `resource/access_application`: Add support for `auto_redirect_to_identity` ([#730](https://github.com/cloudflare/terraform-provider-cloudflare/issues/730))
* `resource/access_application`: Add CORS support ([#725](https://github.com/cloudflare/terraform-provider-cloudflare/issues/725))
* `resource/cloudflare_custom_ssl`: Allow `geo_restrictions` to be `nil` and not included in the request payload ([#714](https://github.com/cloudflare/terraform-provider-cloudflare/issues/714))
* `datasource/cloudflare_zones`: Filtering is now performed on the server side and the `name` parameter is no longer a regex. Instead, `name` is a string to match on and `match` is a regex. See the website documentation for more examples and updated references ([#708](https://github.com/cloudflare/terraform-provider-cloudflare/issues/708)) in order to make your code compatible with this release.


## 2.8.0 (June 22, 2020)

* **New Resource:** `cloudflare_waf_override` ([#691](https://github.com/cloudflare/terraform-provider-cloudflare/issues/691))

**Improvements**

* `resource/cloudflare_argo`: Allow `tiered_caching` and `smart_routing` to be toggled individually allowing for entitlement differences ([#703](https://github.com/cloudflare/terraform-provider-cloudflare/issues/703))
* `resource/cloudflare_page_rule`: Add support for `cache_ttl_by_status` ([#706](https://github.com/cloudflare/terraform-provider-cloudflare/issues/706))
* `resource/cloudflare_worker_script`: Add support for `plain_text` and `secret_text` bindings ([#710](https://github.com/cloudflare/terraform-provider-cloudflare/issues/710))

**Fixes**

* `resource/cloudflare_record`: Update `TestAccCloudflareRecord_LOC` test asserted value to use less precise floats and match the API responses ([#712](https://github.com/cloudflare/terraform-provider-cloudflare/issues/712))
* `resource/cloudflare_record`: Update `TestAccCloudflareRecord_Basic` test `metadata` attributes to match updated API payload ([#713](https://github.com/cloudflare/terraform-provider-cloudflare/issues/713))

## 2.7.0 (May 20, 2020)

* **New Resource:** `cloudflare_byo_ip_prefix` ([#671](https://github.com/cloudflare/terraform-provider-cloudflare/issues/671))
* **New Resource:** `cloudflare_logpull_retention` ([#678](https://github.com/cloudflare/terraform-provider-cloudflare/issues/678))
* **New Resource:** `cloudflare_healthcheck` ([#680](https://github.com/cloudflare/terraform-provider-cloudflare/issues/680))

**Improvements:**

* `resource/cloudflare_worker_route`: Improve documentation to mention using `account_id` for the underlying APIs ([#669](https://github.com/cloudflare/terraform-provider-cloudflare/issues/669))
* `resource/cloudflare_worker_script`: Improve documentation to mention using `account_id` for the underlying APIs ([#670](https://github.com/cloudflare/terraform-provider-cloudflare/issues/670))
* `resource/cloudflare_load_balancer_pool`: Improve documentation to mention `notification_email` accepts a comma delimited list of emails ([#687](https://github.com/cloudflare/terraform-provider-cloudflare/issues/687))
* `resource/cloudflare_page_rule`: Add support for `cache_key_fields` Page Rule action ([#662](https://github.com/cloudflare/terraform-provider-cloudflare/issues/662))

**Fixes:**
* `resource/cloudflare_zone_settings_override`: Fix regression where if you didn't have universal SSL settings defined, it would error when setting them ([#663](https://github.com/cloudflare/terraform-provider-cloudflare/issues/663))
* `resource/cloudflare_zone`: Handle changing zone rate plan from "free" to "enterprise" ([#668](https://github.com/cloudflare/terraform-provider-cloudflare/issues/668))
* `resource/cloudflare_record`: Update validation to allow PTR records ([9a8fd43](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9a8fd43))

## 2.6.0 (April 22, 2020)

**Improvements:**

* `resource/cloudflare_zone_settings_override`: Add `universal_ssl` to control enablement of Universal SSL on a zone ([#658](https://github.com/cloudflare/terraform-provider-cloudflare/issues/658))
* provider: API keys and API tokens are now validated to help differentiate incorrect usage before making API calls ([#661](https://github.com/cloudflare/terraform-provider-cloudflare/issues/661))
* `resource/cloudflare_logpush_job`: Add support for "firewall_events" dataset parameter ([#660](https://github.com/cloudflare/terraform-provider-cloudflare/issues/660))
* `resource/cloudflare_logpush_job`: Add support for "dataset" parameter ([#649](https://github.com/cloudflare/terraform-provider-cloudflare/issues/649))
* `resource/cloudflare_zone_settings_override`: Remove `edge_cache_ttl` ([#654](https://github.com/cloudflare/terraform-provider-cloudflare/issues/654))
* `resource/cloudflare_access_group`: Allow Access conditions for `include`/`require`/`exclude` to be used consistently between Access Groups and Access Policies ([#646](https://github.com/cloudflare/terraform-provider-cloudflare/issues/646))

**Fixes:**
* `resource/cloudflare_logpush_job`: fix for `strconv.Atoi: parsing ""` error while creating Logpush job

## 2.5.1 (April 03, 2020)

**Improvements:**

* `resource/cloudflare_zone_settings_override`: Update `image_resizing` options to include `"open"` ([#639](https://github.com/cloudflare/terraform-provider-cloudflare/issues/639))

**Fixes:**

* `resource/cloudflare_access_group`: Fixed misspelt Okta in JSON payload ([cloudflare/cloudflare-go#440](https://github.com/cloudflare/cloudflare-go/issues/440))

## 2.5.0 (March 27, 2020)

**Improvements:**

* `resource/cloudflare_access_policy`: Add support for `service_token` and `any_valid_service_token` ([#612](https://github.com/cloudflare/terraform-provider-cloudflare/issues/612))
* `resource/cloudflare_waf_group`: Handle WAF group deletions in the API responses ([#623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/623))
* `resource/cloudflare_waf_package`: Handle WAF package deletions in the API responses ([#623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/623))
* `resource/cloudflare_waf_rule`: Handle WAF rule deletions in the API responses ([#623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/623))
* `resource/cloudflare_access_policy`: Add support for `group` ([#626](https://github.com/cloudflare/terraform-provider-cloudflare/issues/626))
* `resource/cloudflare_firewall_rule`: Add support for bypassing specific `products` ([#630](https://github.com/cloudflare/terraform-provider-cloudflare/issues/630))
* `resource/cloudflare_spectrum_application`: Add support for `edge_ips`, `argo_smart_routing` and `edge_ip_connectivity` ([#631](https://github.com/cloudflare/terraform-provider-cloudflare/issues/631))
* `resource/cloudflare_access_group`: Add support for using external providers (`gsuite`, `github`, `azure`, `okta`, `saml`, `mTLS certificate`, `common name`
) ([#633](https://github.com/cloudflare/terraform-provider-cloudflare/issues/633))

## 2.4.1 (March 12, 2020)

**Improvements:**

* `resource/cloudflare_logpush_job`: Support `Import` on the resource ([#618](https://github.com/cloudflare/terraform-provider-cloudflare/issues/618))

**Fixes:**

* `resource/cloudflare_record`: Missing CAA in DNS validation ([#619](https://github.com/cloudflare/terraform-provider-cloudflare/issues/619))

## 2.4.0 (March 09, 2020)

* **New Resource:** `cloudflare_workers_kv` ([#595](https://github.com/cloudflare/terraform-provider-cloudflare/issues/595))
* **New Resource:** `cloudflare_access_identity_provider` ([#597](https://github.com/cloudflare/terraform-provider-cloudflare/issues/597))


**Improvements:**

* `resource/cloudflare_record`: Stricter validation for record types ([#610](https://github.com/cloudflare/terraform-provider-cloudflare/issues/610))
* `resource/logpush_job`: Add more verbose error handling ([#564](https://github.com/cloudflare/terraform-provider-cloudflare/issues/564))
* `resource/zone_settings_override`: Update documentation for `cache_level` values ([#606](https://github.com/cloudflare/terraform-provider-cloudflare/issues/606))
* `resource/access_application`: Add documentation for available attributes ([#587](https://github.com/cloudflare/terraform-provider-cloudflare/issues/587))
* `resource/cloudflare_firewall_rule`: Add support for bypassing security configuration rules by URL ([#568](https://github.com/cloudflare/terraform-provider-cloudflare/issues/568))
* `resource/cloudflare_record_migrate`: Use `zone_id` for state migration before attempting to use `domain` ([#566](https://github.com/cloudflare/terraform-provider-cloudflare/issues/566))
* `resource/cloudflare_load_balancer`: Update `session_affinity` validation to allow `"ip_cookie"` ([#573](https://github.com/cloudflare/terraform-provider-cloudflare/issues/573))
* `datasource/ip_ranges`: Update documentation to show 0.12 syntax ([#617](https://github.com/cloudflare/terraform-provider-cloudflare/issues/617))

**Fixes**

* `resource/zone_settings_override`: Handle individual zone settings within `Delete` operations ([#599](https://github.com/cloudflare/terraform-provider-cloudflare/issues/599))

## 2.3.0 (December 18, 2019)

* **New Resource:** `cloudflare_origin_ca_certificate` ([#547](https://github.com/cloudflare/terraform-provider-cloudflare/issues/547))

**Fixes:**

* `resource/cloudflare_zone_settings_override`: Renamed `0rtt` to `zero_rtt` to conform to HCL grammar requirements ([#557](https://github.com/cloudflare/terraform-provider-cloudflare/issues/557))

**Improvements:**

* `resource/cloudflare_access_rule`: Add `ip6` as valid option ([#560](https://github.com/cloudflare/terraform-provider-cloudflare/issues/560))
* `resource/cloudflare_spectrum_application`: Swap `proxy_protocol` to string field with supporting enum values instead ([#561](https://github.com/cloudflare/terraform-provider-cloudflare/issues/561))
* `resource/cloudflare_waf_rule`: Add `package_id` as valid option and export `group_id` ([#552](https://github.com/cloudflare/terraform-provider-cloudflare/issues/552))

## 2.2.0 (December 05, 2019)

* **New Resource:** `cloudflare_access_group` ([#510](https://github.com/cloudflare/terraform-provider-cloudflare/issues/510))
* **New Resource:** `cloudflare_workers_kv_namespace` ([#443](https://github.com/cloudflare/terraform-provider-cloudflare/issues/443))

**Improvements:**

* `resource/cloudflare_zone_settings_override`: Add `non_identity` to allowed `decision` schema ([#541](https://github.com/cloudflare/terraform-provider-cloudflare/issues/541))
* `resource/cloudflare_zone_settings_override`: Add support for `0rtt` and `http3` settings ([#542](https://github.com/cloudflare/terraform-provider-cloudflare/issues/542))
* `resource/cloudflare_load_balancer_monitor`: Allow empty string for `expected_body` ([#539](https://github.com/cloudflare/terraform-provider-cloudflare/issues/539))
* `resource/cloudflare_worker_script`: Add support for Worker KV Namespace Bindings ([#544](https://github.com/cloudflare/terraform-provider-cloudflare/issues/544))
* `data_source/waf_rules`, `resource/cloudflare_waf_rule`, Support allowed modes for WAF Rules ([#550](https://github.com/cloudflare/terraform-provider-cloudflare/issues/550))

**Fixes:**

* `resource/cloudflare_spectrum_application`: Spectrum origin_port is optional ([#549](https://github.com/cloudflare/terraform-provider-cloudflare/issues/549))

## 2.1.0 (November 07, 2019)

* **New datasource:** `cloudflare_waf_rules` ([#525](https://github.com/cloudflare/terraform-provider-cloudflare/issues/525))

**Improvements:**

* `resource/cloudflare_zone`: Expose `verification_key` for partial setups ([#532](https://github.com/cloudflare/terraform-provider-cloudflare/issues/532))
* `resource/cloudflare_worker_route`: Enable API Tokens support from upstream [cloudflare-go](https://github.com/cloudflare/cloudflare-go) release

## 2.0.1 (October 22, 2019)

* **New Resource:** `cloudflare_access_service_tokens` ([#521](https://github.com/cloudflare/terraform-provider-cloudflare/issues/521))
* **New Resource:** `cloudflare_waf_package` ([#475](https://github.com/cloudflare/terraform-provider-cloudflare/issues/475))
* **New Resource:** `cloudflare_waf_group` ([#476](https://github.com/cloudflare/terraform-provider-cloudflare/issues/476))
* **New datasource:** `cloudflare_waf_groups` ([#508](https://github.com/cloudflare/terraform-provider-cloudflare/issues/508))
* **New datasource:** `cloudflare_waf_packages` ([#509](https://github.com/cloudflare/terraform-provider-cloudflare/issues/509))

**Fixes:**

* `resource/cloudflare_page_rule`: Set `h2_prioritization` individually not via bulk endpoint ([#493](https://github.com/cloudflare/terraform-provider-cloudflare/issues/493))
* `resource/cloudflare_zone_settings_override`: Set `zone_id` to prevent unnecessary re-creation of resources ([#502](https://github.com/cloudflare/terraform-provider-cloudflare/issues/502))

**Improvements:**

* `resource/cloudflare_spectrum_application`: Add support for setting `traffic_type` ([#481](https://github.com/cloudflare/terraform-provider-cloudflare/issues/481))
* `resource/cloudflare_zone_settings_override`: Update documentation with default values ([#498](https://github.com/cloudflare/terraform-provider-cloudflare/issues/498))

**Internals:**

* Migrated to Terraform plugin SDK ([#489](https://github.com/cloudflare/terraform-provider-cloudflare/issues/489))

## 2.0.0 (September 30, 2019)

**Breaking changes:**
* `provider/cloudflare`:
 * renamed `token` to `api_key`
 * renamed `org_id` to `account_id`
 * removed `use_org_from_zone`, you need to explicitly specify `account_id`
* Environment variables:
 * renamed `CLOUDFLARE_TOKEN` to `CLOUDFLARE_API_TOKEN`
 * renamed `CLOUDFLARE_ORG_ID` to `CLOUDFLARE_ACCOUNT_ID`
 * removed `CLOUDFLARE_ORG_ZONE`, you need to explicitly specify `CLOUDFLARE_ACCOUNT_ID`
* Changed the following resources to require Zone ID:
 * `cloudflare_access_rule`
 * `cloudflare_filter`
 * `cloudflare_firewall_rule`
 * `cloudflare_load_balancer`
 * `cloudflare_page_rule`
 * `cloudflare_rate_limit`
 * `cloudflare_record`
 * `cloudflare_waf_rule`
 * `cloudflare_worker_route"`
 * `cloudflare_zone_lockdown`
 * `cloudflare_zone_settings_override`
* Workers single-script support removed

Please see [Version 2 Upgrade Guide](https://www.terraform.io/docs/providers/cloudflare/guides/version-2-upgrade.html) for details.

**Improvements:**

* `cloudflare/resource_cloudflare_argo`: Handle errors when fetching tiered caching + smart routing settings ([#477](https://github.com/cloudflare/terraform-provider-cloudflare/issues/477))
* Various documentation updates for 0.12 syntax

## 1.18.1 (August 29, 2019)

**Fixes:**

* `resource/cloudflare_load_balancer`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_page_rule`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_rate_limit`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_waf_rule`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_worker_route`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_worker_script`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
* `resource/cloudflare_zone_lockdown`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))

## 1.18.0 (August 27, 2019)

**Fixes:**

* `resource/cloudflare_page_rule`: Fix a logic condition where setting `edge_cache_ttl` action but then not updating it in subsequent `apply` runs causes it to be blown away ([#453](https://github.com/cloudflare/terraform-provider-cloudflare/issues/453))

**Improvements:**

* provider: You can now use API tokens to authenticate instead of user email and key ([#450](https://github.com/cloudflare/terraform-provider-cloudflare/issues/450))
* `resource/cloudflare_zone_lockdown`: `priority` can now be set on the resource ([#445](https://github.com/cloudflare/terraform-provider-cloudflare/issues/445))
* `resource/cloudflare_custom_ssl`: Updated website documentation navigation to include link for resource ([#442)](https://github.com/cloudflare/terraform-provider-cloudflare/issues/442))

**Deprecations:**

* `resource/cloudflare_access_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_filter`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_firewall_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_load_balancer`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_page_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_rate_limit`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_waf_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_worker_route`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_worker_script`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
* `resource/cloudflare_zone_lockdown`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))

## 1.17.1 (August 09, 2019)

**Fixes:**

* Partially revert [[#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421)] deprecation messages

## 1.17.0 (August 09, 2019)

**Removals:**

* `resource/cloudflare_zone_settings_override`: `sha1_support` has been removed due to Cloudflare no longer supporting SHA1 certificates or the API endpoint ([#415](https://github.com/cloudflare/terraform-provider-cloudflare/issues/415))

**Deprecations:**

* `resource/cloudflare_zone_settings_override`: `tls_1_2_only` has been superseded by using `min_tls_version` instead ([#405](https://github.com/cloudflare/terraform-provider-cloudflare/issues/405))
* `resource/cloudflare_access_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_filter`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_firewall_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_load_balancer`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_page_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_rate_limit`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_waf_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_worker_route`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_worker_script`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
* `resource/cloudflare_zone_lockdown`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))

**Improvements:**

* **New Resource:** `cloudflare_custom_ssl` ([#418](https://github.com/cloudflare/terraform-provider-cloudflare/issues/418))
* `resource/cloudflare_filter`: Strip all surrounding whitespace from filter expressions to match API responses ([#361](https://github.com/cloudflare/terraform-provider-cloudflare/issues/361))
* `resource/cloudflare_zone`: Support unicode zone name values ([#412](https://github.com/cloudflare/terraform-provider-cloudflare/issues/412))
* `resource/cloudflare_page_rule`: Allow setting `origin_pull` for SSL ([#430](https://github.com/cloudflare/terraform-provider-cloudflare/issues/430))
* `resource/cloudflare_load_balancer_monitor`: Add TCP support for load balancer monitor ([#428](https://github.com/cloudflare/terraform-provider-cloudflare/issues/428))

**Fixes:**
* `resource/cloudflare_logpush_job`: Update documentation ([#395](https://github.com/cloudflare/terraform-provider-cloudflare/issues/395))
* `resource/cloudflare_zone_lockdown`: Fix: examples in documentation ([#407](https://github.com/cloudflare/terraform-provider-cloudflare/issues/407))
* `resource/cloudflare_page_rule`: Set nil on changed string-based Page Rule actions

## 1.16.1 (June 27, 2019)

**Fixes:**

* `resource/cloudflare_page_rule`: Fix regression in `browser_cache_ttl` where the value was sent as a string instead of an integer to the remote ([#390](https://github.com/cloudflare/terraform-provider-cloudflare/issues/390))

## 1.16.0 (June 20, 2019)

**Improvements:**

* `resource/cloudflare_zone_settings_override`: Add support for `h2_prioritization` and `image_resizing` ([#381](https://github.com/cloudflare/terraform-provider-cloudflare/issues/381))
* `resource/cloudflare_load_balancer_pool`: Update IP range for tests to not use reserved ranges ([#369](https://github.com/cloudflare/terraform-provider-cloudflare/issues/369))

**Fixes:**

* `resource/cloudflare_page_rule`: Fix issues with `browser_cache_ttl` defaults and when value is `0` (for Enterprise users)   ([#379](https://github.com/cloudflare/terraform-provider-cloudflare/issues/379))

## 1.15.0 (May 24, 2019)

* The provider is now compatible with Terraform v0.12, while retaining compatibility with prior versions. ([#309](https://github.com/cloudflare/terraform-provider-cloudflare/issues/309))

## 1.14.0 (May 15, 2019)

**Improvements:**

* **New Resource:** `cloudflare_argo` Manage Argo features ([#304](https://github.com/cloudflare/terraform-provider-cloudflare/issues/304))
* `cloudflare_zone`: Support management of partial zones ([#303](https://github.com/cloudflare/terraform-provider-cloudflare/issues/303))
* `cloudflare_rate_limit`: Update `modes` documentation ([#293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/212))
* `cloudflare_load_balancer`: Allow steering policy of "random" ([#329](https://github.com/cloudflare/terraform-provider-cloudflare/issues/329))

**Fixes:**

* `cloudflare_page_rule` - Allow setting `browser_cache_ttl` to 0 ([#293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/291))
* `cloudflare_page_rule` - Swap to completely replacing rules ([#338](https://github.com/cloudflare/terraform-provider-cloudflare/issues/338))

## 1.13.0 (April 12, 2019)

**Improvements**

* **New Resource:** `cloudflare_logpush_job` ([#287](https://github.com/cloudflare/terraform-provider-cloudflare/issues/287))
* `cloudflare_zone_settings` - Remove option to toggle `always_on_ddos` ([#253](https://github.com/cloudflare/terraform-provider-cloudflare/issues/253))
* `cloudflare_page_rule` - Update documentation to clarify "0" usage
* `cloudflare_zones` - Return zone ID and zone name ([#275](https://github.com/cloudflare/terraform-provider-cloudflare/issues/275))
* `cloudflare_load_balancer` - Add `enabled` field ([#208](https://github.com/cloudflare/terraform-provider-cloudflare/issues/208))
* `cloudflare_record` - validators: Allow PTR DNS records ([#283](https://github.com/cloudflare/terraform-provider-cloudflare/issues/283))

**Fixes:**

* `cloudflare_custom_pages` - Use correct casing for `zone_id` lookups
* `cloudflare_rate_limit` - Make `correlate` optional and not flap in state management ([#271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/271))
* `cloudflare_spectrum_application` - Fixed integration tests to work ([#275](https://github.com/cloudflare/terraform-provider-cloudflare/issues/275))
* `cloudflare_page_rule` - Better track field changes in `actions` resource. ([#107](https://github.com/cloudflare/terraform-provider-cloudflare/issues/107))

## 1.12.0 (March 07, 2019)

**Improvements:**

* provider: Enable request/response logging ([#212](https://github.com/cloudflare/terraform-provider-cloudflare/issues/212))
* resource/cloudflare_load_balancer_monitor: Add validation for `port` ([#213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/213))
* resource/cloudflare_load_balancer_monitor: Add `allow_insecure` and `follow_redirects` ([#205](https://github.com/cloudflare/terraform-provider-cloudflare/issues/205))
* resource/cloudflare_page_rule: Updated available actions documentation to match what is available ([#228](https://github.com/cloudflare/terraform-provider-cloudflare/issues/228))
* provider: Swap to using go modules for dependency management ([#230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/230))
* provider: Minimum Go version for development is now 1.11 ([#230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/230))

**Fixes:**

* resource/cloudflare_record: Read `data` back from API correctly ([#217](https://github.com/cloudflare/terraform-provider-cloudflare/issues/217))
* resource/cloudflare_rate_limit: Read `correlate` back from API correctly ([#204](https://github.com/cloudflare/terraform-provider-cloudflare/issues/204))
* resource/cloudflare_load_balancer_monitor: Fix incorrect type cast for `port` ([#213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/213))
* resource/cloudflare_load_balancer: Make `steering_policy` computed to avoid spurious diffs ([#214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/214))
* resource/cloudflare_load_balancer: Read `session_affinity` back from API to make import work & detects drifts ([#214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/214))

## 1.11.0 (January 11, 2019)

**Improvements:**
* **New Resource:** `cloudflare_spectrum_app` ([#156](https://github.com/cloudflare/terraform-provider-cloudflare/issues/156))
* **New Data Source:** `cloudflare_zones` ([#168](https://github.com/cloudflare/terraform-provider-cloudflare/issues/168))
* `cloudflare_load_balancer_monitor` - Add optional `port` parameter ([#179](https://github.com/cloudflare/terraform-provider-cloudflare/issues/179))
* `cloudflare_page_rule` - Improved documentation for `priority` attribute ([#182](https://github.com/cloudflare/terraform-provider-cloudflare/issues/182)], missing `explicit_cache_control` [[#185](https://github.com/cloudflare/terraform-provider-cloudflare/issues/185))
* `cloudflare_rate_limit` - Add `challenge` and `js_challenge` rate-limit modes ([#172](https://github.com/cloudflare/terraform-provider-cloudflare/issues/172))

**Fixes:**
* `cloudflare_page_rule` - Page rule `zone` attribute change to trigger new resource ([#183](https://github.com/cloudflare/terraform-provider-cloudflare/issues/183))

## 1.10.0 (December 18, 2018)

**Improvements:**
* `cloudflare_zone_settings_override` - Add `opportunistic_onion` zone setting support ([#170](https://github.com/cloudflare/terraform-provider-cloudflare/issues/170))
* `cloudflare_zone` - Add ability to set zone plan ([#160](https://github.com/cloudflare/terraform-provider-cloudflare/issues/160))

**Fixes:**
* `cloudflare_zone` - Allow zones to be properly imported ([#157](https://github.com/cloudflare/terraform-provider-cloudflare/issues/157))
* `cloudflare_access_policy` - Match access_policy argument requisites with reality ([#158](https://github.com/cloudflare/terraform-provider-cloudflare/issues/158))
* `cloudflare_filter` - Allow `zone_id` to set `zone` and vice versa ([#162](https://github.com/cloudflare/terraform-provider-cloudflare/issues/162))
* `cloudflare_firewall_rule` - Allow `zone_id` to set `zone` and vice versa ([#174](https://github.com/cloudflare/terraform-provider-cloudflare/issues/174))
* `cloudflare_access_rule` - Ensure `zone` and `zone_id` are always set ([#175](https://github.com/cloudflare/terraform-provider-cloudflare/issues/175))
* Minor documentation fixes

## 1.9.0 (November 15, 2018)

**Improvements:**
* **New Resource:** `cloudflare_access_application` ([#145](https://github.com/cloudflare/terraform-provider-cloudflare/issues/145))
* **New Resource:** `cloudflare_access_policy` ([#145](https://github.com/cloudflare/terraform-provider-cloudflare/issues/145))
* `cloudflare_load_balancer` - Add steering policy support ([#147](https://github.com/cloudflare/terraform-provider-cloudflare/issues/147))
* `cloudflare_load_balancer` - Support `session_affinity` ([#153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/153))
* `cloudflare_load_balancer_pool` - Support `weight` ([#153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/153))

**Fixes:**
* `cloudflare_record` - Compare name without the zone name ([#151](https://github.com/cloudflare/terraform-provider-cloudflare/issues/151))
* Minor documentation fixes ([#149](https://github.com/cloudflare/terraform-provider-cloudflare/issues/149)] [[#152](https://github.com/cloudflare/terraform-provider-cloudflare/issues/152))

## 1.8.0 (November 05, 2018)

**Improvements:**
* **New Resource:** `cloudflare_zone` ([#58](https://github.com/cloudflare/terraform-provider-cloudflare/issues/58))
* **New Resource:** `cloudflare_custom_pages` ([#132](https://github.com/cloudflare/terraform-provider-cloudflare/issues/132))
* `cloudflare_zone_settings_override` - Allow setting SSL level to Strict (SSL-Only Origin Pull) ([#122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/122))
* Update provider usage/build docs and how to update a dependency ([#138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/138))
* Improve `Building The Provider` instructions ([#143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/143))
* `cloudflare_access_rule` - Make importable for all rule types ([#141](https://github.com/cloudflare/terraform-provider-cloudflare/issues/141))
* `cloudflare_load_balancer_pool` - Implement `Update` ([#140](https://github.com/cloudflare/terraform-provider-cloudflare/issues/140))

**Fixes:**
* `cloudflare_rate_limit` - Documentation fixes for markdown where \_ALL\_ is italicized ([#125](https://github.com/cloudflare/terraform-provider-cloudflare/issues/125))
* `cloudflare_worker_route` - Correctly set `multi_script` on Enterprise worker imports ([#124](https://github.com/cloudflare/terraform-provider-cloudflare/issues/124))
* `account_member` - Ignore role ID ordering ([#128](https://github.com/cloudflare/terraform-provider-cloudflare/issues/128))
* `cloudflare_rate_limit` - Origin traffic isn't default anymore ([#130](https://github.com/cloudflare/terraform-provider-cloudflare/issues/130))
* `cloudflare_rate_limit` - Update rate limit validation to allow `1` ([#129](https://github.com/cloudflare/terraform-provider-cloudflare/issues/129))
* `cloudflare_record` - Add validation to ensure TTL is not set while `proxied` is true ([#127](https://github.com/cloudflare/terraform-provider-cloudflare/issues/127))
* Updated code for provider version in User-Agent
* `cloudflare_zone_lockdown` - Fix import of zone lockdowns ([#135](https://github.com/cloudflare/terraform-provider-cloudflare/issues/135))

## 1.7.0 (October 09, 2018)

**Improvements:**
* **New Resource:** `cloudflare_account_member` ([#78](https://github.com/cloudflare/terraform-provider-cloudflare/issues/78))

## 1.6.0 (October 05, 2018)

**Improvements:**
* **New Resource:** `cloudflare_filter`
* **New Resource:** `cloudflare_firewall_rule`

## 1.5.0 (September 21, 2018)

**Improvements:**
* **New Resource:** `cloudflare_zone_lockdown` ([#115](https://github.com/cloudflare/terraform-provider-cloudflare/issues/115))

**Fixes:**
* Send User-Agent header with name and version when contacting API
* `cloudflare_page_rule` - Fix page rule polish (off, lossless or lossy) ([#116](https://github.com/cloudflare/terraform-provider-cloudflare/issues/116))

## 1.4.0 (September 11, 2018)

**Improvements:**
* **New Resource:** `cloudflare_worker_route` ([#110](https://github.com/cloudflare/terraform-provider-cloudflare/issues/110))
* **New Resource:** `cloudflare_worker_script` ([#110](https://github.com/cloudflare/terraform-provider-cloudflare/issues/110))

## 1.3.0 (September 04, 2018)

**Improvements:**
* **New Resource:** `cloudflare_access_rule` ([#64](https://github.com/cloudflare/terraform-provider-cloudflare/issues/64))

**Fixes:**
* `cloudflare_zone_settings_override` -  Change Zone Settings Override to use GetOkExists ([#107](https://github.com/cloudflare/terraform-provider-cloudflare/issues/107))

## 1.2.0 (August 13, 2018)

**Improvements:**
* **New Resource:** `cloudflare_waf_rule` ([#98](https://github.com/cloudflare/terraform-provider-cloudflare/issues/98))
* `cloudflare_zone_settings_override` - Add `off` as Security Level setting ([#99](https://github.com/cloudflare/terraform-provider-cloudflare/issues/99))
* `resource_cloudflare_rate_limit` - Add nat support ([#96](https://github.com/cloudflare/terraform-provider-cloudflare/issues/96))
* `resource_cloudflare_zone_settings_override` - Add `zrt` as a value for the `tls_1_3` setting ([#106](https://github.com/cloudflare/terraform-provider-cloudflare/issues/106))
* Minor documentation improvements

**Fixes:**
* `cloudflare_record` - Setting a DNS record's `proxied` flag to false stopped working ([#103](https://github.com/cloudflare/terraform-provider-cloudflare/issues/103))

## 1.1.0 (July 25, 2018)

FIXES:

* `cloudflare_ip_ranges` - IPv6 CIDR blocks should return IPv6 addresses ([#51](https://github.com/cloudflare/terraform-provider-cloudflare/issues/51))
* `cloudflare_zone_settings_override` - Allow `0` for `browser_cache_ttl` ([#71](https://github.com/cloudflare/terraform-provider-cloudflare/issues/71))
* `cloudflare_page_rule` - `forwarding_urls` in page rules are lists ([#79](https://github.com/cloudflare/terraform-provider-cloudflare/issues/79))
* `cloudflare_page_rule` - The API supports `active` and `disabled`, not `paused` ([#84](https://github.com/cloudflare/terraform-provider-cloudflare/issues/84))

IMPROVEMENTS:
* `cloudflare_zone_settings_override` - Add support for `min_tls_version` ([#72](https://github.com/cloudflare/terraform-provider-cloudflare/issues/72))
* `cloudflare_page_rule` - Add support for more settings: `bypass_cache_on_cookie`, `cache_by_device_type`, `cache_deception_armor`, `cache_on_cookie`, `host_header_override`, `polish`, `explicit_cache_control`, `origin_error_page_pass_thru`, `sort_query_string_for_cache`, `resolve_override`, `respect_strong_etag`, `response_buffering`, `true_client_ip_header`, `mirage`, `disable_railgun`, `cache_key`, `waf`, `rocket_loader`, `cname_flattening` ([#68](https://github.com/cloudflare/terraform-provider-cloudflare/issues/68)], [[#81](https://github.com/cloudflare/terraform-provider-cloudflare/issues/81)], [[#85](https://github.com/cloudflare/terraform-provider-cloudflare/issues/85))
* `cloudflare_page_rule` - Add `off` setting to `security_level` ([#81](https://github.com/cloudflare/terraform-provider-cloudflare/issues/81))
* `cloudflare_record` - DNS Record improvements ([#97](https://github.com/cloudflare/terraform-provider-cloudflare/issues/97))
* Various documentation improvements

## 1.0.0 (April 06, 2018)

BACKWARDS INCOMPATIBILITIES / NOTES:

* resource/cloudflare_record: Changing `name` or `domain` now force a recreation
  of the record ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))

FEATURES:

* **New Resource:** `cloudflare_rate_limit` ([#30](https://github.com/cloudflare/terraform-provider-cloudflare/issues/30))
* **New Resource:** `cloudflare_page_rule` ([#38](https://github.com/cloudflare/terraform-provider-cloudflare/issues/38))
* **New Resource:** `cloudflare_load_balancer` ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))
* **New Resource:** `cloudflare_load_balancer_pool` ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))
* **New Resource:** `cloudflare_zone_settings_override` ([#41](https://github.com/cloudflare/terraform-provider-cloudflare/issues/41))
* **New Resource:** `cloudflare_load_balancer_monitor` ([#42](https://github.com/cloudflare/terraform-provider-cloudflare/issues/42))
* **New Data Source:** `cloudflare_ip_ranges` ([#28](https://github.com/cloudflare/terraform-provider-cloudflare/issues/28))

IMPROVEMENTS:

* resource/cloudflare_record: Validate `TXT` records ([#14](https://github.com/cloudflare/terraform-provider-cloudflare/issues/14))
* resource/cloudflare_record: Add `data` input to suppport SRV, LOC records
  ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))
* resource/cloudflare_record: Add computed attributes `created_on`,
  `modified_on`, `proxiable`, and `metadata` to records ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))
* resource/cloudflare_record: Support import of existing records ([#36](https://github.com/cloudflare/terraform-provider-cloudflare/issues/36))
* New Provider configuration options for API rate limiting ([#43](https://github.com/cloudflare/terraform-provider-cloudflare/issues/43))
* New Provider configuration options for using Organizations ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))

## 0.1.0 (June 20, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider
  Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
