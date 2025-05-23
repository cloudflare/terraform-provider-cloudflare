---
page_title: "cloudflare_spectrum_application Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_spectrum_application (Data Source)



## Example Usage

```terraform
data "cloudflare_spectrum_application" "example_spectrum_application" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  app_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) App identifier.
- `zone_id` (String) Zone identifier.

### Read-Only

- `argo_smart_routing` (Boolean) Enables Argo Smart Routing for this application.
Notes: Only available for TCP applications with traffic_type set to "direct".
- `created_on` (String) When the Application was created.
- `dns` (Attributes) The name and type of DNS record for the Spectrum application. (see [below for nested schema](#nestedatt--dns))
- `edge_ips` (Attributes) The anycast edge IP configuration for the hostname of this application. (see [below for nested schema](#nestedatt--edge_ips))
- `id` (String) App identifier.
- `ip_firewall` (Boolean) Enables IP Access Rules for this application.
Notes: Only available for TCP applications.
- `modified_on` (String) When the Application was last modified.
- `origin_direct` (List of String) List of origin IP addresses. Array may contain multiple IP addresses for load balancing.
- `origin_dns` (Attributes) The name and type of DNS record for the Spectrum application. (see [below for nested schema](#nestedatt--origin_dns))
- `origin_port` (Dynamic) The destination port at the origin. Only specified in conjunction with origin_dns. May use an integer to specify a single origin port, for example `1000`, or a string to specify a range of origin ports, for example `"1000-2000"`.
Notes: If specifying a port range, the number of ports in the range must match the number of ports specified in the "protocol" field.
- `protocol` (String) The port configuration at Cloudflare's edge. May specify a single port, for example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
- `proxy_protocol` (String) Enables Proxy Protocol to the origin. Refer to [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/) for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple Proxy Protocol.
Available values: "off", "v1", "v2", "simple".
- `tls` (String) The type of TLS termination associated with the application.
Available values: "off", "flexible", "full", "strict".
- `traffic_type` (String) Determines how data travels from the edge to your origin. When set to "direct", Spectrum will send traffic directly to your origin, and the application's type is derived from the `protocol`. When set to "http" or "https", Spectrum will apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and the application type matches this property exactly.
Available values: "direct", "http", "https".

<a id="nestedatt--dns"></a>
### Nested Schema for `dns`

Read-Only:

- `name` (String) The name of the DNS record associated with the application.
- `type` (String) The type of DNS record associated with the application.
Available values: "CNAME", "ADDRESS".


<a id="nestedatt--edge_ips"></a>
### Nested Schema for `edge_ips`

Read-Only:

- `connectivity` (String) The IP versions supported for inbound connections on Spectrum anycast IPs.
Available values: "all", "ipv4", "ipv6".
- `ips` (List of String) The array of customer owned IPs we broadcast via anycast for this hostname and application.
- `type` (String) The type of edge IP configuration specified. Dynamically allocated edge IPs use Spectrum anycast IPs in accordance with the connectivity you specify. Only valid with CNAME DNS names.
Available values: "dynamic", "static".


<a id="nestedatt--origin_dns"></a>
### Nested Schema for `origin_dns`

Read-Only:

- `name` (String) The name of the DNS record associated with the origin.
- `ttl` (Number) The TTL of our resolution of your DNS record in seconds.
- `type` (String) The type of DNS record associated with the origin. "" is used to specify a combination of A/AAAA records.
Available values: "", "A", "AAAA", "SRV".


