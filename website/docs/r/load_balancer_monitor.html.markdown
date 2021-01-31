---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_load_balancer_monitor"
sidebar_current: "docs-cloudflare-resource-load-balancer-monitor"
description: |-
  Provides a Cloudflare Load Balancer Monitor resource.
---

# cloudflare_load_balancer_monitor

If you're using Cloudflare's Load Balancing to load-balance across multiple origin servers or data centers, you configure one of these Monitors to actively check the availability of those servers over HTTP(S) or TCP.

-> **Note:** When creating a monitor, you have to pass `account_id` to the provider configuration in order to create account level resources. Otherwise, by default, it will be a user level resource.

## Example Usage

### HTTP Monitor
```hcl
resource "cloudflare_load_balancer_monitor" "http_monitor" {
  type = "http"
  expected_body = "alive"
  expected_codes = "2xx"
  method = "GET"
  timeout = 7
  path = "/health"
  interval = 60
  retries = 5
  description = "example http load balancer"
  header {
    header = "Host"
    values = ["example.com"]
  }
  allow_insecure = false
  follow_redirects = true
  probe_zone = "example.com"
}
```

### TCP Monitor
```hcl
resource "cloudflare_load_balancer_monitor" "tcp_monitor" {
  type = "tcp"
  method = "connection_established"
  timeout = 7
  port = 8080
  interval = 60
  retries = 5
  description = "example tcp load balancer"
}
```

## Argument Reference

The following arguments are supported:

* `expected_body` - (Optional) A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy. Only valid if `type` is "http" or "https". Default: "".
* `expected_codes` - (Optional) The expected HTTP response code or code range of the health check. Eg `2xx`. Only valid and required if `type` is "http" or "https".
* `method` - (Optional) The method to use for the health check. Valid values are any valid HTTP verb if `type` is "http" or "https", or `connection_established` if `type` is "tcp". Default: "GET" if `type` is "http" or "https", or "connection_established" if `type` is "tcp" .
* `timeout` - (Optional) The timeout (in seconds) before marking the health check as failed. Default: 5.
* `path` - (Optional) The endpoint path to health check against. Default: "/". Only valid if `type` is "http" or "https".
* `interval` - (Optional) The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations. Default: 60.
* `retries` - (Optional) The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately. Default: 2.
* `header` - (Optional) The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden. Fields documented below. Only valid if `type` is "http" or "https".
* `type` - (Optional) The protocol to use for the healthcheck. Currently supported protocols are 'HTTP', 'HTTPS' and 'TCP'. Default: "http".
* `port` - The port number to use for the healthcheck, required when creating a TCP monitor. Valid values are in the range `0-65535`.
* `description` - (Optional) Free text description.
* `allow_insecure` - (Optional) Do not validate the certificate when monitor use HTTPS. Only valid if `type` is "http" or "https".
* `follow_redirects` - (Optional) Follow redirects if returned by the origin. Only valid if `type` is "http" or "https".
* `probe_zone` - (Optional) Assign this monitor to emulate the specified zone while probing. Only valid if `type` is "http" or "https".

**header** requires the following:

* `header` - (Required) The header name.
* `values` - (Required) A list of string values for the header.

## Attributes Reference

The following attributes are exported:

* `id` - Load balancer monitor ID.
* `created_on` - The RFC3339 timestamp of when the load balancer monitor was created.
* `modified_on` - The RFC3339 timestamp of when the load balancer monitor was last modified.
