---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_zone"
sidebar_current: "docs-cloudflare-datasource-zone"
description: |- Get information on a Cloudflare Zone.
---

# cloudflare_zone

Use this data source to look up [Zone][1] info. This is the singular alternative to `cloudflare_zones`.

## Example usage

Given you have the following zones in Cloudflare.

- example.com

```hcl
# Look for a single zone that you know exists using an exact match.
data "cloudflare_zone" "example" {
  name = "example.com"
}

# you can also lookup by zone_id
data "cloudflare_zone" "example" {
  zone_id = "<zone_id>"
}
```

### Example usage with other resources

The example below fetches the zone information for example.com and then is referenced in the `cloudflare_record`
section.

```hcl
data "cloudflare_zone" "example" {
  name = "example.com"
}

resource "cloudflare_record" "example" {
  zone_id     = cloudflare_zone.example.zone_id # or shorthand cloudflare_zone.example.id
  name        = "www"
  value       = "203.0.113.1"
  type        = "A"
  proxied     = true
}
```

## Argument Reference

-> **Note:** It's only required specify **one of** `zone_id` or `name`.

- `zone_id` - (Optional) The zone ID.
- `name` - (Optional) The name of the zone.

## Attributes Reference

- `id` - The zone ID. Same value as `zone_id`.
- `zone_id` - The zone ID.
- `account_id` - The account ID associated with the zone.
- `name` - The name of the zone.
- `status` - Status of the zone. Values can be: `active`, `pending`, `initializing`, `moved`, `deleted`,
  or `deactivated`.
- `paused` - `true` if cloudflare is enabled on the zone, otherwise `false`.
- `plan` - The name of the plan associated with the zone.
- `name_servers` - Cloudflare-assigned name servers. This is only populated for zones that use Cloudflare DNS.
- `vanity_name_servers` - List of Vanity Nameservers (if set).

[1]: https://api.cloudflare.com/#zone-properties
