resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "WAF Managed Overrides Ruleset %[2]s"
  phase   = "http_request_firewall_managed"
  kind    = "zone"

  rules {
    expression = "true"
    action     = "execute"
    description = "Execute WAF with category overrides"

    action_parameters {
      id = "efb7b8c949ac4650a09736fc376e9aee"

      overrides {
        enabled = true
        action  = "log"

        categories {
          category = "wordpress"
          action   = "block"
          enabled  = true
        }

        categories {
          category = "joomla"
          enabled  = false
        }
      }
    }
  }

}
