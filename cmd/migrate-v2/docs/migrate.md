# Migration Guide: Porting Resources from V1 to V2 Migrator

This guide documents the process of migrating resources from the old migrator (`cmd/migrate`) to the new migrator (`cmd/migrate-v2`). Use this as a reference when porting additional resources.

## Overview

The V2 migrator uses a YAML-first approach with minimal custom Go code. Most transformations are defined in YAML configuration files, with custom Go code only when absolutely necessary (e.g., resource splitting).

## Step-by-Step Migration Process

### Step 1: Copy and Adapt Existing Tests

The V1 migrator has test cases embedded directly in the test files (e.g., `access_application_test.go`). Extract these test cases:

```go
// V1 test structure (cmd/migrate/access_application_test.go)
tests := []TestCase{
    {
        Name: "transform policies",
        Config: `resource "cloudflare_zero_trust_access_application" "test" {
            domain = "test.example.com"
        }`,
        Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
            hostname = "test.example.com"
        }`},
    },
}
```

Convert these to V2 test files:
- Unit tests: `resources/<resource_name>/migrations/v4_to_v5_test.go`
- Integration test data: Create actual `.tf` files in `integration/testdata/v4/` and `v5/`

### Step 2: Create Directory Structure

Create the following directory structure for each resource:

```
resources/
└── <resource_name>/
    ├── registry.go                    # Resource registration
    ├── migrations/
    │   ├── v4_to_v5.yaml             # YAML configuration
    │   ├── v4_to_v5.go               # Migration wrapper (if needed)
    │   └── v4_to_v5_test.go          # Unit tests
    └── integration/
        ├── integration_test.go        # Integration tests
        └── testdata/
            ├── v4/                    # Input test files
            │   ├── basic.tf
            │   └── complex.tf
            └── v5/                    # Expected output files
                ├── basic.tf
                └── complex.tf
```

### Step 3: Analyze V1 Migration Code

The V1 migrator uses a hybrid approach with both YAML configurations and Go code. V1 files to examine:

#### V1 YAML Configurations
Check these YAML files in `cmd/migrate/transformations/config/`:
- `cloudflare_terraform_v5_attribute_renames_configuration.yaml` - Attribute renames for all resources
- `cloudflare_terraform_v5_block_to_attribute_configuration.yaml` - Block-to-list/map conversions
- `cloudflare_terraform_v5_resource_renames_configuration.yaml` - Resource type renames
- `cloudflare_terraform_v5_attribute_renames_state.yaml` - State attribute renames
- `cloudflare_terraform_v5_resource_renames_state.yaml` - State resource renames

#### V1 Go Code
- `cmd/migrate/<resource_name>.go` - Resource-specific custom transformation logic
- `cmd/migrate/<resource_name>_test.go` - Test cases with before/after examples
- `cmd/migrate/state.go` - State transformation logic
- `cmd/migrate/transformations/` - Generic transformation functions

Things to look for:
1. **In YAML files**: Common transformations already defined (renames, block conversions)
2. **In Go files**: Complex custom logic that can't be expressed in YAML
3. **Patterns to migrate**:
   - Attribute renames (e.g., `domain` → `hostname`)
   - Removed attributes (e.g., `zone_id`)
   - Block-to-list conversions (e.g., `cors_headers` block → list)
   - Default values (e.g., `type = "self_hosted"`)
   - Complex transformations (custom Go code in V1)

### Step 4: Create YAML Configuration

Extract transformations from V1 YAML files and V1 Go code, then combine them into a V2 YAML configuration.

#### Extracting from V1 YAML Files

Example: For `cloudflare_access_application`, check V1 YAML files:

From `cloudflare_terraform_v5_attribute_renames_configuration.yaml`:
```yaml
# V1 format - centralized
cloudflare_access_application:
  domain: hostname  # domain renamed to hostname
```

From `cloudflare_terraform_v5_block_to_attribute_configuration.yaml`:
```yaml
# V1 format - centralized  
cloudflare_access_application:
  to_map:
    - cors_headers
    - saas_app
  to_list:
    - footer_links
    - landing_page_design
```

Convert these to V2 format (resource-specific YAML):
```yaml
# V2 format - in resources/access_application/migrations/v4_to_v5.yaml
resource_type: cloudflare_access_application
source_version: v4
target_version: v5

config:
  attribute_renames:
    domain: hostname
  
  blocks_to_lists:
    - cors_headers
    - saas_app
    - footer_links
    - landing_page_design
```

#### Converting V1 Go Code to V2 YAML

Here's how different V1 Go patterns map to V2 YAML:

#### Attribute Renames
V1 Go code:
```go
renamed := p.renameAttribute(resource, "domain", "hostname")
```

V2 YAML:
```yaml
config:
  attribute_renames:
    domain: hostname
```

#### Removed Attributes
V1 Go code:
```go
removed := p.removeAttribute(resource, "zone_id")
```

V2 YAML:
```yaml
config:
  removed_attributes:
    - zone_id
