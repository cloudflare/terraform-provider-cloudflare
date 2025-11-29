# Sweepers Guide

This document describes the standardized sweeper implementation for the Cloudflare Terraform Provider v5.

## Overview

Sweepers are cleanup functions that delete test resources from your Cloudflare account after tests run. All test resources must follow a standard naming convention to ensure sweepers can identify and clean them up without affecting production resources.

## Test Resource Naming Convention

**All test resources MUST use the prefix: `cf-tf-test-`**

### Generating Test Resource Names

Use the `utils.GenerateRandomResourceName()` function to generate test resource names:

```go
rnd := utils.GenerateRandomResourceName()
// Returns: "cf-tf-test-abcdefghij" (10 random lowercase letters)
```

This ensures all test resources follow the standard naming convention.

### Legacy Prefix

The old prefix `tf-acctest-` is temporarily supported during migration but will be removed in the future. Do not use it for new tests.

## Implementing Sweepers

Every `internal/services/<resource_name>/resource_test.go` file should implement a sweeper.

### Standard Sweeper Pattern

```go
package resource_name_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/<service>"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_<resource>", &resource.Sweeper{
		Name: "cloudflare_<resource>",
		F:    testSweepCloudflare<Resource>,
	})
}

func testSweepCloudflare<Resource>(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	// Get required IDs from environment
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// OR: zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	if accountID == "" {
		return fmt.Errorf("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	// List resources using v6 SDK
	resources, err := client.<Service>.<Resource>.List(ctx, <service>.<Resource>ListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		return fmt.Errorf("failed to list resources: %w", err)
	}

	// Filter and delete test resources
	for _, res := range resources.Result {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(res.Name) {
			continue
		}

		_, err := client.<Service>.<Resource>.Delete(ctx, res.ID, <service>.<Resource>DeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			// Log error but continue sweeping other resources
			fmt.Printf("Failed to delete %s (%s): %s\n", res.Name, res.ID, err)
			continue
		}
		fmt.Printf("Deleted %s: %s (%s)\n", "<resource>", res.Name, res.ID)
	}

	return nil
}
```

### Key Points

1. **Use `utils.ShouldSweepResource(name)`** - Standard filtering function that checks for test prefixes
2. **Continue on errors** - Don't fail the entire sweep if one resource fails to delete
3. **Log deletions** - Print what was deleted for debugging
4. **Check environment variables** - Verify required IDs are set before listing resources

## Filtering Helpers

The `internal/utils/sweeper_helpers.go` package provides helper functions:

- `IsTestResource(name string) bool` - Checks if name starts with `cf-tf-test-`
- `IsLegacyTestResource(name string) bool` - Checks if name starts with `tf-acctest-`
- `ShouldSweepResource(name string) bool` - Returns true if resource should be deleted
  - In normal mode: Returns `true` only for resources with test prefixes (`cf-tf-test-*` or `tf-acctest-*`)
  - In danger mode: Returns `true` for ALL resources when `SWEEP_DANGEROUSLY_DELETE_ALL=true` environment variable is set

## Running Sweepers

### Via Script

Use the provided sweep script to run sweepers:

```bash
# Sweep all resources in an account
./scripts/sweep --account <account_id>

# Sweep a specific resource
./scripts/sweep --account <account_id> --resource account_token

# Sweep resources in a specific zone
./scripts/sweep --account <account_id> --zone <zone_id> --resource dns_record

# Dry-run mode (preview what would be deleted)
./scripts/sweep --account <account_id> --resource dns_record --dry-run

# List all available sweepers
./scripts/sweep --list
```

### ⚠️ Dangerous Mode: Delete ALL Resources

**WARNING: This is an extremely dangerous feature that bypasses all safety checks!**

The `--dangerously-delete-resources` flag disables resource name validation and will delete **EVERY** resource of the specified type, including production resources.

```bash
# DANGEROUS: Delete ALL resources (not just test resources)
./scripts/sweep --account <test_account_id> \
  --dangerously-delete-resources \
  --resource dns_record
```

**Safety Features:**
- Prominent warning messages are displayed
- 2-second delay before execution
- Works with `--dry-run` for preview
- Should only be used on isolated test accounts

**When to use:**
- Cleaning up dedicated test accounts
- Resetting isolated test zones
- Environments with no production resources

**When NOT to use:**
- Production accounts
- Shared accounts
- Any account with production resources
- Without explicit confirmation

