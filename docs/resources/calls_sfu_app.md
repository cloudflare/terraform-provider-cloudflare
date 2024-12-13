---
page_title: "cloudflare_calls_sfu_app Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_calls_sfu_app (Resource)



## Example Usage

```terraform
resource "cloudflare_calls_sfu_app" "example_calls_sfu_app" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "production-realtime-app"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) The account identifier tag.

### Optional

- `app_id` (String) A Cloudflare-generated unique identifier for a item.
- `name` (String) A short description of Calls app, not shown to end users.

### Read-Only

- `created` (String) The date and time the item was created.
- `modified` (String) The date and time the item was last modified.
- `secret` (String) Bearer token
- `uid` (String) A Cloudflare-generated unique identifier for a item.

