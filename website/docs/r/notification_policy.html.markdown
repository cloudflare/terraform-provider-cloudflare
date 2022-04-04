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

### Basic Example
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

Example With Filters
```hcl
resource "cloudflare_notification_policy" "example" {
  account_id  = "c4a7362d577a6c3019a474fd6f485821"
  name        = "Policy for Healthcheck notification"
  description = "Notification policy to alert on unhealthy Healthcheck status"
  enabled     =  true
  alert_type  = "health_check_status_notification"

  email_integration {
    id   =  "myemail@example.com"
  }

  webhooks_integration {
    id   =  "1860572c5d964d27aa0f379d13645940"
  }

  pagerduty_integration {
    id   =  "850129d136459401860572c5d964d27k"
  }
  
  filters {
    health_check_id = ["699d98642c564d2e855e9661899b7252"]
    status           = ["Unhealthy"]
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
* `filters` - (Optional) An optional nested block of filters that applies to the selected `alert_type`. A key-value map that specifies the type of filter and the values to match against.

## Filters

| Alert Type                       | Filter          | Description                               | Example Values                                                                                                                                                                                                                                                                                    |
|----------------------------------|-----------------|-------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| billing_usage_alert              |                 | billing usage exceeds threshold           |                                                                                                                                                                                                                                                                                                   |
|                                  | product         | product name                              | "worker_requests", "worker_durable_objects_requests", "worker_durable_objects_duration", "worker_durable_objects_data_transfer", "worker_durable_objects_stored_data", "worker_durable_objects_storage_deletes", "worker_durable_objects_storage_writes", "worker_durable_objects_storage_reads"  |
|                                  | limit           | a numerical limit                         | "100"                                                                                                                                                                                                                                                                                             |
| health_check_status_notification |                 | health check status changes               |                                                                                                                                                                                                                                                                                                   |
|                                  | health_check_id | health check ID                           | 699d98642c564d2e855e9661899b7252                                                                                                                                                                                                                                                                  |
|                                  | status          | status to alert on                        | "Unhealthy", "Healthy"                                                                                                                                                                                                                                                                            |
| g6_pool_toggle_alert             |                 | pool alerts on enable/disable status      |                                                                                                                                                                                                                                                                                                   |
|                                  | pool_id         | load balancing pool id                    | "17b5962d775c646f3f9725cbc7a53df4"                                                                                                                                                                                                                                                                |
|                                  | enabled         | state to alert on                         | "true", "false"                                                                                                                                                                                                                                                                                   |
| real_origin_monitoring           |                 | Cloudflare is unable to reach your origin |                                                                                                                                                                                                                                                                                                   |
| universal_ssl_event_type         |                 | universal certificate notices             |                                                                                                                                                                                                                                                                                                   |
| bgp_hijack_notification          |                 | alerts for BGP hijack                     |                                                                                                                                                                                                                                                                                                   |

## Import

An existing notification policy can be imported using the account ID and the policy ID

```
$ terraform import cloudflare_notification_policy.example 72c379d136459405d964d27aa0f18605/c4a7362d577a6c3019a474fd6f485821
```
