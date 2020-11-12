---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_workers_secret"
sidebar_current: "docs-cloudflare-resource-workers-secret"
description: |-
  Provides the ability to manage Cloudflare Workers secrets individually.
---

# cloudflare_worker_secret

Provides a Worker secret via a Worker script name, secret name, and secret text.  

*NOTE:*

- Must be used **mutually exclusive** to `cloudflare_worker_script` as they conflict within a Terraform plan.
- This resource uses the Cloudflare account APIs. This requires setting the `CLOUDFLARE_ACCOUNT_ID` environment variable or `account_id` provider argument.


## Example Usage

```hcl
resource "cloudflare_worker_secret" "my_secret" {
  script_name = "my_worker"
  name = "MY_EXAMPLE_SECRET_TEXT"
  text = var.secret_foo_value
}
```

## Argument Reference

The following arguments are supported:

* `script_name` - (Required) The name of the Worker script against which you want to create the secret
* `name` - (Required) The secret name
* `secret_text` - (Required) The string value to be stored as a secret
