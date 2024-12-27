# Generating acceptance tests

For convenience, we have a tool for generating test files inside of the directory of a newly generated service. If the service directory does not exist yet, the service is not configured in the Stainless tooling. Stainless generates the resource, data source, schema, and model files for us, but the acceptance tests and test data have to be added and maintained by Cloudflare.

## Pre-requisites
- The service is onboarded to the Stainless configuration.
- The service exists within this repo, e.g. `internal/services/example_svc`. This is automatic, so long as there is a valid configuration in Stainless.
- You have `go` installed locally.
- The service doesn't already have test files.

## Usage
0. (Optional) Execute the program with `dryrun` enabled to preview the generated file paths:

    `go run cmd/acctestgen/main.go -dryrun example_svc`

1. Run the tool from the root of the repo:

    `go run cmd/acctestgen/main.go example_svc`

2. Confirm the file paths with `y` or `n`.
3. Done! There should be files ready at the paths specified. You can optionally run the tests, but you will need to have configured the necessary environment variables first. You will need access to a test account and zone, as well as a valid API key.

## Files Explained
- `testdata/*.tf`: Terraform configurations used by the acceptance tests.
- `resource_test.go`: Resource acceptance tests, where you can add attribute checks to validate the behavior of a resource.
- `data_source_test.go`: Data source acceptance tests, where you can add attribute checks to validate the behavior of a data source.

Some scaffolding will be auto-generated to help get started, but the tests will fail by default and you will likely have schema validation errors caused by the empty Terraform configuration blocks. To get the tests to pass, first add valid configurations in the `testdata/*.tf` files, then add attribute checks in the `resource.ComposeTestCheckFunc()` of your test cases. You can find many examples in other services that already have acceptance tests.
