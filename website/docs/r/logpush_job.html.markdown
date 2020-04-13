---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_logpush_job"
sidebar_current: "docs-cloudflare-logpush-job"
description: |-
  Provides a resource which manages Cloudflare logpush jobs.
---

# cloudflare_logpush_job

Provides a resource which manages Cloudflare logpush jobs.

## Example Usage

```hcl
resource "cloudflare_logpush_job" "example_job" {
  enabled = true
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  name = "My-logpush-job"
  logpull_options = "fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"
  destination_conf = "s3://my-bucket-path?region=us-west-2"
  ownership_challenge = "00000000000000000"
  dataset = "http_requests"
}
```

## Argument Reference

The following arguments are supported:


* `name` - (Required) The name of the logpush job to create. Must match the regular expression `^[a-zA-Z0-9\-\.]*$`.
* `zone_id` - (Required) The zone ID where the logpush job should be created.
* `destination_conf` - (Required) Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included. See [Logpush destination documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#destination).
* `ownership_challenge` - (Required) Ownership challenge token to prove destination ownership. See [Developer documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#usage).
* `dataset` - (Required) Which type of dataset resource to use. Available values are `"http_requests"` or `"spectrum_events"`.
* `logpull_options` - (Optional) Configuration string for the Logshare API. It specifies things like requested fields and timestamp formats. See [Logpull options documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#options).
* `enable` - (Optional) Whether to enable to job to create or not.
