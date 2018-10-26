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
  domain = "${var.cloudflare_zone}"
  name   = "terraform"
  value  = "192.168.0.11"
  type   = "A"
  ttl    = 3600
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required) The DNS zone to add the record to
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
* `metadata` - A key-value map of string metadata cloudflare associates with the record
* `zone_id` - The zone id of the record

## Import

Records can be imported using a composite ID formed of zone name and record ID, e.g.

```
$ terraform import cloudflare_record.default example.com/d41d8cd98f00b204e9800998ecf8427e
```

where:

* `example.com` - the zone name
* `d41d8cd98f00b204e9800998ecf8427e` - record ID as returned by [API](https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records)
