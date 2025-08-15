
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_transform"

  rules = [{
    action = "rewrite"
    action_parameters = {
      uri = {
        query = {
          value = "a=b"
        }
      }
    }

    expression  = "(http.host eq \"%[4]s\")"
    description = "URI transformation query example"
    enabled     = false
  }]
}