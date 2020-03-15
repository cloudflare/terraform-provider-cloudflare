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

* `application_id` - (Required) The ID of the application the policy is
  associated with.
* `zone_id` - (Required) The DNS zone to which the access rule should be
  added.
* `decision` - (Required) Defines the action Access will take if the policy matches the user.
  Allowed values: `allow`, `deny`, `non_identity`, `bypass`
* `name` - (Required) Friendly name of the Access Application.
* `precedence` - (Optional) The unique precedence for policies on a single application. Integer.
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
* `service_token` - (Optional) A list of service token ids. Example:
  `service_token = cloudflare_access_service_token.demo.id`
* `any_valid_service_token` - (Optional) Boolean indicating if allow
  all tokens to be granted. Example: `any_valid_service_token = true`
* `everyone` - (Optional) Boolean indicating permitting access for all
  requests. Example: `everyone = true`


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
