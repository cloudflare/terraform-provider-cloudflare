
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "%[1]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_transform"

  rules = [{
    action = "rewrite"
    action_parameters = {
      uri = {
        path = {
          value = "/"
        }
      }
    }
    expression  = "true"
    description = "strip off path"
    enabled     = true
  }]
}