---
page_title: "cloudflare_snippets Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_snippets (Data Source)



## Example Usage

```terraform
data "cloudflare_snippets" "example_snippets" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  snippet_name = "snippet_name_01"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `snippet_name` (String) Snippet identifying name
- `zone_id` (String) Identifier

### Read-Only

- `created_on` (String) Creation time of the snippet
- `modified_on` (String) Modification time of the snippet


