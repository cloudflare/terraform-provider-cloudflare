---
page_title: "cloudflare_zone_lockdown Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zone_lockdown (Resource)



## Example Usage

```terraform
# Restrict access to these endpoints to requests from a known IP address range.
resource "cloudflare_zone_lockdown" "example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  paused      = "false"
  description = "Restrict access to these endpoints to requests from a known IP address range"
  urls = [
    "api.mysite.com/some/endpoint*",
  ]
  configurations = [{
    target = "ip_range"
    value  = "192.0.2.0/24"
  }]
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_identifier` (String) Identifier

### Read-Only

- `configurations` (Attributes) A list of IP addresses or CIDR ranges that will be allowed to access the URLs specified in the Zone Lockdown rule. You can include any number of `ip` or `ip_range` configurations. (see [below for nested schema](#nestedatt--configurations))
- `created_on` (String) The timestamp of when the rule was created.
- `description` (String) An informative summary of the rule.
- `id` (String) The unique identifier of the Zone Lockdown rule.
- `modified_on` (String) The timestamp of when the rule was last modified.
- `paused` (Boolean) When true, indicates that the rule is currently paused.
- `urls` (List of String) The URLs to include in the rule definition. You can use wildcards. Each entered URL will be escaped before use, which means you can only use simple wildcard patterns.

<a id="nestedatt--configurations"></a>
### Nested Schema for `configurations`

Optional:

- `target` (String) The configuration target. You must set the target to `ip` when specifying an IP address in the Zone Lockdown rule.
- `value` (String) The IP address to match. This address will be compared to the IP address of incoming requests.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_zone_lockdown.example <zone_id>/<lockdown_id>
```