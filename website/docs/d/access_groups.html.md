---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_access_groups"
sidebar_current: "docs-cloudflare-datasource-access-groups"
description: |-
  List available Cloudflare Access Groups.
---

# cloudflare_access_groups

Use this data source to look up [Access Groups][1].

## Example Usage

The example below matches all Access Groups that are part of the `zone_id` `12345`.  Access groups can also be looked up by `account_id`. The matched Access Groups are then returned as output.

```hcl
data "cloudflare_access_groups" "test" {
  zone_id = "12345"
}

output "access_groups" {
  value = data.cloudflare_access_groups.test.groups
}
```

## Argument Reference

The following arguments are supported:

-> **Note:** It's required that an `account_id` or `zone_id` is provided and in most cases using either is fine. However, if you're using a scoped access token, you must provide the argument that matches the token's scope. For example, an access token that is scoped to the "example.com" zone needs to use the `zone_id` argument.

- `zone_id` - (Optional) The ID of the DNS zone in which to search for the Access Groups.  Conflicts with `account_id`.
- `account_id` - (Optional) The ID of the account in which to search for the Access Groups.  Conflicts with `zone_id`.

## Attributes Reference

- `groups` - An list of Access Groups. Object format:

**groups**

- `id` - The Access Group ID
* `name` - Friendly name of the Access Group.
* `require` - A series of access conditions, see [Access Groups](/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).
* `exclude` - A series of access conditions, see [Access Groups](/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).
* `include` - A series of access conditions, see [Access Groups](/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).
