---
layout: "cloudflare"
page_title: "Upgrading to version 5 (from 4.x)"
description: Terraform Cloudflare Provider Version 5 Upgrade Guide
---

# Terraform Cloudflare Provider Version 5 Upgrade Guide

Version 5 of the Cloudflare Terraform Provider is a ground-up rewrite of the
provider, using code generation from our OpenAPI spec. While this introduces
attribute and resource changes, it moves the provider to align more closely
with the service endpoints and makes automatic support going forward possible.

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

## Resource renames

For an in depth guide on how to perform migrations for resources or datasources that
have been renamed, check out [migrating renamed resources].

## Automatic migration

For assisting with automatic migrations, we have provided a [GritQL] pattern.

This will allow you to rewrite the parts of your Terraform configuration (not state)
that have changed automatically. Once you [install Grit], you can run the following
command in the directory where your Terraform configuration is located.

```bash
$ grit apply cloudflare_terraform_v5
```

We recommend ensuring you are using version control for these changes or make a
backup prior to initiating the change to enable reverting if needed.

~> If you are using modules or other dynamic features of HCL, the provided
   codemods may not be as effective. We recommend reviewing the migration notes below
   to verify all the changes.

<!-- This code block is only used for confirming grit patterns -->
```grit
language hcl

terraform_cloudflare_v5()
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

- Renamed to `cloudflare_zero_trust_device_profiles`

## cloudflare_dlp_custom_profile

- Renamed to `cloudflare_zero_trust_dlp_custom_profile`

## cloudflare_dlp_predefined_profile

- Renamed to `cloudflare_zero_trust_dlp_predefined_profile`

## cloudflare_dlp_profile

- Renamed to `cloudflare_zero_trust_dlp_profile`

## cloudflare_fallback_domain

- Renamed to `cloudflare_zero_trust_local_domain_fallback`

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

- Renamed to `cloudflare_zero_trust_split_tunnels`

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

- Renamed to `cloudflare_zero_trust_tunnel_route`

## cloudflare_tunnel_virtual_network

- Renamed to `cloudflare_zero_trust_tunnel_virtual_network`

## cloudflare_worker_cron_trigger

- Renamed to `cloudflare_workers_cron_trigger`

## cloudflare_worker_domain

- Renamed to `cloudflare_workers_custom_domain`

## cloudflare_worker_script

- Renamed to `cloudflare_workers_script`

## cloudflare_worker_secret

- Renamed to `cloudflare_workers_secret`

## cloudflare_workers_for_platforms_namespace

- Renamed to `cloudflare_workers_for_platforms_dispatch_namespace`

## cloudflare_zone_dnssec

- Renamed to `cloudflare_dns_zone_dnssec`

## cloudflare_managed_headers

- Renamed to `cloudflare_managed_transforms`

## cloudflare_api_token

- `policy` is now `policies`.

  Before
  ```hcl
  resource "cloudflare_api_token" "example" {
    name = "example"
		policy = [{
			effect = "allow"
			permission_groups = [ "%[2]s" ]
			resources = { "com.cloudflare.api.account.zone.*" = "*" }
		}]
		condition = {
      request_ip = {
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
			effect = "allow"
			permission_groups = [ "%[2]s" ]
			resources = { "com.cloudflare.api.account.zone.*" = "*" }
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
    zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
    setting = "min_tls_version"
    value = "1.2"
  }
  ```

  After
  ```hcl
  resource "cloudflare_hostname_tls_setting" "example" {
    zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
    setting_id = "min_tls_version"
    value      = "1.2"
  }
  ```

## cloudflare_load_balancer

- `fallback_pool_id` is now `fallback_pool`.

  Before
  ```hcl
  resource "cloudflare_load_balancer" "example" {
    zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
    fallback_pool_id = "520636c63a13852db69ca56570b0abf4"
  }
  ```

  After
  ```hcl
  resource "cloudflare_load_balancer" "example" {
    zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
    fallback_pool = "520636c63a13852db69ca56570b0abf4"
  }
  ```

- `default_pool_ids` is now `default_pools`.

  Before
  ```hcl
  resource "cloudflare_load_balancer" "example" {
    zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
    default_pool_ids = ["520636c63a13852db69ca56570b0abf4", "4cc60288984088b5188246199f87daa7"]
  }
  ```

  After
  ```hcl
  resource "cloudflare_load_balancer" "example" {
    zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
    default_pools = ["520636c63a13852db69ca56570b0abf4", "4cc60288984088b5188246199f87daa7"]
  }
  ```

## cloudflare_r2_bucket

- `location_hint` is now `location`.

  Before
  ```hcl
  resource "cloudflare_r2_bucket" "example" {
   	account_id    = "f037e56e89293a057740de681ac9abbe"
   	name          = "example"
   	location_hint = "APAC"
  }
  ```

  After
  ```hcl
  resource "cloudflare_r2_bucket" "example" {
   	account_id = "f037e56e89293a057740de681ac9abbe"
   	name       = "example"
   	location   = "APAC"
  }
  ```

## cloudflare_teams_list

- `items` is now an object of `value = $item` instead of `items = [$item1, $item2]`

  Before
  ```hcl
  resource "cloudflare_teams_list" "example" {
   	account_id  = "f037e56e89293a057740de681ac9abbe"
   	name        = "example"
   	description = "My description"
   	type        = "SERIAL"
   	items       = ["item-1234", "item-5678"	]
  }
  ```

  After
  ```hcl
  resource "cloudflare_teams_list" "example" {
   	account_id  = "f037e56e89293a057740de681ac9abbe"
   	name        = "example"
   	description = "My description"
   	type        = "SERIAL"
   	items       = [{ value = "item-1234" }, { value = "item-5678" }]
  }
  ```

## cloudflare_workers_kv

- `key` is now `key_name`.

  Before
  ```hcl
  resource "cloudflare_workers_kv" "example" {
    account_id = "0da42c8d2132a9ddaf714f9e7c920711"
    namespace_id = "9e5bd5c4acd7201064fe42d4e46cc48c"
    key = "my-simple-key"
    value = "foo"
  }
  ```

  After
  ```hcl
  resource "cloudflare_workers_kv" "example" {
    account_id = "0da42c8d2132a9ddaf714f9e7c920711"
    namespace_id = "9e5bd5c4acd7201064fe42d4e46cc48c"
    key_name= "my-simple-key"
    value = "foo"
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

[GritQL]: https://www.grit.io/
[install Grit]: https://docs.grit.io/cli/quickstart
[migrating renamed resources]: https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/guides/migrating-renamed-resources
