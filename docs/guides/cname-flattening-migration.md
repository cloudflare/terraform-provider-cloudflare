---
page_title: "Migrate CNAME flattening to DNS settings"
subcategory: "Guides"
description: |-
  Migrate CNAME flattening from cloudflare_zone_setting to cloudflare_zone_dns_settings.
---

# Migrate CNAME flattening to DNS settings

The `cname_flattening` zone setting is no longer supported. Use the `flatten_all_cnames` attribute of `cloudflare_zone_dns_settings` instead.

Map the previous setting value as follows:

| Previous value | `flatten_all_cnames` |
|---|---|
| `on` or `flatten_all` | `true` |
| `off`, `apex`, or `flatten_at_root` | `false` |

## Migrate a standalone setting

Replace the `cloudflare_zone_setting` resource with `cloudflare_zone_dns_settings`, then add a `moved` block. Terraform 1.8 and later transfers the state without calling the deprecated endpoint.

```terraform
resource "cloudflare_zone_dns_settings" "example" {
  zone_id            = var.zone_id
  flatten_all_cnames = true
}

moved {
  from = cloudflare_zone_setting.cname_flattening
  to   = cloudflare_zone_dns_settings.example
}
```

## Merge with existing DNS settings

Only one `cloudflare_zone_dns_settings` resource should manage a zone. If one already exists, add `flatten_all_cnames` to it and remove the old setting from state without changing the remote setting:

```terraform
removed {
  from = cloudflare_zone_setting.cname_flattening

  lifecycle {
    destroy = false
  }
}
```

For Terraform versions that do not support `removed` blocks, run:

```sh
terraform state rm cloudflare_zone_setting.cname_flattening
```

## Migrate data sources

Replace `data.cloudflare_zone_setting` with `data.cloudflare_zone_dns_settings` and use its `flatten_all_cnames` attribute.
