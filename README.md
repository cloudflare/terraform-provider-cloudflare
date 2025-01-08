# Cloudflare Terraform Provider

The Cloudflare Terraform provider

The [Cloudflare Terraform provider](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs) provides convenient access to
[the Cloudflare REST API](https://developers.cloudflare.com/api) from Terraform.

It is generated with [Stainless](https://www.stainlessapi.com/).

## Installation

<!-- x-release-please-start-version -->

```
terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 0.0.1-alpha.0"
    }
  }
}
```

<!-- x-release-please-end -->

And initialize your project by running `terraform init`.

## Requirements

This library requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on Hashicorp's website.

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/cloudflare/terraform-provider-cloudflare/issues) with questions, bugs, or suggestions.
