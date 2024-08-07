---
page_title: "cloudflare_zero_trust_access_custom_page Resource - Cloudflare"
subcategory: ""
description: |-
  Provides a resource to customize the pages your end users will see
  when trying to reach applications behind Cloudflare Access.
---

# cloudflare_zero_trust_access_custom_page (Resource)

Provides a resource to customize the pages your end users will see
when trying to reach applications behind Cloudflare Access.

## Example Usage

```terraform
resource "cloudflare_zero_trust_access_custom_page" "example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "example"
  type        = "forbidden"
  custom_html = "<html><body><h1>Forbidden</h1></body></html>"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Friendly name of the Access Custom Page configuration.
- `type` (String) Type of Access custom page to create. Available values: `identity_denied`, `forbidden`.

### Optional

- `account_id` (String) The account identifier to target for the resource. Conflicts with `zone_id`. **Modifying this attribute will force creation of a new resource.**
- `app_count` (Number) Number of apps to display on the custom page.
- `custom_html` (String) Custom HTML to display on the custom page.
- `zone_id` (String) The zone identifier to target for the resource. Conflicts with `account_id`. **Modifying this attribute will force creation of a new resource.**

### Read-Only

- `id` (String) The ID of this resource.


