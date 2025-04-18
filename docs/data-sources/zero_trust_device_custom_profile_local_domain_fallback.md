---
page_title: "cloudflare_zero_trust_device_custom_profile_local_domain_fallback Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_device_custom_profile_local_domain_fallback (Data Source)



## Example Usage

```terraform
data "cloudflare_zero_trust_device_custom_profile_local_domain_fallback" "example_zero_trust_device_custom_profile_local_domain_fallback" {
  account_id = "699d98642c564d2e855e9661899b7252"
  policy_id = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String)
- `policy_id` (String)

### Read-Only

- `description` (String) A description of the fallback domain, displayed in the client UI.
- `dns_server` (List of String) A list of IP addresses to handle domain resolution.
- `suffix` (String) The domain suffix to match when resolving locally.


