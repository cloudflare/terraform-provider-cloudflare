---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_notification_policy_webhooks"
description: Provides a resource to create and manage webhooks destinations for Cloudflare's notification policies.
---

# cloudflare_notification_policy_webhooks

Provides a resource, that manages a webhook destination. These destinations can be tied to the notification policies created for Cloudflare's products.

## Example Usage

```hcl
resource "cloudflare_notification_policy_webhooks" "example" {
  account_id = "c4a7362d577a6c3019a474fd6f485821"
  name       = "Webhooks destination"
  url        = "https://example.com"
  secret     =  "my-secret"
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The ID of the account for which the webhook destination has to be connected.
- `name` - (Required) The name of the webhook destination.
- `url` - (Required) The URL of the webhook destinations.
- `secret` - (Optional) An optional secret can be provided that will be passed in the `cf-webhook-auth` header when dispatching a webhook notification.
  Secrets are not returned in any API response body.
  Refer to the documentation for more details - https://api.cloudflare.com/#notification-webhooks-create-webhook.

## Import

An existing notification policy can be imported using the account ID and the webhook ID

```
$ terraform import cloudflare_notification_policy_webhooks.example 72c379d136459405d964d27aa0f18605/c4a7362d577a6c3019a474fd6f485821
```
