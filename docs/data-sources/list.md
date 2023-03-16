---
page_title: "cloudflare_list Data Source - Cloudflare"
subcategory: ""
description: |-
  Data source for looking up a Cloudflare List.
---

# cloudflare_list (Data Source)

Data source for looking up a Cloudflare List.

## Example Usage

```terraform
data "cloudflare_list" "example" {
  account_id = "01234567890123456789012345678901"
  name       = "list_name"
}
```

## Schema

### Required

- `account_id` (String) The account id to target for the resource.
- `name` (String) The name of the list for the resource.

### Read-Only

- `id` (String) The ID of this resource.
- `name` (String)
- `description` (String)
- `kind` (String)
- `numitems` (Number)


