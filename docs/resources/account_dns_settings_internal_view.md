---
page_title: "cloudflare_account_dns_settings_internal_view Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_account_dns_settings_internal_view (Resource)



## Example Usage

```terraform
resource "cloudflare_account_dns_settings_internal_view" "example_account_dns_settings_internal_view" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "my view"
  zones = ["372e67954025e0ba6aaa6d586b9e0b59"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier.
- `name` (String) The name of the view.
- `zones` (List of String) The list of zones linked to this view.

### Read-Only

- `created_time` (String) When the view was created.
- `id` (String) Identifier.
- `modified_time` (String) When the view was last modified.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_account_dns_settings_internal_view.example '<account_id>/<view_id>'
```
