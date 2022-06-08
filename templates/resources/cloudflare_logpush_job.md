---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_logpush_job"
description: Provides a resource which manages Cloudflare logpush jobs.
---

# cloudflare_logpush_job

Provides a resource which manages Cloudflare Logpush jobs. For Logpush jobs pushing to Amazon S3, Google Cloud Storage,
Microsoft Azure or Sumo Logic, this resource cannot be automatically created. In order to have this automated, you must
have:

- `cloudflare_logpush_ownership_challenge`: Configured to generate the challenge
  to confirm ownership of the destination.
- Either manual inspection or another Terraform Provider to get the contents of
  the `ownership_challenge_filename` value from the
  `cloudflare_logpush_ownership_challenge` resource.
- `cloudflare_logpush_job`: Create and manage the Logpush Job itself.

## Example Usage (Cloudflare R2)

When using Cloudflare R2, no ownership challenge is required.

```hcl
data "cloudflare_api_token_permission_groups" "all" {}

resource "cloudflare_api_token" "logpush_r2_token" {
  name = "logpush_r2_token"

  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.permissions["Workers R2 Storage Write"],
    ]
    resources = {
      "com.cloudflare.api.account.*" = "*"
    }
  }
}

resource "cloudflare_logpush_job" "http_requests" {
  enabled          = true
  zone_id          = var.zone_id
  name             = "http_requests"
  logpull_options  = "fields=ClientIP,ClientRequestHost,ClientRequestMethod,ClientRequestURI,EdgeEndTimestamp,EdgeResponseBytes,EdgeResponseStatus,EdgeStartTimestamp,RayID&timestamps=rfc3339"
  destination_conf = "r2://cloudflare-logs/http_requests/date={DATE}?account-id=${var.account_id}&access-key-id=${cloudflare_api_token.logpush_r2_token.id}&secret-access-key=${sha256(cloudflare_api_token.logpush_r2_token.value)}"
  dataset          = "http_requests"
}
```

## Example Usage (with AWS provider)

Please see
[`cloudflare_logpush_ownership_challenge`](/docs/providers/cloudflare/r/logpush_ownership_challenge.html)
for how to use that resource and the third party provider documentation if you
choose to automate the intermediate step of fetching the ownership challenge contents.

~> **Important:** If you're using this approach, the `destination_conf` values must
match identically in all resources. Otherwise the challenge validation will fail.

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
  ownership_challenge = data.aws_s3_bucket_object.challenge_file.body
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
  frequency = "high"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) The name of the logpush job to create. Must match the regular expression `^[a-zA-Z0-9\-\.]*$`.
- `destination_conf` - (Required) Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included. See [Logpush destination documentation](https://developers.cloudflare.com/logs/reference/logpush-api-configuration#destination).
- `dataset` - (Required) Which type of dataset resource to use. Available values are
  - [account-scoped](https://developers.cloudflare.com/logs/reference/log-fields/account): `"audit_logs"`, `"gateway_dns"`, `"gateway_http"`, `"gateway_network"`, `"network_analytics_logs"`
  - [zone-scoped](https://developers.cloudflare.com/logs/reference/log-fields/zone): `"firewall_events"`, `"http_requests"`, `"spectrum_events"`, `"nel_reports", "dns_logs"`
- `account_id` - (Optional) The account ID where the logpush job should be created. Either `account_id` or `zone_id` are required.
- `zone_id` - (Optional) The zone ID where the logpush job should be created. Either `account_id` or `zone_id` are required.
- `logpull_options` - (Optional) Configuration string for the Logshare API. It specifies things like requested fields and timestamp formats. See [Logpull options documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#options).
- `ownership_challenge` - (Optional) Ownership challenge token to prove destination ownership, required when destination is Amazon S3, Google Cloud Storage,
  Microsoft Azure or Sumo Logic. See [Developer documentation](https://developers.cloudflare.com/logs/logpush/logpush-configuration-api/understanding-logpush-api/#usage).
* `enabled` - (Optional) Whether to enable the job.
* `frequency` - (Optional) `"high"` or `"low"`. A higher frequency will result in logs being pushed on faster with smaller files. `"low"` frequency will push logs less often with larger files.
* `filter` - (Optional) Use filters to select the events to include and/or remove from your logs. For more information, refer to [Filters](https://developers.cloudflare.com/logs/reference/logpush-api-configuration/filters/).

## Import

Logpush jobs can be imported using a composite ID formed of:

- `identifierType` - Either `account` or `zone`.
- `identifierID` - The ID of the account or zone.
- `jobID` - The Logpush Job ID to import.

Import an account-scoped job using `account/:accountID/:jobID`

```
$ terraform import cloudflare_logpush_job.example account/1d5fdc9e88c8a8c4518b068cd94331fe/54321
```

Import a zone-scoped job using `zone/:zoneID/:jobID`

```
$ terraform import cloudflare_logpush_job.example zone/d41d8cd98f00b204e9800998ecf8427e/54321
```
