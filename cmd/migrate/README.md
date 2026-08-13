# Cloudflare Terraform Provider Migration Tool

A Command-line tool for migrating Terraform configurations and state files for the Cloudflare Terraform Provider.

## Overview

The migration tool automates the transformation of Terraform configurations and state files to handle breaking changes between provider versions. It operates in two phases:

1. **Grit transformations** - Applies predefined patterns for common migrations
2. **Go transformations** - Handles complex structural changes requiring AST manipulation

## Architecture

### Components

- **main.go** - CLI entry point and orchestration
- **grit.go** - Grit pattern application for bulk transformations
- **state.go** - Terraform state file JSON manipulation
- **renames.go** - Attribute renaming logic
- other resource-specific transformations

### Dependencies

- `github.com/hashicorp/hcl/v2/hclwrite` - HCL AST manipulation
- `github.com/zclconf/go-cty/cty` - Type system for HCL values
- `grit` CLI tool - Pattern-based code transformations

## How It Works

### Phase 1: Grit Transformations

Applies Grit patterns in sequence:
1. `cloudflare_terraform_v5` - Main configuration migrations
2. `cloudflare_terraform_v5_attribute_renames_state` - State attribute renames
3. `cloudflare_terraform_v5_resource_renames_configuration` - Resource type renames in config
4. `cloudflare_terraform_v5_resource_renames_state` - Resource type renames in state

Patterns are sourced from:
- GitHub: `github.com/cloudflare/terraform-provider-cloudflare#<pattern_name>`
- Local: Specified via `--patterns-dir` flag

### Phase 2: Go Transformations

#### Configuration Files (.tf)
1. Parses HCL using `hclwrite.ParseConfig()`
2. Applies transformations:
   - **Attribute renames** - Maps old attribute names to new ones
   - **Zone settings splitting** - Converts `cloudflare_zone_settings_override` to individual `cloudflare_zone_setting` resources
   - **Load balancer pool headers** - Restructures header blocks in origins
3. Formats and writes back using `hclwrite.Format()`

#### State Files (.tfstate)
1. Parses JSON state structure
2. Removes empty arrays from load balancer pool attributes (`load_shedding`, `origin_steering`)
3. Preserves state file formatting with `json.MarshalIndent()`

## Command Line Options

```
migrate [options]
```

| Flag | Default | Description |
|------|---------|-------------|
| `--config <dir>` | Current directory | Directory containing .tf files to migrate |
| `--state <file>` | First .tfstate in current dir | State file to migrate |
| `--grit` | true | Enable Grit transformations |
| `--patterns-dir <dir>` | (none) | Local directory containing Grit patterns |
| `--dryrun` | false | Preview changes without modifying files |

### Special Values
- Set `--config false` to skip configuration migration
- Set `--state false` to skip state migration

## Usage Examples

```bash
# Migrate current directory (config and state)
migrate

# Dry run to preview changes
migrate --dryrun

# Migrate specific directory and state file
migrate --config ./terraform --state ./terraform.tfstate

# Use local Grit patterns (for development)
migrate --patterns-dir ../grit-patterns

# Skip Grit, only run Go transformations
migrate --grit=false

# Migrate only configuration files
migrate --state false

# Migrate only state file
migrate --config false
```

## Testing

Test files use `hclwrite` to verify transformations:
- `*_test.go` - Unit tests for each transformation
- `test_helpers.go` - Shared testing utilities

## Integration with Provider

Called automatically by `acctest.MigrationTestStep()` in provider acceptance tests with:
- Working directory set to test directory
- `--grit` flag enabled
- Patterns fetched from GitHub repository

## Prerequisites

- Go 1.21+
- Grit CLI: `npm install -g @getgrit/cli`
- Write permissions to target files