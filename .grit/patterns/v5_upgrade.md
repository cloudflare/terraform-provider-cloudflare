---
layout: "cloudflare"
page_title: "Upgrading to version 5 (from 4.x)"
description: Terraform Cloudflare Provider Version 5 Upgrade Guide
---

# Terraform Cloudflare Provider Version 5 Upgrade Guide

Version 5 of the Cloudflare Terraform Provider is a ground-up rewrite of the provider, using code generation from our OpenAPI spec.

```grit
language hcl

or {
  terraform_cloudflare_v5(),
  `provider "cloudflare" { $provider }` where {
    $provider <: contains `version = $old` => `version = "~> 5"`,
    $old <: not includes "5"
  }
}
```

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

~> If you are using modules or other dynamic features of HCL, the provided
codemods may not be as effective. We recommend reviewing the migration notes below to verify all the changes.

## Block attributes

All blocks used for configuration have been converted to attributes, which must be set with an `=` sign.

For example, the `config` block in the `cloudflare_device_posture_integration` resource must be converted from this:

```hcl
resource "cloudflare_device_posture_integration" "example" {
  # old stuff
  config {
    api_url       = "https://example.com/api"
    auth_url      = "https://example.com/connect/token"
    client_id     = "client-id"
    client_secret = "client-secret"
  }
}
```

Afterwards it will look like this:

```hcl
resource "cloudflare_device_posture_integration" "example" {
  # old stuff
  config = {
    api_url       = "https://example.com/api"
    auth_url      = "https://example.com/connect/token"
    client_id     = "client-id"
    client_secret = "client-secret"
  }
}
```

## Renames

## Removals

[GritQL]: https://www.grit.io/
[install Grit]: https://docs.grit.io/cli/quickstart
