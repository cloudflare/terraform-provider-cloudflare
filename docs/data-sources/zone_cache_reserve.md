---
page_title: "cloudflare_zone_cache_reserve Data Source - Cloudflare"
subcategory: ""
description: |-
  Provides a Cloudflare data source to look up Cache Reserve
  status for a given zone.
---

# cloudflare_zone_cache_reserve (Data Source)

Provides a Cloudflare data source to look up Cache Reserve
status for a given zone.

## Example Usage

```terraform
data "cloudflare_zone_cache_reserve" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) The zone identifier to target for the resource.

### Read-Only

- `enabled` (Boolean) The status of Cache Reserve support.
- `id` (String) The ID of this resource.


