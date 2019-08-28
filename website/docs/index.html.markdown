---
layout: "cloudflare"
page_title: "Provider: Cloudflare"
sidebar_current: "docs-cloudflare-index"
description: |-
  The Cloudflare provider is used to interact with resources supported by Cloudflare. The provider needs to be configured with the proper credentials before it can be used.
---

# Cloudflare Provider

The Cloudflare provider is used to interact with resources supported by Cloudflare. The provider needs to be configured
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

# Create a page rule
resource "cloudflare_page_rule" "www" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `email` - (Optional) The email associated with the account. This can also be
  specified with the `CLOUDFLARE_EMAIL` shell environment variable.
* `token` - (Optional) The Cloudflare API key. This can also be specified
  with the `CLOUDFLARE_TOKEN` shell environment variable.
* `api_token` - (Optional) The Cloudflare API Token. This can also be specified with
  the `CLOUDFLARE_API_TOKEN` shell environment variable. This is an alternative to
  `email`+`token`(key). If both are specified, `api_token` will be used over `email`+`token`(key) fields.
* `rps` - (Optional) RPS limit to apply when making calls to the API. Default: 4.
  This can also be specified with the `CLOUDFLARE_RPS` shell environment variable.
* `retries` - (Optional) Maximum number of retries to perform when an API request fails. Default: 3.
  This can also be specified with the `CLOUDFLARE_RETRIES` shell environment variable.
* `min_backoff` - (Optional) Minimum backoff period in seconds after failed API calls. Default: 1.
  This can also be specified with the `CLOUDFLARE_MIN_BACKOFF` shell environment variable.
* `max_backoff` - (Optional) Maximum backoff period in seconds after failed API calls Default: 30.
  This can also be specified with the `CLOUDFLARE_MAX_BACKOFF` shell environment variable.
* `api_client_logging` - (Optional) Whether to print logs from the API client (using the default log library logger). Default: false.
  This can also be specified with the `CLOUDFLARE_API_CLIENT_LOGGING` shell environment variable.
* `account_id` - (Optional) Configure API client with this account ID, so calls use the account API rather than the (default) user API.
  This is required for other users in your account to have access to the resources you manage.
  This can also be specified with the `CLOUDFLARE_ACCOUNT_ID` shell environment variable.
* `use_account_from_zone` - (Optional) Takes a zone name value. This is used to lookup the account ID that owns this zone,
  which will be used to configure the API client. If `account_id` is also specified, this field will be ignored.
  This can also be specified with the `CLOUDFLARE_FROM_ACCOUNT_ZONE` shell environment variable.
