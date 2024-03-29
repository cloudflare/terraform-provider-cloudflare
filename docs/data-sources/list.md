---
page_title: "cloudflare_list Data Source - Cloudflare"
subcategory: ""
description: |-
  Use this data source to lookup a List https://developers.cloudflare.com/api/operations/lists-get-lists.
---

# cloudflare_list (Data Source)

Use this data source to lookup a [List](https://developers.cloudflare.com/api/operations/lists-get-lists).

## Example Usage

```terraform
data "cloudflare_list" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "list_name"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) The account identifier to target for the resource.
- `name` (String) The list name to target for the resource.

### Read-Only

- `description` (String) List description.
- `id` (String) The ID of this resource.
- `kind` (String) List kind.
- `numitems` (Number) Number of items in list.


