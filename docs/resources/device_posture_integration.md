---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_device_posture_integration"
description: Provides a Cloudflare Device Posture Integration resource.
---

# cloudflare_device_posture_integration

Provides a Cloudflare Device Posture Integration resource. Device posture integrations configure third-party data providers for device posture rules.

## Example Usage

```hcl
resource "cloudflare_device_posture_integration" "third_party_devices_posture_integration" {
  account_id  = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name        = "Device posture integration"
  type        = "workspace_one"
  interval    = "24h"
  config {
      api_url       = "https://example.com/api"
      auth_url      = "https://example.com/connect/token"
      client_id     = "client-id"
      client_secret = "client-secret"
  }
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The account to which the device posture integration should be added.
- `name` - (Optional) Name of the device posture integration.
- `type` - (Required) The device posture integration type. Valid values are `workspace_one`.
- `interval` - (Optional) Indicates the frequency with which to poll the third-party API.
  Must be in the format `"1h"` or `"30m"`. Valid units are `h` and `m`.
- `config` - (Required) The device posture integration's connection authorization parameters.

### Config argument

The The config structure depends on the integration type.

**ws1** allows the following:

- `api_url` - (Required) The third-party API's URL.
- `auth_url` - (Required) The third-party authorization API URL.
- `client_id` - (Required) The client identifier for authenticating API calls.
- `client_secret` - (Required) The client secret for authenticating API calls.

**crowdstrike_s2s** allows the following:

* `api_url` - (Required) The third-party API's URL.
* `customer_id` - (Required)  The customer identifier for authenticating API calls.
* `client_id` - (Required) The client identifier for authenticating API calls.
* `client_secret` - (Required) The client secret for authenticating API calls.

**uptycs** allows the following:

* `client_key` - (Required) The client key for authenticating API calls.
* `client_id` - (Required) The client identifier for authenticating API calls.
* `client_secret` - (Required) The client secret for authenticating API calls.

**intune** allows the following:
* `customer_id` - (Required)  The customer identifier for authenticating API calls.
* `client_id` - (Required) The client identifier for authenticating API calls.
* `client_secret` - (Required) The client secret for authenticating API calls.

## Attributes Reference

The following additional attributes are exported:

- `id` - ID of the device posture integration.

## Import

Device posture integrations can be imported using a composite ID formed of account
ID and device posture integration ID.

```
$ terraform import cloudflare_device_posture_integration.corporate_devices cb029e245cfdd66dc8d2e570d5dd3322/0ade592a-62d6-46ab-bac8-01f47c7fa792
```
