---
layout: "cloudflare"
page_title: "Upgrading to version 5 (from 4.x)"
description: Terraform Cloudflare Provider Version 5 Upgrade Guide
---

# Terraform Cloudflare Provider Version 5 Upgrade Guide

Version 5 of the Cloudflare Terraform Provider is a ground-up rewrite of the provider, using code generation from our OpenAPI spec.

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
   latest version of 4 to ensure any transitional updates are applied to your
   existing configuration.

Once ready, make the following change to use the latest 5.x release:

```hcl
provider "cloudflare" {
  version = "~> 5"
  # ... any other configuration
}
```

## Automatic migration

For assisting with automatic migrations, we have provided a [GritQL] pattern.
This will allow you to rewrite the parts of your Terraform configuration that have changed automatically. Once you [install Grit], you can run the following
command in the directory where your Terraform configuration is located.

```bash
$ grit apply terraform_cloudflare_v5
```

~> If you are using modules or other dynamic features of HCL, the provided
   codemods may not be as effective. We recommend reviewing the migration notes below to verify all the changes.

<!-- This code block is only used for confirming grit patterns -->
```grit
language hcl

terraform_cloudflare_v5()
```

## Additions

<!-- TODO: grab a dump of all new resources and datasources just before release -->

## Removals

- `cloudflare_zone_settings_override`. Use `cloudflare_zone_setting` instead on a per setting basis.

## Renames

## cloudflare_access_application
## cloudflare_access_group
## cloudflare_access_identity_provider
## cloudflare_access_mutual_tls_hostname_settings
## cloudflare_access_organization
## cloudflare_access_policy

- `application_id` and `precedence` no longer used.

  Before
  ```hcl
  resource "cloudflare_access_policy" "example" {
    account_id      = "f037e56e89293a057740de681ac9abbe"
    application_id  = "foo"
    name            = "example"
    precedence      = 3
  }
  ```

  After
  ```hcl
  resource "cloudflare_access_policy" "example" {
    account_id = "f037e56e89293a057740de681ac9abbe"
    name       = "example"
  }
  ```

## cloudflare_access_rule
## cloudflare_address_map
## cloudflare_api_shield
## cloudflare_api_token
## cloudflare_certificate_pack
## cloudflare_custom_hostname
## cloudflare_custom_ssl
## cloudflare_device_dex_test
## cloudflare_device_managed_networks
## cloudflare_device_posture_integration
## cloudflare_device_posture_rule
## cloudflare_dlp_profile
## cloudflare_email_routing_catch_all
## cloudflare_email_routing_rule
## cloudflare_fallback_domain
## cloudflare_healthcheck
## cloudflare_list
## cloudflare_list_item
## cloudflare_load_balancer
## cloudflare_load_balancer_monitor
## cloudflare_load_balancer_pool
## cloudflare_logpush_job
## cloudflare_managed_headers
## cloudflare_notification_policy
## cloudflare_page_rule
## cloudflare_pages_project
## cloudflare_rate_limit
## cloudflare_record
## cloudflare_risk_behavior
## cloudflare_ruleset
## cloudflare_r2_bucket
## cloudflare_spectrum_application
## cloudflare_split_tunneclououdflare_teams_account
## cloudflare_teams_location
## cloudflare_teams_rule
## cloudflare_tunnel_config
## cloudflare_user_agent_blocking_rule
## cloudflare_waiting_room
## cloudflare_waiting_room_rules
## cloudflare_worker_script
## cloudflare_zone_lockdown

## cloudflare_zone

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
    zone   = "example.com"
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
      zone   = "example.com"
    }
    ```

[GritQL]: https://www.grit.io/
[install Grit]: https://docs.grit.io/cli/quickstart
