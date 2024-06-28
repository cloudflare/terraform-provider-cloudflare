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

## Renames

## Removals

[GritQL]: https://www.grit.io/
[install Grit]: https://docs.grit.io/cli/quickstart
