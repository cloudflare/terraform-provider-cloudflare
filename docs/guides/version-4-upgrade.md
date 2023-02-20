---
layout: "cloudflare"
page_title: "Upgrading to version 4 (from 3.x)"
description: Terraform Cloudflare Provider Version 4 Upgrade Guide
---

# Terraform Cloudflare Provider Version 4 Upgrade Guide

Version 4 of the Cloudflare Terraform Provider is introducing several breaking
changes to accomodate better user experience and broader resource support for
more customer types.

## Provider Version Configuration

If you are not ready to make a move to version 4 of the Cloudflare provider,
you may keep the 3.x branch active for your Terraform project by specifying:

```hcl
provider "cloudflare" {
  version = "~> 3.0"
  # ... any other configuration
}
```

We highly recommend reviewing this guide, make necessary changes and move to
4.x branch, as further 3.x releases are unlikely to happen outside of critical
security fixes.

Once ready, make the following change to use the latest 4.x release:

```hcl
provider "cloudflare" {
  version = "~> 4.0"
  # ... any other configuration
}
```

To rewrite your HCL configurations, you could use a combination of `grep`/`ripgrep`
and `sed` for simple replacements however we will be providing examples using
[comby] which is a more advanced tool for searching and changing code
structure. NB: the attached examples are intentionally simple and you may want
to make them more specific to suit your use case or environment.

## Provider configuration for `account_id` is removed

In 3.x, `account_id` was available as a provider configuration attribute to
globally define the account you wanted to target for resources. While convenient,
this introduced problems for managing multiple account level configurations as
it was a global configuration stored on the underlying HTTP client.

Instead, all resources now require a `zone_id` or `account_id` attribute
explicitly.

Before:

```hcl
provider "cloudflare" {
  account_id = "..."
}

resource "cloudflare_resource" "example" {
  name = "..."
}
```

After:

```hcl
provider "cloudflare" {}

resource "cloudflare_resource" "example" {
  account_id = "..."
  name = "..."
}
```

Check the specific resource documentation for what resources now require this
attribute.

## `account_id` required for some resources

To accomodate the removal of `account_id` at the provider level, the following
resources now require the `account_id` value to be defined.

- `cloudflare_zone`
- `cloudflare_load_balancer_pool`
- `cloudflare_load_balancer_monitor`
- `cloudflare_account_member`
- `cloudflare_workers_kv_namespace`
- `cloudflare_workers_kv`
- `cloudflare_workers_script`
- `cloudflare_worker_cron_trigger`


## Removed ability to create user level resources

User level resources are an older concept that are superseded by account level
resources. In previous versions, the following resources allowed a fallback in
the event `zone_id` or `account_id` were provided.

- `cloudflare_access_rule`
- `cloudflare_load_balancer`
- `cloudflare_load_balancer_monitor`
- `cloudflare_load_balancer_pool`

In 4.x, these resources do not have automatic fallback to user level resources
but instead, require an explicit `account_id` to use the account level. Zone
level resources remain unchanged.

## `userBaseUrl` is no longer used internally

`userBaseUrl` is a helper method provided by `cloudflare-go` to automatically
detect and build parts of the HTTP request URL if the global `AccountID` value
is provided. While internal, this has been removed in favour of explicit
`account_id` configurations mentioned above.

## `cloudflare_spectrum_application`

- `edge_ips` is now a nested block that holds all edge IP configuration such as
  `type`, `connectivity` and `ips`.
- `edge_ip_connectivity` is now nested under `edge_ips` as `connectivity`.
- `type` is now a required field.

```hcl
resource "cloudflare_spectrum_application" "..." {
  zone_id = "..."
  edge_ip_connectivity = "all"
  edge_ips = ["203.0.113.1", "203.0.113.2"]
}
```

After:

```hcl
resource "cloudflare_spectrum_application" "..." {
  zone_id = "..."
  edge_ips {
    type = "static"
    ips = ["203.0.113.1", "203.0.113.2"]
  }
}
```

## `cloudflare_load_balancer`

- `session_affinity_attributes` has been migrated from `TypeMap` to `TypeSet`.

Before:

```hcl
resource "cloudflare_load_balancer" "..." {
  zone_id = "..."
  session_affinity_attributes = {
    ...
  }
}
```

After:

```hcl
resource "cloudflare_load_balancer" "..." {
  zone_id = "..."
  session_affinity_attributes {
    ...
  }
}
```

There is no automatic schema migration for this change as `TypeMap` does not have
constraints and the current values may have been invalid in previous states.

- `drain_duration` is now an integer (previously a string)

Before:

```hcl
resource "cloudflare_load_balancer" "..." {
  zone_id = "..."
  session_affinity_attributes = {
    secure = "Always"
    drain_duration = "5"
  }
}
```

After:

```hcl
resource "cloudflare_load_balancer" "..." {
  zone_id = "..."
  session_affinity_attributes { # note the type change from above
    secure = "Always"
    drain_duration = 5
  }
}
```

## Renames

### Resources

- `cloudflare_argo_tunnel` is now `cloudflare_tunnel`

## Removals

### Resources

- `cloudflare_access_bookmark`: Use `cloudflare_access_application` configuration
  attributes instead.
- `cloudflare_waf_group`: Use `cloudflare_ruleset` instead.
- `cloudflare_waf_override`: Use `cloudflare_ruleset` instead.
- `cloudflare_waf_package`: Use `cloudflare_ruleset` instead.
- `cloudflare_waf_rule`: Use `cloudflare_ruleset` instead.
- `cloudflare_ip_list`: Use `cloudflare_list` instead.

### Data source

- `cloudflare_waf_groups`: Use `cloudflare_ruleset` instead.
- `cloudflare_waf_packages`: Use `cloudflare_ruleset` instead.
- `cloudflare_waf_rules`: Use `cloudflare_ruleset` instead.
