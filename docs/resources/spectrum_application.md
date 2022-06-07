---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_spectrum_application"
description: Provides a Cloudflare Spectrum Application resource.
---

# cloudflare_spectrum_application

Provides a Cloudflare Spectrum Application. You can extend the power of Cloudflare's DDoS, TLS, and IP Firewall to your other TCP-based services.

## Example Usage

```hcl
# Define a spectrum application proxies ssh traffic
resource "cloudflare_spectrum_application" "ssh_proxy" {
  zone_id      = var.cloudflare_zone_id
  protocol     = "tcp/22"
  traffic_type = "direct"
  dns {
    type = "CNAME"
    name = "ssh.example.com"
  }

  origin_direct = [
    "tcp://109.151.40.129:22"
  ]
}
```

## Argument Reference

- `zone_id` - (Required) The DNS zone ID to add the application to
- `protocol` - (Required) The port configuration at Cloudflare’s edge. e.g. `tcp/22`.
- `dns` - (Required) The name and type of DNS record for the Spectrum application. Fields documented below.
- `origin_direct` - (Optional) A list of destination addresses to the origin. e.g. `tcp://192.0.2.1:22`.
- `origin_dns` - (Optional) A destination DNS addresses to the origin. Fields documented below.
- `origin_port` - (Optional) If using `origin_dns` and not `origin_port_range`, this is a required attribute. Origin port to proxy traffice to e.g. `22`.
- `origin_port_range` - (Optional) If using `origin_dns` and not `origin_port`, this is a required attribute. Origin port range to proxy traffice to. When using a range, the protocol field must also specify a range, e.g. `tcp/22-23`. Fields documented below.
- `tls` - (Optional) TLS configuration option for Cloudflare to connect to your origin. Valid values are: `off`, `flexible`, `full` and `strict`. Defaults to `off`.
- `ip_firewall` - (Optional) Enables the IP Firewall for this application. Defaults to `true`.
- `proxy_protocol` - (Optional) Enables a proxy protocol to the origin. Valid values are: `off`, `v1`, `v2`, and `simple`. Defaults to `off`.
- `traffic_type` - (Optional) Sets application type. Valid values are: `direct`, `http`, `https`. Defaults to `direct`.
- `argo_smart_routing` - (Optional). Enables Argo Smart Routing. Defaults to `false`.
- `edge_ip_connectivity` - (Optional). Choose which types of IP addresses will be provisioned for this subdomain. Valid values are: `all`, `ipv4`, `ipv6`. Defaults to `all`.
- `edge_ips` - (Optional). A list of edge IPs (IPv4 and/or IPv6) to configure Spectrum application to. Requires [Bring Your Own IP](https://developers.cloudflare.com/spectrum/getting-started/byoip/) provisioned.

**dns**

- `type` - (Required) The type of DNS record associated with the application. Valid values: `CNAME`.
- `name` - (Required) The name of the DNS record associated with the application.i.e. `ssh.example.com`.

**origin_dns**

- `name` - (Required) Fully qualified domain name of the origin e.g. origin-ssh.example.com.

**origin_port_range**

- `start` - (Required) Lower bound of the origin port range, e.g. `1000`
- `end` - (Required) Upper bound of the origin port range, e.g. `2000`

## Attributes Reference

The following attributes are exported:

- `id` - Unique identifier in the API for the spectrum application.

## Import

Spectrum resource can be imported using a zone ID and Application ID, e.g.

```
$ terraform import cloudflare_spectrum_application.example d41d8cd98f00b204e9800998ecf8427e/9a7806061c88ada191ed06f989cc3dac
```

where:

- `d41d8cd98f00b204e9800998ecf8427e` - zone ID, as returned from [API](https://api.cloudflare.com/#zone-list-zones)
- `9a7806061c88ada191ed06f989cc3dac` - Application ID
