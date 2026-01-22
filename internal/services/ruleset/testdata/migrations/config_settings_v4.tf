resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "Config Settings Ruleset %[2]s"
  phase   = "http_config_settings"
  kind    = "zone"

  rules {
    expression = "true"
    action     = "set_config"
    description = "Enable automatic HTTPS rewrites"

    action_parameters {
      automatic_https_rewrites = true
      bic                     = true
      disable_zaraz           = true
    }
  }

  rules {
    expression = "http.host eq \"example.com\""
    action     = "set_config"
    description = "Configure security settings"

    action_parameters {
      email_obfuscation     = true
      hotlink_protection    = true
      mirage                = false
      opportunistic_encryption = true
      polish                = "lossless"
      rocket_loader         = false
      security_level        = "medium"
      server_side_excludes  = true
      ssl                   = "flexible"
    }
  }

  rules {
    expression = "http.request.uri.path contains \"/api\""
    action     = "set_config"
    description = "API-specific config"

    action_parameters {
      automatic_https_rewrites = false
      sxg                      = false
    }
  }
}
