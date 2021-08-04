---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_notification_policy"
sidebar_current: "docs-cloudflare-notification_policy"
description: |-
  Provides a resource to create and manage notification policies for Cloudflare's products.
---

# cloudflare_notification_policy

Provides a resource, that manages a notification policy for Cloudflare's products. The delivery
mechanisms supported are email, webhooks, and PagerDuty.

## Example Usage

```hcl
resource "cloudflare_notification_policy" "example" {
  account_id = "c4a7362d577a6c3019a474fd6f485821"
  name = "Policy for SSL notification events"
  description = "Notification policy to alert when my SSL certificates are modified"
  enabled     =  true
  alert_type  = "universal_ssl_event_type"

  email_integration {
    id   =  "myemail@example.com"
  }

  webhooks_integration {
    id   =  "1860572c5d964d27aa0f379d13645940"
  }

  pagerduty_integration {
    id   =  "850129d136459401860572c5d964d27k"
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The ID of the account for which the notification policy has to be created.
* `name` - (Required) The name of the notification policy.
* `enabled` - (Required) The status of the notification policy, a boolean value.
* `alert_type` - (Required) The event type that will trigger the dispatch of a notification.
* `email_integration` - (Optional) The email id to which the notification should be dispatched. One of email, webhooks, or PagerDuty mechanisms is required.
* `webhooks_integration` - (Optional) The unique id of a configured webhooks endpoint to which the notification should be dispatched. One of email, webhooks, or PagerDuty mechanisms is required.
* `pagerduty_integration` - (Optional) The unique id of a configured pagerduty endpoint to which the notification should be dispatched. One of email, webhooks, or PagerDuty mechanisms is required.
* `description` - (Optional) Description of the notification policy.
* `filters` - (Optional) Optional filterable items for a policy.

## Import

An existing notification policy can be imported using the account ID and the policy ID

```
$ terraform import cloudflare_notification_policy.example 72c379d136459405d964d27aa0f18605/c4a7362d577a6c3019a474fd6f485821
```
