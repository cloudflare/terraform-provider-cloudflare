---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_worker_cron_trigger"
sidebar_current: "docs-cloudflare-resource-worker-cron-trigger"
description: |-
  Provides a Cloudflare Worker Cron Trigger resource.
---

# cloudflare_worker_cron_trigger

Worker Cron Triggers allow users to map a cron expression to a Worker script
using a `ScheduledEvent` listener that enables Workers to be executed on a
schedule. Worker Cron Triggers are ideal for running periodic jobs for
maintenance or calling third-party APIs to collect up-to-date data.

## Example Usage

```hcl
resource "cloudflare_worker_script" "example_script" {
  name    = "example-script"
  content = file("path/to/my.js")
}

resource "cloudflare_worker_cron_trigger" "example_trigger" {
  script_name = cloudflare_worker_script.example_script.name
  schedules   = [
    "*/5 * * * *",      # every 5 minutes
    "10 7 * * mon-fri", # 7:10am every weekday
  ]
}
```

## Argument Reference

The following arguments are supported:


* `script_name` - (Required) Worker script to target for the schedules
* `schedules` - (Required) List of cron expressions to execute the Worker Script

## Attributes Reference

The following additional attributes are exported:

* `id` - md5 checksum of the script name
* `script_name` - Name of the Worker Script being targeted
* `schedules` - List of cron expressions in use

## Import

Worker Cron Triggers can be imported using the script name of the Worker they
are targeting.

```
$ terraform import cloudflare_worker_cron_trigger.example my-script
```
