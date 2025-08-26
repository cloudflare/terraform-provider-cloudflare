# Test configuration to reproduce drift issues reported in GitHub issue #5517
# Based on actual user reports from the issue comments

# Test 1: CNAME with proxied=true and flatten_cname setting
# This reproduces the issue reported by ricohomewood where settings block
# causes drift when proxied=true
resource "cloudflare_dns_record" "%[1]s_proxied_cname_with_settings" {
  zone_id = "%[2]s"
  name    = "api-gateway.%[1]s.%[3]s"  # Using FQDN as recommended
  content = "target.%[3]s"
  type    = "CNAME"
  proxied = true
  ttl     = 1
  
  # This setting should be ignored/cause issues when proxied=true
  # as flatten_cname cannot be set on proxied records
  settings = {
    flatten_cname = false
  }
}

# Test 2: CNAME with mixed case content
# This reproduces the case sensitivity issue reported by RafPe
resource "cloudflare_dns_record" "%[1]s_cname_mixed_case" {
  zone_id = "%[2]s"
  name    = "selector2._domainkey.%[1]s.%[3]s"
  content = "selector2-test._domainkey.MixedCase.onmicrosoft.com"  # Mixed case
  type    = "CNAME"
  proxied = false
  ttl     = 60
}

# Test 3: Record with partial settings
# Based on perlboy's report
resource "cloudflare_dns_record" "%[1]s_cname_partial_settings" {
  zone_id = "%[2]s"
  name    = "_abcd1234.%[1]s.%[3]s"
  content = "example.%[3]s"
  type    = "CNAME"
  proxied = false
  ttl     = 60
  
  settings = {
    flatten_cname = false
    # Not specifying ipv4_only and ipv6_only to test if they cause drift
  }
}

# Test 4: A record with empty tags
# Testing if empty tags array causes drift
resource "cloudflare_dns_record" "%[1]s_a_record_empty_tags" {
  zone_id = "%[2]s"
  name    = "test-empty-tags.%[1]s.%[3]s"
  content = "192.168.0.30"
  type    = "A"
  proxied = false
  ttl     = 3600
  tags    = []
}

# Test 5: Record without any optional fields
# Testing if computed fields cause drift when not specified
resource "cloudflare_dns_record" "%[1]s_minimal_a_record" {
  zone_id = "%[2]s"
  name    = "minimal.%[1]s.%[3]s"
  content = "192.168.0.31"
  type    = "A"
  proxied = false
  ttl     = 3600
  # No tags, no settings, no comment - testing if these cause drift
}