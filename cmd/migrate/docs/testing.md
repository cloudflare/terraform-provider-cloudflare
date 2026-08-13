## Testing Pattern for Transformations

### Use RunTransformationTests Helper
When writing tests for HCL transformations, **always use the `RunTransformationTests` helper function** from `test_helpers.go` instead of writing custom test logic.

#### Example Test Structure
```go
func TestYourResourceTransformation(t *testing.T) {
    tests := []TestCase{
        {
            Name: "descriptive test name",
            Config: `
resource "cloudflare_your_resource" "example" {
  attribute = "value"
}`,
            Expected: []string{`
resource "cloudflare_your_resource" "example" {
  attribute = "value"
}`},
        },
    }
    
    RunTransformationTests(t, tests, transformFile)
}
```

#### What RunTransformationTests Does
1. Parses the input config using `hclwrite.ParseConfig`
2. Runs it through `transformFile` (which applies both string-level and AST transformations)
3. Formats the output using `hclwrite.Format`
4. Checks that each expected string is contained in the output

#### Important Notes
- The helper automatically handles HCL formatting differences (whitespace, alignment)
- Expected output should include the exact formatting that `hclwrite.Format` produces
- For multiple expected outputs (like split resources), provide multiple strings in the `Expected` array
- Don't write custom parsing, transformation, or comparison logic - let the helper handle it

#### Examples to Follow
- `zone_settings_test.go` - Shows how to test resource splitting transformations
- `renames_test.go` - Shows how to test attribute renaming
- `load_balancer_pool_test.go` - Shows basic transformation testing

## Additional Documentation Resources

### HCL and Terraform Language
* [HCL Language Reference](https://developer.hashicorp.com/terraform/language/syntax/configuration) - Terraform configuration syntax
* [HCL v2 Library Documentation](https://pkg.go.dev/github.com/hashicorp/hcl/v2) - Main HCL Go package
* [hclwrite Package](https://pkg.go.dev/github.com/hashicorp/hcl/v2/hclwrite) - AST manipulation for preserving formatting
* [HCL Native Syntax Specification](https://github.com/hashicorp/hcl/blob/main/hclsyntax/spec.md) - Detailed language spec
* [go-cty Type System](https://pkg.go.dev/github.com/zclconf/go-cty/cty) - Type system used with HCL

### Grit Transformation Engine
* [Grit Documentation](https://docs.grit.io/) - Main documentation site
* [Grit Pattern Language](https://docs.grit.io/language/overview) - Writing transformation patterns
* [Grit CLI Reference](https://docs.grit.io/cli/reference) - Command-line usage
* [GritQL Tutorial](https://docs.grit.io/tutorials/gritql) - Pattern query language

### Terraform Provider Development
* [Provider Framework](https://developer.hashicorp.com/terraform/plugin/framework) - Modern provider development
* [Testing Migration Paths](https://developer.hashicorp.com/terraform/plugin/framework/migrating/testing) - Migration test patterns
* [State Upgrade Functions](https://developer.hashicorp.com/terraform/plugin/framework/resources/state-upgrade) - Handling state schema changes
* [Acceptance Testing](https://developer.hashicorp.com/terraform/plugin/testing/acceptance-tests) - Test framework