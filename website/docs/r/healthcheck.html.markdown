---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_healthcheck"
sidebar_current: "docs-cloudflare-resource-healthcheck"
description: |-
  Provides the ability to create a Standalone Health Check without needing a Cloudflare Load Balancer.
---

# cloudflare_healthcheck

Standalone Health Checks provide a way to monitor origin servers without needing a Cloudflare Load Balancer. 

## Example Usage

The resource supports HTTP, HTTPS and TCP type health checks.

### HTTPS Health Check
```hcl
resource "cloudflare_healthcheck" "http_health_check" {
  zone_id = var.cloudflare_zone_id
  name = "http-health-check"
  description = "example http health check"
  address = "example.com"
  suspended = false
  check_regions = [
    "WEU",
    "EEU"
  ]
  notification_suspended = false
  notification_email_addresses = [
    "hostmaster@example.com"
  ]
  type = "HTTPS"
  port = "443"
  method = "GET"
  path = "/health"
  expected_body = "alive"
  expected_codes = [
    "2xx",
    "301"
  ]
  follow_redirects = true
  allow_insecure = false
  header {
    header = "Host"
    values = ["example.com"]
  }
  timeout = 10
  retries = 2
  interval = 60
  consecutive_fails = 3
  consecutive_successes = 2
}
```

### TCP Monitor
```hcl
resource "cloudflare_healthcheck" "tcp_health_check" {
  zone_id = var.cloudflare_zone_id
  name = "tcp-health-check"
  description = "example tcp health check"
  address = "example.com"
  suspended = false
  check_regions = [
    "WEU",
    "EEU"
  ]
  notification_suspended = false
  notification_email_addresses = [
    "hostmaster@example.com"
  ]
  type = "TCP"
  port = "22"
  method = "connection_established"
  timeout = 10
  retries = 2
  interval = 60
  consecutive_fails = 3
  consecutive_successes = 2
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone ID to which apply settings.
* `name` - (Required) A short name to identify the health check. Only alphanumeric characters, hyphens and underscores are allowed.
* `description` - (Optional) A human-readable description of the health check.
* `address` - (Required) The hostname or IP address of the origin server to run health checks on.
* `suspended` - (Optional) If suspended, no health checks are sent to the origin. Valid values: `true` or `false` (Default: `false`).
* `check_regions` - (Optional) A list of regions from which to run health checks. If not set Cloudflare will pick a default region. Valid values: `WNAM`, `ENAM`, `WEU`, `EEU`, `NSAM`, `SSAM`, `OC`, `ME`, `NAF`, `SAF`, `IN`, `SEAS`, `NEAS`, `ALL_REGIONS`.
* `notification_suspended` - (Optional) Whether the notifications are suspended or not. Useful for maintenance periods. Valid values: `true` or `false` (Default: `false`).
* `notification_email_addresses` - (Optional) A list of email addresses we want to send the notifications to.
* `type` - (Required) The protocol to use for the health check. Valid values: `HTTP`, `HTTPS`, `TCP`.
* `port` - (Optional) Port number to connect to for the health check.  Valid values are in the rage `0-65535` (Default: `80`).
* `timeout` - (Optional) The timeout (in seconds) before marking the health check as failed. (Default: `5`)
* `retries` - (Optional) The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately. (Default: `2`)
* `interval` - (Optional) The interval between each health check. Shorter intervals may give quicker notifications if the origin status changes, but will increase load on the origin as we check from multiple locations. (Default: `60`)
* `consecutive_fails` - (Optional) The number of consecutive fails required from a health check before changing the health to unhealthy. (Default: `1`)
* `consecutive_successes` - (Optional) The number of consecutive successes required from a health check before changing the health to healthy. (Default: `1`)

### HTTP/HTTPS specific arguments
* `method` - (Optional) The HTTP method to use for the health check. Valid values: `GET` or `HEAD` (Default: `GET`).
* `path` - (Optional) The endpoint path to health check against. (Default: `/`)
* `expected_body` - (Optional) A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy.
* `expected_codes` - (Optional) The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all codes starting with 2) of the health check. (Default: `["200"]`)
* `follow_redirects` - (Optional) Follow redirects if the origin returns a 3xx status code. Valid values: `true` or `false` (Default: `false`).
* `allow_insecure` - (Optional) Do not validate the certificate when the health check uses HTTPS. Valid values: `true` or `false` (Default: `false`).
* `header` - (Optional) The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden.

**header** requires the following:

* `header` - (Required) The header name.
* `values` - (Required) A list of string values for the header.

### TCP specific arguments
* `method` - (Optional) The TCP connection method to use for the health check. Valid values: `connection_established` (Default: `connection_established`).