data "cloudflare_zero_trust_resource_library_categories" "test" {
  account_id = "%[2]s"
}

resource "cloudflare_zero_trust_resource_library_application" "%[1]s" {
  account_id  = "%[2]s"
  category_id = min(data.cloudflare_zero_trust_resource_library_categories.test.result[*].id...)
  human_id    = "%[1]s"
  name        = "%[1]s"

  hostnames       = ["%[3]s.example.com"]
  ip_subnets      = ["%[4]s"]
  port_protocols  = ["%[5]s"]
  support_domains = ["support-%[3]s.example.com"]
}
