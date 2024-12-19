---
page_title: "cloudflare_magic_transit_site_acl Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_magic_transit_site_acl (Data Source)



## Example Usage

```terraform
data "cloudflare_magic_transit_site_acl" "example_magic_transit_site_acl" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  site_id = "023e105f4ecef8ad9ca31a8372d0c353"
  acl_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) Identifier
- `acl_id` (String) Identifier
- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))
- `site_id` (String) Identifier

### Read-Only

- `description` (String) Description for the ACL.
- `forward_locally` (Boolean) The desired forwarding action for this ACL policy. If set to "false", the policy will forward traffic to Cloudflare. If set to "true", the policy will forward traffic locally on the Magic Connector. If not included in request, will default to false.
- `id` (String) Identifier
- `lan_1` (Attributes) (see [below for nested schema](#nestedatt--lan_1))
- `lan_2` (Attributes) (see [below for nested schema](#nestedatt--lan_2))
- `name` (String) The name of the ACL.
- `protocols` (List of String)

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Required:

- `account_id` (String) Identifier
- `site_id` (String) Identifier


<a id="nestedatt--lan_1"></a>
### Nested Schema for `lan_1`

Read-Only:

- `lan_id` (String) The identifier for the LAN you want to create an ACL policy with.
- `lan_name` (String) The name of the LAN based on the provided lan_id.
- `ports` (List of Number) Array of ports on the provided LAN that will be included in the ACL. If no ports are provided, communication on any port on this LAN is allowed.
- `subnets` (List of String) Array of subnet IPs within the LAN that will be included in the ACL. If no subnets are provided, communication on any subnets on this LAN are allowed.


<a id="nestedatt--lan_2"></a>
### Nested Schema for `lan_2`

Read-Only:

- `lan_id` (String) The identifier for the LAN you want to create an ACL policy with.
- `lan_name` (String) The name of the LAN based on the provided lan_id.
- `ports` (List of Number) Array of ports on the provided LAN that will be included in the ACL. If no ports are provided, communication on any port on this LAN is allowed.
- `subnets` (List of String) Array of subnet IPs within the LAN that will be included in the ACL. If no subnets are provided, communication on any subnets on this LAN are allowed.

