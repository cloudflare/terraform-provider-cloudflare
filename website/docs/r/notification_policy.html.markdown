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

### With Filters

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
* `alert_type` - (Required) The event type that will trigger the dispatch of a notification (refer to the [nested schema](#nestedblock--alert-type)).
* `email_integration` - (Optional) The email id to which the notification should be dispatched. One of email, webhooks, or PagerDuty mechanisms is required.
* `webhooks_integration` - (Optional) The unique id of a configured webhooks endpoint to which the notification should be dispatched. One of email, webhooks, or PagerDuty mechanisms is required.
* `pagerduty_integration` - (Optional) The unique id of a configured pagerduty endpoint to which the notification should be dispatched. One of email, webhooks, or PagerDuty mechanisms is required.
* `description` - (Optional) Description of the notification policy.
* `filters` - (Optional) An optional nested block of filters that applies to the selected `alert_type`. A key-value map that specifies the type of filter and the values to match against (refer to the alert type block for available fields).

<a id="nestedblock--alert-type"></a>
**Nested schema for `alert_type`**

* `billing_usage_alert` - (Optional) Billing usage exceeds threshold (refer to the [nested schema](#nestedblock--alert-type-billing-usage-alert)).
* `health_check_status_notification` - (Optional) Health check status changes (refer to the [nested schema](#nestedblock--alert-type-health-check-status-notification)).
* `g6_pool_toggle_alert` - (Optional) Pool alerts on enable/disable status (refer to the [nested schema](#nestedblock--alert-type-g6-pool-toggle-alert)).
* `real_origin_monitoring` - (Optional) Cloudflare is unable to reach your origin (refer to the [nested schema](#nestedblock--alert-type-real-origin-monitoring)).
* `universal_ssl_event_type` - (Optional) Universal certificate notices (refer to the [nested schema](#nestedblock--alert-type-universal-ssl-event-type)).
* `bgp_hijack_notification` - (Optional) Alerts for BGP hijack (refer to the [nested schema](#nestedblock--alert-type-bgp-hijack-notification)).
* `http_alert_origin_error` - (Optional) HTTP origin error rate alert  (refer to the [nested schema](#nestedblock--alert-type-http-alert-origin-error)).

<a id="nestedblock--alert-type-billing-usage-alert"></a>
**Nested schema for `billing_usage_alert`**

* `product` - (Optional) Product name. Available values: `"worker_requests"`, `"worker_durable_objects_requests"`, `"worker_durable_objects_duration"`, `"worker_durable_objects_data_transfer"`, `"worker_durable_objects_stored_data"`, `"worker_durable_objects_storage_deletes"`, `"worker_durable_objects_storage_writes"`, `"worker_durable_objects_storage_reads"`.
* `limit` - (Optional) A numerical limit. Example: `"100"`

<a id="nestedblock--alert-type-health-check-status-notification"></a>
**Nested schema for `health_check_status_notification`**

* `health_check_id` - (Optional) Identifier health check.
* `status` - (Optional) Status to alert on. Example: `"Unhealthy"`, `"Healthy"`.

<a id="nestedblock--alert-type-g6-pool-toggle-alert"></a>
**Nested schema for `g6_pool_toggle_alert`**

* `pool_id` - (Optional) Load balancer pool identifier.
* `enabled` - (Optional) State of the pool to alert on. Example: `"true"`, `"false"`.

<a id="#nestedblock--alert-type-real-origin-monitoring"></a>
**Nested schema for `real_origin_monitoring`**

<a id="#nestedblock--alert-type-universal-ssl-event-type"></a>
**Nested schema for `universal_ssl_event_type`**

<a id="#nestedblock--alert-type-bgp-hijack-notification"></a>
**Nested schema for `bgp_hijack_notification`**

<a id="nestedblock--alert-type-http-alert-origin-error"></a>
**Nested schema for `http_alert_origin_error`**

* `zones` - (Optional) A list of zone identifiers.
* `slo` - (Optional) A numerical limit. Example: `"99.9"`


## Import

An existing notification policy can be imported using the account ID and the policy ID

```
$ terraform import cloudflare_notification_policy.example 72c379d136459405d964d27aa0f18605/c4a7362d577a6c3019a474fd6f485821
```
