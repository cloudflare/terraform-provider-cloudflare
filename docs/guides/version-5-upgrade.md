---
layout: "cloudflare"
page_title: "Upgrading to version 5 (from 4.x)"
description: Terraform Cloudflare Provider Version 5 Upgrade Guide
---

# Terraform Cloudflare Provider Version 5 Upgrade Guide

Version 5 of the Cloudflare Terraform Provider is a ground-up rewrite of the
provider, using code generation from our OpenAPI spec. While this introduces
attribute and resource changes, it moves the provider to align more closely
with the service endpoints. This allows automation the steps to get changes
into the provider lowering the delay between new features and complete
coverage.

## Provider Version Configuration

If you are not ready to make a move to version 5 of the Cloudflare provider,
you may keep the 4.x branch active for your Terraform project by specifying:

```hcl
provider "cloudflare" {
  version = "~> 4"
  # ... any other configuration
}
```

We highly recommend reviewing this guide, make necessary changes and move to
5.x branch, as further 4.x releases are unlikely to happen outside of critical
security fixes.

~> Before attempting to upgrade to version 5, you should first upgrade to the
latest version of 4.x to ensure any transitional updates are applied to your
existing configuration.

Once ready, make the following change to use the latest 5.x release:

```hcl
provider "cloudflare" {
  version = "~> 5"
  # ... any other configuration
}
```

## Approach

At a high level, there are two parts to the migration. The first is the migration of the
configuration (HCL) and the second is the migration of the state. Within each of those
sections, there is the need to migrate attributes and potentially the resource rename.

### Automatic

For assisting with automatic migrations, we have provided [GritQL] patterns.

This will allow you to rewrite the parts of your Terraform configuration and state
that have changed automatically. Once you [install Grit], you can run the commands
in the directory where your Terraform configuration is located.

~> While all efforts have been made to ease the transition, some of the more complex
resources that may contain difficult to reconcile resources have been intentionally
skipped for the automatic migration and are only manually documented. If you are
using modules or other dynamic features of HCL, the provided codemods may not be
as effective. We recommend reviewing the manual migration notes to verify all the
changes.

We recommend ensuring you are using version control for these changes or make a
backup prior to initiating the change to enable reverting if needed.

1. Update the resource attributes in your configuration. _Note: this will not update
  your state file. The next step will determine how your state file is updated._
  ```bash
  $ grit apply github.com/cloudflare/terraform-provider-cloudflare#cloudflare_terraform_v5
  ```
2. Choose the appropriate method from [migrating renamed resources] that best suits
  your situation and use case to migrate the attribute changes. If you are choosing to
  use the provided GritQL patterns, the pattern name is
  `cloudflare_terraform_v5_attribute_renames_state`. Otherwise, you can reimport the
  resources without manually managing the state file.
3. Perform the resource renames. _Note: this will not update your state file.
  The next step will determine how your state file is updated._
  ```bash
  $ grit apply github.com/cloudflare/terraform-provider-cloudflare#cloudflare_terraform_v5_resource_renames_configuration
  ```
4. Choose the appropriate method from [migrating renamed resources] that best suits
  your situation and use case to migrate the resource renames. If you are choosing to
  use the provided GritQL patterns, the pattern name is
  `cloudflare_terraform_v5_resource_renames_state`.

### Manual

1. Update the resource attributes in your configuration using the migration notes.
2. Choose the appropriate method from [migrating renamed resources] that best suits
  your situation and use case to migrate the attribute changes.
3. Perform the resource renames using the migration notes.
4. Choose the appropriate method from [migrating renamed resources] that best suits
  your situation and use case to migrate the resource renames.

<!-- This code block is only used for confirming grit patterns -->

## Changelog

```grit
language hcl

cloudflare_terraform_v5()
```

## cloudflare_access_application

- Renamed to `cloudflare_zero_trust_access_application`

## cloudflare_access_ca_certificate

- Renamed to `cloudflare_zero_trust_access_short_lived_certificate`

## cloudflare_access_custom_page

- Renamed to `cloudflare_zero_trust_access_custom_page`

## cloudflare_access_group

- Renamed to `cloudflare_zero_trust_access_group`

## cloudflare_access_identity_provider

- Renamed to `cloudflare_zero_trust_access_identity_provider`

## cloudflare_access_keys_configuration

- Renamed to `cloudflare_zero_trust_access_key_configuration`

## cloudflare_access_mutual_tls_certificate

- Renamed to `cloudflare_zero_trust_access_mtls_certificate`

## cloudflare_access_mutual_tls_hostname_settings

- Renamed to `cloudflare_zero_trust_access_mtls_hostname_settings`

## cloudflare_access_organization

- Renamed to `cloudflare_zero_trust_organization`

## cloudflare_access_policy

- Renamed to `cloudflare_zero_trust_access_policy`

## cloudflare_access_service_token

- Renamed to `cloudflare_zero_trust_access_service_token`

## cloudflare_access_tag

- Renamed to `cloudflare_zero_trust_access_tag`

## cloudflare_device_dex_test

- Renamed to `cloudflare_zero_trust_dex_test`

## cloudflare_device_managed_networks

- Renamed to `cloudflare_zero_trust_device_managed_networks`

## cloudflare_device_policy_certificates

- Renamed to `cloudflare_zero_trust_device_certificates`

## cloudflare_device_posture_integration

- Renamed to `cloudflare_zero_trust_device_posture_integration`

## cloudflare_device_posture_rule

- Renamed to `cloudflare_zero_trust_device_posture_rule`

## cloudflare_device_settings_policy

- Renamed to `cloudflare_zero_trust_device_custom_profile` or `cloudflare_zero_trust_device_custom_profile` depending on your intended usage.

## cloudflare_dlp_custom_profile

- Renamed to `cloudflare_zero_trust_dlp_custom_profile`

## cloudflare_dlp_predefined_profile

- Renamed to `cloudflare_zero_trust_dlp_predefined_profile`

## cloudflare_dlp_profile

- Renamed to `cloudflare_zero_trust_custom_dlp_profile` or `cloudflare_zero_trust_predefined_dlp_profile` depending on which you are targeting.

## cloudflare_fallback_domain / cloudflare_zero_trust_local_fallback_domain

