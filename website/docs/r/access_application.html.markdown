---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_access_application"
sidebar_current: "docs-cloudflare-resource-access-application"
description: |-
  Provides a Cloudflare Access Application resource.
---

# cloudflare_access_application

Provides a Cloudflare Access Application resource. Access Applications
are used to restrict access to a whole application using an
authorisation gateway managed by Cloudflare.

## Example Usage

```hcl
resource "cloudflare_access_application" "staging_app" {
  zone_id          = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name             = "staging application"
  domain           = "staging.example.com"
  session_duration = "24h"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone to which the access rule should be added.
* `name` - (Required) Friendly name of the Access Application.
* `domain` - (Required) The complete URL of the asset you wish to put
  Cloudflare Access in front of. Can include subdomains or paths. Or both.
* `session_duration` - (Optional) How often a user will be forced to
  re-authorise. Must be one of `30m`, `6h`, `12h`, `24h`, `168h`, `730h`.
  
## Attributes Reference

The following additional attributes are exported:

* `id` - ID of the application
* `aud` - Application Audience (AUD) Tag of the application
* `domain` - Domain of the application
* `session_duration` - Session duriation of the application


## Import

Access Applications can be imported using a composite ID formed of zone
ID and application ID.

```
$ terraform import cloudflare_access_application.staging cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e
```
