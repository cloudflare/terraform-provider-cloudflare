
  resource "cloudflare_ruleset" "%[1]s" {
    account_id  = "%[2]s"
    name        = "%[1]s managed WAF"
    description = "%[1]s managed WAF ruleset description"
    kind        = "root"
    phase       = "http_request_firewall_managed"

    rules =[ {
      action = "skip"
      action_parameters = {
    rules = {
          "4814384a9e5d4991b9815dcfc25d2f1f" = "a6be45d4905042b9964ff81dc12e41d2,fa54f3d75ed446e78c22b4ea57b90acf,ec42fac3279943388b6be5ee9182835e,37da7855d2f94f69865365d894a556a4,f2db062052cf453fbe9e93f058ecf7e7,1129dfb383bb42e48466488cf3b37cb1"
        }
  }
      expression = "(cf.zone.name eq \"%[3]s\") and (cf.zone.plan eq \"ENT\")"
      description = "Account skip rules OWASP"
      enabled = true
	  logging = {
    enabled = true
  }
    },
    {
    action = "execute"
      action_parameters = {
    id = "4814384a9e5d4991b9815dcfc25d2f1f"
        overrides = { rules =[ {
            id = "6179ae15870a4bb7b2d480d4843b323c"
            action = "block"
            score_threshold = 25
          }]
          enabled = true }
        matched_data = {
    public_key = "zpUlcpNtaNiSUN6LL6NiNz8XgIJZWWG3iSZDdPbMszM="
  }
  }
      expression  = "(cf.zone.name eq \"%[3]s\") and (cf.zone.plan eq \"ENT\")"
      description = "Account OWASP %[3]s"
      enabled     = true
    }]

  }
