---
layout: "cloudflare"
page_title: "Upgrading to version 2 (from 1.x)"
description: Terraform Cloudflare Provider Version 2 Upgrade Guide
---

# Terraform Cloudflare Provider Version 2 Upgrade Guide

Version 2 of the Cloudflare Terraform Provider is introducing several breaking changes intended to remove confusion
around different ways of specifying zones and Worker resources, and accommodates for API changes around Workers product.

## Provider Version Configuration

If you are not ready to make a move to version 2 of the Cloudflare provider, you may keep the 1.x branch active for
your Terraform project by specifying:

```hcl
provider "cloudflare" {
  version = "~> 1.0"
  # ... any other configuration
}
```

We highly recommend to review this guide, make necessary changes and move to 2.x branch, as further 1.x releases are
unlikely to happen.

~> Before attempting to upgrade to version 2, you should first upgrade to the
   latest version of 1 to ensure any transitional updates are applied to your
   existing configuration.

Once ready, make the following change to use the latest 2.x release:

```hcl
provider "cloudflare" {
  version = "~> 2.0"
  # ... any other configuration
}
```

## Provider global configuration changes

The following changes have been made to fields:

- renamed `token` to `api_key`
- renamed `org_id` to `account_id`
- removed `use_org_from_zone`, you need to explicitly specify `account_id`

The following changes have been made to environment variables:

- renamed `CLOUDFLARE_TOKEN` to `CLOUDFLARE_API_KEY`
- renamed `CLOUDFLARE_ORG_ID` to `CLOUDFLARE_ACCOUNT_ID`
- removed `CLOUDFLARE_ORG_ZONE`, you need to explicitly specify `CLOUDFLARE_ACCOUNT_ID`

Before:

```hcl
provider "cloudflare" {
  version = "~> 1.0"

  email  = "terraform@example.com"
  token  = "a647b7f10e7b7374d206817a7f92b642"
  org_id = "975ecf5a45e3bcb680dba0722a420ad9"
}
```

After:

```hcl
provider "cloudflare" {
  version = "~> 2.0"

  email      = "terraform@example.com"
  api_key    = "a647b7f10e7b7374d206817a7f92b642"
  account_id = "975ecf5a45e3bcb680dba0722a420ad9"
}
```

## Zone Name to Zone ID changes

All resources that accepted Zone Name have been changed to accept Zone ID instead. You can find the Zone ID in the Cloudflare Dashboard on the overview page in the right hand side navigation.

The following resources now require Zone IDs:

- `cloudflare_access_rule`
- `cloudflare_filter`
- `cloudflare_firewall_rule`
- `cloudflare_load_balancer`
- `cloudflare_page_rule`
- `cloudflare_rate_limit`
- `cloudflare_record`
- `cloudflare_waf_rule`
- `cloudflare_worker_route`
- `cloudflare_zone_lockdown`
- `cloudflare_zone_settings_override`

Before:

```hcl
resource "cloudflare_zone_lockdown" "example" {
  zone = "example.com"
  # ...
}

resource "cloudflare_record" "foobar" {
  domain = "example.com"
  name   = "terraform"
  value  = "192.168.0.11"
  type   = "A"
  ttl    = 3600
}
```

After:

```hcl
resource "cloudflare_zone_lockdown" "example" {
  zone_id = "43feed7a08b85f654aa54ca9d61bb0c0"
  # ...
}

resource "cloudflare_record" "foobar" {
  zone_id = "43feed7a08b85f654aa54ca9d61bb0c0"
  name    = "terraform"
  value   = "192.168.0.11"
  type    = "A"
  ttl     = 3600
}
```

## Workers single-script support removed

Formerly Enterprise-only APIs for configuring multiple Worker scripts are now available for all customers. Therefore,
there is no longer need for single-script support, which works in compatibility mode now.

Before:

```hcl
resource "cloudflare_worker_script" "my_script" {
  zone = "example.com"
  content = "${file("script.js")}"
}

resource "cloudflare_worker_route" "my_route" {
  zone = "example.com"
  pattern = "example.com/*"
  enabled = true
  depends_on = ["cloudflare_worker_script.my_script"]
}
```

After:

```hcl
# Sets the script with the name "script_1"
resource "cloudflare_worker_script" "my_script" {
  name = "script_1"
  content = "${file("script.js")}"
}

# Runs the specified worker script for all URLs that match `example.com/*`
resource "cloudflare_worker_route" "my_route" {
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  pattern = "example.com/*"
  script_name = "${cloudflare_worker_script.my_script.name}"
}
```
