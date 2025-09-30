# Cloudflare Zone Tagging Proof of Concept

This proof of concept demonstrates two different approaches to implementing tagging support for Cloudflare zones in the Terraform provider.

## Option A: Inline Tags Support

Modified existing `cloudflare_zone` resource to support tags directly.

### Files Modified:
- `internal/services/zone/schema.go` - Added tags map attribute
- `internal/services/zone/model.go` - Added Tags field to ZoneModel
- `internal/services/zone/resource.go` - Added tag handling in CRUD operations
- `internal/services/zone/resource_test.go` - Added test example

### Usage:
```hcl
resource "cloudflare_zone" "example_zone" {
  account = {
    id = "023e105f4ecef8ad9ca31a8372d0c353"
  }
  name = "example.com"
  type = "full"
  tags = {
    env = "production"
    CostCenter = "4300-systems-engineering"
  }
}
```

### Benefits:
- Single resource manages both zone and tags
- Consistent with other Terraform providers
- Tags are part of zone lifecycle

### Drawbacks:
- Tags are tightly coupled to zone resource
- Cannot manage tags independently
- Must recreate entire resource if zone needs to be recreated

## Option B: Separate Zone Tags Resource

Created new `cloudflare_zone_tags` resource for independent tag management.

### Files Created:
- `internal/services/zone_tags/schema.go` - Zone tags resource schema
- `internal/services/zone_tags/model.go` - ZoneTagsModel struct
- `internal/services/zone_tags/resource.go` - Full CRUD implementation
- `internal/services/zone_tags/resource_test.go` - Test examples

### Usage:
```hcl
resource "cloudflare_zone" "example_zone" {
  account = {
    id = "023e105f4ecef8ad9ca31a8372d0c353"
  }
  name = "example.com"
  type = "full"
}

resource "cloudflare_zone_tags" "example_zone_tags" {
  zone_id = cloudflare_zone.example_zone.id
  tags = {
    env = "production"
    CostCenter = "4300-systems-engineering"
  }
}
```

### Benefits:
- Independent tag lifecycle management
- Can manage tags for externally created zones
- More flexible for complex scenarios
- Clear separation of concerns

### Drawbacks:
- Requires two resources for full setup
- More complex configuration
- Additional state to manage

## Assumed API Functions

Both implementations assume these Cloudflare API functions exist:

```go
// Update/set zone tags
client.Zones.UpdateTags(ctx, zones.ZoneUpdateTagsParams{
    ZoneID: cloudflare.F(zoneID),
    Tags:   map[string]string{"key": "value"},
})

// Get zone tags
client.Zones.GetTags(ctx, zones.ZoneGetTagsParams{
    ZoneID: cloudflare.F(zoneID),
})
```

## Implementation Notes

- Uses Terraform Plugin Framework v1.15.0 patterns
- Proper type handling with `types.Map` for tags
- Standard CRUD lifecycle management
- Resource state management
- Import functionality (Option B)
- Comprehensive test examples (non-runnable for demonstration)

## Next Steps

1. Register new `cloudflare_zone_tags` resource in provider configuration
2. Add actual Cloudflare API implementation once endpoints are available
3. Add comprehensive validation and error handling
4. Update documentation and examples
5. Add acceptance tests once API is available

## Usage Examples

See test files for complete working examples:
- Option A: `internal/services/zone/resource_test.go`
- Option B: `internal/services/zone_tags/resource_test.go`