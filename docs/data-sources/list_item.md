---
page_title: "cloudflare_list_item Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_list_item (Data Source)



## Example Usage

```terraform
data "cloudflare_list_item" "example_list_item" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  list_id = "2c0fc9fa937b11eaa1b71c4d701ab86e"
  item_id = "34b12448945f11eaa1b71c4d701ab86e"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Defines an identifier.
- `item_id` (String) Defines the unique ID of the item in the List.
- `list_id` (String) The unique ID of the list.

### Read-Only

- `asn` (Number) Defines a non-negative 32 bit integer.
- `comment` (String) Defines an informative summary of the list item.
- `created_on` (String) The RFC 3339 timestamp of when the item was created.
- `hostname` (Attributes) Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the hyphen (-). (see [below for nested schema](#nestedatt--hostname))
- `id` (String) The unique ID of the list.
- `ip` (String) An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.
- `modified_on` (String) The RFC 3339 timestamp of when the item was last modified.
- `redirect` (Attributes) The definition of the redirect. (see [below for nested schema](#nestedatt--redirect))

<a id="nestedatt--hostname"></a>
### Nested Schema for `hostname`

Read-Only:

- `url_hostname` (String)


<a id="nestedatt--redirect"></a>
### Nested Schema for `redirect`

Read-Only:

- `include_subdomains` (Boolean)
- `preserve_path_suffix` (Boolean)
- `preserve_query_string` (Boolean)
- `source_url` (String)
- `status_code` (Number) Available values: 301, 302, 307, 308.
- `subpath_matching` (Boolean)
- `target_url` (String)


