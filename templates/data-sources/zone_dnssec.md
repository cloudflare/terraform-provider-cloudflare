---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_zone_dnssec"
description: Get information on a Cloudflare Zone DNSSEC.
---

# cloudflare_zone_dnssec

Use this data source to look up [Zone][1] DNSSEC settings.

## Example usage

```hcl
data "cloudflare_zone_dnssec" "example" {
  zone_id = "<zone_id>"
}
```

## Argument Reference

- `zone_id` - (Required) The zone id for the zone.

## Attributes Reference

The following attributes are exported:

- `status` - The status of the Zone DNSSEC.
- `flags` - Zone DNSSEC flags.
- `algorithm` - Zone DNSSEC algorithm.
- `key_type` - Key type used for Zone DNSSEC.
- `digest_type` - Digest Type for Zone DNSSEC.
- `digest_algorithm` - Digest algorithm use for Zone DNSSEC.
- `digest` - Zone DNSSEC digest.
- `ds` - DS for the Zone DNSSEC.
- `key_tag` - Key Tag for the Zone DNSSEC.
- `public_key` - Public Key for the Zone DNSSEC.
- `modified_on` - Zone DNSSEC updated time.
