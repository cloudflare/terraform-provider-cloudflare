---
page_title: "cloudflare_zone_cache_variants Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zone_cache_variants (Data Source)



## Example Usage

```terraform
data "cloudflare_zone_cache_variants" "example_zone_cache_variants" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) Identifier

### Read-Only

- `editable` (Boolean) Whether the setting is editable
- `id` (String) ID of the zone setting.
Available values: "variants".
- `modified_on` (String) Last time this setting was modified.
- `value` (String) The value of the feature


