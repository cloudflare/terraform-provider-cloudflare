---
layout: "cloudflare"
page_title: "Provider: Cloudflare"
sidebar_current: "docs-cloudflare-index"
description: |-
  The Cloudflare provider is used to interact with the DNS resources supported by Cloudflare. The provider needs to be configured with the proper credentials before it can be used.
---

# Cloudflare Provider

The Cloudflare provider is used to interact with the
DNS resources supported by Cloudflare. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Cloudflare provider
provider "cloudflare" {
  email = "${var.cloudflare_email}"
  token = "${var.cloudflare_token}"
}

# Create a record
resource "cloudflare_record" "www" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `email` - (Required) The email associated with the account. This can also be
  specified with the `CLOUDFLARE_EMAIL` shell environment variable.
* `token` - (Required) The Cloudflare API token. This can also be specified
  with the `CLOUDFLARE_TOKEN` shell environment variable.
* `rps` - (Optional) RPS limit to apply when making calls to the API. Default: 4.
* `retries` - (Optional) Maximum number of retries to perform when an API request fails. Default: 3.
* `min_backoff` - (Optional) Minimum backoff period in seconds after failed API calls. Default: 1.
* `max_backoff` - (Optional) Maximum backoff period in seconds after failed API calls Default: 30.
* `api_client_logging` - (Optional) Whether to print logs from the API client (using the default log library logger). Default: false.

