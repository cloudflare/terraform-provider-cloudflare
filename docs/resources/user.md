---
page_title: "cloudflare_user Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_user (Resource)



## Example Usage

```terraform
resource "cloudflare_user" "example_user" {
  country = "US"
  first_name = "John"
  last_name = "Appleseed"
  telephone = "+1 123-123-1234"
  zipcode = "12345"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `country` (String) The country in which the user lives.
- `first_name` (String) User's first name
- `last_name` (String) User's last name
- `telephone` (String) User's telephone number
- `zipcode` (String) The zipcode or postal code where the user lives.


