---
page_title: "cloudflare_authenticated_origin_pulls_certificates Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_authenticated_origin_pulls_certificates (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) Identifier

### Optional

- `max_items` (Number) Max items to fetch, default: 1000

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Optional:

- `certificate` (String) The zone's leaf certificate.
- `id` (String) Identifier
- `status` (String) Status of the certificate activation.
- `uploaded_on` (String) This is the time the certificate was uploaded.

Read-Only:

- `expires_on` (String) When the certificate from the authority expires.
- `issuer` (String) The certificate authority that issued the certificate.
- `signature` (String) The type of hash used for the certificate.

