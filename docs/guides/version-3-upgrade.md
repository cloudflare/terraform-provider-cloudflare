---
layout: "cloudflare"
page_title: "Upgrading to version 3 (from 2.x)"
description: Terraform Cloudflare Provider Version 3 Upgrade Guide
---

# Terraform Cloudflare Provider Version 3 Upgrade Guide

Version 3 of the Cloudflare Terraform Provider is introducing several breaking
changes largely to accommodate the underlying upgrade of the `terraform-sdk-plugin`
to version 2.

## Provider Version Configuration

If you are not ready to make a move to version 3 of the Cloudflare provider,
you may keep the 2.x branch active for your Terraform project by specifying:

```hcl
provider "cloudflare" {
  version = "~> 2.0"
  # ... any other configuration
}
```

We highly recommend reviewing this guide, make necessary changes and move to
3.x branch, as further 2.x releases are unlikely to happen.

~> Before attempting to upgrade to version 3, you should first upgrade to the
   latest version of 2 to ensure any transitional updates are applied to your
   existing configuration.

Once ready, make the following change to use the latest 3.x release:

```hcl
provider "cloudflare" {
  version = "~> 3.0"
  # ... any other configuration
}
```

To rewrite your HCL configurations, you could use a combination of `grep`/`ripgrep`
and `sed` for simple replacements however we will be providing examples using
[comby] which is a more advanced tool for searching and changing code
structure. NB: the attached examples are intentionally simple and you may want
to make them more specific to suit your use case or environment.

As some schema changes have been made, state migrators have been included in the
newer versions. As the majority of the changes are `TypeMap` => `TypeList` we
recommend verifying the upgrade yourself on a non-production resource first
as `TypeMap` did include some undocumented behaviour which has been removed in
`terraform-sdk-plugin` v2 and will not be the same across the versions.

## Terraform 0.13 and older versions no longer supported

In line with [HashiCorp's Terraform 1.x compatibility promises],
we will be dropping support for Terraform 0.13 and older within in the
Cloudflare provider. This will allow us to focus on moving the provider forward
with Terraform core. Please be aware, should you raise an issue with the
Cloudflare provider using Terraform core < 0.14, you will be asked to replicate
on a newer version before the issue is triaged by maintainers.

## HTTP user agent changes

While generally treated as internal, we do know of customers having specific
network policies associated with the HTTP user agent produced by the Cloudflare
Terraform Provider. In version 3, the format has changed.

Before: `HashiCorp Terraform/1.0.5 (+https://www.terraform.io) Terraform Plugin SDK/1.17.0 terraform-provider-cloudflare/2.26.1`

After: `terraform/1.0.5 terraform-plugin-sdk/2.7.1 terraform-provider-cloudflare/2.26.1`

## cloudflare_access_rule

- `configuration` is now a `TypeList` instead of a `TypeMap`.

Before:

```hcl
resource "cloudflare_access_rule" "..." {
  configuration = {
    target = "..."
    value  = "..."
  }
}
```

After:

```hcl
resource "cloudflare_access_rule" "..." {
  configuration {
    target = "..."
    value  = "..."
  }
}
```

[comby.live playground URL](https://bit.ly/3ChB8uh)

## cloudflare_custom_hostname

- `status` is now `Computed` as the value isn't managed by an end user.
- `settings` is now `Optional`/`Computed` to reflect the stricter schema
  validation introduced in terraform-plugin-sdk v2.
- `settings.ciphers` is now a `TypeSet` internally to handle suppress ordering
  changes. Schema representation remains the same.

## cloudflare_custom_ssl

- `custom_ssl_options` is now a `TypeList` instead of `TypeMap`.

Before:

```hcl
resource "cloudflare_custom_ssl" "" {
  zone_id = "..."
  custom_ssl_options = {
    certificate = "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n"
    private_key = "-----BEGIN RSA PRIVATE KEY-----\n...\n-----END RSA PRIVATE KEY-----\n"
    bundle_method = "ubiquitous"
    type = "legacy_custom"
  }
}
```

After:

```hcl
resource "cloudflare_custom_ssl" "..." {
  zone_id = "..."
  custom_ssl_options {
    certificate = "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n"
    private_key = "-----BEGIN RSA PRIVATE KEY-----\n...\n-----END RSA PRIVATE KEY-----\n"
    bundle_method = "ubiquitous"
    type = "legacy_custom"
  }
}
```

[comby.live playground URL](https://bit.ly/3C9kEUX)

## cloudflare_load_balancer

- `fixed_response` is now a `TypeList` instead of a `TypeMap`.
- `fixed_response.status_code` is now a `TypeInt` instead of a `TypeString`.

Before:

```hcl
resource "cloudflare_load_balancer" "..." {
  zone_id = "..."
  name    = "..."

  rules {
    name = "..."
    condition = "..."
    fixed_response = {
      message_body = "hello"
      status_code  = "200"
      content_type = "html"
      location     = "www.example.com"
    }
  }
}
```

After:

```hcl
resource "cloudflare_load_balancer" "..." {
  zone_id = "..."
  name    = "..."

  rules {
    name = "..."
    condition = "..."
    fixed_response {
      message_body = "hello"
      status_code  = 200
      content_type = "html"
      location     = "www.example.com"
    }
  }
}
```

[comby.live playground URL](https://bit.ly/3EkySnS)

## cloudflare_record

- `data` is now a `TypeList` instead of a `TypeMap`.

Before:

```hcl
resource "cloudflare_record" "..." {
  zone_id = "..."
  name    = "..."
  type    = "..."

  data = {
    service  = "..."
    proto    = "..."
    name     = "..."
  }
}
```

After:

```hcl
resource "cloudflare_record" "..." {
  zone_id = "..."
  name    = "..."
  type    = "..."

  data {
    service  = "..."
    proto    = "..."
    name     = "..."
  }
}
```

[comby.live playground URL](https://bit.ly/3C9zfj6)

[comby]: https://comby.dev
[hashicorp's terraform 1.x compatibility promises]: https://www.terraform.io/docs/language/v1-compatibility-promises.html
