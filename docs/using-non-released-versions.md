# Running a non-released version of the Cloudflare Terraform provider

From time to time, you may need to run a non-released version of the provider against your infrastructure. You have may have custom patches, trialing a beta version or just want that bugfix or feature the moment it lands. Below are a few ways to achieve that in various CI/CD providers.

## Using the custom provider in Terraform Cloud

Hashicorp maintain a [help centre article](https://support.hashicorp.com/hc/en-us/articles/360016992613-Using-custom-and-community-providers-in-Terraform-Cloud-and-Enterprise) and an [in-depth walk through](https://www.terraform.io/docs/cloud/run/install-software.html#installing-terraform-providers) on this process.

## Using the custom provider in env0

Refer to [this example env.yml](https://github.com/env0/templates/blob/aab3b93db25cbf79395cec869e1e87a2a493bbd7/community-providers/pingdom/env0.yml). The basic procedure is to build the custom provider, host it somewhere 
accessible to env0 and install it into the environment as a part of your `deploy` 
and `destroy` steps.
