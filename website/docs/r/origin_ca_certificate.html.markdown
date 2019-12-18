---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_origin_ca_certificate"
sidebar_current: "docs-cloudflare-resource-origin-ca-certificate"
description: |-
  Provides a Cloudflare Origin CA certificate resource.
---

# cloudflare_origin_ca_certificate

Provides a Cloudflare Origin CA certificate used to protect traffic to your origin without involving a third party Certificate Authority.

**This resource requires you use your Origin CA Key as the [`api_user_service_key`](../index.html#api_user_service_key).**

## Example Usage

```hcl
# Create a CSR and generate a CA certificate
resource "tls_private_key" "example" {
  algorithm = "RSA"
}

resource "tls_cert_request" "example" {
  key_algorithm   = tls_private_key.example.algorithm
  private_key_pem = tls_private_key.example.private_key_pem

  subject {
    common_name  = ""
    organization = "Terraform Test"
  }
}

resource "cloudflare_origin_ca_certificate" "example" {
  csr                = tls_cert_request.example.cert_request_pem
  hostnames          = [ "example.com" ]
  request_type       = "origin-rsa"
  requested_validity = 7
}
```

## Argument Reference

* `csr`  - (Required) The Certificate Signing Request. Must be newline-encoded.
* `hostnames` - (Required) An array of hostnames or wildcard names bound to the certificate.
* `request_type` - (Required) The signature type desired on the certificate.
* `requested_validity` - (Required) The number of days for which the certificate should be valid.

## Attributes Reference

The following attributes are exported:

* `id` - The x509 serial number of the Origin CA certificate.
* `certificate` - The Origin CA certificate
* `expires_on` - The datetime when the certificate will expire.

## Import

Origin CA certificate resource can be imported using an ID, e.g.

```
$ terraform import cloudflare_origin_ca_certificate.example 276266538771611802607153687288146423901027769273
```
