# Running a non-released version of the Cloudflare Terraform provider

From time to time, you may need to run a non-released version of the provider against your infrastructure. You have may have custom patches, trialing a beta version or just want that bugfix or feature the moment it lands. Below are a few ways to achieve that in various CI/CD providers.

## Using the custom provider in Terraform Cloud

Hashicorp maintain a [help centre article](https://support.hashicorp.com/hc/en-us/articles/360016992613-Using-custom-and-community-providers-in-Terraform-Cloud-and-Enterprise) and an [in-depth walk through](https://www.terraform.io/docs/cloud/run/install-software.html#installing-terraform-providers) on this process.

## Using the custom provider in env0

Refer to [this example env.yml](https://github.com/env0/templates/blob/aab3b93db25cbf79395cec869e1e87a2a493bbd7/community-providers/pingdom/env0.yml). The basic procedure is to build the custom provider, host it somewhere 
accessible to env0 and install it into the environment as a part of your `deploy` 
and `destroy` steps.

## Everything else

Make sure your git repository is at the point in time you want to generate the provider and run `make build`. It will generate the binary which you can then use with [implied local mirror directories](https://www.terraform.io/docs/cli/config/config-file.html#implied-local-mirror-directories) or [explicit installation method configuration](https://www.terraform.io/docs/cli/config/config-file.html#explicit-installation-method-configuration) in your CI/CD environment.

If you'd like to use local overrides, see ["using the provider" section from the development environment setup](development.md#using-the-provider) guide.
