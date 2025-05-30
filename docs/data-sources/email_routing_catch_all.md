---
page_title: "cloudflare_email_routing_catch_all Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_email_routing_catch_all (Data Source)



## Example Usage

```terraform
data "cloudflare_email_routing_catch_all" "example_email_routing_catch_all" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) Identifier.

### Read-Only

- `actions` (Attributes List) List actions for the catch-all routing rule. (see [below for nested schema](#nestedatt--actions))
- `enabled` (Boolean) Routing rule status.
- `id` (String) Routing rule identifier.
- `matchers` (Attributes List) List of matchers for the catch-all routing rule. (see [below for nested schema](#nestedatt--matchers))
- `name` (String) Routing rule name.
- `tag` (String, Deprecated) Routing rule tag. (Deprecated, replaced by routing rule identifier)

<a id="nestedatt--actions"></a>
### Nested Schema for `actions`

Read-Only:

- `type` (String) Type of action for catch-all rule.
Available values: "drop", "forward", "worker".
- `value` (List of String)


<a id="nestedatt--matchers"></a>
### Nested Schema for `matchers`

Read-Only:

- `type` (String) Type of matcher. Default is 'all'.
Available values: "all".


