---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_teams_account"
sidebar_current: "docs-cloudflare-resource-teams-account"
description: Provides a Cloudflare Teams Account resource.
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
  
  antivirus {
    enabled_download_phase = true
    enabled_upload_phase = false
    fail_closed = true
  }
  
  proxy {
    tcp = true
    udp = true
  }
  
  url_browser_isolation_enabled = true
  
  logging {
    redact_pii = true
    settings_by_rule_type {
      dns {
        log_all = false
        log_blocks = true
      }
      http {
        log_all = true
        log_blocks = true
      }
      l4 {
        log_all = false
        log_blocks = true
      }
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The account to which the teams location should be added.
* `tls_decrypt_enabled` - (Optional) Indicator that decryption of TLS traffic is enabled.
* `block_page` - (Optional) Configuration for a custom block page.
* `fips` - (Optional) Configure compliance with Federal Information Processing Standards
* `antivirus` - (Optional) Configuration for antivirus traffic scanning.
* `url_browser_isolation_enabled` - (Optional) Safely browse websites in Browser Isolation through a URL.

The **block_page** block supports:

* `name` - (Optional) Name of block page configuration.
* `enabled` - (Optional) Indicator of enablement.
* `footer_text` - (Optional) Block page header text.
* `header_text` - (Optional) Block page footer text.
* `logo_path` - (Optional) URL of block page logo.
* `background_color` - (Optional) Hex code of block page background color.

The **FIPS** block supports:
* `tls` - (Optional) Only allow FIPS-compliant TLS configuration

The **antivirus** block supports:

* `enabled_download_phase` - (Optional) Scan on file download.
* `enabled_upload_phase` - (Optional) Scan on file upload.
* `fail_closed` - (Optional) Block requests for files that cannot be scanned.

* The **proxy** block supports:

* `tcp` - (Required) Whether gateway proxy is enabled on gateway devices for tcp traffic.
* `udp` - (Required) Whether gateway proxy is enabled on gateway devices for udp traffic.

The **logging** block supports:

* `redact_pii` - (Required) Redact personally identifiable information from activity logging (PII fields are: source IP,
  user email, user ID, device ID, URL, referrer, user agent).
* `settings_by_rule_type` - (Required) Represents whether all requests are logged or only the blocked requests are
  logged in DNS, HTTP and L4 filters.

## Import

Since a Teams account does not have a unique resource ID, configuration can be imported using the account ID.

```
$ terraform import cloudflare_teams_account.example cb029e245cfdd66dc8d2e570d5dd3322
```
