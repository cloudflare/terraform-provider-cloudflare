variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_request_dynamic_redirect"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "redirect"
      action_parameters = {
        from_value = {
          preserve_query_string = true
          status_code           = 301
          target_url = {
            expression = "concat(\"https://m.example.com\", http.request.uri.path)"
          }
        }
      }
    }
  ]
}
