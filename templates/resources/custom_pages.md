---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_custom_pages"
description: Provides a resource which manages Cloudflare custom pages.
---

# cloudflare_custom_pages

Provides a resource which manages Cloudflare custom error pages.

## Example Usage

```hcl
resource "cloudflare_custom_pages" "basic_challenge" {
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  type    = "basic_challenge"
  url     = "https://example.com/challenge.html"
  state   = "customized"
}
```

## Argument Reference

The following arguments are supported:

- `zone_id` - (Optional) The zone ID where the custom pages should be
  updated. Either `zone_id` or `account_id` must be provided.
- `account_id` - (Optional) The account ID where the custom pages should be
  updated. Either `account_id` or `zone_id` must be provided. If
  `account_id` is present, it will override the zone setting.
- `type` - (Required) The type of custom page you wish to update. Must
  be one of `basic_challenge`, `waf_challenge`, `waf_block`,
  `ratelimit_block`, `country_challenge`, `ip_block`, `under_attack`,
  `500_errors`, `1000_errors`, `always_online`, `managed_challenge`.
- `url` - (Required) URL of where the custom page source is located.
- `state` - (Required) Managed state of the custom page. Must be one of
  `default`, `customized`. If the value is `default` it will be removed
  from the Terraform state management.

## Import

Custom pages can be imported using a composite ID formed of:

- `customPageLevel` - Either `account` or `zone`.
- `identifier` - The ID of the account or zone you intend to manage.
- `pageType` - The value from the `type` argument.

Example for a zone:

```
$ terraform import cloudflare_custom_pages.basic_challenge zone/d41d8cd98f00b204e9800998ecf8427e/basic_challenge
```

Example for an account:

```
$ terraform import cloudflare_custom_pages.basic_challenge account/e268443e43d93dab7ebef303bbe9642f/basic_challenge
```
