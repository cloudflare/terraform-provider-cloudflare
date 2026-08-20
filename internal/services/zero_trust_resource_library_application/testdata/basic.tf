data "cloudflare_zero_trust_resource_library_categories" "test" {
  account_id = "%[2]s"
  max_items  = 1
}

resource "cloudflare_zero_trust_resource_library_application" "%[1]s" {
  account_id  = "%[2]s"
  category_id = tonumber(data.cloudflare_zero_trust_resource_library_categories.test.result[0].id)
  human_id    = "%[1]s"
  name        = "%[1]s"

  hostnames       = ["%[3]s.example.com"]
  ip_subnets      = ["%[4]s"]
  port_protocols  = ["%[5]s"]
  support_domains = ["support-%[3]s.example.com"]
}
