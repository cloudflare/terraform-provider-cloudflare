variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_config_settings"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "set_config"
      action_parameters = {
        autominify = {
          css  = false
          html = false
          js   = false
        }
        bic                      = true
        email_obfuscation        = true
        fonts                    = true
        hotlink_protection       = true
        mirage                   = true
        opportunistic_encryption = true
        rocket_loader            = true
        server_side_excludes     = true
        sxg                      = true
      }
    }
  ]
}
