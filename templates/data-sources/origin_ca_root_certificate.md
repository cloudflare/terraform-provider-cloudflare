---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_origin_ca_root_certificate"
description: Get Cloudflare Origin CA root certificate.
---

# cloudflare_origin_ca_root_certificate

Use this data source to get the [Origin CA root certificate][1] for a given algorithm.

## Example Usage

```hcl
data "cloudflare_origin_ca_root_certificate" "origin_ca" {
  algorithm = "<algorithm>"
}
```

## Arguments Reference

- `algorithm` - (Required) The name of the algorithm used when creating an Origin CA certificate. Currently-supported values are "rsa" and "ecc" (case-insensitive).

## Attributes Reference

- `cert_pem` - The Origin CA root certificate in PEM format.

[1]: https://developers.cloudflare.com/ssl/origin-configuration/origin-ca#4-required-for-some-add-cloudflare-origin-ca-root-certificates
