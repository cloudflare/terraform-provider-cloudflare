---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_access_service_token"
sidebar_current: "docs-cloudflare-resource-access-service-token"
description: |-
  Provides a Cloudflare Access Service Token resource.
---

# cloudflare_access_service_token

Access Service Tokens are used for service-to-service communication
when an application is behind Cloudflare Access.

## Example Usage

```hcl
resource "cloudflare_access_service_token" "my_app" {
  account_id = "d41d8cd98f00b204e9800998ecf8427e"
  name       = "CI/CD app"
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The ID of the account where the Access
Service is being created.
* `name` - (Required) Friendly name of the token's intent.

## Attributes Reference

The following attributes are exported:

* `client_id` - UUID client ID associated with the Service Token.
* `client_secret` - A secret for interacting with Access protocols.

## Import

~> **Important:** If you are importing an Access Service Token you will
not have the `client_secret` available in the state for use. The
`client_secret` is only available once, at creation. In most cases, it
is better to just create a new resource should you need to reference it
in other resources.

Access Service Tokens can be imported using a composite ID formed of
account ID and Service Token ID.

```
$ terraform import cloudflare_access_service_token.my_app cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e
```

where

* `cb029e245cfdd66dc8d2e570d5dd3322` - Account ID
* `d41d8cd98f00b204e9800998ecf8427e` - Access Service Token ID
