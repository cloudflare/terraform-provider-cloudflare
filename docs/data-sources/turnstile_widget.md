---
page_title: "cloudflare_turnstile_widget Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_turnstile_widget (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) Identifier
- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))
- `secret` (String) Secret key for this widget.
- `sitekey` (String) Widget item identifier tag.

### Read-Only

- `bot_fight_mode` (Boolean) If bot_fight_mode is set to `true`, Cloudflare issues computationally
expensive challenges in response to malicious bots (ENT only).
- `clearance_level` (String) If Turnstile is embedded on a Cloudflare site and the widget should grant challenge clearance,
this setting can determine the clearance level to be set
- `created_on` (String) When the widget was created.
- `domains` (List of String)
- `mode` (String) Widget Mode
- `modified_on` (String) When the widget was modified.
- `name` (String) Human readable widget name. Not unique. Cloudflare suggests that you
set this to a meaningful string to make it easier to identify your
widget, and where it is used.
- `offlabel` (Boolean) Do not show any Cloudflare branding on the widget (ENT only).
- `region` (String) Region where this widget can be used.

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Required:

- `account_id` (String) Identifier

Optional:

- `direction` (String) Direction to order widgets.
- `order` (String) Field to order widgets by.
- `page` (Number) Page number of paginated results.
- `per_page` (Number) Number of items per page.

