variable "account_id" {}

resource "cloudflare_list" "my_list" {
  account_id = var.account_id
  kind       = "redirect"
  name       = "my_list"
}

resource "cloudflare_ruleset" "my_ruleset" {
  account_id = var.account_id
  name       = "My ruleset"
  phase      = "http_request_redirect"
  kind       = "root"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "redirect"
      action_parameters = {
        from_list = {
          key  = "http.request.full_uri"
          name = cloudflare_list.my_list.name
        }
      }
    }
  ]
}
