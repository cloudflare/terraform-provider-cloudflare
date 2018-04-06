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

# Create a page rule
resource "cloudflare_page_rule" "www" {
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
  This can also be specified with the `CLOUDFLARE_RPS` shell environment variable.
* `retries` - (Optional) Maximum number of retries to perform when an API request fails. Default: 3.
  This can also be specified with the `CLOUDFLARE_RETRIES` shell environment variable.
* `min_backoff` - (Optional) Minimum backoff period in seconds after failed API calls. Default: 1.
  This can also be specified with the `CLOUDFLARE_MIN_BACKOFF` shell environment variable.
* `max_backoff` - (Optional) Maximum backoff period in seconds after failed API calls Default: 30.
  This can also be specified with the `CLOUDFLARE_MAX_BACKOFF` shell environment variable.
* `api_client_logging` - (Optional) Whether to print logs from the API client (using the default log library logger). Default: false.
  This can also be specified with the `CLOUDFLARE_API_CLIENT_LOGGING` shell environment variable.
* `org_id` - (Optional) Configure API client with this organisation ID, so calls use the organization API rather than the (default) user API.
  This is required for other users in your organization to have access to the resources you manage.
  This can also be specified with the `CLOUDFLARE_ORG_ID` shell environment variable.
* `use_org_from_zone` - (Optional) Takes a zone name value. This is used to lookup the organization ID that owns this zone, 
  which will be used to configure the API client. If `org_id` is also specified, this field will be ignored.
  This can also be specified with the `CLOUDFLARE_ORG_ZONE` shell environment variable.

