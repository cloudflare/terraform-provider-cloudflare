---
page_title: "cloudflare_waiting_room_settings Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_waiting_room_settings (Data Source)



## Example Usage

```terraform
data "cloudflare_waiting_room_settings" "example_waiting_room_settings" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) Identifier

### Read-Only

- `search_engine_crawler_bypass` (Boolean) Whether to allow verified search engine crawlers to bypass all waiting rooms on this zone.
Verified search engine crawlers will not be tracked or counted by the waiting room system,
and will not appear in waiting room analytics.

