---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_load_balancer_monitor"
sidebar_current: "docs-cloudflare-resource-load-balancer-monitor"
description: |-
  Provides a Cloudflare Load Balancer Monitor resource.
---

# cloudflare_load_balancer_monitor

If you're using Cloudflare's Load Balancing to load-balance across multiple origin servers or data centers, you configure one of these Monitors to actively check the availability of those servers over HTTP(S).

## Example Usage

```hcl
resource "cloudflare_load_balancer_monitor" "test" {
  expected_body = "alive"
  expected_codes = "2xx"
  method = "GET"
  timeout = 7
  path = "/health"
  interval = 60
  retries = 5
  description = "example load balancer"
  header {
    header = "Host"
    values = ["example.com"]
  }
  allow_insecure = false
  follow_redirects = true
}
```

## Argument Reference

The following arguments are supported:

* `expected_body` - (Required) A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy.
* `expected_codes` - (Required) The expected HTTP response code or code range of the health check. Eg `2xx`
* `method` - (Optional) The HTTP method to use for the health check. Default: "GET".
* `timeout` - (Optional) The timeout (in seconds) before marking the health check as failed. Default: 5.
* `path` - (Optional) The endpoint path to health check against. Default: "/".
* `interval` - (Optional) The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations. Default: 60.
* `retries` - (Optional) The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately. Default: 2.
* `header` - (Optional) The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden. Fields documented below.
* `type` - (Optional) The protocol to use for the healthcheck. Currently supported protocols are 'HTTP' and 'HTTPS'. Default: "http".
* `description` - (Optional) Free text description.
* `allow_insecure` - (Optional) Do not validate the certificate when monitor use HTTPS.
* `follow_redirects` - (Optional) Follow redirects if returned by the origin.

**header** requires the following:

* `header` - (Required) The header name.
* `values` - (Required) A list of string values for the header.

## Attributes Reference

The following attributes are exported:

* `id` - Load balancer monitor ID.
* `created_on` - The RFC3339 timestamp of when the load balancer monitor was created.
* `modified_on` - The RFC3339 timestamp of when the load balancer monitor was last modified.
