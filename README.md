# Cloudflare Terraform Provider

## Requirements

Terraform 1.2 or newer. We recommend running the [latest version](https://developer.hashicorp.com/terraform/downloads?product_intent=terraform) for optimal compatibility with the Cloudflare provider. Terraform versions older than 1.2 have known issues with newer features and internals.

## Documentation

Full, comprehensive documentation is available on the [Terraform Registry](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs). [API documentation](https://api.cloudflare.com) and [Developer documentation](https://developers.cloudflare.com) is also available
for non-Terraform or service specific information.

## Usage

<!-- x-release-please-start-version -->

```hcl
# Declare the provider and version
terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 5.0.0"
    }
  }
}

# Initialize the provider
provider "cloudflare" {
  api_token = "Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY" # or set CLOUDFLARE_API_TOKEN env variable
}

# Configure a resource
resource "cloudflare_zone" "example_zone" {
  account = {
    id = "023e105f4ecef8ad9ca31a8372d0c353"
  }
  name = "example.com"
  type = "full"
}
```

<!-- x-release-please-end -->

Initialize your project by running `terraform init` in the directory.

## Migrating to Terraform from using the Dashboard

Do you have an existing Cloudflare account (or many!) that you'd like to transition
to be managed via Terraform? Check out [cf-terraforming](https://github.com/cloudflare/cf-terraforming)
which is a tool Cloudflare has built to help dump the existing resources and
import them into Terraform.

## Maintenance

This SDK is actively maintained, however, many issues are tracked outside of GitHub on internal Cloudflare systems. Members of the community are welcome to join and discuss your issues during our weekly triage meetings. For urgent issues, please contact [Cloudflare support](https://www.support.cloudflare.com/s/?language=en_US). 

* [Community triage meeting](https://calendar.google.com/calendar/embed?src=c_dbf6ce250643f2e60f806d28f3fc09a9de24cbe0ab3ffb699838303d2adfc9e4%40group.calendar.google.com&ctz=America%2FLos_Angeles)