```

#### Blocks to Lists
V1 Go code:
```go
p.convertBlocksToList(resource, "cors_headers")
```

V2 YAML:
```yaml
config:
  blocks_to_lists:
    - cors_headers
```

#### Default Values
V1 Go code:
```go
p.setDefaultValue(resource, "type", "self_hosted")
```

V2 YAML:
```yaml
config:
  default_values:
    type: "self_hosted"
```

### Step 5: Create Migration Files

#### 5a. Create `migrations/v4_to_v5.yaml`

Example for `access_application`:

```yaml
resource_type: cloudflare_access_application
source_version: v4
target_version: v5
description: Migrate access_application from v4 to v5

config:
  blocks_to_lists:
    - cors_headers
    - saas_app
    - footer_links
    - landing_page_design
  
  attribute_renames:
    domain: hostname
  
  removed_attributes:
    - zone_id
  
  default_values:
    type: "self_hosted"

state:
  attribute_renames:
    domain: hostname
  schema_version: 0
```

#### 5b. Create `migrations/v4_to_v5.go` (Simple Case)

For resources that only need YAML transformations:

```go
package migrations

import (
    _ "embed"
    "fmt"
    
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
)

//go:embed v4_to_v5.yaml
var v4ToV5Config []byte

// RegisterV4ToV5 registers the v4 to v5 migration
func RegisterV4ToV5(registry internal.MigrationRegistry) error {
    migration, err := internal.NewMigration(v4ToV5Config)
    if err != nil {
        return fmt.Errorf("failed to create migration: %w", err)
    }
    
    return registry.Register(migration)
}
```

#### 5c. Create `migrations/v4_to_v5.go` (Complex Case - Resource Splitting)

For resources that need file-level transformations (like `argo`):

```go
package migrations

import (
    _ "embed"
    "fmt"
    
    "github.com/hashicorp/hcl/v2/hclwrite"
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
)

//go:embed v4_to_v5.yaml
var v4ToV5Config []byte

// ArgoMigration extends base migration for resource splitting
type ArgoMigration struct {
    *internal.Migration
}

// RequiresFileTransformation returns true for file-level processing
func (m *ArgoMigration) RequiresFileTransformation() bool {
    return true
}

// MigrateConfig delegates to base migration
func (m *ArgoMigration) MigrateConfig(block *hclwrite.Block, ctx *internal.MigrationContext) error {
    return m.Migration.MigrateConfig(block, ctx)
}

// MigrateState delegates to base migration
func (m *ArgoMigration) MigrateState(state map[string]interface{}, ctx *internal.MigrationContext) error {
    return m.Migration.MigrateState(state, ctx)
}

// RegisterV4ToV5 registers the migration
func RegisterV4ToV5(registry internal.MigrationRegistry) error {
    baseMigration, err := internal.NewMigration(v4ToV5Config)
    if err != nil {
        return fmt.Errorf("failed to create argo migration: %w", err)
    }
    
    migration := &ArgoMigration{
        Migration: baseMigration,
    }
    
    return registry.Register(migration)
}
```

### Step 6: Create Registry File

Create `registry.go` in the resource directory:

```go
package <resource_name>

import (
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources/<resource_name>/migrations"
)

// RegisterMigrations registers all migrations for this resource
func RegisterMigrations(registry internal.MigrationRegistry) error {
    if err := migrations.RegisterV4ToV5(registry); err != nil {
        return err
    }
    
    return nil
}
```

### Step 7: Port Tests

#### 7a. Unit Tests (`migrations/v4_to_v5_test.go`)

Convert V1 test patterns to V2:

```go
package migrations

import (
    "testing"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/utils/testhelpers"
)

func TestMigration_Config(t *testing.T) {
    migration, err := internal.NewMigration(v4ToV5Config)
    require.NoError(t, err)
    
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {
            name: "renames domain to hostname",
            input: `resource "cloudflare_access_application" "example" {
                domain = "example.com"
            }`,
            expected: `resource "cloudflare_access_application" "example" {
                hostname = "example.com"
            }`,
        },
    }
    
    runner := testhelpers.NewMigrationTestRunner(migration, "v4", "v5")
    runner.RunConfigTests(t, tests)
}
```

#### 7b. Integration Tests (`integration/integration_test.go`)

```go
package integration

import (
    "testing"
    
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/utils/testhelpers"
)

func TestAccessApplicationMigration(t *testing.T) {
    test := testhelpers.NewIntegrationTest(
        "cloudflare_access_application",
        "v4", "v5",
    )
    
    test.Run(t)
}
```

### Step 8: Register the Resource (REQUIRED)

**IMPORTANT**: You MUST add the resource to `cmd/migrate-v2/resources/imports.go` for it to be discovered:

```go
package resources

import (
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources/access_application"
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources/argo"
    // Add your new resource import here
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources/<new_resource>"
)

