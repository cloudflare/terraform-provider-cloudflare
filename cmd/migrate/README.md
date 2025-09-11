# Cloudflare Terraform Provider Migration Tool

A Command-line tool for migrating Terraform configurations and state files for the Cloudflare Terraform Provider.

## Overview

The migration tool automates the transformation of Terraform configurations and state files to handle breaking changes between provider versions. It operates in two phases:

1. **YAML-based transformations** - Applies transformation rules defined in YAML configuration files
2. **Go transformations** - Handles complex structural changes requiring AST manipulation

## Architecture

### Components

- **main.go** - CLI entry point and orchestration
- **transformations.go** - YAML-based transformation engine
- **state.go** - Terraform state file JSON manipulation
- **renames.go** - Attribute renaming logic
- **transformations/config/** - YAML transformation configuration files
- other resource-specific transformations

### Dependencies

- `github.com/hashicorp/hcl/v2/hclwrite` - HCL AST manipulation
- `github.com/zclconf/go-cty/cty` - Type system for HCL values
- `gopkg.in/yaml.v3` - YAML configuration parsing

## How It Works

### Phase 1: YAML-based Transformations

Applies transformation rules from YAML configuration files:
1. `cloudflare_terraform_v5_attribute_renames_configuration.yaml` - Configuration attribute renames
2. `cloudflare_terraform_v5_attribute_renames_state.yaml` - State attribute renames
3. `cloudflare_terraform_v5_resource_renames_configuration.yaml` - Resource type renames in config
4. `cloudflare_terraform_v5_resource_renames_state.yaml` - Resource type renames in state
5. `cloudflare_terraform_v5_block_to_attribute_configuration.yaml` - Block to attribute conversions

Transformations are loaded from:
- Default: Embedded configs from GitHub repository
- Custom: Local directory via `--transformer-dir` flag

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
| `--transformer-dir <dir>` | Embedded configs | Path to directory containing transformer YAML configs |
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

# Use local transformer configs (for development)
migrate --transformer-dir ./transformations/config

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
- YAML transformations applied from embedded configs
- Go transformations applied for complex structural changes

## Prerequisites

- Go 1.21+
- Write permissions to target files