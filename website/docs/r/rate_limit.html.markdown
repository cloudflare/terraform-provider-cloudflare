---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_rate_limit"
sidebar_current: "docs-cloudflare-resource-rate-limit"
description: |-
  Provides a Cloudflare rate limit resource for a particular zone.
---

# cloudflare_rate_limit

Provides a Cloudflare rate limit resource for a given zone. This can be used to limit the traffic you receive zone-wide, or matching more specific types of requests/responses.

## Example Usage

```hcl
resource "cloudflare_rate_limit" "example" {
  zone = var.cloudflare_zone_id
  threshold = 2000
  period = 2
  match {
    request {
      url_pattern = "${var.cloudflare_zone}/*"
      schemes = ["HTTP", "HTTPS"]
      methods = ["GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"]
    }
    response {
      statuses = [200, 201, 202, 301, 429]
      origin_traffic = false
    }
  }
  action {
    mode = "simulate"
    timeout = 43200
    response {
      content_type = "text/plain"
      body = "custom response body"
    }
  }
  correlate {
    by = "nat"
  }
  disabled = false
  description = "example rate limit for a zone"
  bypass_url_patterns = ["${var.cloudflare_zone}/bypass1","${var.cloudflare_zone}/bypass2"]
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone ID to apply rate limiting to.
* `threshold` - (Required) The threshold that triggers the rate limit mitigations, combine with period. i.e. threshold per period (min: 2, max: 1,000,000).
* `period` - (Required) The time in seconds to count matching traffic. If the count exceeds threshold within this period the action will be performed (min: 1, max: 86,400).
* `action` - (Required) The action to be performed when the threshold of matched traffic within the period defined is exceeded.
* `match` - (Optional) Determines which traffic the rate limit counts towards the threshold. By default matches all traffic in the zone. See definition below.
* `disabled` - (Optional) Whether this ratelimit is currently disabled. Default: `false`.
* `description` - (Optional) A note that you can use to describe the reason for a rate limit. This value is sanitized and all tags are removed.
* `bypass_url_patterns` - (Optional) URLs matching the patterns specified here will be excluded from rate limiting.
* `correlate` - (Optional) Determines how rate limiting is applied. By default if not specified, rate limiting applies to the clients IP address.

The **match** block supports:

* `request` - (Optional) Matches HTTP requests (from the client to Cloudflare). See definition below.
* `response` (Optional) Matches HTTP responses before they are returned to the client from Cloudflare. If this is defined, then the entire counting of traffic occurs at this stage. This field is not required.

The **match.request** block supports:

* `methods` - (Optional) HTTP Methods, can be a subset ['POST','PUT'] or all ['\_ALL\_']. Default: ['\_ALL\_'].
* `schemes` - (Optional) HTTP Schemes, can be one ['HTTPS'], both ['HTTP','HTTPS'] or all ['\_ALL\_'].  Default: ['\_ALL\_'].
* `url_pattern` - (Optional) The URL pattern to match comprised of the host and path, i.e. example.org/path. Wildcard are expanded to match applicable traffic, query strings are not matched. Use * for all traffic to your zone. Default: '*'.

The **match.response** block supports:

* `statuses` - (Optional) HTTP Status codes, can be one [403], many [401,403] or indicate all by not providing this value.
* `origin_traffic` - (Optional) Only count traffic that has come from your origin servers. If true, cached items that Cloudflare serve will not count towards rate limiting. Default: `true`.

The **action** block supports:

* `mode` - (Required) The type of action to perform. Allowable values are 'simulate', 'ban', 'challenge' and 'js_challenge'.
* `timeout` - (Optional) The time in seconds as an integer to perform the mitigation action. This field is required if the `mode` is either `simulate` or `ban`. Must be the same or greater than the period (min: 1, max: 86400).
* `response` - (Optional) Custom content-type and body to return, this overrides the custom error for the zone. This field is not required. Omission will result in default HTML error page. Definition below.

The **action.response** block supports:

* `content_type` - (Required) The content-type of the body, must be one of: 'text/plain', 'text/xml', 'application/json'.
* `body` - (Required) The body to return, the content here should conform to the content_type.

The **correlate** block supports:

* `by` - (Optional) If set to 'nat', NAT support will be enabled for rate limiting.


## Attributes Reference

The following attributes are exported:

* `id` - The Rate limit ID.

## Import

Rate limits can be imported using a composite ID formed of zone name and rate limit ID, e.g.

```
$ terraform import cloudflare_rate_limit.default d41d8cd98f00b204e9800998ecf8427e/ch8374ftwdghsif43
```
