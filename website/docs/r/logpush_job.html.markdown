---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_logpush_job"
sidebar_current: "docs-cloudflare-logpush-job"
description: |-
  Provides a resource which manages Cloudflare logpush jobs.
---

# cloudflare_logpush_job

Provides a resource which manages Cloudflare Logpush jobs. On it's own, this
resource cannot be automatically created. In order to have this automated, you
must have:

- `cloudflare_logpush_ownership_challenge`: Configured to generate the challenge
  to confirm ownership of the destination.
- Either manual inspection or another Terraform Provider to get the contents of
  the `ownership_challenge_filename` value from the
  `cloudflare_logpush_ownership_challenge` resource.
- `cloudflare_logpush_job`: Create and manage the Logpush Job itself.

## Example Usage (with AWS provider)

Please see
[`cloudflare_logpush_ownership_challenge`](/docs/providers/cloudflare/r/logpush_ownership_challenge.html)
for how to use that resource and the third party provider documentation if you
choose to automate the intermediate step of fetching the ownership challenge contents.

```hcl
resource "cloudflare_logpush_ownership_challenge" "ownership_challenge" {
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  destination_conf = "s3://my-bucket-path?region=us-west-2"
}

data "aws_s3_bucket_object" "challenge_file" {
  bucket = "my-bucket-path"
  key    = cloudflare_logpush_ownership_challenge.ownership_challenge.ownership_challenge_filename
}

resource "cloudflare_logpush_job" "example_job" {
  enabled = true
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  name = "My-logpush-job"
  logpull_options = "fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"
  destination_conf = "s3://my-bucket-path?region=us-west-2"
  ownership_challenge = aws_s3_bucket_object.challenge_file.body
  dataset = "http_requests"
}
```

## Example Usage (manual inspection of S3 bucket)

- Create `cloudflare_logpush_ownership_challenge` resource

```hcl
resource "cloudflare_logpush_ownership_challenge" "ownership_challenge" {
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  destination_conf = "s3://my-bucket-path?region=us-west-2"
}
```

- Check S3 bucket for your ownership challenge filename and grab the contents.
- Create the `cloudflare_logpush_job` substituting in your manual `ownership_challenge`.

```hcl
resource "cloudflare_logpush_job" "example_job" {
  enabled = true
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  name = "My-logpush-job"
  logpull_options = "fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"
  destination_conf = "s3://my-bucket-path?region=us-west-2"
  ownership_challenge = "0000000000000"
  dataset = "http_requests"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the logpush job to create. Must match the regular expression `^[a-zA-Z0-9\-\.]*$`.
* `zone_id` - (Required) The zone ID where the logpush job should be created.
* `destination_conf` - (Required) Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included. See [Logpush destination documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#destination).
* `ownership_challenge` - (Required) Ownership challenge token to prove destination ownership. See [Developer documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#usage).
* `dataset` - (Required) Which type of dataset resource to use. Available values are `"firewall_events"`, `"http_requests"`, and `"spectrum_events"`.
* `logpull_options` - (Optional) Configuration string for the Logshare API. It specifies things like requested fields and timestamp formats. See [Logpull options documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#options).
* `enabled` - (Optional) Whether to enable to job to create or not.
