## 4.33.0 (Unreleased)

## 4.32.0 (May 8th, 2024)

NOTES:

* resource/cloudflare_rate_limit: This resource is being deprecated in favor of the cloudflare_rulesets resource ([#3279](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3279))

ENHANCEMENTS:

* resource/cloudflare_access_application: add support for SCIM provisioning configuration ([#3291](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3291))
* resource/cloudflare_access_group: Add the option for email_list to be used in require, include and exclude fields ([#3247](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3247))
* resource/cloudflare_device_posture_rules: added support for os_version_extra ([#3281](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3281))

BUG FIXES:

* resource/cloudflare_turnstile: Fix error handling corrupting state ([#3284](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3284))

DEPENDENCIES:

* provider: bump github.com/cloudflare/cloudflare-go from 0.94.0 to 0.95.0 ([#3294](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3294))
* provider: bump github.com/hashicorp/terraform-plugin-go from 0.22.2 to 0.23.0 ([#3289](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3289))
* provider: bump golang.org/x/net from 0.24.0 to 0.25.0 ([#3290](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3290))
* provider: bump golangci/golangci-lint-action from 5 to 6 ([#3293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3293))

## 4.31.0 (May 1st, 2024)

ENHANCEMENTS:

* resource/cloudflare_access_application: added support for options_preflight_bypass ([#3267](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3267))
* resource/cloudflare_dlp_profile: Added support for `ocr_enabled` field to profiles ([#3224](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3224))
* resource/cloudflare_notification_policy: add 'target_ip' atrribute to 'filter' nested block ([#3263](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3263))
* resource/cloudflare_teams_account: add `custom_certificate` setting support ([#3253](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3253))
* resource/cloudflare_teams_location: added `ecs_support` field ([#3264](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3264))

BUG FIXES:

* resource/cloudflare_hyperdrive_config: Fix 'HyperdriveID' not included in Update call ([#3251](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3251))
* resource/cloudflare_managed_headers: disable header if it is deleted from terraform state ([#3260](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3260))
* resource/cloudflare_worker_script: fix namespaced script delete trying to delete from account rather than the namespace ([#3238](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3238))

INTERNAL:

* provider: introduce a muxed client to support using cloudflare-go/v0 and cloudflare-go/v2 together ([#3262](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3262))

DEPENDENCIES:

* provider: bump github.com/cloudflare/cloudflare-go from 0.93.0 to 0.94.0 ([#3265](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3265))
* provider: bump github.com/cloudflare/cloudflare-go/v2 from 2.0.0 to 2.1.0 ([#3274](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3274))
* provider: bump github.com/hashicorp/terraform-plugin-framework from 1.5.0 to 1.8.0 ([#3255](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3255))
* provider: bump github.com/hashicorp/terraform-plugin-go from 0.21.0 to 0.22.2 ([#3254](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3254))
* provider: bump golang.org/x/net from 0.19.0 to 0.23.0 in /tools ([#3258](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3258))
* provider: bump golangci/golangci-lint-action from 4 to 5 ([#3271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3271))

## 4.30.0 (April 17th, 2024)

ENHANCEMENTS:

* cloudflare/resource_logpush_job: Add support for `page_shield_events` ([#3237](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3237))
* resource/cloudflare_access_group: added support for common_names rule list type to allow for more than one common_name rule in a policy block ([#3229](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3229))
* resource/cloudflare_access_policy: added support for common_names rule list type to allow for more than one common_name rule in a policy block ([#3229](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3229))
* resource/cloudflare_ipsec_tunnel: added support for replay_protection ([#3249](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3249))

BUG FIXES:

* resource/cloudflare_email_routing_address: Make sure schema is correctly upgraded. ([#3245](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3245))

DEPENDENCIES:

* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.27.10 to 1.27.11 ([#3232](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3232))
* provider: bump `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.10 to 1.17.11 ([#3232](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3232))
* provider: bump github.com/cloudflare/cloudflare-go from 0.92.0 to 0.93.0 ([#3239](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3239))
* provider: bump golang.org/x/net from 0.22.0 to 0.23.0 ([#3225](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3225))
* provider: bump golang.org/x/net from 0.23.0 to 0.24.0 ([#3230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3230))

## 4.29.0 (April 3rd, 2024)

BREAKING CHANGES:

* data_source/record: Remove `locked` flag which is always false ([#3220](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3220))

ENHANCEMENTS:

* datasource/cloudflare_tunnel: Add the option to filter deleted tunnels ([#3201](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3201))
* resource/cloudflare_teams_rule: Add support for resolver policies ([#3198](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3198))

DEPENDENCIES:

* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.27.9 to 1.27.10 ([#3222](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3222))
* provider: bump `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.9 to 1.17.10 ([#3222](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3222))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.53.0 to 1.53.1 ([#3222](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3222))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.26.0 to 1.26.1 ([#3222](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3222))

## 4.28.0 (March 28th, 2024)

ENHANCEMENTS:

* resource/cloudflare_access_application: adds saml_attribute_transform_jsonata` to SaaS applications ([#3187](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3187))
* resource/cloudflare_device_posture_rule: update support for new fields for crowdstrike_s2s posture rule. ([#3216](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3216))
* resource/cloudflare_ipsec_tunnel: Adds IPsec tunnel health_check_direction & health_check_rate parameters ([#3112](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3112))

DEPENDENCIES:

* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.27.8 to 1.27.9 ([#3207](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3207))
* provider: bump `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.8 to 1.17.9 ([#3207](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3207))
* provider: bump github.com/cloudflare/cloudflare-go from 0.90.0 to 0.91.0 ([#3208](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3208))
* provider: bump github.com/cloudflare/cloudflare-go from 0.91.0 to 0.92.0 ([#3218](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3218))

## 4.27.0 (March 20th, 2024)

FEATURES:

* **New Resource:** `cloudflare_access_mutual_tls_hostname_settings` ([#3173](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3173))
* **New Resource:** `cloudflare_hyperdrive_config` ([#3111](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3111))

ENHANCEMENTS:

* resource/cloudflare_dlp_profile: Added support for `context_awareness` field to profiles ([#3158](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3158))
* resource/cloudflare_logpush_job: Add `output_options` parameter ([#3171](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3171))
* resource/cloudflare_notification_policy: Implement the `airport_code` filter ([#3183](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3183))
* resource/cloudflare_worker_script: Add `dispatch_namespace` to support uploading to a Workers for Platforms namespace ([#3154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3154))
* resource/cloudflare_worker_script: Add `tags` to support tagging Workers for Platforms Workers ([#3154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3154))

BUG FIXES:

* resource/cloudflare_access_application: Add Sensitive to oidc client_secret and preserve client_secret across apply ([#3168](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3168))
* resource/cloudflare_list_item: fix id parsing for imports ([#3191](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3191))
* resource/cloudflare_logpush_job: only set the value in state when it is defined ([#3188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3188))

DEPENDENCIES:

* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.27.6 to 1.27.7 ([#3172](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3172))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.27.7 to 1.27.8 ([#3197](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3197))
* provider: bump `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.6 to 1.17.7 ([#3172](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3172))
* provider: bump `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.7 to 1.17.8 ([#3197](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3197))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.51.3 to 1.51.4 ([#3172](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3172))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.51.4 to 1.52.0 ([#3182](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3182))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.52.0 to 1.52.1 ([#3190](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3190))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.52.1 to 1.53.0 ([#3197](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3197))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.25.2 to 1.25.3 ([#3172](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3172))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.25.3 to 1.26.0 ([#3197](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3197))
* provider: bump github.com/cloudflare/cloudflare-go from 0.89.0 to 0.90.0 ([#3178](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3178))
* provider: bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /tools ([#3180](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3180))
* provider: bump google.golang.org/protobuf from 1.32.0 to 1.33.0 ([#3181](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3181))

## 4.26.0 (March 6th, 2024)

FEATURES:

* **New Data Source:** `cloudflare_dlp_datasets` ([#3135](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3135))

ENHANCEMENTS:

* resource/cloudflare_access_application: adds `name_id_transform_jsonata` to SaaS applications ([#3132](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3132))

BUG FIXES:

* resource/cloudflare_access_application: Fix issue with sending allow_authenticate_via_warp on updates when it is not provided ([#3140](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3140))

DEPENDENCIES:

* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.27.1 to 1.27.2 ([#3136](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3136))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.27.2 to 1.27.3 ([#3138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3138))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.27.3 to 1.27.4 ([#3141](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3141))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.27.4 to 1.27.5 ([#3159](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3159))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.27.5 to 1.27.6 ([#3161](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3161))
* provider: bump `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.1 to 1.17.2 ([#3136](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3136))
* provider: bump `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.2 to 1.17.3 ([#3138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3138))
* provider: bump `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.3 to 1.17.4 ([#3141](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3141))
* provider: bump `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.4 to 1.17.5 ([#3159](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3159))
* provider: bump `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.5 to 1.17.6 ([#3161](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3161))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.50.2 to 1.50.3 ([#3136](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3136))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.50.3 to 1.51.0 ([#3138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3138))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.51.0 to 1.51.1 ([#3141](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3141))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.51.1 to 1.51.2 ([#3159](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3159))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.51.2 to 1.51.3 ([#3161](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3161))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.25.0 to 1.25.1 ([#3136](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3136))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.25.1 to 1.25.2 ([#3141](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3141))
* provider: bump github.com/cloudflare/cloudflare-go from 0.88.0 to 0.89.0 ([#3148](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3148))
* provider: bump github.com/hashicorp/terraform-plugin-go from 0.21.0 to 0.22.0 ([#3139](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3139))
* provider: bump github.com/hashicorp/terraform-plugin-mux from 0.14.0 to 0.15.0 ([#3149](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3149))
* provider: bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.32.0 to 2.33.0 ([#3142](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3142))
* provider: bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.32.0 to 2.33.0 ([#3147](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3147))
* provider: bump github.com/hashicorp/terraform-plugin-testing from 1.6.0 to 1.7.0 ([#3162](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3162))
* provider: bump github.com/stretchr/testify from 1.8.4 to 1.9.0 ([#3157](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3157))
* provider: bump golang.org/x/net from 0.21.0 to 0.22.0 ([#3160](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3160))

## 4.25.0 (February 21st, 2024)

BREAKING CHANGES:

* resource/cloudflare_custom_pages: Removed the `always_online` variant. This page is never generated anymore, if a requested page is unavailable in the archive the error page that would have been shown if always online wasn't enabled is shown. ([#3117](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3117))

ENHANCEMENTS:

* resource/cloudflare_access_application: adds oidc saas application support ([#3133](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3133))
* resource/cloudflare_access_application: adds the ability to set allow_authenticate_via_warp. ([#3103](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3103))
* resource/cloudflare_access_organization: adds the ability to set allow_authenticate_via_warp and warp_auth_session_duration. ([#3103](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3103))
* resource/cloudflare_teams_account: Add support for extended e-mail matching ([#3089](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3089))
* resource/cloudflare_teams_accounts: Added notification settings to teams antivirus settings ([#3124](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3124))
* resource/pages_project: Add `build_caching` attribute ([#3110](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3110))

BUG FIXES:

* resource/cloudflare_email_routing_address: add schema migrator ([#3119](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3119))

DEPENDENCIES:

* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.26.6 to 1.27.0 ([#3118](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3118))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.27.0 to 1.27.1 ([#3134](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3134))
* provider: bump `github.com/aws/aws-sdk-go-v2/credentials` from 1.16.16 to 1.17.0 ([#3118](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3118))
* provider: bump `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.0 to 1.17.1 ([#3134](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3134))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.48.1 to 1.49.0 ([#3118](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3118))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.49.0 to 1.50.0 ([#3125](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3125))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.50.0 to 1.50.1 ([#3128](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3128))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.50.1 to 1.50.2 ([#3134](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3134))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.24.1 to 1.25.0 ([#3118](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3118))
* provider: bump github.com/cloudflare/cloudflare-go from 0.87.0 to 0.88.0 ([#3122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3122))
* provider: bump golang.org/x/net from 0.20.0 to 0.21.0 ([#3108](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3108))
* provider: bump golangci/golangci-lint-action from 3 to 4 ([#3115](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3115))

## 4.24.0 (February 7th, 2023)

ENHANCEMENTS:

* datasource/cloudflare_record: Add the option to filter by "content" ([#3084](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3084))

BUG FIXES:

* resource/cloudflare_access_application: leave existence error handling checks to the `Read` operation when performing imports. ([#3075](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3075))
* resource/cloudflare_device_settings_policy: updated docs that `auto_connect` is in seconds, not in minutes ([#3080](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3080))
* resource/cloudflare_dlp_profile: fixed plan flapping with DLP custom entries ([#3090](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3090))
* resource/email_routing_rule: add schema migration for upgrading 4.22.0 to 4.23.0 ([#3102](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3102))

DEPENDENCIES:

* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.48.0 to 1.48.1 ([#3078](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3078))
* provider: bump github.com/cloudflare/cloudflare-go from 0.86.0 to 0.87.0 ([#3095](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3095))
* provider: bump github.com/google/uuid from 1.5.0 to 1.6.0 ([#3076](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3076))
* provider: bump github.com/hashicorp/terraform-plugin-go from 0.20.0 to 0.21.0 ([#3081](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3081))
* provider: bump github.com/hashicorp/terraform-plugin-mux from 0.13.0 to 0.14.0 ([#3085](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3085))
* provider: bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.31.0 to 2.32.0 ([#3086](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3086))
* provider: bump peter-evans/create-or-update-comment from 3 to 4 ([#3079](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3079))

## 4.23.0 (January 24th, 2023)

BREAKING CHANGES:

* resource/cloudflare_list_item: `include_subdomains` is now a boolean value. If you previously set it to `"enabled"`, you should update your configuration to use `true` instead or if you set it to "`disabled`", you should update it to `false`. The rest will be handled by the internal state migrator. ([#3026](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3026))
* resource/cloudflare_list_item: `preserve_path_suffix` is now a boolean value. If you previously set it to `"enabled"`, you should update your configuration to use `true` instead or if you set it to "`disabled`", you should update it to `false`. The rest will be handled by the internal state migrator. ([#3026](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3026))
* resource/cloudflare_list_item: `preserve_query_string` is now a boolean value. If you previously set it to `"enabled"`, you should update your configuration to use `true` instead or if you set it to "`disabled`", you should update it to `false`. The rest will be handled by the internal state migrator. ([#3026](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3026))
* resource/cloudflare_list_item: `subpath_matching` is now a boolean value. If you previously set it to `"enabled"`, you should update your configuration to use `true` instead or if you set it to "`disabled`", you should update it to `false`. The rest will be handled by the internal state migrator. ([#3026](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3026))

ENHANCEMENTS:

* resource/cloudflare_access_application: adds the ability to set default_relay_state on saas applications. ([#3053](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3053))
* resource/cloudflare_email_routing_address: add ability to import ([#2977](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2977))
* resource/cloudflare_email_routing_rule: add ability to import ([#2998](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2998))
* resource/cloudflare_notification_policy: Implement the `affected_components` option ([#3009](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3009))

INTERNAL:

* cloudflare_email_routing_rule: migrate to plugin framework ([#2998](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2998))
* resource/cloudflare_email_routing_address: migrate to framework provider ([#2977](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2977))
* resource/cloudflare_list_item: migrate to plugin framework. Due to this migration, we are removing some workaround field values that were previously in place to account for the known zero value issues in the underlying SDKv2. See the release notes for the end user facing changes that need to be made for the internal state migrator to handle the internals. ([#3026](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3026))

DEPENDENCIES:

* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.26.3 to 1.26.4 ([#3065](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3065))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.26.4 to 1.26.5 ([#3071](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3071))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.26.5 to 1.26.6 ([#3074](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3074))
* provider: bump actions/cache from 3 to 4 ([#3067](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3067))
* provider: bump github.com/cloudflare/cloudflare-go from 0.85.0 to 0.86.0 ([#3066](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3066))
* provider: bump github.com/hashicorp/terraform-plugin-framework from 1.4.2 to 1.5.0 ([#3058](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3058))

## 4.22.0 (January 10th, 2024)

FEATURES:

* **New Resource:** `cloudflare_worker_secret` ([#3035](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3035))

ENHANCEMENTS:

* resource/cloudflare_notification_policy: Add tunnel_id filter for tunnel_health_event policies ([#3038](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3038))
* resource/cloudflare_worker_script: adds D1 binding support ([#2960](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2960))

BUG FIXES:

* cloudflare_notification_policy: revert `ExactlyOneOf` ([#3032](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3032))
* resource/cloudflare_dlp_profile: Prevent misidentified changes in dlp resources ([#3044](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3044))
* resource/cloudflare_teams_rule: changed type & validation on the notification settings url ([#3030](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3030))
* resource/cloudflare_teams_rules: fix block_page_enabled behaviour ([#3010](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3010))
* resource/cloudflare_turnstile_widget: Support empty list of domains ([#3046](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3046))

DEPENDENCIES:

* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.26.2 to 1.26.3 ([#3042](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3042))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.47.7 to 1.47.8 ([#3042](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3042))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.47.8 to 1.48.0 ([#3043](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3043))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.24.0 to 1.24.1 ([#3042](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3042))
* provider: bump github.com/cloudflare/circl from 1.3.3 to 1.3.7 ([#3047](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3047))
* provider: bump github.com/cloudflare/circl from 1.3.3 to 1.3.7 in /tools ([#3048](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3048))
* provider: bump github.com/cloudflare/cloudflare-go from 0.84.0 to 0.85.0 ([#3034](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3034))
* provider: bump github.com/go-git/go-git/v5 from 5.4.2 to 5.11.0 in /tools ([#3029](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3029))
* provider: bump golang.org/x/net from 0.19.0 to 0.20.0 ([#3050](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3050))

## 4.21.0 (December 27th, 2023)

ENHANCEMENTS:

* resource/cloudflare_access_application: adds the ability to set customization fields on the app launcher application. ([#2777](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2777))
* resource/cloudflare_access_organization: remove default value for `session_duration`. ([#2995](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2995))
* resource/cloudflare_access_policy: remove default value for `session_duration`. ([#2995](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2995))
* resource/cloudflare_device_posture_integration: add support for `access_client_id` and `access_client_secret` fields ([#3013](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3013))
* resource/cloudflare_logpush_job: add support for `magic_ids_detections`. ([#2983](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2983))
* resource/cloudflare_notification_policy: enable `selector` filter and add `traffic_anomalies_alert` as a policy alert type ([#2976](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2976))
* resource/cloudflare_pages_project: support `standard` usage model for functions ([#2963](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2963))
* resource/cloudflare_tunnel_config: Destroying tunnel configurations now applies an empty configuration rather than deleting the parent cloudflare_tunnel resource ([#2769](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2769))

BUG FIXES:

* resource/cloudflare_list_item: fix issue preventing usage of redirect item type ([#2975](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2975))

DEPENDENCIES:

* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.25.10 to 1.25.11 ([#2973](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2973))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.25.11 to 1.25.12 ([#2987](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2987))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.25.12 to 1.26.0 ([#2993](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2993))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.25.12 to 1.26.0 ([#2993](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2993))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.25.5 to 1.25.8 ([#2968](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2968))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.25.8 to 1.25.9 ([#2969](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2969))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.25.9 to 1.25.10 ([#2971](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2971))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.26.0 to 1.26.1 ([#2997](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2997))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.26.1 to 1.26.2 ([#3022](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3022))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.44.0 to 1.46.0 ([#2968](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2968))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.46.0 to 1.47.0 ([#2969](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2969))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.47.0 to 1.47.1 ([#2971](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2971))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.47.1 to 1.47.2 ([#2973](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2973))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.47.2 to 1.47.3 ([#2987](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2987))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.47.3 to 1.47.4 ([#2993](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2993))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.47.3 to 1.47.4 ([#2993](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2993))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.47.4 to 1.47.5 ([#2997](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2997))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.47.5 to 1.47.6 ([#3016](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3016))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.47.6 to 1.47.7 ([#3022](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3022))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.23.1 to 1.23.2 ([#2968](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2968))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.23.2 to 1.23.3 ([#2969](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2969))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.23.3 to 1.23.4 ([#2971](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2971))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.23.4 to 1.23.5 ([#2973](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2973))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.23.5 to 1.24.0 ([#2993](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2993))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.23.5 to 1.24.0 ([#2993](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2993))
* provider: bump actions/setup-go from 4 to 5 ([#2989](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2989))
* provider: bump actions/stale from 8 to 9 ([#2992](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2992))
* provider: bump github.com/cloudflare/cloudflare-go from 0.82.0 to 0.83.0 ([#2988](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2988))
* provider: bump github.com/cloudflare/cloudflare-go from 0.83.0 to 0.84.0 ([#3019](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3019))
* provider: bump github.com/google/uuid from 1.4.0 to 1.5.0 ([#3002](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3002))
* provider: bump github.com/hashicorp/terraform-plugin-mux from 0.12.0 to 0.13.0 ([#3006](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3006))
* provider: bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.30.0 to 2.31.0 ([#3007](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3007))
* provider: bump github.com/hashicorp/terraform-plugin-testing from 1.5.1 to 1.6.0 ([#2984](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2984))
* provider: bump github/codeql-action from 2 to 3 ([#3005](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3005))
* provider: bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /tools ([#3015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3015))
* provider: bump golang.org/x/crypto from 0.16.0 to 0.17.0 ([#3017](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3017))
* resource/cloudflare_teams_rule: Added support for notification settings at teams rule ([#3021](https://github.com/cloudflare/terraform-provider-cloudflare/issues/3021))

## 4.20.0 (November 29th, 2023)

FEATURES:

* **New Data Source:** `cloudflare_origin_ca_certificate` ([#2961](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2961))

ENHANCEMENTS:

* resource/cloudflare_email_routing_rule: `action.value` is now optional to support `drop` rules not requiring it ([#2449](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2449))
* resource/cloudflare_email_routing_rule: add action type `drop` ([#2449](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2449))
* resource/cloudflare_notification_policy: add support for `brand_protection_alert` alert type ([#2937](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2937))
* resource/cloudflare_notification_policy: add support for `brand_protection_digest` alert type ([#2937](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2937))
* resource/cloudflare_notification_policy: add support for `logo_match_alert` alert type ([#2937](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2937))
* resource/cloudflare_notification_policy: add support for `magic_tunnel_health_check_event` alert type ([#2937](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2937))
* resource/cloudflare_notification_policy: add support for `maintenance_event_notification` alert type ([#2937](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2937))
* resource/cloudflare_notification_policy: add support for `mtls_certificate_store_certificate_expiration_type` alert type ([#2937](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2937))
* resource/cloudflare_notification_policy: add support for `radar_notification` alert type ([#2937](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2937))
* resource/cloudflare_ruleset: make rate limiting `requests_to_origin` optional with a default value of `false` to match the API behaviour ([#2954](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2954))

BUG FIXES:

* resource/cloudflare_list_item: fix list_item for `asn` and `hostname` types ([#2951](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2951))
* resource/cloudflare_notification_policy: Fix missing new_status filter required by tunnel_health_event policies ([#2390](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2390))

DEPENDENCIES:

* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.25.1 to 1.25.3 ([#2948](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2948))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.25.3 to 1.25.4 ([#2953](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2953))
* provider: bump `github.com/aws/aws-sdk-go-v2/config` from 1.25.4 to 1.25.5 ([#2956](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2956))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.42.2 to 1.43.0 ([#2948](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2948))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.43.0 to 1.43.1 ([#2953](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2953))
* provider: bump `github.com/aws/aws-sdk-go-v2/service/s3` from 1.43.1 to 1.44.0 ([#2956](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2956))
* provider: bump `github.com/aws/aws-sdk-go-v2` from 1.23.0 to 1.23.1 ([#2953](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2953))
* provider: bump github.com/cloudflare/cloudflare-go from 0.81.0 to 0.82.0 ([#2957](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2957))
* provider: bump github.com/hashicorp/terraform-plugin-go from 0.19.0 to 0.19.1 ([#2942](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2942))
* provider: bump golang.org/x/net from 0.18.0 to 0.19.0 ([#2967](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2967))
* provider: updates `github.com/aws/aws-sdk-go-v2/config` from 1.24.0 to 1.25.1 ([#2945](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2945))
* provider: updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.42.1 to 1.42.2 ([#2945](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2945))
* provider: updates `github.com/aws/aws-sdk-go-v2` from 1.22.2 to 1.23.0 ([#2945](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2945))

## 4.19.0 (15th November, 2023)

NOTES:

* resource/cloudflare_argo: `tiered_caching` attribute is deprecated in favour of the dedicated `cloudflare_tiered_cache` resource. ([#2906](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2906))

FEATURES:

* **New Resource:** `cloudflare_keyless_certificate` ([#2779](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2779))

ENHANCEMENTS:

* resource/cloudflare_notification_policy: Add support for `incident_alert` type ([#2901](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2901))
* resource/cloudflare_zone: add support for `secondary` zone types ([#2939](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2939))

BUG FIXES:

* resource/cloudflare_list_item: ensure each `item` has its own ID and is not based on the latest created entry ([#2922](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2922))

INTERNAL:

* provider: prevent new resources and datasources from being created with `terraform-plugin-sdk` ([#2871](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2871))

DEPENDENCIES:

* provider: bumps github.com/aws/aws-sdk-go-v2 from 1.21.2 to 1.22.0 ([#2899](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2899))
* provider: bumps github.com/aws/aws-sdk-go-v2 from 1.22.0 to 1.22.1 ([#2904](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2904))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.19.1 to 1.20.0 ([#2898](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2898))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.20.0 to 1.21.0 ([#2902](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2902))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.21.0 to 1.22.0 ([#2908](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2908))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.22.0 to 1.22.1 ([#2912](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2912))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.22.1 to 1.22.2 ([#2917](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2917))
* provider: bumps github.com/aws/aws-sdk-go-v2/service/s3 from 1.40.2 to 1.41.0 ([#2897](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2897))
* provider: bumps github.com/aws/aws-sdk-go-v2/service/s3 from 1.41.0 to 1.42.0 ([#2905](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2905))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.80.0 to 0.81.0 ([#2919](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2919))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.29.0 to 2.30.0 ([#2925](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2925))
* provider: bumps golang.org/x/net from 0.17.0 to 0.18.0 ([#2921](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2921))
* provider: updates `github.com/aws/aws-sdk-go-v2/config` from 1.22.2 to 1.23.0 ([#2931](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2931))
* provider: updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.42.0 to 1.42.1 ([#2931](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2931))
* provider: updates `github.com/aws/aws-sdk-go-v2` from 1.22.1 to 1.22.2 ([#2931](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2931))

## 4.18.0 (1st November, 2023)

FEATURES:

* **New Data Source:** `cloudflare_device_posture_rules` ([#2868](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2868))
* **New Data Source:** `cloudflare_tunnel` ([#2866](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2866))
* **New Data Source:** `cloudflare_tunnel_virtual_network` ([#2867](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2867))
* **New Resource:** `cloudflare_api_shield_operation_schema_validation_settings` ([#2852](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2852))
* **New Resource:** `cloudflare_api_shield_schema_validation_settings` ([#2841](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2841))

ENHANCEMENTS:

* resource/cloudflare_load_balancer: Add support for least_connections steering ([#2818](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2818))
* resource/cloudflare_load_balancer_pool: Add support for least_connections origin steering ([#2818](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2818))
* resource/cloudflare_logpush_job: add support for `casb_findings` dataset ([#2859](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2859))
* resource/cloudflare_teams_account: Add `non_identity_browser_isolation_enabled` field ([#2878](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2878))
* resource/cloudflare_teams_account: add support for `body_scanning` config ([#2887](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2887))
* resource/cloudflare_workers_script: add support for `placement` config ([#2893](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2893))

BUG FIXES:

* resource/cloudflare_observatory_scheduled_test: Add missing 'asia-south1' region ([#2891](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2891))
* resource/cloudflare_rulesets: Allow zero to not default to null for mitigation_timeout ([#2874](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2874))

DEPENDENCIES:

* ci: drop separate misspell installation ([#2814](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2814))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.19.0 to 1.19.1 ([#2877](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2877))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.79.0 to 0.80.0 ([#2883](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2883))
* provider: bumps github.com/google/uuid from 1.3.1 to 1.4.0 ([#2889](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2889))
* provider: bumps github.com/hashicorp/terraform-plugin-framework from 1.4.1 to 1.4.2 ([#2876](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2876))

## 4.17.0 (18th October, 2023)

FEATURES:

* **New Resource:** `cloudflare_access_tag` ([#2776](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2776))
* **New Resource:** `cloudflare_api_shield_schema` ([#2784](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2784))
* **New Resource:** `cloudflare_d1_database` ([#2850](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2850))
* **New Resource:** `cloudflare_observatory_scheduled_test` ([#2807](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2807))

ENHANCEMENTS:

* provider: allow defining a user agent operator suffix through the schema field (`user_agent_operator_suffix`) and via the environment variable (`CLOUDFLARE_USER_AGENT_OPERATOR_SUFFIX`) ([#2831](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2831))
* resource/cloudflare_access_application: Add idp_entity_id, public_key and sso_endpoint attributes to saas_app ([#2838](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2838))
* resource/cloudflare_access_application: adds the ability to associate a tag with an application. ([#2776](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2776))
* resource/cloudflare_access_organization: Add session_duration field ([#2857](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2857))
* resource/cloudflare_access_policy: Add session_duration field ([#2857](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2857))
* resource/cloudflare_ruleset: Add support for the use of Additional Cacheable Ports option in the Rulesets API ([#2854](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2854))
* resource/cloudflare_teams_accounts: Add support for setting ssh encryption key in ZT settings ([#2826](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2826))
* resource/cloudflare_zone_settings_override: Add support for `fonts` ([#2773](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2773))

BUG FIXES:

* resource/cloudflare_access_application: fix import of cloudflare_access_application not reading saas_app config ([#2843](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2843))
* resource/cloudflare_access_policy: Send purpose justification settings properly on updates ([#2836](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2836))
* resource/cloudflare_bot_management: fix fight mode not being sent to API ([#2833](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2833))
* resource/cloudflare_pages_project: Fix 'preview_branch_includes' always showing it has changes if not provided ([#2796](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2796))
* resource/cloudflare_ruleset: Add note that logging is only supported with the skip action ([#2851](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2851))

INTERNAL:

* provider: updated user agent string to now be `terraform-provider-cloudflare/<version> <plugin> <operator suffix>` ([#2831](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2831))

DEPENDENCIES:

* provider: bumps github.com/aws/aws-sdk-go-v2 from 1.21.0 to 1.21.1 ([#2820](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2820))
* provider: bumps github.com/aws/aws-sdk-go-v2 from 1.21.1 to 1.21.2 ([#2847](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2847))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.43 to 1.18.44 ([#2823](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2823))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.44 to 1.18.45 ([#2846](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2846))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.45 to 1.19.0 ([#2853](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2853))
* provider: bumps github.com/aws/aws-sdk-go-v2/credentials from 1.13.41 to 1.13.42 ([#2821](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2821))
* provider: bumps github.com/aws/aws-sdk-go-v2/service/s3 from 1.40.0 to 1.40.1 ([#2822](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2822))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.78.0 to 0.79.0 ([#2832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2832))
* provider: bumps github.com/google/go-cmp from 0.5.9 to 0.6.0 ([#2830](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2830))
* provider: bumps github.com/hashicorp/terraform-plugin-framework from 1.4.0 to 1.4.1 ([#2828](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2828))
* provider: bumps golang.org/x/net from 0.15.0 to 0.16.0 ([#2819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2819))
* provider: bumps golang.org/x/net from 0.16.0 to 0.17.0 ([#2829](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2829))
* provider: bumps golang.org/x/net from 0.7.0 to 0.17.0 ([#2837](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2837))

## 4.16.0 (4th October, 2023)

BREAKING CHANGES:

* resource/cloudflare_spectrum_application: Remove default values, make `edge_ips` parameter optional. ([#2629](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2629))

FEATURES:

* **New Resource:** `cloudflare_api_shield_operation` ([#2760](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2760))

ENHANCEMENTS:

* resource/cloudflare_authenticated_origin_pulls: Improve import, update documentation ([#2771](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2771))
* resource/cloudflare_notification_policy: Add `advanced_http_alert_error` alert_type ([#2789](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2789))
* resource/cloudflare_notification_policy: Implement the `group_by`, `where` and `actions` options ([#2789](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2789))
* resource/cloudflare_ruleset: Add support for cache bypass by default in Edge TTL modes ([#2764](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2764))

BUG FIXES:

* resource/cloudflare_access_identity_provider: Fix cloudflare_access_identity_provider incorrectly discards SCIM configuration secret ([#2744](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2744))
* resource/cloudflare_notification_policy: handle manually deleted policies by removing them from state ([#2791](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2791))
* resource/cloudflare_ruleset: ability to use exclude_origin=true in cache_key.custom_key.header without the need of specifying include or check_presence. ([#2802](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2802))
* resource/cloudflare_ruleset: mark `requests_to_origin` required for ratelimit blocks ([#2808](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2808))

DEPENDENCIES:

* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.40 to 1.18.41 ([#2781](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2781))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.41 to 1.18.42 ([#2792](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2792))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.42 to 1.18.43 ([#2811](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2811))
* provider: bumps github.com/aws/aws-sdk-go-v2/credentials from 1.13.39 to 1.13.40 ([#2793](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2793))
* provider: bumps github.com/aws/aws-sdk-go-v2/credentials from 1.13.40 to 1.13.41 ([#2810](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2810))
* provider: bumps github.com/aws/aws-sdk-go-v2/service/s3 from 1.38.5 to 1.39.0 ([#2782](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2782))
* provider: bumps github.com/aws/aws-sdk-go-v2/service/s3 from 1.39.0 to 1.40.0 ([#2795](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2795))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.77.0 to 0.78.0 ([#2797](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2797))

## 4.15.0 (20th September, 2023)

ENHANCEMENTS:

* resource/cloudflare_access_identity_provider: Support email_claim_name, Okta authorization_server_id, and pingone ([#2765](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2765))
* resource/cloudflare_ruleset: Add support for a new Browser Mode that allows bypass of downstream caches ([#2756](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2756))
* resource/cloudflare_ruleset: Add support for the use of Origin Cache Control in the Rulesets API ([#2753](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2753))
* resource/cloudflare_ruleset: Add support for the use of Proxy Read Timeout field in Rulesets API ([#2755](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2755))

BUG FIXES:

* resource/cloudflare_list: Fix import for cloudflare_list resource ([#2663](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2663))
* resource/cloudflare_record: Updates the cast to a pointer to match changes in the SDK ([#2763](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2763))
* resource/pages_project: force replace when changing pages source ([#2750](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2750))

DEPENDENCIES:

* provider: bumps crazy-max/ghaction-import-gpg from 5 to 6 ([#2758](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2758))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.39 to 1.18.40 ([#2775](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2775))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.76.0 to 0.77.0 ([#2761](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2761))
* provider: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.5 to 1.4.0 ([#2745](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2745))
* provider: bumps github.com/hashicorp/terraform-plugin-mux from 0.11.2 to 0.12.0 ([#2746](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2746))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.28.0 to 2.29.0 ([#2748](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2748))
* provider: bumps goreleaser/goreleaser-action from 4.6.0 to 5.0.0 ([#2757](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2757))

## 4.14.0 (6th September, 2023)

FEATURES:

* **New Resource:** `cloudflare_web_analytics_rule` ([#2686](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2686))
* **New Resource:** `cloudflare_web_analytics_site` ([#2686](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2686))

ENHANCEMENTS:

* resource/cloudflare_access_application: Add custom_non_identity_deny_url field ([#2721](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2721))
* resource/cloudflare_access_group: Improve documentation for access_group usage ([#2718](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2718))
* resource/cloudflare_load_balancer_monitor: add support for `consecutive_up` and `consecutive_down` ([#2723](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2723))
* resource/cloudflare_total_tls: add support for importing existing resources ([#2734](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2734))

BUG FIXES:

* resource/cloudflare_access_identity_provider: Fix access IDPs not importing config obj ([#2735](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2735))

DEPENDENCIES:

* provider: bumps actions/checkout from 3 to 4 ([#2736](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2736))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.36 to 1.18.37 ([#2714](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2714))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.37 to 1.18.38 ([#2731](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2731))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.38 to 1.18.39 ([#2741](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2741))
* provider: bumps github.com/aws/aws-sdk-go-v2/credentials from 1.13.35 to 1.13.36 ([#2732](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2732))
* provider: bumps github.com/aws/aws-sdk-go-v2/credentials from 1.13.36 to 1.13.37 ([#2740](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2740))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.75.0 to 0.76.0 ([#2726](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2726))
* provider: bumps github.com/hashicorp/terraform-plugin-framework-validators from 0.11.0 to 0.12.0 ([#2727](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2727))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.27.0 to 2.28.0 ([#2719](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2719))
* provider: bumps github.com/hashicorp/terraform-plugin-testing from 1.4.0 to 1.5.1 ([#2730](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2730))
* provider: bumps golang.org/x/net from 0.14.0 to 0.15.0 ([#2739](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2739))
* provider: bumps goreleaser/goreleaser-action from 4.4.0 to 4.6.0 ([#2742](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2742))

## 4.13.0 (23rd August, 2023)

FEATURES:

* **New Data Source:** `cloudflare_user` ([#2691](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2691))
* **New Resource:** `cloudflare_bot_management` ([#2672](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2672))
* **New Resource:** `cloudflare_hostname_tls_setting` ([#2700](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2700))
* **New Resource:** `cloudflare_hostname_tls_setting_ciphers` ([#2700](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2700))
* **New Resource:** `cloudflare_zone_hold` ([#2671](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2671))

ENHANCEMENTS:

* datasource/api_token_permission_groups: Add R2 scopes ([#2687](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2687))
* datasource/api_token_permission_groups: Convert to plugin framework ([#2687](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2687))
* resource/cloudflare_access_application: adds support for custom saml attributes in saas access apps ([#2676](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2676))
* resource/cloudflare_access_group: add support for AccessGroupAzureAuthContext ([#2654](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2654))
* resource/cloudflare_access_identity_provider: add conditional_access_enabled attr ([#2654](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2654))
* resource/cloudflare_access_service_token: add support for managing `Duration` ([#2647](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2647))
* resource/cloudflare_device_posture_integration: update support for managing `tanium_s2s` third party posture provider. ([#2674](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2674))
* resource/cloudflare_device_posture_rule: update support for new fields for tanium_s2s posture rule. ([#2674](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2674))
* resource/cloudflare_notification_policy: Add possibility to configure Pages Alerts. ([#2694](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2694))
* resource/cloudflare_waiting_room: Add `queueing_status_code` to the Waiting Room resource ([#2666](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2666))
* resource/cloudflare_worker_domain: add support for `Import` operations ([#2679](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2679))

BUG FIXES:

* resource/cloudflare_access_group: Fix issue where saml rules would not read the IDP id from the API ([#2683](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2683))
* resource/cloudflare_rulest: allow configuring an origin `Port` value without the `Host` (and vice versa) ([#2677](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2677))

DEPENDENCIES:

* provider: bumps github.com/aws/aws-sdk-go-v2 from 1.20.1 to 1.20.2 ([#2695](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2695))
* provider: bumps github.com/aws/aws-sdk-go-v2 from 1.20.3 to 1.21.0 ([#2710](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2710))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.33 to 1.18.34 ([#2697](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2697))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.34 to 1.18.35 ([#2706](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2706))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.35 to 1.18.36 ([#2708](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2708))
* provider: bumps github.com/aws/aws-sdk-go-v2/credentials from 1.13.32 to 1.13.33 ([#2696](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2696))
* provider: bumps github.com/aws/aws-sdk-go-v2/credentials from 1.13.33 to 1.13.34 ([#2703](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2703))
* provider: bumps github.com/aws/aws-sdk-go-v2/credentials from 1.13.34 to 1.13.35 ([#2709](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2709))
* provider: bumps github.com/aws/aws-sdk-go-v2/service/s3 from 1.38.2 to 1.38.3 ([#2698](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2698))
* provider: bumps github.com/aws/aws-sdk-go-v2/service/s3 from 1.38.3 to 1.38.4 ([#2705](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2705))
* provider: bumps github.com/aws/aws-sdk-go-v2/service/s3 from 1.38.4 to 1.38.5 ([#2707](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2707))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.74.0 to 0.75.0 ([#2685](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2685))
* provider: bumps github.com/google/uuid from 1.3.0 to 1.3.1 ([#2711](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2711))
* provider: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.4 to 1.3.5 ([#2699](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2699))
* provider: bumps goreleaser/goreleaser-action from 4.3.0 to 4.4.0 ([#2675](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2675))

## 4.12.0 (9th August, 2023)

BREAKING CHANGES:

* resource/cloudflare_ruleset: remove `shareable_entitlement_name` per the Go library changes since it hasn't ever been controllable by users ([#2652](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2652))

FEATURES:

* **New Data Source:** `cloudflare_zone_cache_reserve` ([#2642](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2642))
* **New Resource:** `cloudflare_access_custom_page` ([#2643](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2643))
* **New Resource:** `cloudflare_zone_cache_reserve` ([#2642](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2642))

ENHANCEMENTS:

* resource/cloudflare_access_application: adds the ability to associate a custom page with an application. ([#2643](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2643))
* resource/cloudflare_access_organization: adds the ability to associate a custom page with an organization. ([#2643](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2643))
* resource/cloudflare_notification_policy: Add support for `pages_event_alert` alert type ([#2602](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2602))
* resource/cloudflare_pages_project: Allow renaming projects without destroying and recreating ([#2602](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2602))
* resource/cloudflare_teams_account: Adds support for protocol detection feature ([#2625](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2625))
* resource/cloudflare_user_agent_blocking_rules: add support for importing resources ([#2640](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2640))

BUG FIXES:

* resource/cloudflare_custom_hostname: prevent infinite loop when `wait_for_ssl_pending_validation` is set if SSL status is already `active` ([#2638](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2638))
* resource/cloudflare_load_balancer: fix full deletion of pop_pools, region_pools, country_pools on update ([#2673](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2673))
* resource/cloudflare_load_balancer: handle inconsistent sorting bug in `schema.HashResource` resulting in resources incorrectly being updated when no changes have been made ([#2635](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2635))
* resource/cloudflare_pages_project: `deployment_configs` are now computed ([#2602](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2602))

DEPENDENCIES:

* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.29 to 1.18.32 ([#2651](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2651))
* provider: bumps github.com/aws/aws-sdk-go-v2/config from 1.18.32 to 1.18.33 ([#2670](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2670))
* provider: bumps github.com/aws/aws-sdk-go-v2/credentials from 1.13.28 to 1.13.31 ([#2648](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2648))
* provider: bumps github.com/aws/aws-sdk-go-v2/service/s3 from 1.37.0 to 1.38.1 ([#2650](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2650))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.73.0 to 0.74.0 ([#2652](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2652))
* provider: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.3 to 1.3.4 ([#2657](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2657))
* provider: bumps github.com/hashicorp/terraform-plugin-framework-validators from 0.10.0 to 0.11.0 ([#2658](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2658))
* provider: bumps golang.org/x/net from 0.12.0 to 0.13.0 ([#2646](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2646))
* provider: bumps golang.org/x/net from 0.13.0 to 0.14.0 ([#2661](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2661))

## 4.11.0 (26th July, 2023)

FEATURES:

* **New Resource:** `cloudflare_regional_tiered_cache` ([#2624](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2624))

ENHANCEMENTS:

* resource/cloudflare_device_posture_integration: add support for managing `sentinelone_s2s` third party posture provider. ([#2618](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2618))
* resource/cloudflare_device_posture_rule: add ability to create client_certificate and sentinelone_s2s posture rule ([#2618](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2618))
* resource/cloudflare_load_balancer: support header session affinity policy ([#2521](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2521))
* resource/record: Allow SVCB DNS record ([#2632](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2632))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.72.0 to 0.73.0 ([#2626](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2626))
* provider: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.2 to 1.3.3 ([#2627](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2627))
* provider: bumps github.com/hashicorp/terraform-plugin-mux from 0.11.1 to 0.11.2 ([#2616](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2616))
* provider: bumps github.com/hashicorp/terraform-plugin-testing from 1.3.0 to 1.4.0 ([#2631](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2631))

## 4.10.0 (12th July, 2023)

FEATURES:

* **New Data Source:** `clouflare_access_application` ([#2547](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2547))

ENHANCEMENTS:

* resource/cloudflare_access_ca_certificate: remove redundant `certificate_id` from `Import` requirements as it is never used ([#2547](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2547))
* resource/cloudflare_load_balancer_monitor: Add example import. ([#2572](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2572))

BUG FIXES:

* resource/cloudflare_load_balancer: fix import of load_balancer when rules included overrides or fixed_response ([#2571](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2571))
* resource/cloudflare_record: fix importing of DNSKEY record types ([#2568](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2568))
* resource/cloudflare_ruleset: Fix detection of conflicting entrypoint rulesets ([#2566](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2566))

DEPENDENCIES:

* provider: bumps dependabot/fetch-metadata from 1.5.1 to 1.6.0 ([#2557](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2557))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.70.0 to 0.72.0 ([#2584](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2584))
* provider: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.1 to 1.3.2 ([#2563](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2563))
* provider: bumps github.com/hashicorp/terraform-plugin-go from 0.17.0 to 0.18.0 ([#2580](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2580))
* provider: bumps github.com/hashicorp/terraform-plugin-mux from 0.10.0 to 0.11.0 ([#2564](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2564))
* provider: bumps github.com/hashicorp/terraform-plugin-mux from 0.11.0 to 0.11.1 ([#2567](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2567))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.26.1 to 2.27.0 ([#2565](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2565))
* provider: bumps golang.org/x/net from 0.11.0 to 0.12.0 ([#2589](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2589))

## 4.9.0 (28th June, 2023)

NOTES:

* resource/cloudflare_pages_project: Clarify example projects resource ([#2543](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2543))

ENHANCEMENTS:

* resource/cloudflare_notification_policy: Add `alert_trigger_preferences` to the filters block. ([#2535](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2535))
* resource/cloudflare_waiting_room: Add `additional_routes` and `cookie_suffix` to the Waiting Room resource ([#2528](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2528))

BUG FIXES:

* resource/cloudflare_access_ca_certificate: Fix issue with importing existing certificate as the application id was not being set. ([#2539](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2539))
* resource/cloudflare_teams_rules: handle state correctly when `rules_setting` is empty ([#2532](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2532))
* resource/cloudflare_tunnel_config: fix sending incorrect values for various timeouts in the origin configuration block ([#2510](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2510))
* tunnel_config: fix nil pointers for time.Durations ([#2504](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2504))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.69.0 to 0.70.0 ([#2541](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2541))
* provider: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.0 to 1.3.1 ([#2529](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2529))
* provider: bumps github.com/hashicorp/terraform-plugin-go from 0.15.0 to 0.16.0 ([#2536](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2536))

## 4.8.0 (14th June, 2023)

BREAKING CHANGES:

* resource/cloudflare_ruleset: Prevent the rule ID, version and last updated attributes from being set ([#2511](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2511))

ENHANCEMENTS:

* cloudflare_pages_project: add `placement` to deployment config ([#2480](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2480))
* resource/access_application: add support for self_hosted_domains ([#2441](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2441))
* resource/cloudflare_custom_hostname: add support for `bundle_method` TLS configuration ([#2494](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2494))
* resource/cloudflare_device_posture_rule: add ability to create intune and kolide s2s posture rule creation ([#2474](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2474))
* resource/cloudflare_device_settings_policy: add `description` to device settings policy ([#2474](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2474))
* resource/cloudflare_load_balancer: Add support for least_outstanding_requests steering ([#2472](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2472))
* resource/cloudflare_load_balancer_pool: Add support for least_outstanding_requests origin steering ([#2472](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2472))
* resource/cloudflare_page_rule: removes ability to set wildcards for include and exclude, provides guidance on proper values to use instead ([#2491](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2491))
* resource/cloudflare_teams_account: add ability to set `root_ca` for ZT Accounts ([#2474](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2474))

BUG FIXES:

* cloudflare_pages_project: use user provided configuration for secrets in the state handler since the API does not return them ([#2480](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2480))
* resource/cloudflare_certificate_pack: handle UI deletion scenarios for HTTP 404s and `status = "deleted"` responses ([#2497](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2497))
* resource/cloudflare_custom_hostname: use user provided values for state management when the API response isn't provided ([#2494](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2494))
* resource/cloudflare_origin_ca_certificate: mark `csr` as Required ([#2496](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2496))
* resource/cloudflare_ruleset: Mark that the ruleset must be re-created if the shareable entitlement name attribute changes ([#2511](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2511))
* resource/cloudflare_ruleset: Populate the rule ID, ref, version and last updated attributes in API requests and from API responses ([#2511](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2511))
* resource/cloudflare_ruleset: Populate the shareable entitlement name attribute in API requests and from API responses ([#2511](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2511))
* resource/cloudflare_ruleset: handle `Import` operations where the required values are missing for providing a nicer error message ([#2503](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2503))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.68.0 to 0.69.0 ([#2507](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2507))
* provider: bumps github.com/hashicorp/terraform-plugin-framework from 1.2.0 to 1.3.0 ([#2509](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2509))
* provider: bumps github.com/hashicorp/terraform-plugin-log from 0.8.0 to 0.9.0 ([#2489](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2489))
* provider: bumps github.com/hashicorp/terraform-plugin-testing from 1.2.0 to 1.3.0 ([#2524](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2524))
* provider: bumps golang.org/x/net from 0.10.0 to 0.11.0 ([#2523](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2523))
* provider: bumps goreleaser/goreleaser-action from 4.2.0 to 4.3.0 ([#2519](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2519))

## 4.7.1 (31st May, 2023)

BUG FIXES:

* resource/cloudflare_list: remove `IsIPAddress` validation that doesn't take into account CIDR notation ([#2486](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2486))

## 4.7.0 (31st May, 2023)

NOTES:

* resource/cloudflare_filter: This resource is being deprecated in favor of the `cloudflare_rulesets` resource. See https://developers.cloudflare.com/waf/reference/migration-guides/firewall-rules-to-custom-rules/#relevant-changes-for-terraform-users for more details. ([#2442](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2442))
* resource/cloudflare_firewall_rule: This resource is being deprecated in favor of the `cloudflare_rulesets` resource. See https://developers.cloudflare.com/waf/reference/migration-guides/firewall-rules-to-custom-rules/#relevant-changes-for-terraform-users for more details. ([#2442](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2442))

FEATURES:

* **New Resource:** `cloudflare_r2_bucket` ([#2378](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2378))

ENHANCEMENTS:

* resource/cloudflare_account: provide account ID for error handling in `resourceCloudflareAccountDelete` ([#2436](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2436))
* resource/cloudflare_device_posture_integration: add `api_url` to `uptycs` posture integration config. ([#2468](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2468))
* resource/cloudflare_list: add support for Hostname and ASN lists. ([#2483](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2483))
* resource/cloudflare_tunnel_config: add support for origin config on ingress rule and access ([#2477](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2477))

BUG FIXES:

* resource/cloudflare_logpush_job: Properly set dataset field when importing logpush jobs ([#2444](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2444))
* resource/cloudflare_pages_project: suggest a better default value for root_dir ([#2440](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2440))
* resource/cloudflare_ruleset: Validation of ttls for action_parameters with edge_ttl or browser_ttl mode of override_origin ([#2454](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2454))
* resource/cloudflare_workers_kv: Fix import to properly parse the id ([#2434](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2434))

DEPENDENCIES:

* provider: bumps dependabot/fetch-metadata from 1.4.0 to 1.5.0 ([#2463](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2463))
* provider: bumps dependabot/fetch-metadata from 1.5.0 to 1.5.1 ([#2469](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2469))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.67.0 to 0.68.0 ([#2466](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2466))
* provider: bumps github.com/stretchr/testify from 1.8.2 to 1.8.3 ([#2457](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2457))
* provider: bumps github.com/stretchr/testify from 1.8.3 to 1.8.4 ([#2484](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2484))

## 4.6.0 (17th May, 2023)

ENHANCEMENTS:

* resource/cloudflare_ruleset: add support for `auto` compression in the `compress_response` action ([#2409](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2409))
* resource/cloudflare_waiting_room_settings: add support for waiting room zone-level settings. ([#2419](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2419))

BUG FIXES:

* resource/cloudflare_notification_policy: Fix unexpected crashes when setting target_hostname with a filters attribute ([#2425](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2425))
* resource/cloudflare_ruleset: allow `FromValue.PreserveQueryString` to be nullable and handled correctly ([#2414](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2414))
* resource/cloudflare_ruleset: allow using `0` as an edge TTL value without conflicting with Go types for zeros ([#2415](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2415))
* resource/cloudflare_turnstile_widget: align schema to match what is returned by the API and fix updating the widget ([#2413](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2413))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.66.0 to 0.67.0 ([#2429](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2429))
* provider: bumps golang.org/x/net from 0.9.0 to 0.10.0 ([#2421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2421))

## 4.5.0 (3rd May, 2023)

FEATURES:

* **New Resource:** `cloudflare_regional_hostname` ([#2396](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2396))
* **New Resource:** `cloudflare_turnstile_widget` ([#2380](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2380))

ENHANCEMENTS:

* resource/cloudflare_device_posture_rule: Add support for `sentinelone` type. ([#2279](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2279))
* resource/cloudflare_logpush_job: Fix schema for logpush job `dataset` field ([#2397](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2397))
* resource/cloudflare_logpush_job: add max upload parameters ([#2394](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2394))
* resource/cloudflare_logpush_job: add support for `device_posture_results` and `zero_trust_network_sessions`. ([#2405](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2405))
* resource/cloudflare_notification_policy: Added support for setting Megabits per second threshold for dos alert in Cloudflare notification policy resource. ([#2404](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2404))
* resource/cloudflare_pages_project: added secrets to Pages project. Secrets are encrypted environment variables, ideal for secrets such as API tokens. See documentation here: https://developers.cloudflare.com/pages/platform/functions/bindings/#secrets ([#2399](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2399))
* resource/cloudflare_ruleset: add support for the `compress_response` action ([#2372](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2372))
* resource/cloudflare_ruleset: add support for the `http_response_compression` phase ([#2372](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2372))

BUG FIXES:

* resource/cloudflare_load_balancer: fixes random_steering being unset on value updates ([#2403](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2403))
* resource/cloudflare_pages_project: fixes pages project acceptance test ([#2402](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2402))
* resource/cloudflare_ruleset: ensure custom cache keys using query parameters are defined as known values for state handling ([#2388](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2388))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.65.0 to 0.66.0 ([#2398](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2398))
* provider: bumps github.com/hashicorp/terraform-plugin-mux from 0.9.0 to 0.10.0 ([#2395](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2395))

## 4.4.0 (19th April, 2023)

NOTES:

* resource/cloudflare_ruleset: introduced future deprecation warning for the `http_request_sbfm` phase. ([#2382](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2382))

ENHANCEMENTS:

* resource/cloudflare_access_organization: Add auto_redirect_to_identity flag ([#2356](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2356))
* resource/cloudflare_access_policy: Add isolation_required flag ([#2351](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2351))
* resource/cloudflare_tunnel: Adds config_src parameter ([#2369](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2369))
* resource/cloudflare_worker_script: Add `logpush` attribute ([#2375](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2375))

INTERNAL:

* scripts/generate-changelog-entry: make error message match the executable we are expecting ([#2357](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2357))

DEPENDENCIES:

* provider: bumps dependabot/fetch-metadata from 1.3.6 to 1.4.0 ([#2383](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2383))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.64.0 to 0.65.0 ([#2370](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2370))
* provider: bumps golang.org/x/net from 0.8.0 to 0.9.0 ([#2359](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2359))
* provider: bumps peter-evans/create-or-update-comment from 2 to 3 ([#2355](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2355))

## 4.3.0 (5th April, 2023)

NOTES:

* adds support for a basic `flox` environment project ([#2345](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2345))

FEATURES:

* **New Resource:** `cloudflare_device_dex_tests` ([#2250](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2250))
* **New Resource:** `cloudflare_worker_domain` ([#2339](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2339))

ENHANCEMENTS:

* resource/cloudflare_access_group: Add example of usage of Azure ([#2332](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2332))
* resource/cloudflare_access_identity_provider: add `claims` and `scopes` fields ([#2313](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2313))
* resource/cloudflare_access_identity_provider: add ability for users to enable SCIM provisioning on their Identity Providers ([#2147](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2147))
* resource/cloudflare_device_posture_integration: add support for managing `kolide` third party posture provider. ([#2321](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2321))
* resource/cloudflare_device_settings_policy: use new `cloudflare.ServiceMode` type ([#2331](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2331))
* resource/cloudflare_ruleset: enforce schema validation of conflicting cache key parameters ([#2326](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2326))
* resource/cloudflare_teams_rules: updated gateway rule action audit ssh and rule settings ([#2303](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2303))
* resource/cloudflare_worker_script: Add `compatibility_flags` attribute ([#2324](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2324))
* resources/device_settings_policy: add validation for possible `service_mode_v2_mode` values ([#2331](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2331))

BUG FIXES:

* datasource/cloudflare_devices: Fix cloudflare_devices data source to return devices correctly and not error ([#2348](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2348))
* resource/cloudflare_custom_ssl: fix json sent to API when geo_restrictions are not used ([#2319](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2319))

DEPENDENCIES:

* provider: bumps actions/stale from 7 to 8 ([#2322](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2322))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.63.0 to 0.64.0 ([#2344](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2344))
* provider: bumps github.com/hashicorp/terraform-plugin-go from 0.14.3 to 0.15.0 ([#2333](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2333))
* provider: bumps github.com/hashicorp/terraform-plugin-testing from 1.1.0 to 1.2.0 ([#2320](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2320))

## 4.2.0 (22nd March, 2023)

BREAKING CHANGES:

* resource/cloudflare_ruleset: `status` has been removed in favour of `enabled` now that the workaround for zero values is no longer required ([#2271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2271))

NOTES:

* `cloudflare_ruleset` has been migrated to the `terraform-plugin-framework` in doing so addresses issues with the internal representation of zero values. A downside to this is that to get the full benefits, you will need to remove the resource from your Terraform state (`terraform state rm ...`) and then import the resource back into your state. Along with this, you will need to update any references to `status` which was the previous workaround for the `enabled` values. If you have `status = "enabled"` you will need to replace it with `enabled = true` and similar for `status = "disabled"` to be replaced with `enabled = false`. ([#2271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2271))

FEATURES:

* **New Data Source:** `cloudflare_list` ([#2296](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2296))
* **New Data Source:** `cloudflare_lists` ([#2296](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2296))
* **New Resource:** `cloudflare_address_map` ([#2290](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2290))
* **New Resource:** `cloudflare_list_item` ([#2304](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2304))

ENHANCEMENTS:

* resource/access_organization: add ui_read_only_toggle_reason field ([#2175](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2175))
* resource/cloudflare_device_posture_rule: Support `check_disks` in the `input` block schema. ([#2280](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2280))
* resource/cloudflare_notification_policy_webhooks: ensure `url` triggers recreation, not in-place updates ([#2302](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2302))
* resource/cloudflare_tunnel: rename references of cloudflare_argo_tunnel to cloudflare_tunnel in documentation ([#2281](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2281))
* resource/cloudflare_tunnel_config: add support for import of `cloudflare_tunnel_config` ([#2298](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2298))
* resource/cloudflare_tunnel_config: rename references of cloudflare_argo_tunnel to cloudflare_tunnel in documentation ([#2281](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2281))
* resource/cloudflare_tunnel_route: rename references of cloudflare_argo_tunnel to cloudflare_tunnel in documentation ([#2281](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2281))
* resource/cloudflare_worker_script: Add `compatibility_date` attribute ([#2300](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2300))

BUG FIXES:

* resource/cloudflare_ruleset: support cache rules for status range >= and =< operations ([#2307](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2307))
* resource/cloudflare_teams_account: fixes an issue where accounts that had never configured DLP payload logging would error upon reading this resource ([#2284](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2284))

INTERNAL:

* resource/cloudflare_ruleset: migrate from SDKv2 to `terraform-plugin-framework` ([#2271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2271))
* test: swap SDKv2 testing harness to github.com/hashicorp/terraform-plugin-testing ([#2272](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2272))

DEPENDENCIES:

* provider: bumps actions/setup-go from 3 to 4 ([#2291](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2291))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.62.0 to 0.63.0 ([#2289](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2289))
* provider: bumps github.com/hashicorp/terraform-plugin-framework from 1.1.1 to 1.2.0 ([#2314](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2314))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.25.1-0.20230317190757-53a4ec42ea7e to 2.26.0 ([#2308](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2308))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.26.0 to 2.26.1 ([#2315](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2315))

## 4.1.0 (March 8th, 2023)

ENHANCEMENTS:

* resource/cloudflare_cloudflare_teams_rules: Add untrusted_cert setting to teams rules settings ([#2256](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2256))
* resource/cloudflare_teams_account: Add support for DLP payload logging public key ([#2267](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2267))
* resource/cloudflare_teams_rule: Add support for enabling DLP payload logging per-rule ([#2267](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2267))
* resource/cloudflare_waiting_room: add 'ru-RU' and 'fa-IR' to default_template_language field ([#2262](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2262))

BUG FIXES:

* resource/cloudflare_access_group: fixes an issue where Azure group rules with different identity provider ids would override each other ([#2270](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2270))
* resource/cloudflare_notification_policy: ensure all emails are saved if multiple `email_integration` values specified ([#2248](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2248))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.61.0 to 0.62.0 ([#2268](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2268))
* provider: bumps github.com/stretchr/testify from 1.8.1 to 1.8.2 ([#2263](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2263))
* provider: bumps golang.org/x/net from 0.7.0 to 0.8.0 ([#2274](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2274))

## 4.0.0 (February 21st, 2023)

> **Warning** Prior to upgrading you should ensure you have adequate backups in the event you need to rollback to version 3. This is a major version bump and involves backwards incompatible changes.

[3.x to 4.x upgrade guide](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/guides/version-4-upgrade)

BREAKING CHANGES:

* datasource/cloudflare_waf_groups: removed in favour of `cloudflare_rulesets` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* datasource/cloudflare_waf_packages: removed in favour of `cloudflare_rulesets` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* datasource/cloudflare_waf_rules: removed in favour of `cloudflare_rulesets` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* provider: `account_id` is no longer available as a global configuration option. Instead, use the resource specific attributes. ([#2139](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2139))
* resource/cloudflare_access_bookmark: resource has been removed in favour of configuration on `cloudflare_access_application` ([#2136](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2136))
* resource/cloudflare_access_rule: require explicit `zone_id` or `account_id` and remove implicit fallback to user level rules ([#2157](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2157))
* resource/cloudflare_account_member: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_account_member: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_argo_tunnel: resource has been renamed to `cloudflare_tunnel` ([#2135](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2135))
* resource/cloudflare_ip_list: removed in favour of `cloudflare_list` ([#2137](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2137))
* resource/cloudflare_load_balancer: Migrate session_affinity_attributes from TypeMap to TypeSet ([#1959](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1959))
* resource/cloudflare_load_balancer: `session_affinity_attributes.drain_duration` is now `TypeInt` instead of `TypeString` ([#1959](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1959))
* resource/cloudflare_load_balancer_monitor: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_load_balancer_monitor: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_load_balancer_pool: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_load_balancer_pool: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_spectrum_application: `edge_ip_connectivity` is now nested under `edge_ips` as `connectivity` ([#2219](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2219))
* resource/cloudflare_spectrum_application: `edge_ips.type` is now a required field ([#2219](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2219))
* resource/cloudflare_spectrum_application: `edge_ips` now contains nested attributes other than IP ranges. `type` and `connectivity` have been added. `edge_ips.ips` contains the static IP addresses that used to reside at `edge_ips`. ([#2219](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2219))
* resource/cloudflare_waf_group: removed in favour of `cloudflare_ruleset` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* resource/cloudflare_waf_override: removed in favour of `cloudflare_ruleset` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* resource/cloudflare_waf_package: removed in favour of `cloudflare_ruleset` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* resource/cloudflare_waf_rule: removed in favour of `cloudflare_ruleset` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* resource/cloudflare_workers_kv: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_workers_kv: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_workers_kv_namespace: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_workers_kv_namespace: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_workers_script: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_workers_script: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_zone: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_zone: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))

## 3.35.0 (February 20th, 2023)

FEATURES:

* **New Data Source:** `cloudflare_rulesets` ([#2220](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2220))

ENHANCEMENTS:

* resource/cloudflare_argo_tunnel: mark `tunnel_token` as sensitive ([#2231](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2231))
* resource/cloudflare_device_settings_policy: Add new flag MS IP Exclusion for device policies ([#2236](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2236))
* resource/cloudflare_dlp_profile: Add new `allowed_match_count` field to profiles ([#2210](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2210))

BUG FIXES:

* resource/cloudflare_logpush_job: fixing typo in comment ([#2238](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2238))
* resource/cloudflare_record: always send tags object which allows removal of unwanted tags ([#2205](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2205))
* resource/cloudflare_tunnel_config: use correct notation for nested lists ([#2235](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2235))

INTERNAL:

* internal: bump Go version to 1.20 ([#2243](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2243))

DEPENDENCIES:

* provider: bump golang.org/x/net to v0.7.0 ([#2245](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2245))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.60.0 to 0.61.0 ([#2240](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2240))
* provider: bumps github.com/hashicorp/terraform-plugin-framework-validators from 0.9.0 to 0.10.0 ([#2227](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2227))
* provider: bumps github.com/hashicorp/terraform-plugin-mux from 0.8.0 to 0.9.0 ([#2228](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2228))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.24.1 to 2.25.0 ([#2239](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2239))
* provider: bumps golang.org/x/net from 0.6.0 to 0.7.0 ([#2241](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2241))

## 3.34.0 (February 8th, 2023)

BREAKING CHANGES:

* datasource/cloudflare_waf_groups: removed with no current replacement ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* datasource/cloudflare_waf_packages: removed with no current replacement ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* datasource/cloudflare_waf_rules: removed with no current replacement ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* provider: `account_id` is no longer available as a global configuration option. Instead, use the resource specific attributes. ([#2139](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2139))
* resource/cloudflare_access_bookmark: resource has been removed in favour of configuration on `cloudflare_access_application` ([#2136](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2136))
* resource/cloudflare_access_rule: require explicit `zone_id` or `account_id` and remove implicit fallback to user level rules ([#2157](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2157))
* resource/cloudflare_account_member: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_account_member: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_argo_tunnel: resource has been renamed to `cloudflare_tunnel` ([#2135](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2135))
* resource/cloudflare_ip_list: removed in favour of `cloudflare_list` ([#2137](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2137))
* resource/cloudflare_load_balancer: Migrate session_affinity_attributes from TypeMap to TypeSet ([#1959](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1959))
* resource/cloudflare_load_balancer: `session_affinity_attributes.drain_duration` is now `TypeInt` instead of `TypeString` ([#1959](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1959))
* resource/cloudflare_load_balancer_monitor: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_load_balancer_monitor: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_load_balancer_pool: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_load_balancer_pool: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_notification_policy: alert types `block_notification_review_accepted` and `workers_uptime` have been removed. ([#2215](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2215))
* resource/cloudflare_notification_policy: alert types `g6_health_alert` has been renamed to `load_balancing_health_alert` ([#2215](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2215))
* resource/cloudflare_notification_policy: alert types `g6_pool_toggle_alert` has been renamed to `load_balancing_pool_enablement_alert` ([#2215](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2215))
* resource/cloudflare_notification_policy: alert types `scriptmonitor_alert_new_max_length_script_url` has been renamed to `scriptmonitor_alert_new_max_length_resource_url` ([#2215](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2215))
* resource/cloudflare_notification_policy: alert types `scriptmonitor_alert_new_scripts` has been renamed to `scriptmonitor_alert_new_resources` ([#2215](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2215))
* resource/cloudflare_waf_group: removed in favour of `cloudflare_ruleset` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* resource/cloudflare_waf_override: removed in favour of `cloudflare_ruleset` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* resource/cloudflare_waf_package: removed in favour of `cloudflare_ruleset` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* resource/cloudflare_waf_rule: removed in favour of `cloudflare_ruleset` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* resource/cloudflare_workers_kv: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_workers_kv: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_workers_kv_namespace: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_workers_kv_namespace: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_workers_script: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_workers_script: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_zone: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_zone: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))

FEATURES:

* **New Resource:** `cloudflare_mtls_certificate` ([#2182](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2182))
* **New Resource:** `cloudflare_queue` ([#2134](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2134))

ENHANCEMENTS:

* resource/cloudflare_notification_policy: alert types `block_notification_block_removed`, `fbm_dosd_attack`, `scriptmonitor_alert_new_max_length_resource_url`, `scriptmonitor_alert_new_resources`, `tunnel_health_event`, `tunnel_update_event` have been added. ([#2215](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2215))
* resource/cloudflare_ruleset: Preserve IDs of unmodified rules when updating rulesets ([#2172](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2172))
* resource/cloudflare_ruleset: add support for `score_per_period` and `score_response_header_name` ([#2177](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2177))
* resource/cloudflare_worker_script: add support for `queue_binding` ([#2134](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2134))

BUG FIXES:

* resource/cloudflare_account_member: allow `status` to be computed when not provided ([#2217](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2217))
* resource/cloudflare_page_rule: fix failing page rules acceptance tests ([#2213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2213))
* resource/cloudflare_page_rule: make cache_key_fields optional to align with API constraints ([#2192](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2192))
* resource/cloudflare_page_rule: remove empty cookie and header fields when applying this resource ([#2208](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2208))
* resource/cloudflare_pages_project: changing `name` will now force recreation of the project ([#2216](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2216))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.59.0 to 0.60.0 ([#2204](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2204))
* provider: bumps goreleaser/goreleaser-action from 4.1.0 to 4.2.0 ([#2201](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2201))

## 3.33.1 (January 25th, 2023)

BUG FIXES:

* provider: remove conflicting `ExactlyOneOf` schema validation from framework schema ([#2185](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2185))

## 3.33.0 (January 25th, 2023)

BREAKING CHANGES:

* datasource/cloudflare_waf_groups: removed with no current replacement ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* datasource/cloudflare_waf_packages: removed with no current replacement ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* datasource/cloudflare_waf_rules: removed with no current replacement ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* provider: `account_id` is no longer available as a global configuration option. Instead, use the resource specific attributes. ([#2139](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2139))
* resource/cloudflare_access_bookmark: resource has been removed in favour of configuration on `cloudflare_access_application` ([#2136](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2136))
* resource/cloudflare_access_rule: require explicit `zone_id` or `account_id` and remove implicit fallback to user level rules ([#2157](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2157))
* resource/cloudflare_account_member: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_account_member: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_argo_tunnel: resource has been renamed to `cloudflare_tunnel` ([#2135](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2135))
* resource/cloudflare_ip_list: removed in favour of `cloudflare_list` ([#2137](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2137))
* resource/cloudflare_load_balancer_monitor: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_load_balancer_monitor: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_load_balancer_pool: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_load_balancer_pool: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_waf_group: removed in favour of `cloudflare_ruleset` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* resource/cloudflare_waf_override: removed in favour of `cloudflare_ruleset` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* resource/cloudflare_waf_package: removed in favour of `cloudflare_ruleset` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* resource/cloudflare_waf_rule: removed in favour of `cloudflare_ruleset` ([#2138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2138))
* resource/cloudflare_workers_kv: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_workers_kv: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_workers_kv_namespace: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_workers_kv_namespace: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_workers_script: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_workers_script: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))
* resource/cloudflare_zone: `account_id` is now required ([#2153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2153))
* resource/cloudflare_zone: no longer sets `client.AccountID` internally and relies on the resource provided value ([#2154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2154))

ENHANCEMENTS:

* provider: mux `terraform-plugin-sdk/v2` and `terraform-plugin-framework` ([#2170](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2170))
* resource/cloudflare_access_group: supports ip_list property. ([#2073](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2073))
* resource/cloudflare_access_organization: add support for `user_seat_expiration_inactive_time` ([#2115](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2115))
* resource/cloudflare_ruleset: do not let edge_ttl: default be zero ([#2143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2143))
* resource/cloudflare_teams_accounts: adds support for `mailto_address` and `mailto_subject` blockpage settings ([#2146](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2146))
* resource/cloudflare_teams_rules: adds egress rule settings. ([#2159](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2159))

BUG FIXES:

* resource/cloudflare_record: fix issue with DNS comments and tags not being set for new records ([#2148](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2148))

DEPENDENCIES:

* provider: bumps dependabot/fetch-metadata from 1.3.5 to 1.3.6 ([#2183](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2183))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.58.1 to 0.59.0 ([#2166](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2166))

## 3.32.0 (January 11th, 2023)

FEATURES:

* **New Resource:** `cloudflare_device_managed_networks` ([#2126](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2126))

ENHANCEMENTS:

* provider: `X-Auth-Email`, `X-Auth-Key`, `X-Auth-User-Service-Key` and `Authorization` values are now automatically redacted from debug logs ([#2123](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2123))
* provider: use inbuilt cloudflare-go logger for HTTP interactions ([#2123](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2123))
* resource/cloudflare_device_posture_rule: add ability to create crowdstrike s2s posture rule creation ([#2128](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2128))
* resource/cloudflare_origin_ca: support all authentication schemes ([#2124](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2124))
* resource/cloudflare_pages_project: adds support for `always_use_latest_compatibility_date`, `fail_open`, `service_binding` and `usage_model` ([#2083](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2083))
* resource/cloudflare_record: add support for tags and comments. ([#2105](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2105))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.57.1 to 0.58.1 ([#2122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2122))

## 3.31.0 (December 28th, 2022)

NOTES:

* resource/cloudflare_worker_script: supports explicit `account_id` instead of inheriting global values ([#2102](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2102))

FEATURES:

* **New Resource:** `cloudflare_tiered_cache` ([#2101](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2101))

ENHANCEMENTS:

* resource/cloudflare_access_application: makes allowed_idps type to set ([#2094](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2094))
* resource/cloudflare_custom_hostname: add support for defining custom metadata ([#2107](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2107))

BUG FIXES:

* resource/cloudflare_api_shield: allow for empty auth_id_characteristics ([#2091](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2091))
* resource/cloudflare_ruleset: allow edge_ttl -> default to be optional ([#2097](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2097))

DEPENDENCIES:

* provider: bumps actions/stale from 6 to 7 ([#2098](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2098))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.56.0 to 0.57.0 ([#2102](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2102))

## 3.30.0 (December 14th, 2022)

FEATURES:

* **New Data Source:** `cloudflare_load_balancer_pools` ([#1228](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1228))
* **New Resource:** `cloudflare_url_normalization_settings` ([#1878](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1878))

ENHANCEMENTS:

* resource/cloudflare_workers_script: add support for `analytics_engine_binding` bindings ([#2051](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2051))

BUG FIXES:

* resource/access_application: fix issue where session_duration always showed a diff for bookmark apps ([#2076](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2076))
* resource/cloudflare_ruleset: fix issue where SSL setting is based of security level ([#2088](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2088))
* resource/cloudflare_split_tunnel: handle nested attribute changes and ignore ordering ([#2066](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2066))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.55.0 to 0.56.0 ([#2075](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2075))
* provider: bumps goreleaser/goreleaser-action from 3.2.0 to 4.1.0 ([#2087](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2087))

## 3.29.0 (November 30th, 2022)

NOTES:

* datasource/api_token_permission_groups: `permissions` attribute has been deprecated in favour of individual resource level attributes. ([#1960](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1960))

FEATURES:

* **New Resource:** `cloudflare_device_settings_policy` ([#1926](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1926))
* **New Resource:** `cloudflare_tunnel_config` ([#2041](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2041))

ENHANCEMENTS:

* resource/cloudflare_fallback_domain: Add creating fallback domains for device policies ([#1926](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1926))
* resource/cloudflare_logpush_job: add support for `workers_trace_events` ([#2025](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2025))
* resource/cloudflare_origin_ca_certificate: add logic to renew certificate and add a new flag to set if we should renew earlier ([#2048](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2048))
* resource/cloudflare_origin_ca_certificate: trigger a replacement when `csr` is changed ([#2055](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2055))
* resource/cloudflare_origin_ca_certificate: trigger a replacement when `validity` is changed ([#2046](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2046))
* resource/cloudflare_pages_domain: add note about needing to make a separate `cloudflare_record`. ([#2060](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2060))
* resource/cloudflare_pages_project: add note about linking git accounts to Cloudflare account. ([#2060](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2060))
* resource/cloudflare_ruleset: add support for importing existing resources ([#2054](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2054))
* resource/cloudflare_split_tunnel: Add configuring split tunnel for device policies ([#1926](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1926))
* resource/cloudflare_workers_kv: add support for explicitly setting `account_id` on the resource ([#2049](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2049))
* resource/cloudflare_workers_kv_namespace: add support for explicitly setting `account_id` on the resource ([#2049](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2049))
* resource/cloudflare_workers_kv_namespace: swap internals to use new method signatures from cloudflare-go release ([#2049](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2049))

BUG FIXES:

* datasource/api_token_permission_groups: add `user`, `account` and `zone` attributes to contain only those specific resource level permissions. ([#1960](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1960))
* resource/access_policy: Fix issue where only last SAML rule group was applied in
Access policy ([#2033](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2033))
* resource/cloudflare_account: Fix uninitialized cloudflare.Account.Settings ([#2034](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2034))
* resource/cloudflare_custom_hostname: remove `ForceNew` on `wait_for_ssl_pending_validation` ([#2027](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2027))
* resource/cloudflare_list: Do not reapply changes if only list order changed. ([#2063](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2063))
* resource/cloudflare_record: Fix null MX record creation ([#2038](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2038))
* resource/cloudflare_spectrum_application: ignore ordering of `edge_ips` ([#2032](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2032))
* resource/cloudflare_workers_kv: `key` changes force creation of a new resource ([#2044](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2044))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.54.0 to 0.55.0 ([#2049](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2049))

## 3.28.0 (November 16th, 2022)

ENHANCEMENTS:

* resource/cloudflare_zone: add new plans for zone subscriptions ([#2023](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2023))

BUG FIXES:

* resource/access_application: Fix issue where empty CORS headers state causes panics ([#2010](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2010))

DEPENDENCIES:

* provider: bumps dependabot/fetch-metadata from 1.3.4 to 1.3.5 ([#2008](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2008))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.53.0 to 0.54.0 ([#2016](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2016))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.24.0 to 2.24.1 ([#2024](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2024))

## 3.27.0 (November 2nd, 2022)

FEATURES:

* **New Resource:** `cloudflare_access_organization` ([#1961](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1961))
* **New Resource:** `cloudflare_dlp_profile` ([#1984](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1984))
* **New Resource:** `cloudflare_total_tls` ([#1979](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1979))
* **New Resource:** `cloudflare_waiting_room_rules` ([#1957](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1957))

ENHANCEMENTS:

* resource/cloudflare_access_application: add support for `app_launcher`, `biso`, `dash_sso` and `warp` to the schema ([#1988](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1988))
* resource/cloudflare_load_balancer_monitor: support defining explicit `account_id` for resources ([#1986](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1986))
* resource/cloudflare_load_balancer_pool: support defining explicit `account_id` for resources ([#1986](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1986))
* resource/cloudflare_logpush_job: add support for "access_requests" dataset parameter ([#2001](https://github.com/cloudflare/terraform-provider-cloudflare/issues/2001))
* resource/cloudflare_teams_list: handle pagination for larger Team List accounts ([#1706](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1706))
* test: use `T.Setenv` to set env vars in provider tests ([#1985](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1985))

BUG FIXES:

* resource/cloudflare_access_group: fix issue where policy groups were always showing a diff during plans ([#1983](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1983))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.52.0 to 0.53.0 ([#1995](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1995))
* provider: bumps github.com/stretchr/testify from 1.8.0 to 1.8.1 ([#1993](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1993))

## 3.26.0 (October 19th, 2022)

ENHANCEMENTS:

* resource/cloudflare_custom_hostname: Add `wait_for_ssl_pending_validation` attribute ([#1953](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1953))
* resource/cloudflare_device_posture_rule: Add chromeos and unique_client_id values ([#1950](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1950))
* resource/cloudflare_load_balancer: Migrate to autogen docs, improve docs ([#1954](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1954))
* resource/cloudflare_pages_domain: add Pages project domain importer. ([#1973](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1973))
* resource/cloudflare_ruleset: add support for overriding sensitivity levels for ruleset rules ([#1965](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1965))

BUG FIXES:

* resource/cloudflare_byo_ip_prefix: set correct prefix ID for the byoip prefix during import. ([#1951](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1951))
* resource/cloudflare_custom_ssl: check GeoRestrictions is not nil before attempting to compare it ([#1964](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1964))
* resource/cloudflare_pages_project: add defaults to Pages project deployment config ([#1973](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1973))
* resource/cloudflare_zone_settings_override: Fetch/modify `origin_max_http_version` as a single setting. ([#1805](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1805))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.51.0 to 0.52.0 ([#1962](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1962))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.23.0 to 2.24.0 ([#1969](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1969))
* provider: bumps goreleaser/goreleaser-action from 3.1.0 to 3.2.0 ([#1977](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1977))

## 3.25.0 (October 5th, 2022)

NOTES:

* resource/device_posture_rule: update device posture rule to reflect new linux posture fields ([#1842](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1842))

ENHANCEMENTS:

* resource/cloudflare_account_member: permit setting status in terraform schema if desired ([#1920](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1920))
* resource/cloudflare_email_routing_catch_all: switch to a dedicated scheme to allow type = "drop" ([#1947](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1947))
* resource/cloudflare_load_balancer: Add support for adaptive_routing, location_strategy, random_steering, and zero_downtime_failover ([#1941](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1941))
* resource/cloudflare_load_balancer: update internal method signatures to match upstream library ([#1932](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1932))
* resource/cloudflare_load_balancer_monitor: update internal method signatures to match upstream library ([#1932](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1932))
* resource/cloudflare_load_balancer_pool: update internal method signatures to match upstream library ([#1932](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1932))

BUG FIXES:

* provider: allow individual setting of x-auth-service-key ([#1923](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1923))
* provider: fix versioning injection during release builds ([#1935](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1935))
* resource/cloudflare_byo_ip_prefix: fix `Import` to set `account_id` ([#1930](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1930))
* resource/cloudflare_record: update Read method to pull from remote API instead of local configuration which is empty during `Import` ([#1942](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1942))
* resource/cloudflare_zone_settings_override: Fix array manipulation bug related to single zone settings ([#1925](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1925))

DEPENDENCIES:

* provider: bumps actions/stale from 5 to 6 ([#1922](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1922))
* provider: bumps dependabot/fetch-metadata from 1.3.3 to 1.3.4 ([#1945](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1945))

## 3.24.0 (September 21st, 2022)

NOTES:

* resource/cloudflare_access_bookmark: Bookmark resource is deprecated in favor of using the `cloudflare_access_application` resource. ([#1914](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1914))
* resource/cloudflare_email_routing_rule: Fix example resource to use correct syntax ([#1895](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1895))
* resource/cloudflare_email_routing_rule_catch_all: Fix example resource to use correct syntax ([#1895](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1895))

FEATURES:

* **New Data Source:** `cloudflare_accounts` ([#1899](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1899))
* **New Data Source:** `cloudflare_record` ([#1906](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1906))
* **New Resource:** `cloudflare_account` ([#1902](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1902))
* **New Resource:** `cloudflare_user_agent_blocking_rule` ([#1894](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1894))

ENHANCEMENTS:

* resource/cloudflare_pages_project: Adds importer for pages_project ([#1886](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1886))
* tools: add devcontainer for local development ([#1892](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1892))

BUG FIXES:

* provider: allow setting `api_user_service_key` without token and/or key ([#1907](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1907))
* resource/cloudflare_load_balancer_monitor: fix detection of headers values changing ([#1903](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1903))
* resource/cloudflare_pages_project: fix null source on project create ([#1898](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1898))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.49.0 to 0.50.0 ([#1910](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1910))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.21.0 to 2.22.0 ([#1900](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1900))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.22.0 to 2.23.0 ([#1913](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1913))

## 3.23.0 (September 7th, 2022)

FEATURES:

* **New Resource:** `cloudflare_api_shield` ([#1874](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1874))
* **New Resource:** `cloudflare_email_routing_address` ([#1856](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1856))
* **New Resource:** `cloudflare_email_routing_catch_all` ([#1856](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1856))
* **New Resource:** `cloudflare_email_routing_rules` ([#1856](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1856))
* **New Resource:** `cloudflare_email_routing_settings` ([#1856](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1856))
* **New Resource:** `cloudflare_web3_hostname` ([#1882](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1882))

ENHANCEMENTS:

* resource/cloudflare_access_service_token: updates internals to allow in place refreshing instead of full replacement based on the `expires_at` and `min_days_for_renewal` values ([#1872](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1872))
* resource/cloudflare_pages_domain: Adds support for Pages domains ([#1835](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1835))
* resource/cloudflare_pages_project: Adds support for Pages Projects ([#1835](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1835))
* resource/cloudflare_record: Add HTTPS DNS record type ([#1887](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1887))
* resource/cloudflare_worker: provide js module option to allow service bindings ([#1865](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1865))

BUG FIXES:

* resource/cloudflare_authenticated_origin_pulls: fix improper handling of enabled=false ([#1861](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1861))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.48.0 to 0.49.0 ([#1871](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1871))
* provider: bumps github.com/golangci/golangci-lint from 1.48.0 to 1.49.0 ([#1855](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1855))
* provider: bumps goreleaser/goreleaser-action from 3.0.0 to 3.1.0 ([#1868](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1868))

## 3.22.0 (August 24th, 2022)

NOTES:

* update local setup documentation to reflect newer required Go version ([#1847](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1847))

ENHANCEMENTS:

* resource/cloudflare_ruleset: add support for `http_config_settings` ([#1837](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1837))
* resources/worker_script: add support for r2_bucket_binding ([#1825](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1825))

BUG FIXES:

* resource/cloudflare_fallback_domain: fix perpetual changes due to ordering ([#1828](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1828))
* resource/cloudflare_notification_policy: add missing alert types and filters to validation and docs ([#1830](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1830))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.46.0 to 0.47.1 ([#1844](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1844))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.20.0 to 2.21.0 ([#1838](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1838))
* provider: bumps github.com/hcloudflare-go from 0.47.1 to 0.48.0 ([#1848](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1848))

## 3.21.0 (August 10th, 2022)

BREAKING CHANGES:

* resource/cloudflare_page_rule: Removed `always_online` from page rules since this action has been decommissioned from page rules ([#1817](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1817))

ENHANCEMENTS:

* resource/cloudflare_custom_ssl: handle when remote ID changes during updates ([#1824](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1824))
* resource/cloudflare_ruleset: add support and configuration for `serve_errors` action ([#1794](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1794))
* resource/cloudflare_ruleset: add support for sni override in route action ([#1816](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1816))

BUG FIXES:

* resource/cloudflare_account_member: actually use the `account_id` value ([#1823](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1823))
* resource/cloudflare_zone_settings_override: add missing allowed value of 120 for `browser_cache_ttl` ([#1822](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1822))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.45.0 to 0.46.0 ([#1815](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1815))
* provider: bumps github.com/golangci/golangci-lint from 1.47.2 to 1.47.3 ([#1813](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1813))
* provider: bumps github.com/golangci/golangci-lint from 1.47.3 to 1.48.0 ([#1820](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1820))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.19.0 to 2.20.0 ([#1804](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1804))

## 3.20.0 (July 27th, 2022)

BREAKING CHANGES:

* resource/cloudflare_healthcheck: deprecates `notification_email_addresses` and `notification_suspended` in favour of `cloudflare_notification_policy` ([#1789](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1789))

NOTES:

* resource/cloudflare_access_rule: this resource now supports an explicit `account_id` instead of the implied one from the client configuration. You should update your configuration to include `account_id` and remove permadiffs. ([#1790](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1790))
* resource/cloudflare_account_member: this resource now supports an explicit `account_id` instead of the implied one from the client configuration. You should update your configuration to include `account_id` and remove permadiffs. ([#1767](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1767))
* resource/cloudflare_certificate_pack: remove references to long-deprecated dedicated certs (replaced by `advanced`) ([#1778](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1778))
* resource/cloudflare_rulesets: Cache Rules use cache flag instead of bypass_cache ([#1785](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1785))
* resource/cloudflare_zone: this resource now supports an explicit `account_id` instead of the implied one from the client configuration. You should update your configuration to include `account_id` and remove permadiffs. ([#1767](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1767))

ENHANCEMENTS:

* resource/cloudflare_access_application: Add support for Saas applications ([#1762](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1762))
* resource/cloudflare_access_rule: add support for `account_id` ([#1790](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1790))
* resource/cloudflare_account_member: add support for `account_id` ([#1767](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1767))
* resource/cloudflare_api_token: add support for `not_before` and `expires_on` ([#1792](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1792))
* resource/cloudflare_certificate_pack: fix some of the custom hostname docs copy ([#1778](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1778))
* resource/cloudflare_certificate_pack: update the list of allowed certificate authorities ([#1778](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1778))
* resource/cloudflare_load_balancer: Add support for LB country pools ([#1797](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1797))
* resource/cloudflare_managed_headers: swap filtering to use API instead of custom logic ([#1765](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1765))
* resource/cloudflare_ruleset: add support for `from_value` action parameter when using redirect action ([#1781](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1781))
* resource/cloudflare_zone: add support for `account_id` ([#1767](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1767))

BUG FIXES:

* resource/cloudflare_waiting_room: fix default waiting room `session_duration` and `path` values ([#1766](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1766))
* resource/cloudflare_zone_lockdown: Fix crash when logging upstream error message ([#1777](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1777))

DEPENDENCIES:

* provider: bumps github.com/cloudflare/cloudflare-go from 0.44.0 to 0.45.0 ([#1793](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1793))
* provider: bumps github.com/golangci/golangci-lint from 1.46.2 to 1.47.0 ([#1786](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1786))
* provider: bumps github.com/golangci/golangci-lint from 1.47.0 to 1.47.1 ([#1788](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1788))
* provider: bumps github.com/golangci/golangci-lint from 1.47.1 to 1.47.2 ([#1795](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1795))
* provider: bumps github.com/hashicorp/terraform-plugin-log from 0.4.1 to 0.5.0 ([#1773](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1773))
* provider: bumps github.com/hashicorp/terraform-plugin-log from 0.5.0 to 0.6.0 ([#1780](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1780))
* provider: bumps github.com/hashicorp/terraform-plugin-log from 0.6.0 to 0.7.0 ([#1798](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1798))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.18.0 to 2.19.0 ([#1779](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1779))

## 3.19.0 (July 13th, 2022)

ENHANCEMENTS:

* resource/cloudflare_ipsec_tunnel: add allow_null_cipher to ipsec tunnel ([#1736](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1736))
* resource/cloudflare_record: Validate that DNS record names are non-empty ([#1740](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1740))
* resource/cloudflare_ruleset: add support for `from_list` action parameter when using redirect action ([#1744](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1744))
* resource/cloudflare_waiting_room: Add queueing_method field. ([#1759](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1759))
* resource/cloudflare_workers_script: add support for `service_binding` bindings ([#1760](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1760))
* resource/cloudflare_zone_settings_override: Add support for `origin_max_http_version` ([#1755](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1755))

BUG FIXES:

* resource/cloudflare_list: fix default values for redirect list updates ([#1746](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1746))
* resource/cloudflare_logpush_job: fix logpush job name validation regex ([#1743](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1743))
* resource/cloudflare_tunnel_route: Fix incorrect indexing of resource data id attributes ([#1753](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1753))

DEPENDENCIES:

* provider: bumps dependabot/fetch-metadata from 1.3.1 to 1.3.2 ([#1747](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1747))
* provider: bumps dependabot/fetch-metadata from 1.3.2 to 1.3.2 ([#1748](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1748))
* provider: bumps github.com/cloudflare/cloudflare-go from 0.43.0 to 0.44.0 ([#1757](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1757))
* provider: bumps github.com/hashicorp/terraform-plugin-docs from 0.12.0 to 0.13.0 ([#1763](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1763))
* provider: bumps github.com/hashicorp/terraform-plugin-sdk/v2 from 2.17.0 to 2.18.0 ([#1758](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1758))
* provider: bumps github.com/stretchr/testify from 1.7.5 to 1.8.0 ([#1738](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1738))

## 3.18.0 (June 29th, 2022)

NOTES:

* resource/cloudflare_ip_list: Deprecated cloudflare_ip_list in favor of cloudflare_list. ([#1700](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1700))

FEATURES:

* **New Resource:** `cloudflare_managed_headers` ([#1688](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1688))
* **New Resource:** `resource/cloudflare_list: Added support for generic list types, including redirect lists.` ([#1700](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1700))

ENHANCEMENTS:

* resource/cloudflare_logpush_job: adds support for `kind` attribute ([#1718](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1718))
* resource/cloudflare_logpush_job: validate name attribute ([#1717](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1717))
* resource/cloudflare_ruleset: add support for set cache settings ([#1701](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1701))

BUG FIXES:

* resource/cloudflare_logpush_job: Fix for optional `filter` attribute ([#1712](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1712))
* resource/cloudflare_logpush_job: fix unmarhalling job with empty/no filter ([#1723](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1723))
* resource/cloudflare_record: ensure trailing `.` in `value` don't cause surious diffs ([#1713](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1713))

## 3.17.0 (June 15th, 2022)

BREAKING CHANGES:

* resource/cloudflare_ruleset: deprecates `enabled` in overridden configurations immediately in favour of `status` ([#1689](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1689))

FEATURES:

* **New Resource:** `cloudflare_tunnel_virtual_network` ([#1672](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1672))

ENHANCEMENTS:

* resource/cloudflare_access_identity_provider: Add support for PKCE when creating IDPS. ([#1667](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1667))
* resource/cloudflare_device_posture_integration: add support for managing `uptycs`, `intune` and `crowdstrike` third party posture providers. ([#1628](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1628))
* resource/cloudflare_ipsec_tunnel: add support for `healthcheck_enabled`, `health_check_target`, `healthcheck_type`, `psk` ([#1685](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1685))
* resource/cloudflare_logpush_job: Add `filter` field support ([#1660](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1660))
* resource/cloudflare_tunnel_route: Add `virtual_network_id` attribute ([#1668](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1668))

BUG FIXES:

* resource/cloudflare_teams_rule: Fixes issue with rule precedence updates by using a generated version of precendence in API calls to reduce clashing versions ([#1663](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1663))

## 3.16.0 (June 1st, 2022)

NOTES:

* provider: swap internal logging mechanism to use `tflog` ([#1638](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1638))
* provider: updated internal package structure of repository ([#1636](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1636))

ENHANCEMENTS:

* resource/cloudflare_access_group: add support for external evaluation as a new access group rule ([#1623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1623))
* resource/cloudflare_argo_tunnel: add `tunnel_token` support ([#1590](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1590))
* resource/cloudflare_logpush_job: add support for specifying `frequency` ([#1634](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1634))
* resource/cloudflare_ruleset: add support for custom fields logging ([#1630](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1630))
* resource/cloudflare_waiting_room: Add default_template_language field. ([#1651](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1651))

BUG FIXES:

* resource/cloudflare_access_application: Fix inability to update `http_only_cookie_attribute` to false ([#1602](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1602))
* resource/cloudflare_waiting_room_event: handle time pointer for nullable struct member ([#1648](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1648))
* resource/cloudflare_workers_kv: handle invalid id during terraform import ([#1635](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1635))

## 3.15.0 (May 18th, 2022)

NOTES:

* provider: internally swapped to using `diag.Diagnostics` for CRUD return types and using `context.Context` passed in from the provider itself instead of instantiating our own in each operation ([#1592](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1592))

ENHANCEMENTS:

* resource/cloudflare_device_posture_rule: Add `expiration` to device posture rule ([#1585](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1585))
* resource/cloudflare_logpush_job: add support for managing `network_analytics_logs` ([#1627](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1627))
* resource/cloudflare_logpush_job: allow r2 logpush destinations without ownership validation ([#1597](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1597))
* resource/ruleset: add support for `origin` and `host_header` attributes ([#1620](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1620))

BUG FIXES:

* resource/cloudflare_access_rule: Fix lifecycle of access_rule update ([#1601](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1601))
* resource/cloudflare_spectrum_application: prevent panic when configuration does not include `edge_ips.connectivity` ([#1599](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1599))
* resource/cloudflare_teams_rule: fixed detection of deleted teams rules ([#1622](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1622))

## 3.14.0 (May 4th, 2022)

FEATURES:

* **New Resource:** `cloudflare_tunnel_route` ([#1572](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1572))

ENHANCEMENTS:

* resource/cloudflare_certificate_pack: add support for new option (`wait_for_active_status`) to block creation until certificate pack is active ([#1567](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1567))
* resource/cloudflare_notification_policy: Add `slo` to notification policy filters ([#1573](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1573))
* resource/cloudflare_teams_list: Add support for IP type ([#1550](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1550))

BUG FIXES:

* cloudflare_tunnel_routes: Fix reads matching routers with larger CIDRs ([#1581](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1581))
* resource/cloudflare_access_group: allow github access groups to be created without a list of teams ([#1589](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1589))
* resource/cloudflare_logpush_job: make ownership challenge check for https not required ([#1588](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1588))
* resource/cloudflare_tunnel_route: Fix importing resource ([#1580](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1580))
* resource/cloudflare_zone: update plan identifier for professional rate plans ([#1583](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1583))

## 3.13.0 (April 20th, 2022)

NOTES:

* resource/cloudflare_byo_ip_prefix: now requires an explicit `account_id` parameter instead of implicitly relying on `client.AccountID` ([#1563](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1563))
* resource/cloudflare_ip_list: no longer sets `client.AccountID` internally for resources ([#1563](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1563))
* resource/cloudflare_magic_firewall_ruleset: no longer sets `client.AccountID` internally for resources ([#1563](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1563))
* resource/cloudflare_static_route: no longer sets `client.AccountID` internally for resources ([#1563](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1563))
* resource/cloudflare_worker_cron_trigger: now requires an explicit `account_id` parameter instead of implicitly relying on `client.AccountID` ([#1563](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1563))

ENHANCEMENTS:

* resource/cloudflare_custom_pages: add support for managed_challenge action ([#1478](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1478))
* resource/cloudflare_ruleset: add support for rule `logging` ([#1538](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1538))

## 3.12.2 (April 13th, 2022)

ENHANCEMENTS:

* resource/cloudflare_ruleset: Setting description to `Optional` to better reflect API requirements ([#1556](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1556))

## 3.12.1 (April 9th, 2022)

BUG FIXES:

* resource/cloudflare_zone: don't get stuck in endless loop for partner zone rate plans ([#1547](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1547))

## 3.12.0 (April 6th, 2022)

NOTES:

* resource/cloudflare_healthcheck: `notification_suspended` and `notification_email_addresses` attributes are being deprecated in favour of `cloudflare_notification_policy` resource instead. ([#1529](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1529))

FEATURES:

* **New Resource:** `cloudflare_access_bookmark` ([#1539](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1539))

ENHANCEMENTS:

* resource/cloudflare_access_application: Add service_auth_401_redirect field. ([#1540](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1540))

BUG FIXES:

* resource/cloudflare_api_token: ignore ordering changes in `permission_groups` ([#1545](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1545))
* resource/cloudflare_notification_policy: Fix unexpected crashes when using cloudflare_notification_policy with a filters attribute ([#1542](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1542))
* resource/cloudflare_zone_dnssec: don't try to enable DNSSEC when state is "pending" ([#1530](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1530))

## 3.11.0 (March 23rd, 2022)

NOTES:

* resource/cloudflare_origin_ca_certificate: `requested_validity` no longer decrements until the `expires_on` value but is now the amount of days the certificate was requested for. ([#1502](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1502))

FEATURES:

* **New Resource:** `cloudflare_teams_proxy_endpoint` ([#1517](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1517))
* **New Resource:** `cloudflare_waiting_room_event` ([#1509](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1509))

ENHANCEMENTS:

* resource/cloudflare_page_rule: add support for `actions.disable_zaraz` ([#1523](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1523))
* resource/cloudflare_ruleset: add support for `action_parameters.response` to control the response when triggering a WAF rule ([#1507](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1507))
* resource/cloudflare_ruleset: add support for `ratelimit.requests_to_origin` ([#1507](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1507))

BUG FIXES:

* resource/cloudflare_device_posture_integration: remove superfluous `id` from schema ([#1504](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1504))
* resource/cloudflare_spectrum_application: Fix 'edge_ip_connectivity' state persistence ([#1515](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1515))

## 3.10.1 (March 10th, 2022)

BUG FIXES:

- resource/cloudflare_ruleset: don't attempt to upgrade ratelimit if it isn't set ([#1501](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1501))

## 3.10.0 (March 9th, 2022)

BREAKING CHANGES:

- resource/cloudflare_ruleset: rename `mitigation_expression` to `counting_expression` ([#1477](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1477))

ENHANCEMENTS:

- resource/cloudflare_access_rule: add support for managed_challenge action ([#1457](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1457))
- resource/cloudflare_custom_hostname: adds support for custom_origin_sni ([#1482](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1482))
- resource/cloudflare_device_policy_certificates: add support for device policy certificate settings ([#1467](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1467))
- resource/cloudflare_teams_rules: Add `insecure_disable_dnssec_validation` option to settings ([#1469](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1469))
- resource/cloudflare_zone: add support for partner rate plans ([#1464](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1464))

BUG FIXES:

- resource/cloudflare_record: no need to pass the resourceCloudflareRecordUpdate to the NonRetryable handler ([#1496](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1496))

## 3.9.1 (February 15th, 2022)

NOTES:

- resource/cloudflare_api_token: revert swap from TypeList to TypeSet due to broken migration ([#1455](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1455))

FEATURES:

- **New Data Source:** `cloudflare_devices` ([#1453](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1453))

## 3.9.0 (February 14th, 2022)

FEATURES:

- **New Resource:** `cloudflare_gre_tunnel` ([#1423](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1423))
- **New Resource:** `cloudflare_zone_cache_variants` ([#1444](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1444))

ENHANCEMENTS:

- cloudflare_ruleset: add support for "managed_challenge" action ([#1442](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1442))
- resource/certificate_pack: adds `validation_errors` and `validation_records` with same format as custom hostnames. ([#1424](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1424))
- resource/custom_hostname: also adds missing `validation_errors`, and `certificate_authority` ([#1424](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1424))
- resource/custom_hostname: validation tokens are now an array (`validation_records`) instead of a top level, but the only top level record that was previously here was for cname validation, txt/http/email were entirely missing. ([#1424](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1424))

BUG FIXES:

- cloudflare_argo_tunnel: conditionally fetch settings based on the provided configuration ([#1451](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1451))
- resource/cloudflare_api_token: ignore ordering of `permission_group` IDs ([#1425](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1425))

## 3.8.0 (January 28th, 2022)

FEATURES:

- **New Resource:** `cloudflare_ipsec_tunnel` ([#1404](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1404))

ENHANCEMENTS:

- datasource/cloudflare_zones: allow filtering by account_id ([#1401](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1401))
- resource/cloudflare_cloudflare_teams_rules: Add `check_session` and `add_headers` attributes to settings ([#1402](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1402))
- resource/cloudflare_cloudflare_teams_rules: Add `disable_download`, `disable_keyboard`, and `disable_upload` attributes to `BISOAdminControls` ([#1402](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1402))
- resource/cloudflare_logpush_job: add support for managing `dns_logs` ([#1400](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1400))
- resource/cloudflare_ruleset: add skip support for `products` and `phases` ([#1391](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1391))
- resource/cloudflare_ruleset: smoother handling of UI/API collisions during migrations ([#1393](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1393))
- resource/cloudflare_teams_accounts: Add the `fips` field for configuring FIPS-compliant TLS. ([#1380](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1380))

BUG FIXES:

- resource/cloudflare_fallback_domain: default entries are now restored on delete. ([#1399](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1399))
- resource/cloudflare_ruleset: conditionally set action parameter "version" ([#1388](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1388))
- resource/cloudflare_ruleset: fix handling of `false` values for category/rule overrides ([#1405](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1405))

## 3.7.0 (January 13th, 2022)

FEATURES:

- **New Resource:** `cloudflare_device_posture_integration` ([#1340](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1340))
- **New Resource:** `cloudflare_fallback_domain` ([#1356](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1356))

ENHANCEMENTS:

- resource/cloudflare_firewall_rule: add support for managed_challenge action ([#1378](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1378))
- resource/cloudflare_load_balancer_monitor: added support for smtp, icmp_ping, and udp_icmp monitors ([#1371](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1371))
- resource/cloudflare_logpush_job: add support for account-level logpush jobs ([#1311](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1311))
- resource/cloudflare_logpush_ownership_challenge: add support for account-level logpush ownership challenges ([#1311](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1311))

BUG FIXES:

- resource/cloudflare_api_token: modified_on is now read correctly ([#1368](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1368))

DEPENDENCIES:

- `github.com/cloudflare/cloudflare-go` v0.29.0 => v0.30.0 ([#1379](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1379))

## 3.6.0 (December 23rd, 2021)

ENHANCEMENTS:

- resource/cloudflare_access_application: add bookmark type to apptypes ([#1343](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1343))
- resource/cloudflare_teams_rules: GATE-2273: Adds support for device posture gateway rules ([#1353](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1353))

BUG FIXES:

- resource/cloudflare_load_balancer: handle empty `rules` for `resourceCloudflareLoadBalancerStateUpgradeV1` ([#1257](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1257))
- resource/cloudflare_split_tunnel: import will now use correct import function ([#1345](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1345))

## 3.5.0 (December 14th, 2021)

NOTES:

- provider: split schema definition from resource CRUD operations ([#1321](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1321))

FEATURES:

- **New Data Source:** `cloudflare_access_identity_provider` ([#1300](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1300))

ENHANCEMENTS:

- resource/cloudflare_access_application: add support for `app_launcher_visible` to the schema ([#1303](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1303))
- resource/cloudflare_ruleset: add support for rewriting HTTP response headers ([#1339](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1339))
- resource/cloudflare_zone: support changing `type` values ([#1301](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1301))

BUG FIXES:

- resource/cloudflare_access_group: fix mapping error for AzureAD ([#1341](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1341))
- resource/cloudflare_access_rule: allow "ip6" to be a padded or unpadded value and compare correctly ([#1294](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1294))
- resource/cloudflare_argo: call `Read` for `Import` operations ([#1295](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1295))
- resource/cloudflare_argo_tunnel: fix import mechanism ([#1329](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1329))
- resource/cloudflare_argo_tunnel: update CNAME to use `cfargotunnel.com` ([#1293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1293))
- resource/cloudflare_origin_ca_certificate: reintroduce `DiffSuppressFunc` for `requested_validity` changes to handle all schema/SDK combinations ([#1289](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1289))
- resource/cloudflare_split_tunnel: import now works by specifying accountId/mode ([#1313](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1313))
- resource/cloudflare_teams_list: ignore `items` ordering ([#1338](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1338))

## 3.4.0 (November 1st, 2021)

ENHANCEMENTS:

- provider: add the ability to configure a different hostname and base path for the API client ([#1270](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1270))
- resource/cloudflare_access_application: add support for 'skip_interstitial' and 'logo_url' properties ([#1262](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1262))
- resource/cloudflare_custom_hostname: add `settings.early_hints` to ssl schema ([#1286](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1286))
- resource/cloudflare_ruleset: add support for exposed credential checks ([#1263](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1263))
- resource/cloudflare_zone_setting_override: add support for overriding `early_hints` ([#1285](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1285))

BUG FIXES:

- resource/cloudflare_ruleset: allow action parameter override `enabled` to be true/false or uninitialised ([#1275](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1275))
- resource/cloudflare_ruleset: allow setting `uri` and `path` action parmeters together in a single rule ([#1271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1271))

## 3.3.0 (October 20th, 2021)

FEATURES:

- **New Data Source:** `cloudflare_account_roles` ([#1238](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1238))

ENHANCEMENTS:

- resource/cloudflare_access_application: add support for 'SameSite' and 'HttpOnly' cookie attributes ([#1241](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1241))
- resource/cloudflare_argo_tunnel: add `cname` as exported attribute ([#1259](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1259))
- resource/cloudflare_load_balancer_pool: add support for origin steering ([#1240](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1240))
- resource/cloudflare_ruleset: add support for 'Action' and 'Enabled' action_parameters > overrides attributes ([#1249](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1249))
- resource/cloudflare_zone_setting_override: add support for overriding `binary_ast` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
- resource/cloudflare_zone_setting_override: add support for overriding `filter_logs_to_cloudflare` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
- resource/cloudflare_zone_setting_override: add support for overriding `log_to_cloudflare` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
- resource/cloudflare_zone_setting_override: add support for overriding `orange_to_orange` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
- resource/cloudflare_zone_setting_override: add support for overriding `proxy_read_timeout` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))
- resource/cloudflare_zone_setting_override: add support for overriding `visitor_ip` ([#1261](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1261))

BUG FIXES:

- resource/cloudflare_access_policy: handle empty `nil` values for building policies ([#1237](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1237))
- resource/cloudflare_ruleset: don't attempt to update "custom" rulesets using the phase entrypoint ([#1245](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1245))

## 3.2.0 (October 7th, 2021)

NOTES:

- provider: cloudflare-go has been upgraded to v0.25.0 ([#1236](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1236))

FEATURES:

- **New Data Source:** `cloudflare_zone` ([#1213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1213))
- **New Resource:** `cloudflare_split_tunnel` ([#1207](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1207))

ENHANCEMENTS:

- provider: add support for debugging via debuggers (like delve) ([#1217](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1217))
- resource/cloudflare_access_policy: add support for approval_required flag ([#1230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1230))

BUG FIXES:

- resource/cloudflare_account_member: handle role changes made in the dashboard ([#1202](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1202))
- resource/cloudflare_origin_ca_certificate: ignore `requested_validity` changes due to the value decreasing but still store it ([#1214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1214))
- resource/cloudflare_record: handle `Update`s for records with `data` blocks ([#1229](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1229))

## 3.1.0 (September 21st, 2021)

ENHANCEMENTS:

- resource/cloudflare_ruleset: add support for ddos_l7 configuration ([#1212](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1212))

## 3.0.1 (September 21st, 2021)

ENHANCEMENTS:

- resource/cloudflare_access_rule: add state migrator for 3.x ([#1211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1211))
- resource/cloudflare_custom_ssl: add state migrator for 3.x ([#1211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1211))
- resource/cloudflare_load_balancer: add state migrator for 3.x ([#1211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1211))
- resource/cloudflare_record: add state migrator for 3.x ([#1211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1211))

## 3.0.0 (September 20th, 2021)

[2.x to 3.x upgrade guide](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/guides/version-3-upgrade)

BREAKING CHANGES:

- resource/cloudflare_access_rule: `configuration` is now a `TypeList` instead of a `TypeMap` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
- resource/cloudflare_custom_ssl: `custom_ssl_options` is now a `TypeList` instead of `TypeMap` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
- resource/cloudflare_load_balancer: `fixed_response` is now a `TypeList` instead of a `TypeMap` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
- resource/cloudflare_load_balancer: fixed_response.status_code`is now a`TypeInt`instead of a`TypeString` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
- resource/cloudflare_record: `data` is now a `TypeList` instead of a `TypeMap` ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))

NOTES:

- provider: Golang version has been upgraded to 1.17 ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
- provider: HTTP user agent is now "terraform/:version terraform-plugin-sdk/:version terraform-provider-cloudflare/:version" ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
- provider: Minimum Terraform core version is now 0.14 ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
- provider: terraform-plugin-sdk has been upgraded to 2.x ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))

ENHANCEMENTS:

- resource/cloudflare_custom_hostname: `settings.ciphers` is now a `TypeSet` internally to handle suppress ordering changes. Schema representation remains the same ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
- resource/cloudflare_custom_hostname: `settings` is now `Optional`/`Computed` to reflect the stricter schema validation introduced in terraform-plugin-sdk v2 ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))
- resource/cloudflare_custom_hostname: `status` is now `Computed` as the value isn't managed by an end user ([#1188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1188))

## 2.27.0 (September 20th, 2021)

NOTES:

- provider: Update to cloudflare-go v0.22.0 ([#1184](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1184))

FEATURES:

- **New Resource:** `cloudflare_access_keys_configuration` ([#1186](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1186))
- **New Resource:** `cloudflare_teams_account` ([#1173](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1173))
- **New Resource:** `cloudflare_teams_rule` ([#1173](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1173))

ENHANCEMENTS:

- resource/cloudflare_access_policy: add support for purpose justification and approvals ([#1199](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1199))
- resource/cloudflare_ruleset: add support for HTTP rate limiting ([#1179](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1179))
- resource/cloudflare_ruleset: add support for Transform Rules ([#1169](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1169))
- resource/cloudflare_ruleset: add support for WAF payload logging ([#1174](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1174))
- resource/cloudflare_ruleset: add support for more complex skip ruleset configurations ([#1201](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1201))

BUG FIXES:

- resource/cloudflare_ruleset: fix state handling for terraform-plugin-sdk v2 ([#1183](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1183))
- resource/cloudflare_zone_settings_override: remap `zero_rtt` => `0rtt` for resource delete ([#1175](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1175))

## 2.26.1 (August 30th, 2021)

**Fixes**

- `resource/cloudflare_ruleset`: Send a single payload for rules instead of many individual payloads to prevent overwriting previous rules ([#1171](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1171))

## 2.26.0 (August 27th, 2021)

- **New resource**: `cloudflare_notification_policy` ([#1138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1138))
- **New resource**: `cloudflare_notification_policy_webhooks` ([#1151](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1151))
- **New resource**: `cloudflare_ruleset` ([#1143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1143))
- **New resource**: `cloudflare_teams_location` ([#1154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1154))
- **New datasource**: `cloudflare_origin_ca_root_certificate` ([#1158](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1158))

**Improvements**

- `resource/cloudflare_waiting_room`: Add support for `json_response_enabled` as an argument ([#1122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1122))

## 2.25.0 (August 4th, 2021)

**Improvements**

- `resource/cloudflare_access_device_posture_rule`: Add support for `domain_joined`, `firewall`, `os_version`, and `disk_encryption` ([#1137](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1137))
- provider: bump `cloudflare-go` to v0.20.0 ([#1146](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1146))

## 2.24.0 (July 19th, 2021)

**Improvements**

- `resource/cloudflare_logpush_job`: Add support for `"nel_reports"` as a dataset ([#1122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1122))
- `resource/cloudflare_custom_hostname`: Allow SSL options to be optional when not required ([#1131](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1131))
- `resource/cloudflare_access_identity_provider`: Support optional Okta API token ([#1119](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1119))
- `resource/cloudflare_load_balancer_pool`: Add support for load shedding ([#1108](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1108))
- `resource/cloudflare_load_balancer_pool`: Add support for longitude and latitude ([#1093](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1093))

**Fixes**

- `resource/cloudflare_record`: Use correct `Import` method on resource ([#1116](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1116))
- `resource/cloudflare_worker_cron_trigger`: Account for deletion of scripts and force a refresh of triggers ([#1121](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1121))
- `resource/cloudflare_rate_limit`: Handle `origin_traffic` missing from API response ([#1125](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1125))
- `resource/cloudflare_record`: Support `allow_overwrite` for root records ([#1129](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1129))

## 2.23.0 (June 30th, 2021)

- **New resource**: `cloudflare_waiting_room` ([#1053](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1053))

**Improvements**

- `datasource/cloudflare_waf_rules`: Export `default_mode` as an attribute ([#1079](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1079))

**Fixes**

- `resource/cloudflare_access_application`: Revert removal of schema changes causing existing applications unable to re-apply ([#1118](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1118))

## 2.22.0 (June 25th, 2021)

- **New resource**: `cloudflare_static_route` ([#1098](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1098))

**Improvements**

- `resource/cloudflare_origin_ca`: Ignore decreasing `requested_validity` ([#1043](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1078))
- `resource/waf_override`: Allow `rules` to be optional ([#1090](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1090))
- `resource/cloudflare_zone`: Don't attempt to set free zone rate plans as that is already the default ([#1102](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1102))
- `resource/cloudflare_access_application`: Ability to set `type` for Applications ([#1076](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1076))
- `resource/cloudflare_zone_lockdown`: Update documentation to show examples of multiple configurations ([#1106](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1106))

## 2.21.0 (May 26th, 2021)

- **New resource**: `cloudflare_device_posture_rule` ([#1058](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1058))
- **New resource**: `cloudflare_teams_list` ([#1058](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1058))

**Improvements**

- provider: Update to terraform-plugin-sdk v1.17.1 ([#1035](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1035), [#1043](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1043))
- `resource/cloudflare_logpush_job`: Allow `ownership_challenge` to be optional to account for Datadog, Splunk or S3-Compatible endpoints ([#1048](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1048))
- `resource/cloudflare_access_group`: Add support for `login_method` ([#1066](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1066))
- `resource/cloudflare_load_balancer`: Add support for `promixity` based steering ([#1072](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1072))
- `resource/cloudflare_access_application`: Prevent bad CORS configuration when credentials and all origins are permitted ([#1073](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1073))
- `resource/cloudflare_access_service_tokens`: Allow configuration to manage automatic renewal when the threshold is crossed and Terraform operations are performed within the window ([#1057](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1057))
- `resource/cloudflare_load_balancer_pool`: Allow support for `Host` header settings ([#1042](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1042))

**Fixes**

- `resource/cloudflare_access_policy`: Allow empty slices in blocks when building policies ([#1034](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1034))
- `resource/cloudflare_load_balancer`: Fix `override` attributes `pop_pools` and `region_pools` referencing incorrect values causing a panic ([#1039](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1039))

## 2.20.0 (April 15th, 2021)

**New resource**: `cloudflare_access_ca_certificate` ([#995](https://github.com/cloudflare/terraform-provider-cloudflare/issues/995))

**Improvements**

- `resource/cloudflare_access_application`: Improve documentation for `Import` usage ([#1002](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1002))
- `resource/cloudflare_logpush_job`: Update documentation to reflect requirements for `destination_conf` to match across all uses ([#1024](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1024))
- `resource/cloudflare_custom_hostname_fallback`: Better handle service lag when updating existing resources by attempting retries ([#1014](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1014))
- `resource/cloudflare_waf_group`: Simplify error handling using inbuilt helpers ([#1015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1015))
- `resource/cloudflare_waf_rule`: Simplify error handling using inbuilt helpers ([#1015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1015))
- `resource/cloudflare_waf_package`: Simplify error handling using inbuilt helpers ([#1015](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1015))
- `resource/cloudflare_access_group`: Add support for `login_method` ([#1018](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1018))
- provider: Update to cloudflare-go v0.16.0 ([#1018](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1018))
- provider: Update to terraform-plugin-sdk v1.16.1 ([#1003](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1003))
- `resource/cloudflare_load_balancer`: Add support for `rules` ([#1016](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1016))

## 2.19.2 (March 15th, 2021)

**Fixes**

- `resource/cloudflare_record`: Address regression from 2.19.1 by checking the API response instead of the schema output for `Priority` ([#992](https://github.com/cloudflare/terraform-provider-cloudflare/issues/992))

## 2.19.1 (March 11th, 2021)

**Fixes**

- `resource/cloudflare_record`: Update `Priority` handling for MX parked records ([#986](https://github.com/cloudflare/terraform-provider-cloudflare/issues/986))

## 2.19.0 (March 10th, 2021)

**Fixes**

- `resource/cloudflare_access_group`: Fix crash when constructing a GSuite group ([#940](https://github.com/cloudflare/terraform-provider-cloudflare/issues/940))
- `resource/cloudflare_access_policy`: Make `precedence` required ([#941](https://github.com/cloudflare/terraform-provider-cloudflare/issues/941))
- `resource/cloudflare_access_group`: Fix crash when constructing a SAML group ([#948](https://github.com/cloudflare/terraform-provider-cloudflare/issues/948))
- `resource/cloudflare_zone`: Update `Retry` logic to look at an available field for passing conditions ([#973](https://github.com/cloudflare/terraform-provider-cloudflare/issues/973))
- `resource/cloudflare_page_rule`: Allow ignoring/including all query string parameters for `cache_key_fields` ([#975](https://github.com/cloudflare/terraform-provider-cloudflare/issues/975))

**Improvements**

- `resource/cloudflare_access_policy`: Enable zone and account level resources to be imported ([#956](https://github.com/cloudflare/terraform-provider-cloudflare/issues/956))
- `resource/cloudflare_origin_ca_certificate`: Smoother import process with less recreation ([#955](https://github.com/cloudflare/terraform-provider-cloudflare/issues/955))
- provider: Update internals to match `cloudflare-go` 0.14 for better error handling and context aware methods ([#976](https://github.com/cloudflare/terraform-provider-cloudflare/issues/976))

## 2.18.0 (February 3rd, 2021)

- **New Resource:** `cloudflare_argo_tunnel` ([#905](https://github.com/cloudflare/terraform-provider-cloudflare/issues/905))
- **New Resource:** `cloudflare_worker_cron_trigger` ([#926](https://github.com/cloudflare/terraform-provider-cloudflare/issues/926))

**Fixes**

- `datasource/cloudflare_zones`: Pagination is now correctly handled internally and will return more than the single page of results ([cloudflare/cloudflare-go#534](https://github.com/cloudflare/cloudflare-go/pull/534)).
- `resource/cloudflare_access_policy`: Correctly handle transforming API responses to schema ([#917](https://github.com/cloudflare/terraform-provider-cloudflare/issues/917))
- `resource/cloudflare_access_group`: Correctly handle transforming API responses to schema ([#918](https://github.com/cloudflare/terraform-provider-cloudflare/issues/918))
- `resource/cloudflare_ip_list`: Ensure account ID is persisted during `Import` ([#916](https://github.com/cloudflare/terraform-provider-cloudflare/issues/916))

**Improvements**

- `resource/cloudflare_access_application`: Allow any `session_duration` that is `time.ParseDuration` compatible ([#910](https://github.com/cloudflare/terraform-provider-cloudflare/issues/910))
- `resource/cloudflare_rate_limit`: Add the ability to configure `match.response.headers` in rate limits ([#911](https://github.com/cloudflare/terraform-provider-cloudflare/issues/911))
- `resource/cloudflare_access_rule`: Validate IP masks within schema ([#921](https://github.com/cloudflare/terraform-provider-cloudflare/issues/921))

## 2.17.0 (January 5th, 2021)

- **New Resource:** `cloudflare_magic_firewall_ruleset` ([#884](https://github.com/cloudflare/terraform-provider-cloudflare/issues/884))

**Fixes**

- `resource/cloudfare_api_token`: Omitting `conditions` will no longer send empty arrays causing IP restriction issues and unusable tokens ([#902](https://github.com/cloudflare/terraform-provider-cloudflare/pull/902))

## 2.16.0 (January 5th, 2021)

**Improvements**

- `resource/cloudflare_access_application`: Add support for `custom_deny_message` and `custom_deny_url` values ([#895](https://github.com/cloudflare/terraform-provider-cloudflare/issues/895))
- `resource/cloudflare_load_balancer_monitor`: Add support for `probe_zone` for monitors ([#903](https://github.com/cloudflare/terraform-provider-cloudflare/issues/903))

## 2.15.0 (December 29th, 2020)

**Improvements**

- `resource/cloudflare_load_balancer`: Add support for `session_affinity_ttl` ([#882](https://github.com/cloudflare/terraform-provider-cloudflare/issues/882))
- `resource/cloudflare_load_balancer`: Add support for `session_affinity_attributes` ([#883](https://github.com/cloudflare/terraform-provider-cloudflare/issues/883))

**Fixes**

- `resource/cloudflare_page_rule`: Fixed crash during update when using custom cache key ([#894](https://github.com/cloudflare/terraform-provider-cloudflare/pull/894))

## 2.14.0 (November 26th, 2020)

- **New Resource:** `cloudflare_api_token` ([#862](https://github.com/cloudflare/terraform-provider-cloudflare/issues/862))
- **New Datasource:** `cloudflare_api_token_permission_groups` ([#862](https://github.com/cloudflare/terraform-provider-cloudflare/issues/862))
- **New Resource:** `cloudflare_zone_dnssec` ([#852](https://github.com/cloudflare/terraform-provider-cloudflare/issues/852))
- **New Datasource:** `cloudflare_zone_dnssec` ([#852](https://github.com/cloudflare/terraform-provider-cloudflare/issues/852))

**Improvements**

- `resource/cloudflare_record`: Add explicit fields for CAA records instead of relying on the map value ([#866](https://github.com/cloudflare/terraform-provider-cloudflare/issues/866))
- `resource/cloudflare_account_member`: Swap schema `role_ids` to `TypeSet` to better handle internal ordering changes ([#876](https://github.com/cloudflare/terraform-provider-cloudflare/issues/876))

**Fixes**

- `datasource/cloudflare_waf_groups`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))
- `datasource/cloudflare_waf_packages`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))
- `datasource/cloudflare_waf_rules`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))
- `datasource/cloudflare_zones`: Make `d.Id()` a consistent string value to prevent Terraform thinking it requires an update ([#869](https://github.com/cloudflare/terraform-provider-cloudflare/issues/869))

## 2.13.2 (November 6th, 2020)

**Fixes**

- `resource/cloudflare_filter`: Remove schema based validation for filters ([#863](https://github.com/cloudflare/terraform-provider-cloudflare/issues/863))

## 2.13.1 (November 5th, 2020)

**Improvements**

- `resource/cloudflare_filter`: Pass missing credential error through to end user ([#860](https://github.com/cloudflare/terraform-provider-cloudflare/issues/860))

## 2.13.0 (November 5th, 2020)

**Improvements**

- `datasource/cloudflare_ip_ranges`: Add the ability to query `china_ipv4_cidr_blocks` and `china_ipv6_cidr_blocks` ([#833](https://github.com/cloudflare/terraform-provider-cloudflare/issues/833))
- `resource/cloudflare_filter`: Improve validation of expressions using the schema ([#848](https://github.com/cloudflare/terraform-provider-cloudflare/issues/848))

**Fixes**

- `resource/cloudflare_page_rule`: Set default for `cache_key_fields.host.resolved` to prevent panics ([#832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/832))
- `resource/cloudflare_authenticated_origin_pulls`: Fix off-by-one error check in `Import` ([#832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/859))
- `resource/cloudflare_authenticated_origin_pulls_certificate`: Fix off-by-one error check in `Import` ([#832](https://github.com/cloudflare/terraform-provider-cloudflare/issues/859))

## 2.12.0 (October 22nd, 2020)

**Improvements**

- `resource/cloudflare_certificate_pack`: Swap internal representation of `hosts` to remove inconsistent ordering issues ([#800](https://github.com/cloudflare/terraform-provider-cloudflare/issues/800))
- `resource/cloudflare_logpush_job`: Handle deletion outside of Terraform ([#798](https://github.com/cloudflare/terraform-provider-cloudflare/issues/798))
- `resource/cloudflare_access_group`: Add support for `geo` conditionals ([#803](https://github.com/cloudflare/terraform-provider-cloudflare/issues/803))
- `resource/cloudflare_access_application`: Add support for `enable_binding_cookie` ([#802](https://github.com/cloudflare/terraform-provider-cloudflare/issues/802))
- `resource/cloudflare_waf_rule`: Improve documentation for `mode` ([#824](https://github.com/cloudflare/terraform-provider-cloudflare/issues/824))
- `datasource/cloudflare_waf_rule`: Improve documentation for `mode` ([#824](https://github.com/cloudflare/terraform-provider-cloudflare/issues/824))
- `resource/cloudflare_access_application`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))
- `resource/cloudflare_access_group`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))
- `resource/cloudflare_access_identity_provider`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))
- `resource/cloudflare_access_policy`: Add support for zone-level routes to Access resources ([#819](https://github.com/cloudflare/terraform-provider-cloudflare/issues/819))

**Fixes**

- `resource/cloudflare_custom_hostname_fallback_origin`: Don't retry the "active" status of custom hostnames fallbacks ([#818](https://github.com/cloudflare/terraform-provider-cloudflare/issues/818))
- `resource/cloudflare_zone`: Remove `DiffSuppressFunc` causing `jump_start` issues ([#830](https://github.com/cloudflare/terraform-provider-cloudflare/issues/830))

## 2.11.0 (September 11th, 2020)

- **New Resource:** `cloudflare_certificate_pack` ([#778](https://github.com/cloudflare/terraform-provider-cloudflare/issues/778))

**Improvements**

- `resource/cloudflare_access_group`: Add support for `auth_method` ([#762](https://github.com/cloudflare/terraform-provider-cloudflare/issues/762))
- `resource/cloudflare_access_group`: De-duplicate blocks in groups by accepting lists instead ([#739](https://github.com/cloudflare/terraform-provider-cloudflare/issues/739))
- `resource/cloudflare_worker_script`: Adds support for `webassembly_binding` ([#780](https://github.com/cloudflare/terraform-provider-cloudflare/issues/780))
- `resource/cloudflare_healthcheck`: Retry hostname resolution errors when encountering "no such host" responses ([#789](https://github.com/cloudflare/terraform-provider-cloudflare/issues/789))
- `resource/cloudflare_access_application`: Better validation for allowed methods and origin combinations to prevent getting state into an unrecoverable state ([#793](https://github.com/cloudflare/terraform-provider-cloudflare/issues/793))

**Fixes**

- `resource/cloudflare_healthcheck`: Handle resource deletion outside of Terraform ([#787](https://github.com/cloudflare/terraform-provider-cloudflare/issues/787))
- `resource/cloudflare_custom_hostname`: Ensure `Import` sets hostname to prevent recreation ([#788](https://github.com/cloudflare/terraform-provider-cloudflare/issues/788))
- `resource/cloudflare_ip_list`: Handle resource deletion outside of Terraform ([#794](https://github.com/cloudflare/terraform-provider-cloudflare/issues/794))
- `resource/cloudflare_ip_list`: Remove `item`.`id` from schema ([#796](https://github.com/cloudflare/terraform-provider-cloudflare/issues/796))

## 2.10.1 (August 24th, 2020)

**Fixes**

- `resource/cloudflare_access_application`: Handle the `zone_id` => `account_id` move internally ([#724](https://github.com/cloudflare/terraform-provider-cloudflare/issues/724))

## 2.10.0 (August 24th, 2020)

- **New Resource:** `cloudflare_custom_hostname_origin_fallback` ([#757](https://github.com/cloudflare/terraform-provider-cloudflare/issues/757))
- **New Resource:** `cloudflare_authenticated_origin_pulls` ([#749](https://github.com/cloudflare/terraform-provider-cloudflare/issues/749))
- **New Resource:** `cloudflare_authenticated_origin_pulls_certificate` ([#749](https://github.com/cloudflare/terraform-provider-cloudflare/issues/749))
- **New Resource:** `cloudflare_ip_list` ([#766](https://github.com/cloudflare/terraform-provider-cloudflare/issues/766))

**Improvements**

- `resource/cloudflare_spectrum_application`: Add support for port ranges ([#745](https://github.com/cloudflare/terraform-provider-cloudflare/issues/745))
- `resource/cloudflare_custom_hostname`: Force creation of a new resource if the `zone_id` value changes ([#761](https://github.com/cloudflare/terraform-provider-cloudflare/issues/761))
- `resource/cloudflare_record`: Retry record creation/update if the response includes an "already exists" exception for handling race conditions ([#773](https://github.com/cloudflare/terraform-provider-cloudflare/issues/773))

**Fixes**

- `resource/cloudflare_firewall_rule`: Compare descriptions after converting unicode + HTML entities to prevent unnecessary diffs ([#758](https://github.com/cloudflare/terraform-provider-cloudflare/issues/758))
- `resource/cloudflare_filter`: Compare descriptions after converting unicode + HTML entities to prevent unnecessary diffs ([#758](https://github.com/cloudflare/terraform-provider-cloudflare/issues/758))

## 2.9.0 (July 30th, 2020)

- **New Resource:** `cloudflare_custom_hostname` (SSL for SaaS) ([#746](https://github.com/cloudflare/terraform-provider-cloudflare/issues/746))

**Improvements**

- `resource/access_application`: Add support for `allowed_idps` and restricting which Identity Providers are associated with an Application ([#734](https://github.com/cloudflare/terraform-provider-cloudflare/issues/734))
- `resource/access_application`: Add support for `auto_redirect_to_identity` ([#730](https://github.com/cloudflare/terraform-provider-cloudflare/issues/730))
- `resource/access_application`: Add CORS support ([#725](https://github.com/cloudflare/terraform-provider-cloudflare/issues/725))
- `resource/cloudflare_custom_ssl`: Allow `geo_restrictions` to be `nil` and not included in the request payload ([#714](https://github.com/cloudflare/terraform-provider-cloudflare/issues/714))
- `datasource/cloudflare_zones`: Filtering is now performed on the server side and the `name` parameter is no longer a regex. Instead, `name` is a string to match on and `match` is a regex. See the website documentation for more examples and updated references ([#708](https://github.com/cloudflare/terraform-provider-cloudflare/issues/708)) in order to make your code compatible with this release.

## 2.8.0 (June 22, 2020)

- **New Resource:** `cloudflare_waf_override` ([#691](https://github.com/cloudflare/terraform-provider-cloudflare/issues/691))

**Improvements**

- `resource/cloudflare_argo`: Allow `tiered_caching` and `smart_routing` to be toggled individually allowing for entitlement differences ([#703](https://github.com/cloudflare/terraform-provider-cloudflare/issues/703))
- `resource/cloudflare_page_rule`: Add support for `cache_ttl_by_status` ([#706](https://github.com/cloudflare/terraform-provider-cloudflare/issues/706))
- `resource/cloudflare_worker_script`: Add support for `plain_text` and `secret_text` bindings ([#710](https://github.com/cloudflare/terraform-provider-cloudflare/issues/710))

**Fixes**

- `resource/cloudflare_record`: Update `TestAccCloudflareRecord_LOC` test asserted value to use less precise floats and match the API responses ([#712](https://github.com/cloudflare/terraform-provider-cloudflare/issues/712))
- `resource/cloudflare_record`: Update `TestAccCloudflareRecord_Basic` test `metadata` attributes to match updated API payload ([#713](https://github.com/cloudflare/terraform-provider-cloudflare/issues/713))

## 2.7.0 (May 20, 2020)

- **New Resource:** `cloudflare_byo_ip_prefix` ([#671](https://github.com/cloudflare/terraform-provider-cloudflare/issues/671))
- **New Resource:** `cloudflare_logpull_retention` ([#678](https://github.com/cloudflare/terraform-provider-cloudflare/issues/678))
- **New Resource:** `cloudflare_healthcheck` ([#680](https://github.com/cloudflare/terraform-provider-cloudflare/issues/680))

**Improvements:**

- `resource/cloudflare_worker_route`: Improve documentation to mention using `account_id` for the underlying APIs ([#669](https://github.com/cloudflare/terraform-provider-cloudflare/issues/669))
- `resource/cloudflare_worker_script`: Improve documentation to mention using `account_id` for the underlying APIs ([#670](https://github.com/cloudflare/terraform-provider-cloudflare/issues/670))
- `resource/cloudflare_load_balancer_pool`: Improve documentation to mention `notification_email` accepts a comma delimited list of emails ([#687](https://github.com/cloudflare/terraform-provider-cloudflare/issues/687))
- `resource/cloudflare_page_rule`: Add support for `cache_key_fields` Page Rule action ([#662](https://github.com/cloudflare/terraform-provider-cloudflare/issues/662))

**Fixes:**

- `resource/cloudflare_zone_settings_override`: Fix regression where if you didn't have universal SSL settings defined, it would error when setting them ([#663](https://github.com/cloudflare/terraform-provider-cloudflare/issues/663))
- `resource/cloudflare_zone`: Handle changing zone rate plan from "free" to "enterprise" ([#668](https://github.com/cloudflare/terraform-provider-cloudflare/issues/668))
- `resource/cloudflare_record`: Update validation to allow PTR records ([9a8fd43](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9a8fd43))

## 2.6.0 (April 22, 2020)

**Improvements:**

- `resource/cloudflare_zone_settings_override`: Add `universal_ssl` to control enablement of Universal SSL on a zone ([#658](https://github.com/cloudflare/terraform-provider-cloudflare/issues/658))
- provider: API keys and API tokens are now validated to help differentiate incorrect usage before making API calls ([#661](https://github.com/cloudflare/terraform-provider-cloudflare/issues/661))
- `resource/cloudflare_logpush_job`: Add support for "firewall_events" dataset parameter ([#660](https://github.com/cloudflare/terraform-provider-cloudflare/issues/660))
- `resource/cloudflare_logpush_job`: Add support for "dataset" parameter ([#649](https://github.com/cloudflare/terraform-provider-cloudflare/issues/649))
- `resource/cloudflare_zone_settings_override`: Remove `edge_cache_ttl` ([#654](https://github.com/cloudflare/terraform-provider-cloudflare/issues/654))
- `resource/cloudflare_access_group`: Allow Access conditions for `include`/`require`/`exclude` to be used consistently between Access Groups and Access Policies ([#646](https://github.com/cloudflare/terraform-provider-cloudflare/issues/646))

**Fixes:**

- `resource/cloudflare_logpush_job`: fix for `strconv.Atoi: parsing ""` error while creating Logpush job

## 2.5.1 (April 03, 2020)

**Improvements:**

- `resource/cloudflare_zone_settings_override`: Update `image_resizing` options to include `"open"` ([#639](https://github.com/cloudflare/terraform-provider-cloudflare/issues/639))

**Fixes:**

- `resource/cloudflare_access_group`: Fixed misspelt Okta in JSON payload ([cloudflare/cloudflare-go#440](https://github.com/cloudflare/cloudflare-go/issues/440))

## 2.5.0 (March 27, 2020)

**Improvements:**

- `resource/cloudflare_access_policy`: Add support for `service_token` and `any_valid_service_token` ([#612](https://github.com/cloudflare/terraform-provider-cloudflare/issues/612))
- `resource/cloudflare_waf_group`: Handle WAF group deletions in the API responses ([#623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/623))
- `resource/cloudflare_waf_package`: Handle WAF package deletions in the API responses ([#623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/623))
- `resource/cloudflare_waf_rule`: Handle WAF rule deletions in the API responses ([#623](https://github.com/cloudflare/terraform-provider-cloudflare/issues/623))
- `resource/cloudflare_access_policy`: Add support for `group` ([#626](https://github.com/cloudflare/terraform-provider-cloudflare/issues/626))
- `resource/cloudflare_firewall_rule`: Add support for bypassing specific `products` ([#630](https://github.com/cloudflare/terraform-provider-cloudflare/issues/630))
- `resource/cloudflare_spectrum_application`: Add support for `edge_ips`, `argo_smart_routing` and `edge_ip_connectivity` ([#631](https://github.com/cloudflare/terraform-provider-cloudflare/issues/631))
- `resource/cloudflare_access_group`: Add support for using external providers (`gsuite`, `github`, `azure`, `okta`, `saml`, `mTLS certificate`, `common name`
  ) ([#633](https://github.com/cloudflare/terraform-provider-cloudflare/issues/633))

## 2.4.1 (March 12, 2020)

**Improvements:**

- `resource/cloudflare_logpush_job`: Support `Import` on the resource ([#618](https://github.com/cloudflare/terraform-provider-cloudflare/issues/618))

**Fixes:**

- `resource/cloudflare_record`: Missing CAA in DNS validation ([#619](https://github.com/cloudflare/terraform-provider-cloudflare/issues/619))

## 2.4.0 (March 09, 2020)

- **New Resource:** `cloudflare_workers_kv` ([#595](https://github.com/cloudflare/terraform-provider-cloudflare/issues/595))
- **New Resource:** `cloudflare_access_identity_provider` ([#597](https://github.com/cloudflare/terraform-provider-cloudflare/issues/597))

**Improvements:**

- `resource/cloudflare_record`: Stricter validation for record types ([#610](https://github.com/cloudflare/terraform-provider-cloudflare/issues/610))
- `resource/logpush_job`: Add more verbose error handling ([#564](https://github.com/cloudflare/terraform-provider-cloudflare/issues/564))
- `resource/zone_settings_override`: Update documentation for `cache_level` values ([#606](https://github.com/cloudflare/terraform-provider-cloudflare/issues/606))
- `resource/access_application`: Add documentation for available attributes ([#587](https://github.com/cloudflare/terraform-provider-cloudflare/issues/587))
- `resource/cloudflare_firewall_rule`: Add support for bypassing security configuration rules by URL ([#568](https://github.com/cloudflare/terraform-provider-cloudflare/issues/568))
- `resource/cloudflare_record_migrate`: Use `zone_id` for state migration before attempting to use `domain` ([#566](https://github.com/cloudflare/terraform-provider-cloudflare/issues/566))
- `resource/cloudflare_load_balancer`: Update `session_affinity` validation to allow `"ip_cookie"` ([#573](https://github.com/cloudflare/terraform-provider-cloudflare/issues/573))
- `datasource/ip_ranges`: Update documentation to show 0.12 syntax ([#617](https://github.com/cloudflare/terraform-provider-cloudflare/issues/617))

**Fixes**

- `resource/zone_settings_override`: Handle individual zone settings within `Delete` operations ([#599](https://github.com/cloudflare/terraform-provider-cloudflare/issues/599))

## 2.3.0 (December 18, 2019)

- **New Resource:** `cloudflare_origin_ca_certificate` ([#547](https://github.com/cloudflare/terraform-provider-cloudflare/issues/547))

**Fixes:**

- `resource/cloudflare_zone_settings_override`: Renamed `0rtt` to `zero_rtt` to conform to HCL grammar requirements ([#557](https://github.com/cloudflare/terraform-provider-cloudflare/issues/557))

**Improvements:**

- `resource/cloudflare_access_rule`: Add `ip6` as valid option ([#560](https://github.com/cloudflare/terraform-provider-cloudflare/issues/560))
- `resource/cloudflare_spectrum_application`: Swap `proxy_protocol` to string field with supporting enum values instead ([#561](https://github.com/cloudflare/terraform-provider-cloudflare/issues/561))
- `resource/cloudflare_waf_rule`: Add `package_id` as valid option and export `group_id` ([#552](https://github.com/cloudflare/terraform-provider-cloudflare/issues/552))

## 2.2.0 (December 05, 2019)

- **New Resource:** `cloudflare_access_group` ([#510](https://github.com/cloudflare/terraform-provider-cloudflare/issues/510))
- **New Resource:** `cloudflare_workers_kv_namespace` ([#443](https://github.com/cloudflare/terraform-provider-cloudflare/issues/443))

**Improvements:**

- `resource/cloudflare_zone_settings_override`: Add `non_identity` to allowed `decision` schema ([#541](https://github.com/cloudflare/terraform-provider-cloudflare/issues/541))
- `resource/cloudflare_zone_settings_override`: Add support for `0rtt` and `http3` settings ([#542](https://github.com/cloudflare/terraform-provider-cloudflare/issues/542))
- `resource/cloudflare_load_balancer_monitor`: Allow empty string for `expected_body` ([#539](https://github.com/cloudflare/terraform-provider-cloudflare/issues/539))
- `resource/cloudflare_worker_script`: Add support for Worker KV Namespace Bindings ([#544](https://github.com/cloudflare/terraform-provider-cloudflare/issues/544))
- `data_source/waf_rules`, `resource/cloudflare_waf_rule`, Support allowed modes for WAF Rules ([#550](https://github.com/cloudflare/terraform-provider-cloudflare/issues/550))

**Fixes:**

- `resource/cloudflare_spectrum_application`: Spectrum origin_port is optional ([#549](https://github.com/cloudflare/terraform-provider-cloudflare/issues/549))

## 2.1.0 (November 07, 2019)

- **New datasource:** `cloudflare_waf_rules` ([#525](https://github.com/cloudflare/terraform-provider-cloudflare/issues/525))

**Improvements:**

- `resource/cloudflare_zone`: Expose `verification_key` for partial setups ([#532](https://github.com/cloudflare/terraform-provider-cloudflare/issues/532))
- `resource/cloudflare_worker_route`: Enable API Tokens support from upstream [cloudflare-go](https://github.com/cloudflare/cloudflare-go) release

## 2.0.1 (October 22, 2019)

- **New Resource:** `cloudflare_access_service_tokens` ([#521](https://github.com/cloudflare/terraform-provider-cloudflare/issues/521))
- **New Resource:** `cloudflare_waf_package` ([#475](https://github.com/cloudflare/terraform-provider-cloudflare/issues/475))
- **New Resource:** `cloudflare_waf_group` ([#476](https://github.com/cloudflare/terraform-provider-cloudflare/issues/476))
- **New datasource:** `cloudflare_waf_groups` ([#508](https://github.com/cloudflare/terraform-provider-cloudflare/issues/508))
- **New datasource:** `cloudflare_waf_packages` ([#509](https://github.com/cloudflare/terraform-provider-cloudflare/issues/509))

**Fixes:**

- `resource/cloudflare_page_rule`: Set `h2_prioritization` individually not via bulk endpoint ([#493](https://github.com/cloudflare/terraform-provider-cloudflare/issues/493))
- `resource/cloudflare_zone_settings_override`: Set `zone_id` to prevent unnecessary re-creation of resources ([#502](https://github.com/cloudflare/terraform-provider-cloudflare/issues/502))

**Improvements:**

- `resource/cloudflare_spectrum_application`: Add support for setting `traffic_type` ([#481](https://github.com/cloudflare/terraform-provider-cloudflare/issues/481))
- `resource/cloudflare_zone_settings_override`: Update documentation with default values ([#498](https://github.com/cloudflare/terraform-provider-cloudflare/issues/498))

**Internals:**

- Migrated to Terraform plugin SDK ([#489](https://github.com/cloudflare/terraform-provider-cloudflare/issues/489))

## 2.0.0 (September 30, 2019)

**Breaking changes:**

- `provider/cloudflare`:
- renamed `token` to `api_key`
- renamed `org_id` to `account_id`
- removed `use_org_from_zone`, you need to explicitly specify `account_id`
- Environment variables:
- renamed `CLOUDFLARE_TOKEN` to `CLOUDFLARE_API_TOKEN`
- renamed `CLOUDFLARE_ORG_ID` to `CLOUDFLARE_ACCOUNT_ID`
- removed `CLOUDFLARE_ORG_ZONE`, you need to explicitly specify `CLOUDFLARE_ACCOUNT_ID`
- Changed the following resources to require Zone ID:
- `cloudflare_access_rule`
- `cloudflare_filter`
- `cloudflare_firewall_rule`
- `cloudflare_load_balancer`
- `cloudflare_page_rule`
- `cloudflare_rate_limit`
- `cloudflare_record`
- `cloudflare_waf_rule`
- `cloudflare_worker_route"`
- `cloudflare_zone_lockdown`
- `cloudflare_zone_settings_override`
- Workers single-script support removed

Please see [Version 2 Upgrade Guide](https://www.terraform.io/docs/providers/cloudflare/guides/version-2-upgrade.html) for details.

**Improvements:**

- `cloudflare/resource_cloudflare_argo`: Handle errors when fetching tiered caching + smart routing settings ([#477](https://github.com/cloudflare/terraform-provider-cloudflare/issues/477))
- Various documentation updates for 0.12 syntax

## 1.18.1 (August 29, 2019)

**Fixes:**

- `resource/cloudflare_load_balancer`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
- `resource/cloudflare_page_rule`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
- `resource/cloudflare_rate_limit`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
- `resource/cloudflare_waf_rule`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
- `resource/cloudflare_worker_route`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
- `resource/cloudflare_worker_script`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))
- `resource/cloudflare_zone_lockdown`: Mark `zone` as Computed to allow deprecations ([#462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/462))

## 1.18.0 (August 27, 2019)

**Fixes:**

- `resource/cloudflare_page_rule`: Fix a logic condition where setting `edge_cache_ttl` action but then not updating it in subsequent `apply` runs causes it to be blown away ([#453](https://github.com/cloudflare/terraform-provider-cloudflare/issues/453))

**Improvements:**

- provider: You can now use API tokens to authenticate instead of user email and key ([#450](https://github.com/cloudflare/terraform-provider-cloudflare/issues/450))
- `resource/cloudflare_zone_lockdown`: `priority` can now be set on the resource ([#445](https://github.com/cloudflare/terraform-provider-cloudflare/issues/445))
- `resource/cloudflare_custom_ssl`: Updated website documentation navigation to include link for resource ([#442)](https://github.com/cloudflare/terraform-provider-cloudflare/issues/442))

**Deprecations:**

- `resource/cloudflare_access_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_filter`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_firewall_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_load_balancer`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_page_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_rate_limit`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_waf_rule`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_worker_route`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_worker_script`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))
- `resource/cloudflare_zone_lockdown`: `zone` has been superseded by using `zone_id` ([#452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/452))

## 1.17.1 (August 09, 2019)

**Fixes:**

- Partially revert [[#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421)] deprecation messages

## 1.17.0 (August 09, 2019)

**Removals:**

- `resource/cloudflare_zone_settings_override`: `sha1_support` has been removed due to Cloudflare no longer supporting SHA1 certificates or the API endpoint ([#415](https://github.com/cloudflare/terraform-provider-cloudflare/issues/415))

**Deprecations:**

- `resource/cloudflare_zone_settings_override`: `tls_1_2_only` has been superseded by using `min_tls_version` instead ([#405](https://github.com/cloudflare/terraform-provider-cloudflare/issues/405))
- `resource/cloudflare_access_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_filter`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_firewall_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_load_balancer`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_page_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_rate_limit`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_waf_rule`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_worker_route`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_worker_script`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))
- `resource/cloudflare_zone_lockdown`: `zone` has been superseded by using `zone_id` ([#421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/421))

**Improvements:**

- **New Resource:** `cloudflare_custom_ssl` ([#418](https://github.com/cloudflare/terraform-provider-cloudflare/issues/418))
- `resource/cloudflare_filter`: Strip all surrounding whitespace from filter expressions to match API responses ([#361](https://github.com/cloudflare/terraform-provider-cloudflare/issues/361))
- `resource/cloudflare_zone`: Support unicode zone name values ([#412](https://github.com/cloudflare/terraform-provider-cloudflare/issues/412))
- `resource/cloudflare_page_rule`: Allow setting `origin_pull` for SSL ([#430](https://github.com/cloudflare/terraform-provider-cloudflare/issues/430))
- `resource/cloudflare_load_balancer_monitor`: Add TCP support for load balancer monitor ([#428](https://github.com/cloudflare/terraform-provider-cloudflare/issues/428))

**Fixes:**

- `resource/cloudflare_logpush_job`: Update documentation ([#395](https://github.com/cloudflare/terraform-provider-cloudflare/issues/395))
- `resource/cloudflare_zone_lockdown`: Fix: examples in documentation ([#407](https://github.com/cloudflare/terraform-provider-cloudflare/issues/407))
- `resource/cloudflare_page_rule`: Set nil on changed string-based Page Rule actions

## 1.16.1 (June 27, 2019)

**Fixes:**

- `resource/cloudflare_page_rule`: Fix regression in `browser_cache_ttl` where the value was sent as a string instead of an integer to the remote ([#390](https://github.com/cloudflare/terraform-provider-cloudflare/issues/390))

## 1.16.0 (June 20, 2019)

**Improvements:**

- `resource/cloudflare_zone_settings_override`: Add support for `h2_prioritization` and `image_resizing` ([#381](https://github.com/cloudflare/terraform-provider-cloudflare/issues/381))
- `resource/cloudflare_load_balancer_pool`: Update IP range for tests to not use reserved ranges ([#369](https://github.com/cloudflare/terraform-provider-cloudflare/issues/369))

**Fixes:**

- `resource/cloudflare_page_rule`: Fix issues with `browser_cache_ttl` defaults and when value is `0` (for Enterprise users) ([#379](https://github.com/cloudflare/terraform-provider-cloudflare/issues/379))

## 1.15.0 (May 24, 2019)

- The provider is now compatible with Terraform v0.12, while retaining compatibility with prior versions. ([#309](https://github.com/cloudflare/terraform-provider-cloudflare/issues/309))

## 1.14.0 (May 15, 2019)

**Improvements:**

- **New Resource:** `cloudflare_argo` Manage Argo features ([#304](https://github.com/cloudflare/terraform-provider-cloudflare/issues/304))
- `cloudflare_zone`: Support management of partial zones ([#303](https://github.com/cloudflare/terraform-provider-cloudflare/issues/303))
- `cloudflare_rate_limit`: Update `modes` documentation ([#293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/212))
- `cloudflare_load_balancer`: Allow steering policy of "random" ([#329](https://github.com/cloudflare/terraform-provider-cloudflare/issues/329))

**Fixes:**

- `cloudflare_page_rule` - Allow setting `browser_cache_ttl` to 0 ([#293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/291))
- `cloudflare_page_rule` - Swap to completely replacing rules ([#338](https://github.com/cloudflare/terraform-provider-cloudflare/issues/338))

## 1.13.0 (April 12, 2019)

**Improvements**

- **New Resource:** `cloudflare_logpush_job` ([#287](https://github.com/cloudflare/terraform-provider-cloudflare/issues/287))
- `cloudflare_zone_settings` - Remove option to toggle `always_on_ddos` ([#253](https://github.com/cloudflare/terraform-provider-cloudflare/issues/253))
- `cloudflare_page_rule` - Update documentation to clarify "0" usage
- `cloudflare_zones` - Return zone ID and zone name ([#275](https://github.com/cloudflare/terraform-provider-cloudflare/issues/275))
- `cloudflare_load_balancer` - Add `enabled` field ([#208](https://github.com/cloudflare/terraform-provider-cloudflare/issues/208))
- `cloudflare_record` - validators: Allow PTR DNS records ([#283](https://github.com/cloudflare/terraform-provider-cloudflare/issues/283))

**Fixes:**

- `cloudflare_custom_pages` - Use correct casing for `zone_id` lookups
- `cloudflare_rate_limit` - Make `correlate` optional and not flap in state management ([#271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/271))
- `cloudflare_spectrum_application` - Fixed integration tests to work ([#275](https://github.com/cloudflare/terraform-provider-cloudflare/issues/275))
- `cloudflare_page_rule` - Better track field changes in `actions` resource. ([#107](https://github.com/cloudflare/terraform-provider-cloudflare/issues/107))

## 1.12.0 (March 07, 2019)

**Improvements:**

- provider: Enable request/response logging ([#212](https://github.com/cloudflare/terraform-provider-cloudflare/issues/212))
- resource/cloudflare_load_balancer_monitor: Add validation for `port` ([#213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/213))
- resource/cloudflare_load_balancer_monitor: Add `allow_insecure` and `follow_redirects` ([#205](https://github.com/cloudflare/terraform-provider-cloudflare/issues/205))
- resource/cloudflare_page_rule: Updated available actions documentation to match what is available ([#228](https://github.com/cloudflare/terraform-provider-cloudflare/issues/228))
- provider: Swap to using go modules for dependency management ([#230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/230))
- provider: Minimum Go version for development is now 1.11 ([#230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/230))

**Fixes:**

- resource/cloudflare_record: Read `data` back from API correctly ([#217](https://github.com/cloudflare/terraform-provider-cloudflare/issues/217))
- resource/cloudflare_rate_limit: Read `correlate` back from API correctly ([#204](https://github.com/cloudflare/terraform-provider-cloudflare/issues/204))
- resource/cloudflare_load_balancer_monitor: Fix incorrect type cast for `port` ([#213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/213))
- resource/cloudflare_load_balancer: Make `steering_policy` computed to avoid spurious diffs ([#214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/214))
- resource/cloudflare_load_balancer: Read `session_affinity` back from API to make import work & detects drifts ([#214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/214))

## 1.11.0 (January 11, 2019)

**Improvements:**

- **New Resource:** `cloudflare_spectrum_app` ([#156](https://github.com/cloudflare/terraform-provider-cloudflare/issues/156))
- **New Data Source:** `cloudflare_zones` ([#168](https://github.com/cloudflare/terraform-provider-cloudflare/issues/168))
- `cloudflare_load_balancer_monitor` - Add optional `port` parameter ([#179](https://github.com/cloudflare/terraform-provider-cloudflare/issues/179))
- `cloudflare_page_rule` - Improved documentation for `priority` attribute ([#182](https://github.com/cloudflare/terraform-provider-cloudflare/issues/182)], missing `explicit_cache_control` [[#185](https://github.com/cloudflare/terraform-provider-cloudflare/issues/185))
- `cloudflare_rate_limit` - Add `challenge` and `js_challenge` rate-limit modes ([#172](https://github.com/cloudflare/terraform-provider-cloudflare/issues/172))

**Fixes:**

- `cloudflare_page_rule` - Page rule `zone` attribute change to trigger new resource ([#183](https://github.com/cloudflare/terraform-provider-cloudflare/issues/183))

## 1.10.0 (December 18, 2018)

**Improvements:**

- `cloudflare_zone_settings_override` - Add `opportunistic_onion` zone setting support ([#170](https://github.com/cloudflare/terraform-provider-cloudflare/issues/170))
- `cloudflare_zone` - Add ability to set zone plan ([#160](https://github.com/cloudflare/terraform-provider-cloudflare/issues/160))

**Fixes:**

- `cloudflare_zone` - Allow zones to be properly imported ([#157](https://github.com/cloudflare/terraform-provider-cloudflare/issues/157))
- `cloudflare_access_policy` - Match access_policy argument requisites with reality ([#158](https://github.com/cloudflare/terraform-provider-cloudflare/issues/158))
- `cloudflare_filter` - Allow `zone_id` to set `zone` and vice versa ([#162](https://github.com/cloudflare/terraform-provider-cloudflare/issues/162))
- `cloudflare_firewall_rule` - Allow `zone_id` to set `zone` and vice versa ([#174](https://github.com/cloudflare/terraform-provider-cloudflare/issues/174))
- `cloudflare_access_rule` - Ensure `zone` and `zone_id` are always set ([#175](https://github.com/cloudflare/terraform-provider-cloudflare/issues/175))
- Minor documentation fixes

## 1.9.0 (November 15, 2018)

**Improvements:**

- **New Resource:** `cloudflare_access_application` ([#145](https://github.com/cloudflare/terraform-provider-cloudflare/issues/145))
- **New Resource:** `cloudflare_access_policy` ([#145](https://github.com/cloudflare/terraform-provider-cloudflare/issues/145))
- `cloudflare_load_balancer` - Add steering policy support ([#147](https://github.com/cloudflare/terraform-provider-cloudflare/issues/147))
- `cloudflare_load_balancer` - Support `session_affinity` ([#153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/153))
- `cloudflare_load_balancer_pool` - Support `weight` ([#153](https://github.com/cloudflare/terraform-provider-cloudflare/issues/153))

**Fixes:**

- `cloudflare_record` - Compare name without the zone name ([#151](https://github.com/cloudflare/terraform-provider-cloudflare/issues/151))
- Minor documentation fixes ([#149](https://github.com/cloudflare/terraform-provider-cloudflare/issues/149)] [[#152](https://github.com/cloudflare/terraform-provider-cloudflare/issues/152))

## 1.8.0 (November 05, 2018)

**Improvements:**

- **New Resource:** `cloudflare_zone` ([#58](https://github.com/cloudflare/terraform-provider-cloudflare/issues/58))
- **New Resource:** `cloudflare_custom_pages` ([#132](https://github.com/cloudflare/terraform-provider-cloudflare/issues/132))
- `cloudflare_zone_settings_override` - Allow setting SSL level to Strict (SSL-Only Origin Pull) ([#122](https://github.com/cloudflare/terraform-provider-cloudflare/issues/122))
- Update provider usage/build docs and how to update a dependency ([#138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/138))
- Improve `Building The Provider` instructions ([#143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/143))
- `cloudflare_access_rule` - Make importable for all rule types ([#141](https://github.com/cloudflare/terraform-provider-cloudflare/issues/141))
- `cloudflare_load_balancer_pool` - Implement `Update` ([#140](https://github.com/cloudflare/terraform-provider-cloudflare/issues/140))

**Fixes:**

- `cloudflare_rate_limit` - Documentation fixes for markdown where \_ALL\_ is italicized ([#125](https://github.com/cloudflare/terraform-provider-cloudflare/issues/125))
- `cloudflare_worker_route` - Correctly set `multi_script` on Enterprise worker imports ([#124](https://github.com/cloudflare/terraform-provider-cloudflare/issues/124))
- `account_member` - Ignore role ID ordering ([#128](https://github.com/cloudflare/terraform-provider-cloudflare/issues/128))
- `cloudflare_rate_limit` - Origin traffic isn't default anymore ([#130](https://github.com/cloudflare/terraform-provider-cloudflare/issues/130))
- `cloudflare_rate_limit` - Update rate limit validation to allow `1` ([#129](https://github.com/cloudflare/terraform-provider-cloudflare/issues/129))
- `cloudflare_record` - Add validation to ensure TTL is not set while `proxied` is true ([#127](https://github.com/cloudflare/terraform-provider-cloudflare/issues/127))
- Updated code for provider version in User-Agent
- `cloudflare_zone_lockdown` - Fix import of zone lockdowns ([#135](https://github.com/cloudflare/terraform-provider-cloudflare/issues/135))

## 1.7.0 (October 09, 2018)

**Improvements:**

- **New Resource:** `cloudflare_account_member` ([#78](https://github.com/cloudflare/terraform-provider-cloudflare/issues/78))

## 1.6.0 (October 05, 2018)

**Improvements:**

- **New Resource:** `cloudflare_filter`
- **New Resource:** `cloudflare_firewall_rule`

## 1.5.0 (September 21, 2018)

**Improvements:**

- **New Resource:** `cloudflare_zone_lockdown` ([#115](https://github.com/cloudflare/terraform-provider-cloudflare/issues/115))

**Fixes:**

- Send User-Agent header with name and version when contacting API
- `cloudflare_page_rule` - Fix page rule polish (off, lossless or lossy) ([#116](https://github.com/cloudflare/terraform-provider-cloudflare/issues/116))

## 1.4.0 (September 11, 2018)

**Improvements:**

- **New Resource:** `cloudflare_worker_route` ([#110](https://github.com/cloudflare/terraform-provider-cloudflare/issues/110))
- **New Resource:** `cloudflare_worker_script` ([#110](https://github.com/cloudflare/terraform-provider-cloudflare/issues/110))

## 1.3.0 (September 04, 2018)

**Improvements:**

- **New Resource:** `cloudflare_access_rule` ([#64](https://github.com/cloudflare/terraform-provider-cloudflare/issues/64))

**Fixes:**

- `cloudflare_zone_settings_override` - Change Zone Settings Override to use GetOkExists ([#107](https://github.com/cloudflare/terraform-provider-cloudflare/issues/107))

## 1.2.0 (August 13, 2018)

**Improvements:**

- **New Resource:** `cloudflare_waf_rule` ([#98](https://github.com/cloudflare/terraform-provider-cloudflare/issues/98))
- `cloudflare_zone_settings_override` - Add `off` as Security Level setting ([#99](https://github.com/cloudflare/terraform-provider-cloudflare/issues/99))
- `resource_cloudflare_rate_limit` - Add nat support ([#96](https://github.com/cloudflare/terraform-provider-cloudflare/issues/96))
- `resource_cloudflare_zone_settings_override` - Add `zrt` as a value for the `tls_1_3` setting ([#106](https://github.com/cloudflare/terraform-provider-cloudflare/issues/106))
- Minor documentation improvements

**Fixes:**

- `cloudflare_record` - Setting a DNS record's `proxied` flag to false stopped working ([#103](https://github.com/cloudflare/terraform-provider-cloudflare/issues/103))

## 1.1.0 (July 25, 2018)

FIXES:

- `cloudflare_ip_ranges` - IPv6 CIDR blocks should return IPv6 addresses ([#51](https://github.com/cloudflare/terraform-provider-cloudflare/issues/51))
- `cloudflare_zone_settings_override` - Allow `0` for `browser_cache_ttl` ([#71](https://github.com/cloudflare/terraform-provider-cloudflare/issues/71))
- `cloudflare_page_rule` - `forwarding_urls` in page rules are lists ([#79](https://github.com/cloudflare/terraform-provider-cloudflare/issues/79))
- `cloudflare_page_rule` - The API supports `active` and `disabled`, not `paused` ([#84](https://github.com/cloudflare/terraform-provider-cloudflare/issues/84))

IMPROVEMENTS:

- `cloudflare_zone_settings_override` - Add support for `min_tls_version` ([#72](https://github.com/cloudflare/terraform-provider-cloudflare/issues/72))
- `cloudflare_page_rule` - Add support for more settings: `bypass_cache_on_cookie`, `cache_by_device_type`, `cache_deception_armor`, `cache_on_cookie`, `host_header_override`, `polish`, `explicit_cache_control`, `origin_error_page_pass_thru`, `sort_query_string_for_cache`, `resolve_override`, `respect_strong_etag`, `response_buffering`, `true_client_ip_header`, `mirage`, `disable_railgun`, `cache_key`, `waf`, `rocket_loader`, `cname_flattening` ([#68](https://github.com/cloudflare/terraform-provider-cloudflare/issues/68)], [[#81](https://github.com/cloudflare/terraform-provider-cloudflare/issues/81)], [[#85](https://github.com/cloudflare/terraform-provider-cloudflare/issues/85))
- `cloudflare_page_rule` - Add `off` setting to `security_level` ([#81](https://github.com/cloudflare/terraform-provider-cloudflare/issues/81))
- `cloudflare_record` - DNS Record improvements ([#97](https://github.com/cloudflare/terraform-provider-cloudflare/issues/97))
- Various documentation improvements

## 1.0.0 (April 06, 2018)

BACKWARDS INCOMPATIBILITIES / NOTES:

- resource/cloudflare_record: Changing `name` or `domain` now force a recreation
  of the record ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))

FEATURES:

- **New Resource:** `cloudflare_rate_limit` ([#30](https://github.com/cloudflare/terraform-provider-cloudflare/issues/30))
- **New Resource:** `cloudflare_page_rule` ([#38](https://github.com/cloudflare/terraform-provider-cloudflare/issues/38))
- **New Resource:** `cloudflare_load_balancer` ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))
- **New Resource:** `cloudflare_load_balancer_pool` ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))
- **New Resource:** `cloudflare_zone_settings_override` ([#41](https://github.com/cloudflare/terraform-provider-cloudflare/issues/41))
- **New Resource:** `cloudflare_load_balancer_monitor` ([#42](https://github.com/cloudflare/terraform-provider-cloudflare/issues/42))
- **New Data Source:** `cloudflare_ip_ranges` ([#28](https://github.com/cloudflare/terraform-provider-cloudflare/issues/28))

IMPROVEMENTS:

- resource/cloudflare_record: Validate `TXT` records ([#14](https://github.com/cloudflare/terraform-provider-cloudflare/issues/14))
- resource/cloudflare_record: Add `data` input to suppport SRV, LOC records
  ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))
- resource/cloudflare_record: Add computed attributes `created_on`,
  `modified_on`, `proxiable`, and `metadata` to records ([#29](https://github.com/cloudflare/terraform-provider-cloudflare/issues/29))
- resource/cloudflare_record: Support import of existing records ([#36](https://github.com/cloudflare/terraform-provider-cloudflare/issues/36))
- New Provider configuration options for API rate limiting ([#43](https://github.com/cloudflare/terraform-provider-cloudflare/issues/43))
- New Provider configuration options for using Organizations ([#40](https://github.com/cloudflare/terraform-provider-cloudflare/issues/40))

## 0.1.0 (June 20, 2017)

NOTES:

- Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider
  Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
