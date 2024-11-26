
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_config_settings"

  rules = [{
    action = "set_config"
    action_parameters = {
      automatic_https_rewrites = true
      autominify = {
        html = true
        css  = true
        js   = true
      }
      bic                      = true
      disable_apps             = true
      disable_zaraz            = true
      disable_rum              = true
      fonts                    = true
      email_obfuscation        = true
      mirage                   = true
      opportunistic_encryption = true
      polish                   = "off"
      rocket_loader            = true
      security_level           = "off"
      server_side_excludes     = true
      ssl                      = "off"
      sxg                      = true
      hotlink_protection       = true
    }
    expression  = "true"
    description = "%[1]s set config rule"
    enabled     = true
  }]
}
