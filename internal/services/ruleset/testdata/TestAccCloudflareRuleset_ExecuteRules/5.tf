variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "ddos_l7"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "execute"
      action_parameters = {
        id = "4d21379b4f9f4bb088e0729962c8b3cf"
        overrides = {
          categories = [
            {
              category          = "botnets"
              sensitivity_level = "medium"
            }
          ]
          rules = [
            {
              id                = "8fc7efb08f984ced8d61b34b254da96a"
              sensitivity_level = "low"
            }
          ]
          sensitivity_level = "eoff"
        }
      }
    }
  ]
}

data "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  id      = cloudflare_ruleset.my_ruleset.id
}
