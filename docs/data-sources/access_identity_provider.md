---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_access_identity_provider"
description: Get information on a Cloudflare Access Identity Provider by name.
---

# cloudflare_access_identity_provider

Use this data source to lookup a single [Access Identity Provider][access_identity_provider_guide] by name.

## Example usage

```hcl
data "cloudflare_access_identity_provider" "main" {
  name = "Google SSO"
  account_id = "example-account-id"
}

resource "cloudflare_access_application" "main" {
  zone_id                   = "example.com"
  name                      = "name"
  domain                    = "name.example.com"
  type                      = "self_hosted"
  session_duration          = "24h"
  allowed_idps              = [data.cloudflare_access_identity_provider.main.id]
  auto_redirect_to_identity = true
}
```

## Argument Reference

The following arguments are supported:

-> **Note:** It's required that an `account_id` or `zone_id` is provided.

- `account_id` - (Optional) The account for which to look for an Access Identity Provider. Conflicts with `zone_id`.
- `zone_id` - (Optional) The Zone's ID. Conflicts with `account_id`.
- `name` - (Required) Access Identity Provider name to search for.

## Attributes Reference

- `id` - Access Identity Provider ID
- `name` - Access Identity Provider Name
- `type` - Access Identity Provider Type

[access_identity_provider_guide]: https://developers.cloudflare.com/cloudflare-one/identity/idp-integration
