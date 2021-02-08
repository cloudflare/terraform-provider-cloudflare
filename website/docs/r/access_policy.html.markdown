---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_access_policy"
sidebar_current: "docs-cloudflare-resource-access-policy"
description: |-
  Provides a Cloudflare Access Policy resource.
---

# cloudflare_access_policy

Provides a Cloudflare Access Policy resource. Access Policies are used
in conjunction with Access Applications to restrict access to a
particular resource.

## Example Usage

```hcl
# Allowing access to `test@example.com` email address only
resource "cloudflare_access_policy" "test_policy" {
  application_id = "cb029e245cfdd66dc8d2e570d5dd3322"
  zone_id        = "d41d8cd98f00b204e9800998ecf8427e"
  name           = "staging policy"
  precedence     = "1"
  decision       = "allow"

  include {
    email = ["test@example.com"]
  }
}

# Allowing `test@example.com` to access but only when coming from a
# specific IP.
resource "cloudflare_access_policy" "test_policy" {
  application_id = "cb029e245cfdd66dc8d2e570d5dd3322"
  zone_id        = "d41d8cd98f00b204e9800998ecf8427e"
  name           = "staging policy"
  precedence     = "1"
  decision       = "allow"

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

-> **Note:** It's required that an `account_id` or `zone_id` is provided and in most cases using either is fine. However, if you're using a scoped access token, you must provide the argument that matches the token's scope. For example, an access token that is scoped to the "example.com" zone needs to use the `zone_id` argument.

* `application_id` - (Required) The ID of the application the policy is associated with.
* `account_id` - (Optional) The account to which the access rule should be added. Conflicts with `zone_id`.
* `zone_id` - (Optional) The DNS zone to which the access rule should be added. Conflicts with `account_id`.
* `decision` - (Required) Defines the action Access will take if the policy matches the user.
  Allowed values: `allow`, `deny`, `non_identity`, `bypass`
* `name` - (Required) Friendly name of the Access Application.
* `precedence` - (Required) The unique precedence for policies on a single application. Integer.
* `require` - (Optional) A series of access conditions, see [Access Groups](/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).
* `exclude` - (Optional) A series of access conditions, see [Access Groups](/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).
* `include` - (Required) A series of access conditions, see [Access Groups](/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).


## Import

Access Policies can be imported using a composite ID formed of zone
ID, application ID and policy ID.

```
$ terraform import cloudflare_access_policy.staging cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e/67ea780ce4982c1cfbe6b7293afc765d
```

where

* `cb029e245cfdd66dc8d2e570d5dd3322` - Zone ID
* `d41d8cd98f00b204e9800998ecf8427e` - Access Application ID
* `67ea780ce4982c1cfbe6b7293afc765d` - Access Policy ID
