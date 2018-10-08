---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_account_member"
sidebar_current: "docs-cloudflare-resource-account-member"
description: |-
  Provides a resource which manages Cloudflare account members.
---

# cloudflare_account_member

Provides a resource which manages Cloudflare account members.

## Example Usage

```hcl
resource "cloudflare_account_member" "example_user" {
  email_address = "user@example.com"
  role_ids = [
    "68b329da9893e34099c7d8ad5cb9c940",
    "d784fa8b6d98d27699781bd9a7cf19f0"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `email_address` - (Required) The email address of the user who you wish to manage. Note: Following creation, this field becomes read only via the API and cannot be updated.
* `role_ids` - (Required) Array of account role IDs that you want to assign to a member.
