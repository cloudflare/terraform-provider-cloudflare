---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_teams_proxy_endpoint"
sidebar_current: "docs-cloudflare-resource-teams-proxy-endpoint"
description: |-
Provides a Cloudflare Teams Proxy Endpoint resource.
---

# cloudflare_teams_proxy_endpoint

Provides a Cloudflare Teams Proxy Endpoint resource. Teams Proxy Endpoints are used for pointing proxy clients at
Cloudflare Secure Gateway.

## Example Usage

```hcl
resource "cloudflare_teams_proxy_endpoint" "corporate_office" {
  name = "office"
  account_id  = "1d5fdc9e88c8a8c4518b068cd94331fe"
  ips  = ["192.0.2.0/24"]
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The account to which the teams proxy endpoint should be added.
* `name` - (Required) Name of the teams proxy endpoint.
* `ips` - (Required) The networks CIDRs that will be allowed to initiate proxy connections.

## Attributes Reference

The following additional attributes are exported:

* `id` - ID of the teams proxy endpoint.
* `subdomain` - The FQDN that proxy clients should be pointed at.

## Import

Teams Proxy Endpoints can be imported using a composite ID formed of account
ID and teams proxy_endpoint ID.

```
$ terraform import cloudflare_teams_proxy_endpoint.corporate_office cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e
```
