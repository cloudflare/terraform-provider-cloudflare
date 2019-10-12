---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_access_group"
sidebar_current: "docs-cloudflare-resource-access-group"
description: |-
  Provides a Cloudflare Access Group resource.
---

# cloudflare_access_group

Provides a Cloudflare Access Group resource. Access Groups are used
in conjunction with Access Plolicies to restrict access to a
particular resource.

## Example Usage

```hcl
# Allowing access to `test@example.com` email address only
resource "cloudflare_access_group" "test_group" {
  account_id     = "975ecf5a45e3bcb680dba0722a420ad9"
  name           = "staging group"

  include {
    email = ["test@example.com"]
  }
}

# Allowing `test@example.com` to access but only when coming from a
# specific IP.
resource "cloudflare_access_group" "test_group" {
  account_id     = "975ecf5a45e3bcb680dba0722a420ad9"
  name           = "staging group"

  include {
    email = ["test@example.com"]
  }

  require = {
    ip = [var.office_ip]
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The ID of the account the group is
  associated with.
* `name` - (Required) Friendly name of the Access Group.
* `require` - (Optional) A series of access conditions, see below for
  full list.
* `exclude` - (Optional) A series of access conditions, see below for
  full list.
* `include` - (Required) A series of access conditions, see below for
  full list.

## Conditions

`require`, `exclude` and `include` arguments share the available
conditions which can be applied. The conditions are:

* `ip` - (Optional) A list of IP addresses or ranges. Example:
  `ip = ["1.2.3.4", "10.0.0.0/2"]`
* `email` - (Optional) A list of email addresses. Example:
  `email = ["test@example.com"]`
* `email_domain` - (Optional) A list of email domains. Example:
  `email_domain = ["example.com"]`
* `everyone` - (Optional) Boolean indicating permitting access for all
  requests. Example: `everyone = true`


## Import

Access Groups can be imported using a composite ID formed of account
ID and group ID.

```
$ terraform import cloudflare_access_group.staging 975ecf5a45e3bcb680dba0722a420ad9/67ea780ce4982c1cfbe6b7293afc765d
```

where

* `975ecf5a45e3bcb680dba0722a420ad9` - Account ID
* `67ea780ce4982c1cfbe6b7293afc765d` - Access Group ID
