---
page_title: "cloudflare_web3_hostname Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_web3_hostname (Resource)



## Example Usage

```terraform
resource "cloudflare_web3_hostname" "example_web3_hostname" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "gateway.example.com"
  target = "ipfs"
  description = "This is my IPFS gateway."
  dnslink = "/ipns/onboarding.ipfs.cloudflare.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The hostname that will point to the target gateway via CNAME.
- `target` (String) Target gateway of the hostname.
Available values: "ethereum", "ipfs", "ipfs_universal_path".
- `zone_id` (String) Identifier

### Optional

- `description` (String) An optional description of the hostname.
- `dnslink` (String) DNSLink value used if the target is ipfs.

### Read-Only

- `created_on` (String)
- `id` (String) Identifier
- `modified_on` (String)
- `status` (String) Status of the hostname's activation.
Available values: "active", "pending", "deleting", "error".

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_web3_hostname.example '<zone_id>/<identifier>'
```
