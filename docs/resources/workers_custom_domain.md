---
page_title: "cloudflare_workers_custom_domain Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_workers_custom_domain (Resource)



## Example Usage

```terraform
resource "cloudflare_workers_custom_domain" "example_workers_custom_domain" {
  account_id = "9a7806061c88ada191ed06f989cc3dac"
  environment = "production"
  hostname = "foo.example.com"
  service = "foo"
  zone_id = "593c9c94de529bbbfaac7c53ced0447d"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifer of the account.
- `environment` (String) Worker environment associated with the zone and hostname.
- `hostname` (String) Hostname of the Worker Domain.
- `service` (String) Worker service associated with the zone and hostname.
- `zone_id` (String) Identifier of the zone.

### Read-Only

- `id` (String) Identifer of the Worker Domain.
- `zone_name` (String) Name of the zone.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_workers_custom_domain.example '<account_id>/<domain_id>'
```
