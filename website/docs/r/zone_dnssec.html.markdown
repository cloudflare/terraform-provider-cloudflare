---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_zone_dnssec"
sidebar_current: "docs-cloudflare-resource-zone-dnssec"
description: |-
  Provides a Cloudflare resource to create and modify a zone DNSSEC.
---

# cloudflare_zone

Provides a Cloudflare Zone DNSSEC resource.

## Example Usage

```hcl
resource "cloudflare_zone" "example" {
    zone = "example.com"
}

resource "cloudflare_zone_dnssec" "example" {
    zone_id = cloudflare_zone.example.id
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The zone id for the zone.

## Attributes Reference

The following attributes are exported:

* `status` - The status of the Zone DNSSEC.
* `flags` - Zone DNSSEC flags.
* `algorithm` - Zone DNSSEC algorithm.
* `key_type` - Key type used for Zone DNSSEC.
* `digest_type` - Digest Type for Zone DNSSEC.
* `digest_algorithm` - Digest algorithm use for Zone DNSSEC.
* `digest` - Zone DNSSEC digest.
* `ds` - DS for the Zone DNSSEC.
* `key_tag` - Key Tag for the Zone DNSSEC.
* `public_key` - Public Key for the Zone DNSSEC.
* `modified_on` - Zone DNSSEC updated time.

## Import

Zone DNSSEC resource can be imported using a zone ID, e.g.

```
$ terraform import cloudflare_zone_dnssec.example d41d8cd98f00b204e9800998ecf8427e
```

where:

* `d41d8cd98f00b204e9800998ecf8427e` - zone ID, as returned from [API](https://api.cloudflare.com/#zone-list-zones)
