---
page_title: "cloudflare_api_shield_operation Resource - Cloudflare"
subcategory: ""
description: |-
  Provides a resource to manage an operation in API Shield Endpoint Management.
---

# cloudflare_api_shield_operation (Resource)

Provides a resource to manage an operation in API Shield Endpoint Management.

## Example Usage

```terraform
resource "cloudflare_api_shield_operation" "example" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  method   = "GET"
  host     = "api.example.com"
  endpoint = "/path"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `endpoint` (String) The endpoint which can contain path parameter templates in curly braces, each will be replaced from left to right with `{varN}`, starting with `{var1}`. This will then be [Cloudflare-normalized](https://developers.cloudflare.com/rules/normalization/how-it-works/). **Modifying this attribute will force creation of a new resource.**
- `host` (String) RFC3986-compliant host. **Modifying this attribute will force creation of a new resource.**
- `method` (String) The HTTP method used to access the endpoint. **Modifying this attribute will force creation of a new resource.**
- `zone_id` (String) The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**

### Read-Only

- `id` (String) The ID of this resource.


