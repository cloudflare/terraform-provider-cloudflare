resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  description = "desc"
  precedence = 12303
  action = "block"
  filters = ["dns"]
  traffic = "any(dns.domains[*] == \"example.com\")"
  identity = "any(identity.groups.name[*] in {\"finance\"})"
  rule_settings = {
    block_page_enabled = true
    block_reason = "cuzs"
    ip_categories = true
    ip_indicator_feeds = true
    ignore_cname_category_matches = true
  }
  schedule = {
    fri = "08:00-12:30,13:30-17:00"
    mon = "08:00-12:30,13:30-17:00"
    sat = "08:00-12:30,13:30-17:00"
    sun = "08:00-12:30,13:30-17:00"
    thu = "08:00-12:30,13:30-17:00"
    time_zone = "America/New_York"
    tue = "08:00-12:30,13:30-17:00"
    wed = "08:00-12:30,13:30-17:00"
  }
}
