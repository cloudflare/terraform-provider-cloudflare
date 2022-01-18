---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_fallback_domain"
sidebar_current: "docs-cloudflare-resource-fallback-domain"
description: |-
  Provides a Cloudflare Fallback Domain resource.
---

# cloudflare_fallback_domain

Provides a Cloudflare Fallback Domain resource. Fallback domains are used to ignore DNS requests to a given list of domains. These DNS requests will be passed back to other DNS servers configured on existing network interfaces on the device.

## Example Usage

```hcl
# Adding example.com to the default fallback domain list.
# Upon resource deletion, restore the default entries.
resource "cloudflare_fallback_domain" "with_defaults" {
  account_id                        = "1d5fdc9e88c8a8c4518b068cd94331fe"
  include_default_domains           = true
  restore_default_domains_on_delete = true
  domains {
    suffix      = "example.com"
    description = "Example domain"
    dns_server  = ["1.1.1.1", "2.2.2.2"]
  }
}
```

```hcl
# Specify all domains explicitly, including default domains.
# Upon deletion, leave domains empty.
resource "cloudflare_fallback_domain" "explicit" {
  account_id                        = "1d5fdc9e88c8a8c4518b068cd94331fe"
  include_default_domains           = false
  restore_default_domains_on_delete = false
  dynamic "domains" {
    for_each = toset(["intranet", "internal", "private", "localdomain", "domain", "lan", "home", "host", "corp", "local", "localhost", "home.arpa", "invalid", "test"])
    content {
      suffix = domains.value
    }
  }
  domains {
    suffix      = "example.com"
    description = "Example domain"
    dns_server  = ["1.1.1.1", "2.2.2.2"]
  }
}
```


## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The account to which the device posture rule should be added.
- `domains` - (Required) The value of the domain attributes (refer to the [nested schema](#nestedblock--domains)).
- `include_default_domains` - (Optional) Indicator to add specified domains to default list of domains (shown in `explicit` example).  Defaults to `false`.
- `restore_default_domains_on_delete` - (Optional) Indicator to reset domains to default list upon delete.  Otherwise, list will be empty.  Defaults to `true`.

<a id="nestedblock--domains"></a>
**Nested schema for `domains`**

- `suffix` - (Required) The domain to ignore DNS requests.
- `description` - (Optional) The description of the domain.
- `dns_server` - (Optional) The DNS servers to receive the redirected request.

## Import

Fallback Domains can be imported using the account identifer.

```
$ terraform import cloudflare_fallback_domain.example 1d5fdc9e88c8a8c4518b068cd94331fe
```
