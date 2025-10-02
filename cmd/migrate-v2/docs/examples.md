# Migration Examples

Real-world examples of common migration scenarios.

## Simple Attribute Rename

**YAML Configuration:**
```yaml
attribute_renames:
  enabled: is_enabled
```

**Before:**
```hcl
resource "cloudflare_feature" "example" {
  name    = "my-feature"
  enabled = true
}
```

**After:**
```hcl
resource "cloudflare_feature" "example" {
  name       = "my-feature"
  is_enabled = true
}
```

## Blocks to List Conversion

**YAML Configuration:**
```yaml
blocks_to_lists:
  - rule
```

**Before:**
```hcl
resource "cloudflare_firewall" "example" {
  zone_id = var.zone_id
  
  rule {
    action = "block"
    expression = "ip.src eq 192.168.1.1"
  }
  
  rule {
    action = "challenge"
    expression = "ip.geoip.country eq \"CN\""
  }
}
```

**After:**
```hcl
resource "cloudflare_firewall" "example" {
  zone_id = var.zone_id
  
  rules = [
    {
      action = "block"
      expression = "ip.src eq 192.168.1.1"
    },
    {
      action = "challenge"
      expression = "ip.geoip.country eq \"CN\""
    }
  ]
}
```

## Complex Migration Example

**YAML Configuration:**
```yaml
resource_type: cloudflare_access_application
migration:
  from: "4.*"
  to: "5.*"

blocks_to_lists:
  - destinations

attribute_renames:
  domain: hostname

remove_fields:
  - domain_type
  - skip_app_launcher_login_page

default_values:
  type: "self_hosted"
```

**Before:**
```hcl
resource "cloudflare_access_application" "example" {
  zone_id     = var.zone_id
  name        = "Admin Portal"
  domain      = "admin.example.com"
  domain_type = "self_hosted"
  
  destinations {
    uri = "https://admin.example.com"
    path = "/"
  }
  
  destinations {
    uri = "https://admin.example.com"  
    path = "/api"
  }
}
```

**After:**
```hcl
resource "cloudflare_access_application" "example" {
  zone_id  = var.zone_id
  name     = "Admin Portal"
  hostname = "admin.example.com"
  type     = "self_hosted"
  
  destinations = [
    {
      uri = "https://admin.example.com"
      path = "/"
    },
    {
      uri = "https://admin.example.com"
      path = "/api"
    }
  ]
}
```

## Flatten Nested Objects

**YAML Configuration:**
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
resource "cloudflare_example" "test" {
  address = {
    street = {
      name = "Main St"
      number = 123
    }
    city = "San Francisco"
  }
}
```

**After:**
```hcl
resource "cloudflare_example" "test" {
  address_street_name = "Main St"
  address_street_number = 123
  address_city = "San Francisco"
}
```

## State Migration

**YAML Configuration with State Changes:**
```yaml
resource_type: cloudflare_access_application
migration:
  from: "4.*"
  to: "5.*"

# Configuration transformations
attribute_renames:
  session_duration: session_timeout

remove_fields:
  - domain_type

# State-specific transformations  
state:
  attribute_renames:
    session_duration: session_timeout
  schema_version: 2
```

## Complete YAML Example

```yaml
version: "1.0"
resource_type: cloudflare_access_application
migration:
  from: "4.*"
  to: "5.*"

# Simple transformations
attribute_renames:
  domain: hostname
  enabled: is_enabled

remove_fields:
  - deprecated_field
  - legacy_option

default_values:
  type: "self_hosted"
  enabled: true

# Structural transformations
blocks_to_lists:
  - cors_headers
  - destinations

# Conditional removals
conditional_removals:
  - target: auto_redirect
    condition_attr: type
    allowed_values: ["saml", "oidc"]

# State transformations
state:
  attribute_renames:
    internal_id: resource_id
  schema_version: 2
```