---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_record"
sidebar_current: "docs-cloudflare-resource-record"
description: |-
  Provides a Cloudflare record resource.
---

# cloudflare_record

Provides a Cloudflare record resource.

> **Hands-on:** Try the [Host a Static Website with S3 and Cloudflare](https://learn.hashicorp.com/tutorials/terraform/provider-release-publish?utm_source=WEBSITE&utm_medium=WEB_IO&utm_offer=ARTICLE_PAGE&utm_content=DOCS) tutorial on HashiCorp Learn. In this tutorial, you will set up a static website using AWS S3 as an object store and Cloudflare for DNS, SSL and CDN (using `cloudflare_record`), then create Cloudflare page rules to always redirect HTTPS and temporarily redirect certain paths.

## Example Usage

```hcl
# Add a record to the domain
resource "cloudflare_record" "foobar" {
  zone_id = var.cloudflare_zone_id
  name    = "terraform"
  value   = "192.168.0.11"
  type    = "A"
  ttl     = 3600
}

# Add a record requiring a data map
resource "cloudflare_record" "_sip_tls" {
  zone_id = var.cloudflare_zone_id
  name    = "_sip._tls"
  type    = "SRV"

  data = {
    service  = "_sip"
    proto    = "_tls"
    name     = "terraform-srv"
    priority = 0
    weight   = 0
    port     = 443
    target   = "example.com"
  }
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone ID to add the record to
* `name` - (Required) The name of the record
* `type` - (Required) The type of the record
* `value` - (Optional) The (string) value of the record. Either this or `data` must be specified
* `data` - (Optional) Map of attributes that constitute the record value. Primarily used for LOC and SRV record types. Either this or `value` must be specified
* `ttl` - (Optional) The TTL of the record ([automatic: '1'](https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record))
* `priority` - (Optional) The priority of the record
* `proxied` - (Optional) Whether the record gets Cloudflare's origin protection; defaults to `false`.

## Attributes Reference

The following attributes are exported:

* `id` - The record ID
* `hostname` - The FQDN of the record
* `proxiable` - Shows whether this record can be proxied, must be true if setting `proxied=true`
* `created_on` - The RFC3339 timestamp of when the record was created
* `modified_on` - The RFC3339 timestamp of when the record was last modified
* `metadata` - A key-value map of string metadata Cloudflare associates with the record

## Import

Records can be imported using a composite ID formed of zone ID and record ID, e.g.

```
$ terraform import cloudflare_record.default ae36f999674d196762efcc5abb06b345/d41d8cd98f00b204e9800998ecf8427e
```

where:

* `ae36f999674d196762efcc5abb06b345` - the zone ID
* `d41d8cd98f00b204e9800998ecf8427e` - record ID as returned by [API](https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records)
