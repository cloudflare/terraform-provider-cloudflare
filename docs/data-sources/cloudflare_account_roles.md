---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_account_roles"
description: Get information on a Cloudflare Account Roles.
---

# cloudflare_account_roles

Use this data source to lookup [Account Roles][1].

## Example usage

```hcl
data "cloudflare_account_roles" "account_roles" {
    account_id = var.cloudflare_account_id
}

locals {
  roles_by_name = {
    for role in data.cloudflare_account_roles.account_roles.roles :
      role.name => role
  }
}

resource "cloudflare_account_member" "member" {
  ...
  role_ids = [
    local.roles_by_name["Administrator"].id
  ]
}
```

## Argument Reference

- `account_id` - (Required) The account for which to list the roles.

## Attributes Reference

- `roles` - A list of roles object. See below for nested attributes.

**roles**

- `id` - Role identifier tag
- `name` - Role Name
- `description` - Description of role's permissions

[1]: https://api.cloudflare.com/#account-roles-properties
