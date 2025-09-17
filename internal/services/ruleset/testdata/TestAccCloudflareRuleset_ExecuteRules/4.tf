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
        overrides = {
          action = "log"
          categories = [
            {
              category = "language-java"
              action   = "block"
            },
            {
              category = "language-php"
              enabled  = false
            },
            {
              category = "language-shell"
              action   = "block"
              enabled  = true
            }
          ]
          enabled = true
          rules = [
            {
              id     = "04116d14d7524986ba314d11c8a41e11"
              action = "block"
            },
            {
              id      = "55b58c71f653446fa0942cf7700f8c8e"
              enabled = false
            },
            {
              id              = "7683285d70b14023ac407b67eccbb280"
              action          = "block"
              enabled         = true
              score_threshold = 40
            }
          ]
        }
      }
    }
  ]
}

data "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  id      = cloudflare_ruleset.my_ruleset.id
}
