---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_zone"
sidebar_current: "docs-cloudflare-datasource-zone"
description: |-
  Get information on a Cloudflare Zone.
---

# cloudflare_zone

Use this data source to look up a [Zone][1] record.

## Example Usage

```hcl
data "cloudflare_zone" "test" {
  zone = "example.com"
}

resource "cloudflare_spectrum_application" "ssh" {
  zone_id  = "${data.cloudflare_zone.test.id}"
  protocol = "tcp/22"

  dns = {
    type = "CNAME"
    name = "ssh.${data.cloudflare_zone.test.zone}"
  }

  origin_direct = ["tcp://81.120.102.10:23"]
  origin_port   = 22
}
```

## Argument Reference

- `zone` - (Required) The name of the zone to match.

## Attributes Reference

- `id`           - Id of the zone
- `status`       - Zone status. Valid values: active, pending, initializing, moved, deleted, deactivated and read only
- `name_servers` - List of name servers
- `type`         - Zone type. Valid values: full or partial

[1]: https://api.cloudflare.com/#zone-properties
