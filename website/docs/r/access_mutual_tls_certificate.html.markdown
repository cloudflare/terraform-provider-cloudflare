---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_access_mutual_tls_certificate"
sidebar_current: "docs-cloudflare-resource-access-mutual-tls-certificate"
description: |-
  Provides a Cloudflare Access Mutual TLS Certificate resource.
---

# cloudflare_access_mutual_tls_certificate

Provides a Cloudflare Access Mutual TLS Certificate resource. Mutual TLS authentication ensures that the traffic is secure and trusted in both directions between a client and server and can be used with Access to only allows requests from devices with a corresponding client certificate.

## Example Usage

```hcl
resource "cloudflare_access_mutual_tls_certificate" "my_cert" {
  zone_id              = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name                 = "My Root Cert"
  certificate          = var.ca_pem
  associated_hostnames = ["staging.example.com"]
}
```

## Argument Reference

The following arguments are supported:

-> **Note:** It's required that an `account_id` or `zone_id` is provided and in most cases using either is fine. However, if you're using a scoped access token, you must provide the argument that matches the token's scope. For example, an access token that is scoped to the "example.com" zone needs to use the `zone_id` argument.

* `account_id` - (Optional) The account to which the certificate should be added. Conflicts with `zone_id`.
* `zone_id` - (Optional) The DNS zone to which the certificate should be added. Conflicts with `account_id`.
* `name` - (Required) The name of the certificate.
* `certificate` - (Required) The Root CA for your certificates.
* `associated_hostnames` - (Optional) The hostnames that will be prompted for this certificate.

## Attributes Reference

The following additional attributes are exported:

* `id` - ID of the Access Mutual TLS Certificate resource

## Import

Access Mutual TLS Certificate can be imported using a composite ID composed of the account or zone and the mutual TLS certificate ID in the form of: `account/ACCOUNT_ID/MUTUAL_TLS_CERTIFICATE_ID` or `zone/ZONE_ID/MUTUAL_TLS_CERTIFICATE_ID`.

```
$ terraform import cloudflare_access_mutual_tls_certificate.staging account/cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e
```
