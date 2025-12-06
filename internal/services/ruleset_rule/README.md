# Cloudflare Ruleset Rule Resource

This resource manages individual rules within existing Cloudflare Rulesets. It provides a way to manage rules separately from the parent ruleset, similar to how AWS manages security group rules.

## Key Features

- **Individual Rule Management**: Manage single rules within existing rulesets without affecting other rules
- **Account and Zone Support**: Works with both account-level and zone-level rulesets
- **Full CRUD Operations**: Create, read, update, and delete individual rules
- **Import Support**: Import existing rules using various ID formats
- **No GET Endpoint Handling**: Gracefully handles the lack of individual rule GET endpoint by reading the entire ruleset

## API Endpoints Used

- `POST /{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}/rules` - Create rule
- `PATCH /{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}/rules/{rule_id}` - Update rule  
- `DELETE /{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}/rules/{rule_id}` - Delete rule
- `GET /{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}` - Read ruleset (to find individual rule)

## Import Formats

The resource supports multiple import formats:

1. `{account_or_zone_id}/{ruleset_id}/{rule_id}` - Auto-detects account vs zone
2. `account/{account_id}/{ruleset_id}/{rule_id}` - Explicit account-level
3. `zone/{zone_id}/{ruleset_id}/{rule_id}` - Explicit zone-level

## Example Usage

```hcl
# Create a parent ruleset first
resource "cloudflare_ruleset" "example" {
  zone_id     = "your_zone_id"
  name        = "Example Ruleset"
  description = "Example ruleset for testing"
  kind        = "zone"
  phase       = "http_request_firewall_custom"
}

# Create individual rules within the ruleset
resource "cloudflare_ruleset_rule" "block_bad_ips" {
  zone_id     = "your_zone_id"
  ruleset_id  = cloudflare_ruleset.example.id
  action      = "block"
  expression  = "ip.src in {192.0.2.0/24}"
  description = "Block traffic from bad IP range"
  enabled     = true
}

resource "cloudflare_ruleset_rule" "redirect_example" {
  zone_id     = "your_zone_id"
  ruleset_id  = cloudflare_ruleset.example.id
  action      = "redirect"
  expression  = "http.request.uri.path eq \"/old-path\""
  description = "Redirect old path to new location"
  enabled     = true

  action_parameters {
    from_value {
      status_code = 301
      target_url {
        value = "https://example.com/new-path"
      }
      preserve_query_string = true
    }
  }
}
```

## Implementation Notes

### Handling Missing GET Endpoint

Since Cloudflare doesn't provide a GET endpoint for individual rules, this resource:

1. **On Read**: Fetches the entire parent ruleset and searches for the rule by ID
2. **On Create**: Creates the rule and extracts the ID from the returned ruleset
3. **On Update/Delete**: Uses the rule ID directly with the appropriate endpoints

### State Management

- The resource tracks the parent ruleset's `last_updated` and `version` fields
- If a rule is not found during read, it's considered deleted and removed from state
- The resource requires either `account_id` or `zone_id` to be specified

### Action Parameters

The resource supports comprehensive action parameters for different rule types:

- **Managed Ruleset Execution**: `id`, `ruleset`, `rulesets`, `overrides`
- **URI Rewriting**: `uri.path`, `uri.query` 
- **Header Modifications**: `headers` array with operations
- **Custom Responses**: `response` with status code and content
- **Redirects**: `from_value` with target URL and options
- **Rate Limiting**: `ratelimit` configuration
- **Skip Actions**: `increment` parameter

### Validation

- Validates account/zone ID format (32-character hex)
- Validates ruleset ID format  
- Ensures exactly one of `account_id` or `zone_id` is specified
- Validates action parameter combinations based on action type

## Testing

The resource includes comprehensive tests:

- Basic CRUD operations
- Update scenarios
- Import functionality
- Error handling
- Account vs zone level operations

Run tests with:
```bash
go test ./internal/services/ruleset_rule/...
```