**Best practice:** Always use `--dry-run` first:

```bash
# Preview what would be deleted
./scripts/sweep --account <test_account_id> \
  --dangerously-delete-resources \
  --resource dns_record \
  --dry-run

# If you're absolutely sure, run without --dry-run
./scripts/sweep --account <test_account_id> \
  --dangerously-delete-resources \
  --resource dns_record
```

**How it works:**
- Sets the `SWEEP_DANGEROUSLY_DELETE_ALL` environment variable
- The `utils.ShouldSweepResource()` function checks this variable
- When set, returns `true` for ALL resources (bypassing name checks)

For complete documentation on danger mode, including implementation details and testing, see [DANGEROUS_SWEEP_MODE.md](../DANGEROUS_SWEEP_MODE.md).

## Linting

A linter is available to detect sweeper and naming convention violations:

```bash
# Run the sweeper linter
./scripts/lint-sweepers

# Or run all lints (includes sweeper linter)
./scripts/lint
```

The linter checks for:

1. **Missing sweeper registration** - All `resource_test.go` files should have a sweeper
2. **Legacy filtering patterns** - Sweepers should use `utils.ShouldSweepResource()`

Currently, the linter only produces **warnings** and does not fail CI. This will change to errors once all resources are migrated.

## Example: account_token Sweeper

This is a complete, working example from `internal/services/account_token/resource_test.go`:

```go
func init() {
	resource.AddTestSweepers("cloudflare_account_token", &resource.Sweeper{
		Name: "cloudflare_account_token",
		F:    testSweepCloudflareAccountToken,
	})
}

func testSweepCloudflareAccountToken(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	// List all API tokens
	tokens, err := client.Accounts.Tokens.List(ctx, accounts.TokenListParams{})
	if err != nil {
		return fmt.Errorf("failed to fetch account tokens: %w", err)
	}

	// Delete test tokens (those created by our tests)
	// Uses utils.ShouldSweepResource() to filter by standard test naming convention
	for _, token := range tokens.Result {
		if !utils.ShouldSweepResource(token.Name) {
			continue
		}

		_, err := client.Accounts.Tokens.Delete(ctx, token.ID, accounts.TokenDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			// Log error but continue sweeping other resources
			fmt.Printf("Failed to delete account token %s (%s): %s\n", token.Name, token.ID, err)
			continue
		}
		fmt.Printf("Deleted account token: %s (%s)\n", token.Name, token.ID)
	}

	return nil
}
```

## Migration Plan

### Current Status

- ✅ Helper functions created (`utils.ShouldSweepResource()`, etc.)
- ✅ `GenerateRandomResourceName()` updated to use `cf-tf-test-` prefix
- ✅ Linter implemented and integrated into CI
- ✅ Pilot resource (`account_token`) migrated
- ⏳ Remaining resources need migration

### Migration Process

For each resource:

1. Update the sweeper function to use `utils.ShouldSweepResource()`
2. Verify tests use `utils.GenerateRandomResourceName()` (most already do)
3. Run the sweeper to test: `./scripts/sweep --account <test_account> --resource <name>`
4. Run the linter: `./scripts/lint-sweepers`

## Best Practices

1. **Never delete production resources** - Always filter by test prefix
2. **Test your sweepers** - Run them on a test account before merging
3. **Handle errors gracefully** - Continue sweeping even if one resource fails
4. **Use the v6 SDK** - Prefer cloudflare-go/v6 over legacy v1 SDK
5. **Be specific with environment variables** - Check for required IDs early
6. **Add logging** - Print what's being deleted for debugging

## Troubleshooting

### Sweeper not finding resources

- Verify environment variables are set (`CLOUDFLARE_ACCOUNT_ID`, `CLOUDFLARE_ZONE_ID`)
- Check that test resources are being created with the correct prefix
- Run linter to detect naming violations: `./scripts/lint-sweepers`

### Sweeper failing with permissions error

- Ensure your API token has the necessary permissions
- Some resources require account-level permissions, others zone-level

## References

- [Terraform Testing Framework - Sweepers](https://developer.hashicorp.com/terraform/plugin/testing/acceptance-tests/sweepers)
- Sweep script: `scripts/sweep`
- Linter implementation: `internal/tools/sweeper-lint/`
- Helper functions: `internal/utils/sweeper_helpers.go`
