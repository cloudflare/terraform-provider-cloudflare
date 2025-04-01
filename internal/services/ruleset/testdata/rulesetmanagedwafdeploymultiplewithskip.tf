
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_managed"

  rules = [{
    action = "skip"
    action_parameters = {
      ruleset = "current"
    }
    description = "not this zone"
    expression  = "(http.host eq \"%[4]s\" and http.request.method eq \"GET\")"
    enabled     = true
    logging = {
      enabled = true
    }
    },
    {
      action = "execute"
      action_parameters = {
        id = "4814384a9e5d4991b9815dcfc25d2f1f"
      }
      expression  = "true"
      description = "zone deployment test"
      enabled     = true
    },
    {
      action = "execute"
      action_parameters = {
        id = "efb7b8c949ac4650a09736fc376e9aee"
      }
      expression  = "true"
      description = "zone deployment test"
      enabled     = true
    },
    {
      action = "execute"
      action_parameters = {
        id = "c2e184081120413c86c3ab7e14069605"
      }
      expression  = "true"
      description = "zone deployment test"
      enabled     = true
  }]



}