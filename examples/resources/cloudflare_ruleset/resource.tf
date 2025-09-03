resource "cloudflare_ruleset" "example_ruleset" {
  zone_id     = "9f1839b6152d298aca64c4e906b6d074"
  name        = "My ruleset"
  phase       = "http_request_firewall_custom"
  kind        = "root"
  description = "A description for my ruleset."
  rules = [
    {
      description = "Block the request."
      expression  = "ip.src ne 1.1.1.1"
      action      = "block"
      ref         = "my_rule"
    }
  ]
}
