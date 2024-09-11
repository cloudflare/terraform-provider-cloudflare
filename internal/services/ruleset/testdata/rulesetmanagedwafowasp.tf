
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_managed"

  rules = [{
    action = "execute"
    action_parameters = {
      id = "4814384a9e5d4991b9815dcfc25d2f1f"
    }
    expression  = "true"
    description = "Execute Cloudflare Managed OWASP Ruleset on my zone-level phase ruleset"
    enabled     = true
  }]
}