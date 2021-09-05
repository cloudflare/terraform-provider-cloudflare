---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_teams_location"
sidebar_current: "docs-cloudflare-resource-teams-location"
description: |-
Provides a Cloudflare Teams Location resource.
---

# cloudflare_teams_location

Provides a Cloudflare Teams Location resource. Teams Locations are referenced
when creating secure web gateway policies.

## Example Usage

```hcl
resource "cloudflare_teams_location" "corporate_office" {
  name = "office"
  account_id  = "1d5fdc9e88c8a8c4518b068cd94331fe"
  client_default = true
  networks {
    network = "203.0.113.1/32"
  }
  networks {
    network = "203.0.113.2/32"
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The account to which the teams location should be added.
* `name` - (Required) Name of the teams location.
* `networks` - (Optional) The networks CIDRs that comprise the location.
* `client_default` - (Optional) Indicator that this is the default location.

## Attributes Reference

The following additional attributes are exported:

* `id` - ID of the teams location.
* `ip` - Client IP address
* `doh_subdomain` - The FQDN that DoH clients should be pointed at.
* `anonymized_logs_enabled` - Indicator that anonymized logs are enabled.
* `ipv4_destination` - IP to direct all IPv4 DNS queries too.

## Import

Teams locations can be imported using a composite ID formed of account
ID and teams location ID.

```
$ terraform import cloudflare_teams_location.corporate_office cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e
```