- Renamed to `cloudflare_zero_trust_device_custom_profile_local_domain_fallback` or `cloudflare_zero_trust_device_default_profile_local_domain_fallback` depending on which you are targeting.

## cloudflare_gateway_app_types

- Renamed to `cloudflare_zero_trust_gateway_app_types`

## cloudflare_gre_tunnel

- Renamed to `cloudflare_magic_wan_gre_tunnel`

## cloudflare_ipsec_tunnel

- Renamed to `cloudflare_magic_wan_ipsec_tunnel`

## cloudflare_record

- Renamed to `cloudflare_dns_record`

## cloudflare_risk_behavior

- Renamed to `cloudflare_zero_trust_risk_behavior`

## cloudflare_split_tunnel

- Renamed to `cloudflare_zero_trust_device_default_profile` and `cloudflare_zero_trust_device_custom_profile` depending on which you are targeting.

## cloudflare_static_route

- Renamed to `cloudflare_magic_wan_static_route`

## cloudflare_teams_account

- Renamed to `cloudflare_zero_trust_gateway_settings`

## cloudflare_teams_list

- Renamed to `cloudflare_zero_trust_list`

## cloudflare_teams_location

- Renamed to `cloudflare_zero_trust_dns_location`

## cloudflare_teams_proxy_endpoint

- Renamed to `cloudflare_zero_trust_gateway_proxy_endpoint`

## cloudflare_teams_rule

- Renamed to `cloudflare_zero_trust_gateway_policy`

## cloudflare_tunnel

- Renamed to `cloudflare_zero_trust_tunnel_cloudflared`

## cloudflare_tunnel_config

- Renamed to `cloudflare_zero_trust_tunnel_cloudflared_config`

## cloudflare_tunnel_route

- Renamed to `cloudflare_zero_trust_tunnel_cloudflared_route`

## cloudflare_tunnel_virtual_network

- Renamed to `cloudflare_zero_trust_tunnel_cloudflared_virtual_network`

## cloudflare_worker_cron_trigger

- Renamed to `cloudflare_workers_cron_trigger`

## cloudflare_worker_domain

- Renamed to `cloudflare_workers_custom_domain`

## cloudflare_worker_script

- Renamed to `cloudflare_workers_script`

## cloudflare_worker_secret

This has been removed. Users should instead use the:

