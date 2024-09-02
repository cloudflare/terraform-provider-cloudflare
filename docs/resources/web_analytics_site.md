---
page_title: "cloudflare_web_analytics_site Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_web_analytics_site (Resource)



## Example Usage

```terraform
resource "cloudflare_web_analytics_site" "example" {
  account_id   = "f037e56e89293a057740de681ac9abbe"
  zone_tag     = "0da42c8d2132a9ddaf714f9e7c920711"
  auto_install = true
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier

### Optional

- `auto_install` (Boolean) If enabled, the JavaScript snippet is automatically injected for orange-clouded sites.
- `host` (String) The hostname to use for gray-clouded sites.
- `zone_tag` (String) The zone identifier.

### Read-Only

- `created` (String)
- `id` (String) The Web Analytics site identifier.
- `rules` (Attributes List) A list of rules. (see [below for nested schema](#nestedatt--rules))
- `ruleset` (Attributes) (see [below for nested schema](#nestedatt--ruleset))
- `site_tag` (String) The Web Analytics site identifier.
- `site_token` (String) The Web Analytics site token.
- `snippet` (String) Encoded JavaScript snippet.

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

- `host` (String) The hostname the rule will be applied to.
- `id` (String) The Web Analytics rule identifier.
- `inclusive` (Boolean) Whether the rule includes or excludes traffic from being measured.
- `is_paused` (Boolean) Whether the rule is paused or not.
- `paths` (List of String) The paths the rule will be applied to.
- `priority` (Number)

Read-Only:

- `created` (String)


<a id="nestedatt--ruleset"></a>
### Nested Schema for `ruleset`

Optional:

- `enabled` (Boolean) Whether the ruleset is enabled.
- `id` (String) The Web Analytics ruleset identifier.
- `zone_name` (String)
- `zone_tag` (String) The zone identifier.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_web_analytics_site.example <account_id>/<site_tag>
```