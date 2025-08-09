
resource "cloudflare_list" "list-%[1]s" {
  account_id  = "%[2]s"
  name        = "redirect_list_%[1]s"
  description = "%[1]s list description"
  kind        = "redirect"
}

resource "cloudflare_ruleset" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "%[1]s ruleset description"
  kind        = "root"
  phase       = "http_request_redirect"

  rules = [{
    action = "redirect"
    action_parameters = {
      from_list = {
        name = cloudflare_list.list-%[1]s.name
        key  = "http.request.full_uri"
      }
    }
    expression  = "http.request.full_uri in $redirect_list_%[1]s"
    description = "Apply redirects from redirect list"
    enabled     = true
  }]
}
