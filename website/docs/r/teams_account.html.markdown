---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_teams_account"
sidebar_current: "docs-cloudflare-resource-teams-account"
description: |-
Provides a Cloudflare Teams Account resource.
---

# cloudflare_teams_account

Provides a Cloudflare Teams Account resource. The Teams Account resource defines configuration for secure web gateway.

## Example Usage

```hcl
resource "cloudflare_teams_account" "main" {
  account_id  = "1d5fdc9e88c8a8c4518b068cd94331fe"
  tls_decrypt_enabled = true

  block_page {
    footer_text = "hello"
    header_text = "hello"
    logo_path = "https://google.com"
    background_color = "#000000"
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The account to which the teams location should be added.
* `tls_decrypt_enabled` - (Optional) Indicator that decryption of TLS traffic is enabled.
* `block_page` - (Optional) Configuration for a custom block page.
* `antivirus` - (Optional) Configuration for antivirus traffic scanning.

The **block_page** block supports:
* `name` - (Optional) Name of block page configuration.
* `enabled` - (Optional) Indicator of enablement.
* `footer_text` - (Optional) Block page header text.
* `header_text` - (Optional) Block page footer text.
* `logo_path` - (Optional) URL of block page logo.
* `background_color` - (Optional) Hex code of block page background color.

## Import

Since a Teams account does not have a unique resource ID, configuration can be
imported using the account ID.

```
$ terraform import cloudflare_teams_account.example cb029e245cfdd66dc8d2e570d5dd3322
```
