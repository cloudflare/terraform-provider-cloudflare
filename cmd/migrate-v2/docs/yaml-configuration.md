# YAML Configuration Reference

## Basic Structure

```yaml
resource_type: cloudflare_<resource_name>
source_version: v4
target_version: v5
description: Brief description of the migration

config:
  # Configuration transformations
  
state:
  # State transformations
```

## Configuration Transformations

### Simple Transformations

```yaml
config:
  # Rename attributes
  attribute_renames:
    old_name: new_name
    domain: hostname
  
  # Remove deprecated fields
  removed_attributes:
    - deprecated_field
    - unused_attribute
  
  # Add default values
  default_values:
    type: "self_hosted"
    enabled: true
```

### Structural Transformations

```yaml
config:
  # Convert blocks to lists
  blocks_to_lists:
    - cors_headers
    - allowed_idps
  
  # Value mappings (transform values)
  value_mappings:
    - attribute: cache_type
      mappings:
        "on": "smart"
        "off": "off"
```

### Resource Splitting

For splitting one resource into multiple:

```yaml
requires_file_transformation: true

config:
  resource_splits:
    - source_resource: cloudflare_argo
      splits:
        - when_attribute_exists: smart_routing
          create_resource: cloudflare_argo_smart_routing
          attribute_mappings:
            smart_routing: value
        - when_attribute_exists: tiered_caching
          create_resource: cloudflare_tiered_cache
          attribute_mappings:
            tiered_caching: cache_type
      fallback:
        create_resource: cloudflare_argo_smart_routing
        set_attributes:
          value: "off"
```

## State Transformations

```yaml
state:
  # Same options as config
  attribute_renames:
    old_field: new_field
  
  # Schema version update
  schema_version: 1
```

## Complete Example

```yaml
resource_type: cloudflare_access_application
source_version: v4
target_version: v5
description: Migrate access_application from v4 to v5

config:
  blocks_to_lists:
    - cors_headers
    - saas_app
  
  attribute_renames:
    domain: hostname
  
  removed_attributes:
    - zone_id
  
  default_values:
    type: "self_hosted"

state:
  attribute_renames:
    domain: hostname
  schema_version: 1
```

## Tips

- Order matters: transformations apply in the sequence defined
- Test with `--dry-run` first
- Use `--verbose` to see detailed transformation steps
- Keep YAML simple - use custom Go code only when necessary