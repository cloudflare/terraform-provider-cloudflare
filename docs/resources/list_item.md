---
page_title: "cloudflare_list_item Resource - Cloudflare"
subcategory: ""
description: |-
  Provides Lists Items (IPs, Redirects) to be used in Edge Rules Engine
  across all zones within the same account.
---

# cloudflare_list_item (Resource)

Provides List Items (IPs, Redirects) to be used in Edge Rules Engine
across all zones within the same account.

~> The Cloudflare Terraform Provider supports list items being specified in-line when creating a `cloudflare_list` resource. The `cloudflare_list_item` resource can only be used on lists that are created with no in-line items or lists managed outside of Terraform. If managing a list via Terraform, set the `cloudflare_list` attribute `ignore_inline_items` to `true` when using the `cloudflare_list_item` resource as it will prevent the list resource from trying to reconcile an empty in-line item list.

## Example Usage

```terraform
# IP List Item
resource "cloudflare_list" "example_ip_list" {
  account_id          = "01234567890123456789012345678901"
  name                = "example_list"
  description         = "example IPs for a list"
  kind                = "ip"
  ignore_inline_items = true
}

resource "cloudflare_list_item" example_ip_item" {
  account_id = "01234567890123456789012345678901"
  list_id    = cloudflare_list.example_ip_list.id
  comment    = "List Item Comment"
  ip         = "192.0.2.0"
}
```
```
# Redirect List Item
resource "cloudflare_list_item" "test_two" {
  account_id = "01234567890123456789012345678901"
  list_id    = cloudflare_list.example_ip_list.id
  redirect {
    source_url       = "https://source.tld"
    target_url       = "https://target.tld"
    status_code      = 302
    subpath_matching = "enabled"
  }
}
```


## Schema

### Required

- `account_id` (String) The account identifier to target for the resource.
- `list_id` (String) The list identifier to target for the resource.

### Optional

- `comment` (String) An optional description of the list.
- `ip` (String) The IP address for lists of `kind = "ip"`.
- `redirect` (Block Set) (see [below for nested schema](#nestedblock--redirect))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--redirect"></a>
### Nested Schema for `redirect`

Required:

- `source_url` (String) The source url of the redirect.
- `target_url` (String) The target url of the redirect.

Optional:

- `include_subdomains` (String) Whether the redirect also matches subdomains of the source url. Available values: `disabled`, `enabled`.
- `preserve_path_suffix` (String) Whether to preserve the path suffix when doing subpath matching. Available values: `disabled`, `enabled`.
- `preserve_query_string` (String) Whether the redirect target url should keep the query string of the request's url. Available values: `disabled`, `enabled`.
- `status_code` (Number) The status code to be used when redirecting a request.
- `subpath_matching` (String) Whether the redirect also matches subpaths of the source url. Available values: `disabled`, `enabled`.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_list_item.example <account_id>/<list_id>/<item_id>
```
