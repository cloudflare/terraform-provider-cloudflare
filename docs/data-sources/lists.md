---
page_title: "cloudflare_lists Data Source - Cloudflare"
subcategory: ""
description: |-
  Data source for looking up Cloudflare Lists.
---

# cloudflare_lists (Data Source)

Data source for looking up Cloudflare Lists.

## Example Usage

```terraform
data "cloudflare_lists" "example" {
  account_id = "01234567890123456789012345678901"
}
```

## Schema

### Required

- `account_id` (String) The account id to target for the resource.

### Read-Only

- `lists` (List of Object) (see [below for nested schema](#nestedatt--lists))
- `id` (String) The ID of this resource.

<a id="nestedatt--lists"></a>
### Nested Schema for `lists`

Read-Only:

- `name` (String)
- `id` (String)
- `description` (String)
- `kind` (String)
- `numitems` (Number)


