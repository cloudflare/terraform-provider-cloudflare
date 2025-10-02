# YAML Configuration Guide

This guide explains how to configure migrations using YAML files.

## Basic Structure

Every migration YAML file follows this structure:

```yaml
version: "1.0"
resource_type: cloudflare_access_application
migration:
  from: "4.*"
  to: "5.*"

# Transformations go here
attribute_renames:
  old_name: new_name

# ... more transformations
```

## Required Fields

- `version`: Configuration format version (currently "1.0")
- `resource_type`: The Terraform resource type (e.g., `cloudflare_access_application`)  
- `migration.from`: Source version pattern (supports wildcards like "4.*")
- `migration.to`: Target version pattern

## Transformation Sections

### Simple Transformations

```yaml
# Rename attributes
attribute_renames:
  domain: hostname
  api_key: client_secret

# Remove fields
remove_fields:
  - deprecated_field
  - legacy_option

# Add default values
default_values:
  type: "self_hosted"
  enabled: true
```

### Structural Transformations

```yaml
# Convert blocks to lists
blocks_to_lists:
  - cors_headers
  - allowed_origins

# Convert lists to blocks  
lists_to_blocks:
  - destinations
  - rules

# Complex structural changes
structural_changes:
  - type: flatten_nested
    source: address
    parameters:
      separator: "_"
      max_depth: 2
      
  - type: split_object
    source: config
    parameters:
      attributes:
        - host
        - port
      prefix: "server_"
```

### Conditional Logic

```yaml
# Remove attributes conditionally
conditional_removals:
  - target: auto_redirect
    condition_attr: type
    allowed_values: ["saml", "oidc"]
    
# Different state transformations
state:
  attribute_renames:
    internal_id: resource_id
  schema_version: 2
```

## Complete Example

Here's a real-world migration configuration:

```yaml
version: "1.0"
resource_type: cloudflare_access_application
migration:
  from: "4.*"
  to: "5.*"

# Convert structure first
blocks_to_lists:
  - cors_headers

# Rename attributes
attribute_renames:
  domain: hostname
  zone_id: account_id

# Remove deprecated fields
remove_fields:
  - legacy_auth_domain
  - deprecated_session_duration

# Add required defaults
default_values:
  type: "self_hosted"
  app_launcher_visible: true

# Complex transformations
structural_changes:
  - type: flatten_nested
    source: saml_config
    parameters:
      separator: "_"
      max_depth: 1

# State-specific transformations
state:
  attribute_renames:
    domain: hostname
  schema_version: 2
```

## File Naming Convention

Migration files should follow this naming pattern:

```
resources/<resource_name>/migrations/v<from>_to_v<to>.yaml
```

Examples:
- `v4_to_v5.yaml`
- `v5_to_v6.yaml`
- `v3_to_v4.yaml`

## Testing Your Configuration

Always test your YAML configuration:

1. Create test Terraform configuration
2. Run migration in dry-run mode
3. Verify the output

```bash
# Test with dry-run
./bin/migrate-v2 migrate -f test.tf --dry-run

# Check the preview
./bin/migrate-v2 preview -f test.tf
```

## Tips

1. **Order matters**: Transformations are applied in a specific order regardless of how you write them
2. **Use wildcards**: Version patterns support wildcards (e.g., "4.*" matches 4.0, 4.1, etc.)
3. **Test incrementally**: Test each transformation separately before combining
4. **Document complex logic**: Add comments to explain non-obvious transformations