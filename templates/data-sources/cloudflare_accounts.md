---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_accounts"
description: Get information on Cloudflare accounts.
---

# cloudflare_accounts

Use this data source to lookup [Accounts][1].

## Example usage

```hcl
data "cloudflare_accounts" "accounts" {
    name = var.cloudflare_account_name
}

locals {
  account_id = accounts.accounts[0].id
}
```

## Argument Reference

- `name` - (Optional) Account name to filter accounts

## Attributes Reference

- `accounts` - A list of accounts object. See below for nested attributes.

**accounts**

- `id` - Account ID
- `name` - Account Name
- `type` - Account subscription type
- `enforce_twofactor` - Enforcement of 2 factors authentication

[1]: https://api.cloudflare.com/#accounts
