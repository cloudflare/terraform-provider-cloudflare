---
page_title: "cloudflare_email_security_trusted_domains Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_email_security_trusted_domains (Resource)



## Example Usage

```terraform
resource "cloudflare_email_security_trusted_domains" "example_email_security_trusted_domains" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  is_recent = true
  is_regex = false
  is_similarity = false
  pattern = "example.com"
  comments = null
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Account Identifier

### Optional

- `body` (Attributes List) (see [below for nested schema](#nestedatt--body))
- `comments` (String)
- `is_recent` (Boolean) Select to prevent recently registered domains from triggering a
Suspicious or Malicious disposition.
- `is_regex` (Boolean)
- `is_similarity` (Boolean) Select for partner or other approved domains that have similar
spelling to your connected domains. Prevents listed domains from
triggering a Spoof disposition.
- `pattern` (String)

### Read-Only

- `created_at` (String)
- `id` (Number) The unique identifier for the trusted domain.
- `last_modified` (String)

<a id="nestedatt--body"></a>
### Nested Schema for `body`

Required:

- `is_recent` (Boolean) Select to prevent recently registered domains from triggering a
Suspicious or Malicious disposition.
- `is_regex` (Boolean)
- `is_similarity` (Boolean) Select for partner or other approved domains that have similar
spelling to your connected domains. Prevents listed domains from
triggering a Spoof disposition.
- `pattern` (String)

Optional:

- `comments` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_email_security_trusted_domains.example '<account_id>/<trusted_domain_id>'
```