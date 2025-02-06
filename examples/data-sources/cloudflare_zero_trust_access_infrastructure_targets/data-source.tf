data "cloudflare_zero_trust_access_infrastructure_targets" "example_zero_trust_access_infrastructure_targets" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  created_after = "2019-12-27T18:11:19.117Z"
  created_before = "2019-12-27T18:11:19.117Z"
  direction = "asc"
  hostname = "hostname"
  hostname_contains = "hostname_contains"
  ip_like = "ip_like"
  ip_v4 = "ip_v4"
  ip_v6 = "ip_v6"
  ips = ["string"]
  ipv4_end = "ipv4_end"
  ipv4_start = "ipv4_start"
  ipv6_end = "ipv6_end"
  ipv6_start = "ipv6_start"
  modified_after = "2019-12-27T18:11:19.117Z"
  modified_before = "2019-12-27T18:11:19.117Z"
  order = "hostname"
  target_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
  virtual_network_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
}
