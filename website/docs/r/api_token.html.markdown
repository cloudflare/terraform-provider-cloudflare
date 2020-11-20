---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_api_token"
sidebar_current: "docs-cloudflare-resource-api-token"
description: |-
  Provides a resource which manages Cloudflare API tokens.
---

# cloudflare_api_token

Provides a resource which manages Cloudflare API tokens.

Read more about permission groups and their applicable scopes in
[the official documentation][1].

## Example Usage

### User Permissions

```hcl
data "cloudflare_api_token_permission_groups" "all" {}

// Token allowed to create new tokens.
// Can only be used from specific ip range.
resource "cloudflare_api_token" "api_token_create" {
  name = "api_token_create"

  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.permissions["API Tokens Write"],
    ]
    resources = {
      "com.cloudflare.api.user.${var.user_id}" = "*"
    }
  }

  request_ip_in = [
    "10.0.0.0/8",
    "2400:cb00::/32",
  ]
  request_ip_not_in = [
    "10.0.0.10/16",
    "2400:cb00::/64",
  ]
}
```

### Account permissions

```hcl
data "cloudflare_api_token_permission_groups" "all" {}

// Token allowed to read audit logs from all accounts.
resource "cloudflare_api_token" "logs_account_all" {
  name = "logs_account_all"

  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.permissions["Access: Audit Logs Read"],
    ]
    resources = {
      "com.cloudflare.api.account.*" = "*"
    }
  }
}

// Token allowed to read audit logs from specific account.
resource "cloudflare_api_token" "logs_account" {
  name = "logs_account"

  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.permissions["Access: Audit Logs Read"],
    ]
    resources = {
      "com.cloudflare.api.account.${var.account_id}" = "*"
    }
  }
}
```
### Zone Permissions

```hcl
data "cloudflare_api_token_permission_groups" "all" {}

// Token allowed to edit DNS entries and TLS certs for specific zone.
resource "cloudflare_api_token" "dns_tls_edit" {
  name = "dns_tls_edit"

  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.permissions["DNS Write"],
      data.cloudflare_api_token_permission_groups.all.permissions["SSL and Certificates Write"],
    ]
    resources = {
      "com.cloudflare.api.account.zone.${var.zone_id}" = "*"
    }
  }
}

// Token allowed to edit DNS entries for all zones except one.
resource "cloudflare_api_token" "dns_tls_edit_all_except_one" {
  name = "dns_tls_edit_all_except_one"

  // include all zones
  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.permissions["DNS Write"],
    ]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }

  // exclude (deny) specific zone
  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.permissions["DNS Write"],
    ]
    resources = {
      "com.cloudflare.api.account.zone.${var.zone_id}" = "*"
    }
    effect = "deny"
  }
}


// Token allowed to edit DNS entries for all zones from specific account.
resource "cloudflare_api_token" "dns_edit_all_account" {
  name = "dns_edit_all_account"

  // include all zones from specific account
  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.permissions["DNS Write"],
    ]
    resources = {
      "com.cloudflare.api.account.${var.account_id}" = jsonencode({
        "com.cloudflare.api.account.zone.*" = "*"
      })
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the APIToken.
* `policy` - (Required) Permissions policy. Multiple policy blocks can be defined.
See the definition below.
* `request_ip_in` - (Optional) List of IPv4/IPv6 CIDR addresses where
the Token can be used from.
* `request_ip_not_in` - (Optional) List of IPv4/IPv6 CIDR addresses where
the Token cannot be used from.

The **policy** block supports:

* `permission_groups` - (Required) List of permissions groups
ids ([see official docs][1]).
* `resources` - (Required) Map describes what operations against which resources
are allowed or denied.
* `effect` - (Optional) Policy effect. Valid values are `allow` or `deny`. `allow` 
   is set as default.

[1]: https://developers.cloudflare.com/api/tokens/create/permissions
