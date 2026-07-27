data "cloudflare_email_routing_rules" "example_email_routing_rules" {
  account_id = "account_id"
  zone_id = "zone_id"
  enabled = true
}
