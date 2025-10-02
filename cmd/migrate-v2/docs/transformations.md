# Transformation Reference

This guide covers all available transformations for migrating Terraform configurations.

## Common Transformations

These transformations work with both configuration and state files.

### Attribute Renames

Rename attributes while preserving their values.

```yaml
attribute_renames:
  old_name: new_name
  domain: hostname
  api_key: client_secret
```

### Field Removals

Remove deprecated or unnecessary attributes.

```yaml
remove_fields:
  - deprecated_field
  - legacy_option
  - unused_attribute
```

### Default Values

Add default values for new required fields.

```yaml
default_values:
  type: "self_hosted"
  enabled: true
  priority: 1
```

### Conditional Removals

Remove attributes based on conditions.

```yaml
conditional_removals:
  - target: auto_redirect
    condition_attr: type
    allowed_values: ["saml", "oidc"]
```

## Structural Transformations

These transformations change the structure of your configuration.

### Blocks to Lists

Convert HCL blocks to list attributes.

```yaml
blocks_to_lists:
  - cors_headers
  - allowed_methods
```

**Before:**
```hcl
cors_headers {
  allow_credentials = true
  max_age = 3600
}
```

**After:**
```hcl
cors_headers = [{
  allow_credentials = true
  max_age = 3600
}]
```

### Lists to Blocks

Convert list attributes to multiple blocks.

```yaml
lists_to_blocks:
  - destinations
```

**Before:**
```hcl
destinations = [
  { uri = "https://app1.com" },
  { uri = "https://app2.com" }
]
```

**After:**
```hcl
destinations {
  uri = "https://app1.com"
}
destinations {
  uri = "https://app2.com"  
}
```

### Flatten Nested Objects

Flatten nested object structures into flat attributes.

```yaml
structural_changes:
  - type: flatten_nested
    source: address
    parameters:
      separator: "_"
      max_depth: 2
```

**Before:**
```hcl
address = {
  street = {
    name = "Main St"
    number = 123
  }
  city = "San Francisco"
}
```

**After:**
```hcl
address_street_name = "Main St"
address_street_number = 123
address_city = "San Francisco"
```

### Split Objects

Split an object into multiple flat attributes.

```yaml
structural_changes:
  - type: split_object
    source: config
    parameters:
      attributes:
        - host
        - port
      prefix: "server_"
```

**Before:**
```hcl
config = {
  host = "localhost"
  port = 8080
  timeout = 30
}
```

**After:**
```hcl
server_host = "localhost"
server_port = 8080
config = {
  timeout = 30
}
```

### Merge Attributes

Merge multiple attributes into a single object or list.

```yaml
structural_changes:
  - type: merge_attributes
    target: address
    parameters:
      source_attributes:
        - street
        - city
        - state
        - zip
      format: object
```

**Before:**
```hcl
street = "123 Main St"
city = "San Francisco"
state = "CA"
zip = "94102"
```

**After:**
```hcl
address = {
  street = "123 Main St"
  city = "San Francisco"
  state = "CA"
  zip = "94102"
}
```

## Advanced Transformations

### List to Blocks with Mapping

Convert lists to blocks with field renaming.

```yaml
structural_changes:
  - type: list_to_blocks
    source: items
    parameters:
      attribute_map:
        old_field: new_field
        legacy_name: modern_name
```

### Nested Restructuring

Restructure nested attributes with path mapping.

```yaml
structural_changes:
  - type: nested_restructure
    source: config
    parameters:
      paths:
        "database.host": "db_host"
        "database.port": "db_port"
        "cache.ttl": "cache_timeout"
```

## Type Conversions

Convert between different data types.

```yaml
type_conversions:
  port: string_to_int
  enabled: string_to_bool
  tags: csv_to_list
```

Supported conversions:
- `string_to_int`
- `string_to_bool`
- `int_to_string`
- `bool_to_string`
- `csv_to_list`
- `list_to_csv`

## Using Multiple Transformations

Transformations are applied in a specific order:

1. Structural changes (blocks_to_lists, flatten_nested, etc.)
2. Attribute renames
3. Field removals
4. Type conversions
5. Default values

```yaml
# They will be applied in the correct order regardless of how you write them
blocks_to_lists:
  - cors_headers

attribute_renames:
  domain: hostname

remove_fields:
  - deprecated_field

default_values:
  type: "self_hosted"
```