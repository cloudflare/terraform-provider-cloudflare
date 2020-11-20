---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_api_token_permission_groups"
sidebar_current: "docs-cloudflare-api-token-permissions-groups"
description: |-
  List available API Token Permission Group IDs.
---

# cloudflare_api_token_permission_groups

Use this data source to look up [API Token Permission Groups](https://developers.cloudflare.com/api/tokens/create/permissions). Commonly used as references within [`cloudflare_api_token`](/docs/providers/cloudflare/r/api_token.html) resources.

## Example Usage

```hcl
data "cloudflare_api_token_permission_groups" "test" {}

output "dns_read_permission_id" {
  value = data.cloudflare_api_token_permission_groups.test.permissions["DNS Read"] // 82e64a83756745bbbb1c9c2701bf816b
}
```

## Attributes Reference

- `permissions` - A map of permission groups where keys are human-readable permission names
and values are permission IDs.
