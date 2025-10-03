# Adding New Resource Migrations

To add a new resource migration to the migrate-v2 tool, follow these steps:

## 1. Create Resource Directory Structure

```
resources/
├── your_resource/
│   ├── registry.go           # Registration function
│   ├── integration/           # Integration tests
│   │   ├── integration_test.go
│   │   └── testdata/
│   │       ├── v4/           # Input test files
│   │       └── v5/           # Expected output files
│   └── migrations/
│       ├── v4_to_v5.yaml     # Migration configuration
│       ├── v4_to_v5.go       # Migration implementation (if needed)
│       └── v4_to_v5_test.go  # Unit tests
```

## 2. Create Migration Configuration (YAML)

Create `migrations/v4_to_v5.yaml` with your migration rules:

```yaml
resource_type: cloudflare_your_resource
source_version: v4
target_version: v5
description: Migrate your_resource from v4 to v5

config:
  attribute_renames:
    old_name: new_name
  
  removed_attributes:
    - deprecated_field
  
  # Add other transformations as needed

state:
  attribute_renames:
    old_name: new_name
  schema_version: 1
```

## 3. Create Registry File

Create `registry.go` in your resource directory:

```go
package your_resource

import (
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources/your_resource/migrations"
)

// RegisterMigrations registers all migrations for the your_resource resource
func RegisterMigrations(registry internal.MigrationRegistry) error {
    // Register v4 to v5 migration
    if err := migrations.RegisterV4ToV5(registry); err != nil {
        return err
    }
    
    return nil
}
```

## 4. Register in imports.go

Add your resource to `resources/imports.go`:

```go
import (
    // ... existing imports ...
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources/your_resource"
)

func RegisterAll(registry internal.MigrationRegistry) error {
    // ... existing registrations ...
    
    if err := your_resource.RegisterMigrations(registry); err != nil {
        return err
    }
    
    return nil
}
```

## 5. Add Tests

Create unit tests in `migrations/v4_to_v5_test.go` and integration tests in `integration/integration_test.go`.

## Special Cases

### Resource Splitting

If your resource needs to be split into multiple resources (like `cloudflare_argo` → `cloudflare_argo_smart_routing` + `cloudflare_tiered_cache`), you'll need:

1. Set `requires_file_transformation: true` in your YAML
2. Implement a custom migration wrapper with `TransformFile()` method
3. See the `argo` resource for a complete example

### Complex Transformations

For transformations that can't be expressed in YAML:

1. Create a custom migration struct that embeds `*internal.Migration`
2. Override the necessary methods (`MigrateConfig`, `MigrateState`, etc.)
3. See `access_application` or `argo` for examples

## Testing

Run tests for your new resource:

```bash
# Unit tests
go test ./resources/your_resource/migrations/...

# Integration tests  
go test ./resources/your_resource/integration/...

# Build to ensure compilation
go build -o ./bin/migrate-v2 .
```

## Notes

- The registration is now automatic - just add the import and RegisterAll call
- No need to modify main.go directly
- Each resource manages its own migrations independently
- YAML-first approach: use custom Go code only when YAML isn't sufficient