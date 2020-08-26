---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_certificate_pack"
sidebar_current: "docs-cloudflare-resource-certificate-pack"
description: |-
  Provides a Cloudflare Certificate Pack resource.
---

# cloudflare_certificate_pack

Provides a Cloudflare Certificate Pack resource that is used to provision
managed TLS certificates.

~> **Important:** Certificate packs are not able to be updated in place and if
you require a zero downtime rotation, you need to use Terraform's meta-arguments
for [`lifecycle`](https://www.terraform.io/docs/configuration/resources.html#lifecycle-lifecycle-customizations) blocks.
`create_before_destroy` should be suffice for most scenarios (exceptions are
things like missing entitlements, high ranking domain). To completely
de-risk rotations, use you can create multiple resources using a 2-phase change
where you have both resources live at once and you remove the old one once
you've confirmed the certificate is available.

## Example Usage

```hcl
resource "cloudflare_certificate_pack" "dedicated_custom_example" {
  zone_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  type    = "dedicated_custom"
  hosts   = ["example.com", "sub.example.com"]
}

# Advanced certificate manager for DigiCert
resource "cloudflare_certificate_pack" "advanced_example_for_digicert" {
  zone_id               = "1d5fdc9e88c8a8c4518b068cd94331fe"
  type                  = "advanced"
  hosts                 = ["example.com", "sub.example.com"]
  validation_method     = "txt"
  validity_days         = 30
  certificate_authority = "digicert"
  cloudflare_branding   = false
}

# Advanced certificate manager for Let's Encrypt
resource "cloudflare_certificate_pack" "advanced_example_for_lets_encrypt" {
  zone_id               = "1d5fdc9e88c8a8c4518b068cd94331fe"
  type                  = "advanced"
  hosts                 = ["example.com", "*.example.com"]
  validation_method     = "http"
  validity_days         = 90
  certificate_authority = "lets_encrypot"
  cloudflare_branding   = false
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone to which the certificate pack should be added.
* `type` - (Required) Certificate pack configuration type.
  Allowed values: `"custom"`, `"dedicated_custom"`, `"advanced"`.
* `hosts` - (Required) List of hostnames to provision the certificate pack for.
  Note: If using Let's Encrypt, you cannot use individual subdomains and only a
  wildcard for subdomain is available.
* `validation_method` - (Optional based on `type`) Which validation method to
  use in order to prove domain ownership. Allowed values: `"txt"`, `"http"`, `"email"`.
* `validity_days` - (Optional based on `type`) How long the certificate is valid
  for. Note: If using Let's Encrypt, this value can only be 90 days.
  Allowed values: 14, 30, 90, 365.
* `certificate_authority` - (Optional based on `type`) Which certificate
  authority to issue the certificate pack. Allowed values: `"digicert"`,
  `"lets_encrypt"`.
* `cloudflare_branding` - (Optional based on `type`) Whether or not to include
  Cloudflare branding. This will add `sni.cloudflaressl.com` as the Common Name
  if set to `true`.

## Import

Certificate packs can be imported using a composite ID of the zone ID and
certificate pack ID. This isn't recommended and it is advised to replace the
certificate entirely instead.

```
$ terraform import cloudflare_certificate_pack.example cb029e245cfdd66dc8d2e570d5dd3322/8fda82e2-6af9-4eb2-992a-5ab65b792ef1
```
