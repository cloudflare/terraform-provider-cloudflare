---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_ip_ranges"
description: Get information on Cloudflare IP ranges.
---

# cloudflare_ip_ranges

Use this data source to get the [IP ranges][1] of Cloudflare edge nodes.

## Example Usage

```hcl
data "cloudflare_ip_ranges" "cloudflare" {}

resource "google_compute_firewall" "allow_cloudflare_ingress" {
  name    = "from-cloudflare"
  network = "default"

  source_ranges = data.cloudflare_ip_ranges.cloudflare.ipv4_cidr_blocks

  allow {
    ports    = "443"
    protocol = "tcp"
  }
}
```

## Attributes Reference

- `cidr_blocks` - The lexically ordered list of all non-China CIDR blocks.
- `ipv4_cidr_blocks` - The lexically ordered list of only the IPv4 CIDR blocks.
- `ipv6_cidr_blocks` - The lexically ordered list of only the IPv6 CIDR blocks.
- `china_ipv4_cidr_blocks` - The lexically ordered list of only the IPv4 China CIDR blocks.
- `china_ipv6_cidr_blocks` - The lexically ordered list of only the IPv6 China CIDR blocks.

[1]: https://www.cloudflare.com/ips/
