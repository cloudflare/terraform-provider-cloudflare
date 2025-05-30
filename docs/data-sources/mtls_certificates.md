---
page_title: "cloudflare_mtls_certificates Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_mtls_certificates (Data Source)



## Example Usage

```terraform
data "cloudflare_mtls_certificates" "example_mtls_certificates" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier.

### Optional

- `max_items` (Number) Max items to fetch, default: 1000

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `ca` (Boolean) Indicates whether the certificate is a CA or leaf certificate.
- `certificates` (String) The uploaded root CA certificate.
- `expires_on` (String) When the certificate expires.
- `id` (String) Identifier.
- `issuer` (String) The certificate authority that issued the certificate.
- `name` (String) Optional unique name for the certificate. Only used for human readability.
- `serial_number` (String) The certificate serial number.
- `signature` (String) The type of hash used for the certificate.
- `uploaded_on` (String) This is the time the certificate was uploaded.


