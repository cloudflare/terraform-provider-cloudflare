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
        automatic_https_rewrites = true
        autominify               = {}
        bic                      = false
        disable_apps             = true
        disable_rum              = true
        disable_zaraz            = true
        email_obfuscation        = false
        fonts                    = false
        hotlink_protection       = false
        mirage                   = false
        opportunistic_encryption = false
        polish                   = "off"
        rocket_loader            = false
        security_level           = "off"
        server_side_excludes     = false
        ssl                      = "off"
        sxg                      = false
      }
    }
  ]
}
