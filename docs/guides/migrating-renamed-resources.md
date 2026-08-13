---
layout: "cloudflare"
page_title: "Migrating renamed resources"
description: Guide for handling renamed resources in the Cloudflare Provider
---

## Before you start

-> For v4 to v5 migrations, the recommended approach is the
[version 5 migration guide](version-5-migration) which uses built-in state
upgraders that handle both configuration and state migration automatically.

This guide covers general strategies for handling renamed resources across
provider versions.

- Lock or ensure no other changes are happening while you are performing the
  migration.
- Make backups of your configuration and state file.

~> It is recommended to perform testing in a non-production, non-critical
environment or a small subset resources before performing the changes to
all resources.

## Using tf-migrate (Automatic)

-> Recommended for v4 to v5 migrations.

For automatic configuration migrations, use [tf-migrate], the official
Cloudflare Terraform Provider migration tool.

**Installation:**

Download the latest release from the [tf-migrate releases] page, or install via Go:

```bash
go install github.com/cloudflare/tf-migrate/cmd/tf-migrate@latest
```

**Migrate your configuration:**

```bash
# Preview changes (dry run)
tf-migrate migrate --dry-run --source-version v4 --target-version v5

# Apply the migration
tf-migrate migrate --source-version v4 --target-version v5
```

`tf-migrate` handles:
- Resource type renames (e.g., `cloudflare_access_application` → `cloudflare_zero_trust_access_application`)
- Attribute renames and restructuring
- Block structure changes

After running `tf-migrate`, run `terraform plan` to allow the v5 provider's
built-in state upgraders to automatically migrate your state.

## Using import (Manual)

-> Recommended when tf-migrate doesn't support a specific resource or for
granular control.

We'll assume we're migrating from the `cloudflare_old` resource to the
`cloudflare_new` resource and is applicable for any changes where the changes
are 1:1 drop in replacements.

Example:

```hcl
resource "cloudflare_dependant_thing" "example" {
  zone_id = var.staging_zone_id
  name    = "Example Thing"
  enabled = true
}

resource "cloudflare_old" "example" {
  name           = "Example"
  other_thing_id = cloudflare_dependant_thing.example.id
}
```

1. Navigate to the directory where your HCL configuration is located.
2. Update your configuration to use the new resource name (manually or with tf-migrate).
3. List the resources in state using `terraform state list`.
   ```
   $ terraform state list
   cloudflare_dependant_thing.example
   cloudflare_old.example
   ```
4. Remove the **old** resource using `terraform state rm <resource>`.
5. Invoke `terraform import` to import the new resource (refer to the resource
   documentation for the correct import string).
   ```
   $ terraform import cloudflare_new.example 0da42c8d2132a9ddaf714f9e7c920711/8295065e70782f633792396e73af14bb
   ```
6. Confirm no further changes are required by running `terraform plan`.

At this point, you've switched over to using the new resource and should be able
to continue using Terraform as normal.

## Using moved blocks (Terraform 1.1+)

-> Recommended for Terraform 1.1+ users who want to track renames in code.

Terraform's native `moved` blocks provide a declarative way to handle resource
renames without manual state manipulation.

Example:

```hcl
# Tell Terraform the resource was renamed
moved {
  from = cloudflare_old.example
  to   = cloudflare_new.example
}

resource "cloudflare_new" "example" {
  zone_id = var.staging_zone_id
  name    = "Example Thing"
  enabled = true
}
```

When you run `terraform plan`, Terraform will automatically move the state
from the old address to the new address.

~> `moved` blocks work for address changes but don't handle attribute renames.
For v4 to v5 migrations with attribute changes, the provider's built-in state
upgraders handle this automatically.

## Two phase swap over

-> Recommend for incremental swapping or introducing new resources with older
resources.

We'll assume we're migrating from the `cloudflare_old` resource to the
`cloudflare_new` resource as a dependent resource.

Example:

```hcl
resource "cloudflare_old" "example" {
  zone_id = var.staging_zone_id
  name    = "Example Thing"
  enabled = true
}

resource "cloudflare_thing" "example" {
  name           = "Example"
  other_thing_id = cloudflare_old.example.id
}
```

1. Add your new resource to your configuration.

```hcl
resource "cloudflare_old" "example" {
  zone_id = var.staging_zone_id
  name    = "Example Thing"
  enabled = true
}

resource "cloudflare_new" "example" {
  zone_id = var.staging_zone_id
  name    = "Example Thing"
  enabled = true
}

resource "cloudflare_thing" "example" {
  name           = "Example"
  other_thing_id = cloudflare_old.example.id
}
```

2. Update your existing resource to reference the newer resource.

```hcl
resource "cloudflare_old" "example" {
  zone_id = var.staging_zone_id
  name    = "Example Thing"
  enabled = true
}

resource "cloudflare_new" "example" {
  zone_id = var.staging_zone_id
  name    = "Example Thing"
  enabled = true
}

resource "cloudflare_thing" "example" {
  name           = "Example"
  other_thing_id = cloudflare_new.example.id
}
```

3. Remove the old resource once it is no longer referenced.

```hcl
resource "cloudflare_new" "example" {
  zone_id = var.staging_zone_id
  name    = "Example Thing"
  enabled = true
}

resource "cloudflare_thing" "example" {
  name           = "Example"
  other_thing_id = cloudflare_new.example.id
}
```

At this point, you've switched over to using the new resource and should be able
to continue using Terraform as normal.

[tf-migrate]: https://github.com/cloudflare/tf-migrate
[tf-migrate releases]: https://github.com/cloudflare/tf-migrate/releases
