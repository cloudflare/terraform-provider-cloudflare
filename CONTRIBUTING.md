## Setting up the environment

To set up the repository, run:

```sh
$ ./scripts/bootstrap
$ ./scripts/build
```

This will install all the required dependencies and build the Provider binary into the root directory.

You can also [install go 1.22+ manually](https://go.dev/doc/install).

## Running the Provider locally

You can build the provider locally and have your `.tf` files refer to the local build instead of the one in the Hashicorp registry.

First, build the provider binary:

```sh
$ ./scripts/build
```

Then edit (or create) your `~/.terraformrc` to look something like this:

```hcl
  provider_installation {
    dev_overrides {
      "cloudflare/cloudflare" = "/local/path/to/this/repo"
    }
    direct {}
  }
```

## Running tests

To execute the schema and unit tests, run:

```sh
$ ./scripts/test
```

Note that this does not run [acceptance tests](https://developer.hashicorp.com/terraform/plugin/framework/acctests) by default, because
those tests interact with real resources in the cloud and could incur fees.

To enable the running of acceptance tests, use the `TF_ACC` environment variable:

```sh
$ TF_ACC=1 ./scripts/test
```

## Formatting

This library uses the standard gofmt code formatter:

```sh
$ ./scripts/format
```

## Commit Format

We use [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) in this repository. If the change is specific to a single resource, include it in the scope of the commit.

Example commit: `fix(account_member): fix detected drift in post-import refresh plan`

### Overview of Commit Types

| Commit Type | Title                    | Description                                                                                                 |
| ----------- | ------------------------ | ----------------------------------------------------------------------------------------------------------- |
| `feat`      | Features                 | A new feature                                                                                               |
| `fix`       | Bug Fixes                | A bug fix                                                                                                   |
| `docs`      | Documentation            | Documentation only changes                                                                                  |
| `style`     | Styles                   | Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)      |
| `refactor`  | Code Refactoring         | A code change that neither fixes a bug nor adds a feature                                                   |
| `perf`      | Performance Improvements | A code change that improves performance                                                                     |
| `test`      | Tests                    | Adding or modifying tests and/or test data                                                                  |
| `build`     | Builds                   | Changes that affect the build system or external dependencies                                               |
| `ci`        | Continuous Integrations  | Changes to our CI configuration files and scripts                                                           |
| `chore`     | Chores                   | Other changes that don't modify source code or test files                                                   |
| `revert`    | Reverts                  | Reverts a previous commit                                                                                   |