func RegisterAll(registry internal.MigrationRegistry) error {
    if err := access_application.RegisterMigrations(registry); err != nil {
        return err
    }
    
    if err := argo.RegisterMigrations(registry); err != nil {
        return err
    }
    
    // Add registration call for new resource
    if err := <new_resource>.RegisterMigrations(registry); err != nil {
        return err
    }
    
    return nil
}
```

**Note**: Without this step, your migration will not be discovered or executed. This is the ONLY file you need to modify outside of your resource directory.

## Code Style Guidelines

### File Naming
- Use snake_case for file names: `v4_to_v5.yaml`, `v4_to_v5_test.go`
- Test data files should be descriptive: `basic.tf`, `complex.tf`, `with_cors.tf`

### YAML Structure
- Always include all required fields: `resource_type`, `source_version`, `target_version`
- Group transformations logically under `config:` and `state:`
- Use descriptive comments for complex transformations
- Keep indentation consistent (2 spaces)

### Go Code Style
- Embed YAML files using `//go:embed`
- Keep custom Go code minimal
- Always delegate to base migration when possible
- Use clear, descriptive names for custom migration types
- Add comments explaining why custom code is needed

### Test Structure
- Separate unit tests (transformations) from integration tests (full file processing)
- Use table-driven tests for multiple scenarios
- Test both config and state transformations
- Include edge cases and error conditions

## Common Patterns

### Pattern 1: Simple YAML-Only Migration

Most resources only need YAML configuration:
- Attribute renames
- Removed attributes
- Default values
- Blocks to lists conversions

### Pattern 2: Resource Splitting

Some resources (like `argo`) split into multiple resources:

1. Set `requires_file_transformation: true` in YAML
2. Create custom migration type with `TransformFile()` method
3. Implement `file_transformer.go` with splitting logic
4. Generate `moved` blocks for Terraform state migration

### Pattern 3: Complex Value Transformations

Use `value_mappings` in YAML:

```yaml
config:
  value_mappings:
    - attribute: cache_type
      mappings:
        "on": "smart"
        "off": "off"
```

## Testing Checklist

- [ ] Unit tests pass for all transformations
- [ ] Integration tests pass for all test cases
- [ ] Test files include various scenarios (basic, complex, edge cases)
- [ ] State transformations are tested if applicable
- [ ] Resource splitting generates correct `moved` blocks (if applicable)

## Troubleshooting

### Test Failures
1. Check that expected output files have trailing newlines
2. Verify attribute names match exactly (case-sensitive)
3. Ensure HCL formatting matches (use `hclwrite.Format`)

### Migration Not Applied
1. Verify resource type matches exactly in YAML
2. Check that migration is registered in `resources/imports.go`
3. Ensure version strings match (`v4`, `v5`)

### Custom Code Needed
Only use custom Go code when:
1. Resource splitting is required
2. Complex conditional logic can't be expressed in YAML
3. File-level transformations are needed

## V1 vs V2 Architecture Comparison

### V1 Migrator Structure
The V1 migrator (`cmd/migrate/`) uses:
- **Centralized YAML files** in `transformations/config/` - All resources defined in a few large YAML files
- **Resource-specific Go files** - Each resource has custom Go code for complex transformations
- **Mixed approach** - Both YAML and Go code handle transformations

### V2 Migrator Structure
The V2 migrator (`cmd/migrate-v2/`) uses:
- **Resource-specific YAML** - Each resource has its own YAML configuration
- **Minimal custom Go code** - Only for file-level transformations (resource splitting)
- **YAML-first approach** - Everything possible is done in YAML

### Key Differences

| Aspect | V1 Migrator | V2 Migrator |
|--------|------------|-------------|
| YAML Location | `transformations/config/` (shared) | `resources/<name>/migrations/` (per-resource) |
| YAML Structure | One YAML with all resources | One YAML per resource |
| Custom Code | Common, in `<resource>.go` files | Rare, only when YAML insufficient |
| Transformations | `transformations/` directory | Built into internal framework |
| Registration | Hardcoded in main.go | Auto-registration via `imports.go` |

## Available YAML Transformations in V2

The V2 migrator provides these transformations (defined in each resource's YAML):

### Basic Transformations
- `attribute_renames`: Simple field renaming
- `removed_attributes`: Remove deprecated fields
- `default_values`: Add default values for new fields
- `value_mappings`: Transform attribute values

### Structural Transformations
- `blocks_to_lists`: Convert HCL blocks to list attributes
- `resource_splits`: Split one resource into multiple (requires custom code)

### V2 Transformation Framework
V2 transformations are implemented as Go functions in `cmd/migrate-v2/transformations/`:
- `config/basic/`: Simple transformations (renames, defaults, removals)
- `config/structural/`: Complex structure changes (blocks to lists, flattening)
- `config/conditional/`: Conditional transformations
- `state/`: State file transformations

These are automatically applied based on YAML configuration - you don't write Go code to use them.

## Summary

1. **Start with tests**: Copy existing V1 tests and adapt them
2. **YAML first**: Express all transformations in YAML if possible
3. **Minimal custom code**: Only write Go code when YAML isn't sufficient
4. **Test thoroughly**: Include unit and integration tests
5. **Register properly**: Add to `resources/imports.go`

This approach ensures consistency, maintainability, and ease of adding new resource migrations.