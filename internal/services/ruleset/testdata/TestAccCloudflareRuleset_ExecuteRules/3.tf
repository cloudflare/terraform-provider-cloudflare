variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_request_firewall_managed"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "execute"
      action_parameters = {
        id = "4814384a9e5d4991b9815dcfc25d2f1f"
        matched_data = {
          public_key = "iGqBmyIUxuWt1rvxoAharN9FUXneUBxA/Y19PyyrEG0="
        }
        overrides = {
          enabled = false
        }
      }
    }
  ]
}
