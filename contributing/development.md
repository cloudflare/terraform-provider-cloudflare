# Development Environment Setup

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.14+ (to run acceptance tests)
- [Go](https://golang.org/doc/install) 1.18 (to build the provider plugin)

## Quick Start

If you wish to work on the provider, you'll first need [Go](http://www.golang.org)
installed on your machine (version 1.18+ is _required_). You'll also need to
correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well
as adding `$GOPATH/bin` to your `$PATH`.

See above for which option suits your workflow for building the provider.

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

To run a subset of the acceptance test suite, you can run

```sh
# SDKv2 example
TESTARGS='-run "^<regex target of tests>" -count 1 -parallel 1' make testacc
TESTARGS='-run "^TestAccTestPagesProject" -count 1 -parallel 1' make testacc

# Framework example
TEST=./internal/framework/service/rulesets TESTARGS='-run "^<regex target of tests>" -count 1' make testacc
TEST=./internal/framework/service/rulesets TESTARGS='-run "^TestAccCloudflareRuleset_" -count 1' make testacc
```

You can also install other optional (but great to have tools) using `make tools`.
Most of these tools run in CI automatically but helps having these locally to
either hook into your editor or debug CI failures.

## Using the Provider

With Terraform v0.14 and later, [development overrides for provider developers](https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers) can be leveraged in order to use the provider built from source.

To do this, populate a Terraform CLI configuration file (`~/.terraformrc` for
all platforms other than Windows; `terraform.rc` in the `%APPDATA%` directory
when using Windows) with at least the following options:

```
provider_installation {
  dev_overrides {
    "cloudflare/cloudflare" = "<GOPATH>/src/github.com/cloudflare/terraform-provider-cloudflare"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

You will need to replace `<GOPATH>` with the **full path** to your GOPATH where
the repository lives, no `~` shorthand.

Once you have this file in place, you can run `make build-dev` which will
build a development version of the binary in the repository that Terraform
will use instead of the version from the remote registry.

## Pulling in unreleased changes from Go library (cloudflare-go)

In some scenarios, you may be wanting to make changes in `cloudflare-go` and have
them reflected in the Terraform provider before they are released upstream. To do
this, you can update the `go.mod` file locally (do not commit it) using `go mod replace`.

```
replace github.com/cloudflare/cloudflare-go => ../cloudflare-go
```

Updating the second path to wherever the local library sits on your machine.