- [Secrets Store](https://developers.cloudflare.com/secrets-store/) with the `secrets_store_secret` binding on the [cloudflare_workers_script resource](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/resources/workers_script)
- `secret_text` binding
- [Workers Secrets API](https://developers.cloudflare.com/api/resources/workers/subresources/scripts/subresources/secrets/)

## cloudflare_workers_for_platforms_namespace

- Renamed to `cloudflare_workers_for_platforms_dispatch_namespace`

## cloudflare_managed_headers

- Renamed to `cloudflare_managed_transforms`

## cloudflare_account_member

- `email_address` is now `email`.
- `role_ids` is now `roles`.

  Before

  ```hcl
  resource "cloudflare_account_member" "example" {
    email_address = "me@example.com"
    role_ids = ["a", "b", "c"]
  }
  ```

  After

  ```hcl
  resource "cloudflare_account_member" "example" {
    email = "me@example.com"
    roles = ["a", "b", "c"]
  }
  ```

## cloudflare_account

- `enforce_twofactor` is now `settings.enforce_twofactor`.

  Before

  ```hcl
  resource "cloudflare_account" "example" {
    enforce_twofactor = true
  }
  ```

  After

  ```hcl
  resource "cloudflare_account" "example" {
    settings = {
      enforce_twofactor = true
    }
  }
  ```

## cloudflare_byo_ip_prefix

- `advertisement` is now `advertised`.

  Before

  ```hcl
  resource "cloudflare_byo_ip_prefix" "example" {
    advertisement = true
  }
  ```

  After

  ```hcl
  resource "cloudflare_byo_ip_prefix" "example" {
    advertised = true
  }
  ```

## cloudflare_healthcheck

- `allow_insecure` is now nested under `http_config`
- `expected_body` is now nested under `http_config`.
- `expected_codes` is now nested under `http_config`.
- `follow_redirects` is now nested under `http_config`.
- `method` is now nested under `http_config`.
- `path` is now nested under `http_config`.
- `port` is now nested under `http_config`.

## cloudflare_zone_cache_reserve

- `enabled` is now `value`.

  Before

  ```hcl
  resource "cloudflare_zone_cache_reserve" "example" {
    enabled = true
  }

  resource "cloudflare_zone_cache_reserve" "example" {
    enabled = false
  }
  ```

  After

  ```hcl
  resource "cloudflare_zone_cache_reserve" "example" {
    value = "on"
  }

  resource "cloudflare_zone_cache_reserve" "example" {
    value = "off"
  }
  ```

## cloudflare_api_token

- `condition` is now a single nested attribute (`condition = { ... }`) instead of a block (`condition { ... }`).
- `request_ip` is now a single nested attribute (`request_ip = { ... }`) instead of a block (`request_ip { ... }`).
- `policy` is now `policies`.

Before
```hcl
resource "cloudflare_api_token" "example" {
  name = "example"
  policy = [{
    effect            = "allow"
    permission_groups = ["%[2]s"]
    resources         = { "com.cloudflare.api.account.zone.*" = "*" }
  }]
  condition {
    request_ip {
      in     = ["192.0.2.1/32"]
      not_in = ["198.51.100.1/32"]
    }
  }
}
```

After
```hcl
resource "cloudflare_api_token" "example" {
  name = "example"
  policies = [{
    effect            = "allow"
    permission_groups = ["%[2]s"]
    resources         = { "com.cloudflare.api.account.zone.*" = "*" }
  }]
  condition = {
    request_ip = {
      in     = ["192.0.2.1/32"]
      not_in = ["198.51.100.1/32"]
    }
  }
}
```

## cloudflare_hostname_tls_setting

- `setting` is now `setting_id`.

  Before

  ```hcl
  resource "cloudflare_hostname_tls_setting" "example" {
    setting = "min_tls_version"
  }
  ```

  After

  ```hcl
  resource "cloudflare_hostname_tls_setting" "example" {
    setting_id = "min_tls_version"
  }
  ```

## cloudflare_logpull_retention

- `enabled` is now `flag`.

  Before

  ```hcl
  resource "cloudflare_logpull_retention" "example" {
    enabled = true
  }
  ```

  After

  ```hcl
  resource "cloudflare_logpull_retention" "example" {
    flag = true
  }
  ```

## cloudflare_logpush_ownership_challenge

- `ownership_challenge_filename` is now `filename`.

  Before

  ```hcl
  resource "cloudflare_logpush_ownership_challenge" "example" {
    ownership_challenge_filename = "example"
  }
  ```

  After

  ```hcl
  resource "cloudflare_logpush_ownership_challenge" "example" {
    filename = "example"
  }
  ```

## cloudflare_magic_wan_gre_tunnel

- `health_check_enabled` is now `health_check.enabled`.
- `health_check_type` is now `health_check.type`.

  Before

  ```hcl
  resource "cloudflare_magic_wan_gre_tunnel" "example" {
    health_check_enabled = true
    health_check_type = "reply"
  }
  ```

  After

  ```hcl
  resource "cloudflare_magic_wan_gre_tunnel" "example" {
    health_check = {
      enabled = true
      type = "reply"
    }
  }
  ```

## cloudflare_magic_wan_ipsec_tunnel

- `health_check_direction` is now `health_check.direction`.
- `health_check_enabled` is now `health_check.enabled`.
- `health_check_rate` is now `health_check.rate`.
- `health_check_target` is now `health_check.target` with further nested attributes.
- `health_check_type` is now `health_check.type`.

  Before

  ```hcl
  resource "cloudflare_magic_wan_ipsec_tunnel" "example" {
    health_check_direction = "unidirectional"
    health_check_type = "reply"
    health_check_enabled = true
    health_check_rate = "low"
  }
  ```

  After

  ```hcl
  resource "cloudflare_magic_wan_ipsec_tunnel" "example" {
    health_check = {
      direction = "unidirectional"
      type = "reply"
      enabled = true
      rate = "low"
    }
  }
  ```

## cloudflare_zero_trust_tunnel_cloudflared

- `secret` is now `tunnel_secret`.

  Before

  ```hcl
  resource "cloudflare_zero_trust_tunnel_cloudflared" "example" {
    secret = "my_s3cr3t!"
  }
  ```

  After

  ```hcl
  resource "cloudflare_zero_trust_tunnel_cloudflared" "example" {
    tunnel_secret = "my_s3cr3t!"
  }
  ```

## cloudflare_zero_trust_access_short_lived_certificate

- `application_id` is now `app_id`.

  Before

  ```hcl
  resource "cloudflare_zero_trust_access_short_lived_certificate" "example" {
    application_id = "21d2f241bab5a0af65be098a0ba3d6c1"
  }
  ```

  After

  ```hcl
  resource "cloudflare_zero_trust_access_short_lived_certificate" "example" {
    app_id = "21d2f241bab5a0af65be098a0ba3d6c1"
  }
  ```

## cloudflare_workers_secret

This has been removed. Users should instead use the:

- [Secrets Store](https://developers.cloudflare.com/secrets-store/) with the `secrets_store_secret` binding on the [cloudflare_workers_script resource](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/resources/workers_script)
- `secret_text` binding
- [Workers Secrets API](https://developers.cloudflare.com/api/resources/workers/subresources/scripts/subresources/secrets/)

## cloudflare_workers_kv

- `key` is now `key_name`.

  Before

  ```hcl
  resource "cloudflare_workers_kv" "example" {
    key = "my_key"
  }
  ```

  After

  ```hcl
  resource "cloudflare_workers_kv" "example" {
    key_name = "my_key"
  }
  ```

## cloudflare_tiered_cache

- `cache_type` is now `value`.

  Before

  ```hcl
  resource "cloudflare_tiered_cache" "example" {
    cache_type = "on"
  }
  ```

  After

  ```hcl
  resource "cloudflare_tiered_cache" "example" {
    value = "on"
  }
  ```

## cloudflare_web_analytics_site

- `ruleset_id` is now `ruleset.id`.

  Before

  ```hcl
  resource "cloudflare_web_analytics_site" "example" {
    ruleset_id = "deadbeef"
  }
  ```

  After

  ```hcl
  resource "cloudflare_web_analytics_site" "example" {
    ruleset = {
      id = "deadbeef"
    }
  }
  ```

## cloudflare_zero_trust_gateway_settings

- `activity_log_enabled` is now `settings.activity_log.enabled`.
- `non_identity_browser_isolation_enabled` is now `settings.browser_isolation.non_identity_enabled`.
- `protocol_detection_enabled` is now `settings.protocol_detection.enabled`.
- `tls_decrypt_enabled` is now `settings.tls_decrypt.enabled`.
- `url_browser_isolation_enabled` is now `settings.browser_isolation.url_browser_isolation_enabled`.

## cloudflare_load_balancer

- `adaptive_routing` is now a single nested attribute (`adaptive_routing = { ... }`) instead of a block (`adaptive_routing { ... }`).
- `country_pools` is now a list of objects (`country_pools = [{ ... }]`) instead of multiple block attribute (`country_pools { ... }`).
- `fixed_response` is now a single nested attribute (`fixed_response = { ... }`) instead of a block (`fixed_response { ... }`).
- `location_strategy` is now a single nested attribute (`location_strategy = { ... }`) instead of a block (`location_strategy { ... }`).
- `overrides` is now a single nested attribute (`overrides = { ... }`) instead of a block (`overrides { ... }`).
- `pop_pools` is now a list of objects (`pop_pools = [{ ... }]`) instead of multiple block attribute (`pop_pools { ... }`).
- `random_steering` is now a single nested attribute (`random_steering = { ... }`) instead of a block (`random_steering { ... }`).
- `region_pools` is now a list of objects (`region_pools = [{ ... }]`) instead of multiple block attribute (`region_pools { ... }`).
- `rules` is now a list of objects (`rules = [{ ... }]`) instead of multiple block attribute (`rules { ... }`).
- `session_affinity_attributes` is now a single nested attribute (`session_affinity_attributes = { ... }`) instead of a block (`session_affinity_attributes { ... }`).
- `fallback_pool_id` is now `fallback_pool`.

  Before

  ```hcl
  resource "cloudflare_load_balancer" "example" {
    fallback_pool_id = "520636c63a13852db69ca56570b0abf4"
  }
  ```

  After

  ```hcl
  resource "cloudflare_load_balancer" "example" {
    fallback_pool = "520636c63a13852db69ca56570b0abf4"
  }
  ```

- `default_pool_ids` is now `default_pools`.

  Before

  ```hcl
  resource "cloudflare_load_balancer" "example" {
    default_pool_ids = ["520636c63a13852db69ca56570b0abf4", "4cc60288984088b5188246199f87daa7"]
  }
  ```

  After

  ```hcl
  resource "cloudflare_load_balancer" "example" {
    default_pools = ["520636c63a13852db69ca56570b0abf4", "4cc60288984088b5188246199f87daa7"]
  }
  ```

## cloudflare_r2_bucket

- `location_hint` is now `location`.

  Before

  ```hcl
  resource "cloudflare_r2_bucket" "example" {
   	location_hint = "APAC"
  }
  ```

  After

  ```hcl
  resource "cloudflare_r2_bucket" "example" {
   	location   = "APAC"
  }
  ```

## cloudflare_teams_list

- `items` is now a list of objects (`[{ value = $item }]`) instead of `items = [$item1, $item2]`

  Before

  ```hcl
  resource "cloudflare_teams_list" "example" {
   	items = ["item-1234", "item-5678"	]
  }
  ```

  After

  ```hcl
  resource "cloudflare_teams_list" "example" {
   	items = [{ value = "item-1234" }, { value = "item-5678" }]
  }
  ```

## cloudflare_workers_kv

- `key` is now `key_name`.

  Before

  ```hcl
  resource "cloudflare_workers_kv" "example" {
    key = "my-simple-key"
  }
  ```

  After

  ```hcl
  resource "cloudflare_workers_kv" "example" {
    key_name= "my-simple-key"
  }
  ```

## cloudflare_zero_trust_tunnel_cloudflared

- `secret` is now `tunnel_secret`.
- `cname` is no longer available.

  Before

  ```hcl
  resource "zero_trust_tunnel_cloudflared" "example" {
    account_id = "0da42c8d2132a9ddaf714f9e7c920711"
    secret = "example-secret"
    cname = "foo.example.com"
  }
  ```

  After

  ```hcl
  resource "zero_trust_tunnel_cloudflared" "example" {
    account_id = "0da42c8d2132a9ddaf714f9e7c920711"
    tunnel_secret = "example-secret"
  }
  ```

## cloudflare_zone_cache_variants

- file extensions are now nested under the `value` object.

  Before

  ```hcl
  resource "cloudflare_zone_cache_variants" "example" {
    zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
    avif    = ["image/avif", "image/webp"]
    bmp     = ["image/bmp", "image/webp"]
    gif     = ["image/gif", "image/webp"]
    jpeg    = ["image/jpeg", "image/webp"]
    jpg     = ["image/jpg", "image/webp"]
    jp2     = ["image/jp2", "image/webp"]
    jpg2    = ["image/jpg2", "image/webp"]
    png     = ["image/png"]
    tif     = ["image/tif", "image/webp"]
    tiff    = ["image/tiff", "image/webp"]
    webp    = ["image/webp"]
  }
  ```

  After

  ```hcl
  resource "cloudflare_zone_cache_variants" "example" {
    zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
    value = {
      avif = ["image/avif", "image/webp"]
      bmp  = ["image/bmp", "image/webp"]
      gif  = ["image/gif", "image/webp"]
      jpeg = ["image/jpeg", "image/webp"]
      jpg  = ["image/jpg", "image/webp"]
      jp2  = ["image/jp2", "image/webp"]
      jpg2 = ["image/jpg2", "image/webp"]
      png  = ["image/png"]
      tif  = ["image/tif", "image/webp"]
      tiff = ["image/tiff", "image/webp"]
      webp = ["image/webp"]
    }
  }
  ```

## cloudflare_custom_page

- `cloudflare_custom_page` has been removed.

## cloudflare_zone_settings_override

- `cloudflare_zone_settings_override` has been removed. Use `cloudflare_zone_setting` instead on a per setting basis.

## cloudflare_zone

- Zone subscriptions are now controlled independently using `cloudflare_zone_subscription` resource.
- `zone` is now an `name`.

  Before

  ```hcl
  resource "cloudflare_zone" "example" {
    zone = "example.com"
  }
  ```

  After

  ```hcl
  resource "cloudflare_zone" "example" {
    name = "example.com"
  }
  ```

- `account_id` is now an `account` object with the `id` attribute inside.

  Before

  ```hcl
  resource "cloudflare_zone" "example" {
    account_id = "f037e56e89293a057740de681ac9abbe"
    zone       = "example.com"
  }
  ```

  After

  ```hcl
  resource "cloudflare_zone" "example" {
    account = {
      id = "f037e56e89293a057740de681ac9abbe"
    }
    name   = "example.com"
  }
  ```

  - `jump_start` is no longer an attribute.

    Before

    ```hcl
    resource "cloudflare_zone" "example" {
      account = {
        id = "f037e56e89293a057740de681ac9abbe"
      }
      zone       = "example.com"
      jump_start = false
    }
    ```

    After

    ```hcl
    resource "cloudflare_zone" "example" {
      account = {
        id = "f037e56e89293a057740de681ac9abbe"
      }
      name   = "example.com"
    }
    ```

## cloudflare_zero_trust_access_group

- `auth_context` is now a list of objects (`auth_context = [{ ... }]`) instead of multiple block attribute (`auth_context { ... }`).
- `azure` is now a single nested attribute (`azure = { ... }`) instead of a block (`azure { ... }`).
- `exclude` is now a list of objects (`exclude = [{ ... }]`) instead of multiple block attribute (`exclude { ... }`).
- `external_evaluation` is now a single nested attribute (`external_evaluation = { ... }`) instead of a block (`external_evaluation { ... }`).
- `github` is now a list of objects (`github = [{ ... }]`) instead of multiple block attribute (`github { ... }`).
- `gsuite` is now a single nested attribute (`gsuite = { ... }`) instead of a block (`gsuite { ... }`).
- `include` is now a list of objects (`include = [{ ... }]`) instead of multiple block attribute (`include { ... }`).
- `okta` is now a single nested attribute (`okta = { ... }`) instead of a block (`okta { ... }`).
- `require` is now a list of objects (`require = [{ ... }]`) instead of multiple block attribute (`require { ... }`).
- `saml` is now a single nested attribute (`saml = { ... }`) instead of a block (`saml { ... }`).
- `application_id` and `precedence` no longer used.

  Before

  ```hcl
  resource "cloudflare_zero_trust_access_group" "example" {
    account_id      = "f037e56e89293a057740de681ac9abbe"
    application_id  = "foo"
    name            = "example"
    precedence      = 3
  }
  ```

  After

  ```hcl
  resource "cloudflare_zero_trust_access_group" "example" {
    account_id = "f037e56e89293a057740de681ac9abbe"
    name       = "example"
  }
  ```

## cloudflare_zero_trust_access_service_token

- `min_days_for_renewal` is no longer used. If you would like similar functionality, you can use `duration = "forever"` instead.

  Before

  ```hcl
  resource "cloudflare_zero_trust_access_service_token" "example" {
    account_id           = "f037e56e89293a057740de681ac9abbe"
    name                 = "CI/CD app renewed"
    min_days_for_renewal = 30
  }
  ```

  After

  ```hcl
  resource "cloudflare_zero_trust_access_service_token" "example" {
    account_id = "f037e56e89293a057740de681ac9abbe"
    name       = "CI/CD app renewed"
  }
  ```

## cloudflare_zero_trust_access_application

- `authentication` is now a single nested attribute (`authentication = { ... }`) instead of a block (`authentication { ... }`).
- `cors_headers` is now a single nested attribute (`cors_headers = { ... }`) instead of a block (`cors_headers { ... }`).
- `custom_attribute` is now a list of objects (`custom_attribute = [{ ... }]`) instead of multiple block attribute (`custom_attribute { ... }`).
- `custom_claim` is now a list of objects (`custom_claim = [{ ... }]`) instead of multiple block attribute (`custom_claim { ... }`).
- `footer_links` is now a list of objects (`footer_links = [{ ... }]`) instead of multiple block attribute (`footer_links { ... }`).
- `hybrid_and_implicit_options` is now a single nested attribute (`hybrid_and_implicit_options = { ... }`) instead of a block (`hybrid_and_implicit_options { ... }`).
- `landing_page_design` is now a list of objects (`landing_page_design = [{ ... }]`) instead of multiple block attribute (`landing_page_design { ... }`).
- `mappings` is now a list of objects (`mappings = [{ ... }]`) instead of multiple block attribute (`mappings { ... }`).
- `operations` is now a single nested attribute (`operations = { ... }`) instead of a block (`operations { ... }`).
- `refresh_token_options` is now a single nested attribute (`refresh_token_options = { ... }`) instead of a block (`refresh_token_options { ... }`).
- `saas_app` is now a single nested attribute (`saas_app = { ... }`) instead of a block (`saas_app { ... }`).
- `scim_config` is now a single nested attribute (`scim_config = { ... }`) instead of a block (`scim_config { ... }`).
- `source` is now a list of objects (`source = [{ ... }]`) instead of multiple block attribute (`source { ... }`).

## cloudflare_zero_trust_access_mtls_hostname_settings

- `settings` is now a list of objects (`settings = [{ ... }]`) instead of multiple block attribute (`settings { ... }`).

## cloudflare_zero_trust_access_policy

- `approval_group` is now a list of objects (`approval_group = [{ ... }]`) instead of multiple block attribute (`approval_group { ... }`).
- `auth_context` is now a list of objects (`auth_context = [{ ... }]`) instead of multiple block attribute (`auth_context { ... }`).
- `azure` is now a single nested attribute (`azure = { ... }`) instead of a block (`azure { ... }`).
- `exclude` is now a list of objects (`exclude = [{ ... }]`) instead of multiple block attribute (`exclude { ... }`).
- `external_evaluation` is now a single nested attribute (`external_evaluation = { ... }`) instead of a block (`external_evaluation { ... }`).
- `github` is now a list of objects (`github = [{ ... }]`) instead of multiple block attribute (`github { ... }`).
- `gsuite` is now a single nested attribute (`gsuite = { ... }`) instead of a block (`gsuite { ... }`).
- `include` is now a list of objects (`include = [{ ... }]`) instead of multiple block attribute (`include { ... }`).
- `okta` is now a single nested attribute (`okta = { ... }`) instead of a block (`okta { ... }`).
- `require` is now a list of objects (`require = [{ ... }]`) instead of multiple block attribute (`require { ... }`).
- `saml` is now a single nested attribute (`saml = { ... }`) instead of a block (`saml { ... }`).

## cloudflare_zero_trust_access_identity_provider

- `config` is now a single nested attribute (`config = { ... }`) instead of a block (`config { ... }`).
- `scim_config` is now a single nested attribute (`scim_config = { ... }`) instead of a block (`scim_config { ... }`).

## cloudflare_zero_trust_organization

- `custom_pages` is now a single nested attribute (`custom_pages = { ... }`) instead of a block (`custom_pages { ... }`).
- `login_design` is now a single nested attribute (`login_design = { ... }`) instead of a block (`login_design { ... }`).

## cloudflare_address_map

- `ips` is now a list of objects (`ips = [{ ... }]`) instead of multiple block attribute (`ips { ... }`).
- `memberships` is now a list of objects (`memberships = [{ ... }]`) instead of multiple block attribute (`memberships { ... }`).

## cloudflare_api_shield

- `auth_id_characteristics` is now a list of objects (`auth_id_characteristics = [{ ... }]`) instead of multiple block attribute (`auth_id_characteristics { ... }`).

## cloudflare_certificate_pack

- `validation_errors` is now a list of objects (`validation_errors = [{ ... }]`) instead of multiple block attribute (`validation_errors { ... }`).
- `validation_records` is now a list of objects (`validation_records = [{ ... }]`) instead of multiple block attribute (`validation_records { ... }`).

## cloudflare_custom_hostname

- `settings` is now a single nested attribute (`settings = { ... }`) instead of a block (`settings { ... }`).
- `ssl` is now a single nested attribute (`ssl = { ... }`) instead of a block (`ssl { ... }`).

## cloudflare_custom_ssl

- `custom_ssl_options` is now a list of objects (`custom_ssl_options = [{ ... }]`) instead of multiple block attribute (`custom_ssl_options { ... }`).
- `custom_ssl_priority` is now a list of objects (`custom_ssl_priority = [{ ... }]`) instead of multiple block attribute (`custom_ssl_priority { ... }`).

## cloudflare_zero_trust_dex_test

- `data` is now a single nested attribute (`data = { ... }`) instead of a block (`data { ... }`).

## cloudflare_zero_trust_device_managed_networks

- `config` is now a single nested attribute (`config = { ... }`) instead of a block (`config { ... }`).

## cloudflare_zero_trust_device_posture_integration

- `config` is now a single nested attribute (`config = { ... }`) instead of a block (`config { ... }`).

## cloudflare_zero_trust_device_posture_rule

- `input` is now a single nested attribute (`input = { ... }`) instead of a block (`input { ... }`).
- `match` is now a list of objects (`match = [{ ... }]`) instead of multiple block attribute (`match { ... }`).

## cloudflare_zero_trust_custom_dlp_profile / cloudflare_zero_trust_predefined_dlp_profile

- `context_awareness` is now a list of objects (`context_awareness = [{ ... }]`) instead of multiple block attribute (`context_awareness { ... }`).
- `entry` is now a list of objects (`entry = [{ ... }]`) instead of multiple block attribute (`entry { ... }`).
- `pattern` is now a list of objects (`pattern = [{ ... }]`) instead of multiple block attribute (`pattern { ... }`).
- `skip` is now a list of objects (`skip = [{ ... }]`) instead of multiple block attribute (`skip { ... }`).

## cloudflare_email_routing_catch_all

- `action` is now a list of objects (`action = [{ ... }]`) instead of multiple block attribute (`action { ... }`).
- `matcher` is now a list of objects (`matcher = [{ ... }]`) instead of multiple block attribute (`matcher { ... }`).

## cloudflare_email_routing_rule

- `action` is now a list of objects (`action = [{ ... }]`) instead of multiple block attribute (`action { ... }`).
- `matcher` is now a list of objects (`matcher = [{ ... }]`) instead of multiple block attribute (`matcher { ... }`).

## cloudflare_healthcheck

- `header` is now a list of objects (`header = [{ ... }]`) instead of multiple block attribute (`header { ... }`).

## cloudflare_list

- `hostname` is now a list of objects (`hostname = [{ ... }]`) instead of multiple block attribute (`hostname { ... }`).
- `item` is now a list of objects (`item = [{ ... }]`) instead of multiple block attribute (`item { ... }`).
- `redirect` is now a list of objects (`redirect = [{ ... }]`) instead of multiple block attribute (`redirect { ... }`).
- `value` is now a list of objects (`value = [{ ... }]`) instead of multiple block attribute (`value { ... }`).

## cloudflare_list_item

- `hostname` is now a single nested attribute (`hostname = { ... }`) instead of a block (`hostname { ... }`).
- `redirect` is now a single nested attribute (`redirect = { ... }`) instead of a block (`redirect { ... }`).

## cloudflare_load_balancer_monitor

- `header` is now a list of objects (`header = [{ ... }]`) instead of multiple block attribute (`header { ... }`).

## cloudflare_load_balancer_pool

- `header` is now a single nested attribute (`header = { ... }`) instead of a block (`header { ... }`).
- `load_shedding` is now a single nested attribute (`load_shedding = { ... }`) instead of a block (`load_shedding { ... }`).
- `origin_steering` is now a single nested attribute (`origin_steering = { ... }`) instead of a block (`origin_steering { ... }`).
- `origins` is now a list of objects (`origins = [{ ... }]`) instead of multiple block attribute (`origins { ... }`).

## cloudflare_logpush_job

- `output_options` is now a single nested attribute (`output_options = { ... }`) instead of a block (`output_options { ... }`).

## cloudflare_managed_transforms

- `managed_request_headers` is now a list of objects (`managed_request_headers = [{ ... }]`) instead of multiple block attribute (`managed_request_headers { ... }`).
- `managed_response_headers` is now a list of objects (`managed_response_headers = [{ ... }]`) instead of multiple block attribute (`managed_response_headers { ... }`).

## cloudflare_notification_policy

- `email_integration` is now a list of objects (`email_integration = [{ ... }]`) instead of multiple block attribute (`email_integration { ... }`).
- `filters` is now a single nested attribute (`filters = { ... }`) instead of a block (`filters { ... }`).
- `pagerduty_integration` is now a list of objects (`pagerduty_integration = [{ ... }]`) instead of multiple block attribute (`pagerduty_integration { ... }`).
- `webhooks_integration` is now a list of objects (`webhooks_integration = [{ ... }]`) instead of multiple block attribute (`webhooks_integration { ... }`).

## cloudflare_pages_project

- `build_config` is now a single nested attribute (`build_config = { ... }`) instead of a block (`build_config { ... }`).
- `config` is now a list of objects (`config = [{ ... }]`) instead of multiple block attribute (`config { ... }`).
- `deployment_configs` is now a single nested attribute (`deployment_configs = { ... }`) instead of a block (`deployment_configs { ... }`).
- `placement` is now a single nested attribute (`placement = { ... }`) instead of a block (`placement { ... }`).
- `preview` is now a single nested attribute (`preview = { ... }`) instead of a block (`preview { ... }`).
- `production` is now a single nested attribute (`production = { ... }`) instead of a block (`production { ... }`).
- `service_binding` is now a list of objects (`service_binding = [{ ... }]`) instead of multiple block attribute (`service_binding { ... }`).
- `source` is now a list of objects (`source = [{ ... }]`) instead of multiple block attribute (`source { ... }`).

## cloudflare_dns_record

- `data` is now a single nested attribute (`data = { ... }`) instead of a block (`data { ... }`).
- `data.flag` is now a number (`flag = 0`) instead of a string (`flag = "0"`).
- `hostname` has been removed. Instead, you should use a combination of data source and resource attributes to get the same value.
- `allow_overwrite` has been removed.

## cloudflare_zero_trust_risk_behavior

- `behavior` is now a list of objects (`behavior = [{ ... }]`) instead of multiple block attribute (`behavior { ... }`).

## cloudflare_ruleset

- `action_parameters` is now a single nested attribute (`action_parameters = { ... }`) instead of a block (`action_parameters { ... }`).
- `algorithms` is now a list of objects (`algorithms = [{ ... }]`) instead of multiple block attribute (`algorithms { ... }`).
- `autominify` is now a list of objects (`autominify = [{ ... }]`) instead of multiple block attribute (`autominify { ... }`).
- `browser_ttl` is now a list of objects (`browser_ttl = [{ ... }]`) instead of multiple block attribute (`browser_ttl { ... }`).
- `cache_key` is now a list of objects (`cache_key = [{ ... }]`) instead of multiple block attribute (`cache_key { ... }`).
- `categories` is now a list of objects (`categories = [{ ... }]`) instead of multiple block attribute (`categories { ... }`).
- `cookie` is now a list of objects (`cookie = [{ ... }]`) instead of multiple block attribute (`cookie { ... }`).
- `custom_key` is now a list of objects (`custom_key = [{ ... }]`) instead of multiple block attribute (`custom_key { ... }`).
- `edge_ttl` is now a list of objects (`edge_ttl = [{ ... }]`) instead of multiple block attribute (`edge_ttl { ... }`).
- `exposed_credential_check` is now a list of objects (`exposed_credential_check = [{ ... }]`) instead of multiple block attribute (`exposed_credential_check { ... }`).
- `from_list` is now a list of objects (`from_list = [{ ... }]`) instead of multiple block attribute (`from_list { ... }`).
- `from_value` is now a list of objects (`from_value = [{ ... }]`) instead of multiple block attribute (`from_value { ... }`).
- `header` is now a list of objects (`header = [{ ... }]`) instead of multiple block attribute (`header { ... }`).
- `headers` is now a map of attributes keyed by the name instead of multiple block attribute (`headers { ... }`).
- `host` is now a list of objects (`host = [{ ... }]`) instead of multiple block attribute (`host { ... }`).
- `logging` is now a single nested attribute (`logging = { ... }`) instead of a block (`logging { ... }`).
- `matched_data` is now a list of objects (`matched_data = [{ ... }]`) instead of multiple block attribute (`matched_data { ... }`).
- `origin` is now a list of objects (`origin = [{ ... }]`) instead of multiple block attribute (`origin { ... }`).
- `overrides` is now a list of objects (`overrides = [{ ... }]`) instead of multiple block attribute (`overrides { ... }`).
- `path` is now a list of objects (`path = [{ ... }]`) instead of multiple block attribute (`path { ... }`).
- `query_string` is now a list of objects (`query_string = [{ ... }]`) instead of multiple block attribute (`query_string { ... }`).
- `query` is now a list of objects (`query = [{ ... }]`) instead of multiple block attribute (`query { ... }`).
- `ratelimit` is now a list of objects (`ratelimit = [{ ... }]`) instead of multiple block attribute (`ratelimit { ... }`).
- `response` is now a single nested attribute (`response = { ... }`) instead of a block (`response { ... }`).
- `rules` is now a list of objects (`rules = [{ ... }]`) instead of multiple block attribute (`rules { ... }`).
- `serve_stale` is now a list of objects (`serve_stale = [{ ... }]`) instead of multiple block attribute (`serve_stale { ... }`).
- `sni` is now a list of objects (`sni = [{ ... }]`) instead of multiple block attribute (`sni { ... }`).
- `status_code_range` is now a list of objects (`status_code_range = [{ ... }]`) instead of multiple block attribute (`status_code_range { ... }`).
- `status_code_ttl` is now a list of objects (`status_code_ttl = [{ ... }]`) instead of multiple block attribute (`status_code_ttl { ... }`).
- `target_url` is now a list of objects (`target_url = [{ ... }]`) instead of multiple block attribute (`target_url { ... }`).
- `uri` is now a list of objects (`uri = [{ ... }]`) instead of multiple block attribute (`uri { ... }`).
- `user` is now a list of objects (`user = [{ ... }]`) instead of multiple block attribute (`user { ... }`).

## cloudflare_spectrum_application

- `dns` is now a single nested attribute (`dns = { ... }`) instead of a block (`dns { ... }`).
- `edge_ips` is now a single nested attribute (`edge_ips = { ... }`) instead of a block (`edge_ips { ... }`).
- `origin_dns` is now a single nested attribute (`origin_dns = { ... }`) instead of a block (`origin_dns { ... }`).
- `origin_port_range` is now a list of objects (`origin_port_range = [{ ... }]`) instead of multiple block attribute (`origin_port_range { ... }`).

## cloudflare_zero_trust_split_tunnels

- `tunnels` is now a list of objects (`tunnels = [{ ... }]`) instead of multiple block attribute (`tunnels { ... }`).

## cloudflare_zero_trust_gateway_settings

- `antivirus` is now a list of objects (`antivirus = [{ ... }]`) instead of multiple block attribute (`antivirus { ... }`).
- `block_page` is now a list of objects (`block_page = [{ ... }]`) instead of multiple block attribute (`block_page { ... }`).
- `body_scanning` is now a list of objects (`body_scanning = [{ ... }]`) instead of multiple block attribute (`body_scanning { ... }`).
- `custom_certificate` is now a list of objects (`custom_certificate = [{ ... }]`) instead of multiple block attribute (`custom_certificate { ... }`).
- `dns` is now a list of objects (`dns = [{ ... }]`) instead of multiple block attribute (`dns { ... }`).
- `extended_email_matching` is now a list of objects (`extended_email_matching = [{ ... }]`) instead of multiple block attribute (`extended_email_matching { ... }`).
- `fips` is now a list of objects (`fips = [{ ... }]`) instead of multiple block attribute (`fips { ... }`).
- `http` is now a list of objects (`http = [{ ... }]`) instead of multiple block attribute (`http { ... }`).
- `l4` is now a list of objects (`l4 = [{ ... }]`) instead of multiple block attribute (`l4 { ... }`).
- `logging` is now a list of objects (`logging = [{ ... }]`) instead of multiple block attribute (`logging { ... }`).
- `notification_settings` is now a list of objects (`notification_settings = [{ ... }]`) instead of multiple block attribute (`notification_settings { ... }`).
- `payload_log` is now a list of objects (`payload_log = [{ ... }]`) instead of multiple block attribute (`payload_log { ... }`).
- `proxy` is now a list of objects (`proxy = [{ ... }]`) instead of multiple block attribute (`proxy { ... }`).
- `settings_by_rule_type` is now a list of objects (`settings_by_rule_type = [{ ... }]`) instead of multiple block attribute (`settings_by_rule_type { ... }`).
- `ssh_session_log` is now a list of objects (`ssh_session_log = [{ ... }]`) instead of multiple block attribute (`ssh_session_log { ... }`).

## cloudflare_zero_trust_dns_location

- `networks` is now a list of objects (`networks = [{ ... }]`) instead of multiple block attribute (`networks { ... }`).

## cloudflare_zero_trust_gateway_policy

- `audit_ssh` is now a single nested attribute (`audit_ssh = { ... }`) instead of a block (`audit_ssh { ... }`).
- `biso_admin_controls` is now a single nested attribute (`biso_admin_controls = { ... }`) instead of a block (`biso_admin_controls { ... }`).
- `check_session` is now a single nested attribute (`check_session = { ... }`) instead of a block (`check_session { ... }`).
- `dns_resolvers` is now a single nested attribute (`dns_resolvers = { ... }`) instead of a block (`dns_resolvers { ... }`).
- `egress` is now a single nested attribute (`egress = { ... }`) instead of a block (`egress { ... }`).
- `ipv4` is now a list of objects (`ipv4 = [{ ... }]`) instead of multiple block attribute (`ipv4 { ... }`).
- `ipv6` is now a list of objects (`ipv6 = [{ ... }]`) instead of multiple block attribute (`ipv6 { ... }`).
- `l4override` is now a single nested attribute (`l4override = { ... }`) instead of a block (`l4override { ... }`).
- `notification_settings` is now a single nested attribute (`notification_settings = { ... }`) instead of a block (`notification_settings { ... }`).
- `payload_log` is now a single nested attribute (`payload_log = { ... }`) instead of a block (`payload_log { ... }`).
- `rule_settings` is now a single nested attribute (`rule_settings = { ... }`) instead of a block (`rule_settings { ... }`).
- `untrusted_cert` is now a single nested attribute (`untrusted_cert = { ... }`) instead of a block (`untrusted_cert { ... }`).

## cloudflare_zero_trust_tunnel_cloudflared_config

- `access` is now a list of objects (`access = [{ ... }]`) instead of multiple block attribute (`access { ... }`).
- `access` is now a single nested attribute (`access = { ... }`) instead of a block (`access { ... }`).
- `config` is now a single nested attribute (`config = { ... }`) instead of a block (`config { ... }`).
- `ingress_rule` is now a list of objects (`ingress_rule = [{ ... }]`) instead of multiple block attribute (`ingress_rule { ... }`).
- `ip_rules` is now a list of objects (`ip_rules = [{ ... }]`) instead of multiple block attribute (`ip_rules { ... }`).
- `origin_request` is now a list of objects (`origin_request = [{ ... }]`) instead of multiple block attribute (`origin_request { ... }`).
- `origin_request` is now a single nested attribute (`origin_request = { ... }`) instead of a block (`origin_request { ... }`).
- `warp_routing` is now a single nested attribute (`warp_routing = { ... }`) instead of a block (`warp_routing { ... }`).

## cloudflare_waiting_room

- `additional_routes` is now a list of objects (`additional_routes = [{ ... }]`) instead of multiple block attribute (`additional_routes { ... }`).

## cloudflare_waiting_room_rules

- `rules` is now a list of objects (`rules = [{ ... }]`) instead of multiple block attribute (`rules { ... }`).

## cloudflare_workers_script

- `name` is now `script_name`.
- `analytics_engine_binding` is now a list of objects (`analytics_engine_binding = [{ ... }]`) instead of multiple block attribute (`analytics_engine_binding { ... }`).
- `d1_database_binding` is now a list of objects (`d1_database_binding = [{ ... }]`) instead of multiple block attribute (`d1_database_binding { ... }`).
- `kv_namespace_binding` is now a list of objects (`kv_namespace_binding = [{ ... }]`) instead of multiple block attribute (`kv_namespace_binding { ... }`).
- `placement` is now a list of objects (`placement = [{ ... }]`) instead of multiple block attribute (`placement { ... }`).
- `plain_text_binding` is now a list of objects (`plain_text_binding = [{ ... }]`) instead of multiple block attribute (`plain_text_binding { ... }`).
- `queue_binding` is now a list of objects (`queue_binding = [{ ... }]`) instead of multiple block attribute (`queue_binding { ... }`).
- `r2_bucket_binding` is now a list of objects (`r2_bucket_binding = [{ ... }]`) instead of multiple block attribute (`r2_bucket_binding { ... }`).
- `secret_text_binding` is now a list of objects (`secret_text_binding = [{ ... }]`) instead of multiple block attribute (`secret_text_binding { ... }`).
- `service_binding` is now a list of objects (`service_binding = [{ ... }]`) instead of multiple block attribute (`service_binding { ... }`).
- `webassembly_binding` is now a list of objects (`webassembly_binding = [{ ... }]`) instead of multiple block attribute (`webassembly_binding { ... }`).

## cloudflare_magic_wan_static_route

- `colo_names` is now `scope.colo_names`
- `colo_regions` is now `scope.colo_regions`

## cloudflare_page_rule

- `actions`is now a single nested attribute instead of a block.
- `ignore = true` is now `exclude = ["*"]`
- `ignore = false` is now `include = ["*"]`
- `cache_ttl_by_status` is now a map (`cache_ttl_by_status = { ... }`) instead of a list of objects (`cache_ttl_by_status = [{ ... }]`)

Before

```
resource "cloudflare_page_rule" "example" {
  target = "example.com"
  actions {
    cache_key_fields = {
      query_string = {
        ignore = true
        ignore = false
      }
    }
  }
}
```

After

```
resource "cloudflare_page_rule" "example" {
  target = "example.com"
  actions = {
    cache_key_fields = {
      query_string = {
        exclude = ["*"]
        include = ["*"]
      }
    }
  }
}
```

[GritQL]: https://www.grit.io/
[install Grit]: https://docs.grit.io/cli/quickstart
[migrating renamed resources]: https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/guides/migrating-renamed-resources
