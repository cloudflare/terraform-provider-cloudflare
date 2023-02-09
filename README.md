# Cloudflare Terraform Provider

## Quickstarts

- [Getting started with Cloudflare and Terraform](https://developers.cloudflare.com/terraform/installing)
- [Developing the provider](contributing/development.md)

## Documentation

Full, comprehensive documentation is available on the [Terraform Registry](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs). [API documentation](https://api.cloudflare.com) and [Developer documentation](https://developers.cloudflare.com) is also available
for non-Terraform or service specific information.

## Migrating to Terraform from using the Dashboard

Do you have an existing Cloudflare account (or many!) that you'd like to transition
to be managed via Terraform? Check out [cf-terraforming](https://github.com/cloudflare/cf-terraforming)
which is a tool Cloudflare has built to help dump the existing resources and
import them into Terraform.

## Version 4.x early release candidates

> **Warning** Release candidates may contain bugs and backwards incompatible state modifications. **You should not use it in production you are clear on the ramifications and have a clear backup plan in the event of breakages.**<br><br>For production usage, the 3.x release is recommended using the `~> 3` provider version selector.

We are working on releasing the next major version of the Cloudflare Terraform Provider and want your help! 

If you have suitable workloads and would like to test out the next release before everyone else, you can opt-in by updating your provider `version` to explicitly match one of the release candidate versions ([`~>`, `>` or `>=` will not work](https://developer.hashicorp.com/terraform/language/expressions/version-constraints#version-constraint-behavior)). See the [releases](https://github.com/cloudflare/terraform-provider-cloudflare/releases) page for available versions.

```hcl
terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "4.0.0-rc1"
    }
  }
}
```

Be sure to check out the [version 4 upgrade guide](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/guides/version-4-upgrade) and make any modifications. If you hit bugs, please [open a new issue](https://github.com/cloudflare/terraform-provider-cloudflare/issues/new/choose).

## Contributing

To contribute, please read the [contribution guidelines](contributing/README.md).

## Feedback

If you would like to provide feedback (not a bug or feature request) on the Cloudflare Terraform provider, you're welcome to via [this form](https://forms.gle/6ofUoRY2QmPMSqoR6).
