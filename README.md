# Cloudflare Terraform Provider

## Requirements

Terraform 1.2 or newer. We recommend running the [latest version](https://developer.hashicorp.com/terraform/downloads?product_intent=terraform) for optimal compatibility with the Cloudflare provider. Terraform versions older than 1.2 have known issues with newer features and internals.

## Documentation

Full, comprehensive documentation is available on the [Terraform Registry](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs). [API documentation](https://api.cloudflare.com) and [Developer documentation](https://developers.cloudflare.com) is also available
for non-Terraform or service specific information.

## Usage

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

## Migrating to Terraform from using the Dashboard

Do you have an existing Cloudflare account (or many!) that you'd like to transition
to be managed via Terraform? Check out [cf-terraforming](https://github.com/cloudflare/cf-terraforming)
which is a tool Cloudflare has built to help dump the existing resources and
import them into Terraform.
