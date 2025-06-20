---
layout: "cloudflare"
page_title: "Migrating renamed resources"
description: Guide for handling renamed resources in the Cloudflare Provider
---

## Before you start

- You will need to follow the [Grit CLI install] steps.
- Locate the Grit patterns you'll need to perform the upgrade. They will be in
  `.grit/patterns` directory of the [GitHub repository]. The filename convention
  is `cloudflare_terraform_<version>_<what is changing>_state` for operating on
  the JSON tfstate file and `cloudflare_terraform_<version>_<what is changing>_configuration`
  for changing the HCL configuration. The changelog will call out the
  pattern names as well.
- Lock or ensure no other changes are happening while you are performing the
  migration.
- Make backups of your configuration and state file.

~> It is recommended to perform testing in a non-production, non-critical
environment or a small subset resources before performing the changes to
all resources.

NOTE: use the environment variable `GRIT_MAX_FILE_SIZE_BYTES=0` if the state file is too big and grit errors out.

## Using import

-> Recommended for most users and migrations.

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
2. Apply the **configuration** pattern (replacing the respective parts of the
   filename). This will update your local configuration but not the state file.
   ```bash
   $ grit apply github.com/cloudflare/terraform-provider-cloudflare#cloudflare_terraform_<version>_<what is changing>_configuration
   ```
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

## Manually modifying the state file

-> Recommended for a large number of resources where importing individually is
not feasible.

1. Pull the state file locally using `terraform state pull` and output it to a
   file.

```bash
$ terraform state pull > terraform.tfstate
```

2. Run the **state** pattern (replacing the respective parts of the filename).

```bash
$ grit apply github.com/cloudflare/terraform-provider-cloudflare#cloudflare_terraform_<version>_<what is changing>_state terraform.tfstate
```

3. Navigate to the directory where your HCL configuration is located.
4. Apply the **configuration** pattern (replacing the respective parts of the
   filename).

```bash
$ grit apply github.com/cloudflare/terraform-provider-cloudflare#cloudflare_terraform_<version>_<what is changing>_configuration
```

5. Push your state file back to the remote using `terraform state push terraform.tfstate`.
6. Confirm no further changes are required by running `terraform plan`.

At this point, you've switched over to using the new resource and should be able
to continue using Terraform as normal.

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

[Grit CLI install]: https://docs.grit.io/cli/quickstart
[GitHub repository]: https://github.com/cloudflare/terraform-provider-cloudflare/tree/master/.grit/patterns
