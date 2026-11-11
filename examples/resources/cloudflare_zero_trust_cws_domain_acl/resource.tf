resource "cloudflare_zero_trust_cws_domain_acl" "example_zero_trust_cws_domain_acl" {
  account_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  default_action = "ALLOW"
  entries = [{
    action = "ALLOW"
    domain = "x"
  }]
  name = "x"
}
