---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_record"
sidebar_current: "docs-cloudflare-resource-record"
description: |-
  Provides a Cloudflare record resource.
---

# cloudflare_record

Provides a Cloudflare record resource.

## Example Usage

```hcl
# Add a record to the domain
resource "cloudflare_record" "foobar" {
  domain = "${var.cloudflare_domain}"
  name   = "terraform"
  value  = "192.168.0.11"
  type   = "A"
  ttl    = 3600
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required) The domain to add the record to
* `name` - (Required) The name of the record
* `value` - (Required) The value of the record
* `type` - (Required) The type of the record
* `ttl` - (Optional) The TTL of the record ([automatic: '1'](https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record))
* `priority` - (Optional) The priority of the record
* `proxied` - (Optional) Whether the record gets Cloudflare's origin protection.

## Attributes Reference

The following attributes are exported:

* `id` - The record ID
* `name` - The name of the record
* `value` - The value of the record
* `type` - The type of the record
* `ttl` - The TTL of the record
* `priority` - The priority of the record
* `hostname` - The FQDN of the record
* `proxied` - (Optional) Whether the record gets Cloudflare's origin protection; defaults to `false`.
* `created_on` - (Computed) The RFC3339 timestamp of when the record was created
* `data` - (Computed) A key-value map of strings showing individual values for SRV and LOC record types
* `modified_on` - (Computed) The RFC3339 timestamp of when the record was last modified
* `metadata` - (Computed) A key-value map of string metadata cloudflare associates with the record
* `proxiable` - (Computed) Shows whether this record can be proxied, must be true if setting `proxied=true`
* `zone_id` - (Computed) the zone id of the record
