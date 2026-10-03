resource "cloudflare_zero_trust_resource_library_application" "example_zero_trust_resource_library_application" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  hostnames = ["example.com", "foo.com"]
  category_id = 12
  human_id = "HR"
  ip_subnets = ["192.168.1.0/24", "2001:db8::/48"]
  name = "HR"
  port_protocols = ["tcp/80", "tcp/443"]
  support_domains = ["example.com", "foo.com"]
}
