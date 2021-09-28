---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_zone"
sidebar_current: "docs-cloudflare-datasource-zone"
description: Get information on a Cloudflare Zone.
---

# cloudflare_zone

Use this data source to look up [zone] info. This is the singular alternative
to `cloudflare_zones`.

~> **Note** Cloudflare zone names **are not unique**. It is possible for multiple
  accounts to have the same zone created but in different states. If you are
  using this setup, it is advised to use the `account_id` attribute on this
  resource or swap to `cloudflare_zones` to further filter the results.

## Example usage

```hcl
# Look for a single zone that you know exists using an exact match.
data "cloudflare_zone" "example" {
  name = "example.com"
}

# Look for a zone in a specific account by the zone name.
data "cloudflare_zone" "example" {
  name       = "example.com"
  account_id = "d41d8cd98f00b204e9800998ecf8427e"
}

# You can also lookup by zone_id if you prefer.
data "cloudflare_zone" "example" {
  zone_id = "0b6d347b01d437a092be84c2edfce72c"
}
```

### Example usage with other resources

The example below fetches the zone information for example.com and then is
referenced in the `cloudflare_record` section.

```hcl
data "cloudflare_zone" "example" {
  name = "example.com"
}

resource "cloudflare_record" "example" {
  zone_id = cloudflare_zone.example.id
  name    = "www"
  value   = "203.0.113.1"
  type    = "A"
  proxied = true
}
```

## Argument Reference

- `zone_id` - (Optional) The zone ID. Conflicts with `"name"`.
- `name` - (Optional) The name of the zone. Conflicts with `"zone_id"`.

## Attributes Reference

- `id` - The zone ID.
- `account_id` - The account ID associated with the zone.
- `name` - The name of the zone.
- `status` - Status of the zone. Values can be: `"active"`, `"pending"`, `"initializing"`, `"moved"`, `"deleted"`,
  or `"deactivated"`.
- `paused` - `true` if cloudflare is enabled on the zone, otherwise `false`.
- `plan` - The name of the plan associated with the zone.
- `name_servers` - Cloudflare assigned name servers. This is only populated for zones that use Cloudflare DNS.
- `vanity_name_servers` - List of Vanity Nameservers (if set).

[zone]: https://api.cloudflare.com/#zone-properties
