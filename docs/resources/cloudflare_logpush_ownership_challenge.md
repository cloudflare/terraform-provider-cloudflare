---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_logpush_ownership_challenge"
description: |-
  Provides a resource which manages Cloudflare Logpush ownership challenges to use
  in a Logpush Job.
---

# cloudflare_logpush_ownership_challenge

Provides a resource which manages Cloudflare Logpush ownership challenges to use
in a Logpush Job. On it's own, doesn't do much however this resource should
be used in conjunction to create Logpush jobs.

## Example Usage

```hcl
resource "cloudflare_logpush_ownership_challenge" "example" {
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  destination_conf = "s3://my-bucket-path?region=us-west-2"
}
```

```hcl
resource "cloudflare_logpush_ownership_challenge" "example" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  destination_conf = "s3://my-bucket-path?region=us-west-2"
}
```

## Argument Reference

The following arguments are supported:

- `destination_conf` - (Required) Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included. See [Logpush destination documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#destination).
- `account_id` - (Optional) The account ID where the logpush ownership challenge should be created. Either `account_id` or `zone_id` are required.
- `zone_id` - (Optional) The zone ID where the logpush ownership challenge should be created. Either `account_id` or `zone_id` are required.

## Attributes Reference

The following attributes are exported:

- `ownership_challenge_filename` - The filename of the ownership challenge which
  contains the contents required for Logpush Job creation.
